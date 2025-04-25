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

// v2MoneyManagementInboundTransferService is used to invoke inboundtransfer related APIs.
type v2MoneyManagementInboundTransferService struct {
	B   Backend
	Key string
}

// InboundTransfers APIs are used to create, retrieve or list InboundTransfers.
func (c v2MoneyManagementInboundTransferService) Create(ctx context.Context, params *V2MoneyManagementInboundTransferCreateParams) (*V2MoneyManagementInboundTransfer, error) {
	inboundtransfer := &V2MoneyManagementInboundTransfer{}
	if params == nil {
		params = &V2MoneyManagementInboundTransferCreateParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/inbound_transfers", c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Retrieve an InboundTransfer by ID.
func (c v2MoneyManagementInboundTransferService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementInboundTransferRetrieveParams) (*V2MoneyManagementInboundTransfer, error) {
	path := FormatURLPath("/v2/money_management/inbound_transfers/%s", id)
	inboundtransfer := &V2MoneyManagementInboundTransfer{}
	if params == nil {
		params = &V2MoneyManagementInboundTransferRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Retrieves a list of InboundTransfers.
func (c v2MoneyManagementInboundTransferService) List(ctx context.Context, listParams *V2MoneyManagementInboundTransferListParams) Seq2[*V2MoneyManagementInboundTransfer, error] {
	return NewV2List("/v2/money_management/inbound_transfers", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementInboundTransfer], error) {
		page := &V2Page[*V2MoneyManagementInboundTransfer]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
