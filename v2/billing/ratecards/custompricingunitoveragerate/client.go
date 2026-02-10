//
//
// File generated from our OpenAPI spec
//
//

// Package custompricingunitoveragerate provides the custompricingunitoveragerate related APIs
package custompricingunitoveragerate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke custompricingunitoveragerate related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Rate Card Custom Pricing Unit Overage Rate on a Rate Card.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingRateCardsCustomPricingUnitOverageRateParams) (*stripe.V2BillingRateCardCustomPricingUnitOverageRate, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_cards/%s/custom_pricing_unit_overage_rates", stripe.StringValue(
			params.RateCardID))
	ratecardcustompricingunitoveragerate := &stripe.V2BillingRateCardCustomPricingUnitOverageRate{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, ratecardcustompricingunitoveragerate)
	return ratecardcustompricingunitoveragerate, err
}

// Retrieve a Rate Card Custom Pricing Unit Overage Rate from a Rate Card.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingRateCardsCustomPricingUnitOverageRateParams) (*stripe.V2BillingRateCardCustomPricingUnitOverageRate, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_cards/%s/custom_pricing_unit_overage_rates/%s", stripe.StringValue(
			params.RateCardID), id)
	ratecardcustompricingunitoveragerate := &stripe.V2BillingRateCardCustomPricingUnitOverageRate{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, ratecardcustompricingunitoveragerate)
	return ratecardcustompricingunitoveragerate, err
}

// Delete a Rate Card Custom Pricing Unit Overage Rate from a Rate Card.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.V2BillingRateCardsCustomPricingUnitOverageRateParams) (*stripe.V2DeletedObject, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_cards/%s/custom_pricing_unit_overage_rates/%s", stripe.StringValue(
			params.RateCardID), id)
	deletedObj := &stripe.V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}

// List all Rate Card Custom Pricing Unit Overage Rates on a Rate Card.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingRateCardsCustomPricingUnitOverageRateListParams) stripe.Seq2[*stripe.V2BillingRateCardCustomPricingUnitOverageRate, error] {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_cards/%s/custom_pricing_unit_overage_rates", stripe.StringValue(
			listParams.RateCardID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingRateCardCustomPricingUnitOverageRate], error) {
		page := &stripe.V2Page[*stripe.V2BillingRateCardCustomPricingUnitOverageRate]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
