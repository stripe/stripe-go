//
//
// File generated from our OpenAPI spec
//
//

// Package physicalbundle provides the /issuing/physical_bundles APIs
package physicalbundle

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /issuing/physical_bundles APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an issuing physical bundle.
func Get(id string, params *stripe.IssuingPhysicalBundleParams) (*stripe.IssuingPhysicalBundle, error) {
	return getC().Get(id, params)
}

// Get returns the details of an issuing physical bundle.
func (c Client) Get(id string, params *stripe.IssuingPhysicalBundleParams) (*stripe.IssuingPhysicalBundle, error) {
	path := stripe.FormatURLPath("/v1/issuing/physical_bundles/%s", id)
	physicalbundle := &stripe.IssuingPhysicalBundle{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, physicalbundle)
	return physicalbundle, err
}

// List returns a list of issuing physical bundles.
func List(params *stripe.IssuingPhysicalBundleListParams) *Iter {
	return getC().List(params)
}

// List returns a list of issuing physical bundles.
func (c Client) List(listParams *stripe.IssuingPhysicalBundleListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingPhysicalBundleList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/physical_bundles", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing physical bundles.
type Iter struct {
	*stripe.Iter
}

// IssuingPhysicalBundle returns the issuing physical bundle which the iterator is currently pointing to.
func (i *Iter) IssuingPhysicalBundle() *stripe.IssuingPhysicalBundle {
	return i.Current().(*stripe.IssuingPhysicalBundle)
}

// IssuingPhysicalBundleList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingPhysicalBundleList() *stripe.IssuingPhysicalBundleList {
	return i.List().(*stripe.IssuingPhysicalBundleList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
