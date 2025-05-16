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

// v1TestHelpersTreasuryOutboundPaymentService is used to invoke /v1/treasury/outbound_payments APIs.
type v1TestHelpersTreasuryOutboundPaymentService struct {
	B   Backend
	Key string
}

// Updates a test mode created OutboundPayment with tracking details. The OutboundPayment must not be cancelable, and cannot be in the canceled or failed states.
func (c v1TestHelpersTreasuryOutboundPaymentService) Update(ctx context.Context, id string, params *TestHelpersTreasuryOutboundPaymentUpdateParams) (*TreasuryOutboundPayment, error) {
	if params == nil {
		params = &TestHelpersTreasuryOutboundPaymentUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/treasury/outbound_payments/%s", id)
	outboundpayment := &TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Transitions a test mode created OutboundPayment to the failed status. The OutboundPayment must already be in the processing state.
func (c v1TestHelpersTreasuryOutboundPaymentService) Fail(ctx context.Context, id string, params *TestHelpersTreasuryOutboundPaymentFailParams) (*TreasuryOutboundPayment, error) {
	if params == nil {
		params = &TestHelpersTreasuryOutboundPaymentFailParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/treasury/outbound_payments/%s/fail", id)
	outboundpayment := &TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Transitions a test mode created OutboundPayment to the posted status. The OutboundPayment must already be in the processing state.
func (c v1TestHelpersTreasuryOutboundPaymentService) Post(ctx context.Context, id string, params *TestHelpersTreasuryOutboundPaymentPostParams) (*TreasuryOutboundPayment, error) {
	if params == nil {
		params = &TestHelpersTreasuryOutboundPaymentPostParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/treasury/outbound_payments/%s/post", id)
	outboundpayment := &TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}

// Transitions a test mode created OutboundPayment to the returned status. The OutboundPayment must already be in the processing state.
func (c v1TestHelpersTreasuryOutboundPaymentService) ReturnOutboundPayment(ctx context.Context, id string, params *TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentParams) (*TreasuryOutboundPayment, error) {
	if params == nil {
		params = &TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/treasury/outbound_payments/%s/return", id)
	outboundpayment := &TreasuryOutboundPayment{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, outboundpayment)
	return outboundpayment, err
}
