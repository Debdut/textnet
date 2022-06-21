package fetch

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	readability "github.com/go-shiori/go-readability"
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

func getURL(path string) (*url.URL, error) {
	url, err := url.ParseRequestURI(path)
	if err != nil {
		return url, err
	}

	if !strings.HasPrefix(url.Scheme, "http") {
		return url, errors.New("non http/https url")
	}

	return url, err
}

func getArticleText(buf io.Reader, url *url.URL) (string, error) {
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
	url, err := getURL(pageURL)
	if err != nil {
		return site, nil
	}

	// get request valiud url
	res, err := http.Get(pageURL)
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
	article, err := getArticleText(buf, url)
	if err != nil {
		site.FullText = body.Text()
	} else {
		site.FullText = article
	}

	site.Links = GetLinks(body)
	site.Images = GetImages(body)

	return site, nil
}

// TODO
func GetSearchLinks(query string, page uint8) []Link {
	return []Link{}
}
