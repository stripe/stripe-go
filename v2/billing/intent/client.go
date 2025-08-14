//
//
// File generated from our OpenAPI spec
//
//

// Package intent provides the intent related APIs
package intent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke intent related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a BillingIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingIntentParams) (*stripe.V2BillingIntent, error) {
	intent := &stripe.V2BillingIntent{}
	err := c.B.Call(http.MethodPost, "/v2/billing/intents", c.Key, params, intent)
	return intent, err
}

// Retrieve a BillingIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingIntentParams) (*stripe.V2BillingIntent, error) {
	path := stripe.FormatURLPath("/v2/billing/intents/%s", id)
	intent := &stripe.V2BillingIntent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, intent)
	return intent, err
}

// Cancel a BillingIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2BillingIntentCancelParams) (*stripe.V2BillingIntent, error) {
	path := stripe.FormatURLPath("/v2/billing/intents/%s/cancel", id)
	intent := &stripe.V2BillingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, intent)
	return intent, err
}

// Commit a BillingIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Commit(id string, params *stripe.V2BillingIntentCommitParams) (*stripe.V2BillingIntent, error) {
	path := stripe.FormatURLPath("/v2/billing/intents/%s/commit", id)
	intent := &stripe.V2BillingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, intent)
	return intent, err
}

// Release a BillingIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReleaseReservation(id string, params *stripe.V2BillingIntentReleaseReservationParams) (*stripe.V2BillingIntent, error) {
	path := stripe.FormatURLPath("/v2/billing/intents/%s/release_reservation", id)
	intent := &stripe.V2BillingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, intent)
	return intent, err
}

// Reserve a BillingIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Reserve(id string, params *stripe.V2BillingIntentReserveParams) (*stripe.V2BillingIntent, error) {
	path := stripe.FormatURLPath("/v2/billing/intents/%s/reserve", id)
	intent := &stripe.V2BillingIntent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, intent)
	return intent, err
}

// List BillingIntents.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingIntentListParams) stripe.Seq2[*stripe.V2BillingIntent, error] {
	return stripe.NewV2List("/v2/billing/intents", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingIntent], error) {
		page := &stripe.V2Page[*stripe.V2BillingIntent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
