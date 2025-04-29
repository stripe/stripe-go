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

// v2MoneyManagementAdjustmentService is used to invoke adjustment related APIs.
type v2MoneyManagementAdjustmentService struct {
	B   Backend
	Key string
}

// Retrieves the details of an Adjustment by ID.
func (c v2MoneyManagementAdjustmentService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementAdjustmentRetrieveParams) (*V2MoneyManagementAdjustment, error) {
	path := FormatURLPath("/v2/money_management/adjustments/%s", id)
	adjustment := &V2MoneyManagementAdjustment{}
	if params == nil {
		params = &V2MoneyManagementAdjustmentRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, adjustment)
	return adjustment, err
}

// Returns a list of Adjustments that match the provided filters.
func (c v2MoneyManagementAdjustmentService) List(ctx context.Context, listParams *V2MoneyManagementAdjustmentListParams) Seq2[*V2MoneyManagementAdjustment, error] {
	if listParams == nil {
		listParams = &V2MoneyManagementAdjustmentListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/money_management/adjustments", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementAdjustment], error) {
		page := &V2Page[*V2MoneyManagementAdjustment]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
