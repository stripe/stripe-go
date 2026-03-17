//
//
// File generated from our OpenAPI spec
//
//

package stripe

type BillingAlertTriggered struct {
	// A billing alert is a resource that notifies you when a certain usage threshold on a meter is crossed. For example, you might create a billing alert to notify you when a certain user made 100 API requests.
	Alert *BillingAlert `json:"alert"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// ID of customer for which the alert triggered
	Customer string `json:"customer"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The value triggering the alert
	Value float64 `json:"value,string"`
}
