//
//
// File generated from our OpenAPI spec
//
//

// Package version provides the version related APIs
package version

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke version related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a specific Pricing Plan Version of a Pricing Plan.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingPricingPlansVersionParams) (*stripe.V2BillingPricingPlanVersion, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/pricing_plans/%s/versions/%s", stripe.StringValue(
			params.PricingPlanID), id)
	pricingplanversion := &stripe.V2BillingPricingPlanVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, pricingplanversion)
	return pricingplanversion, err
}

// List all Pricing Plan Versions of a Pricing Plan.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingPricingPlansVersionListParams) stripe.Seq2[*stripe.V2BillingPricingPlanVersion, error] {
	path := stripe.FormatURLPath(
		"/v2/billing/pricing_plans/%s/versions", stripe.StringValue(
			listParams.PricingPlanID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingPricingPlanVersion], error) {
		page := &stripe.V2Page[*stripe.V2BillingPricingPlanVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
