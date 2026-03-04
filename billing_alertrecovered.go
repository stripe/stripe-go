//
//
// File generated from our OpenAPI spec
//
//

package stripe

type BillingAlertRecovered struct {
	// A billing alert is a resource that notifies you when a certain usage threshold on a meter is crossed. For example, you might create a billing alert to notify you when a certain user made 100 API requests.
	Alert *BillingAlert `json:"alert"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Currency for the threshold value
	Currency Currency `json:"currency"`
	// ID of customer for which the alert recovered
	Customer string `json:"customer"`
	// Custom pricing unit for the threshold value
	CustomPricingUnit string `json:"custom_pricing_unit"`
	// External customer ID for the customer for which the alert recovered
	ExternalCustomerID string `json:"external_customer_id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The value at which the alert recovered
	Value float64 `json:"value,string"`
}
