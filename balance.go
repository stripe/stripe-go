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
	Value       int64                       `json:"amount"`
	Currency    Currency                    `json:"currency"`
	SourceTypes map[BalanceSourceType]int64 `json:"source_types"`
}
type BalanceDetails struct {
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
	Available        []*Amount       `json:"available"`
	ConnectReserved  []*Amount       `json:"connect_reserved"`
	InstantAvailable []*Amount       `json:"instant_available"`
	Issuing          *BalanceDetails `json:"issuing"`
	Livemode         bool            `json:"livemode"`
	Object           string          `json:"object"`
	Pending          []*Amount       `json:"pending"`
}
