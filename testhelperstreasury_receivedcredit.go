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

// Optional fields for `ach`.
type TestHelpersTreasuryReceivedCreditNetworkDetailsACHParams struct {
	// ACH Addenda record
	Addenda *string `form:"addenda"`
}

// Details about the network used for the ReceivedCredit.
type TestHelpersTreasuryReceivedCreditNetworkDetailsParams struct {
	// Optional fields for `ach`.
	ACH *TestHelpersTreasuryReceivedCreditNetworkDetailsACHParams `form:"ach"`
	// The type of flow that originated the ReceivedCredit.
	Type *string `form:"type"`
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
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The FinancialAccount to send funds to.
	FinancialAccount *string `form:"financial_account"`
	// Initiating payment method details for the object.
	InitiatingPaymentMethodDetails *TestHelpersTreasuryReceivedCreditInitiatingPaymentMethodDetailsParams `form:"initiating_payment_method_details"`
	// The rails used for the object.
	Network *string `form:"network"`
	// Details about the network used for the ReceivedCredit.
	NetworkDetails *TestHelpersTreasuryReceivedCreditNetworkDetailsParams `form:"network_details"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryReceivedCreditParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
