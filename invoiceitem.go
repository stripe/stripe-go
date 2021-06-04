//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// InvoiceItemDiscountParams represents the parameters associated with the discounts to apply to an invoice item.
type InvoiceItemDiscountParams struct {
	Coupon   *string `form:"coupon"`
	Discount *string `form:"discount"`
}

// InvoiceItemPeriodParams represents the period associated with that invoice item.
type InvoiceItemPeriodParams struct {
	End   *int64 `form:"end"`
	Start *int64 `form:"start"`
}

// InvoiceItemPriceDataParams is a structure representing the parameters to create an inline price.
type InvoiceItemPriceDataParams struct {
	Currency          *string  `form:"currency"`
	Product           *string  `form:"product"`
	TaxBehavior       *string  `form:"tax_behavior"`
	UnitAmount        *int64   `form:"unit_amount"`
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// InvoiceItemParams is the set of parameters that can be used when creating or updating an invoice item.
// For more details see https://stripe.com/docs/api#create_invoiceitem and https://stripe.com/docs/api#update_invoiceitem.
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

// InvoiceItemListParams is the set of parameters that can be used when listing invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
type InvoiceItemListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Customer     *string           `form:"customer"`
	Invoice      *string           `form:"invoice"`
	Pending      *bool             `form:"pending"`
}

// InvoiceItem is the resource represneting a Stripe invoice item.
// For more details see https://stripe.com/docs/api#invoiceitems.
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

// InvoiceItemList is a list of invoice items as retrieved from a list endpoint.
type InvoiceItemList struct {
	APIResource
	ListMeta
	Data []*InvoiceItem `json:"data"`
}
