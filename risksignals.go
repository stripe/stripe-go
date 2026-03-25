//
//
// File generated from our OpenAPI spec
//
//

package stripe

type RiskSignals struct {
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Represents the status of risk signal session metadata collection. When false, the account has payouts and payments disabled.
	SessionMetadata bool `json:"session_metadata"`
}
