//
//
// File generated from our OpenAPI spec
//
//

// Package ratecard provides the ratecard related APIs
package ratecard

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke ratecard related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a RateCard object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingRateCardParams) (*stripe.V2BillingRateCard, error) {
	ratecard := &stripe.V2BillingRateCard{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/rate_cards", c.Key, params, ratecard)
	return ratecard, err
}

// Retrieve the latest version of a RateCard object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingRateCardParams) (*stripe.V2BillingRateCard, error) {
	path := stripe.FormatURLPath("/v2/billing/rate_cards/%s", id)
	ratecard := &stripe.V2BillingRateCard{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, ratecard)
	return ratecard, err
}

// Update a RateCard object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingRateCardParams) (*stripe.V2BillingRateCard, error) {
	path := stripe.FormatURLPath("/v2/billing/rate_cards/%s", id)
	ratecard := &stripe.V2BillingRateCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, ratecard)
	return ratecard, err
}

// List all RateCard objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingRateCardListParams) stripe.Seq2[*stripe.V2BillingRateCard, error] {
	return stripe.NewV2List("/v2/billing/rate_cards", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingRateCard], error) {
		page := &stripe.V2Page[*stripe.V2BillingRateCard]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
