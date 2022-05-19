//
//
// File generated from our OpenAPI spec
//
//

// Package outboundpayment provides the /treasury/outbound_payments APIs
package outboundpayment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /treasury/outbound_payments APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Fail is the method for the `POST /v1/test_helpers/treasury/outbound_payments/{id}/fail` API.
func Fail(id string, params *stripe.TestHelpersTreasuryOutboundPaymentFailParams) (*stripe.TreasuryOutboundPayment, error) {
	return getC().Fail(id, params)
}

// Fail is the method for the `POST /v1/test_helpers/treasury/outbound_payments/{id}/fail` API.
func (c Client) Fail(id string, params *stripe.TestHelpersTreasuryOutboundPaymentFailParams) (*stripe.TreasuryOutboundPayment, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_payments/%s/fail",
		id,
	)
	outboundpayment := &stripe.TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Post is the method for the `POST /v1/test_helpers/treasury/outbound_payments/{id}/post` API.
func Post(id string, params *stripe.TestHelpersTreasuryOutboundPaymentPostParams) (*stripe.TreasuryOutboundPayment, error) {
	return getC().Post(id, params)
}

// Post is the method for the `POST /v1/test_helpers/treasury/outbound_payments/{id}/post` API.
func (c Client) Post(id string, params *stripe.TestHelpersTreasuryOutboundPaymentPostParams) (*stripe.TreasuryOutboundPayment, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_payments/%s/post",
		id,
	)
	outboundpayment := &stripe.TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// ReturnOutboundPayment is the method for the `POST /v1/test_helpers/treasury/outbound_payments/{id}/return` API.
func ReturnOutboundPayment(id string, params *stripe.TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	return getC().ReturnOutboundPayment(id, params)
}

// ReturnOutboundPayment is the method for the `POST /v1/test_helpers/treasury/outbound_payments/{id}/return` API.
func (c Client) ReturnOutboundPayment(id string, params *stripe.TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_payments/%s/return",
		id,
	)
	outboundpayment := &stripe.TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
