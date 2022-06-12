package fetch

import (
	"io"
	"net/http"
)

var BASE string = "http://frogfind.com/read.php?a="

func Get(url string) (string, error) {
	// frogFindURL := fmt.Sprintf(BASE, url)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// TODO
func GetText(html string) string {
	return ""
}

type Link struct {
	Text string
	URL  string
}

// TODO
func GetLinks(html string) []Link {
	return []Link{}
}

// TODO
func GetImages(html string) []Link {
	return []Link{}
}

type Site struct {
	URL      string
	FullText string
	Links    []Link
	Images   []Link
}

// TODO
func GetSite(url string) (Site, error) {
	var site Site

	html, err := Get(url)
	if err != nil {
		return site, err
	}

	site.FullText = GetText(html)
	site.Links = GetLinks(html)
	site.Images = GetImages(html)

	return site, nil
}

// TODO
func GetSearchLinks(query string, page uint8) []Link {
	return []Link{}
}
