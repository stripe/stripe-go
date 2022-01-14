//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The discounts applied to the line item.
type LineItemDiscount struct {
	// The amount discounted.
	Amount int64 `json:"amount"`
	// A discount represents the actual application of a coupon to a particular
	// customer. It contains information about when the discount began and when it
	// will end.
	//
	// Related guide: [Applying Discounts to Subscriptions](https://stripe.com/docs/billing/subscriptions/discounts).
	Discount *Discount `json:"discount"`
}

// The taxes applied to the line item.
type LineItemTax struct {
	// Amount of tax applied for this rate.
	Amount int64 `json:"amount"`
	// Tax rates can be applied to [invoices](https://stripe.com/docs/billing/invoices/tax-rates), [subscriptions](https://stripe.com/docs/billing/subscriptions/taxes) and [Checkout Sessions](https://stripe.com/docs/payments/checkout/set-up-a-subscription#tax-rates) to collect tax.
	//
	// Related guide: [Tax Rates](https://stripe.com/docs/billing/taxes/tax-rates).
	Rate *TaxRate `json:"rate"`

	// This property never worked. Use Rate instead.
	// TODO: Remove in the next major version
	TaxRate *TaxRate `json:"tax_rate"`
}

// A line item.
type LineItem struct {
	APIResource
	// Total before any discounts or taxes are applied.
	AmountSubtotal int64 `json:"amount_subtotal"`
	// Total after discounts and taxes.
	AmountTotal int64 `json:"amount_total"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	Deleted  bool     `json:"deleted"`
	// An arbitrary string attached to the object. Often useful for displaying to users. Defaults to product name.
	Description string `json:"description"`
	// The discounts applied to the line item.
	Discounts []*LineItemDiscount `json:"discounts"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The price used to generate the line item.
	Price *Price `json:"price"`
	// The quantity of products being purchased.
	Quantity int64 `json:"quantity"`
	// The taxes applied to the line item.
	Taxes []*LineItemTax `json:"taxes"`
}

// UnmarshalJSON handles deserialization of a LineItem.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (l *LineItem) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		l.ID = id
		return nil
	}

	type lineItem LineItem
	var v lineItem
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*l = LineItem(v)
	return nil
}

// LineItemList is a list of LineItems as retrieved from a list endpoint.
type LineItemList struct {
	APIResource
	ListMeta
	Data []*LineItem `json:"data"`
}
