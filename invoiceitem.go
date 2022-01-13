//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Returns a list of your invoice items. Invoice items are returned sorted by creation date, with the most recently created invoice items appearing first.
type InvoiceItemListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Customer     *string           `form:"customer"`
	Invoice      *string           `form:"invoice"`
	Pending      *bool             `form:"pending"`
}

// The coupons to redeem into discounts for the invoice item or invoice line item.
type InvoiceItemDiscountParams struct {
	Coupon   *string `form:"coupon"`
	Discount *string `form:"discount"`
}

// The period associated with this invoice item.
type InvoiceItemPeriodParams struct {
	End   *int64 `form:"end"`
	Start *int64 `form:"start"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object inline.
type InvoiceItemPriceDataParams struct {
	Currency          *string  `form:"currency"`
	Product           *string  `form:"product"`
	TaxBehavior       *string  `form:"tax_behavior"`
	UnitAmount        *int64   `form:"unit_amount"`
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Creates an item to be added to a draft invoice (up to 250 items per invoice). If no invoice is specified, the item will be on the next invoice created for the customer specified.
type InvoiceItemParams struct {
	Params            `form:"*"`
	Amount            *int64                       `form:"amount"`
	Currency          *string                      `form:"currency"`
	Customer          *string                      `form:"customer"`
	Description       *string                      `form:"description"`
	Discountable      *bool                        `form:"discountable"`
	Discounts         []*InvoiceItemDiscountParams `form:"discounts"`
	Invoice           *string                      `form:"invoice"`
	Period            *InvoiceItemPeriodParams     `form:"period"`
	Price             *string                      `form:"price"`
	PriceData         *InvoiceItemPriceDataParams  `form:"price_data"`
	Quantity          *int64                       `form:"quantity"`
	Subscription      *string                      `form:"subscription"`
	TaxRates          []*string                    `form:"tax_rates"`
	UnitAmount        *int64                       `form:"unit_amount"`
	UnitAmountDecimal *float64                     `form:"unit_amount_decimal,high_precision"`
}

// Sometimes you want to add a charge or credit to a customer, but actually
// charge or credit the customer's card only at the end of a regular billing
// cycle. This is useful for combining several charges (to minimize
// per-transaction fees), or for having Stripe tabulate your usage-based billing
// totals.
//
// Related guide: [Subscription Invoices](https://stripe.com/docs/billing/invoices/subscription#adding-upcoming-invoice-items).
type InvoiceItem struct {
	APIResource
	Amount            int64             `json:"amount"`
	Currency          Currency          `json:"currency"`
	Customer          *Customer         `json:"customer"`
	Date              int64             `json:"date"`
	Deleted           bool              `json:"deleted"`
	Description       string            `json:"description"`
	Discountable      bool              `json:"discountable"`
	Discounts         []*Discount       `json:"discounts"`
	ID                string            `json:"id"`
	Invoice           *Invoice          `json:"invoice"`
	Livemode          bool              `json:"livemode"`
	Metadata          map[string]string `json:"metadata"`
	Object            string            `json:"object"`
	Period            *Period           `json:"period"`
	Plan              *Plan             `json:"plan"`
	Price             *Price            `json:"price"`
	Proration         bool              `json:"proration"`
	Quantity          int64             `json:"quantity"`
	Subscription      *Subscription     `json:"subscription"`
	SubscriptionItem  string            `json:"subscription_item"`
	TaxRates          []*TaxRate        `json:"tax_rates"`
	UnitAmount        int64             `json:"unit_amount"`
	UnitAmountDecimal float64           `json:"unit_amount_decimal,string"`
}

// UnmarshalJSON handles deserialization of an InvoiceItem.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *InvoiceItem) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type invoiceItem InvoiceItem
	var v invoiceItem
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = InvoiceItem(v)
	return nil
}

// InvoiceItemList is a list of InvoiceItems as retrieved from a list endpoint.
type InvoiceItemList struct {
	APIResource
	ListMeta
	Data []*InvoiceItem `json:"data"`
}
