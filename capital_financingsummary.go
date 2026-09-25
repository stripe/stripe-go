//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of disclaimer to use for a financing offer in user-facing surfaces. The corresponding disclaimer text to use for each disclaimer_variant value can be found in the [regulatory compliance docs](https://docs.stripe.com/capital/regulatory-compliance).
type CapitalFinancingSummaryDetailsDisclaimerVariant string

// List of values that CapitalFinancingSummaryDetailsDisclaimerVariant can take
const (
	CapitalFinancingSummaryDetailsDisclaimerVariantCelticUSLoan       CapitalFinancingSummaryDetailsDisclaimerVariant = "celtic_us_loan"
	CapitalFinancingSummaryDetailsDisclaimerVariantFundboxAuFinancing CapitalFinancingSummaryDetailsDisclaimerVariant = "fundbox_au_financing"
	CapitalFinancingSummaryDetailsDisclaimerVariantFundboxCaFinancing CapitalFinancingSummaryDetailsDisclaimerVariant = "fundbox_ca_financing"
	CapitalFinancingSummaryDetailsDisclaimerVariantYoulendDEFinancing CapitalFinancingSummaryDetailsDisclaimerVariant = "youlend_de_financing"
	CapitalFinancingSummaryDetailsDisclaimerVariantYoulendFRFinancing CapitalFinancingSummaryDetailsDisclaimerVariant = "youlend_fr_financing"
	CapitalFinancingSummaryDetailsDisclaimerVariantYoulendUkMca       CapitalFinancingSummaryDetailsDisclaimerVariant = "youlend_uk_mca"
	CapitalFinancingSummaryDetailsDisclaimerVariantYoulendUSMca       CapitalFinancingSummaryDetailsDisclaimerVariant = "youlend_us_mca"
)

// The ways the connected account can pay toward its financing(s).
type CapitalFinancingSummaryDetailsEnabledPaymentType string

// List of values that CapitalFinancingSummaryDetailsEnabledPaymentType can take
const (
	CapitalFinancingSummaryDetailsEnabledPaymentTypeAutomaticDebits CapitalFinancingSummaryDetailsEnabledPaymentType = "automatic_debits"
	CapitalFinancingSummaryDetailsEnabledPaymentTypeManualPayment   CapitalFinancingSummaryDetailsEnabledPaymentType = "manual_payment"
	CapitalFinancingSummaryDetailsEnabledPaymentTypeWithholding     CapitalFinancingSummaryDetailsEnabledPaymentType = "withholding"
)

// The financing status of the connected account.
type CapitalFinancingSummaryStatus string

// List of values that CapitalFinancingSummaryStatus can take
const (
	CapitalFinancingSummaryStatusAccepted  CapitalFinancingSummaryStatus = "accepted"
	CapitalFinancingSummaryStatusDelivered CapitalFinancingSummaryStatus = "delivered"
	CapitalFinancingSummaryStatusNone      CapitalFinancingSummaryStatus = "none"
)

// Retrieve the financing summary object for the account.
type CapitalFinancingSummaryParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CapitalFinancingSummaryParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieve the financing summary object for the account.
type CapitalFinancingSummaryRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CapitalFinancingSummaryRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The chronologically current repayment interval for the financing offer.
type CapitalFinancingSummaryDetailsCurrentRepaymentInterval struct {
	// The time at which the minimum payment amount will be due. If not met through withholding, the Connected account's linked bank account or account balance will be debited.
	// Given in seconds since unix epoch.
	DueAt float64 `json:"due_at"`
	// The balance for the current repayment interval, in minor units. This does not account for any amount paid down during the interval.
	IncrementalIntervalTargetAmount int64 `json:"incremental_interval_target_amount,omitempty"`
	// The amount that has already been paid in the current repayment interval, in minor units. For example, 100 USD is represented as 10000.
	PaidAmount int64 `json:"paid_amount"`
	// The amount that is yet to be paid in the current repayment interval, in minor units. For example, 100 USD is represented as 10000.
	RemainingAmount int64 `json:"remaining_amount"`
	// The time at which the current repayment interval started. Given in seconds since unix epoch.
	StartsAt int64 `json:"starts_at,omitempty"`
}

// Additional information about the financing summary. Describes currency, advance amount,
// fee amount, withhold rate, remaining amount, paid amount, current repayment interval,
// repayment start date, and advance payout date.
//
// Only present for financing offers with a `status` other than `none`.
type CapitalFinancingSummaryDetails struct {
	// Amount of financing offered, in minor units. For example, 1,000 USD is represented as 100000.
	AdvanceAmount int64 `json:"advance_amount"`
	// The time at which the funds were paid out to the connected account's Stripe balance. Given in milliseconds since unix epoch.
	AdvancePaidOutAt float64 `json:"advance_paid_out_at"`
	// Currency that the financing offer is transacted in. For example, `usd`.
	Currency Currency `json:"currency"`
	// The chronologically current repayment interval for the financing offer.
	CurrentRepaymentInterval *CapitalFinancingSummaryDetailsCurrentRepaymentInterval `json:"current_repayment_interval"`
	// The type of disclaimer to use for a financing offer in user-facing surfaces. The corresponding disclaimer text to use for each disclaimer_variant value can be found in the [regulatory compliance docs](https://docs.stripe.com/capital/regulatory-compliance).
	DisclaimerVariant CapitalFinancingSummaryDetailsDisclaimerVariant `json:"disclaimer_variant,omitempty"`
	// The ways the connected account can pay toward its financing(s).
	EnabledPaymentTypes []CapitalFinancingSummaryDetailsEnabledPaymentType `json:"enabled_payment_types,omitempty"`
	// Fixed fee amount, in minor units. For example, 100 USD is represented as 10000.
	FeeAmount int64 `json:"fee_amount"`
	// The amount the Connected account has paid toward the financing debt so far, in minor units. For example, 1,000 USD is represented as 100000.
	PaidAmount int64 `json:"paid_amount"`
	// The balance remaining to be paid on the financing, in minor units. For example, 1,000 USD is represented as 100000.
	RemainingAmount int64 `json:"remaining_amount"`
	// The time at which Capital will begin withholding from payments. Given in seconds since unix epoch.
	RepaymentsBeginAt float64 `json:"repayments_begin_at"`
	// Total amount to be paid, independent of what's already been paid, in minor units. For example, 100 USD is represented as 10000.
	TotalDueAmount int64 `json:"total_due_amount,omitempty"`
	// Per-transaction rate at which Stripe withholds funds to repay the financing.
	WithholdRate float64 `json:"withhold_rate"`
}

// A financing summary object describes a connected account's financing details in real time.
type CapitalFinancingSummary struct {
	APIResource
	// Additional information about the financing summary. Describes currency, advance amount,
	// fee amount, withhold rate, remaining amount, paid amount, current repayment interval,
	// repayment start date, and advance payout date.
	//
	// Only present for financing offers with a `status` other than `none`.
	Details *CapitalFinancingSummaryDetails `json:"details"`
	// The unique identifier of the Financing Offer object that corresponds to the Financing Summary object.
	FinancingOffer string `json:"financing_offer"`
	// The object type: financing_summary
	Object string `json:"object"`
	// The financing status of the connected account.
	Status CapitalFinancingSummaryStatus `json:"status"`
}
