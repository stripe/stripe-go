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

// v2MoneyManagementInboundTransferMandateService is used to invoke inboundtransfermandate related APIs.
type v2MoneyManagementInboundTransferMandateService struct {
	B   Backend
	Key string
}

// Create an InboundTransferMandate for a v2 credential. If a pending or
// active mandate already exists for the same user and credential, that
// mandate is returned instead of creating a new one.
func (c v2MoneyManagementInboundTransferMandateService) Create(ctx context.Context, params *V2MoneyManagementInboundTransferMandateCreateParams) (*V2MoneyManagementInboundTransferMandate, error) {
	if params == nil {
		params = &V2MoneyManagementInboundTransferMandateCreateParams{}
	}
	params.Context = ctx
	inboundtransfermandate := &V2MoneyManagementInboundTransferMandate{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/inbound_transfer_mandates", c.Key, params, inboundtransfermandate)
	return inboundtransfermandate, err
}

// Retrieve an InboundTransferMandate by ID.
func (c v2MoneyManagementInboundTransferMandateService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementInboundTransferMandateRetrieveParams) (*V2MoneyManagementInboundTransferMandate, error) {
	if params == nil {
		params = &V2MoneyManagementInboundTransferMandateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/inbound_transfer_mandates/%s", id)
	inboundtransfermandate := &V2MoneyManagementInboundTransferMandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, inboundtransfermandate)
	return inboundtransfermandate, err
}

// Cancel a pending or active InboundTransferMandate.
func (c v2MoneyManagementInboundTransferMandateService) Cancel(ctx context.Context, id string, params *V2MoneyManagementInboundTransferMandateCancelParams) (*V2MoneyManagementInboundTransferMandate, error) {
	if params == nil {
		params = &V2MoneyManagementInboundTransferMandateCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/inbound_transfer_mandates/%s/cancel", id)
	inboundtransfermandate := &V2MoneyManagementInboundTransferMandate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfermandate)
	return inboundtransfermandate, err
}

// Retrieve a list of InboundTransferMandates for the authenticated compartment.
func (c v2MoneyManagementInboundTransferMandateService) List(ctx context.Context, listParams *V2MoneyManagementInboundTransferMandateListParams) *V2List[*V2MoneyManagementInboundTransferMandate] {
	if listParams == nil {
		listParams = &V2MoneyManagementInboundTransferMandateListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/money_management/inbound_transfer_mandates", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2MoneyManagementInboundTransferMandate], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2MoneyManagementInboundTransferMandate]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
