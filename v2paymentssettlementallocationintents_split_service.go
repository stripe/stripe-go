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

// List SettlementAllocationIntentSplits API.
func (c v2PaymentsSettlementAllocationIntentsSplitService) List(ctx context.Context, listParams *V2PaymentsSettlementAllocationIntentsSplitListParams) *V2List[*V2PaymentsSettlementAllocationIntentSplit] {
	if listParams == nil {
		listParams = &V2PaymentsSettlementAllocationIntentsSplitListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/payments/settlement_allocation_intents/%s/splits", StringValue(
			listParams.SettlementAllocationIntentID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2PaymentsSettlementAllocationIntentSplit], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2PaymentsSettlementAllocationIntentSplit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
