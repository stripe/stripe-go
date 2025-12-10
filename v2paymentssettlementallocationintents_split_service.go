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

// v2PaymentsSettlementAllocationIntentsSplitService is used to invoke split related APIs.
type v2PaymentsSettlementAllocationIntentsSplitService struct {
	B   Backend
	Key string
}

// Create SettlementAllocationIntentSplit API.
func (c v2PaymentsSettlementAllocationIntentsSplitService) Create(ctx context.Context, params *V2PaymentsSettlementAllocationIntentsSplitCreateParams) (*V2PaymentsSettlementAllocationIntentSplit, error) {
	if params == nil {
		params = &V2PaymentsSettlementAllocationIntentsSplitCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/splits", StringValue(
			params.SettlementAllocationIntentID))
	settlementallocationintentsplit := &V2PaymentsSettlementAllocationIntentSplit{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, settlementallocationintentsplit)
	return settlementallocationintentsplit, err
}

// Retrieve SettlementAllocationIntentSplit API.
func (c v2PaymentsSettlementAllocationIntentsSplitService) Retrieve(ctx context.Context, id string, params *V2PaymentsSettlementAllocationIntentsSplitRetrieveParams) (*V2PaymentsSettlementAllocationIntentSplit, error) {
	if params == nil {
		params = &V2PaymentsSettlementAllocationIntentsSplitRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/splits/%s", StringValue(
			params.SettlementAllocationIntentID), id)
	settlementallocationintentsplit := &V2PaymentsSettlementAllocationIntentSplit{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, settlementallocationintentsplit)
	return settlementallocationintentsplit, err
}

// Cancel SettlementAllocationIntentSplit API.
func (c v2PaymentsSettlementAllocationIntentsSplitService) Cancel(ctx context.Context, id string, params *V2PaymentsSettlementAllocationIntentsSplitCancelParams) (*V2PaymentsSettlementAllocationIntentSplit, error) {
	if params == nil {
		params = &V2PaymentsSettlementAllocationIntentsSplitCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/splits/%s/cancel", StringValue(
			params.SettlementAllocationIntentID), id)
	settlementallocationintentsplit := &V2PaymentsSettlementAllocationIntentSplit{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, settlementallocationintentsplit)
	return settlementallocationintentsplit, err
}
