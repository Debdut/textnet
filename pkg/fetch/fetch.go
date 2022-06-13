package fetch

import (
	"github.com/PuerkitoBio/goquery"
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

type Site struct {
	URL      string
	FullText string
	Links    []Link
	Images   []Link
}

func GetSite(url string) (Site, error) {
	var site Site
	document, err := goquery.NewDocument(url)
	if err != nil {
		return site, err
	}

	body := document.Find("body")
	site.FullText = body.Text()
	site.Links = GetLinks(body)
	site.Images = GetImages(body)

	return site, nil
}

// TODO
func GetSearchLinks(query string, page uint8) []Link {
	return []Link{}
}
