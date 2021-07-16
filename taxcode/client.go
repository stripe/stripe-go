//
//
// File generated from our OpenAPI spec
//
//

// Package taxcode provides the /tax_codes APIs
package taxcode

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /tax_codes APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a tax code.
func Get(id string, params *stripe.TaxCodeParams) (*stripe.TaxCode, error) {
	return getC().Get(id, params)
}

// Get returns the details of a tax code.
func (c Client) Get(id string, params *stripe.TaxCodeParams) (*stripe.TaxCode, error) {
	path := stripe.FormatURLPath("/v1/tax_codes/%s", id)
	taxcode := &stripe.TaxCode{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, taxcode)
	return taxcode, err
}

// List returns a list of tax codes.
func List(params *stripe.TaxCodeListParams) *Iter {
	return getC().List(params)
}

// List returns a list of tax codes.
func (c Client) List(listParams *stripe.TaxCodeListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxCodeList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/tax_codes", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for tax codes.
type Iter struct {
	*stripe.Iter
}

// TaxCode returns the tax code which the iterator is currently pointing to.
func (i *Iter) TaxCode() *stripe.TaxCode {
	return i.Current().(*stripe.TaxCode)
}

// TaxCodeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxCodeList() *stripe.TaxCodeList {
	return i.List().(*stripe.TaxCodeList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
