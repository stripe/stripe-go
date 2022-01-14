//
//
// File generated from our OpenAPI spec
//
//

package stripe

// BalanceSourceType is the list of allowed values for the balance amount's source_type field keys.
type BalanceSourceType string

// List of values that BalanceSourceType can take.
const (
	BalanceSourceTypeBankAccount BalanceSourceType = "bank_account"
	BalanceSourceTypeCard        BalanceSourceType = "card"
	BalanceSourceTypeFPX         BalanceSourceType = "fpx"
)

// BalanceTransactionStatus is the list of allowed values for the balance transaction's status.
type BalanceTransactionStatus string

// Retrieves the current account balance, based on the authentication that was used to make the request.
//  For a sample request, see [Accounting for negative balances](https://stripe.com/docs/connect/account-balances#accounting-for-negative-balances).
type BalanceParams struct {
	Params `form:"*"`
}

// Funds that are available to be transferred or paid out, whether automatically by Stripe or explicitly via the [Transfers API](https://stripe.com/docs/api#transfers) or [Payouts API](https://stripe.com/docs/api#payouts). The available balance for each currency and payment type can be found in the `source_types` property.
type Amount struct {
	// Balance amount.
	Value int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency    Currency                    `json:"currency"`
	SourceTypes map[BalanceSourceType]int64 `json:"source_types"`
}
type BalanceDetails struct {
	// Funds that are available for use.
	Available []*Amount `json:"available"`
}

// This is an object representing your Stripe balance. You can retrieve it to see
// the balance currently on your Stripe account.
//
// You can also retrieve the balance history, which contains a list of
// [transactions](https://stripe.com/docs/reporting/balance-transaction-types) that contributed to the balance
// (charges, payouts, and so forth).
//
// The available and pending amounts for each currency are broken down further by
// payment source types.
//
// Related guide: [Understanding Connect Account Balances](https://stripe.com/docs/connect/account-balances).
type Balance struct {
	APIResource
	// Funds that are available to be transferred or paid out, whether automatically by Stripe or explicitly via the [Transfers API](https://stripe.com/docs/api#transfers) or [Payouts API](https://stripe.com/docs/api#payouts). The available balance for each currency and payment type can be found in the `source_types` property.
	Available []*Amount `json:"available"`
	// Funds held due to negative balances on connected Custom accounts. The connect reserve balance for each currency and payment type can be found in the `source_types` property.
	ConnectReserved []*Amount `json:"connect_reserved"`
	// Funds that can be paid out using Instant Payouts.
	InstantAvailable []*Amount       `json:"instant_available"`
	Issuing          *BalanceDetails `json:"issuing"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Funds that are not yet available in the balance, due to the 7-day rolling pay cycle. The pending balance for each currency, and for each payment type, can be found in the `source_types` property.
	Pending []*Amount `json:"pending"`
}
