//
//
// File generated from our OpenAPI spec
//
//

// Package product provides the /v1/climate/products APIs
package product

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/climate/products APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Climate product with the given ID.
func Get(id string, params *stripe.ClimateProductParams) (*stripe.ClimateProduct, error) {
	return getC().Get(id, params)
}

// Retrieves the details of a Climate product with the given ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ClimateProductParams) (*stripe.ClimateProduct, error) {
	path := stripe.FormatURLPath("/v1/climate/products/%s", id)
	product := &stripe.ClimateProduct{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, product)
	return product, err
}

// Lists all available Climate product objects.
func List(params *stripe.ClimateProductListParams) *Iter {
	return getC().List(params)
}

// Lists all available Climate product objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ClimateProductListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ClimateProductList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/climate/products", c.Key, []byte(b.Encode()), p, list)

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
