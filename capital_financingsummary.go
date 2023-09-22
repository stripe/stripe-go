//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Status of the Connected Account's financing. [/v1/capital/financing_summary](https://stripe.com/docs/api/capital/financing_summary) will only return `details` for `paid_out` financing.
type CapitalFinancingSummaryStatus string

// List of values that CapitalFinancingSummaryStatus can take
const (
	CapitalFinancingSummaryStatusAccepted  CapitalFinancingSummaryStatus = "accepted"
	CapitalFinancingSummaryStatusDelivered CapitalFinancingSummaryStatus = "delivered"
	CapitalFinancingSummaryStatusNone      CapitalFinancingSummaryStatus = "none"
)

// Retrieve the financing state for the account that was authenticated in the request.
type CapitalFinancingSummaryParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CapitalFinancingSummaryParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The chronologically current repayment interval for the financing offer.
type CapitalFinancingSummaryDetailsCurrentRepaymentInterval struct {
	// The time at which the minimum payment amount will be due. If not met through withholding, the Connected account's linked bank account will be debited.
	// Given in seconds since unix epoch.
	DueAt int64 `json:"due_at"`
	// The amount that has already been paid in the current repayment interval.
	PaidAmount int64 `json:"paid_amount"`
	// The amount that is yet to be paid in the current repayment interval.
	RemainingAmount int64 `json:"remaining_amount"`
}

// Additional information about the financing summary. Describes currency, advance amount,
// fee amount, withhold rate, remaining amount, paid amount, current repayment interval,
// repayment start date, and advance payout date.
type CapitalFinancingSummaryDetails struct {
	// Amount of financing offered, in minor units.
	AdvanceAmount int64 `json:"advance_amount"`
	// The time at which the funds were paid out the the Connected account's Stripe balance. Given in milliseconds since unix epoch.
	AdvancePaidOutAt int64 `json:"advance_paid_out_at"`
	// Currency that the financing offer is transacted in. For example, `usd`.
	Currency Currency `json:"currency"`
	// The chronologically current repayment interval for the financing offer.
	CurrentRepaymentInterval *CapitalFinancingSummaryDetailsCurrentRepaymentInterval `json:"current_repayment_interval"`
	// Fixed fee amount, in minor units.
	FeeAmount int64 `json:"fee_amount"`
	// The amount the Connected account has paid toward the financing debt so far.
	PaidAmount int64 `json:"paid_amount"`
	// The balance remaining to be paid on the financing, in minor units.
	RemainingAmount int64 `json:"remaining_amount"`
	// The time at which Capital will begin withholding from payments. Given in seconds since unix epoch.
	RepaymentsBeginAt int64 `json:"repayments_begin_at"`
	// Per-transaction rate at which Stripe will withhold funds to repay the financing.
	WithholdRate float64 `json:"withhold_rate"`
}

// A financing object describes an account's current financing state. Used by Connect
// platforms to read the state of Capital offered to their connected accounts.
type CapitalFinancingSummary struct {
	APIResource
	// Additional information about the financing summary. Describes currency, advance amount,
	// fee amount, withhold rate, remaining amount, paid amount, current repayment interval,
	// repayment start date, and advance payout date.
	Details *CapitalFinancingSummaryDetails `json:"details"`
	// The Financing Offer ID this Financing Summary corresponds to
	FinancingOffer string `json:"financing_offer"`
	// The object type: financing_summary
	Object string `json:"object"`
	// Status of the Connected Account's financing. [/v1/capital/financing_summary](https://stripe.com/docs/api/capital/financing_summary) will only return `details` for `paid_out` financing.
	Status CapitalFinancingSummaryStatus `json:"status"`
}
