//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of the credit note line item, one of `invoice_line_item` or `custom_line_item`. When the type is `invoice_line_item` there is an additional `invoice_line_item` property on the resource the value of which is the id of the credited line item on the invoice.
type CreditNoteLineItemType string

// List of values that CreditNoteLineItemType can take
const (
	CreditNoteLineItemTypeCustomLineItem  CreditNoteLineItemType = "custom_line_item"
	CreditNoteLineItemTypeInvoiceLineItem CreditNoteLineItemType = "invoice_line_item"
)

// The integer amount in %s representing the discount being credited for this line item.
type CreditNoteLineItemDiscountAmount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// CreditNoteLineItem is the resource representing a Stripe credit note line item.
// For more details see https://stripe.com/docs/api/credit_notes/line_item
type CreditNoteLineItem struct {
	Amount            int64                               `json:"amount"`
	Description       string                              `json:"description"`
	DiscountAmount    int64                               `json:"discount_amount"`
	DiscountAmounts   []*CreditNoteLineItemDiscountAmount `json:"discount_amounts"`
	ID                string                              `json:"id"`
	InvoiceLineItem   string                              `json:"invoice_line_item"`
	Livemode          bool                                `json:"livemode"`
	Object            string                              `json:"object"`
	Quantity          int64                               `json:"quantity"`
	TaxAmounts        []*CreditNoteTaxAmount              `json:"tax_amounts"`
	TaxRates          []*TaxRate                          `json:"tax_rates"`
	Type              CreditNoteLineItemType              `json:"type"`
	UnitAmount        int64                               `json:"unit_amount"`
	UnitAmountDecimal float64                             `json:"unit_amount_decimal,string"`
}

// CreditNoteLineItemList is a list of CreditNoteLineItems as retrieved from a list endpoint.
type CreditNoteLineItemList struct {
	APIResource
	ListMeta
	Data []*CreditNoteLineItem `json:"data"`
}
