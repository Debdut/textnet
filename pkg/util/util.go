package util

import (
	"net/url"
)

// tries to fix url, returns same if ok
func GetURL(path string) (string, *url.URL, error) {
	URL, err := url.ParseRequestURI(path)
	if err != nil {
		fixedPath := "https://" + path
		URL, err = url.ParseRequestURI(fixedPath)
		if err != nil {
			return path, URL, err
		}

		return fixedPath, URL, nil
	}

	return path, URL, nil
}
