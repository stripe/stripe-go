//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Method to be used to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentsQuoteDeliveryOptionsParams struct {
	// Open Enum. Method for bank account.
	BankAccount *string `form:"bank_account" json:"bank_account,omitempty"`
}

// Request details about the sender of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentsQuoteFromParams struct {
	// Describes the FinancialAccount's currency drawn from.
	Currency *string `form:"currency" json:"currency"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount *string `form:"financial_account" json:"financial_account"`
}

// Request details about the recipient of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentsQuoteToParams struct {
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
type V2MoneyManagementOutboundPaymentsQuoteParams struct {
	Params `form:"*"`
	// The "presentment amount" to be sent to the recipient.
	Amount *Amount `form:"amount" json:"amount"`
	// Method to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentsQuoteDeliveryOptionsParams `form:"delivery_options" json:"delivery_options,omitempty"`
	// Request details about the sender of an OutboundPaymentQuote.
	From *V2MoneyManagementOutboundPaymentsQuoteFromParams `form:"from" json:"from"`
	// Request details about the recipient of an OutboundPaymentQuote.
	To *V2MoneyManagementOutboundPaymentsQuoteToParams `form:"to" json:"to"`
}
