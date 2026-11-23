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

// v2MoneyManagementDisputeService is used to invoke dispute related APIs.
type v2MoneyManagementDisputeService struct {
	B   Backend
	Key string
}

// Creates a Dispute for the specified money movement.
func (c v2MoneyManagementDisputeService) Create(ctx context.Context, params *V2MoneyManagementDisputeCreateParams) (*V2MoneyManagementDispute, error) {
	if params == nil {
		params = &V2MoneyManagementDisputeCreateParams{}
	}
	params.Context = ctx
	dispute := &V2MoneyManagementDispute{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/disputes", c.Key, params, dispute)
	return dispute, err
}

// Retrieves a Dispute.
func (c v2MoneyManagementDisputeService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementDisputeRetrieveParams) (*V2MoneyManagementDispute, error) {
	if params == nil {
		params = &V2MoneyManagementDisputeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/disputes/%s", id)
	dispute := &V2MoneyManagementDispute{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, dispute)
	return dispute, err
}

// Retrieves a list of Disputes.
func (c v2MoneyManagementDisputeService) List(ctx context.Context, listParams *V2MoneyManagementDisputeListParams) *V2List[*V2MoneyManagementDispute] {
	if listParams == nil {
		listParams = &V2MoneyManagementDisputeListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/disputes", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementDispute], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementDispute]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
