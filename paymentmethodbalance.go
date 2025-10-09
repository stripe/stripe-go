//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The hashes of balances and amounts for available balances.
type PaymentMethodBalanceBalanceFRMealVoucherAvailable struct {
	// The amount of the balance.
	Amount int64 `json:"amount"`
	// The currency of the balance.
	Currency Currency `json:"currency"`
}

// The available FR Meal Voucher balances.
type PaymentMethodBalanceBalanceFRMealVoucher struct {
	// The hashes of balances and amounts for available balances.
	Available []*PaymentMethodBalanceBalanceFRMealVoucherAvailable `json:"available"`
}

// BalanceEntry contain information about every individual balance type of a card.
type PaymentMethodBalanceBalance struct {
	// The available FR Meal Voucher balances.
	FRMealVoucher *PaymentMethodBalanceBalanceFRMealVoucher `json:"fr_meal_voucher"`
}

// PaymentMethodBalance objects represent balances available on a payment method.
// You can use v1/payment_methods/:id/check_balance to check the balance of a payment method.
type PaymentMethodBalance struct {
	APIResource
	// The time at which the balance was calculated. Measured in seconds since the Unix epoch.
	AsOf int64 `json:"as_of"`
	// BalanceEntry contain information about every individual balance type of a card.
	Balance *PaymentMethodBalanceBalance `json:"balance"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
