//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The current status of the billing profile.
type V2BillingProfileStatus string

// List of values that V2BillingProfileStatus can take
const (
	V2BillingProfileStatusActive   V2BillingProfileStatus = "active"
	V2BillingProfileStatusInactive V2BillingProfileStatus = "inactive"
)

type V2BillingProfile struct {
	APIResource
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// The ID of the customer object.
	Customer string `json:"customer"`
	// The ID of the payment method object.
	DefaultPaymentMethod string `json:"default_payment_method"`
	// A customer-facing name for the billing profile.
	// Maximum length of 250 characters.
	DisplayName string `json:"display_name"`
	// The ID of the billing profile object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// An internal key you can use to search for a particular billing profile.
	// Maximum length of 200 characters.
	LookupKey string `json:"lookup_key"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The current status of the billing profile.
	Status V2BillingProfileStatus `json:"status"`
}
