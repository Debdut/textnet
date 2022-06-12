package state

var ActionTypes = []string{
	"Site: Send",
	"Site: Next",
	"Site: Prev",
	"Search: Send",
	"Search: Next",
	"Search: Prev",
	"Search: Choose",
	"Link",
	"Link: Next",
	"Link: Prev",
	"Link: Choose",
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

type SearchNextPayload struct{}

func (payload *SearchNextPayload) Type() string {
	return "SearchNextPayload"
}

type SearchPrevPayload struct{}

func (payload *SearchPrevPayload) Type() string {
	return "SearchPrevPayload"
}

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

type Payload interface {
	Type() string
}

type Action struct {
	Type    string
	Payload Payload
}
