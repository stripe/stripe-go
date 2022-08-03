//
//
// File generated from our OpenAPI spec
//
//

// Package session provides the /checkout/sessions APIs
package session

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
	"github.com/stripe/stripe-go/v72/lineitem"
)

// Client is used to invoke /checkout/sessions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new checkout session.
func New(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	return getC().New(params)
}

// New creates a new checkout session.
func (c Client) New(params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	session := &stripe.CheckoutSession{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/checkout/sessions",
		c.Key,
		params,
		session,
	)
	return session, err
}

// Get returns the details of a checkout session.
func Get(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	return getC().Get(id, params)
}

// Get returns the details of a checkout session.
func (c Client) Get(id string, params *stripe.CheckoutSessionParams) (*stripe.CheckoutSession, error) {
	path := stripe.FormatURLPath("/v1/checkout/sessions/%s", id)
	session := &stripe.CheckoutSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, session)
	return session, err
}

// Expire is the method for the `POST /v1/checkout/sessions/{session}/expire` API.
func Expire(id string, params *stripe.CheckoutSessionExpireParams) (*stripe.CheckoutSession, error) {
	return getC().Expire(id, params)
}

// Expire is the method for the `POST /v1/checkout/sessions/{session}/expire` API.
func (c Client) Expire(id string, params *stripe.CheckoutSessionExpireParams) (*stripe.CheckoutSession, error) {
	path := stripe.FormatURLPath("/v1/checkout/sessions/%s/expire", id)
	session := &stripe.CheckoutSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, session)
	return session, err
}

// List returns a list of checkout sessions.
func List(params *stripe.CheckoutSessionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of checkout sessions.
func (c Client) List(listParams *stripe.CheckoutSessionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CheckoutSessionList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/checkout/sessions", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for checkout sessions.
type Iter struct {
	*stripe.Iter
}

// CheckoutSession returns the checkout session which the iterator is currently pointing to.
func (i *Iter) CheckoutSession() *stripe.CheckoutSession {
	return i.Current().(*stripe.CheckoutSession)
}

// CheckoutSessionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CheckoutSessionList() *stripe.CheckoutSessionList {
	return i.List().(*stripe.CheckoutSessionList)
}

// ListLineItems is the method for the `GET /v1/checkout/sessions/{session}/line_items` API.
func ListLineItems(id string, params *stripe.CheckoutSessionListLineItemsParams) *LineItemIter {
	return getC().ListLineItems(id, params)
}

// ListLineItems is the method for the `GET /v1/checkout/sessions/{session}/line_items` API.
func (c Client) ListLineItems(id string, listParams *stripe.CheckoutSessionListLineItemsParams) *LineItemIter {
	path := stripe.FormatURLPath("/v1/checkout/sessions/%s/line_items", id)
	return &LineItemIter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.LineItemList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

type LineItemIter = lineitem.Iter

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
