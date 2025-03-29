//
//
// File generated from our OpenAPI spec
//
//

// Package quote provides the quote related APIs
package quote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke quote related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an OutboundPaymentQuote usable in an OutboundPayment.
func (c Client) New(params *stripe.V2MoneyManagementOutboundPaymentsQuoteParams) (*stripe.V2MoneyManagementOutboundPaymentQuote, error) {
	outboundpaymentquote := &stripe.V2MoneyManagementOutboundPaymentQuote{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/outbound_payments/quotes", c.Key, params, outboundpaymentquote)
	return outboundpaymentquote, err
}
