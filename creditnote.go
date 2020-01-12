package stripe

import "encoding/json"

// CreditNoteReason is the reason why a given credit note was created.
type CreditNoteReason string

// List of values that CreditNoteReason can take.
const (
	CreditNoteReasonDuplicate             CreditNoteReason = "duplicate"
	CreditNoteReasonFraudulent            CreditNoteReason = "fraudulent"
	CreditNoteReasonOrderChange           CreditNoteReason = "order_change"
	CreditNoteReasonProductUnsatisfactory CreditNoteReason = "product_unsatisfactory"
)

// CreditNoteStatus is the list of allowed values for the credit note's status.
type CreditNoteStatus string

// List of values that CreditNoteStatus can take.
const (
	CreditNoteStatusIssued CreditNoteStatus = "issued"
	CreditNoteStatusVoid   CreditNoteStatus = "void"
)

// CreditNoteType is the list of allowed values for the credit note's type.
type CreditNoteType string

// List of values that CreditNoteType can take.
const (
	CreditNoteTypePostPayment CreditNoteType = "post_payment"
	CreditNoteTypePrePayment  CreditNoteType = "pre_payment"
)

// CreditNoteLineItemType is the list of allowed values for the credit note line item's type.
type CreditNoteLineItemType string

// List of values that CreditNoteType can take.
const (
	CreditNoteLineItemTypeCustomLineItem  CreditNoteLineItemType = "custom_line_item"
	CreditNoteLineItemTypeInvoiceLineItem CreditNoteLineItemType = "invoice_line_item"
)

// CreditNoteParams is the set of parameters that can be used when creating or updating a credit note.
// For more details see https://stripe.com/docs/api/credit_notes/create, https://stripe.com/docs/api/credit_notes/update.
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

// CreditNoteListParams is the set of parameters that can be used when listing credit notes.
// For more details see https://stripe.com/docs/api/credit_notes/list.
type CreditNoteListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"customer"`
	Invoice    *string `form:"invoice"`
}

// CreditNoteLineItemListParams is the set of parameters that can be used when listing credit note line items.
type CreditNoteLineItemListParams struct {
	ListParams `form:"*"`

	// ID is the credit note ID to list line items for.
	ID *string `form:"-"` // Goes in the URL
}

// CreditNoteLineItemListPreviewParams is the set of parameters that can be used when previewing a credit note's line items
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

// CreditNoteLineParams is the set of parameters that can be used for a line item when creating
// or previewing a credit note.
type CreditNoteLineParams struct {
	Amount            *int64    `form:"amount"`
	Description       *string   `form:"description"`
	InvoiceLineItem   *string   `form:"invoice_line_item"`
	Quantity          *int64    `form:"quantity"`
	TaxRates          []*string `form:"tax_rates"`
	UnitAmount        *int64    `form:"unit_amount"`
	UnitAmountDecimal *float64  `form:"unit_amount_decimal,high_precision"`
	Type              *string   `form:"type"`
}

// CreditNotePreviewParams is the set of parameters that can be used when previewing a credit note.
// For more details see https://stripe.com/docs/api/credit_notes/preview.
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

// CreditNoteVoidParams is the set of parameters that can be used when voiding invoices.
type CreditNoteVoidParams struct {
	Params `form:"*"`
}

// CreditNoteTaxAmount represent the tax amount applied to a credit note.
type CreditNoteTaxAmount struct {
	Amount    int64    `json:"amount"`
	Inclusive bool     `json:"inclusive"`
	TaxRate   *TaxRate `json:"tax_rate"`
}

// CreditNote is the resource representing a Stripe credit note.
// For more details see https://stripe.com/docs/api/credit_notes/object.
type CreditNote struct {
	Amount                     int64                       `json:"amount"`
	Created                    int64                       `json:"created"`
	Currency                   Currency                    `json:"currency"`
	Customer                   *Customer                   `json:"customer"`
	CustomerBalanceTransaction *CustomerBalanceTransaction `json:"customer_balance_transaction"`
	DiscountAmount             int64                       `json:"discount_amount"`
	Invoice                    *Invoice                    `json:"invoice"`
	ID                         string                      `json:"id"`
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

// CreditNoteLineItem is the resource representing a Stripe credit note line item.
// For more details see https://stripe.com/docs/api/credit_notes/line_item
type CreditNoteLineItem struct {
	Amount            int64                  `json:"amount"`
	Description       string                 `json:"description"`
	DiscountAmount    int64                  `json:"discount_amount"`
	ID                string                 `json:"id"`
	InvoiceLineItem   string                 `json:"invoice_line_item"`
	Livemode          bool                   `json:"livemode"`
	Object            string                 `json:"object"`
	Quantity          int64                  `json:"quantity"`
	TaxAmounts        []*CreditNoteTaxAmount `json:"tax_amounts"`
	TaxRates          []*TaxRate             `json:"tax_rates"`
	Type              CreditNoteLineItemType `json:"type"`
	UnitAmount        int64                  `json:"unit_amount"`
	UnitAmountDecimal float64                `json:"unit_amount_decimal,string"`
}

// CreditNoteList is a list of credit notes as retrieved from a list endpoint.
type CreditNoteList struct {
	ListMeta
	Data []*CreditNote `json:"data"`
}

// CreditNoteLineItemList is a list of credit note line items as retrieved from a list endpoint.
type CreditNoteLineItemList struct {
	ListMeta
	Data []*CreditNoteLineItem `json:"data"`
}

// UnmarshalJSON handles deserialization of a CreditNote.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *CreditNote) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type note CreditNote
	var v note
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = CreditNote(v)
	return nil
}
