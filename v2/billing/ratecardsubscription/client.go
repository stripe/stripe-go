//
//
// File generated from our OpenAPI spec
//
//

// Package ratecardsubscription provides the ratecardsubscription related APIs
package ratecardsubscription

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke ratecardsubscription related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Rate Card Subscription to bill a Rate Card on a specified Billing Cadence.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingRateCardSubscriptionParams) (*stripe.V2BillingRateCardSubscription, error) {
	ratecardsubscription := &stripe.V2BillingRateCardSubscription{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/rate_card_subscriptions", c.Key, params, ratecardsubscription)
	return ratecardsubscription, err
}

// Retrieve a Rate Card Subscription by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingRateCardSubscriptionParams) (*stripe.V2BillingRateCardSubscription, error) {
	path := stripe.FormatURLPath("/v2/billing/rate_card_subscriptions/%s", id)
	ratecardsubscription := &stripe.V2BillingRateCardSubscription{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, ratecardsubscription)
	return ratecardsubscription, err
}

// Update fields on an existing, active Rate Card Subscription.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingRateCardSubscriptionParams) (*stripe.V2BillingRateCardSubscription, error) {
	path := stripe.FormatURLPath("/v2/billing/rate_card_subscriptions/%s", id)
	ratecardsubscription := &stripe.V2BillingRateCardSubscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, ratecardsubscription)
	return ratecardsubscription, err
}

// Cancel an existing, active Rate Card Subscription.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2BillingRateCardSubscriptionCancelParams) (*stripe.V2BillingRateCardSubscription, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/rate_card_subscriptions/%s/cancel", id)
	ratecardsubscription := &stripe.V2BillingRateCardSubscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, ratecardsubscription)
	return ratecardsubscription, err
}

// List all Rate Card Subscription objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingRateCardSubscriptionListParams) stripe.Seq2[*stripe.V2BillingRateCardSubscription, error] {
	return stripe.NewV2List("/v2/billing/rate_card_subscriptions", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingRateCardSubscription], error) {
		page := &stripe.V2Page[*stripe.V2BillingRateCardSubscription]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
