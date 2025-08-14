//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

type V2BillingRateCardVersion struct {
	APIResource
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// The ID of the RateCardVersion.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the RateCard that this version belongs to.
	RateCardID string `json:"rate_card_id"`
}
