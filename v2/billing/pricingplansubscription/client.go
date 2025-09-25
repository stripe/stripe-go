//
//
// File generated from our OpenAPI spec
//
//

// Package pricingplansubscription provides the pricingplansubscription related APIs
package pricingplansubscription

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke pricingplansubscription related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a Pricing Plan Subscription object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingPricingPlanSubscriptionParams) (*stripe.V2BillingPricingPlanSubscription, error) {
	path := stripe.FormatURLPath("/v2/billing/pricing_plan_subscriptions/%s", id)
	pricingplansubscription := &stripe.V2BillingPricingPlanSubscription{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, pricingplansubscription)
	return pricingplansubscription, err
}

// Update a Pricing Plan Subscription object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingPricingPlanSubscriptionParams) (*stripe.V2BillingPricingPlanSubscription, error) {
	path := stripe.FormatURLPath("/v2/billing/pricing_plan_subscriptions/%s", id)
	pricingplansubscription := &stripe.V2BillingPricingPlanSubscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pricingplansubscription)
	return pricingplansubscription, err
}

// List all Pricing Plan Subscription objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingPricingPlanSubscriptionListParams) stripe.Seq2[*stripe.V2BillingPricingPlanSubscription, error] {
	return stripe.NewV2List("/v2/billing/pricing_plan_subscriptions", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingPricingPlanSubscription], error) {
		page := &stripe.V2Page[*stripe.V2BillingPricingPlanSubscription]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
