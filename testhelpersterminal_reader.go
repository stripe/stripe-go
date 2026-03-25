//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Simulated data for the card payment method.
type TestHelpersTerminalReaderPresentPaymentMethodCardParams struct {
	// Card security code.
	CVC *string `form:"cvc" json:"cvc,omitempty"`
	// Two-digit number representing the card's expiration month.
	ExpMonth *int64 `form:"exp_month" json:"exp_month"`
	// Two- or four-digit number representing the card's expiration year.
	ExpYear *int64 `form:"exp_year" json:"exp_year"`
	// The card number, as a string without any separators.
	Number *string `form:"number" json:"number"`
}

// Simulated data for the card_present payment method.
type TestHelpersTerminalReaderPresentPaymentMethodCardPresentParams struct {
	// The card number, as a string without any separators.
	Number *string `form:"number" json:"number,omitempty"`
}

// Simulated data for the interac_present payment method.
type TestHelpersTerminalReaderPresentPaymentMethodInteracPresentParams struct {
	// The Interac card number.
	Number *string `form:"number" json:"number,omitempty"`
}

// Presents a payment method on a simulated reader. Can be used to simulate accepting a payment, saving a card or refunding a transaction.
type TestHelpersTerminalReaderPresentPaymentMethodParams struct {
	Params `form:"*"`
	// Simulated on-reader tip amount.
	AmountTip *int64 `form:"amount_tip" json:"amount_tip,omitempty"`
	// Simulated data for the card payment method.
	Card *TestHelpersTerminalReaderPresentPaymentMethodCardParams `form:"card" json:"card,omitempty"`
	// Simulated data for the card_present payment method.
	CardPresent *TestHelpersTerminalReaderPresentPaymentMethodCardPresentParams `form:"card_present" json:"card_present,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Simulated data for the interac_present payment method.
	InteracPresent *TestHelpersTerminalReaderPresentPaymentMethodInteracPresentParams `form:"interac_present" json:"interac_present,omitempty"`
	// Simulated payment type.
	Type *string `form:"type" json:"type,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTerminalReaderPresentPaymentMethodParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Use this endpoint to trigger a successful input collection on a simulated reader.
type TestHelpersTerminalReaderSucceedInputCollectionParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// This parameter defines the skip behavior for input collection.
	SkipNonRequiredInputs *string `form:"skip_non_required_inputs" json:"skip_non_required_inputs,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTerminalReaderSucceedInputCollectionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Use this endpoint to complete an input collection with a timeout error on a simulated reader.
type TestHelpersTerminalReaderTimeoutInputCollectionParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTerminalReaderTimeoutInputCollectionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
