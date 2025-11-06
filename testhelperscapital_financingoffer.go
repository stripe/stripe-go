//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a test financing offer for a connected account.
type TestHelpersCapitalFinancingOfferParams struct {
	Params `form:"*"`
	// Amount of financing offered, in minor units. For example, 1,000 USD is represented as 100000.
	AdvanceAmount *int64 `form:"advance_amount"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Fixed fee amount, in minor units. For example, 100 USD is represented as 10000.
	FeeAmount *int64 `form:"fee_amount"`
	// The type of financing offer.
	FinancingType *string `form:"financing_type"`
	// The status of the financing offer.
	Status *string `form:"status"`
	// Per-transaction rate at which Stripe withholds funds to repay the financing.
	WithholdRate *float64 `form:"withhold_rate"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersCapitalFinancingOfferParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Refills a test financing offer for a connected account.
type TestHelpersCapitalFinancingOfferRefillParams struct {
	Params `form:"*"`
	// Amount of financing offered, in minor units. For example, 1,000 USD is represented as 100000.
	AdvanceAmount *int64 `form:"advance_amount"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Fixed fee amount, in minor units. For example, 100 USD is represented as 10000.
	FeeAmount *int64 `form:"fee_amount"`
	// The type of financing offer
	FinancingType *string `form:"financing_type"`
	// The status of the financing offer
	Status *string `form:"status"`
	// Per-transaction rate at which Stripe withholds funds to repay the financing.
	WithholdRate *float64 `form:"withhold_rate"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersCapitalFinancingOfferRefillParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a test financing offer for a connected account.
type TestHelpersCapitalFinancingOfferCreateParams struct {
	Params `form:"*"`
	// Amount of financing offered, in minor units. For example, 1,000 USD is represented as 100000.
	AdvanceAmount *int64 `form:"advance_amount"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Fixed fee amount, in minor units. For example, 100 USD is represented as 10000.
	FeeAmount *int64 `form:"fee_amount"`
	// The type of financing offer.
	FinancingType *string `form:"financing_type"`
	// The status of the financing offer.
	Status *string `form:"status"`
	// Per-transaction rate at which Stripe withholds funds to repay the financing.
	WithholdRate *float64 `form:"withhold_rate"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersCapitalFinancingOfferCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
