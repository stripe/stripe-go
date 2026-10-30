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

// v2MoneyManagementInboundFundingService is used to invoke inboundfunding related APIs.
type v2MoneyManagementInboundFundingService struct {
	B   Backend
	Key string
}

// Retrieves an InboundFunding by ID.
func (c v2MoneyManagementInboundFundingService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementInboundFundingRetrieveParams) (*V2MoneyManagementInboundFunding, error) {
	if params == nil {
		params = &V2MoneyManagementInboundFundingRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/inbound_fundings/%s", id)
	inboundfunding := &V2MoneyManagementInboundFunding{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, inboundfunding)
	return inboundfunding, err
}

// Lists InboundFundings.
func (c v2MoneyManagementInboundFundingService) List(ctx context.Context, listParams *V2MoneyManagementInboundFundingListParams) *V2List[*V2MoneyManagementInboundFunding] {
	if listParams == nil {
		listParams = &V2MoneyManagementInboundFundingListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/inbound_fundings", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementInboundFunding], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementInboundFunding]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
