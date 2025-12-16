//
//
// File generated from our OpenAPI spec
//
//

// Package plan provides the /v1/reserve/plans APIs
package plan

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
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

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
