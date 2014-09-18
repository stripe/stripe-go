// package event provides the /events APIs
package event

import (
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /events APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an event
// For more details see https://stripe.com/docs/api#retrieve_event.
func Get(id string) (*stripe.Event, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*stripe.Event, error) {
	event := &stripe.Event{}
	err := c.B.Call("GET", "/events/"+id, c.Key, nil, event)

	return event, err
}

// List returns a list of events.
// For more details see https://stripe.com/docs/api#list_events
func List(params *stripe.EventListParams) *stripe.EventIter {
	return getC().List(params)
}

func (c Client) List(params *stripe.EventListParams) *stripe.EventIter {
	type eventList struct {
		stripe.ListMeta
		Values []*stripe.Event `json:"data"`
	}

	var body *url.Values
	var lp *stripe.ListParams

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if len(params.Type) > 0 {
			body.Add("type", params.Type)
		}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &stripe.EventIter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &eventList{}
		err := c.B.Call("GET", "/events", c.Key, &b, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func getC() Client {
	return Client{stripe.GetBackend(), stripe.Key}
}
