//
//
// File generated from our OpenAPI spec
//
//

package stripe

type DelegatedCheckoutOrderLineItemProductDetails struct {
	// The item description.
	Description string `json:"description"`
	// The item images.
	Images []string `json:"images"`
	// The item title.
	Title string `json:"title"`
}
type DelegatedCheckoutOrderLineItemQuantity struct {
	// The current quantity.
	Current int64 `json:"current"`
	// The ordered quantity.
	Ordered int64 `json:"ordered"`
	// The shipped quantity.
	Shipped int64 `json:"shipped"`
}

// The totals for this line item.
type DelegatedCheckoutOrderLineItemTotals struct {
	// The base amount for the line item.
	BaseAmount int64 `json:"base_amount"`
	// The discount amount for the line item.
	Discount int64 `json:"discount"`
	// The subtotal amount for the line item.
	Subtotal int64 `json:"subtotal"`
	// The tax amount for the line item.
	Tax int64 `json:"tax"`
	// The total amount for the line item.
	Total int64 `json:"total"`
}

// The line items in this order.
type DelegatedCheckoutOrderLineItem struct {
	// The order line item key.
	Key            string                                        `json:"key"`
	ProductDetails *DelegatedCheckoutOrderLineItemProductDetails `json:"product_details"`
	Quantity       *DelegatedCheckoutOrderLineItemQuantity       `json:"quantity"`
	// The SKU ID of the line item.
	SKUID string `json:"sku_id"`
	// The totals for this line item.
	Totals *DelegatedCheckoutOrderLineItemTotals `json:"totals"`
	// The line item unit amount.
	UnitAmount int64 `json:"unit_amount"`
}

// The totals for this order.
type DelegatedCheckoutOrderTotals struct {
	// The discount amount for the order.
	Discount int64 `json:"discount"`
	// The fulfillment amount for the order.
	Fulfillment int64 `json:"fulfillment"`
	// The subtotal amount for the order.
	Subtotal int64 `json:"subtotal"`
	// The tax amount for the order.
	Tax int64 `json:"tax"`
	// The total amount for the order.
	Total int64 `json:"total"`
}

// An order represents the post-checkout lifecycle of a delegated checkout purchase.
type DelegatedCheckoutOrder struct {
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The latest order event for this order.
	LatestOrderEvent *DelegatedCheckoutOrderEvent `json:"latest_order_event"`
	// The line items in this order.
	LineItems []*DelegatedCheckoutOrderLineItem `json:"line_items"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The permalink URL for this order.
	PermalinkURL string `json:"permalink_url"`
	// The requested session associated with this order.
	RequestedSession string `json:"requested_session"`
	// The seller reference for this order.
	SellerReference string `json:"seller_reference"`
	// The totals for this order.
	Totals *DelegatedCheckoutOrderTotals `json:"totals"`
}

// DelegatedCheckoutOrderList is a list of Orders as retrieved from a list endpoint.
type DelegatedCheckoutOrderList struct {
	APIResource
	ListMeta
	Data []*DelegatedCheckoutOrder `json:"data"`
}
