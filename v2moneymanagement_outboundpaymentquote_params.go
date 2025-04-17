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
}

// Request details about the sender of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteFromParams struct {
	// Describes the FinancialAccount's currency drawn from.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
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
	// To which account the OutboundPayment is sent.
	Recipient *string `form:"recipient" json:"recipient"`
}

// Creates an OutboundPaymentQuote usable in an OutboundPayment.
type V2MoneyManagementOutboundPaymentQuoteParams struct {
	Params `form:"*"`
	// The "presentment amount" to be sent to the recipient.
	Amount *Amount `form:"amount" json:"amount"`
	// Method to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsParams `form:"delivery_options" json:"delivery_options,omitempty"`
	// Request details about the sender of an OutboundPaymentQuote.
	From *V2MoneyManagementOutboundPaymentQuoteFromParams `form:"from" json:"from"`
	// Request details about the recipient of an OutboundPaymentQuote.
	To *V2MoneyManagementOutboundPaymentQuoteToParams `form:"to" json:"to"`
}

// Method to be used to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentQuoteCreateDeliveryOptionsParams struct {
	// Open Enum. Method for bank account.
	BankAccount *string `form:"bank_account" json:"bank_account,omitempty"`
}

// Request details about the sender of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteCreateFromParams struct {
	// Describes the FinancialAccount's currency drawn from.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
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
