package stripe

import (
	"encoding/json"
)

// LineItemDiscount represent the details of one discount applied to a line item.
type LineItemDiscount struct {
	Amount   int64     `json:"amount"`
	Discount *Discount `json:"discount"`
}

// LineItemTax represent the details of one tax rate applied to a line item.
type LineItemTax struct {
	Amount  int64    `json:"amount"`
	TaxRate *TaxRate `json:"tax_rate"`
}

// LineItem is the resource representing a line item.
type LineItem struct {
	APIResource
	AmountSubtotal int64               `json:"amount_subtotal"`
	AmountTotal    int64               `json:"amount_total"`
	Currency       Currency            `json:"currency"`
	Description    string              `json:"description"`
	Discounts      []*LineItemDiscount `json:"discounts"`
	Deleted        bool                `json:"deleted"`
	ID             string              `json:"id"`
	Object         string              `json:"object"`
	Price          *Price              `json:"price"`
	Quantity       int64               `json:"quantity"`
	Taxes          []*LineItemTax      `json:"taxes"`
}

// LineItemList is a list of prices as returned from a list endpoint.
type LineItemList struct {
	APIResource
	ListMeta
	Data []*LineItem `json:"data"`
}

// UnmarshalJSON handles deserialization of a LineItem.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *LineItem) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type price LineItem
	var v price
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = LineItem(v)
	return nil
}
