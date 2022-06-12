package state

type Action struct {
	Type    string
	Payload struct {
		Query string
	}
}

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
