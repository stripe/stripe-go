package stripe

// BalanceSourceType is the list of allowed values for the balance amount's source_type field keys.
type BalanceSourceType string

// List of values that BalanceSourceType can take.
const (
	BalanceSourceTypeAlipayAccount   BalanceSourceType = "alipay_account"
	BalanceSourceTypeBankAccount     BalanceSourceType = "bank_account"
	BalanceSourceTypeBitcoinReceiver BalanceSourceType = "bitcoin_receiver"
	BalanceSourceTypeCard            BalanceSourceType = "card"
	BalanceSourceTypeFPX             BalanceSourceType = "fpx"
)

// BalanceTransactionStatus is the list of allowed values for the balance transaction's status.
type BalanceTransactionStatus string

// BalanceParams is the set of parameters that can be used when retrieving a balance.
// For more details see https://stripe.com/docs/api#balance.
type BalanceParams struct {
	Params `form:"*"`
}

// Balance is the resource representing your Stripe balance.
// For more details see https://stripe.com/docs/api/#balance.
type Balance struct {
	Available       []*Amount `json:"available"`
	ConnectReserved []*Amount `json:"connect_reserved"`
	Livemode        bool      `json:"livemode"`
	Pending         []*Amount `json:"pending"`
}

// Amount is a structure wrapping an amount value and its currency.
type Amount struct {
	Currency    Currency                    `json:"currency"`
	SourceTypes map[BalanceSourceType]int64 `json:"source_types"`
	Value       int64                       `json:"amount"`
}
