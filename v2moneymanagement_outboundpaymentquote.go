//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Open Enum. Method for bank account.
type V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsBankAccount string

// List of values that V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsBankAccount can take
const (
	V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsBankAccountAutomatic V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsBankAccount = "automatic"
	V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsBankAccountLocal     V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsBankAccount = "local"
	V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsBankAccountWire      V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsBankAccount = "wire"
)

// The fee type.
type V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType string

// List of values that V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType can take
const (
	V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTypeCrossBorderFee V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType = "cross_border_fee"
	V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTypeFxFee          V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType = "fx_fee"
	V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTypePayoutFee      V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType = "payout_fee"
)

// Delivery options to be used to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentQuoteDeliveryOptions struct {
	// Open Enum. Method for bank account.
	BankAccount V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsBankAccount `json:"bank_account"`
}

// The estimated fees for the OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteEstimatedFee struct {
	// The fee amount for corresponding fee type.
	Amount Amount `json:"amount"`
	// The fee type.
	Type V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType `json:"type"`
}

// Details about the sender of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteFrom struct {
	// The monetary amount debited from the sender, only set on responses.
	Debited Amount `json:"debited"`
	// The FinancialAccount that funds were pulled from.
	FinancialAccount string `json:"financial_account"`
}

// Key pair: from currency Value: exchange rate going from_currency -> to_currency.
type V2MoneyManagementOutboundPaymentQuoteFxQuoteRates struct {
	// The exchange rate going from_currency -> to_currency.
	ExchangeRate string `json:"exchange_rate"`
}

// The underlying FXQuote details for the OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteFxQuote struct {
	// Key pair: from currency Value: exchange rate going from_currency -> to_currency.
	Rates map[string]*V2MoneyManagementOutboundPaymentQuoteFxQuoteRates `json:"rates"`
	// The currency that the transaction is exchanging to.
	ToCurrency Currency `json:"to_currency"`
}

// Details about the recipient of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteTo struct {
	// The monetary amount being credited to the destination.
	Credited Amount `json:"credited"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod string `json:"payout_method"`
	// To which account the OutboundPayment is sent.
	Recipient string `json:"recipient"`
}

// OutboundPaymentQuote represents a quote
type V2MoneyManagementOutboundPaymentQuote struct {
	APIResource
	// The "presentment amount" for the OutboundPaymentQuote.
	Amount Amount `json:"amount"`
	// Time at which the OutboundPaymentQuote was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Delivery options to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentQuoteDeliveryOptions `json:"delivery_options"`
	// The estimated fees for the OutboundPaymentQuote.
	EstimatedFees []*V2MoneyManagementOutboundPaymentQuoteEstimatedFee `json:"estimated_fees"`
	// Details about the sender of an OutboundPaymentQuote.
	From *V2MoneyManagementOutboundPaymentQuoteFrom `json:"from"`
	// The underlying FXQuote details for the OutboundPaymentQuote.
	FxQuote *V2MoneyManagementOutboundPaymentQuoteFxQuote `json:"fx_quote"`
	// Unique identifier for the OutboundPaymentQuote.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Details about the recipient of an OutboundPaymentQuote.
	To *V2MoneyManagementOutboundPaymentQuoteTo `json:"to"`
}
