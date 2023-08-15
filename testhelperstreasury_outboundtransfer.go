//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Transitions a test mode created OutboundTransfer to the failed status. The OutboundTransfer must already be in the processing state.
type TestHelpersTreasuryOutboundTransferFailParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryOutboundTransferFailParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Transitions a test mode created OutboundTransfer to the posted status. The OutboundTransfer must already be in the processing state.
type TestHelpersTreasuryOutboundTransferPostParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryOutboundTransferPostParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details about a returned OutboundTransfer.
type TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams struct {
	// Reason for the return.
	Code *string `form:"code"`
}

// Transitions a test mode created OutboundTransfer to the returned status. The OutboundTransfer must already be in the processing state.
type TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Details about a returned OutboundTransfer.
	ReturnedDetails *TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams `form:"returned_details"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
