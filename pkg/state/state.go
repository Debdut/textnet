package state

import (
	"github.com/debdut/textnet/pkg/fetch"
)

var StateTypes = []string{
	"Null",
	"Site",
	"Search",
	"Link",
}

type NullState struct{}

func (state *NullState) GetState() string {
	return "Null"
}

type SiteState struct {
	URL     string
	Page    uint8
	MaxPage uint8
	Text    string
	Site    fetch.Site
}

func (state *SiteState) GetState() string {
	return "Site"
}

type SearchState struct {
	Query string
	Links []fetch.Link
	Page  uint8
}

func (state *SearchState) GetState() string {
	return "Search"
}

type LinkState struct {
	URL     string
	Links   []fetch.Link
	Page    uint8
	MaxPage uint8
}

func (state *LinkState) GetState() string {
	return "Link"
}

type SubState interface {
	GetState() string
}

type State struct {
	Type  string
	State SubState
}
