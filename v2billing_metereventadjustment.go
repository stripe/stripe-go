//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. The meter event adjustment's status.
type V2BillingMeterEventAdjustmentStatus string

// List of values that V2BillingMeterEventAdjustmentStatus can take
const (
	V2BillingMeterEventAdjustmentStatusComplete V2BillingMeterEventAdjustmentStatus = "complete"
	V2BillingMeterEventAdjustmentStatusPending  V2BillingMeterEventAdjustmentStatus = "pending"
)

// Open Enum. Specifies whether to cancel a single event or a range of events for a time period. Time period cancellation is not supported yet.
type V2BillingMeterEventAdjustmentType string

// List of values that V2BillingMeterEventAdjustmentType can take
const (
	V2BillingMeterEventAdjustmentTypeCancel V2BillingMeterEventAdjustmentType = "cancel"
)

// Specifies which event to cancel.
type V2BillingMeterEventAdjustmentCancel struct {
	// Unique identifier for the event. You can only cancel events within 24 hours of Stripe receiving them.
	Identifier string `json:"identifier"`
}

// A Meter Event Adjustment is used to cancel or modify previously recorded meter events. Meter Event Adjustments allow you to correct billing data by canceling individual events or event ranges, with tracking of adjustment status and creation time.
type V2BillingMeterEventAdjustment struct {
	APIResource
	// Specifies which event to cancel.
	Cancel *V2BillingMeterEventAdjustmentCancel `json:"cancel"`
	// The time the adjustment was created.
	Created time.Time `json:"created"`
	// The name of the meter event. Corresponds with the `event_name` field on a meter.
	EventName string `json:"event_name"`
	// The unique id of this meter event adjustment.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Open Enum. The meter event adjustment's status.
	Status V2BillingMeterEventAdjustmentStatus `json:"status"`
	// Open Enum. Specifies whether to cancel a single event or a range of events for a time period. Time period cancellation is not supported yet.
	Type V2BillingMeterEventAdjustmentType `json:"type"`
}
