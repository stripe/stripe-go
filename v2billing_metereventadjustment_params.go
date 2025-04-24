//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Specifies which event to cancel.
type V2BillingMeterEventAdjustmentCancelParams struct {
	// Unique identifier for the event. You can only cancel events within 24 hours of Stripe receiving them.
	Identifier *string `form:"identifier" json:"identifier"`
}

// Creates a meter event adjustment to cancel a previously sent meter event.
type V2BillingMeterEventAdjustmentParams struct {
	Params `form:"*"`
	// Specifies which event to cancel.
	Cancel *V2BillingMeterEventAdjustmentCancelParams `form:"cancel" json:"cancel"`
	// The name of the meter event. Corresponds with the `event_name` field on a meter.
	EventName *string `form:"event_name" json:"event_name"`
	// Specifies whether to cancel a single event or a range of events for a time period. Time period cancellation is not supported yet.
	Type *string `form:"type" json:"type"`
}

// Specifies which event to cancel.
type V2BillingMeterEventAdjustmentCreateCancelParams struct {
	// Unique identifier for the event. You can only cancel events within 24 hours of Stripe receiving them.
	Identifier *string `form:"identifier" json:"identifier"`
}

// Creates a meter event adjustment to cancel a previously sent meter event.
type V2BillingMeterEventAdjustmentCreateParams struct {
	Params `form:"*"`
	// Specifies which event to cancel.
	Cancel *V2BillingMeterEventAdjustmentCreateCancelParams `form:"cancel" json:"cancel"`
	// The name of the meter event. Corresponds with the `event_name` field on a meter.
	EventName *string `form:"event_name" json:"event_name"`
	// Specifies whether to cancel a single event or a range of events for a time period. Time period cancellation is not supported yet.
	Type *string `form:"type" json:"type"`
}
