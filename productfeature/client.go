//
//
// File generated from our OpenAPI spec
//
//

// Package productfeature provides the /v1/products/{product}/features APIs
package productfeature

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/products/{product}/features APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a product_feature, which represents a feature attachment to a product
func New(params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	return getC().New(params)
}

// Creates a product_feature, which represents a feature attachment to a product
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	path := stripe.FormatURLPath(
		"/v1/products/%s/features", stripe.StringValue(params.Product))
	productfeature := &stripe.ProductFeature{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, productfeature)
	return productfeature, err
}

// Retrieves a product_feature, which represents a feature attachment to a product
func Get(id string, params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	return getC().Get(id, params)
}

// Retrieves a product_feature, which represents a feature attachment to a product
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	path := stripe.FormatURLPath(
		"/v1/products/%s/features/%s", stripe.StringValue(params.Product), id)
	productfeature := &stripe.ProductFeature{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, productfeature)
	return productfeature, err
}

// Deletes the feature attachment to a product
func Del(id string, params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	return getC().Del(id, params)
}

// Deletes the feature attachment to a product
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.ProductFeatureParams) (*stripe.ProductFeature, error) {
	path := stripe.FormatURLPath(
		"/v1/products/%s/features/%s", stripe.StringValue(params.Product), id)
	productfeature := &stripe.ProductFeature{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, productfeature)
	return productfeature, err
}

// Retrieve a list of features for a product
func List(params *stripe.ProductFeatureListParams) *Iter {
	return getC().List(params)
}

// Retrieve a list of features for a product
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ProductFeatureListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/products/%s/features", stripe.StringValue(listParams.Product))
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ProductFeatureList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for product features.
type Iter struct {
	*stripe.Iter
}

// ProductFeature returns the product feature which the iterator is currently pointing to.
func (i *Iter) ProductFeature() *stripe.ProductFeature {
	return i.Current().(*stripe.ProductFeature)
}

// ProductFeatureList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ProductFeatureList() *stripe.ProductFeatureList {
	return i.List().(*stripe.ProductFeatureList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
