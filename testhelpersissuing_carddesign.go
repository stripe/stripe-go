//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Updates the status of the specified testmode card design object to active.
type TestHelpersIssuingCardDesignActivateTestmodeParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingCardDesignActivateTestmodeParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates the status of the specified testmode card design object to inactive.
type TestHelpersIssuingCardDesignDeactivateTestmodeParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingCardDesignDeactivateTestmodeParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The reason(s) the card design was rejected.
type TestHelpersIssuingCardDesignRejectTestmodeRejectionReasonsParams struct {
	// The reason(s) the card logo was rejected.
	CardLogo []*string `form:"card_logo"`
	// The reason(s) the carrier text was rejected.
	CarrierText []*string `form:"carrier_text"`
}

// Updates the status of the specified testmode card design object to rejected.
type TestHelpersIssuingCardDesignRejectTestmodeParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The reason(s) the card design was rejected.
	RejectionReasons *TestHelpersIssuingCardDesignRejectTestmodeRejectionReasonsParams `form:"rejection_reasons"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersIssuingCardDesignRejectTestmodeParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
