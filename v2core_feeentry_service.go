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

// v2CoreFeeEntryService is used to invoke feeentry related APIs.
type v2CoreFeeEntryService struct {
	B   Backend
	Key string
}

// Retrieve a FeeEntry by id.
func (c v2CoreFeeEntryService) Retrieve(ctx context.Context, id string, params *V2CoreFeeEntryRetrieveParams) (*V2CoreFeeEntry, error) {
	if params == nil {
		params = &V2CoreFeeEntryRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/fee_entries/%s", id)
	feeentry := &V2CoreFeeEntry{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, feeentry)
	return feeentry, err
}

// List FeeEntries optionally filtered by incurred_by, fee_batch, or collection_record (at most one filter at a time).
func (c v2CoreFeeEntryService) List(ctx context.Context, listParams *V2CoreFeeEntryListParams) *V2List[*V2CoreFeeEntry] {
	if listParams == nil {
		listParams = &V2CoreFeeEntryListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/core/fee_entries", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2CoreFeeEntry], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2CoreFeeEntry]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
