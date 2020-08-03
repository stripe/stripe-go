// Package event provides the /events APIs
package event

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v71"
	"github.com/stripe/stripe-go/v71/form"
)

// Client is used to invoke /events APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an event.
func Get(id string, params *stripe.EventParams) (*stripe.Event, error) {
	return getC().Get(id, params)
}

// Get returns the details of an event.
func (c Client) Get(id string, params *stripe.EventParams) (*stripe.Event, error) {
	path := stripe.FormatURLPath("/v1/events/%s", id)
	event := &stripe.Event{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, event)
	return event, err
}

// List returns a list of events.
func List(params *stripe.EventListParams) *Iter {
	return getC().List(params)
}

// List returns a list of events.
func (c Client) List(listParams *stripe.EventListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.EventList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/events", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for events.
type Iter struct {
	*stripe.Iter
}

// Event returns the event which the iterator is currently pointing to.
func (i *Iter) Event() *stripe.Event {
	return i.Current().(*stripe.Event)
}

// EventList returns the current list object which the iterator is currently
// using. List objects will change as new API calls are made to continue
// pagination.
func (i *Iter) EventList() *stripe.EventList {
	return i.List().(*stripe.EventList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
