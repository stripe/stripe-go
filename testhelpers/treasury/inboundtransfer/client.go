//
//
// File generated from our OpenAPI spec
//
//

// Package inboundtransfer provides the /v1/treasury/inbound_transfers APIs
package inboundtransfer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/treasury/inbound_transfers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Transitions a test mode created InboundTransfer to the failed status. The InboundTransfer must already be in the processing state.
func Fail(id string, params *stripe.TestHelpersTreasuryInboundTransferFailParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().Fail(id, params)
}

// Transitions a test mode created InboundTransfer to the failed status. The InboundTransfer must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Fail(id string, params *stripe.TestHelpersTreasuryInboundTransferFailParams) (*stripe.TreasuryInboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/inbound_transfers/%s/fail", id)
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the succeeded state.
func ReturnInboundTransfer(id string, params *stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().ReturnInboundTransfer(id, params)
}

// Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the succeeded state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ReturnInboundTransfer(id string, params *stripe.TestHelpersTreasuryInboundTransferReturnInboundTransferParams) (*stripe.TreasuryInboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/inbound_transfers/%s/return", id)
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Transitions a test mode created InboundTransfer to the succeeded status. The InboundTransfer must already be in the processing state.
func Succeed(id string, params *stripe.TestHelpersTreasuryInboundTransferSucceedParams) (*stripe.TreasuryInboundTransfer, error) {
	return getC().Succeed(id, params)
}

// Transitions a test mode created InboundTransfer to the succeeded status. The InboundTransfer must already be in the processing state.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Succeed(id string, params *stripe.TestHelpersTreasuryInboundTransferSucceedParams) (*stripe.TreasuryInboundTransfer, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/treasury/inbound_transfers/%s/succeed", id)
	inboundtransfer := &stripe.TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
