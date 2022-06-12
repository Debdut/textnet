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
func GetText(url string) string {
	return ""
}

type Link struct {
	Text string
	URL  string
}

// TODO
func GetLinks(url string) []Link {
	return []Link{}
}

// TODO
func GetSearchLinks(query string, page uint8) []Link {
	return []Link{}
}
