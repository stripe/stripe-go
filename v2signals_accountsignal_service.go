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

// v2SignalsAccountSignalService is used to invoke accountsignal related APIs.
type v2SignalsAccountSignalService struct {
	B   Backend
	Key string
}

// Retrieves an AccountSignal by its ID.
func (c v2SignalsAccountSignalService) Retrieve(ctx context.Context, id string, params *V2SignalsAccountSignalRetrieveParams) (*V2SignalsAccountSignal, error) {
	if params == nil {
		params = &V2SignalsAccountSignalRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/signals/account_signals/%s", id)
	accountsignal := &V2SignalsAccountSignal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountsignal)
	return accountsignal, err
}

// Lists AccountSignals for a given account or customer, filtered by signal type.
func (c v2SignalsAccountSignalService) List(ctx context.Context, listParams *V2SignalsAccountSignalListParams) *V2List[*V2SignalsAccountSignal] {
	if listParams == nil {
		listParams = &V2SignalsAccountSignalListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/signals/account_signals", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2SignalsAccountSignal], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2SignalsAccountSignal]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
