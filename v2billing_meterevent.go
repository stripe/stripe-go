//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// A Meter Event is a usage record that captures billable activity for usage-based billing. Meter Events contain an event name, timestamp, and payload with customer mapping and usage value, enabling accurate usage tracking and billing.
type V2BillingMeterEvent struct {
	APIResource
	// The creation time of this meter event.
	Created time.Time `json:"created"`
	// The name of the meter event. Corresponds with the `event_name` field on a meter.
	EventName string `json:"event_name"`
	// A unique identifier for the event. If not provided, one will be generated. We recommend using a globally unique identifier for this. We'll enforce uniqueness within a rolling 24 hour period.
	Identifier string `json:"identifier"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The payload of the event. This must contain the fields corresponding to a meter's
	// `customer_mapping.event_payload_key` (default is `stripe_customer_id`) and
	// `value_settings.event_payload_key` (default is `value`). Read more about
	// the [payload](https://docs.stripe.com/billing/subscriptions/usage-based/recording-usage#payload-key-overrides)..
	Payload map[string]string `json:"payload"`
	// The time of the event. Must be within the past 35 calendar days or up to
	// 5 minutes in the future. Defaults to current timestamp if not specified.
	Timestamp time.Time `json:"timestamp"`
}
