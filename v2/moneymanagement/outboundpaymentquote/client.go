//
//
// File generated from our OpenAPI spec
//
//

// Package outboundpaymentquote provides the outboundpaymentquote related APIs
package outboundpaymentquote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke outboundpaymentquote related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an OutboundPaymentQuote usable in an OutboundPayment.
func (c Client) New(params *stripe.V2MoneyManagementOutboundPaymentQuoteParams) (*stripe.V2MoneyManagementOutboundPaymentQuote, error) {
	outboundpaymentquote := &stripe.V2MoneyManagementOutboundPaymentQuote{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/outbound_payment_quotes", c.Key, params, outboundpaymentquote)
	return outboundpaymentquote, err
}

// Retrieves the details of an existing OutboundPaymentQuote by passing the unique OutboundPaymentQuote ID.
func (c Client) Get(id string, params *stripe.V2MoneyManagementOutboundPaymentQuoteParams) (*stripe.V2MoneyManagementOutboundPaymentQuote, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/outbound_payment_quotes/%s", id)
	outboundpaymentquote := &stripe.V2MoneyManagementOutboundPaymentQuote{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundpaymentquote)
	return outboundpaymentquote, err
}
