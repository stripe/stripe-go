//
//
// File generated from our OpenAPI spec
//
//

package stripe

type V2FinancialAddressCreditSimulation struct {
	APIResource
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The status of the request, signifying whether a simulated credit was initiated.
	Status string `json:"status"`
}
