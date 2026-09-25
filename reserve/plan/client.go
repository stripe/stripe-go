//
//
// File generated from our OpenAPI spec
//
//

// Package plan provides the /v1/reserve/plans APIs
package plan

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/reserve/plans APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a ReservePlan.
func Get(id string, params *stripe.ReservePlanParams) (*stripe.ReservePlan, error) {
	return getC().Get(id, params)
}

// Retrieve a ReservePlan.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ReservePlanParams) (*stripe.ReservePlan, error) {
	path := stripe.FormatURLPath("/v1/reserve/plans/%s", id)
	plan := &stripe.ReservePlan{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, plan)
	return plan, err
}

// Returns a list of ReservePlans previously created. The ReservePlans are returned in sorted order, with the most recent ReservePlans appearing first.
func List(params *stripe.ReservePlanListParams) *Iter {
	return getC().List(params)
}

// Returns a list of ReservePlans previously created. The ReservePlans are returned in sorted order, with the most recent ReservePlans appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ReservePlanListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ReservePlanList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/reserve/plans", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for reserve plans.
type Iter struct {
	*stripe.Iter
}

// ReservePlan returns the reserve plan which the iterator is currently pointing to.
func (i *Iter) ReservePlan() *stripe.ReservePlan {
	return i.Current().(*stripe.ReservePlan)
}

// ReservePlanList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ReservePlanList() *stripe.ReservePlanList {
	return i.List().(*stripe.ReservePlanList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
