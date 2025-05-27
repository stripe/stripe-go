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

// v2MoneyManagementOutboundTransferService is used to invoke outboundtransfer related APIs.
type v2MoneyManagementOutboundTransferService struct {
	B   Backend
	Key string
}

// Creates an OutboundTransfer.
func (c v2MoneyManagementOutboundTransferService) Create(ctx context.Context, params *V2MoneyManagementOutboundTransferCreateParams) (*V2MoneyManagementOutboundTransfer, error) {
	if params == nil {
		params = &V2MoneyManagementOutboundTransferCreateParams{}
	}
	params.Context = ctx
	outboundtransfer := &V2MoneyManagementOutboundTransfer{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/outbound_transfers", c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Retrieves the details of an existing OutboundTransfer by passing the unique OutboundTransfer ID from either the OutboundPayment create or list response.
func (c v2MoneyManagementOutboundTransferService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementOutboundTransferRetrieveParams) (*V2MoneyManagementOutboundTransfer, error) {
	if params == nil {
		params = &V2MoneyManagementOutboundTransferRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/outbound_transfers/%s", id)
	outboundtransfer := &V2MoneyManagementOutboundTransfer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Cancels an OutboundTransfer. Only processing OutboundTransfers can be canceled.
func (c v2MoneyManagementOutboundTransferService) Cancel(ctx context.Context, id string, params *V2MoneyManagementOutboundTransferCancelParams) (*V2MoneyManagementOutboundTransfer, error) {
	if params == nil {
		params = &V2MoneyManagementOutboundTransferCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/outbound_transfers/%s/cancel", id)
	outboundtransfer := &V2MoneyManagementOutboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Returns a list of OutboundTransfers that match the provided filters.
func (c v2MoneyManagementOutboundTransferService) List(ctx context.Context, listParams *V2MoneyManagementOutboundTransferListParams) Seq2[*V2MoneyManagementOutboundTransfer, error] {
	if listParams == nil {
		listParams = &V2MoneyManagementOutboundTransferListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/money_management/outbound_transfers", listParams, func(path string, p ParamsContainer) (*V2Page[*V2MoneyManagementOutboundTransfer], error) {
		page := &V2Page[*V2MoneyManagementOutboundTransfer]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
