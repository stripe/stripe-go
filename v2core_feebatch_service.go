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

// v2CoreFeeBatchService is used to invoke feebatch related APIs.
type v2CoreFeeBatchService struct {
	B   Backend
	Key string
}

// Retrieve a FeeBatch by id.
func (c v2CoreFeeBatchService) Retrieve(ctx context.Context, id string, params *V2CoreFeeBatchRetrieveParams) (*V2CoreFeeBatch, error) {
	if params == nil {
		params = &V2CoreFeeBatchRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/fee_batches/%s", id)
	feebatch := &V2CoreFeeBatch{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, feebatch)
	return feebatch, err
}

// List FeeBatches optionally filtered by collection_record.
func (c v2CoreFeeBatchService) List(ctx context.Context, listParams *V2CoreFeeBatchListParams) *V2List[*V2CoreFeeBatch] {
	if listParams == nil {
		listParams = &V2CoreFeeBatchListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/core/fee_batches", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2CoreFeeBatch], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2CoreFeeBatch]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
