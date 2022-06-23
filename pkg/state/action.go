package state

import (
	"strconv"
	"strings"

	"github.com/debdut/textnet/pkg/util"
)

var ActionTypes = []string{
	"Site: Send",
	"Site: Next",
	"Site: Prev",
	"Search: Send",
	// "Search: Next",
	// "Search: Prev",
	"Search: Choose",
	"Link",
	"Link: Next",
	"Link: Prev",
	"Link: Choose",
	"Unknown",
}

type SiteSendPayload struct {
	Query string
}

func (payload *SiteSendPayload) Type() string {
	return "SiteSendPayload"
}

type SiteNextPayload struct{}

func (payload *SiteNextPayload) Type() string {
	return "SiteNextPayload"
}

type SitePrevPayload struct{}

func (payload *SitePrevPayload) Type() string {
	return "SitePrevPayload"
}

type SearchSendPayload struct {
	Query string
}

func (payload *SearchSendPayload) Type() string {
	return "SearchSendPayload"
}

// type SearchNextPayload struct{}

// func (payload *SearchNextPayload) Type() string {
// 	return "SearchNextPayload"
// }

// type SearchPrevPayload struct{}

// func (payload *SearchPrevPayload) Type() string {
// 	return "SearchPrevPayload"
// }

type SearchChoosePayload struct {
	Query uint8
}

func (payload *SearchChoosePayload) Type() string {
	return "SearchChoosePayload"
}

type LinkPayload struct{}

func (payload *LinkPayload) Type() string {
	return "LinkPayload"
}

type LinkNextPayload struct{}

func (payload *LinkNextPayload) Type() string {
	return "LinkNextPayload"
}

type LinkPrevPayload struct{}

func (payload *LinkPrevPayload) Type() string {
	return "LinkPrevPayload"
}

type LinkChoosePayload struct {
	Query uint8
}

func (payload *LinkChoosePayload) Type() string {
	return "LinkChoosePayload"
}

type Unknown struct{}

func (payload *Unknown) Type() string {
	return "Unknown"
}

type Payload interface {
	Type() string
}

type Action struct {
	Type    string
	Payload Payload
}

func IdentifyAction(state State, message string) Action {
	action := Action{
		Type: "Unknown",
	}

	if message == "next" {
		if state.Type == "Site" {
			action.Type = "Site: Next"
			action.Payload = &SiteNextPayload{}

			return action
		}

		if state.Type == "Link" {
			action.Type = "Link: Next"
			action.Payload = &LinkNextPayload{}

			return action
		}
	}

	if message == "prev" {
		if state.Type == "Site" {
			action.Type = "Site: Prev"
			action.Payload = &SitePrevPayload{}

			return action
		}

		if state.Type == "Link" {
			action.Type = "Link: Prev"
			action.Payload = &LinkPrevPayload{}

			return action
		}
	}

	if message == "link" {
		if state.Type == "Site" {
			action.Type = "Link"
			action.Payload = &LinkPayload{}

			return action
		}
	}

	if strings.HasPrefix(message, "search ") {
		action.Type = "Search: Send"
		action.Payload = &SearchSendPayload{
			Query: strings.Replace(message, "search ", "", 1),
		}

		return action
	}

	if index, err := strconv.Atoi(message); err == nil {
		query := uint8(index)

		if state.Type == "Link" {
			action.Type = "Link: Choose"
			action.Payload = &LinkChoosePayload{
				Query: query,
			}

			return action
		}

		if state.Type == "Search" {
			action.Type = "Search: Choose"
			action.Payload = &SearchChoosePayload{
				Query: query,
			}

			return action
		}
	}

	if strings.Contains(message, ".") {
		url, _, err := util.GetURL(message)
		if err == nil {
			action.Type = "Site: Send"
			action.Payload = &SiteSendPayload{
				Query: url,
			}
		}
	}

	return action
}
