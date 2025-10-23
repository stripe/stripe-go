//
//
// File generated from our OpenAPI spec
//
//

// Package component provides the component related APIs
package component

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke component related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a Pricing Plan Subscription's components.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingPricingPlanSubscriptionsComponentParams) (*stripe.V2BillingPricingPlanSubscriptionComponents, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/pricing_plan_subscriptions/%s/components", id)
	pricingplansubscriptioncomponents := &stripe.V2BillingPricingPlanSubscriptionComponents{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, pricingplansubscriptioncomponents)
	return pricingplansubscriptioncomponents, err
}
