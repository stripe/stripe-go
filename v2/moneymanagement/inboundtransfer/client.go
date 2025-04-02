//
//
// File generated from our OpenAPI spec
//
//

// Package inboundtransfer provides the inboundtransfer related APIs
package inboundtransfer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke inboundtransfer related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// InboundTransfers APIs are used to create, retrieve or list InboundTransfers.
func (c Client) New(params *stripe.V2MoneyManagementInboundTransferParams) (*stripe.V2MoneyManagementInboundTransfer, error) {
	inboundtransfer := &stripe.V2MoneyManagementInboundTransfer{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/inbound_transfers", c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Retrieve an InboundTransfer by ID.
func (c Client) Get(id string, params *stripe.V2MoneyManagementInboundTransferParams) (*stripe.V2MoneyManagementInboundTransfer, error) {
	path := stripe.FormatURLPath("/v2/money_management/inbound_transfers/%s", id)
	inboundtransfer := &stripe.V2MoneyManagementInboundTransfer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Retrieves a list of InboundTransfers.
func (c Client) All(listParams *stripe.V2MoneyManagementInboundTransferListParams) stripe.Seq2[stripe.V2MoneyManagementInboundTransfer, error] {
	return stripe.NewV2List("/v2/money_management/inbound_transfers", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[stripe.V2MoneyManagementInboundTransfer], error) {
		page := &stripe.V2Page[stripe.V2MoneyManagementInboundTransfer]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
