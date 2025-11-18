//
//
// File generated from our OpenAPI spec
//
//

// Package component provides the component related APIs
package component

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke component related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Pricing Plan Component object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingPricingPlansComponentParams) (*stripe.V2BillingPricingPlanComponent, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/pricing_plans/%s/components", stripe.StringValue(
			params.PricingPlanID))
	pricingplancomponent := &stripe.V2BillingPricingPlanComponent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pricingplancomponent)
	return pricingplancomponent, err
}

// Retrieve a Pricing Plan Component object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingPricingPlansComponentParams) (*stripe.V2BillingPricingPlanComponent, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/pricing_plans/%s/components/%s", stripe.StringValue(
			params.PricingPlanID), id)
	pricingplancomponent := &stripe.V2BillingPricingPlanComponent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, pricingplancomponent)
	return pricingplancomponent, err
}

// Update a Pricing Plan Component object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingPricingPlansComponentParams) (*stripe.V2BillingPricingPlanComponent, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/pricing_plans/%s/components/%s", stripe.StringValue(
			params.PricingPlanID), id)
	pricingplancomponent := &stripe.V2BillingPricingPlanComponent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pricingplancomponent)
	return pricingplancomponent, err
}

// Remove a Pricing Plan Component from the latest version of a Pricing Plan.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.V2BillingPricingPlansComponentParams) (*stripe.V2DeletedObject, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/pricing_plans/%s/components/%s", stripe.StringValue(
			params.PricingPlanID), id)
	deletedObj := &stripe.V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}

// List all Pricing Plan Component objects for a Pricing Plan.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingPricingPlansComponentListParams) stripe.Seq2[*stripe.V2BillingPricingPlanComponent, error] {
	path := stripe.FormatURLPath(
		"/v2/billing/pricing_plans/%s/components", stripe.StringValue(
			listParams.PricingPlanID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingPricingPlanComponent], error) {
		page := &stripe.V2Page[*stripe.V2BillingPricingPlanComponent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
