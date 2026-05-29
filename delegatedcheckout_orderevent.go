//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The status of the adjustment.
type DelegatedCheckoutOrderEventAdjustmentStatus string

// List of values that DelegatedCheckoutOrderEventAdjustmentStatus can take
const (
	DelegatedCheckoutOrderEventAdjustmentStatusCompleted DelegatedCheckoutOrderEventAdjustmentStatus = "completed"
	DelegatedCheckoutOrderEventAdjustmentStatusFailed    DelegatedCheckoutOrderEventAdjustmentStatus = "failed"
	DelegatedCheckoutOrderEventAdjustmentStatusPending   DelegatedCheckoutOrderEventAdjustmentStatus = "pending"
)

// The type of adjustment.
type DelegatedCheckoutOrderEventAdjustmentType string

// List of values that DelegatedCheckoutOrderEventAdjustmentType can take
const (
	DelegatedCheckoutOrderEventAdjustmentTypeCancellation          DelegatedCheckoutOrderEventAdjustmentType = "cancellation"
	DelegatedCheckoutOrderEventAdjustmentTypeCredit                DelegatedCheckoutOrderEventAdjustmentType = "credit"
	DelegatedCheckoutOrderEventAdjustmentTypeDispute               DelegatedCheckoutOrderEventAdjustmentType = "dispute"
	DelegatedCheckoutOrderEventAdjustmentTypeOriginalPaymentRefund DelegatedCheckoutOrderEventAdjustmentType = "original_payment_refund"
	DelegatedCheckoutOrderEventAdjustmentTypeReturn                DelegatedCheckoutOrderEventAdjustmentType = "return"
	DelegatedCheckoutOrderEventAdjustmentTypeStoreCreditRefund     DelegatedCheckoutOrderEventAdjustmentType = "store_credit_refund"
)

// The status of the fulfillment.
type DelegatedCheckoutOrderEventFulfillmentStatus string

// List of values that DelegatedCheckoutOrderEventFulfillmentStatus can take
const (
	DelegatedCheckoutOrderEventFulfillmentStatusConfirmed DelegatedCheckoutOrderEventFulfillmentStatus = "confirmed"
	DelegatedCheckoutOrderEventFulfillmentStatusDelivered DelegatedCheckoutOrderEventFulfillmentStatus = "delivered"
	DelegatedCheckoutOrderEventFulfillmentStatusFulfilled DelegatedCheckoutOrderEventFulfillmentStatus = "fulfilled"
	DelegatedCheckoutOrderEventFulfillmentStatusPending   DelegatedCheckoutOrderEventFulfillmentStatus = "pending"
	DelegatedCheckoutOrderEventFulfillmentStatusReturned  DelegatedCheckoutOrderEventFulfillmentStatus = "returned"
	DelegatedCheckoutOrderEventFulfillmentStatusShipped   DelegatedCheckoutOrderEventFulfillmentStatus = "shipped"
)

// The type of order event.
type DelegatedCheckoutOrderEventType string

// List of values that DelegatedCheckoutOrderEventType can take
const (
	DelegatedCheckoutOrderEventTypeAdjustment  DelegatedCheckoutOrderEventType = "adjustment"
	DelegatedCheckoutOrderEventTypeFulfillment DelegatedCheckoutOrderEventType = "fulfillment"
)

// The line items associated with the adjustment.
type DelegatedCheckoutOrderEventAdjustmentLineItem struct {
	// The line item key.
	Key string `json:"key"`
	// The quantity associated with the order event.
	Quantity int64 `json:"quantity"`
}

// The adjustment details for this order event.
type DelegatedCheckoutOrderEventAdjustment struct {
	// The amount associated with the adjustment.
	Amount int64 `json:"amount"`
	// The currency associated with the adjustment amount.
	Currency Currency `json:"currency"`
	// The description of the adjustment.
	Description string `json:"description"`
	// The line items associated with the adjustment.
	LineItems []*DelegatedCheckoutOrderEventAdjustmentLineItem `json:"line_items"`
	// The status of the adjustment.
	Status DelegatedCheckoutOrderEventAdjustmentStatus `json:"status"`
	// The type of adjustment.
	Type DelegatedCheckoutOrderEventAdjustmentType `json:"type"`
}

// The line items associated with the fulfillment.
type DelegatedCheckoutOrderEventFulfillmentLineItem struct {
	// The line item key.
	Key string `json:"key"`
	// The quantity associated with the order event.
	Quantity int64 `json:"quantity"`
}

// The fulfillment details for this order event.
type DelegatedCheckoutOrderEventFulfillment struct {
	// The carrier for the fulfillment.
	Carrier string `json:"carrier"`
	// Time at which the fulfillment was delivered. Measured in seconds since the Unix epoch.
	DeliveredAt int64 `json:"delivered_at"`
	// The line items associated with the fulfillment.
	LineItems []*DelegatedCheckoutOrderEventFulfillmentLineItem `json:"line_items"`
	// Time at which the fulfillment shipped. Measured in seconds since the Unix epoch.
	ShippedAt int64 `json:"shipped_at"`
	// The status of the fulfillment.
	Status DelegatedCheckoutOrderEventFulfillmentStatus `json:"status"`
	// The tracking number for the fulfillment.
	TrackingNumber string `json:"tracking_number"`
	// The tracking URL for the fulfillment.
	TrackingURL string `json:"tracking_url"`
}

// An order event represents a change to a delegated checkout order.
type DelegatedCheckoutOrderEvent struct {
	// The adjustment details for this order event.
	Adjustment *DelegatedCheckoutOrderEventAdjustment `json:"adjustment"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The fulfillment details for this order event.
	Fulfillment *DelegatedCheckoutOrderEventFulfillment `json:"fulfillment"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Time at which this event occurred. Measured in seconds since the Unix epoch.
	OccurredAt int64 `json:"occurred_at"`
	// The delegated checkout order associated with this order event.
	Order string `json:"order"`
	// The requested session associated with this order event.
	RequestedSession string `json:"requested_session"`
	// The type of order event.
	Type DelegatedCheckoutOrderEventType `json:"type"`
}
