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
