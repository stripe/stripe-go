package stripe

import (
	"encoding/json"
)

// OrderStatus represents the statuses of an order object.
type OrderStatus string

const (
	StatusCanceled  OrderStatus = "canceled"
	StatusCreated   OrderStatus = "created"
	StatusFulfilled OrderStatus = "fulfilled"
	StatusPaid      OrderStatus = "paid"
	StatusReturned  OrderStatus = "returned"
)

type OrderParams struct {
	Params   `form:"*"`
	Coupon   string             `form:"coupon"`
	Currency Currency           `form:"currency"`
	Customer string             `form:"customer"`
	Email    string             `form:"email"`
	Items    []*OrderItemParams `form:"items,indexed"`
	Shipping *ShippingParams    `form:"shipping"`
}

type ShippingParams struct {
	Address *AddressParams `form:"address"`
	Name    string         `form:"name"`
	Phone   string         `form:"phone"`
}

type OrderUpdateParams struct {
	Params                 `form:"*"`
	Coupon                 string      `form:"coupon"`
	SelectedShippingMethod string      `form:"selected_shipping_method"`
	Status                 OrderStatus `form:"status"`
}

// OrderReturnParams is the set of parameters that can be used when returning
// orders. For more details, see: https://stripe.com/docs/api#return_order.
type OrderReturnParams struct {
	Params `form:"*"`
	Items  []*OrderItemParams `form:"items,indexed"`
}

type Shipping struct {
	Address Address `json:"address"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
}

type ShippingMethod struct {
	Amount           int64             `json:"amount"`
	ID               string            `json:"id"`
	Currency         Currency          `json:"currency"`
	DeliveryEstimate *DeliveryEstimate `json:"delivery_estimate"`
	Description      string            `json:"description"`
}

type EstimateType string

const (
	Exact EstimateType = "exact"
	Range EstimateType = "range"
)

type DeliveryEstimate struct {
	Type EstimateType `json:"type"`
	// If Type == Range
	Earliest string `json:"earliest"`
	Latest   string `json:"latest"`
	// If Type == Exact
	Date string `json:"date"`
}

type Order struct {
	Amount                 int64             `json:"amount"`
	AmountReturned         int64             `json:"amount_returned"`
	Application            string            `json:"application"`
	ApplicationFee         int64             `json:"application_fee"`
	Charge                 Charge            `json:"charge"`
	Created                int64             `json:"created"`
	Currency               Currency          `json:"currency"`
	Customer               Customer          `json:"customer"`
	Email                  string            `json:"email"`
	ID                     string            `json:"id"`
	Items                  []OrderItem       `json:"items"`
	Live                   bool              `json:"livemode"`
	Meta                   map[string]string `json:"metadata"`
	Returns                *OrderReturnList  `json:"returns"`
	SelectedShippingMethod *string           `json:"selected_shipping_method"`
	Shipping               Shipping          `json:"shipping"`
	ShippingMethods        []ShippingMethod  `json:"shipping_methods"`
	Status                 OrderStatus       `json:"status"`
	StatusTransitions      StatusTransitions `json:"status_transitions"`
	Updated                int64             `json:"updated"`
}

// OrderList is a list of orders as retrieved from a list endpoint.
type OrderList struct {
	ListMeta
	Values []*Order `json:"data"`
}

// OrderListParams is the set of parameters that can be used when
// listing orders. For more details, see:
// https://stripe.com/docs/api#list_orders.
type OrderListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	IDs          []string          `form:"ids"`
	Status       OrderStatus       `form:"status"`
}

// StatusTransitions are the timestamps at which the order status was updated
// https://stripe.com/docs/api#order_object
type StatusTransitions struct {
	Canceled  int64 `json:"canceled"`
	Fulfilled int64 `json:"fulfiled"`
	Paid      int64 `json:"paid"`
	Returned  int64 `json:"returned"`
}

// OrderPayParams is the set of parameters that can be used when
// paying orders. For more details, see:
// https://stripe.com/docs/api#pay_order.
type OrderPayParams struct {
	Params         `form:"*"`
	ApplicationFee int64         `form:"application_fee"`
	Customer       string        `form:"customer"`
	Email          string        `form:"email"`
	Source         *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
}

// SetSource adds valid sources to a OrderParams object,
// returning an error for unsupported sources.
func (op *OrderPayParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	op.Source = source
	return err
}

// UnmarshalJSON handles deserialization of an Order.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (o *Order) UnmarshalJSON(data []byte) error {
	type order Order
	var oo order
	err := json.Unmarshal(data, &oo)
	if err == nil {
		*o = Order(oo)
		{
		}
	} else {
		// the id is surrounded by "\" characters, so strip them
		o.ID = string(data[1 : len(data)-1])
	}

	return nil
}
