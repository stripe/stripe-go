//
//
// File generated from our OpenAPI spec
//
//

// Package split provides the split related APIs
package split

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke split related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create SettlementAllocationIntentSplit API.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2PaymentsSettlementAllocationIntentsSplitParams) (*stripe.V2PaymentsSettlementAllocationIntentSplit, error) {
	path := stripe.FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/splits", stripe.StringValue(
			params.SettlementAllocationIntentID))
	settlementallocationintentsplit := &stripe.V2PaymentsSettlementAllocationIntentSplit{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, settlementallocationintentsplit)
	return settlementallocationintentsplit, err
}

// Retrieve SettlementAllocationIntentSplit API.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2PaymentsSettlementAllocationIntentsSplitParams) (*stripe.V2PaymentsSettlementAllocationIntentSplit, error) {
	path := stripe.FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/splits/%s", stripe.StringValue(
			params.SettlementAllocationIntentID), id)
	settlementallocationintentsplit := &stripe.V2PaymentsSettlementAllocationIntentSplit{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, settlementallocationintentsplit)
	return settlementallocationintentsplit, err
}

// Cancel SettlementAllocationIntentSplit API.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2PaymentsSettlementAllocationIntentsSplitCancelParams) (*stripe.V2PaymentsSettlementAllocationIntentSplit, error) {
	path := stripe.FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/splits/%s/cancel", stripe.StringValue(
			params.SettlementAllocationIntentID), id)
	settlementallocationintentsplit := &stripe.V2PaymentsSettlementAllocationIntentSplit{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, settlementallocationintentsplit)
	return settlementallocationintentsplit, err
}
