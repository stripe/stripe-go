//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// A Rate Card Version represents a specific configuration of a Rate Card at a point in time. Versions are created automatically
// when you add or modify rates on a Rate Card, allowing you to track changes and manage which version is active for new
// subscriptions. Each version maintains a record of when it was created.
type V2BillingRateCardVersion struct {
	APIResource
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the Rate Card that this version belongs to.
	RateCardID string `json:"rate_card_id"`
}
