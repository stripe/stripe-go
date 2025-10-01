//
//
// File generated from our OpenAPI spec
//
//

// Package outboundpaymentquote provides the outboundpaymentquote related APIs
package outboundpaymentquote

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke outboundpaymentquote related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an OutboundPaymentQuote usable in an OutboundPayment.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementOutboundPaymentQuoteParams) (*stripe.V2MoneyManagementOutboundPaymentQuote, error) {
	outboundpaymentquote := &stripe.V2MoneyManagementOutboundPaymentQuote{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/outbound_payment_quotes", c.Key, params, outboundpaymentquote)
	return outboundpaymentquote, err
}

// Retrieves the details of an existing OutboundPaymentQuote by passing the unique OutboundPaymentQuote ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementOutboundPaymentQuoteParams) (*stripe.V2MoneyManagementOutboundPaymentQuote, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/outbound_payment_quotes/%s", id)
	outboundpaymentquote := &stripe.V2MoneyManagementOutboundPaymentQuote{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, outboundpaymentquote)
	return outboundpaymentquote, err
}
