//
//
// File generated from our OpenAPI spec
//
//

// Package sku provides the /skus APIs
package sku

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /skus APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new sku.
func New(params *stripe.SKUParams) (*stripe.SKU, error) {
	return getC().New(params)
}

// New creates a new sku.
func (c Client) New(params *stripe.SKUParams) (*stripe.SKU, error) {
	sku := &stripe.SKU{}
	err := c.B.Call(http.MethodPost, "/v1/skus", c.Key, params, sku)
	return sku, err
}

// Get returns the details of a sku.
func Get(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	return getC().Get(id, params)
}

// Get returns the details of a sku.
func (c Client) Get(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	path := stripe.FormatURLPath("/v1/skus/%s", id)
	sku := &stripe.SKU{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, sku)
	return sku, err
}

// Update updates a sku's properties.
func Update(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	return getC().Update(id, params)
}

// Update updates a sku's properties.
func (c Client) Update(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	path := stripe.FormatURLPath("/v1/skus/%s", id)
	sku := &stripe.SKU{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, sku)
	return sku, err
}

// Del removes a sku.
func Del(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	return getC().Del(id, params)
}

// Del removes a sku.
func (c Client) Del(id string, params *stripe.SKUParams) (*stripe.SKU, error) {
	path := stripe.FormatURLPath("/v1/skus/%s", id)
	sku := &stripe.SKU{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, sku)
	return sku, err
}

// List returns a list of skus.
func List(params *stripe.SKUListParams) *Iter {
	return getC().List(params)
}

// List returns a list of skus.
func (c Client) List(listParams *stripe.SKUListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SKUList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/skus", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for skus.
type Iter struct {
	*stripe.Iter
}

// SKU returns the sku which the iterator is currently pointing to.
func (i *Iter) SKU() *stripe.SKU {
	return i.Current().(*stripe.SKU)
}

// SKUList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SKUList() *stripe.SKUList {
	return i.List().(*stripe.SKUList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
