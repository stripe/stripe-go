//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The status of the request, signifying whether simulated EarnedCredit creation was initiated.
type V2MoneyManagementEarnedCreditSimulationStatus string

// List of values that V2MoneyManagementEarnedCreditSimulationStatus can take
const (
	V2MoneyManagementEarnedCreditSimulationStatusAccepted V2MoneyManagementEarnedCreditSimulationStatus = "accepted"
)

// EarnedCredit Simulations represent simulated EarnedCredit creation requests for testing purposes.
type V2MoneyManagementEarnedCreditSimulation struct {
	APIResource
	// Has the value true if the object exists in live mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The status of the request, signifying whether simulated EarnedCredit creation was initiated.
	Status V2MoneyManagementEarnedCreditSimulationStatus `json:"status"`
}
