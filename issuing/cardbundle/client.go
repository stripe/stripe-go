//
//
// File generated from our OpenAPI spec
//
//

// Package cardbundle provides the /issuing/card_bundles APIs
package cardbundle

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/form"
)

// Client is used to invoke /issuing/card_bundles APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an issuing card bundle.
func Get(id string, params *stripe.IssuingCardBundleParams) (*stripe.IssuingCardBundle, error) {
	return getC().Get(id, params)
}

// Get returns the details of an issuing card bundle.
func (c Client) Get(id string, params *stripe.IssuingCardBundleParams) (*stripe.IssuingCardBundle, error) {
	path := stripe.FormatURLPath("/v1/issuing/card_bundles/%s", id)
	cardbundle := &stripe.IssuingCardBundle{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cardbundle)
	return cardbundle, err
}

// List returns a list of issuing card bundles.
func List(params *stripe.IssuingCardBundleListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuing card bundles.
func (c Client) List(listParams *stripe.IssuingCardBundleListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingCardBundleList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/card_bundles", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing card bundles.
type Iter struct {
	*stripe.Iter
}

// IssuingCardBundle returns the issuing card bundle which the iterator is currently pointing to.
func (i *Iter) IssuingCardBundle() *stripe.IssuingCardBundle {
	return i.Current().(*stripe.IssuingCardBundle)
}

// IssuingCardBundleList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingCardBundleList() *stripe.IssuingCardBundleList {
	return i.List().(*stripe.IssuingCardBundleList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
