package fetch

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/debdut/textnet/pkg/util"
	readability "github.com/go-shiori/go-readability"
	google "github.com/rocketlaunchr/google-search"
)

type Link struct {
	Text string
	URL  string
}

func GetLinks(body *goquery.Selection) []Link {
	var links []Link
	as := body.Find("a")
	as.Each(func(i int, a *goquery.Selection) {
		link := Link{Text: a.Text()}
		href, exists := a.Attr("href")
		if exists {
			link.URL = href
			links = append(links, link)
		}
	})

	return links
}

func GetImages(body *goquery.Selection) []Link {
	var images []Link
	imgs := body.Find("img")
	imgs.Each(func(i int, img *goquery.Selection) {
		image := Link{Text: img.Text()}
		src, exists := img.Attr("src")
		if exists {
			image.URL = src
			images = append(images, image)
		}
	})

	return images
}

func GetArticleText(buf io.Reader, url *url.URL) (string, error) {
	var text string
	article, err := readability.FromReader(buf, url)
	// couldn't read valid article from document
	if err != nil {
		return text, err
	} else {
		if article.Title != "" {
			text = article.Title
		}
		if article.Byline != "" {
			text += "\n\n" + article.Byline
		}
		if article.TextContent != "" {
			text += "\n\n" + article.TextContent
		}
		if article.Excerpt != "" {
			text += "\n\n" + article.Excerpt
		}
	}

	return text, nil
}

type Site struct {
	URL      string
	FullText string
	Links    []Link
	Images   []Link
}

func GetSite(pageURL string) (Site, error) {
	var site Site
	URLString, url, err := util.GetURL(pageURL)
	if err != nil {
		return site, err
	}

	// get request valiud url
	res, err := http.Get(URLString)
	// couldn't get error
	if err != nil {
		return site, err
	}

	defer res.Body.Close()

	// Use tee so the reader can be used twice
	buf := bytes.NewBuffer(nil)
	tee := io.TeeReader(res.Body, buf)

	// create document reader like jquery
	document, err := goquery.NewDocumentFromReader(tee)
	if err != nil {
		return site, err
	}

	body := document.Find("body")
	article, err := GetArticleText(buf, url)
	if err != nil {
		site.FullText = body.Text()
	} else {
		site.FullText = article
	}

	site.Links = GetLinks(body)
	site.Images = GetImages(body)

	return site, nil
}

// Search
func GetSearchLinks(query string, page uint8) []Link {
	links, err := GetSearchLinksFromFrogFind(query)
	if err != nil {
		return []Link{}
	}

	return links
}

// search on google
func GetSearchLinksFromGoole(query string, page uint8) ([]Link, error) {
	var links []Link
	// TODO: set proxy, page in options
	results, err := google.Search(context.TODO(), query)
	if err != nil {
		return links, err
	}

	for _, result := range results {
		link := Link{URL: result.URL, Text: result.Title}
		links = append(links, link)
	}

	return links, nil
}

// search on frogfind
func GetSearchLinksFromFrogFind(query string) ([]Link, error) {
	var links []Link
	space := regexp.MustCompile(`\s+`)
	fixedQuery := space.ReplaceAllString(query, " ")
	baseURL := "http://frogfind.com/?q="
	searchURL := baseURL + url.QueryEscape(fixedQuery)

	res, err := http.Get(searchURL)
	if err != nil {
		return links, err
	}

	defer res.Body.Close()

	// create document reader like jquery
	document, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return links, err
	}

	body := document.Find("body")
	as := body.Find("a")
	as.Each(func(i int, a *goquery.Selection) {
		link := Link{}
		href, exists := a.Attr("href")
		if exists {
			// check if a search result
			if strings.HasPrefix(href, "/read.php?a=") {
				link.URL = strings.Replace(href, "/read.php?a=", "", 1)
				link.Text = a.Find("b").Text()

				links = append(links, link)
			}
		}
	})

	return links, nil
}
