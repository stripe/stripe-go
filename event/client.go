// Package event provides the /events APIs
package event

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /events APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an event
// For more details see https://stripe.com/docs/api#retrieve_event.
func Get(id string, params *stripe.Params) (*stripe.Event, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.Params) (*stripe.Event, error) {
	event := &stripe.Event{}
	err := c.B.Call("GET", "/events/"+id, c.Key, nil, params, event)
	return event, err
}

// List returns a list of events.
// For more details see https://stripe.com/docs/api#list_events
func List(params *stripe.EventListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.EventListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.EventList{}
		err := c.B.Call("GET", "/events", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Events.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Event returns the most recent Event
// visited by a call to Next.
func (i *Iter) Event() *stripe.Event {
	return i.Current().(*stripe.Event)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
