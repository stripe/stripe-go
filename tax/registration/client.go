//
//
// File generated from our OpenAPI spec
//
//

// Package registration provides the /tax/registrations APIs
package registration

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /tax/registrations APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new tax registration.
func New(params *stripe.TaxRegistrationParams) (*stripe.TaxRegistration, error) {
	return getC().New(params)
}

// New creates a new tax registration.
func (c Client) New(params *stripe.TaxRegistrationParams) (*stripe.TaxRegistration, error) {
	registration := &stripe.TaxRegistration{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/tax/registrations",
		c.Key,
		params,
		registration,
	)
	return registration, err
}

// Update updates a tax registration's properties.
func Update(id string, params *stripe.TaxRegistrationParams) (*stripe.TaxRegistration, error) {
	return getC().Update(id, params)
}

// Update updates a tax registration's properties.
func (c Client) Update(id string, params *stripe.TaxRegistrationParams) (*stripe.TaxRegistration, error) {
	path := stripe.FormatURLPath("/v1/tax/registrations/%s", id)
	registration := &stripe.TaxRegistration{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, registration)
	return registration, err
}

// List returns a list of tax registrations.
func List(params *stripe.TaxRegistrationListParams) *Iter {
	return getC().List(params)
}

// List returns a list of tax registrations.
func (c Client) List(listParams *stripe.TaxRegistrationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxRegistrationList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/tax/registrations", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for tax registrations.
type Iter struct {
	*stripe.Iter
}

// TaxRegistration returns the tax registration which the iterator is currently pointing to.
func (i *Iter) TaxRegistration() *stripe.TaxRegistration {
	return i.Current().(*stripe.TaxRegistration)
}

// TaxRegistrationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxRegistrationList() *stripe.TaxRegistrationList {
	return i.List().(*stripe.TaxRegistrationList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
