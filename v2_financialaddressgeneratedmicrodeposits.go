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

// The amounts of the microdeposits that were generated.
type V2FinancialAddressGeneratedMicrodepositsAmount struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}
type V2FinancialAddressGeneratedMicrodeposits struct {
	APIResource
	// The amounts of the microdeposits that were generated.
	Amounts []*V2FinancialAddressGeneratedMicrodepositsAmount `json:"amounts"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Closed Enum. The status of the request.
	Status V2FinancialAddressGeneratedMicrodepositsStatus `json:"status"`
}
