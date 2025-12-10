//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2PaymentsSettlementAllocationIntentService is used to invoke settlementallocationintent related APIs.
type v2PaymentsSettlementAllocationIntentService struct {
	B   Backend
	Key string
}

// Create SettlementAllocationIntent API.
func (c v2PaymentsSettlementAllocationIntentService) Create(ctx context.Context, params *V2PaymentsSettlementAllocationIntentCreateParams) (*V2PaymentsSettlementAllocationIntent, error) {
	if params == nil {
		params = &V2PaymentsSettlementAllocationIntentCreateParams{}
	}
	params.Context = ctx
	settlementallocationintent := &V2PaymentsSettlementAllocationIntent{}
	err := c.B.Call(
		http.MethodPost, "/v2/payments/settlement_allocation_intents", c.Key, params, settlementallocationintent)
	return settlementallocationintent, err
}

// Retrieve SettlementAllocationIntent API.
func (c v2PaymentsSettlementAllocationIntentService) Retrieve(ctx context.Context, id string, params *V2PaymentsSettlementAllocationIntentRetrieveParams) (*V2PaymentsSettlementAllocationIntent, error) {
	if params == nil {
		params = &V2PaymentsSettlementAllocationIntentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/payments/settlement_allocation_intents/%s", id)
	settlementallocationintent := &V2PaymentsSettlementAllocationIntent{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, settlementallocationintent)
	return settlementallocationintent, err
}

// Update SettlementAllocationIntent API.
func (c v2PaymentsSettlementAllocationIntentService) Update(ctx context.Context, id string, params *V2PaymentsSettlementAllocationIntentUpdateParams) (*V2PaymentsSettlementAllocationIntent, error) {
	if params == nil {
		params = &V2PaymentsSettlementAllocationIntentUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/payments/settlement_allocation_intents/%s", id)
	settlementallocationintent := &V2PaymentsSettlementAllocationIntent{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, settlementallocationintent)
	return settlementallocationintent, err
}

// Cancel SettlementAllocationIntent API.
func (c v2PaymentsSettlementAllocationIntentService) Cancel(ctx context.Context, id string, params *V2PaymentsSettlementAllocationIntentCancelParams) (*V2PaymentsSettlementAllocationIntent, error) {
	if params == nil {
		params = &V2PaymentsSettlementAllocationIntentCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/cancel", id)
	settlementallocationintent := &V2PaymentsSettlementAllocationIntent{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, settlementallocationintent)
	return settlementallocationintent, err
}

// Submit SettlementAllocationIntent API.
func (c v2PaymentsSettlementAllocationIntentService) Submit(ctx context.Context, id string, params *V2PaymentsSettlementAllocationIntentSubmitParams) (*V2PaymentsSettlementAllocationIntent, error) {
	if params == nil {
		params = &V2PaymentsSettlementAllocationIntentSubmitParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/submit", id)
	settlementallocationintent := &V2PaymentsSettlementAllocationIntent{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, settlementallocationintent)
	return settlementallocationintent, err
}
