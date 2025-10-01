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

// Retrieve a specific version of a Rate Card object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingRateCardsVersionParams) (*stripe.V2BillingRateCardVersion, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_cards/%s/versions/%s", stripe.StringValue(
			params.RateCardID), id)
	ratecardversion := &stripe.V2BillingRateCardVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, ratecardversion)
	return ratecardversion, err
}

// List the versions of a Rate Card object. Results are sorted in reverse chronological order (most recent first).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingRateCardsVersionListParams) stripe.Seq2[*stripe.V2BillingRateCardVersion, error] {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_cards/%s/versions", stripe.StringValue(
			listParams.RateCardID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingRateCardVersion], error) {
		page := &stripe.V2Page[*stripe.V2BillingRateCardVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
