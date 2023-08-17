//
//
// File generated from our OpenAPI spec
//
//

package stripe

// If set to `reset`, the billing_cycle_anchor of the subscription is set to the start of the phase when entering the phase. If unset, then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
type QuotePhaseBillingCycleAnchor string

// List of values that QuotePhaseBillingCycleAnchor can take
const (
	QuotePhaseBillingCycleAnchorReset QuotePhaseBillingCycleAnchor = "reset"
)

// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
type QuotePhaseCollectionMethod string

// List of values that QuotePhaseCollectionMethod can take
const (
	QuotePhaseCollectionMethodChargeAutomatically QuotePhaseCollectionMethod = "charge_automatically"
	QuotePhaseCollectionMethodSendInvoice         QuotePhaseCollectionMethod = "send_invoice"
)

// If the quote will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
type QuotePhaseProrationBehavior string

// List of values that QuotePhaseProrationBehavior can take
const (
	QuotePhaseProrationBehaviorAlwaysInvoice    QuotePhaseProrationBehavior = "always_invoice"
	QuotePhaseProrationBehaviorCreateProrations QuotePhaseProrationBehavior = "create_prorations"
	QuotePhaseProrationBehaviorNone             QuotePhaseProrationBehavior = "none"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason string

// List of values that QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason can take
const (
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonCustomerExempt       QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "customer_exempt"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonNotCollecting        QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "not_collecting"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonNotSubjectToTax      QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "not_subject_to_tax"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonNotSupported         QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "not_supported"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonPortionProductExempt QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "portion_product_exempt"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonPortionReducedRated  QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "portion_reduced_rated"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonPortionStandardRated QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "portion_standard_rated"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonProductExempt        QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonProductExemptHoliday QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "product_exempt_holiday"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonProportionallyRated  QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "proportionally_rated"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonReducedRated         QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "reduced_rated"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonReverseCharge        QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "reverse_charge"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonStandardRated        QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "standard_rated"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonTaxableBasisReduced  QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "taxable_basis_reduced"
	QuotePhaseTotalDetailsBreakdownTaxTaxabilityReasonZeroRated            QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason = "zero_rated"
)

// Returns a list of quote phases.
type QuotePhaseListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The ID of the quote whose phases will be retrieved.
	Quote *string `form:"quote"`
}

// AddExpand appends a new field to expand.
func (p *QuotePhaseListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the quote phase with the given ID.
type QuotePhaseParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuotePhaseParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// When retrieving a quote phase, there is an includable line_items property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type QuotePhaseListLineItemsParams struct {
	ListParams `form:"*"`
	QuotePhase *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *QuotePhaseListLineItemsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The invoice settings applicable during this phase.
type QuotePhaseInvoiceSettings struct {
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int64 `json:"days_until_due"`
}

// The aggregated discounts.
type QuotePhaseTotalDetailsBreakdownDiscount struct {
	// The amount discounted.
	Amount int64 `json:"amount"`
	// A discount represents the actual application of a [coupon](https://stripe.com/docs/api#coupons) or [promotion code](https://stripe.com/docs/api#promotion_codes).
	// It contains information about when the discount began, when it will end, and what it is applied to.
	//
	// Related guide: [Applying discounts to subscriptions](https://stripe.com/docs/billing/subscriptions/discounts)
	Discount *Discount `json:"discount"`
}

// The aggregated tax amounts by rate.
type QuotePhaseTotalDetailsBreakdownTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.
	//
	// Related guide: [Tax rates](https://stripe.com/docs/billing/taxes/tax-rates)
	Rate *TaxRate `json:"rate"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason QuotePhaseTotalDetailsBreakdownTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
}
type QuotePhaseTotalDetailsBreakdown struct {
	// The aggregated discounts.
	Discounts []*QuotePhaseTotalDetailsBreakdownDiscount `json:"discounts"`
	// The aggregated tax amounts by rate.
	Taxes []*QuotePhaseTotalDetailsBreakdownTax `json:"taxes"`
}
type QuotePhaseTotalDetails struct {
	// This is the sum of all the discounts.
	AmountDiscount int64 `json:"amount_discount"`
	// This is the sum of all the shipping amounts.
	AmountShipping int64 `json:"amount_shipping"`
	// This is the sum of all the tax amounts.
	AmountTax int64                            `json:"amount_tax"`
	Breakdown *QuotePhaseTotalDetailsBreakdown `json:"breakdown"`
}

// A quote phase describes the line items, coupons, and trialing status of a subscription for a predefined time period.
type QuotePhase struct {
	APIResource
	// Total before any discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total after discounts and taxes are applied.
	AmountTotal int64 `json:"amount_total"`
	// If set to `reset`, the billing_cycle_anchor of the subscription is set to the start of the phase when entering the phase. If unset, then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor QuotePhaseBillingCycleAnchor `json:"billing_cycle_anchor"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod QuotePhaseCollectionMethod `json:"collection_method"`
	// The default tax rates to apply to the subscription during this phase of the quote.
	DefaultTaxRates []*TaxRate `json:"default_tax_rates"`
	// The stackable discounts that will be applied to the subscription on this phase. Subscription item discounts are applied before subscription discounts.
	Discounts []*Discount `json:"discounts"`
	// The end of this phase of the quote
	EndDate int64 `json:"end_date"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The invoice settings applicable during this phase.
	InvoiceSettings *QuotePhaseInvoiceSettings `json:"invoice_settings"`
	// Integer representing the multiplier applied to the price interval. For example, `iterations=2` applied to a price with `interval=month` and `interval_count=3` results in a phase of duration `2 * 3 months = 6 months`.
	Iterations int64 `json:"iterations"`
	// A list of items the customer is being quoted for.
	LineItems *LineItemList `json:"line_items"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// If the quote will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
	ProrationBehavior QuotePhaseProrationBehavior `json:"proration_behavior"`
	TotalDetails      *QuotePhaseTotalDetails     `json:"total_details"`
	// If set to true the entire phase is counted as a trial and the customer will not be charged for any recurring fees.
	Trial bool `json:"trial"`
	// When the trial ends within the phase.
	TrialEnd int64 `json:"trial_end"`
}

// QuotePhaseList is a list of QuotePhases as retrieved from a list endpoint.
type QuotePhaseList struct {
	APIResource
	ListMeta
	Data []*QuotePhase `json:"data"`
}
