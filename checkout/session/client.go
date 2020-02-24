// Package session provides API functions related to checkout sessions.
package session

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /checkout_sessions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new session.
func New(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	return getC().New(params)
}

// New creates a new session.
func (c Client) New(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	session := &stripe.CheckoutSession{}
	err := c.B.Call(http.MethodPost, "/v1/checkout/sessions", c.Key, params, session)
	return session, err
}

// Get retrieves a session.
func Get(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	return getC().Get(id, params)
}

// Get retrieves a session.
func (c Client) Get(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	path := stripe.FormatURLPath("/v1/checkout/sessions/%s", id)
	session := &stripe.CheckoutSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, session)
	return session, err
}

// List returns a list of sessions.
func List(params *stripe.CheckoutSessionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of sessions.
func (c Client) List(listParams *stripe.CheckoutSessionListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.CheckoutSessionList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/checkout/sessions", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for sessions.
type Iter struct {
	*stripe.Iter
}

// CheckoutSession returns the session which the iterator is currently pointing to.
func (i *Iter) CheckoutSession() *stripe.CheckoutSession {
	return i.Current().(*stripe.CheckoutSession)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
