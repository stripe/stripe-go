//
//
// File generated from our OpenAPI spec
//
//

// Package cadence provides the cadence related APIs
package cadence

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke cadence related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Billing Cadence object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingCadenceParams) (*stripe.V2BillingCadence, error) {
	cadence := &stripe.V2BillingCadence{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/cadences", c.Key, params, cadence)
	return cadence, err
}

// Retrieve a Billing Cadence object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingCadenceParams) (*stripe.V2BillingCadence, error) {
	path := stripe.FormatURLPath("/v2/billing/cadences/%s", id)
	cadence := &stripe.V2BillingCadence{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cadence)
	return cadence, err
}

// Update a Billing Cadence object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingCadenceParams) (*stripe.V2BillingCadence, error) {
	path := stripe.FormatURLPath("/v2/billing/cadences/%s", id)
	cadence := &stripe.V2BillingCadence{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cadence)
	return cadence, err
}

// Cancel the Billing Cadence.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2BillingCadenceCancelParams) (*stripe.V2BillingCadence, error) {
	path := stripe.FormatURLPath("/v2/billing/cadences/%s/cancel", id)
	cadence := &stripe.V2BillingCadence{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cadence)
	return cadence, err
}

// List Billing Cadences.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingCadenceListParams) stripe.Seq2[*stripe.V2BillingCadence, error] {
	return stripe.NewV2List("/v2/billing/cadences", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingCadence], error) {
		page := &stripe.V2Page[*stripe.V2BillingCadence]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
