//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Optional fields for `us_bank_account`.
type TestHelpersTreasuryReceivedCreditInitiatingPaymentMethodDetailsUSBankAccountParams struct {
	// The bank account holder's name.
	AccountHolderName *string `form:"account_holder_name"`
	// The bank account number.
	AccountNumber *string `form:"account_number"`
	// The bank account's routing number.
	RoutingNumber *string `form:"routing_number"`
}

// Initiating payment method details for the object.
type TestHelpersTreasuryReceivedCreditInitiatingPaymentMethodDetailsParams struct {
	// The source type.
	Type *string `form:"type"`
	// Optional fields for `us_bank_account`.
	USBankAccount *TestHelpersTreasuryReceivedCreditInitiatingPaymentMethodDetailsUSBankAccountParams `form:"us_bank_account"`
}

// Use this endpoint to simulate a test mode ReceivedCredit initiated by a third party. In live mode, you can't directly create ReceivedCredits initiated by third parties.
type TestHelpersTreasuryReceivedCreditParams struct {
	Params `form:"*"`
	// Amount (in cents) to be transferred.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// The FinancialAccount to send funds to.
	FinancialAccount *string `form:"financial_account"`
	// Initiating payment method details for the object.
	InitiatingPaymentMethodDetails *TestHelpersTreasuryReceivedCreditInitiatingPaymentMethodDetailsParams `form:"initiating_payment_method_details"`
	// The rails used for the object.
	Network *string `form:"network"`
}
