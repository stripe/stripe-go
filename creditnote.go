//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Type of the pretax credit amount referenced.
type CreditNotePretaxCreditAmountType string

// List of values that CreditNotePretaxCreditAmountType can take
const (
	CreditNotePretaxCreditAmountTypeCreditBalanceTransaction CreditNotePretaxCreditAmountType = "credit_balance_transaction"
	CreditNotePretaxCreditAmountTypeDiscount                 CreditNotePretaxCreditAmountType = "discount"
)

// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
type CreditNoteReason string

// List of values that CreditNoteReason can take
const (
	CreditNoteReasonDuplicate             CreditNoteReason = "duplicate"
	CreditNoteReasonFraudulent            CreditNoteReason = "fraudulent"
	CreditNoteReasonOrderChange           CreditNoteReason = "order_change"
	CreditNoteReasonProductUnsatisfactory CreditNoteReason = "product_unsatisfactory"
)

// Type of the refund, one of `refund` or `payment_record_refund`.
type CreditNoteRefundType string

// List of values that CreditNoteRefundType can take
const (
	CreditNoteRefundTypePaymentRecordRefund CreditNoteRefundType = "payment_record_refund"
	CreditNoteRefundTypeRefund              CreditNoteRefundType = "refund"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type CreditNoteShippingCostTaxTaxabilityReason string

// List of values that CreditNoteShippingCostTaxTaxabilityReason can take
const (
	CreditNoteShippingCostTaxTaxabilityReasonCustomerExempt       CreditNoteShippingCostTaxTaxabilityReason = "customer_exempt"
	CreditNoteShippingCostTaxTaxabilityReasonNotCollecting        CreditNoteShippingCostTaxTaxabilityReason = "not_collecting"
	CreditNoteShippingCostTaxTaxabilityReasonNotSubjectToTax      CreditNoteShippingCostTaxTaxabilityReason = "not_subject_to_tax"
	CreditNoteShippingCostTaxTaxabilityReasonNotSupported         CreditNoteShippingCostTaxTaxabilityReason = "not_supported"
	CreditNoteShippingCostTaxTaxabilityReasonPortionProductExempt CreditNoteShippingCostTaxTaxabilityReason = "portion_product_exempt"
	CreditNoteShippingCostTaxTaxabilityReasonPortionReducedRated  CreditNoteShippingCostTaxTaxabilityReason = "portion_reduced_rated"
	CreditNoteShippingCostTaxTaxabilityReasonPortionStandardRated CreditNoteShippingCostTaxTaxabilityReason = "portion_standard_rated"
	CreditNoteShippingCostTaxTaxabilityReasonProductExempt        CreditNoteShippingCostTaxTaxabilityReason = "product_exempt"
	CreditNoteShippingCostTaxTaxabilityReasonProductExemptHoliday CreditNoteShippingCostTaxTaxabilityReason = "product_exempt_holiday"
	CreditNoteShippingCostTaxTaxabilityReasonProportionallyRated  CreditNoteShippingCostTaxTaxabilityReason = "proportionally_rated"
	CreditNoteShippingCostTaxTaxabilityReasonReducedRated         CreditNoteShippingCostTaxTaxabilityReason = "reduced_rated"
	CreditNoteShippingCostTaxTaxabilityReasonReverseCharge        CreditNoteShippingCostTaxTaxabilityReason = "reverse_charge"
	CreditNoteShippingCostTaxTaxabilityReasonStandardRated        CreditNoteShippingCostTaxTaxabilityReason = "standard_rated"
	CreditNoteShippingCostTaxTaxabilityReasonTaxableBasisReduced  CreditNoteShippingCostTaxTaxabilityReason = "taxable_basis_reduced"
	CreditNoteShippingCostTaxTaxabilityReasonZeroRated            CreditNoteShippingCostTaxTaxabilityReason = "zero_rated"
)

// Status of this credit note, one of `issued` or `void`. Learn more about [voiding credit notes](https://docs.stripe.com/billing/invoices/credit-notes#voiding).
type CreditNoteStatus string

// List of values that CreditNoteStatus can take
const (
	CreditNoteStatusIssued CreditNoteStatus = "issued"
	CreditNoteStatusVoid   CreditNoteStatus = "void"
)

// Whether this tax is inclusive or exclusive.
type CreditNoteTotalTaxTaxBehavior string

// List of values that CreditNoteTotalTaxTaxBehavior can take
const (
	CreditNoteTotalTaxTaxBehaviorExclusive CreditNoteTotalTaxTaxBehavior = "exclusive"
	CreditNoteTotalTaxTaxBehaviorInclusive CreditNoteTotalTaxTaxBehavior = "inclusive"
)

// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
type CreditNoteTotalTaxTaxabilityReason string

// List of values that CreditNoteTotalTaxTaxabilityReason can take
const (
	CreditNoteTotalTaxTaxabilityReasonCustomerExempt       CreditNoteTotalTaxTaxabilityReason = "customer_exempt"
	CreditNoteTotalTaxTaxabilityReasonNotAvailable         CreditNoteTotalTaxTaxabilityReason = "not_available"
	CreditNoteTotalTaxTaxabilityReasonNotCollecting        CreditNoteTotalTaxTaxabilityReason = "not_collecting"
	CreditNoteTotalTaxTaxabilityReasonNotSubjectToTax      CreditNoteTotalTaxTaxabilityReason = "not_subject_to_tax"
	CreditNoteTotalTaxTaxabilityReasonNotSupported         CreditNoteTotalTaxTaxabilityReason = "not_supported"
	CreditNoteTotalTaxTaxabilityReasonPortionProductExempt CreditNoteTotalTaxTaxabilityReason = "portion_product_exempt"
	CreditNoteTotalTaxTaxabilityReasonPortionReducedRated  CreditNoteTotalTaxTaxabilityReason = "portion_reduced_rated"
	CreditNoteTotalTaxTaxabilityReasonPortionStandardRated CreditNoteTotalTaxTaxabilityReason = "portion_standard_rated"
	CreditNoteTotalTaxTaxabilityReasonProductExempt        CreditNoteTotalTaxTaxabilityReason = "product_exempt"
	CreditNoteTotalTaxTaxabilityReasonProductExemptHoliday CreditNoteTotalTaxTaxabilityReason = "product_exempt_holiday"
	CreditNoteTotalTaxTaxabilityReasonProportionallyRated  CreditNoteTotalTaxTaxabilityReason = "proportionally_rated"
	CreditNoteTotalTaxTaxabilityReasonReducedRated         CreditNoteTotalTaxTaxabilityReason = "reduced_rated"
	CreditNoteTotalTaxTaxabilityReasonReverseCharge        CreditNoteTotalTaxTaxabilityReason = "reverse_charge"
	CreditNoteTotalTaxTaxabilityReasonStandardRated        CreditNoteTotalTaxTaxabilityReason = "standard_rated"
	CreditNoteTotalTaxTaxabilityReasonTaxableBasisReduced  CreditNoteTotalTaxTaxabilityReason = "taxable_basis_reduced"
	CreditNoteTotalTaxTaxabilityReasonZeroRated            CreditNoteTotalTaxTaxabilityReason = "zero_rated"
)

// The type of tax information.
type CreditNoteTotalTaxType string

// List of values that CreditNoteTotalTaxType can take
const (
	CreditNoteTotalTaxTypeTaxRateDetails CreditNoteTotalTaxType = "tax_rate_details"
)

// Type of this credit note, one of `pre_payment` or `post_payment`. A `pre_payment` credit note means it was issued when the invoice was open. A `post_payment` credit note means it was issued when the invoice was paid.
type CreditNoteType string

// List of values that CreditNoteType can take
const (
	CreditNoteTypeMixed       CreditNoteType = "mixed"
	CreditNoteTypePostPayment CreditNoteType = "post_payment"
	CreditNoteTypePrePayment  CreditNoteType = "pre_payment"
)

// Returns a list of credit notes.
type CreditNoteListParams struct {
	ListParams `form:"*"`
	// Only return credit notes that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return credit notes that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return credit notes for the customer specified by this customer ID.
	Customer *string `form:"customer"`
	// Only return credit notes for the account representing the customer specified by this account ID.
	CustomerAccount *string `form:"customer_account"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return credit notes for the invoice specified by this invoice ID.
	Invoice *string `form:"invoice"`
}

// AddExpand appends a new field to expand.
func (p *CreditNoteListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A list of up to 10 tax amounts for the credit note line item. Not valid when `tax_rates` is used or if invoice is set up with `automatic_tax[enabled]=true`.
type CreditNoteLineTaxAmountParams struct {
	// The amount, in cents (or local equivalent), of the tax.
	Amount *int64 `form:"amount"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount *int64 `form:"taxable_amount"`
	// The id of the tax rate for this tax amount. The tax rate must have been automatically created by Stripe.
	TaxRate *string `form:"tax_rate"`
}

// Line items that make up the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
type CreditNoteLineParams struct {
	// The line item amount to credit. Only valid when `type` is `invoice_line_item`. If invoice is set up with `automatic_tax[enabled]=true`, this amount is tax exclusive
	Amount *int64 `form:"amount"`
	// The description of the credit note line item. Only valid when the `type` is `custom_line_item`.
	Description *string `form:"description"`
	// The invoice line item to credit. Only valid when the `type` is `invoice_line_item`.
	InvoiceLineItem *string `form:"invoice_line_item"`
	// The line item quantity to credit.
	Quantity *int64 `form:"quantity"`
	// A list of up to 10 tax amounts for the credit note line item. Not valid when `tax_rates` is used or if invoice is set up with `automatic_tax[enabled]=true`.
	TaxAmounts []*CreditNoteLineTaxAmountParams `form:"tax_amounts"`
	// The tax rates which apply to the credit note line item. Only valid when the `type` is `custom_line_item` and `tax_amounts` is not used.
	TaxRates []*string `form:"tax_rates"`
	// Type of the credit note line item, one of `invoice_line_item` or `custom_line_item`. `custom_line_item` is not valid when the invoice is set up with `automatic_tax[enabled]=true`.
	Type *string `form:"type"`
	// The integer unit amount in cents (or local equivalent) of the credit note line item. This `unit_amount` will be multiplied by the quantity to get the full amount to credit for this line item. Only valid when `type` is `custom_line_item`.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// The PaymentRecord refund details to link to this credit note. Required when `type` is `payment_record_refund`.
type CreditNoteRefundPaymentRecordRefundParams struct {
	// The ID of the PaymentRecord with the refund to link to this credit note.
	PaymentRecord *string `form:"payment_record"`
	// The PaymentRecord refund group to link to this credit note. For refunds processed off-Stripe, this will correspond to the `processor_details.custom.refund_reference` field provided when reporting the refund on the PaymentRecord.
	RefundGroup *string `form:"refund_group"`
}

// Refunds to link to this credit note.
type CreditNoteRefundParams struct {
	// Amount of the refund that applies to this credit note, in cents (or local equivalent). Defaults to the entire refund amount.
	AmountRefunded *int64 `form:"amount_refunded"`
	// The PaymentRecord refund details to link to this credit note. Required when `type` is `payment_record_refund`.
	PaymentRecordRefund *CreditNoteRefundPaymentRecordRefundParams `form:"payment_record_refund"`
	// ID of an existing refund to link this credit note to. Required when `type` is `refund`.
	Refund *string `form:"refund"`
	// Type of the refund, one of `refund` or `payment_record_refund`. Defaults to `refund`.
	Type *string `form:"type"`
}

// When shipping_cost contains the shipping_rate from the invoice, the shipping_cost is included in the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
type CreditNoteShippingCostParams struct {
	// The ID of the shipping rate to use for this order.
	ShippingRate *string `form:"shipping_rate"`
}

// Issue a credit note to adjust the amount of a finalized invoice. A credit note will first reduce the invoice's amount_remaining (and amount_due), but not below zero.
// This amount is indicated by the credit note's pre_payment_amount. The excess amount is indicated by post_payment_amount, and it can result in any combination of the following:
//
// Refunds: create a new refund (using refund_amount) or link existing refunds (using refunds).
// Customer balance credit: credit the customer's balance (using credit_amount) which will be automatically applied to their next invoice when it's finalized.
// Outside of Stripe credit: record the amount that is or will be credited outside of Stripe (using out_of_band_amount).
//
// The sum of refunds, customer balance credits, and outside of Stripe credits must equal the post_payment_amount.
//
// You may issue multiple credit notes for an invoice. Each credit note may increment the invoice's pre_payment_credit_notes_amount,
// post_payment_credit_notes_amount, or both, depending on the invoice's amount_remaining at the time of credit note creation.
type CreditNoteParams struct {
	Params `form:"*"`
	// The integer amount in cents (or local equivalent) representing the total amount of the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	Amount *int64 `form:"amount"`
	// The integer amount in cents (or local equivalent) representing the amount to credit the customer's balance, which will be automatically applied to their next invoice.
	CreditAmount *int64 `form:"credit_amount"`
	// The date when this credit note is in effect. Same as `created` unless overwritten. When defined, this value replaces the system-generated 'Date of issue' printed on the credit note PDF.
	EffectiveAt *int64 `form:"effective_at"`
	// Type of email to send to the customer, one of `credit_note` or `none` and the default is `credit_note`.
	EmailType *string `form:"email_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// ID of the invoice.
	Invoice *string `form:"invoice"`
	// Line items that make up the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	Lines []*CreditNoteLineParams `form:"lines"`
	// The credit note's memo appears on the credit note PDF.
	Memo *string `form:"memo"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The integer amount in cents (or local equivalent) representing the amount that is credited outside of Stripe.
	OutOfBandAmount *int64 `form:"out_of_band_amount"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason *string `form:"reason"`
	// The integer amount in cents (or local equivalent) representing the amount to refund. If set, a refund will be created for the charge associated with the invoice.
	RefundAmount *int64 `form:"refund_amount"`
	// Refunds to link to this credit note.
	Refunds []*CreditNoteRefundParams `form:"refunds"`
	// When shipping_cost contains the shipping_rate from the invoice, the shipping_cost is included in the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	ShippingCost *CreditNoteShippingCostParams `form:"shipping_cost"`
}

// AddExpand appends a new field to expand.
func (p *CreditNoteParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CreditNoteParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A list of up to 10 tax amounts for the credit note line item. Not valid when `tax_rates` is used or if invoice is set up with `automatic_tax[enabled]=true`.
type CreditNotePreviewLineTaxAmountParams struct {
	// The amount, in cents (or local equivalent), of the tax.
	Amount *int64 `form:"amount"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount *int64 `form:"taxable_amount"`
	// The id of the tax rate for this tax amount. The tax rate must have been automatically created by Stripe.
	TaxRate *string `form:"tax_rate"`
}

// Line items that make up the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
type CreditNotePreviewLineParams struct {
	// The line item amount to credit. Only valid when `type` is `invoice_line_item`. If invoice is set up with `automatic_tax[enabled]=true`, this amount is tax exclusive
	Amount *int64 `form:"amount"`
	// The description of the credit note line item. Only valid when the `type` is `custom_line_item`.
	Description *string `form:"description"`
	// The invoice line item to credit. Only valid when the `type` is `invoice_line_item`.
	InvoiceLineItem *string `form:"invoice_line_item"`
	// The line item quantity to credit.
	Quantity *int64 `form:"quantity"`
	// A list of up to 10 tax amounts for the credit note line item. Not valid when `tax_rates` is used or if invoice is set up with `automatic_tax[enabled]=true`.
	TaxAmounts []*CreditNotePreviewLineTaxAmountParams `form:"tax_amounts"`
	// The tax rates which apply to the credit note line item. Only valid when the `type` is `custom_line_item` and `tax_amounts` is not used.
	TaxRates []*string `form:"tax_rates"`
	// Type of the credit note line item, one of `invoice_line_item` or `custom_line_item`. `custom_line_item` is not valid when the invoice is set up with `automatic_tax[enabled]=true`.
	Type *string `form:"type"`
	// The integer unit amount in cents (or local equivalent) of the credit note line item. This `unit_amount` will be multiplied by the quantity to get the full amount to credit for this line item. Only valid when `type` is `custom_line_item`.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// The PaymentRecord refund details to link to this credit note. Required when `type` is `payment_record_refund`.
type CreditNotePreviewRefundPaymentRecordRefundParams struct {
	// The ID of the PaymentRecord with the refund to link to this credit note.
	PaymentRecord *string `form:"payment_record"`
	// The PaymentRecord refund group to link to this credit note. For refunds processed off-Stripe, this will correspond to the `processor_details.custom.refund_reference` field provided when reporting the refund on the PaymentRecord.
	RefundGroup *string `form:"refund_group"`
}

// Refunds to link to this credit note.
type CreditNotePreviewRefundParams struct {
	// Amount of the refund that applies to this credit note, in cents (or local equivalent). Defaults to the entire refund amount.
	AmountRefunded *int64 `form:"amount_refunded"`
	// The PaymentRecord refund details to link to this credit note. Required when `type` is `payment_record_refund`.
	PaymentRecordRefund *CreditNotePreviewRefundPaymentRecordRefundParams `form:"payment_record_refund"`
	// ID of an existing refund to link this credit note to. Required when `type` is `refund`.
	Refund *string `form:"refund"`
	// Type of the refund, one of `refund` or `payment_record_refund`. Defaults to `refund`.
	Type *string `form:"type"`
}

// When shipping_cost contains the shipping_rate from the invoice, the shipping_cost is included in the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
type CreditNotePreviewShippingCostParams struct {
	// The ID of the shipping rate to use for this order.
	ShippingRate *string `form:"shipping_rate"`
}

// Get a preview of a credit note without creating it.
type CreditNotePreviewParams struct {
	Params `form:"*"`
	// The integer amount in cents (or local equivalent) representing the total amount of the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	Amount *int64 `form:"amount"`
	// The integer amount in cents (or local equivalent) representing the amount to credit the customer's balance, which will be automatically applied to their next invoice.
	CreditAmount *int64 `form:"credit_amount"`
	// The date when this credit note is in effect. Same as `created` unless overwritten. When defined, this value replaces the system-generated 'Date of issue' printed on the credit note PDF.
	EffectiveAt *int64 `form:"effective_at"`
	// Type of email to send to the customer, one of `credit_note` or `none` and the default is `credit_note`.
	EmailType *string `form:"email_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// ID of the invoice.
	Invoice *string `form:"invoice"`
	// Line items that make up the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	Lines []*CreditNotePreviewLineParams `form:"lines"`
	// The credit note's memo appears on the credit note PDF.
	Memo *string `form:"memo"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The integer amount in cents (or local equivalent) representing the amount that is credited outside of Stripe.
	OutOfBandAmount *int64 `form:"out_of_band_amount"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason *string `form:"reason"`
	// The integer amount in cents (or local equivalent) representing the amount to refund. If set, a refund will be created for the charge associated with the invoice.
	RefundAmount *int64 `form:"refund_amount"`
	// Refunds to link to this credit note.
	Refunds []*CreditNotePreviewRefundParams `form:"refunds"`
	// When shipping_cost contains the shipping_rate from the invoice, the shipping_cost is included in the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	ShippingCost *CreditNotePreviewShippingCostParams `form:"shipping_cost"`
}

// AddExpand appends a new field to expand.
func (p *CreditNotePreviewParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CreditNotePreviewParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A list of up to 10 tax amounts for the credit note line item. Not valid when `tax_rates` is used or if invoice is set up with `automatic_tax[enabled]=true`.
type CreditNotePreviewLinesLineTaxAmountParams struct {
	// The amount, in cents (or local equivalent), of the tax.
	Amount *int64 `form:"amount"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount *int64 `form:"taxable_amount"`
	// The id of the tax rate for this tax amount. The tax rate must have been automatically created by Stripe.
	TaxRate *string `form:"tax_rate"`
}

// Line items that make up the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
type CreditNotePreviewLinesLineParams struct {
	// The line item amount to credit. Only valid when `type` is `invoice_line_item`. If invoice is set up with `automatic_tax[enabled]=true`, this amount is tax exclusive
	Amount *int64 `form:"amount"`
	// The description of the credit note line item. Only valid when the `type` is `custom_line_item`.
	Description *string `form:"description"`
	// The invoice line item to credit. Only valid when the `type` is `invoice_line_item`.
	InvoiceLineItem *string `form:"invoice_line_item"`
	// The line item quantity to credit.
	Quantity *int64 `form:"quantity"`
	// A list of up to 10 tax amounts for the credit note line item. Not valid when `tax_rates` is used or if invoice is set up with `automatic_tax[enabled]=true`.
	TaxAmounts []*CreditNotePreviewLinesLineTaxAmountParams `form:"tax_amounts"`
	// The tax rates which apply to the credit note line item. Only valid when the `type` is `custom_line_item` and `tax_amounts` is not used.
	TaxRates []*string `form:"tax_rates"`
	// Type of the credit note line item, one of `invoice_line_item` or `custom_line_item`. `custom_line_item` is not valid when the invoice is set up with `automatic_tax[enabled]=true`.
	Type *string `form:"type"`
	// The integer unit amount in cents (or local equivalent) of the credit note line item. This `unit_amount` will be multiplied by the quantity to get the full amount to credit for this line item. Only valid when `type` is `custom_line_item`.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// The PaymentRecord refund details to link to this credit note. Required when `type` is `payment_record_refund`.
type CreditNotePreviewLinesRefundPaymentRecordRefundParams struct {
	// The ID of the PaymentRecord with the refund to link to this credit note.
	PaymentRecord *string `form:"payment_record"`
	// The PaymentRecord refund group to link to this credit note. For refunds processed off-Stripe, this will correspond to the `processor_details.custom.refund_reference` field provided when reporting the refund on the PaymentRecord.
	RefundGroup *string `form:"refund_group"`
}

// Refunds to link to this credit note.
type CreditNotePreviewLinesRefundParams struct {
	// Amount of the refund that applies to this credit note, in cents (or local equivalent). Defaults to the entire refund amount.
	AmountRefunded *int64 `form:"amount_refunded"`
	// The PaymentRecord refund details to link to this credit note. Required when `type` is `payment_record_refund`.
	PaymentRecordRefund *CreditNotePreviewLinesRefundPaymentRecordRefundParams `form:"payment_record_refund"`
	// ID of an existing refund to link this credit note to. Required when `type` is `refund`.
	Refund *string `form:"refund"`
	// Type of the refund, one of `refund` or `payment_record_refund`. Defaults to `refund`.
	Type *string `form:"type"`
}

// When shipping_cost contains the shipping_rate from the invoice, the shipping_cost is included in the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
type CreditNotePreviewLinesShippingCostParams struct {
	// The ID of the shipping rate to use for this order.
	ShippingRate *string `form:"shipping_rate"`
}

// Line items that make up the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
type CreditNotePreviewLinesParams struct {
	ListParams `form:"*"`
	// The integer amount in cents (or local equivalent) representing the total amount of the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	Amount *int64 `form:"amount"`
	// The integer amount in cents (or local equivalent) representing the amount to credit the customer's balance, which will be automatically applied to their next invoice.
	CreditAmount *int64 `form:"credit_amount"`
	// The date when this credit note is in effect. Same as `created` unless overwritten. When defined, this value replaces the system-generated 'Date of issue' printed on the credit note PDF.
	EffectiveAt *int64 `form:"effective_at"`
	// Type of email to send to the customer, one of `credit_note` or `none` and the default is `credit_note`.
	EmailType *string `form:"email_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// ID of the invoice.
	Invoice *string `form:"invoice"`
	// Line items that make up the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	Lines []*CreditNotePreviewLinesLineParams `form:"lines"`
	// The credit note's memo appears on the credit note PDF.
	Memo *string `form:"memo"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The integer amount in cents (or local equivalent) representing the amount that is credited outside of Stripe.
	OutOfBandAmount *int64 `form:"out_of_band_amount"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason *string `form:"reason"`
	// The integer amount in cents (or local equivalent) representing the amount to refund. If set, a refund will be created for the charge associated with the invoice.
	RefundAmount *int64 `form:"refund_amount"`
	// Refunds to link to this credit note.
	Refunds []*CreditNotePreviewLinesRefundParams `form:"refunds"`
	// When shipping_cost contains the shipping_rate from the invoice, the shipping_cost is included in the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	ShippingCost *CreditNotePreviewLinesShippingCostParams `form:"shipping_cost"`
}

// AddExpand appends a new field to expand.
func (p *CreditNotePreviewLinesParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CreditNotePreviewLinesParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Marks a credit note as void. Learn more about [voiding credit notes](https://docs.stripe.com/docs/billing/invoices/credit-notes#voiding).
type CreditNoteVoidCreditNoteParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CreditNoteVoidCreditNoteParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// When retrieving a credit note, you'll get a lines property containing the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type CreditNoteListLinesParams struct {
	ListParams `form:"*"`
	CreditNote *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CreditNoteListLinesParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A list of up to 10 tax amounts for the credit note line item. Not valid when `tax_rates` is used or if invoice is set up with `automatic_tax[enabled]=true`.
type CreditNoteCreateLineTaxAmountParams struct {
	// The amount, in cents (or local equivalent), of the tax.
	Amount *int64 `form:"amount"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount *int64 `form:"taxable_amount"`
	// The id of the tax rate for this tax amount. The tax rate must have been automatically created by Stripe.
	TaxRate *string `form:"tax_rate"`
}

// Line items that make up the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
type CreditNoteCreateLineParams struct {
	// The line item amount to credit. Only valid when `type` is `invoice_line_item`. If invoice is set up with `automatic_tax[enabled]=true`, this amount is tax exclusive
	Amount *int64 `form:"amount"`
	// The description of the credit note line item. Only valid when the `type` is `custom_line_item`.
	Description *string `form:"description"`
	// The invoice line item to credit. Only valid when the `type` is `invoice_line_item`.
	InvoiceLineItem *string `form:"invoice_line_item"`
	// The line item quantity to credit.
	Quantity *int64 `form:"quantity"`
	// A list of up to 10 tax amounts for the credit note line item. Not valid when `tax_rates` is used or if invoice is set up with `automatic_tax[enabled]=true`.
	TaxAmounts []*CreditNoteCreateLineTaxAmountParams `form:"tax_amounts"`
	// The tax rates which apply to the credit note line item. Only valid when the `type` is `custom_line_item` and `tax_amounts` is not used.
	TaxRates []*string `form:"tax_rates"`
	// Type of the credit note line item, one of `invoice_line_item` or `custom_line_item`. `custom_line_item` is not valid when the invoice is set up with `automatic_tax[enabled]=true`.
	Type *string `form:"type"`
	// The integer unit amount in cents (or local equivalent) of the credit note line item. This `unit_amount` will be multiplied by the quantity to get the full amount to credit for this line item. Only valid when `type` is `custom_line_item`.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// The PaymentRecord refund details to link to this credit note. Required when `type` is `payment_record_refund`.
type CreditNoteCreateRefundPaymentRecordRefundParams struct {
	// The ID of the PaymentRecord with the refund to link to this credit note.
	PaymentRecord *string `form:"payment_record"`
	// The PaymentRecord refund group to link to this credit note. For refunds processed off-Stripe, this will correspond to the `processor_details.custom.refund_reference` field provided when reporting the refund on the PaymentRecord.
	RefundGroup *string `form:"refund_group"`
}

// Refunds to link to this credit note.
type CreditNoteCreateRefundParams struct {
	// Amount of the refund that applies to this credit note, in cents (or local equivalent). Defaults to the entire refund amount.
	AmountRefunded *int64 `form:"amount_refunded"`
	// The PaymentRecord refund details to link to this credit note. Required when `type` is `payment_record_refund`.
	PaymentRecordRefund *CreditNoteCreateRefundPaymentRecordRefundParams `form:"payment_record_refund"`
	// ID of an existing refund to link this credit note to. Required when `type` is `refund`.
	Refund *string `form:"refund"`
	// Type of the refund, one of `refund` or `payment_record_refund`. Defaults to `refund`.
	Type *string `form:"type"`
}

// When shipping_cost contains the shipping_rate from the invoice, the shipping_cost is included in the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
type CreditNoteCreateShippingCostParams struct {
	// The ID of the shipping rate to use for this order.
	ShippingRate *string `form:"shipping_rate"`
}

// Issue a credit note to adjust the amount of a finalized invoice. A credit note will first reduce the invoice's amount_remaining (and amount_due), but not below zero.
// This amount is indicated by the credit note's pre_payment_amount. The excess amount is indicated by post_payment_amount, and it can result in any combination of the following:
//
// Refunds: create a new refund (using refund_amount) or link existing refunds (using refunds).
// Customer balance credit: credit the customer's balance (using credit_amount) which will be automatically applied to their next invoice when it's finalized.
// Outside of Stripe credit: record the amount that is or will be credited outside of Stripe (using out_of_band_amount).
//
// The sum of refunds, customer balance credits, and outside of Stripe credits must equal the post_payment_amount.
//
// You may issue multiple credit notes for an invoice. Each credit note may increment the invoice's pre_payment_credit_notes_amount,
// post_payment_credit_notes_amount, or both, depending on the invoice's amount_remaining at the time of credit note creation.
type CreditNoteCreateParams struct {
	Params `form:"*"`
	// The integer amount in cents (or local equivalent) representing the total amount of the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	Amount *int64 `form:"amount"`
	// The integer amount in cents (or local equivalent) representing the amount to credit the customer's balance, which will be automatically applied to their next invoice.
	CreditAmount *int64 `form:"credit_amount"`
	// The date when this credit note is in effect. Same as `created` unless overwritten. When defined, this value replaces the system-generated 'Date of issue' printed on the credit note PDF.
	EffectiveAt *int64 `form:"effective_at"`
	// Type of email to send to the customer, one of `credit_note` or `none` and the default is `credit_note`.
	EmailType *string `form:"email_type"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// ID of the invoice.
	Invoice *string `form:"invoice"`
	// Line items that make up the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	Lines []*CreditNoteCreateLineParams `form:"lines"`
	// The credit note's memo appears on the credit note PDF.
	Memo *string `form:"memo"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The integer amount in cents (or local equivalent) representing the amount that is credited outside of Stripe.
	OutOfBandAmount *int64 `form:"out_of_band_amount"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason *string `form:"reason"`
	// The integer amount in cents (or local equivalent) representing the amount to refund. If set, a refund will be created for the charge associated with the invoice.
	RefundAmount *int64 `form:"refund_amount"`
	// Refunds to link to this credit note.
	Refunds []*CreditNoteCreateRefundParams `form:"refunds"`
	// When shipping_cost contains the shipping_rate from the invoice, the shipping_cost is included in the credit note. One of `amount`, `lines`, or `shipping_cost` must be provided.
	ShippingCost *CreditNoteCreateShippingCostParams `form:"shipping_cost"`
}

// AddExpand appends a new field to expand.
func (p *CreditNoteCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CreditNoteCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the credit note object with the given identifier.
type CreditNoteRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CreditNoteRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates an existing credit note.
type CreditNoteUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Credit note memo.
	Memo *string `form:"memo"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *CreditNoteUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CreditNoteUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The integer amount in cents (or local equivalent) representing the total amount of discount that was credited.
type CreditNoteDiscountAmount struct {
	// The amount, in cents (or local equivalent), of the discount.
	Amount int64 `json:"amount"`
	// The discount that was applied to get this discount amount.
	Discount *Discount `json:"discount"`
}

// The pretax credit amounts (ex: discount, credit grants, etc) for all line items.
type CreditNotePretaxCreditAmount struct {
	// The amount, in cents (or local equivalent), of the pretax credit amount.
	Amount int64 `json:"amount"`
	// The credit balance transaction that was applied to get this pretax credit amount.
	CreditBalanceTransaction *BillingCreditBalanceTransaction `json:"credit_balance_transaction"`
	// The discount that was applied to get this pretax credit amount.
	Discount *Discount `json:"discount"`
	// Type of the pretax credit amount referenced.
	Type CreditNotePretaxCreditAmountType `json:"type"`
}

// The PaymentRecord refund details associated with this credit note refund.
type CreditNoteRefundPaymentRecordRefund struct {
	// ID of the payment record.
	PaymentRecord string `json:"payment_record"`
	// ID of the refund group.
	RefundGroup string `json:"refund_group"`
}

// Refunds related to this credit note.
type CreditNoteRefund struct {
	// Amount of the refund that applies to this credit note, in cents (or local equivalent).
	AmountRefunded int64 `json:"amount_refunded"`
	// The PaymentRecord refund details associated with this credit note refund.
	PaymentRecordRefund *CreditNoteRefundPaymentRecordRefund `json:"payment_record_refund"`
	// ID of the refund.
	Refund *Refund `json:"refund"`
	// Type of the refund, one of `refund` or `payment_record_refund`.
	Type CreditNoteRefundType `json:"type"`
}

// The taxes applied to the shipping rate.
type CreditNoteShippingCostTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://docs.stripe.com/invoicing/taxes/tax-rates), [subscriptions](https://docs.stripe.com/billing/taxes/tax-rates) and [Checkout Sessions](https://docs.stripe.com/payments/checkout/use-manual-tax-rates) to collect tax.
	//
	// Related guide: [Tax rates](https://docs.stripe.com/billing/taxes/tax-rates)
	Rate *TaxRate `json:"rate"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason CreditNoteShippingCostTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
}

// The details of the cost of shipping, including the ShippingRate applied to the invoice.
type CreditNoteShippingCost struct {
	// Total shipping cost before any taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total tax amount applied due to shipping costs. If no tax was applied, defaults to 0.
	AmountTax int64 `json:"amount_tax"`
	// Total shipping cost after taxes are applied.
	AmountTotal int64 `json:"amount_total"`
	// The ID of the ShippingRate for this invoice.
	ShippingRate *ShippingRate `json:"shipping_rate"`
	// The taxes applied to the shipping rate.
	Taxes []*CreditNoteShippingCostTax `json:"taxes"`
}

// Additional details about the tax rate. Only present when `type` is `tax_rate_details`.
type CreditNoteTotalTaxTaxRateDetails struct {
	// ID of the tax rate
	TaxRate string `json:"tax_rate"`
}

// The aggregate tax information for all line items.
type CreditNoteTotalTax struct {
	// The amount of the tax, in cents (or local equivalent).
	Amount int64 `json:"amount"`
	// The reasoning behind this tax, for example, if the product is tax exempt. The possible values for this field may be extended as new tax rules are supported.
	TaxabilityReason CreditNoteTotalTaxTaxabilityReason `json:"taxability_reason"`
	// The amount on which tax is calculated, in cents (or local equivalent).
	TaxableAmount int64 `json:"taxable_amount"`
	// Whether this tax is inclusive or exclusive.
	TaxBehavior CreditNoteTotalTaxTaxBehavior `json:"tax_behavior"`
	// Additional details about the tax rate. Only present when `type` is `tax_rate_details`.
	TaxRateDetails *CreditNoteTotalTaxTaxRateDetails `json:"tax_rate_details"`
	// The type of tax information.
	Type CreditNoteTotalTaxType `json:"type"`
}

// Issue a credit note to adjust an invoice's amount after the invoice is finalized.
//
// Related guide: [Credit notes](https://docs.stripe.com/billing/invoices/credit-notes)
type CreditNote struct {
	APIResource
	// The integer amount in cents (or local equivalent) representing the total amount of the credit note, including tax.
	Amount int64 `json:"amount"`
	// This is the sum of all the shipping amounts.
	AmountShipping int64 `json:"amount_shipping"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// ID of the customer.
	Customer *Customer `json:"customer"`
	// ID of the account representing the customer.
	CustomerAccount string `json:"customer_account"`
	// Customer balance transaction related to this credit note.
	CustomerBalanceTransaction *CustomerBalanceTransaction `json:"customer_balance_transaction"`
	// The integer amount in cents (or local equivalent) representing the total amount of discount that was credited.
	DiscountAmount int64 `json:"discount_amount"`
	// The aggregate amounts calculated per discount for all line items.
	DiscountAmounts []*CreditNoteDiscountAmount `json:"discount_amounts"`
	// The date when this credit note is in effect. Same as `created` unless overwritten. When defined, this value replaces the system-generated 'Date of issue' printed on the credit note PDF.
	EffectiveAt int64 `json:"effective_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// ID of the invoice.
	Invoice *Invoice `json:"invoice"`
	// Line items that make up the credit note
	Lines *CreditNoteLineItemList `json:"lines"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Customer-facing text that appears on the credit note PDF.
	Memo string `json:"memo"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// A unique number that identifies this particular credit note and appears on the PDF of the credit note and its associated invoice.
	Number string `json:"number"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Amount that was credited outside of Stripe.
	OutOfBandAmount int64 `json:"out_of_band_amount"`
	// The link to download the PDF of the credit note.
	PDF string `json:"pdf"`
	// The amount of the credit note that was refunded to the customer, credited to the customer's balance, credited outside of Stripe, or any combination thereof.
	PostPaymentAmount int64 `json:"post_payment_amount"`
	// The amount of the credit note by which the invoice's `amount_remaining` and `amount_due` were reduced.
	PrePaymentAmount int64 `json:"pre_payment_amount"`
	// The pretax credit amounts (ex: discount, credit grants, etc) for all line items.
	PretaxCreditAmounts []*CreditNotePretaxCreditAmount `json:"pretax_credit_amounts"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason CreditNoteReason `json:"reason"`
	// Refunds related to this credit note.
	Refunds []*CreditNoteRefund `json:"refunds"`
	// The details of the cost of shipping, including the ShippingRate applied to the invoice.
	ShippingCost *CreditNoteShippingCost `json:"shipping_cost"`
	// Status of this credit note, one of `issued` or `void`. Learn more about [voiding credit notes](https://docs.stripe.com/billing/invoices/credit-notes#voiding).
	Status CreditNoteStatus `json:"status"`
	// The integer amount in cents (or local equivalent) representing the amount of the credit note, excluding exclusive tax and invoice level discounts.
	Subtotal int64 `json:"subtotal"`
	// The integer amount in cents (or local equivalent) representing the amount of the credit note, excluding all tax and invoice level discounts.
	SubtotalExcludingTax int64 `json:"subtotal_excluding_tax"`
	// The integer amount in cents (or local equivalent) representing the total amount of the credit note, including tax and all discount.
	Total int64 `json:"total"`
	// The integer amount in cents (or local equivalent) representing the total amount of the credit note, excluding tax, but including discounts.
	TotalExcludingTax int64 `json:"total_excluding_tax"`
	// The aggregate tax information for all line items.
	TotalTaxes []*CreditNoteTotalTax `json:"total_taxes"`
	// Type of this credit note, one of `pre_payment` or `post_payment`. A `pre_payment` credit note means it was issued when the invoice was open. A `post_payment` credit note means it was issued when the invoice was paid.
	Type CreditNoteType `json:"type"`
	// The time that the credit note was voided.
	VoidedAt int64 `json:"voided_at"`
}

// CreditNoteList is a list of CreditNotes as retrieved from a list endpoint.
type CreditNoteList struct {
	APIResource
	ListMeta
	Data []*CreditNote `json:"data"`
}

// UnmarshalJSON handles deserialization of a CreditNote.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *CreditNote) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type creditNote CreditNote
	var v creditNote
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = CreditNote(v)
	return nil
}
