//
//
// File generated from our OpenAPI spec
//
//

// Package rate provides the rate related APIs
package rate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke rate related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Set the Rate for a Metered Item on the latest version of a Rate Card object. This will create a new Rate Card version
// if the Metered Item already has a rate on the Rate Card.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingRateCardsRateParams) (*stripe.V2BillingRateCardRate, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_cards/%s/rates", stripe.StringValue(params.RateCardID))
	ratecardrate := &stripe.V2BillingRateCardRate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, ratecardrate)
	return ratecardrate, err
}

// Retrieve a Rate object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingRateCardsRateParams) (*stripe.V2BillingRateCardRate, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_cards/%s/rates/%s", stripe.StringValue(
			params.RateCardID), id)
	ratecardrate := &stripe.V2BillingRateCardRate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, ratecardrate)
	return ratecardrate, err
}

// Remove an existing Rate from a Rate Card. This will create a new Rate Card Version without that Rate.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.V2BillingRateCardsRateParams) (*stripe.V2DeletedObject, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_cards/%s/rates/%s", stripe.StringValue(
			params.RateCardID), id)
	deletedObj := &stripe.V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}

// List all Rates associated with a Rate Card for a specific version (defaults to latest). Rates remain active for all subsequent versions until a new rate is created for the same Metered Item.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingRateCardsRateListParams) stripe.Seq2[*stripe.V2BillingRateCardRate, error] {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_cards/%s/rates", stripe.StringValue(
			listParams.RateCardID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingRateCardRate], error) {
		page := &stripe.V2Page[*stripe.V2BillingRateCardRate]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
