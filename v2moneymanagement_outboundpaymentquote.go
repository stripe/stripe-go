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

// Open Enum. Speed of the payout.
type V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsSpeed string

// List of values that V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsSpeed can take
const (
	V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsSpeedInstant         V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsSpeed = "instant"
	V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsSpeedNextBusinessDay V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsSpeed = "next_business_day"
	V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsSpeedStandard        V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsSpeed = "standard"
)

// The fee type.
type V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType string

// List of values that V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType can take
const (
	V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTypeCrossBorderPayoutFee V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType = "cross_border_payout_fee"
	V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTypeForeignExchangeFee   V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType = "foreign_exchange_fee"
	V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTypeInstantPayoutFee     V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType = "instant_payout_fee"
	V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTypeNextDayPayoutFee     V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType = "next_day_payout_fee"
	V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTypeRealTimePayoutFee    V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType = "real_time_payout_fee"
	V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTypeStandardPayoutFee    V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType = "standard_payout_fee"
	V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTypeWirePayoutFee        V2MoneyManagementOutboundPaymentQuoteEstimatedFeeType = "wire_payout_fee"
)

// The duration the exchange rate lock remains valid from creation time. Allowed value is five_minutes or none.
type V2MoneyManagementOutboundPaymentQuoteFxQuoteLockDuration string

// List of values that V2MoneyManagementOutboundPaymentQuoteFxQuoteLockDuration can take
const (
	V2MoneyManagementOutboundPaymentQuoteFxQuoteLockDurationFiveMinutes V2MoneyManagementOutboundPaymentQuoteFxQuoteLockDuration = "five_minutes"
	V2MoneyManagementOutboundPaymentQuoteFxQuoteLockDurationNone        V2MoneyManagementOutboundPaymentQuoteFxQuoteLockDuration = "none"
)

// Lock status of the quote. Transitions from active to expired once past the lock_expires_at timestamp. Value can be active, expired or none.
type V2MoneyManagementOutboundPaymentQuoteFxQuoteLockStatus string

// List of values that V2MoneyManagementOutboundPaymentQuoteFxQuoteLockStatus can take
const (
	V2MoneyManagementOutboundPaymentQuoteFxQuoteLockStatusActive  V2MoneyManagementOutboundPaymentQuoteFxQuoteLockStatus = "active"
	V2MoneyManagementOutboundPaymentQuoteFxQuoteLockStatusExpired V2MoneyManagementOutboundPaymentQuoteFxQuoteLockStatus = "expired"
	V2MoneyManagementOutboundPaymentQuoteFxQuoteLockStatusNone    V2MoneyManagementOutboundPaymentQuoteFxQuoteLockStatus = "none"
)

// Open Enum. ACH submission timing.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission string

// List of values that V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission can take
const (
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmissionNextDay V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission = "next_day"
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmissionSameDay V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission = "same_day"
)

// The transaction purpose for this ACH payment.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose string

// List of values that V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose can take
const (
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurposePayroll V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose = "payroll"
)

// The preferred networks to use for this OutboundPayment.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork string

// List of values that V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork can take
const (
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkACH         V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork = "ach"
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkBECS        V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork = "becs"
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkEft         V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork = "eft"
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkFedwire     V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork = "fedwire"
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkFPS         V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork = "fps"
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkNpp         V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork = "npp"
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkRTP         V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork = "rtp"
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkSEPACredit  V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork = "sepa_credit"
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkSEPAInstant V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork = "sepa_instant"
	V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkSwift       V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork = "swift"
)

// Delivery options to be used to send the OutboundPayment.
type V2MoneyManagementOutboundPaymentQuoteDeliveryOptions struct {
	// Open Enum. Method for bank account.
	BankAccount V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsBankAccount `json:"bank_account,omitempty"`
	// Open Enum. Speed of the payout.
	Speed V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsSpeed `json:"speed,omitempty"`
}

// Tax charged for this fee, if applicable. Value expressed as a decimal string in major units.
type V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTaxAmount struct {
	// Currency code.
	Currency Currency `json:"currency"`
	// Tax amount value represented as a decimal string in major units.
	ValueDecimal string `json:"value_decimal"`
}

