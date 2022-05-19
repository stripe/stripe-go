//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Transitions a test mode created OutboundPayment to the failed status. The OutboundPayment must already be in the processing state.
type TestHelpersTreasuryOutboundPaymentFailParams struct {
	Params `form:"*"`
}

// Transitions a test mode created OutboundPayment to the posted status. The OutboundPayment must already be in the processing state.
type TestHelpersTreasuryOutboundPaymentPostParams struct {
	Params `form:"*"`
}

// Optional hash to set the the return code.
type TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentReturnedDetailsParams struct {
	// The return code to be set on the OutboundPayment object.
	Code *string `form:"code"`
}

// Transitions a test mode created OutboundPayment to the returned status. The OutboundPayment must already be in the processing state.
type TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentParams struct {
	Params `form:"*"`
	// Optional hash to set the the return code.
	ReturnedDetails *TestHelpersTreasuryOutboundPaymentReturnOutboundPaymentReturnedDetailsParams `form:"returned_details"`
}
