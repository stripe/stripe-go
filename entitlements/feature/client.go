//
//
// File generated from our OpenAPI spec
//
//

// Package feature provides the /entitlements/features APIs
package feature

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v77"
	"github.com/stripe/stripe-go/v77/form"
)

// Client is used to invoke /entitlements/features APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new entitlements feature.
func New(params *stripe.EntitlementsFeatureParams) (*stripe.EntitlementsFeature, error) {
	return getC().New(params)
}

// New creates a new entitlements feature.
func (c Client) New(params *stripe.EntitlementsFeatureParams) (*stripe.EntitlementsFeature, error) {
	feature := &stripe.EntitlementsFeature{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/entitlements/features",
		c.Key,
		params,
		feature,
	)
	return feature, err
}

// Get returns the details of an entitlements feature.
func Get(id string, params *stripe.EntitlementsFeatureParams) (*stripe.EntitlementsFeature, error) {
	return getC().Get(id, params)
}

// Get returns the details of an entitlements feature.
func (c Client) Get(id string, params *stripe.EntitlementsFeatureParams) (*stripe.EntitlementsFeature, error) {
	path := stripe.FormatURLPath("/v1/entitlements/features/%s", id)
	feature := &stripe.EntitlementsFeature{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, feature)
	return feature, err
}

// Update updates an entitlements feature's properties.
func Update(id string, params *stripe.EntitlementsFeatureParams) (*stripe.EntitlementsFeature, error) {
	return getC().Update(id, params)
}

// Update updates an entitlements feature's properties.
func (c Client) Update(id string, params *stripe.EntitlementsFeatureParams) (*stripe.EntitlementsFeature, error) {
	path := stripe.FormatURLPath("/v1/entitlements/features/%s", id)
	feature := &stripe.EntitlementsFeature{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, feature)
	return feature, err
}

// List returns a list of entitlements features.
func List(params *stripe.EntitlementsFeatureListParams) *Iter {
	return getC().List(params)
}

// List returns a list of entitlements features.
func (c Client) List(listParams *stripe.EntitlementsFeatureListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.EntitlementsFeatureList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/entitlements/features", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for entitlements features.
type Iter struct {
	*stripe.Iter
}

// EntitlementsFeature returns the entitlements feature which the iterator is currently pointing to.
func (i *Iter) EntitlementsFeature() *stripe.EntitlementsFeature {
	return i.Current().(*stripe.EntitlementsFeature)
}

// EntitlementsFeatureList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) EntitlementsFeatureList() *stripe.EntitlementsFeatureList {
	return i.List().(*stripe.EntitlementsFeatureList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
