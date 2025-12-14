//
//
// File generated from our OpenAPI spec
//
//

// Package settlementallocationintent provides the settlementallocationintent related APIs
package settlementallocationintent

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke settlementallocationintent related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a new SettlementAllocationIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2PaymentsSettlementAllocationIntentParams) (*stripe.V2PaymentsSettlementAllocationIntent, error) {
	settlementallocationintent := &stripe.V2PaymentsSettlementAllocationIntent{}
	err := c.B.Call(
		http.MethodPost, "/v2/payments/settlement_allocation_intents", c.Key, params, settlementallocationintent)
	return settlementallocationintent, err
}

// Retrieve an existing SettlementAllocationIntent.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2PaymentsSettlementAllocationIntentParams) (*stripe.V2PaymentsSettlementAllocationIntent, error) {
	path := stripe.FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s", id)
	settlementallocationintent := &stripe.V2PaymentsSettlementAllocationIntent{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, settlementallocationintent)
	return settlementallocationintent, err
}

// Updates SettlementAllocationIntent. Only SettlementAllocationIntent with status `pending`, `submitted` and `errored` can be updated. Only amount and reference fields can be updated for a SettlementAllocationIntent and at least one must be present. Updating an `amount` moves the SettlementAllocationIntent `pending` status and updating the `reference` for `errored` SettlementAllocationIntent moves it to `submitted`.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2PaymentsSettlementAllocationIntentParams) (*stripe.V2PaymentsSettlementAllocationIntent, error) {
	path := stripe.FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s", id)
	settlementallocationintent := &stripe.V2PaymentsSettlementAllocationIntent{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, settlementallocationintent)
	return settlementallocationintent, err
}

// Cancels an existing SettlementAllocationIntent. Only SettlementAllocationIntent with status `pending`, `submitted` and `errored` can be `canceled`.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2PaymentsSettlementAllocationIntentCancelParams) (*stripe.V2PaymentsSettlementAllocationIntent, error) {
	path := stripe.FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/cancel", id)
	settlementallocationintent := &stripe.V2PaymentsSettlementAllocationIntent{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, settlementallocationintent)
	return settlementallocationintent, err
}

// Submits a SettlementAllocationIntent. Only SettlementAllocationIntent with status `pending` can be `submitted`. The net sum of SettlementAllocationIntentSplit amount must be equal to SettlementAllocationIntent amount to be eligible to be submitted.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Submit(id string, params *stripe.V2PaymentsSettlementAllocationIntentSubmitParams) (*stripe.V2PaymentsSettlementAllocationIntent, error) {
	path := stripe.FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/submit", id)
	settlementallocationintent := &stripe.V2PaymentsSettlementAllocationIntent{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, settlementallocationintent)
	return settlementallocationintent, err
}
