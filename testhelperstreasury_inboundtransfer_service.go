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

// v1TestHelpersTreasuryInboundTransferService is used to invoke /v1/treasury/inbound_transfers APIs.
type v1TestHelpersTreasuryInboundTransferService struct {
	B   Backend
	Key string
}

// Transitions a test mode created InboundTransfer to the failed status. The InboundTransfer must already be in the processing state.
func (c v1TestHelpersTreasuryInboundTransferService) Fail(ctx context.Context, id string, params *TestHelpersTreasuryInboundTransferFailParams) (*TreasuryInboundTransfer, error) {
	if params == nil {
		params = &TestHelpersTreasuryInboundTransferFailParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/treasury/inbound_transfers/%s/fail", id)
	inboundtransfer := &TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the succeeded state.
func (c v1TestHelpersTreasuryInboundTransferService) ReturnInboundTransfer(ctx context.Context, id string, params *TestHelpersTreasuryInboundTransferReturnInboundTransferParams) (*TreasuryInboundTransfer, error) {
	if params == nil {
		params = &TestHelpersTreasuryInboundTransferReturnInboundTransferParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/treasury/inbound_transfers/%s/return", id)
	inboundtransfer := &TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}

// Transitions a test mode created InboundTransfer to the succeeded status. The InboundTransfer must already be in the processing state.
func (c v1TestHelpersTreasuryInboundTransferService) Succeed(ctx context.Context, id string, params *TestHelpersTreasuryInboundTransferSucceedParams) (*TreasuryInboundTransfer, error) {
	if params == nil {
		params = &TestHelpersTreasuryInboundTransferSucceedParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/treasury/inbound_transfers/%s/succeed", id)
	inboundtransfer := &TreasuryInboundTransfer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inboundtransfer)
	return inboundtransfer, err
}
