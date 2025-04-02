//
//
// File generated from our OpenAPI spec
//
//

// Package outboundpayment provides the outboundpayment related APIs
package outboundpayment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke outboundpayment related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an OutboundPayment.
func (c Client) New(params *stripe.V2MoneyManagementOutboundPaymentParams) (*stripe.V2MoneyManagementOutboundPayment, error) {
	outboundpayment := &stripe.V2MoneyManagementOutboundPayment{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/outbound_payments", c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Retrieves the details of an existing OutboundPayment by passing the unique OutboundPayment ID from either the OutboundPayment create or list response.
func (c Client) Get(id string, params *stripe.V2MoneyManagementOutboundPaymentParams) (*stripe.V2MoneyManagementOutboundPayment, error) {
	path := stripe.FormatURLPath("/v2/money_management/outbound_payments/%s", id)
	outboundpayment := &stripe.V2MoneyManagementOutboundPayment{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Cancels an OutboundPayment. Only processing OutboundPayments can be canceled.
func (c Client) Cancel(id string, params *stripe.V2MoneyManagementOutboundPaymentCancelParams) (*stripe.V2MoneyManagementOutboundPayment, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/outbound_payments/%s/cancel", id)
	outboundpayment := &stripe.V2MoneyManagementOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Returns a list of OutboundPayments that match the provided filters.
func (c Client) All(listParams *stripe.V2MoneyManagementOutboundPaymentListParams) stripe.Seq2[stripe.V2MoneyManagementOutboundPayment, error] {
	return stripe.NewV2List("/v2/money_management/outbound_payments", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[stripe.V2MoneyManagementOutboundPayment], error) {
		page := &stripe.V2Page[stripe.V2MoneyManagementOutboundPayment]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
