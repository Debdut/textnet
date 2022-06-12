package state

import (
	"github.com/debdut/textnet/pkg/fetch"
	"github.com/debdut/textnet/pkg/util"
)

func Machine(state State, action Action) State {
	if action.Type == "Site: Send" {

		// fetch site text, compute max pages
		URL := action.Payload.Query
		FullText := fetch.GetText(URL)
		MaxPage := util.GetSiteMaxPages(FullText)
		var Page uint8 = 1
		Text := util.GetSitePage(FullText, Page)

		return State{
			Type: "Site",
			State: &SiteState{
				URL:      URL,
				Page:     Page,
				MaxPage:  MaxPage,
				Text:     Text,
				FullText: FullText,
			},
		}
	}

	if action.Type == "Search: Send" {

		// fetch search results
		Query := action.Payload.Query
		var Page uint8 = 1
		Links := fetch.GetSearchLinks(Query, Page)

		return State{
			Type: "Search",
			State: &SearchState{
				Query: Query,
				Links: Links,
				Page:  Page,
			},
		}
	}

	if state.Type == "Site" {
		prevState := state.State.(*SiteState)
		URL := prevState.URL
		MaxPage := prevState.MaxPage
		FullText := prevState.FullText
		Page := prevState.Page

		if action.Type == "Site: Next" {

			// get next page (part of text) of the site
			if MaxPage > Page {
				Page += 1
			}
			Text := util.GetSitePage(FullText, Page)

			return State{
				Type: "Site",
				State: &SiteState{
					URL:      URL,
					Page:     Page,
					MaxPage:  MaxPage,
					Text:     Text,
					FullText: FullText,
				},
			}
		}

		if action.Type == "Site: Prev" {

			// get prev page (part of text) of the site
			if Page > 1 {
				Page -= 1
			}
			Text := util.GetSitePage(FullText, Page)

			return State{
				Type: "Site",
				State: &SiteState{
					URL:      URL,
					Page:     Page,
					MaxPage:  MaxPage,
					Text:     Text,
					FullText: FullText,
				},
			}
		}

		if action.Type == "Link" {

			// get all links from page
			Links := fetch.GetLinks(URL)
			MaxPage := util.GetLinkMaxPages(Links)

			return State{
				Type: "Link",
				State: &LinkState{
					URL:     URL,
					Links:   Links,
					Page:    1,
					MaxPage: MaxPage,
				},
			}
		}
	}

	if state.Type == "Search" {
		if action.Type == "Search: Next" {

			// TODO

			return State{
				Type: "Search",
				State: &SearchState{
					Query: state.State.(*SearchState).Query,
					Links: []fetch.Link{},
					Page:  state.State.(*SearchState).Page + 1,
				},
			}
		}

		if action.Type == "Search: Prev" {

			// TODO

			return State{
				Type: "Search",
				State: &SearchState{
					Query: state.State.(*SearchState).Query,
					Links: []fetch.Link{},
					Page:  state.State.(*SearchState).Page - 1,
				},
			}
		}

		if action.Type == "Search: Choose" {

			// TODO

			return State{
				Type: "Site",
				State: &SiteState{
					URL:      "",
					Page:     1,
					MaxPage:  1,
					Text:     "",
					FullText: "",
				},
			}
		}
	}

	if state.Type == "Link" {
		if action.Type == "Link: Next" {

			// TODO

			return State{
				Type: "Link",
				State: &LinkState{
					URL:     state.State.(*LinkState).URL,
					Links:   []fetch.Link{},
					Page:    state.State.(*LinkState).Page + 1,
					MaxPage: 1,
				},
			}
		}

		if action.Type == "Link: Prev" {

			// TODO

			return State{
				Type: "Link",
				State: &LinkState{
					URL:     state.State.(*LinkState).URL,
					Links:   []fetch.Link{},
					Page:    state.State.(*LinkState).Page - 1,
					MaxPage: 1,
				},
			}
		}

		if action.Type == "Link: Choose" {

			// TODO

			return State{
				Type: "Site",
				State: &SiteState{
					URL:      "",
					Page:     1,
					MaxPage:  1,
					Text:     "",
					FullText: "",
				},
			}
		}
	}

	return State{Type: "Null"}
}
