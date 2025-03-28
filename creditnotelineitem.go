//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Type of the pretax credit amount referenced.
type CreditNoteLineItemPretaxCreditAmountType string

// List of values that CreditNoteLineItemPretaxCreditAmountType can take
const (
	CreditNoteLineItemPretaxCreditAmountTypeCreditBalanceTransaction CreditNoteLineItemPretaxCreditAmountType = "credit_balance_transaction"
	CreditNoteLineItemPretaxCreditAmountTypeDiscount                 CreditNoteLineItemPretaxCreditAmountType = "discount"
)

// Whether this tax is inclusive or exclusive.
type CreditNoteLineItemTaxTaxBehavior string

// List of values that CreditNoteLineItemTaxTaxBehavior can take
const (
	CreditNoteLineItemTaxTaxBehaviorExclusive CreditNoteLineItemTaxTaxBehavior = "exclusive"
	CreditNoteLineItemTaxTaxBehaviorInclusive CreditNoteLineItemTaxTaxBehavior = "inclusive"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type CreditNoteLineItemTaxTaxabilityReason string

// List of values that CreditNoteLineItemTaxTaxabilityReason can take
const (
	CreditNoteLineItemTaxTaxabilityReasonCustomerExempt       CreditNoteLineItemTaxTaxabilityReason = "customer_exempt"
	CreditNoteLineItemTaxTaxabilityReasonNotAvailable         CreditNoteLineItemTaxTaxabilityReason = "not_available"
	CreditNoteLineItemTaxTaxabilityReasonNotCollecting        CreditNoteLineItemTaxTaxabilityReason = "not_collecting"
	CreditNoteLineItemTaxTaxabilityReasonNotSubjectToTax      CreditNoteLineItemTaxTaxabilityReason = "not_subject_to_tax"
	CreditNoteLineItemTaxTaxabilityReasonNotSupported         CreditNoteLineItemTaxTaxabilityReason = "not_supported"
	CreditNoteLineItemTaxTaxabilityReasonPortionProductExempt CreditNoteLineItemTaxTaxabilityReason = "portion_product_exempt"
	CreditNoteLineItemTaxTaxabilityReasonPortionReducedRated  CreditNoteLineItemTaxTaxabilityReason = "portion_reduced_rated"
	CreditNoteLineItemTaxTaxabilityReasonPortionStandardRated CreditNoteLineItemTaxTaxabilityReason = "portion_standard_rated"
	CreditNoteLineItemTaxTaxabilityReasonProductExempt        CreditNoteLineItemTaxTaxabilityReason = "product_exempt"
	CreditNoteLineItemTaxTaxabilityReasonProductExemptHoliday CreditNoteLineItemTaxTaxabilityReason = "product_exempt_holiday"
	CreditNoteLineItemTaxTaxabilityReasonProportionallyRated  CreditNoteLineItemTaxTaxabilityReason = "proportionally_rated"
	CreditNoteLineItemTaxTaxabilityReasonReducedRated         CreditNoteLineItemTaxTaxabilityReason = "reduced_rated"
	CreditNoteLineItemTaxTaxabilityReasonReverseCharge        CreditNoteLineItemTaxTaxabilityReason = "reverse_charge"
	CreditNoteLineItemTaxTaxabilityReasonStandardRated        CreditNoteLineItemTaxTaxabilityReason = "standard_rated"
	CreditNoteLineItemTaxTaxabilityReasonTaxableBasisReduced  CreditNoteLineItemTaxTaxabilityReason = "taxable_basis_reduced"
	CreditNoteLineItemTaxTaxabilityReasonZeroRated            CreditNoteLineItemTaxTaxabilityReason = "zero_rated"
)

// The type of tax information.
type CreditNoteLineItemTaxType string

// List of values that CreditNoteLineItemTaxType can take
const (
	CreditNoteLineItemTaxTypeTaxRateDetails CreditNoteLineItemTaxType = "tax_rate_details"
)

// The type of the credit note line item, one of `invoice_line_item` or `custom_line_item`. When the type is `invoice_line_item` there is an additional `invoice_line_item` property on the resource the value of which is the id of the credited line item on the invoice.
type CreditNoteLineItemType string

// List of values that CreditNoteLineItemType can take
const (
	CreditNoteLineItemTypeCustomLineItem  CreditNoteLineItemType = "custom_line_item"
	CreditNoteLineItemTypeInvoiceLineItem CreditNoteLineItemType = "invoice_line_item"
)

// The integer amount in cents (or local equivalent) representing the discount being credited for this line item.
type CreditNoteLineItemDiscountAmount struct {
	// The amount, in cents (or local equivalent), of the discount.
	Amount int64 `json:"amount"`
	// The discount that was applied to get this discount amount.
	Discount *Discount `json:"discount"`
}

// The pretax credit amounts (ex: discount, credit grants, etc) for this line item.
type CreditNoteLineItemPretaxCreditAmount struct {
	// The amount, in cents (or local equivalent), of the pretax credit amount.
	Amount int64 `json:"amount"`
	// The credit balance transaction that was applied to get this pretax credit amount.
	CreditBalanceTransaction *BillingCreditBalanceTransaction `json:"credit_balance_transaction"`
	// The discount that was applied to get this pretax credit amount.
	Discount *Discount `json:"discount"`
	// Type of the pretax credit amount referenced.
	Type CreditNoteLineItemPretaxCreditAmountType `json:"type"`
}

// Additional details about the tax rate. Only present when `type` is `tax_rate_details`.
type CreditNoteLineItemTaxTaxRateDetails struct {
	TaxRate string `json:"tax_rate"`
}

// The tax information of the line item.
type CreditNoteLineItemTax struct {
	// The amount of the tax, in cents (or local equivalent).
	Amount int64 `json:"amount"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason CreditNoteLineItemTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
	// Whether this tax is inclusive or exclusive.
	TaxBehavior CreditNoteLineItemTaxTaxBehavior `json:"tax_behavior"`
	// Additional details about the tax rate. Only present when `type` is `tax_rate_details`.
	TaxRateDetails *CreditNoteLineItemTaxTaxRateDetails `json:"tax_rate_details"`
	// The type of tax information.
	Type CreditNoteLineItemTaxType `json:"type"`
}

// CreditNoteLineItem is the resource representing a Stripe credit note line item.
// For more details see https://stripe.com/docs/api/credit_notes/line_item
// The credit note line item object
type CreditNoteLineItem struct {
	// The integer amount in cents (or local equivalent) representing the gross amount being credited for this line item, excluding (exclusive) tax and discounts.
	Amount int64 `json:"amount"`
	// Description of the item being credited.
	Description string `json:"description"`
	// The integer amount in cents (or local equivalent) representing the discount being credited for this line item.
	DiscountAmount int64 `json:"discount_amount"`
	// The amount of discount calculated per discount for this line item
	DiscountAmounts []*CreditNoteLineItemDiscountAmount `json:"discount_amounts"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// ID of the invoice line item being credited
	InvoiceLineItem string `json:"invoice_line_item"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The pretax credit amounts (ex: discount, credit grants, etc) for this line item.
	PretaxCreditAmounts []*CreditNoteLineItemPretaxCreditAmount `json:"pretax_credit_amounts"`
	// The number of units of product being credited.
	Quantity int64 `json:"quantity"`
	// The tax information of the line item.
	Taxes []*CreditNoteLineItemTax `json:"taxes"`
	// The tax rates which apply to the line item.
	TaxRates []*TaxRate `json:"tax_rates"`
	// The type of the credit note line item, one of `invoice_line_item` or `custom_line_item`. When the type is `invoice_line_item` there is an additional `invoice_line_item` property on the resource the value of which is the id of the credited line item on the invoice.
	Type CreditNoteLineItemType `json:"type"`
	// The cost of each unit of product being credited.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
}

// CreditNoteLineItemList is a list of CreditNoteLineItems as retrieved from a list endpoint.
type CreditNoteLineItemList struct {
	APIResource
	ListMeta
	Data []*CreditNoteLineItem `json:"data"`
}
