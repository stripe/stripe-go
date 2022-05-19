//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Transitions a test mode created InboundTransfer to the succeeded status. The InboundTransfer must already be in the processing state.
type TestHelpersTreasuryInboundTransferSucceedParams struct {
	Params `form:"*"`
}

// Details about a failed InboundTransfer.
type TestHelpersTreasuryInboundTransferFailFailureDetailsParams struct {
	// Reason for the failure.
	Code *string `form:"code"`
}

// Transitions a test mode created InboundTransfer to the failed status. The InboundTransfer must already be in the processing state.
type TestHelpersTreasuryInboundTransferFailParams struct {
	Params `form:"*"`
	// Details about a failed InboundTransfer.
	FailureDetails *TestHelpersTreasuryInboundTransferFailFailureDetailsParams `form:"failure_details"`
}

// Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the succeeded state.
type TestHelpersTreasuryInboundTransferReturnInboundTransferParams struct {
	Params `form:"*"`
}
