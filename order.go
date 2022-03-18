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
	// The ID of the SKU being ordered.
	Parent *string `form:"parent"`
	// The quantity of this order item. When type is `sku`, this is the number of instances of the SKU to be ordered.
	Quantity *int64  `form:"quantity"`
	Type     *string `form:"type"`
}

// Shipping address for the order. Required if any of the SKUs are for products that have `shippable` set to true.
type ShippingParams struct {
	// Customer shipping address.
	Address *AddressParams `form:"address"`
	// The name of the carrier like `USPS`, `UPS`, or `FedEx`.
	Carrier *string `form:"carrier"`
	// Customer name.
	Name *string `form:"name"`
	// Customer phone (including extension).
	Phone *string `form:"phone"`
	// The tracking number provided by the carrier.
	TrackingNumber *string `form:"tracking_number"`
}

// Creates a new order object.
type OrderParams struct {
	Params `form:"*"`
	// A coupon code that represents a discount to be applied to this order. Must be one-time duration and in same currency as the order. An order can have multiple coupons.
	Coupon *string `form:"coupon"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// The ID of an existing customer to use for this order. If provided, the customer email and shipping address will be used to create the order. Subsequently, the customer will also be charged to pay the order. If `email` or `shipping` are also provided, they will override the values retrieved from the customer object.
	Customer *string `form:"customer"`
	// The email address of the customer placing the order.
	Email *string `form:"email"`
	// List of items constituting the order. An order can have up to 25 items.
	Items []*OrderItemParams `form:"items"`
	// The shipping method to select for fulfilling this order. If specified, must be one of the `id`s of a shipping method in the `shipping_methods` array. If specified, will overwrite the existing selected shipping method, updating `items` as necessary.
	SelectedShippingMethod *string `form:"selected_shipping_method"`
	// Tracking information once the order has been fulfilled.
	Shipping *ShippingParams `form:"shipping"`
	// Current order status. One of `created`, `paid`, `canceled`, `fulfilled`, or `returned`. More detail in the [Orders Guide](https://stripe.com/docs/orders/guide#understanding-order-statuses).
	Status *string `form:"status"`
}

// Filter orders based on when they were paid, fulfilled, canceled, or returned.
type StatusTransitionsFilterParams struct {
	// Date this order was canceled.
	Canceled *int64 `form:"canceled"`
	// Date this order was canceled.
	CanceledRange *RangeQueryParams `form:"canceled"`
	// Date this order was fulfilled.
	Fulfilled *int64 `form:"fulfilled"`
	// Date this order was fulfilled.
	FulfilledRange *RangeQueryParams `form:"fulfilled"`
	// Date this order was paid.
	Paid *int64 `form:"paid"`
	// Date this order was paid.
	PaidRange *RangeQueryParams `form:"paid"`
	// Date this order was returned.
	Returned *int64 `form:"returned"`
	// Date this order was returned.
	ReturnedRange *RangeQueryParams `form:"returned"`
}

// Returns a list of your orders. The orders are returned sorted by creation date, with the most recently created orders appearing first.
type OrderListParams struct {
	ListParams `form:"*"`
	// Date this order was created.
	Created *int64 `form:"created"`
	// Date this order was created.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return orders for the given customer.
	Customer *string `form:"customer"`
	// Only return orders with the given IDs.
	IDs []*string `form:"ids"`
	// Only return orders that have the given status. One of `created`, `paid`, `fulfilled`, or `refunded`.
	Status *string `form:"status"`
	// Filter orders based on when they were paid, fulfilled, canceled, or returned.
	StatusTransitions *StatusTransitionsFilterParams `form:"status_transitions"`
	// Only return orders with the given upstream order IDs.
	UpstreamIDs []*string `form:"upstream_ids"`
}

// Pay an order by providing a source to create a payment.
type OrderPayParams struct {
	Params `form:"*"`
	// A fee in %s that will be applied to the order and transferred to the application owner's Stripe account. The request must be made with an OAuth key or the `Stripe-Account` header in order to take an application fee. For more information, see the application fees [documentation](https://stripe.com/docs/connect/direct-charges#collecting-fees).
	ApplicationFee *int64 `form:"application_fee"`
	// The ID of an existing customer that will be charged for this order. If no customer was attached to the order at creation, either `source` or `customer` is required. Otherwise, the specified customer will be charged instead of the one attached to the order.
	Customer *string `form:"customer"`
	// The email address of the customer placing the order. Required if not previously specified for the order.
	Email  *string       `form:"email"`
	Source *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
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
	// If `type` is `"exact"`, `date` will be the expected delivery date in the format YYYY-MM-DD.
	Date string `json:"date"`
	// If Type == Range
	// If `type` is `"range"`, `earliest` will be be the earliest delivery date in the format YYYY-MM-DD.
	Earliest string `json:"earliest"`
	// If `type` is `"range"`, `latest` will be the latest delivery date in the format YYYY-MM-DD.
	Latest string `json:"latest"`
	// The type of estimate. Must be either `"range"` or `"exact"`.
	Type OrderDeliveryEstimateType `json:"type"`
}

// A list of supported shipping methods for this order. The desired shipping method can be specified either by updating the order, or when paying it.
type ShippingMethod struct {
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount for the line item.
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The estimated delivery date for the given shipping method. Can be either a specific date or a range.
	DeliveryEstimate *DeliveryEstimate `json:"delivery_estimate"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
}

