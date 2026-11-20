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

// v2MoneyManagementDisputeAdjustmentService is used to invoke disputeadjustment related APIs.
type v2MoneyManagementDisputeAdjustmentService struct {
	B   Backend
	Key string
}

// Retrieves a DisputeAdjustment by ID.
func (c v2MoneyManagementDisputeAdjustmentService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementDisputeAdjustmentRetrieveParams) (*V2MoneyManagementDisputeAdjustment, error) {
	if params == nil {
		params = &V2MoneyManagementDisputeAdjustmentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/dispute_adjustments/%s", id)
	disputeadjustment := &V2MoneyManagementDisputeAdjustment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, disputeadjustment)
	return disputeadjustment, err
}

// Retrieves a list of DisputeAdjustments.
func (c v2MoneyManagementDisputeAdjustmentService) List(ctx context.Context, listParams *V2MoneyManagementDisputeAdjustmentListParams) *V2List[*V2MoneyManagementDisputeAdjustment] {
	if listParams == nil {
		listParams = &V2MoneyManagementDisputeAdjustmentListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/dispute_adjustments", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementDisputeAdjustment], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementDisputeAdjustment]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
