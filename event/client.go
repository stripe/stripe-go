// Package event provides the /events APIs
package event

import (
	"net/http"

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
func Get(id string, params *stripe.EventParams) (*stripe.Event, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.EventParams) (*stripe.Event, error) {
	path := stripe.FormatURLPath("/events/%s", id)
	event := &stripe.Event{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, event)
	return event, err
}

// List returns a list of events.
// For more details see https://stripe.com/docs/api#list_events
func List(params *stripe.EventListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.EventListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.EventList{}
		err := c.B.CallRaw(http.MethodGet, "/events", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
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
