//
//
// File generated from our OpenAPI spec
//
//

// Package hold provides the /v1/reserve/holds APIs
package hold

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/form"
)

// Client is used to invoke /v1/reserve/holds APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a ReserveHold.
func Get(id string, params *stripe.ReserveHoldParams) (*stripe.ReserveHold, error) {
	return getC().Get(id, params)
}

// Retrieve a ReserveHold.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ReserveHoldParams) (*stripe.ReserveHold, error) {
	path := stripe.FormatURLPath("/v1/reserve/holds/%s", id)
	hold := &stripe.ReserveHold{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, hold)
	return hold, err
}

// Returns a list of ReserveHolds previously created. The ReserveHolds are returned in sorted order, with the most recent ReserveHolds appearing first.
func List(params *stripe.ReserveHoldListParams) *Iter {
	return getC().List(params)
}

// Returns a list of ReserveHolds previously created. The ReserveHolds are returned in sorted order, with the most recent ReserveHolds appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ReserveHoldListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ReserveHoldList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/reserve/holds", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for reserve holds.
type Iter struct {
	*stripe.Iter
}

// ReserveHold returns the reserve hold which the iterator is currently pointing to.
func (i *Iter) ReserveHold() *stripe.ReserveHold {
	return i.Current().(*stripe.ReserveHold)
}

// ReserveHoldList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ReserveHoldList() *stripe.ReserveHoldList {
	return i.List().(*stripe.ReserveHoldList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
