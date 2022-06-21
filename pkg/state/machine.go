package state

import (
	"errors"

	"github.com/debdut/textnet/pkg/fetch"
	"github.com/debdut/textnet/pkg/util"
)

func Machine(prevState State, action Action) (State, error) {
	if action.Type == "Site: Send" {

		// fetch site text, compute max pages
		URL := action.Payload.(*SiteSendPayload).Query
		Site, err := fetch.GetSite(URL)
		if err != nil {
			return prevState, err
		}

		MaxPage := util.GetSiteMaxPages(Site.FullText)
		var Page uint8 = 1
		Text := util.GetSitePage(Site.FullText, Page)

		return State{
			Type: "Site",
			State: &SiteState{
				URL:     URL,
				Page:    Page,
				MaxPage: MaxPage,
				Text:    Text,
				Site:    Site,
			},
		}, nil
	}

	if action.Type == "Search: Send" {

		// fetch search results
		Query := action.Payload.(*SearchSendPayload).Query
		var Page uint8 = 1
		Links := fetch.GetSearchLinks(Query, Page)

		return State{
			Type: "Search",
			State: &SearchState{
				Query: Query,
				Links: Links,
				Page:  Page,
			},
		}, nil
	}

	if prevState.Type == "Site" {
		siteState := prevState.State.(*SiteState)
		URL := siteState.URL
		MaxPage := siteState.MaxPage
		Site := siteState.Site
		Page := siteState.Page

		if action.Type == "Site: Next" {

			// get next page (part of text) of the site
			if MaxPage > Page {
				Page += 1
			} else {
				return prevState, errors.New("SITE:END")
			}

			Text := util.GetSitePage(Site.FullText, Page)

			return State{
				Type: "Site",
				State: &SiteState{
					URL:     URL,
					Page:    Page,
					MaxPage: MaxPage,
					Text:    Text,
					Site:    Site,
				},
			}, nil
		}

		if action.Type == "Site: Prev" {

			// get prev page (part of text) of the site
			if Page > 1 {
				Page -= 1
			} else {
				return prevState, errors.New("SITE:START")
			}

			Text := util.GetSitePage(Site.FullText, Page)

			return State{
				Type: "Site",
				State: &SiteState{
					URL:     URL,
					Page:    Page,
					MaxPage: MaxPage,
					Text:    Text,
					Site:    Site,
				},
			}, nil
		}

		if action.Type == "Link" {

			// get all links from page
			Links := Site.Links
			MaxPage := util.GetLinkMaxPages(Links)

			return State{
				Type: "Link",
				State: &LinkState{
					URL:     URL,
					Links:   Links,
					Page:    1,
					MaxPage: MaxPage,
				},
			}, nil
		}
	}

	if prevState.Type == "Search" {
		searchState := prevState.State.(*SearchState)
		Query := searchState.Query
		Page := searchState.Page
		Links := searchState.Links

		if action.Type == "Search: Next" {

			// fetch search results
			Page += 1
			Links := fetch.GetSearchLinks(Query, Page)

			if len(Links) == 0 {
				return prevState, errors.New("SEARCH:END")
			}

			return State{
				Type: "Search",
				State: &SearchState{
					Query: Query,
					Links: Links,
					Page:  Page,
				},
			}, nil
		}

		if action.Type == "Search: Prev" {

			// TODO
			if Page == 1 {
				return prevState, errors.New("SEARCH:START")
			}
			Page -= 1
			Links := fetch.GetSearchLinks(Query, Page)

			return State{
				Type: "Search",
				State: &SearchState{
					Query: Query,
					Links: Links,
					Page:  Page,
				},
			}, nil
		}

		if action.Type == "Search: Choose" {

			// TODO
			Query := action.Payload.(*SearchChoosePayload).Query
			if Query > 10 {
				return prevState, errors.New("SEARCH:OUT_OF_BOUND")
			}

			index := 10*(Page-1) + Query - 1
			URL := Links[index].URL
			Site, err := fetch.GetSite(URL)
			if err != nil {
				return prevState, err
			}

			MaxPage := util.GetSiteMaxPages(Site.FullText)
			var Page uint8 = 1
			Text := util.GetSitePage(Site.FullText, Page)

			return State{
				Type: "Site",
				State: &SiteState{
					URL:     URL,
					Page:    Page,
					MaxPage: MaxPage,
					Text:    Text,
					Site:    Site,
				},
			}, nil
		}
	}

	if prevState.Type == "Link" {
		if action.Type == "Link: Next" {

			// TODO

			return State{
				Type: "Link",
				State: &LinkState{
					URL:     prevState.State.(*LinkState).URL,
					Links:   []fetch.Link{},
					Page:    prevState.State.(*LinkState).Page + 1,
					MaxPage: 1,
				},
			}, nil
		}

		if action.Type == "Link: Prev" {

			// TODO

			return State{
				Type: "Link",
				State: &LinkState{
					URL:     prevState.State.(*LinkState).URL,
					Links:   []fetch.Link{},
					Page:    prevState.State.(*LinkState).Page - 1,
					MaxPage: 1,
				},
			}, nil
		}

		if action.Type == "Link: Choose" {

			// TODO

			return State{
				Type: "Site",
				State: &SiteState{
					URL:     "",
					Page:    1,
					MaxPage: 1,
					Text:    "",
					Site:    fetch.Site{},
				},
			}, nil
		}
	}

	return prevState, errors.New("ACTION:INVALID")
}
