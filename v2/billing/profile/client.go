//
//
// File generated from our OpenAPI spec
//
//

// Package profile provides the profile related APIs
package profile

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke profile related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a BillingProfile object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingProfileParams) (*stripe.V2BillingProfile, error) {
	profile := &stripe.V2BillingProfile{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/profiles", c.Key, params, profile)
	return profile, err
}

// Retrieve a BillingProfile object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingProfileParams) (*stripe.V2BillingProfile, error) {
	path := stripe.FormatURLPath("/v2/billing/profiles/%s", id)
	profile := &stripe.V2BillingProfile{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, profile)
	return profile, err
}

// Update a BillingProfile object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingProfileParams) (*stripe.V2BillingProfile, error) {
	path := stripe.FormatURLPath("/v2/billing/profiles/%s", id)
	profile := &stripe.V2BillingProfile{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, profile)
	return profile, err
}

// List Billing Profiles.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingProfileListParams) stripe.Seq2[*stripe.V2BillingProfile, error] {
	return stripe.NewV2List("/v2/billing/profiles", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingProfile], error) {
		page := &stripe.V2Page[*stripe.V2BillingProfile]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
