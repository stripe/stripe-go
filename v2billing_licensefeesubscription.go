//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// A License Fee Subscription represents a customer's subscription to a License Fee at a specified quantity. It tracks
// the number of units (such as seats or licenses) the customer has subscribed to and bills them according to the service
// interval defined in the License Fee and the Billing Cadence.
type V2BillingLicenseFeeSubscription struct {
	APIResource
	// The ID of the Billing Cadence.
	BillingCadence string `json:"billing_cadence"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The ID of the License Fee.
	LicenseFee string `json:"license_fee"`
	// The ID of the License Fee Version.
	LicenseFeeVersion string `json:"license_fee_version"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Quantity of the License Fee subscribed to.
	Quantity int64 `json:"quantity"`
	// The ID of the Test Clock, if any.
	TestClock string `json:"test_clock,omitempty"`
}
