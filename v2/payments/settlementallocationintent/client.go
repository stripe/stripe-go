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

// Create SettlementAllocationIntent API.
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

// Retrieve SettlementAllocationIntent API.
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

// Update SettlementAllocationIntent API.
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

// Cancel SettlementAllocationIntent API.
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

// Submit SettlementAllocationIntent API.
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
