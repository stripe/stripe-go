//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List of meter events to include in the request.
type V2BillingMeterEventStreamEventParams struct {
	// The name of the meter event. Corresponds with the `event_name` field on a meter.
	EventName *string `form:"event_name" json:"event_name"`
	// A unique identifier for the event. If not provided, one will be generated.
	// We recommend using a globally unique identifier for this. We'll enforce
	// uniqueness within a rolling 24 hour period.
	Identifier *string `form:"identifier" json:"identifier,omitempty"`
	// The payload of the event. This must contain the fields corresponding to a meter's
	// `customer_mapping.event_payload_key` (default is `stripe_customer_id`) and
	// `value_settings.event_payload_key` (default is `value`). Read more about
	// the
	// [payload](https://docs.stripe.com/billing/subscriptions/usage-based/recording-usage#payload-key-overrides).
	Payload map[string]string `form:"payload" json:"payload"`
	// The time of the event. Must be within the past 35 calendar days or up to
	// 5 minutes in the future. Defaults to current timestamp if not specified.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
}

// Creates meter events. Events are processed asynchronously, including validation. Requires a meter event session for authentication. Supports up to 10,000 requests per second in livemode. For even higher rate-limits, contact sales.
type V2BillingMeterEventStreamParams struct {
	Params `form:"*"`
	// List of meter events to include in the request.
	Events []*V2BillingMeterEventStreamEventParams `form:"events" json:"events"`
}

// List of meter events to include in the request.
type V2BillingMeterEventStreamCreateEventParams struct {
	// The name of the meter event. Corresponds with the `event_name` field on a meter.
	EventName *string `form:"event_name" json:"event_name"`
	// A unique identifier for the event. If not provided, one will be generated.
	// We recommend using a globally unique identifier for this. We'll enforce
	// uniqueness within a rolling 24 hour period.
	Identifier *string `form:"identifier" json:"identifier,omitempty"`
	// The payload of the event. This must contain the fields corresponding to a meter's
	// `customer_mapping.event_payload_key` (default is `stripe_customer_id`) and
	// `value_settings.event_payload_key` (default is `value`). Read more about
	// the
	// [payload](https://docs.stripe.com/billing/subscriptions/usage-based/recording-usage#payload-key-overrides).
	Payload map[string]string `form:"payload" json:"payload"`
	// The time of the event. Must be within the past 35 calendar days or up to
	// 5 minutes in the future. Defaults to current timestamp if not specified.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
}

// Creates meter events. Events are processed asynchronously, including validation. Requires a meter event session for authentication. Supports up to 10,000 requests per second in livemode. For even higher rate-limits, contact sales.
type V2BillingMeterEventStreamCreateParams struct {
	Params `form:"*"`
	// List of meter events to include in the request.
	Events []*V2BillingMeterEventStreamCreateEventParams `form:"events" json:"events"`
}
