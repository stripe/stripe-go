// package event provides the /events APIs
package event

import (
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /events APIs.
type Client struct {
	B   Backend
	Tok string
}

// Get returns the details of an event
// For more details see https://stripe.com/docs/api#retrieve_event.
func Get(id string) (*Event, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*Event, error) {
	event := &Event{}
	err := c.B.Call("GET", "/events/"+id, c.Tok, nil, event)

	return event, err
}

// List returns a list of events.
// For more details see https://stripe.com/docs/api#list_events
func List(params *EventListParams) (*EventList, error) {
	return getC().List(params)
}

func (c Client) List(params *EventListParams) (*EventList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Type) > 0 {
			body.Add("type", params.Type)
		}

		params.AppendTo(body)
	}

	list := &EventList{}
	err := c.B.Call("GET", "/events", c.Tok, body, list)

	return list, err
}

func getC() Client {
	return Client{GetBackend(), Key}
}
