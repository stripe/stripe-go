//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a test financing offer for a connected account.
type TestHelpersCapitalFinancingOfferParams struct {
	Params        `form:"*"`
	AdvanceAmount *int64 `form:"advance_amount"`
	// Specifies which fields in the response should be expanded.
	Expand        []*string `form:"expand"`
	FeeAmount     *int64    `form:"fee_amount"`
	FinancingType *string   `form:"financing_type"`
	Status        *string   `form:"status"`
	WithholdRate  *float64  `form:"withhold_rate"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersCapitalFinancingOfferParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Refills a test financing offer for a connected account.
type TestHelpersCapitalFinancingOfferRefillParams struct {
	Params        `form:"*"`
	AdvanceAmount *int64 `form:"advance_amount"`
	// Specifies which fields in the response should be expanded.
	Expand        []*string `form:"expand"`
	FeeAmount     *int64    `form:"fee_amount"`
	FinancingType *string   `form:"financing_type"`
	Status        *string   `form:"status"`
	WithholdRate  *float64  `form:"withhold_rate"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersCapitalFinancingOfferRefillParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a test financing offer for a connected account.
type TestHelpersCapitalFinancingOfferCreateParams struct {
	Params        `form:"*"`
	AdvanceAmount *int64 `form:"advance_amount"`
	// Specifies which fields in the response should be expanded.
	Expand        []*string `form:"expand"`
	FeeAmount     *int64    `form:"fee_amount"`
	FinancingType *string   `form:"financing_type"`
	Status        *string   `form:"status"`
	WithholdRate  *float64  `form:"withhold_rate"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersCapitalFinancingOfferCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
