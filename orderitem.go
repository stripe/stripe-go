//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The type of line item. One of `sku`, `tax`, `shipping`, or `discount`.
type OrderItemType string

// List of values that OrderItemType can take
const (
	OrderItemTypeCoupon   OrderItemType = "coupon"
	OrderItemTypeDiscount OrderItemType = "discount"
	OrderItemTypeSKU      OrderItemType = "sku"
	OrderItemTypeShipping OrderItemType = "shipping"
	OrderItemTypeTax      OrderItemType = "tax"
)

// OrderItemParentType represents the type of order item parent
type OrderItemParentType string

// List of values that OrderItemParentType can take.
const (
	OrderItemParentTypeCoupon   OrderItemParentType = "coupon"
	OrderItemParentTypeDiscount OrderItemParentType = "discount"
	OrderItemParentTypeSKU      OrderItemParentType = "sku"
	OrderItemParentTypeShipping OrderItemParentType = "shipping"
	OrderItemParentTypeTax      OrderItemParentType = "tax"
)

// OrderItemParent describes the parent of an order item.
type OrderItemParent struct {
	ID   string              `json:"id"`
	SKU  *SKU                `json:"-"`
	Type OrderItemParentType `json:"object"`
}

// A representation of the constituent items of any given order. Can be used to
// represent [SKUs](https://stripe.com/docs/api#skus), shipping costs, or taxes owed on the order.
//
// Related guide: [Orders](https://stripe.com/docs/orders/guide).
type OrderItem struct {
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount for the line item.
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Description of the line item, meant to be displayable to the user (e.g., `"Express shipping"`).
	Description string `json:"description"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The ID of the associated object for this line item. Expandable if not null (e.g., expandable to a SKU).
	Parent *OrderItemParent `json:"parent"`
	// A positive integer representing the number of instances of `parent` that are included in this order item. Applicable/present only if `type` is `sku`.
	Quantity int64 `json:"quantity"`
	// The type of line item. One of `sku`, `tax`, `shipping`, or `discount`.
	Type OrderItemType `json:"type"`
}

// UnmarshalJSON handles deserialization of an OrderItemParent.
// This custom unmarshaling is needed because the resulting
// property may be an id or a full SKU struct if it was expanded.
func (p *OrderItemParent) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type orderItemParent OrderItemParent
	var v orderItemParent
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var err error
	*p = OrderItemParent(v)

	switch p.Type {
	case OrderItemParentTypeSKU:
		// Currently only SKUs `parent` is returned as an object when expanded.
		// For other items only IDs are returned.
		if err = json.Unmarshal(data, &p.SKU); err != nil {
			return err
		}
		p.ID = p.SKU.ID
	}

	return nil
}
