//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Optional fields for `us_bank_account`.
type TestHelpersTreasuryReceivedDebitInitiatingPaymentMethodDetailsUSBankAccountParams struct {
	// The bank account holder's name.
	AccountHolderName *string `form:"account_holder_name"`
	// The bank account number.
	AccountNumber *string `form:"account_number"`
	// The bank account's routing number.
	RoutingNumber *string `form:"routing_number"`
}

// Initiating payment method details for the object.
type TestHelpersTreasuryReceivedDebitInitiatingPaymentMethodDetailsParams struct {
	// The source type.
	Type *string `form:"type"`
	// Optional fields for `us_bank_account`.
	USBankAccount *TestHelpersTreasuryReceivedDebitInitiatingPaymentMethodDetailsUSBankAccountParams `form:"us_bank_account"`
}

// Use this endpoint to simulate a test mode ReceivedDebit initiated by a third party. In live mode, you can't directly create ReceivedDebits initiated by third parties.
type TestHelpersTreasuryReceivedDebitParams struct {
	Params `form:"*"`
	// Amount (in cents) to be transferred.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// The FinancialAccount to pull funds from.
	FinancialAccount *string `form:"financial_account"`
	// Initiating payment method details for the object.
	InitiatingPaymentMethodDetails *TestHelpersTreasuryReceivedDebitInitiatingPaymentMethodDetailsParams `form:"initiating_payment_method_details"`
	// The rails used for the object.
	Network *string `form:"network"`
}
