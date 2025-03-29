//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Closed Enum. The status of the request.
type V2FinancialAddressGeneratedMicrodepositsStatus string

// List of values that V2FinancialAddressGeneratedMicrodepositsStatus can take
const (
	V2FinancialAddressGeneratedMicrodepositsStatusAccepted V2FinancialAddressGeneratedMicrodepositsStatus = "accepted"
)

type V2FinancialAddressGeneratedMicrodeposits struct {
	APIResource
	// The amounts of the microdeposits that were generated.
	Amounts []Amount `json:"amounts"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Closed Enum. The status of the request.
	Status V2FinancialAddressGeneratedMicrodepositsStatus `json:"status"`
}
