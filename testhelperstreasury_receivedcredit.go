//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Optional fields for `us_bank_account`.
type TestHelpersTreasuryReceivedCreditInitiatingPaymentMethodDetailsUSBankAccountParams struct {
	// The bank account holder's name.
	AccountHolderName *string `form:"account_holder_name" json:"account_holder_name,omitempty"`
	// The bank account number.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// The bank account's routing number.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
}

// Initiating payment method details for the object.
type TestHelpersTreasuryReceivedCreditInitiatingPaymentMethodDetailsParams struct {
	// The source type.
	Type *string `form:"type" json:"type"`
	// Optional fields for `us_bank_account`.
	USBankAccount *TestHelpersTreasuryReceivedCreditInitiatingPaymentMethodDetailsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// Optional fields for `ach`.
type TestHelpersTreasuryReceivedCreditNetworkDetailsACHParams struct {
	// ACH Addenda record
	Addenda *string `form:"addenda" json:"addenda,omitempty"`
}

// Details about the network used for the ReceivedCredit.
type TestHelpersTreasuryReceivedCreditNetworkDetailsParams struct {
	// Optional fields for `ach`.
	ACH *TestHelpersTreasuryReceivedCreditNetworkDetailsACHParams `form:"ach" json:"ach,omitempty"`
	// The type of flow that originated the ReceivedCredit.
	Type *string `form:"type" json:"type"`
}

// Use this endpoint to simulate a test mode ReceivedCredit initiated by a third party. In live mode, you can't directly create ReceivedCredits initiated by third parties.
type TestHelpersTreasuryReceivedCreditParams struct {
	Params `form:"*"`
	// Amount (in cents) to be transferred.
	Amount *int64 `form:"amount" json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The FinancialAccount to send funds to.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// Initiating payment method details for the object.
	InitiatingPaymentMethodDetails *TestHelpersTreasuryReceivedCreditInitiatingPaymentMethodDetailsParams `form:"initiating_payment_method_details" json:"initiating_payment_method_details,omitempty"`
	// Specifies the network rails to be used. If not set, will default to the PaymentMethod's preferred network. See the [docs](https://docs.stripe.com/treasury/money-movement/timelines) to learn more about money movement timelines for each network type.
	Network *string `form:"network" json:"network"`
	// Details about the network used for the ReceivedCredit.
	NetworkDetails *TestHelpersTreasuryReceivedCreditNetworkDetailsParams `form:"network_details" json:"network_details,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryReceivedCreditParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Optional fields for `us_bank_account`.
type TestHelpersTreasuryReceivedCreditCreateInitiatingPaymentMethodDetailsUSBankAccountParams struct {
	// The bank account holder's name.
	AccountHolderName *string `form:"account_holder_name" json:"account_holder_name,omitempty"`
	// The bank account number.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// The bank account's routing number.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
}

// Initiating payment method details for the object.
type TestHelpersTreasuryReceivedCreditCreateInitiatingPaymentMethodDetailsParams struct {
	// The source type.
	Type *string `form:"type" json:"type"`
	// Optional fields for `us_bank_account`.
	USBankAccount *TestHelpersTreasuryReceivedCreditCreateInitiatingPaymentMethodDetailsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// Optional fields for `ach`.
type TestHelpersTreasuryReceivedCreditCreateNetworkDetailsACHParams struct {
	// ACH Addenda record
	Addenda *string `form:"addenda" json:"addenda,omitempty"`
}

// Details about the network used for the ReceivedCredit.
type TestHelpersTreasuryReceivedCreditCreateNetworkDetailsParams struct {
	// Optional fields for `ach`.
	ACH *TestHelpersTreasuryReceivedCreditCreateNetworkDetailsACHParams `form:"ach" json:"ach,omitempty"`
	// The type of flow that originated the ReceivedCredit.
	Type *string `form:"type" json:"type"`
}

// Use this endpoint to simulate a test mode ReceivedCredit initiated by a third party. In live mode, you can't directly create ReceivedCredits initiated by third parties.
type TestHelpersTreasuryReceivedCreditCreateParams struct {
	Params `form:"*"`
	// Amount (in cents) to be transferred.
	Amount *int64 `form:"amount" json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description" json:"description,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The FinancialAccount to send funds to.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
	// Initiating payment method details for the object.
	InitiatingPaymentMethodDetails *TestHelpersTreasuryReceivedCreditCreateInitiatingPaymentMethodDetailsParams `form:"initiating_payment_method_details" json:"initiating_payment_method_details,omitempty"`
	// Specifies the network rails to be used. If not set, will default to the PaymentMethod's preferred network. See the [docs](https://docs.stripe.com/treasury/money-movement/timelines) to learn more about money movement timelines for each network type.
	Network *string `form:"network" json:"network"`
	// Details about the network used for the ReceivedCredit.
	NetworkDetails *TestHelpersTreasuryReceivedCreditCreateNetworkDetailsParams `form:"network_details" json:"network_details,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersTreasuryReceivedCreditCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
