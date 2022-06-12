package util

import (
	"github.com/debdut/textnet/pkg/fetch"
)

// TODO
func GetSiteMaxPages(text string) uint8 {
	return 1
}

// TODO
func GetSitePage(text string, page uint8) string {
	return ""
}

// TODO
func GetLinkMaxPages(links []fetch.Link) uint8 {
	return 10
}

// TODO
func GetLinkPage(links []fetch.Link, page uint8) []fetch.Link {
	return []fetch.Link{}
}
