//
//
// File generated from our OpenAPI spec
//
//

// Package release provides the /v1/reserve/releases APIs
package release

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/form"
)

// Client is used to invoke /v1/reserve/releases APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a ReserveRelease.
func Get(id string, params *stripe.ReserveReleaseParams) (*stripe.ReserveRelease, error) {
	return getC().Get(id, params)
}

// Retrieve a ReserveRelease.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ReserveReleaseParams) (*stripe.ReserveRelease, error) {
	path := stripe.FormatURLPath("/v1/reserve/releases/%s", id)
	release := &stripe.ReserveRelease{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, release)
	return release, err
}

// Returns a list of ReserveReleases previously created. The ReserveReleases are returned in sorted order, with the most recent ReserveReleases appearing first.
func List(params *stripe.ReserveReleaseListParams) *Iter {
	return getC().List(params)
}

// Returns a list of ReserveReleases previously created. The ReserveReleases are returned in sorted order, with the most recent ReserveReleases appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ReserveReleaseListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ReserveReleaseList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/reserve/releases", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for reserve releases.
type Iter struct {
	*stripe.Iter
}

// ReserveRelease returns the reserve release which the iterator is currently pointing to.
func (i *Iter) ReserveRelease() *stripe.ReserveRelease {
	return i.Current().(*stripe.ReserveRelease)
}

// ReserveReleaseList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ReserveReleaseList() *stripe.ReserveReleaseList {
	return i.List().(*stripe.ReserveReleaseList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
