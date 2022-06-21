//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
type CreditNoteReason string

// List of values that CreditNoteReason can take
const (
	CreditNoteReasonDuplicate             CreditNoteReason = "duplicate"
	CreditNoteReasonFraudulent            CreditNoteReason = "fraudulent"
	CreditNoteReasonOrderChange           CreditNoteReason = "order_change"
	CreditNoteReasonProductUnsatisfactory CreditNoteReason = "product_unsatisfactory"
)

// Status of this credit note, one of `issued` or `void`. Learn more about [voiding credit notes](https://stripe.com/docs/billing/invoices/credit-notes#voiding).
type CreditNoteStatus string

// List of values that CreditNoteStatus can take
const (
	CreditNoteStatusIssued CreditNoteStatus = "issued"
	CreditNoteStatusVoid   CreditNoteStatus = "void"
)

// Type of this credit note, one of `pre_payment` or `post_payment`. A `pre_payment` credit note means it was issued when the invoice was open. A `post_payment` credit note means it was issued when the invoice was paid.
type CreditNoteType string

// List of values that CreditNoteType can take
const (
	CreditNoteTypePostPayment CreditNoteType = "post_payment"
	CreditNoteTypePrePayment  CreditNoteType = "pre_payment"
)

// Line items that make up the credit note.
type CreditNoteLineParams struct {
	// The line item amount to credit. Only valid when `type` is `invoice_line_item`.
	Amount *int64 `form:"amount"`
	// The description of the credit note line item. Only valid when the `type` is `custom_line_item`.
	Description *string `form:"description"`
	// The invoice line item to credit. Only valid when the `type` is `invoice_line_item`.
	InvoiceLineItem *string `form:"invoice_line_item"`
	// The line item quantity to credit.
	Quantity *int64 `form:"quantity"`
	// The tax rates which apply to the credit note line item. Only valid when the `type` is `custom_line_item`.
	TaxRates []*string `form:"tax_rates"`
	// Type of the credit note line item, one of `invoice_line_item` or `custom_line_item`
	Type *string `form:"type"`
	// The integer unit amount in cents (or local equivalent) of the credit note line item. This `unit_amount` will be multiplied by the quantity to get the full amount to credit for this line item. Only valid when `type` is `custom_line_item`.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Issue a credit note to adjust the amount of a finalized invoice. For a status=open invoice, a credit note reduces
// its amount_due. For a status=paid invoice, a credit note does not affect its amount_due. Instead, it can result
// in any combination of the following:
//
//
// Refund: create a new refund (using refund_amount) or link an existing refund (using refund).
// Customer balance credit: credit the customer's balance (using credit_amount) which will be automatically applied to their next invoice when it's finalized.
// Outside of Stripe credit: record the amount that is or will be credited outside of Stripe (using out_of_band_amount).
//
//
// For post-payment credit notes the sum of the refund, credit and outside of Stripe amounts must equal the credit note total.
//
// You may issue multiple credit notes for an invoice. Each credit note will increment the invoice's pre_payment_credit_notes_amount
// or post_payment_credit_notes_amount depending on its status at the time of credit note creation.
type CreditNoteParams struct {
	Params `form:"*"`
	// The integer amount in cents (or local equivalent) representing the total amount of the credit note.
	Amount *int64 `form:"amount"`
	// The integer amount in cents (or local equivalent) representing the amount to credit the customer's balance, which will be automatically applied to their next invoice.
	CreditAmount *int64 `form:"credit_amount"`
	// ID of the invoice.
	Invoice *string `form:"invoice"`
	// Line items that make up the credit note.
	Lines []*CreditNoteLineParams `form:"lines"`
	// Credit note memo.
	Memo *string `form:"memo"`
	// The integer amount in cents (or local equivalent) representing the amount that is credited outside of Stripe.
	OutOfBandAmount *int64 `form:"out_of_band_amount"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason *string `form:"reason"`
	// ID of an existing refund to link this credit note to.
	Refund *string `form:"refund"`
	// The integer amount in cents (or local equivalent) representing the amount to refund. If set, a refund will be created for the charge associated with the invoice.
	RefundAmount *int64 `form:"refund_amount"`
}

// Returns a list of credit notes.
type CreditNoteListParams struct {
	ListParams `form:"*"`
	// Only return credit notes for the customer specified by this customer ID.
	Customer *string `form:"customer"`
	// Only return credit notes for the invoice specified by this invoice ID.
	Invoice *string `form:"invoice"`
}

// Get a preview of a credit note without creating it.
type CreditNotePreviewParams struct {
	Params `form:"*"`
	// The integer amount in cents (or local equivalent) representing the total amount of the credit note.
	Amount *int64 `form:"amount"`
	// The integer amount in cents (or local equivalent) representing the amount to credit the customer's balance, which will be automatically applied to their next invoice.
	CreditAmount *int64 `form:"credit_amount"`
	// ID of the invoice.
	Invoice *string `form:"invoice"`
	// Line items that make up the credit note.
	Lines []*CreditNoteLineParams `form:"lines"`
	// The credit note's memo appears on the credit note PDF.
	Memo *string `form:"memo"`
	// The integer amount in cents (or local equivalent) representing the amount that is credited outside of Stripe.
	OutOfBandAmount *int64 `form:"out_of_band_amount"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason *string `form:"reason"`
	// ID of an existing refund to link this credit note to.
	Refund *string `form:"refund"`
	// The integer amount in cents (or local equivalent) representing the amount to refund. If set, a refund will be created for the charge associated with the invoice.
	RefundAmount *int64 `form:"refund_amount"`
}

// Marks a credit note as void. Learn more about [voiding credit notes](https://stripe.com/docs/billing/invoices/credit-notes#voiding).
type CreditNoteVoidParams struct {
	Params `form:"*"`
}

// When retrieving a credit note preview, you'll get a lines property containing the first handful of those items. This URL you can retrieve the full (paginated) list of line items.
type CreditNoteLineItemListPreviewParams struct {
	ListParams `form:"*"`
	// The integer amount in cents (or local equivalent) representing the total amount of the credit note.
	Amount *int64 `form:"amount"`
	// The integer amount in cents (or local equivalent) representing the amount to credit the customer's balance, which will be automatically applied to their next invoice.
	CreditAmount *int64 `form:"credit_amount"`
	// ID of the invoice.
	Invoice *string `form:"invoice"`
	// Line items that make up the credit note.
	Lines []*CreditNoteLineParams `form:"lines"`
	// The credit note's memo appears on the credit note PDF.
	Memo *string `form:"memo"`
	// The integer amount in cents (or local equivalent) representing the amount that is credited outside of Stripe.
	OutOfBandAmount *int64 `form:"out_of_band_amount"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason *string `form:"reason"`
	// ID of an existing refund to link this credit note to.
	Refund *string `form:"refund"`
	// The integer amount in cents (or local equivalent) representing the amount to refund. If set, a refund will be created for the charge associated with the invoice.
	RefundAmount *int64 `form:"refund_amount"`
}

// When retrieving a credit note, you'll get a lines property containing the the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type CreditNoteLineItemListParams struct {
	ListParams `form:"*"`
	// ID is the credit note ID to list line items for.
	ID *string `form:"-"` // Included in URL
}

// The integer amount in %s representing the total amount of discount that was credited.
type CreditNoteDiscountAmount struct {
	// The amount, in %s, of the discount.
	Amount int64 `json:"amount"`
	// The discount that was applied to get this discount amount.
	Discount *Discount `json:"discount"`
}

// The aggregate amounts calculated per tax rate for all line items.
type CreditNoteTaxAmount struct {
	// The amount, in %s, of the tax.
	Amount int64 `json:"amount"`
	// Whether this tax amount is inclusive or exclusive.
	Inclusive bool `json:"inclusive"`
	// The tax rate that was applied to get this tax amount.
	TaxRate *TaxRate `json:"tax_rate"`
}

// Issue a credit note to adjust an invoice's amount after the invoice is finalized.
//
// Related guide: [Credit Notes](https://stripe.com/docs/billing/invoices/credit-notes).
type CreditNote struct {
	APIResource
	// The integer amount in %s representing the total amount of the credit note, including tax.
	Amount int64 `json:"amount"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// ID of the customer.
	Customer *Customer `json:"customer"`
	// Customer balance transaction related to this credit note.
	CustomerBalanceTransaction *CustomerBalanceTransaction `json:"customer_balance_transaction"`
	// The integer amount in %s representing the total amount of discount that was credited.
	DiscountAmount int64 `json:"discount_amount"`
	// The aggregate amounts calculated per discount for all line items.
	DiscountAmounts []*CreditNoteDiscountAmount `json:"discount_amounts"`
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
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// A unique number that identifies this particular credit note and appears on the PDF of the credit note and its associated invoice.
	Number string `json:"number"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Amount that was credited outside of Stripe.
	OutOfBandAmount int64 `json:"out_of_band_amount"`
	// The link to download the PDF of the credit note.
	PDF string `json:"pdf"`
	// Reason for issuing this credit note, one of `duplicate`, `fraudulent`, `order_change`, or `product_unsatisfactory`
	Reason CreditNoteReason `json:"reason"`
	// Refund related to this credit note.
	Refund *Refund `json:"refund"`
	// Status of this credit note, one of `issued` or `void`. Learn more about [voiding credit notes](https://stripe.com/docs/billing/invoices/credit-notes#voiding).
	Status CreditNoteStatus `json:"status"`
	// The integer amount in %s representing the amount of the credit note, excluding exclusive tax and invoice level discounts.
	Subtotal int64 `json:"subtotal"`
	// The integer amount in %s representing the amount of the credit note, excluding all tax and invoice level discounts.
	SubtotalExcludingTax int64 `json:"subtotal_excluding_tax"`
	// The aggregate amounts calculated per tax rate for all line items.
	TaxAmounts []*CreditNoteTaxAmount `json:"tax_amounts"`
	// The integer amount in %s representing the total amount of the credit note, including tax and all discount.
	Total int64 `json:"total"`
	// The integer amount in %s representing the total amount of the credit note, excluding tax, but including discounts.
	TotalExcludingTax int64 `json:"total_excluding_tax"`
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
