package event

import (
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /events APIs.
type Client struct {
	B     Backend
	Token string
}

var c *Client

// Get returns the details of an event
// For more details see https://stripe.com/docs/api#retrieve_event.
func Get(id string) (*Event, error) {
	refresh()
	return c.Get(id)
}

func (c *Client) Get(id string) (*Event, error) {
	event := &Event{}
	err := c.B.Call("GET", "/events/"+id, c.Token, nil, event)

	return event, err
}

// List returns a list of events.
// For more details see https://stripe.com/docs/api#list_events
func List(params *EventListParams) (*EventList, error) {
	refresh()
	return c.List(params)
}

func (c *Client) List(params *EventListParams) (*EventList, error) {
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
	err := c.B.Call("GET", "/events", c.Token, body, list)

	return list, err
}

func refresh() {
	if c == nil {
		c = &Client{B: GetBackend()}
	}

	c.Token = Key
}