// The timestamps at which the order status was updated.
type StatusTransitions struct {
	// The time that the order was canceled.
	Canceled int64 `json:"canceled"`
	// The time that the order was fulfilled.
	Fulfilled int64 `json:"fulfilled"`
	// The time that the order was paid.
	Paid int64 `json:"paid"`
	// The time that the order was returned.
	Returned int64 `json:"returned"`
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
// Related guide: [Tax, Shipping, and Inventory](https://stripe.com/docs/orders-legacy).
type Order struct {
	APIResource
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount for the order.
	Amount int64 `json:"amount"`
	// The total amount that was returned to the customer.
	AmountReturned int64 `json:"amount_returned"`
	// ID of the Connect Application that created the order.
	Application string `json:"application"`
	// A fee in cents that will be applied to the order and transferred to the application owner's Stripe account. The request must be made with an OAuth key or the Stripe-Account header in order to take an application fee. For more information, see the application fees documentation.
	ApplicationFee int64 `json:"application_fee"`
	// The ID of the payment used to pay for the order. Present if the order status is `paid`, `fulfilled`, or `refunded`.
	Charge *Charge `json:"charge"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The customer used for the order.
	Customer Customer `json:"customer"`
	// The email address of the customer placing the order.
	Email string `json:"email"`
	// External coupon code to load for this order.
	ExternalCouponCode string `json:"external_coupon_code"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// List of items constituting the order. An order can have up to 25 items.
	Items []*OrderItem `json:"items"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// A list of returns that have taken place for this order.
	Returns *OrderReturnList `json:"returns"`
	// The shipping method that is currently selected for this order, if any. If present, it is equal to one of the `id`s of shipping methods in the `shipping_methods` array. At order creation time, if there are multiple shipping methods, Stripe will automatically selected the first method.
	SelectedShippingMethod *string `json:"selected_shipping_method"`
	// The shipping address for the order. Present if the order is for goods to be shipped.
	Shipping *Shipping `json:"shipping"`
	// A list of supported shipping methods for this order. The desired shipping method can be specified either by updating the order, or when paying it.
	ShippingMethods []*ShippingMethod `json:"shipping_methods"`
	// Current order status. One of `created`, `paid`, `canceled`, `fulfilled`, or `returned`. More details in the [Orders Guide](https://stripe.com/docs/orders/guide#understanding-order-statuses).
	Status string `json:"status"`
	// The timestamps at which the order status was updated.
	StatusTransitions StatusTransitions `json:"status_transitions"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	Updated int64 `json:"updated"`
	// The user's order ID if it is different from the Stripe order ID.
	UpstreamID string `json:"upstream_id"`
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
