//
//
// File generated from our OpenAPI spec
//
//

// Package outboundtransfer provides the outboundtransfer related APIs
package outboundtransfer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke outboundtransfer related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an OutboundTransfer.
func (c Client) New(params *stripe.V2MoneyManagementOutboundTransferParams) (*stripe.V2MoneyManagementOutboundTransfer, error) {
	outboundtransfer := &stripe.V2MoneyManagementOutboundTransfer{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/outbound_transfers", c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Retrieves the details of an existing OutboundTransfer by passing the unique OutboundTransfer ID from either the OutboundPayment create or list response.
func (c Client) Get(id string, params *stripe.V2MoneyManagementOutboundTransferParams) (*stripe.V2MoneyManagementOutboundTransfer, error) {
	path := stripe.FormatURLPath("/v2/money_management/outbound_transfers/%s", id)
	outboundtransfer := &stripe.V2MoneyManagementOutboundTransfer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Cancels an OutboundTransfer. Only processing OutboundTransfers can be canceled.
func (c Client) Cancel(id string, params *stripe.V2MoneyManagementOutboundTransferCancelParams) (*stripe.V2MoneyManagementOutboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/outbound_transfers/%s/cancel", id)
	outboundtransfer := &stripe.V2MoneyManagementOutboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundtransfer)
	return outboundtransfer, err
}

// Returns a list of OutboundTransfers that match the provided filters.
func (c Client) All(listParams *stripe.V2MoneyManagementOutboundTransferListParams) stripe.Seq2[*stripe.V2MoneyManagementOutboundTransfer, error] {
	return stripe.NewV2List("/v2/money_management/outbound_transfers", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2MoneyManagementOutboundTransfer], error) {
		page := &stripe.V2Page[*stripe.V2MoneyManagementOutboundTransfer]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
