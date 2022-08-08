//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Transitions a test mode created OutboundTransfer to the failed status. The OutboundTransfer must already be in the processing state.
type TestHelpersTreasuryOutboundTransferFailParams struct {
	Params `form:"*"`
}

// Transitions a test mode created OutboundTransfer to the posted status. The OutboundTransfer must already be in the processing state.
type TestHelpersTreasuryOutboundTransferPostParams struct {
	Params `form:"*"`
}

// Details about a returned OutboundTransfer.
type TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams struct {
	// Reason for the return.
	Code *string `form:"code"`
}

// Transitions a test mode created OutboundTransfer to the returned status. The OutboundTransfer must already be in the processing state.
type TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams struct {
	Params `form:"*"`
	// Details about a returned OutboundTransfer.
	ReturnedDetails *TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams `form:"returned_details"`
}
