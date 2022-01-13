//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The discounts applied to the line item.
type LineItemDiscount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// The taxes applied to the line item.
type LineItemTax struct {
	Amount int64    `json:"amount"`
	Rate   *TaxRate `json:"rate"`

	// This property never worked. Use Rate instead.
	// TODO: Remove in the next major version
	TaxRate *TaxRate `json:"tax_rate"`
}

// A line item.
type LineItem struct {
	APIResource
	AmountSubtotal int64               `json:"amount_subtotal"`
	AmountTotal    int64               `json:"amount_total"`
	Currency       Currency            `json:"currency"`
	Deleted        bool                `json:"deleted"`
	Description    string              `json:"description"`
	Discounts      []*LineItemDiscount `json:"discounts"`
	ID             string              `json:"id"`
	Object         string              `json:"object"`
	Price          *Price              `json:"price"`
	Quantity       int64               `json:"quantity"`
	Taxes          []*LineItemTax      `json:"taxes"`
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
