//
//
// File generated from our OpenAPI spec
//
//

package stripe

// ACH network failure details.
type TestHelpersTreasuryInboundTransferFailFailureDetailsNetworkDetailsACHParams struct {
	// NACHA ACH return code.
	ReturnCode *string `form:"return_code" json:"return_code"`
}

// Network specific details for the failure.
type TestHelpersTreasuryInboundTransferFailFailureDetailsNetworkDetailsParams struct {
	// ACH network failure details.
	ACH *TestHelpersTreasuryInboundTransferFailFailureDetailsNetworkDetailsACHParams `form:"ach" json:"ach,omitempty"`
	// The network used to transfer funds.
	Type *string `form:"type" json:"type"`
}

// Details about a failed InboundTransfer.
type TestHelpersTreasuryInboundTransferFailFailureDetailsParams struct {
	// Reason for the failure.
	Code *string `form:"code" json:"code,omitempty"`
	// Network specific details for the failure.
	NetworkDetails *TestHelpersTreasuryInboundTransferFailFailureDetailsNetworkDetailsParams `form:"network_details" json:"network_details,omitempty"`
}

// Transitions a test mode created InboundTransfer to the failed status. The InboundTransfer must already be in the processing state.
type TestHelpersTreasuryInboundTransferFailParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Details about a failed InboundTransfer.
	FailureDetails *TestHelpersTreasuryInboundTransferFailFailureDetailsParams `form:"failure_details" json:"failure_details,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryInboundTransferFailParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// ACH network failure details.
type TestHelpersTreasuryInboundTransferReturnInboundTransferFailureDetailsNetworkDetailsACHParams struct {
	// NACHA ACH return code.
	ReturnCode *string `form:"return_code" json:"return_code"`
}

// Network specific details for the failure.
type TestHelpersTreasuryInboundTransferReturnInboundTransferFailureDetailsNetworkDetailsParams struct {
	// ACH network failure details.
	ACH *TestHelpersTreasuryInboundTransferReturnInboundTransferFailureDetailsNetworkDetailsACHParams `form:"ach" json:"ach,omitempty"`
	// The network used to transfer funds.
	Type *string `form:"type" json:"type"`
}

// Details about a returned InboundTransfer.
type TestHelpersTreasuryInboundTransferReturnInboundTransferFailureDetailsParams struct {
	// Reason for the failure.
	Code *string `form:"code" json:"code,omitempty"`
	// Network specific details for the failure.
	NetworkDetails *TestHelpersTreasuryInboundTransferReturnInboundTransferFailureDetailsNetworkDetailsParams `form:"network_details" json:"network_details,omitempty"`
}

// Marks the test mode InboundTransfer object as returned and links the InboundTransfer to a ReceivedDebit. The InboundTransfer must already be in the succeeded state.
type TestHelpersTreasuryInboundTransferReturnInboundTransferParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Details about a returned InboundTransfer.
	FailureDetails *TestHelpersTreasuryInboundTransferReturnInboundTransferFailureDetailsParams `form:"failure_details" json:"failure_details,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryInboundTransferReturnInboundTransferParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Transitions a test mode created InboundTransfer to the succeeded status. The InboundTransfer must already be in the processing state.
type TestHelpersTreasuryInboundTransferSucceedParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryInboundTransferSucceedParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
