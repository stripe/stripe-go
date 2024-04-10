//
//
// File generated from our OpenAPI spec
//
//

// Package product provides the /climate/products APIs
package product

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/form"
)

// Client is used to invoke /climate/products APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a climate product.
func Get(id string, params *stripe.ClimateProductParams) (*stripe.ClimateProduct, error) {
	return getC().Get(id, params)
}

// Get returns the details of a climate product.
func (c Client) Get(id string, params *stripe.ClimateProductParams) (*stripe.ClimateProduct, error) {
	path := stripe.FormatURLPath("/v1/climate/products/%s", id)
	product := &stripe.ClimateProduct{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, product)
	return product, err
}

// List returns a list of climate products.
func List(params *stripe.ClimateProductListParams) *Iter {
	return getC().List(params)
}

// List returns a list of climate products.
func (c Client) List(listParams *stripe.ClimateProductListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ClimateProductList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/climate/products", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for climate products.
type Iter struct {
	*stripe.Iter
}

// ClimateProduct returns the climate product which the iterator is currently pointing to.
func (i *Iter) ClimateProduct() *stripe.ClimateProduct {
	return i.Current().(*stripe.ClimateProduct)
}

// ClimateProductList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ClimateProductList() *stripe.ClimateProductList {
	return i.List().(*stripe.ClimateProductList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
