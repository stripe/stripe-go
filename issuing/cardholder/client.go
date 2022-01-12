//
//
// File generated from our OpenAPI spec
//
//

// Package cardholder provides the /issuing/cardholders APIs
// For more details, see: https://stripe.com/docs/api/?lang=go#issuing_cardholders
package cardholder

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /issuing/cardholders APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new issuing cardholder.
func New(params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	return getC().New(params)
}

// New creates a new issuing cardholder.
func (c Client) New(params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	cardholder := &stripe.IssuingCardholder{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/issuing/cardholders",
		c.Key,
		params,
		cardholder,
	)
	return cardholder, err
}

// Get returns the details of an issuing cardholder.
func Get(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	return getC().Get(id, params)
}

// Get returns the details of an issuing cardholder.
func (c Client) Get(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	path := stripe.FormatURLPath("/v1/issuing/cardholders/%s", id)
	cardholder := &stripe.IssuingCardholder{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cardholder)
	return cardholder, err
}

// Update updates an issuing cardholder's properties.
func Update(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	return getC().Update(id, params)
}

// Update updates an issuing cardholder's properties.
func (c Client) Update(id string, params *stripe.IssuingCardholderParams) (*stripe.IssuingCardholder, error) {
	path := stripe.FormatURLPath("/v1/issuing/cardholders/%s", id)
	cardholder := &stripe.IssuingCardholder{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cardholder)
	return cardholder, err
}

// List returns a list of issuing cardholders.
func List(params *stripe.IssuingCardholderListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuing cardholders.
func (c Client) List(listParams *stripe.IssuingCardholderListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingCardholderList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/cardholders", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing cardholders.
type Iter struct {
	*stripe.Iter
}

// IssuingCardholder returns the issuing cardholder which the iterator is currently pointing to.
func (i *Iter) IssuingCardholder() *stripe.IssuingCardholder {
	return i.Current().(*stripe.IssuingCardholder)
}

// IssuingCardholderList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingCardholderList() *stripe.IssuingCardholderList {
	return i.List().(*stripe.IssuingCardholderList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