// The estimated fees for the OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteEstimatedFee struct {
	// The fee amount for corresponding fee type.
	Amount Amount `json:"amount"`
	// Tax charged for this fee, if applicable. Value expressed as a decimal string in major units.
	TaxAmount *V2MoneyManagementOutboundPaymentQuoteEstimatedFeeTaxAmount `json:"tax_amount,omitempty"`
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
	// The duration the exchange rate lock remains valid from creation time. Allowed value is five_minutes or none.
	LockDuration V2MoneyManagementOutboundPaymentQuoteFxQuoteLockDuration `json:"lock_duration"`
	// Time at which the rate lock will expire, measured in seconds since the Unix epoch. Null when rate locking is not supported.
	LockExpiresAt time.Time `json:"lock_expires_at,omitempty"`
	// Lock status of the quote. Transitions from active to expired once past the lock_expires_at timestamp. Value can be active, expired or none.
	LockStatus V2MoneyManagementOutboundPaymentQuoteFxQuoteLockStatus `json:"lock_status"`
	// Key pair: from currency Value: exchange rate going from_currency -> to_currency.
	Rates map[string]*V2MoneyManagementOutboundPaymentQuoteFxQuoteRates `json:"rates"`
	// The currency that the transaction is exchanging to.
	ToCurrency Currency `json:"to_currency"`
}

// ACH-specific network options.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACH struct {
	// Open Enum. ACH submission timing.
	Submission V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHSubmission `json:"submission,omitempty"`
	// The transaction purpose for this ACH payment.
	TransactionPurpose V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACHTransactionPurpose `json:"transaction_purpose,omitempty"`
}

// Per-network configuration options.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptions struct {
	// ACH-specific network options.
	ACH *V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsACH `json:"ach,omitempty"`
}

// Options for bank account payout methods.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccount struct {
	// Per-network configuration options.
	PreferredNetworkOptions *V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptions `json:"preferred_network_options,omitempty"`
	// The preferred networks to use for this OutboundPayment.
	PreferredNetworks []V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetwork `json:"preferred_networks"`
}

// Payout method options for the OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptions struct {
	// Options for bank account payout methods.
	BankAccount *V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccount `json:"bank_account,omitempty"`
}

// Details about the recipient of an OutboundPaymentQuote.
type V2MoneyManagementOutboundPaymentQuoteTo struct {
	// The monetary amount being credited to the destination.
	Credited Amount `json:"credited"`
	// The payout method which the OutboundPayment uses to send payout.
	PayoutMethod string `json:"payout_method"`
	// Payout method options for the OutboundPaymentQuote.
	PayoutMethodOptions *V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptions `json:"payout_method_options,omitempty"`
	// To which account the OutboundPayment is sent.
	Recipient string `json:"recipient"`
}

// OutboundPaymentQuote represents a quote that provides fee and amount estimates for OutboundPayment.
type V2MoneyManagementOutboundPaymentQuote struct {
	APIResource
	// The "presentment amount" for the OutboundPaymentQuote.
	Amount Amount `json:"amount"`
	// Time at which the OutboundPaymentQuote was created.
	// Represented as a RFC 3339 date & time UTC value in millisecond precision, for example: 2022-09-18T13:22:18.123Z.
	Created time.Time `json:"created"`
	// Delivery options to be used to send the OutboundPayment.
	DeliveryOptions *V2MoneyManagementOutboundPaymentQuoteDeliveryOptions `json:"delivery_options,omitempty"`
	// The estimated fees for the OutboundPaymentQuote.
	EstimatedFees []*V2MoneyManagementOutboundPaymentQuoteEstimatedFee `json:"estimated_fees"`
	// Details about the sender of an OutboundPaymentQuote.
	From *V2MoneyManagementOutboundPaymentQuoteFrom `json:"from"`
	// The underlying FXQuote details for the OutboundPaymentQuote.
	FxQuote *V2MoneyManagementOutboundPaymentQuoteFxQuote `json:"fx_quote"`
	// Unique identifier for the OutboundPaymentQuote.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Details about the recipient of an OutboundPaymentQuote.
	To *V2MoneyManagementOutboundPaymentQuoteTo `json:"to"`
}
