//
//
// File generated from our OpenAPI spec
//
//

// Package supplier provides the /climate/suppliers APIs
package supplier

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /climate/suppliers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a climate supplier.
func Get(id string, params *stripe.ClimateSupplierParams) (*stripe.ClimateSupplier, error) {
	return getC().Get(id, params)
}

// Get returns the details of a climate supplier.
func (c Client) Get(id string, params *stripe.ClimateSupplierParams) (*stripe.ClimateSupplier, error) {
	path := stripe.FormatURLPath("/v1/climate/suppliers/%s", id)
	supplier := &stripe.ClimateSupplier{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodGet, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, supplier)
	return supplier, err
}

// List returns a list of climate suppliers.
func List(params *stripe.ClimateSupplierListParams) *Iter {
	return getC().List(params)
}

// List returns a list of climate suppliers.
func (c Client) List(listParams *stripe.ClimateSupplierListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ClimateSupplierList{}
			sr := stripe.StripeRequest{
				Method: http.MethodGet,
				Path:   "/v1/climate/suppliers",
				Key:    c.Key,
			}
			err := sr.SetRawForm(p, b)
			if err != nil {
				return nil, list, err
			}
			err = c.B.Call(sr, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for climate suppliers.
type Iter struct {
	*stripe.Iter
}

// ClimateSupplier returns the climate supplier which the iterator is currently pointing to.
func (i *Iter) ClimateSupplier() *stripe.ClimateSupplier {
	return i.Current().(*stripe.ClimateSupplier)
}

// ClimateSupplierList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ClimateSupplierList() *stripe.ClimateSupplierList {
	return i.List().(*stripe.ClimateSupplierList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
