//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The configuration for how funds that land in the customer cash balance are reconciled.
type CashBalanceSettingsReconciliationMode string

// List of values that CashBalanceSettingsReconciliationMode can take
const (
	CashBalanceSettingsReconciliationModeAutomatic CashBalanceSettingsReconciliationMode = "automatic"
	CashBalanceSettingsReconciliationModeManual    CashBalanceSettingsReconciliationMode = "manual"
)

// Retrieves a customer's cash balance.
type CashBalanceParams struct {
	Params   `form:"*"`
	Settings *CashBalanceSettingsParams `form:"settings"`
}
type CashBalanceSettingsParams struct {
	// Method for using the customer balance to pay outstanding
	// `customer_balance` PaymentIntents. If set to `automatic`, all available
	// funds will automatically be used to pay any outstanding PaymentIntent.
	// If set to `manual`, only customer balance funds from bank transfers
	// with a reference code matching
	// `payment_intent.next_action.display_bank_transfer_intructions.reference_code` will
	// automatically be used to pay the corresponding outstanding
	// PaymentIntent.
	ReconciliationMode *string `form:"reconciliation_mode"`
}
type CashBalanceSettings struct {
	// The configuration for how funds that land in the customer cash balance are reconciled.
	ReconciliationMode CashBalanceSettingsReconciliationMode `json:"reconciliation_mode"`
}

// A customer's `Cash balance` represents real funds. Customers can add funds to their cash balance by sending a bank transfer. These funds can be used for payment and can eventually be paid out to your bank account.
type CashBalance struct {
	APIResource
	// A hash of all cash balances available to this customer. You cannot delete a customer with any cash balances, even if the balance is 0.
	Available map[string]int64 `json:"available"`
	// The ID of the customer whose cash balance this object represents.
	Customer string `json:"customer"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object   string               `json:"object"`
	Settings *CashBalanceSettings `json:"settings"`
}
