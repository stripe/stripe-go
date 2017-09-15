package stripe

import (
	// This is pretty silly. Nowhere else do we import a subpackage from the
	// top-level namespace like this. We should try to deprecate this at some
	// point.
	"github.com/stripe/stripe-go/orderitem"
)

type OrderItemParams struct {
	Amount      int64              `form:"amount"`
	Currency    Currency           `form:"currency"`
	Description string             `form:"description"`
	Parent      string             `form:"parent"`
	Quantity    *int64             `form:"quantity"`
	Type        orderitem.ItemType `form:"type"`
}

type OrderItem struct {
	Amount      int64              `json:"amount"`
	Currency    Currency           `json:"currency"`
	Description string             `json:"description"`
	Parent      string             `json:"parent"`
	Quantity    int64              `json:"quantity"`
	Type        orderitem.ItemType `json:"type"`
}
