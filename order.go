//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// OrderStatus represents the statuses of an order object.
type OrderStatus string

// List of values that OrderStatus can take.
const (
	OrderStatusCanceled  OrderStatus = "canceled"
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusFulfilled OrderStatus = "fulfilled"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusReturned  OrderStatus = "returned"
)

// The type of estimate. Must be either `"range"` or `"exact"`.
type OrderDeliveryEstimateType string

// List of values that OrderDeliveryEstimateType can take
const (
	OrderDeliveryEstimateTypeExact OrderDeliveryEstimateType = "exact"
	OrderDeliveryEstimateTypeRange OrderDeliveryEstimateType = "range"
)

// List of items constituting the order. An order can have up to 25 items.
type OrderItemParams struct {
	Amount      *int64  `form:"amount"`
	Currency    *string `form:"currency"`
	Description *string `form:"description"`
	Parent      *string `form:"parent"`
	Quantity    *int64  `form:"quantity"`
	Type        *string `form:"type"`
}

// Shipping address for the order. Required if any of the SKUs are for products that have `shippable` set to true.
type ShippingParams struct {
	Address        *AddressParams `form:"address"`
	Carrier        *string        `form:"carrier"`
	Name           *string        `form:"name"`
	Phone          *string        `form:"phone"`
	TrackingNumber *string        `form:"tracking_number"`
}

// Creates a new order object.
type OrderParams struct {
	Params                 `form:"*"`
	Coupon                 *string            `form:"coupon"`
	Currency               *string            `form:"currency"`
	Customer               *string            `form:"customer"`
	Email                  *string            `form:"email"`
	Items                  []*OrderItemParams `form:"items"`
	SelectedShippingMethod *string            `form:"selected_shipping_method"`
	Shipping               *ShippingParams    `form:"shipping"`
	Status                 *string            `form:"status"`
}

// Filter orders based on when they were paid, fulfilled, canceled, or returned.
type StatusTransitionsFilterParams struct {
	Canceled       *int64            `form:"canceled"`
	CanceledRange  *RangeQueryParams `form:"canceled"`
	Fulfilled      *int64            `form:"fulfilled"`
	FulfilledRange *RangeQueryParams `form:"fulfilled"`
	Paid           *int64            `form:"paid"`
	PaidRange      *RangeQueryParams `form:"paid"`
	Returned       *int64            `form:"returned"`
	ReturnedRange  *RangeQueryParams `form:"returned"`
}

// Returns a list of your orders. The orders are returned sorted by creation date, with the most recently created orders appearing first.
type OrderListParams struct {
	ListParams        `form:"*"`
	Created           *int64                         `form:"created"`
	CreatedRange      *RangeQueryParams              `form:"created"`
	Customer          *string                        `form:"customer"`
	IDs               []*string                      `form:"ids"`
	Status            *string                        `form:"status"`
	StatusTransitions *StatusTransitionsFilterParams `form:"status_transitions"`
	UpstreamIDs       []*string                      `form:"upstream_ids"`
}

// Pay an order by providing a source to create a payment.
type OrderPayParams struct {
	Params         `form:"*"`
	ApplicationFee *int64        `form:"application_fee"`
	Customer       *string       `form:"customer"`
	Email          *string       `form:"email"`
	Source         *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
}

// SetSource adds valid sources to a OrderPayParams object,
// returning an error for unsupported sources.
func (p *OrderPayParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// The estimated delivery date for the given shipping method. Can be either a specific date or a range.
type DeliveryEstimate struct {
	// If Type == Exact
	Date string `json:"date"`
	// If Type == Range
	Earliest string                    `json:"earliest"`
	Latest   string                    `json:"latest"`
	Type     OrderDeliveryEstimateType `json:"type"`
}

// A list of supported shipping methods for this order. The desired shipping method can be specified either by updating the order, or when paying it.
type ShippingMethod struct {
	Amount           int64             `json:"amount"`
	Currency         Currency          `json:"currency"`
	DeliveryEstimate *DeliveryEstimate `json:"delivery_estimate"`
	Description      string            `json:"description"`
	ID               string            `json:"id"`
}

// The timestamps at which the order status was updated.
type StatusTransitions struct {
	Canceled  int64 `json:"canceled"`
	Fulfilled int64 `json:"fulfilled"`
	Paid      int64 `json:"paid"`
	Returned  int64 `json:"returned"`
}

// OrderUpdateParams is the set of parameters that can be used when updating an order.
type OrderUpdateParams struct {
	Params                 `form:"*"`
	Coupon                 *string                    `form:"coupon"`
	SelectedShippingMethod *string                    `form:"selected_shipping_method"`
	Shipping               *OrderUpdateShippingParams `form:"shipping"`
	Status                 *string                    `form:"status"`
}

// OrderUpdateShippingParams is the set of parameters that can be used for the shipping
// hash on order update.
type OrderUpdateShippingParams struct {
	Carrier        *string `form:"carrier"`
	TrackingNumber *string `form:"tracking_number"`
}

// Shipping describes the shipping hash on an order.
type Shipping struct {
	Address        *Address `json:"address"`
	Carrier        string   `json:"carrier"`
	Name           string   `json:"name"`
	Phone          string   `json:"phone"`
	TrackingNumber string   `json:"tracking_number"`
}

// Order objects are created to handle end customers' purchases of previously
// defined [products](https://stripe.com/docs/api#products). You can create, retrieve, and pay individual orders, as well
// as list all orders. Orders are identified by a unique, random ID.
//
// Related guide: [Tax, Shipping, and Inventory](https://stripe.com/docs/orders).
type Order struct {
	APIResource
	Amount                 int64             `json:"amount"`
	AmountReturned         int64             `json:"amount_returned"`
	Application            string            `json:"application"`
	ApplicationFee         int64             `json:"application_fee"`
	Charge                 *Charge           `json:"charge"`
	Created                int64             `json:"created"`
	Currency               Currency          `json:"currency"`
	Customer               Customer          `json:"customer"`
	Email                  string            `json:"email"`
	ExternalCouponCode     string            `json:"external_coupon_code"`
	ID                     string            `json:"id"`
	Items                  []*OrderItem      `json:"items"`
	Livemode               bool              `json:"livemode"`
	Metadata               map[string]string `json:"metadata"`
	Object                 string            `json:"object"`
	Returns                *OrderReturnList  `json:"returns"`
	SelectedShippingMethod *string           `json:"selected_shipping_method"`
	Shipping               *Shipping         `json:"shipping"`
	ShippingMethods        []*ShippingMethod `json:"shipping_methods"`
	Status                 string            `json:"status"`
	StatusTransitions      StatusTransitions `json:"status_transitions"`
	Updated                int64             `json:"updated"`
	UpstreamID             string            `json:"upstream_id"`
}

// OrderList is a list of Orders as retrieved from a list endpoint.
type OrderList struct {
	APIResource
	ListMeta
	Data []*Order `json:"data"`
}

// UnmarshalJSON handles deserialization of an Order.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (o *Order) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		o.ID = id
		return nil
	}

	type order Order
	var v order
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*o = Order(v)
	return nil
}
