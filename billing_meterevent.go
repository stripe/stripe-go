//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a billing meter event
type BillingMeterEventParams struct {
	Params `form:"*"`
	// The name of the meter event. Corresponds with the `event_name` field on a meter.
	EventName *string `form:"event_name"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A unique identifier for the event. If not provided, one will be generated. We recommend using a globally unique identifier for this. We'll enforce uniqueness within a rolling 24 hour period.
	Identifier *string `form:"identifier"`
	// The payload of the event. This must contain a field with the event's numerical value and a field to map the event to a customer.
	Payload map[string]string `form:"payload"`
	// The time of the event. Measured in seconds since the Unix epoch.
	Timestamp *int64 `form:"timestamp"`
}

// AddExpand appends a new field to expand.
func (p *BillingMeterEventParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A billing meter event represents a customer's usage of a product. Meter events are used to bill a customer based on their usage.
// Meter events are associated with billing meters, which define the shape of the event's payload and how those events are aggregated for billing.
type BillingMeterEvent struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The name of the meter event. Corresponds with the `event_name` field on a meter.
	EventName string `json:"event_name"`
	// A unique identifier for the event.
	Identifier string `json:"identifier"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The payload of the event.
	Payload map[string]string `json:"payload"`
	// The timestamp passed in when creating the event. Measured in seconds since the Unix epoch.
	Timestamp int64 `json:"timestamp"`
}
