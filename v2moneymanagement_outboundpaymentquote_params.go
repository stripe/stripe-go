//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Method to be used to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsParams struct {
	// Open Enum. Method for bank account.
	BankAccount *string `form:"bank_account" json:"bank_account,omitempty"`
	// Open Enum. Speed of the payout.
	Speed *string `form:"speed" json:"speed,omitempty"`
}

// Request details about the sender of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteFromParams struct {
	// Describes the FinancialAccount's currency drawn from.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// ACH-specific network options.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams struct {
	// Open Enum. ACH submission timing.
	Submission *string `form:"submission" json:"submission,omitempty"`
	// The transaction purpose for this ACH payment.
	TransactionPurpose *string `form:"transaction_purpose" json:"transaction_purpose,omitempty"`
}

// Per-network configuration options.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams struct {
	// ACH-specific network options.
	ACH *V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams `form:"ach" json:"ach,omitempty"`
}

// Options for bank account payout methods.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountParams struct {
	// Per-network configuration options.
	PreferredNetworkOptions *V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams `form:"preferred_network_options" json:"preferred_network_options,omitempty"`
	// The preferred networks to use for this OutboundPayment.
	PreferredNetworks []*string `form:"preferred_networks" json:"preferred_networks"`
}

// Payout method options for the OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsParams struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
}

// Request details about the recipient of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteToParams struct {
	// Describes the currency to send to the recipient.
	// If included, this currency must match a currency supported by the destination.
	// Can be omitted in the following cases:
	// - destination only supports one currency
	// - destination supports multiple currencies and one of the currencies matches the FA currency
	// - destination supports multiple currencies and one of the currencies matches the presentment currency
	// Note - when both FA currency and presentment currency are supported, we pick the FA currency to minimize FX.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Payout method options for the OutboundPaymentQuote.
	PayoutMethodOptions *V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsParams `form:"payout_method_options" json:"payout_method_options,omitempty"`
	// To which account the OutboundPayment is sent.
	Recipient *string `form:"recipient" json:"recipient"`
}

// Creates an OutboundPaymentQuote usable in an OutboundPayment.
type V2MoneyManagementOutboundPaymentQuoteParams struct {
	Params `form:"*"`
	// The "presentment amount" to be sent to the recipient.
	Amount *Amount `form:"amount" json:"amount,omitempty"`
	// Method to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsParams `form:"delivery_options" json:"delivery_options,omitempty"`
	// Request details about the sender of an OutboundPaymentQuote.
	From *V2MoneyManagementOutboundPaymentQuoteFromParams `form:"from" json:"from,omitempty"`
	// Request details about the recipient of an OutboundPaymentQuote.
	To *V2MoneyManagementOutboundPaymentQuoteToParams `form:"to" json:"to,omitempty"`
}

// Method to be used to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentQuoteCreateDeliveryOptionsParams struct {
	// Open Enum. Method for bank account.
	BankAccount *string `form:"bank_account" json:"bank_account,omitempty"`
	// Open Enum. Speed of the payout.
	Speed *string `form:"speed" json:"speed,omitempty"`
}

// Request details about the sender of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteCreateFromParams struct {
	// Describes the FinancialAccount's currency drawn from.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// ACH-specific network options.
type V2MoneyManagementOutboundPaymentQuoteCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams struct {
	// Open Enum. ACH submission timing.
	Submission *string `form:"submission" json:"submission,omitempty"`
	// The transaction purpose for this ACH payment.
	TransactionPurpose *string `form:"transaction_purpose" json:"transaction_purpose,omitempty"`
}

// Per-network configuration options.
type V2MoneyManagementOutboundPaymentQuoteCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams struct {
	// ACH-specific network options.
	ACH *V2MoneyManagementOutboundPaymentQuoteCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHParams `form:"ach" json:"ach,omitempty"`
}

// Options for bank account payout methods.
type V2MoneyManagementOutboundPaymentQuoteCreateToPayoutMethodOptionsBankAccountParams struct {
	// Per-network configuration options.
	PreferredNetworkOptions *V2MoneyManagementOutboundPaymentQuoteCreateToPayoutMethodOptionsBankAccountPreferredNetworkOptionsParams `form:"preferred_network_options" json:"preferred_network_options,omitempty"`
	// The preferred networks to use for this OutboundPayment.
	PreferredNetworks []*string `form:"preferred_networks" json:"preferred_networks"`
}

// Payout method options for the OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteCreateToPayoutMethodOptionsParams struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementOutboundPaymentQuoteCreateToPayoutMethodOptionsBankAccountParams `form:"bank_account" json:"bank_account,omitempty"`
}

// Request details about the recipient of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteCreateToParams struct {
	// Describes the currency to send to the recipient.
	// If included, this currency must match a currency supported by the destination.
	// Can be omitted in the following cases:
	// - destination only supports one currency
	// - destination supports multiple currencies and one of the currencies matches the FA currency
	// - destination supports multiple currencies and one of the currencies matches the presentment currency
	// Note - when both FA currency and presentment currency are supported, we pick the FA currency to minimize FX.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// Payout method options for the OutboundPaymentQuote.
	PayoutMethodOptions *V2MoneyManagementOutboundPaymentQuoteCreateToPayoutMethodOptionsParams `form:"payout_method_options" json:"payout_method_options,omitempty"`
	// To which account the OutboundPayment is sent.
	Recipient *string `form:"recipient" json:"recipient"`
}

// Creates an OutboundPaymentQuote usable in an OutboundPayment.
type V2MoneyManagementOutboundPaymentQuoteCreateParams struct {
	Params `form:"*"`
	// The "presentment amount" to be sent to the recipient.
	Amount *Amount `form:"amount" json:"amount"`
	// Method to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentQuoteCreateDeliveryOptionsParams `form:"delivery_options" json:"delivery_options,omitempty"`
	// Request details about the sender of an OutboundPaymentQuote.
	From *V2MoneyManagementOutboundPaymentQuoteCreateFromParams `form:"from" json:"from"`
	// Request details about the recipient of an OutboundPaymentQuote.
	To *V2MoneyManagementOutboundPaymentQuoteCreateToParams `form:"to" json:"to"`
}

// Retrieves the details of an existing OutboundPaymentQuote by passing the unique OutboundPaymentQuote ID.
type V2MoneyManagementOutboundPaymentQuoteRetrieveParams struct {
	Params `form:"*"`
}
