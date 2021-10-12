package stripe

import "encoding/json"

// OrderItemType represents the type of order item
type OrderItemType string

// List of values that OrderItemType can take.
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

// OrderItem is the resource representing an order item.
type OrderItem struct {
	Amount      int64            `json:"amount"`
	Currency    Currency         `json:"currency"`
	Description string           `json:"description"`
	Parent      *OrderItemParent `json:"parent"`
	Quantity    int64            `json:"quantity"`
	Type        OrderItemType    `json:"type"`
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
