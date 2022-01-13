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
	Amount            *int64    `form:"amount"`
	Description       *string   `form:"description"`
	InvoiceLineItem   *string   `form:"invoice_line_item"`
	Quantity          *int64    `form:"quantity"`
	TaxRates          []*string `form:"tax_rates"`
	Type              *string   `form:"type"`
	UnitAmount        *int64    `form:"unit_amount"`
	UnitAmountDecimal *float64  `form:"unit_amount_decimal,high_precision"`
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
	Params          `form:"*"`
	Amount          *int64                  `form:"amount"`
	CreditAmount    *int64                  `form:"credit_amount"`
	Invoice         *string                 `form:"invoice"`
	Lines           []*CreditNoteLineParams `form:"lines"`
	Memo            *string                 `form:"memo"`
	OutOfBandAmount *int64                  `form:"out_of_band_amount"`
	Reason          *string                 `form:"reason"`
	Refund          *string                 `form:"refund"`
	RefundAmount    *int64                  `form:"refund_amount"`
}

// Returns a list of credit notes.
type CreditNoteListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"customer"`
	Invoice    *string `form:"invoice"`
}

// Get a preview of a credit note without creating it.
type CreditNotePreviewParams struct {
	Params          `form:"*"`
	Amount          *int64                  `form:"amount"`
	CreditAmount    *int64                  `form:"credit_amount"`
	Invoice         *string                 `form:"invoice"`
	Lines           []*CreditNoteLineParams `form:"lines"`
	Memo            *string                 `form:"memo"`
	OutOfBandAmount *int64                  `form:"out_of_band_amount"`
	Reason          *string                 `form:"reason"`
	Refund          *string                 `form:"refund"`
	RefundAmount    *int64                  `form:"refund_amount"`
}

// Marks a credit note as void. Learn more about [voiding credit notes](https://stripe.com/docs/billing/invoices/credit-notes#voiding).
type CreditNoteVoidParams struct {
	Params `form:"*"`
}

// When retrieving a credit note preview, you'll get a lines property containing the first handful of those items. This URL you can retrieve the full (paginated) list of line items.
type CreditNoteLineItemListPreviewParams struct {
	ListParams      `form:"*"`
	Amount          *int64                  `form:"amount"`
	CreditAmount    *int64                  `form:"credit_amount"`
	Invoice         *string                 `form:"invoice"`
	Lines           []*CreditNoteLineParams `form:"lines"`
	Memo            *string                 `form:"memo"`
	OutOfBandAmount *int64                  `form:"out_of_band_amount"`
	Reason          *string                 `form:"reason"`
	Refund          *string                 `form:"refund"`
	RefundAmount    *int64                  `form:"refund_amount"`
}

// When retrieving a credit note, you'll get a lines property containing the the first handful of those items. There is also a URL where you can retrieve the full (paginated) list of line items.
type CreditNoteLineItemListParams struct {
	ListParams `form:"*"`
	// ID is the credit note ID to list line items for.
	ID *string `form:"-"` // Included in URL
}

// The integer amount in %s representing the total amount of discount that was credited.
type CreditNoteDiscountAmount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// The aggregate amounts calculated per tax rate for all line items.
type CreditNoteTaxAmount struct {
	Amount    int64    `json:"amount"`
	Inclusive bool     `json:"inclusive"`
	TaxRate   *TaxRate `json:"tax_rate"`
}

// Issue a credit note to adjust an invoice's amount after the invoice is finalized.
//
// Related guide: [Credit Notes](https://stripe.com/docs/billing/invoices/credit-notes).
type CreditNote struct {
	APIResource
	Amount                     int64                       `json:"amount"`
	Created                    int64                       `json:"created"`
	Currency                   Currency                    `json:"currency"`
	Customer                   *Customer                   `json:"customer"`
	CustomerBalanceTransaction *CustomerBalanceTransaction `json:"customer_balance_transaction"`
	DiscountAmount             int64                       `json:"discount_amount"`
	DiscountAmounts            []*CreditNoteDiscountAmount `json:"discount_amounts"`
	ID                         string                      `json:"id"`
	Invoice                    *Invoice                    `json:"invoice"`
	Lines                      *CreditNoteLineItemList     `json:"lines"`
	Livemode                   bool                        `json:"livemode"`
	Memo                       string                      `json:"memo"`
	Metadata                   map[string]string           `json:"metadata"`
	Number                     string                      `json:"number"`
	Object                     string                      `json:"object"`
	OutOfBandAmount            int64                       `json:"out_of_band_amount"`
	PDF                        string                      `json:"pdf"`
	Reason                     CreditNoteReason            `json:"reason"`
	Refund                     *Refund                     `json:"refund"`
	Status                     CreditNoteStatus            `json:"status"`
	Subtotal                   int64                       `json:"subtotal"`
	TaxAmounts                 []*CreditNoteTaxAmount      `json:"tax_amounts"`
	Total                      int64                       `json:"total"`
	Type                       CreditNoteType              `json:"type"`
	VoidedAt                   int64                       `json:"voided_at"`
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
