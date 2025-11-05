//
//
// File generated from our OpenAPI spec
//
//

package stripe

type TransitBalanceBalance struct {
	Available int64 `json:"available"`
	Pending   int64 `json:"pending"`
}

// Funds that are in transit and destined for another balance or another connected account.
type TransitBalance struct {
	Balance *TransitBalanceBalance `json:"balance"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
