//
//
// File generated from our OpenAPI spec
//
//

// Package outboundpayment provides the /v1/treasury/outbound_payments APIs
package outboundpayment

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/treasury/outbound_payments APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Updates a test mode created OutboundPayment with tracking details. The OutboundPayment must not be cancelable, and cannot be in the canceled or failed states.
func Update(id string, params *stripe.TestHelpersTreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	return getC().Update(id, params)
}

// Updates a test mode created OutboundPayment with tracking details. The OutboundPayment must not be cancelable, and cannot be in the canceled or failed states.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TestHelpersTreasuryOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_payments/%s", id)
	outboundpayment := &stripe.TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Transitions a test mode created OutboundPayment to the failed status. The OutboundPayment must already be in the processing state.
func Fail(id string, params *stripe.TestHelpersTreasuryOutboundPaymentFailParams) (*stripe.TreasuryOutboundPayment, error) {
	return getC().Fail(id, params)
}

// Transitions a test mode created OutboundPayment to the failed status. The OutboundPayment must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Fail(id string, params *stripe.TestHelpersTreasuryOutboundPaymentFailParams) (*stripe.TreasuryOutboundPayment, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_payments/%s/fail", id)
	outboundpayment := &stripe.TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Transitions a test mode created OutboundPayment to the posted status. The OutboundPayment must already be in the processing state.
func Post(id string, params *stripe.TestHelpersTreasuryOutboundPaymentPostParams) (*stripe.TreasuryOutboundPayment, error) {
	return getC().Post(id, params)
}

// Transitions a test mode created OutboundPayment to the posted status. The OutboundPayment must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Post(id string, params *stripe.TestHelpersTreasuryOutboundPaymentPostParams) (*stripe.TreasuryOutboundPayment, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_payments/%s/post", id)
	outboundpayment := &stripe.TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Transitions a test mode created OutboundPayment to the returned status. The OutboundPayment must already be in the processing state.
func ReturnOutboundPayment(id string, params *stripe.TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	return getC().ReturnOutboundPayment(id, params)
}

// Transitions a test mode created OutboundPayment to the returned status. The OutboundPayment must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReturnOutboundPayment(id string, params *stripe.TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentParams) (*stripe.TreasuryOutboundPayment, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/outbound_payments/%s/return", id)
	outboundpayment := &stripe.TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
