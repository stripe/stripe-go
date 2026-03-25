//
//
// File generated from our OpenAPI spec
//
//

package stripe

// ACH network tracking details.
type TestHelpersTreasuryOutboundTransferTrackingDetailsACHParams struct {
	// ACH trace ID for funds sent over the `ach` network.
	TraceID *string `form:"trace_id" json:"trace_id"`
}

// US domestic wire network tracking details.
type TestHelpersTreasuryOutboundTransferTrackingDetailsUSDomesticWireParams struct {
	// CHIPS System Sequence Number (SSN) for funds sent over the `us_domestic_wire` network.
	Chips *string `form:"chips" json:"chips,omitempty"`
	// IMAD for funds sent over the `us_domestic_wire` network.
	Imad *string `form:"imad" json:"imad,omitempty"`
	// OMAD for funds sent over the `us_domestic_wire` network.
	Omad *string `form:"omad" json:"omad,omitempty"`
}

// Details about network-specific tracking information.
type TestHelpersTreasuryOutboundTransferTrackingDetailsParams struct {
	// ACH network tracking details.
	ACH *TestHelpersTreasuryOutboundTransferTrackingDetailsACHParams `form:"ach" json:"ach,omitempty"`
	// The US bank account network used to send funds.
	Type *string `form:"type" json:"type"`
	// US domestic wire network tracking details.
	USDomesticWire *TestHelpersTreasuryOutboundTransferTrackingDetailsUSDomesticWireParams `form:"us_domestic_wire" json:"us_domestic_wire,omitempty"`
}

// Updates a test mode created OutboundTransfer with tracking details. The OutboundTransfer must not be cancelable, and cannot be in the canceled or failed states.
type TestHelpersTreasuryOutboundTransferParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Details about network-specific tracking information.
	TrackingDetails *TestHelpersTreasuryOutboundTransferTrackingDetailsParams `form:"tracking_details" json:"tracking_details"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryOutboundTransferParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Transitions a test mode created OutboundTransfer to the failed status. The OutboundTransfer must already be in the processing state.
type TestHelpersTreasuryOutboundTransferFailParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryOutboundTransferFailParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Transitions a test mode created OutboundTransfer to the posted status. The OutboundTransfer must already be in the processing state.
type TestHelpersTreasuryOutboundTransferPostParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryOutboundTransferPostParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details about a returned OutboundTransfer.
type TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams struct {
	// Reason for the return.
	Code *string `form:"code" json:"code,omitempty"`
}

// Transitions a test mode created OutboundTransfer to the returned status. The OutboundTransfer must already be in the processing state.
type TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Details about a returned OutboundTransfer.
	ReturnedDetails *TestHelpersTreasuryOutboundTransferReturnOutboundTransferReturnedDetailsParams `form:"returned_details" json:"returned_details,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryOutboundTransferReturnOutboundTransferParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// ACH network tracking details.
type TestHelpersTreasuryOutboundTransferUpdateTrackingDetailsACHParams struct {
	// ACH trace ID for funds sent over the `ach` network.
	TraceID *string `form:"trace_id" json:"trace_id"`
}

// US domestic wire network tracking details.
type TestHelpersTreasuryOutboundTransferUpdateTrackingDetailsUSDomesticWireParams struct {
	// CHIPS System Sequence Number (SSN) for funds sent over the `us_domestic_wire` network.
	Chips *string `form:"chips" json:"chips,omitempty"`
	// IMAD for funds sent over the `us_domestic_wire` network.
	Imad *string `form:"imad" json:"imad,omitempty"`
	// OMAD for funds sent over the `us_domestic_wire` network.
	Omad *string `form:"omad" json:"omad,omitempty"`
}

// Details about network-specific tracking information.
type TestHelpersTreasuryOutboundTransferUpdateTrackingDetailsParams struct {
	// ACH network tracking details.
	ACH *TestHelpersTreasuryOutboundTransferUpdateTrackingDetailsACHParams `form:"ach" json:"ach,omitempty"`
	// The US bank account network used to send funds.
	Type *string `form:"type" json:"type"`
	// US domestic wire network tracking details.
	USDomesticWire *TestHelpersTreasuryOutboundTransferUpdateTrackingDetailsUSDomesticWireParams `form:"us_domestic_wire" json:"us_domestic_wire,omitempty"`
}

// Updates a test mode created OutboundTransfer with tracking details. The OutboundTransfer must not be cancelable, and cannot be in the canceled or failed states.
type TestHelpersTreasuryOutboundTransferUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Details about network-specific tracking information.
	TrackingDetails *TestHelpersTreasuryOutboundTransferUpdateTrackingDetailsParams `form:"tracking_details" json:"tracking_details"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryOutboundTransferUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
