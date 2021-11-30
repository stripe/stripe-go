//
//
// File generated from our OpenAPI spec
//
//

package stripe

// A string identifying the type of the source of this line item, either an `invoiceitem` or a `subscription`.
type InvoiceLineType string

// List of values that InvoiceLineType can take
const (
	InvoiceLineTypeInvoiceItem  InvoiceLineType = "invoiceitem"
	InvoiceLineTypeSubscription InvoiceLineType = "subscription"
)

// The amount of discount calculated per discount for this line item.
type InvoiceLineDiscountAmount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}
type InvoiceLine struct {
	Amount           int64                        `json:"amount"`
	Currency         Currency                     `json:"currency"`
	Description      string                       `json:"description"`
	Discountable     bool                         `json:"discountable"`
	DiscountAmounts  []*InvoiceLineDiscountAmount `json:"discount_amounts"`
	Discounts        []*Discount                  `json:"discounts"`
	ID               string                       `json:"id"`
	InvoiceItem      string                       `json:"invoice_item"`
	Livemode         bool                         `json:"livemode"`
	Metadata         map[string]string            `json:"metadata"`
	Object           string                       `json:"object"`
	Period           *Period                      `json:"period"`
	Plan             *Plan                        `json:"plan"`
	Price            *Price                       `json:"price"`
	Proration        bool                         `json:"proration"`
	Quantity         int64                        `json:"quantity"`
	Subscription     string                       `json:"subscription"`
	SubscriptionItem string                       `json:"subscription_item"`
	TaxAmounts       []*InvoiceTaxAmount          `json:"tax_amounts"`
	TaxRates         []*TaxRate                   `json:"tax_rates"`
	Type             InvoiceLineType              `json:"type"`
	UnifiedProration bool                         `json:"unified_proration"`
}

// Period is a structure representing a start and end dates.
type Period struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}

// InvoiceLineList is a list of InvoiceLineItems as retrieved from a list endpoint.
type InvoiceLineList struct {
	APIResource
	ListMeta
	Data []*InvoiceLine `json:"data"`
}
