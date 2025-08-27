//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The Custom Pricing Unit object.
type V2BillingCustomPricingUnit struct {
	APIResource
	// Whether the custom pricing unit is active.
	Active bool `json:"active"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Description that customers will see in the invoice line item.
	// Maximum length of 10 characters.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// An internal key you can use to search for a particular Custom Pricing Unit.
	// Maximum length of 200 characters.
	LookupKey string `json:"lookup_key"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
}
