//
//
// File generated from our OpenAPI spec
//
//

// Package pricingplan provides the pricingplan related APIs
package pricingplan

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke pricingplan related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Pricing Plan object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingPricingPlanParams) (*stripe.V2BillingPricingPlan, error) {
	pricingplan := &stripe.V2BillingPricingPlan{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/pricing_plans", c.Key, params, pricingplan)
	return pricingplan, err
}

// Retrieve a Pricing Plan object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingPricingPlanParams) (*stripe.V2BillingPricingPlan, error) {
	path := stripe.FormatURLPath("/v2/billing/pricing_plans/%s", id)
	pricingplan := &stripe.V2BillingPricingPlan{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, pricingplan)
	return pricingplan, err
}

// Update a Pricing Plan object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingPricingPlanParams) (*stripe.V2BillingPricingPlan, error) {
	path := stripe.FormatURLPath("/v2/billing/pricing_plans/%s", id)
	pricingplan := &stripe.V2BillingPricingPlan{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pricingplan)
	return pricingplan, err
}

// List all Pricing Plan objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingPricingPlanListParams) stripe.Seq2[*stripe.V2BillingPricingPlan, error] {
	return stripe.NewV2List("/v2/billing/pricing_plans", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingPricingPlan], error) {
		page := &stripe.V2Page[*stripe.V2BillingPricingPlan]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
