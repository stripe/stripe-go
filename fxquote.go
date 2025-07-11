//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The duration the exchange rate quote remains valid from creation time. Allowed values are none, hour, and day. Note that for the test mode API available in alpha, you can request an extended quote, but it won't be usable for any transactions.
type FxQuoteLockDuration string

// List of values that FxQuoteLockDuration can take
const (
	FxQuoteLockDurationDay         FxQuoteLockDuration = "day"
	FxQuoteLockDurationFiveMinutes FxQuoteLockDuration = "five_minutes"
	FxQuoteLockDurationHour        FxQuoteLockDuration = "hour"
	FxQuoteLockDurationNone        FxQuoteLockDuration = "none"
)

// Lock status of the quote. Transitions from active to expired once past the lock_expires_at timestamp.
//
// Can return value none, active, or expired.
type FxQuoteLockStatus string

// List of values that FxQuoteLockStatus can take
const (
	FxQuoteLockStatusActive  FxQuoteLockStatus = "active"
	FxQuoteLockStatusExpired FxQuoteLockStatus = "expired"
	FxQuoteLockStatusNone    FxQuoteLockStatus = "none"
)

// The reference rate provider.
type FxQuoteRatesRateDetailsReferenceRateProvider string

// List of values that FxQuoteRatesRateDetailsReferenceRateProvider can take
const (
	FxQuoteRatesRateDetailsReferenceRateProviderEcb FxQuoteRatesRateDetailsReferenceRateProvider = "ecb"
)

// The transaction type for which the FX Quote will be used.
//
// Can be 'payment' or 'transfer'.
type FxQuoteUsageType string

// List of values that FxQuoteUsageType can take
const (
	FxQuoteUsageTypePayment  FxQuoteUsageType = "payment"
	FxQuoteUsageTypeTransfer FxQuoteUsageType = "transfer"
)

// Returns a list of FX quotes that have been issued. The FX quotes are returned in sorted order, with the most recent FX quotes appearing first.
type FxQuoteListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *FxQuoteListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The payment transaction details that are intended for the FX Quote.
type FxQuoteUsagePaymentParams struct {
	// The Stripe account ID that the funds will be transferred to.
	//
	// This field should match the account ID that would be used in the PaymentIntent's transfer_data[destination] field.
	Destination *string `form:"destination"`
	// The Stripe account ID that these funds are intended for.
	//
	// This field should match the account ID that would be used in the PaymentIntent's on_behalf_of field.
	OnBehalfOf *string `form:"on_behalf_of"`
}

// The transfer transaction details that are intended for the FX Quote.
type FxQuoteUsageTransferParams struct {
	// The Stripe account ID that the funds will be transferred to.
	//
	// This field should match the account ID that would be used in the Transfer's destination field.
	Destination *string `form:"destination"`
}

// The usage specific information for the quote.
type FxQuoteUsageParams struct {
	// The payment transaction details that are intended for the FX Quote.
	Payment *FxQuoteUsagePaymentParams `form:"payment"`
	// The transfer transaction details that are intended for the FX Quote.
	Transfer *FxQuoteUsageTransferParams `form:"transfer"`
	// Which transaction the FX Quote will be used for
	//
	// Can be “payment” | “transfer”
	Type *string `form:"type"`
}

// Creates an FX Quote object
type FxQuoteParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A list of three letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be [supported currencies](https://stripe.com/docs/currencies).
	FromCurrencies []*string `form:"from_currencies"`
	// The duration that you wish the quote to be locked for. The quote will be usable for the duration specified. The default is `none`. The maximum is 1 day.
	LockDuration *string `form:"lock_duration"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	ToCurrency *string `form:"to_currency"`
	// The usage specific information for the quote.
	Usage *FxQuoteUsageParams `form:"usage"`
}

// AddExpand appends a new field to expand.
func (p *FxQuoteParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The payment transaction details that are intended for the FX Quote.
type FxQuoteCreateUsagePaymentParams struct {
	// The Stripe account ID that the funds will be transferred to.
	//
	// This field should match the account ID that would be used in the PaymentIntent's transfer_data[destination] field.
	Destination *string `form:"destination"`
	// The Stripe account ID that these funds are intended for.
	//
	// This field should match the account ID that would be used in the PaymentIntent's on_behalf_of field.
	OnBehalfOf *string `form:"on_behalf_of"`
}

// The transfer transaction details that are intended for the FX Quote.
type FxQuoteCreateUsageTransferParams struct {
	// The Stripe account ID that the funds will be transferred to.
	//
	// This field should match the account ID that would be used in the Transfer's destination field.
	Destination *string `form:"destination"`
}

// The usage specific information for the quote.
type FxQuoteCreateUsageParams struct {
	// The payment transaction details that are intended for the FX Quote.
	Payment *FxQuoteCreateUsagePaymentParams `form:"payment"`
	// The transfer transaction details that are intended for the FX Quote.
	Transfer *FxQuoteCreateUsageTransferParams `form:"transfer"`
	// Which transaction the FX Quote will be used for
	//
	// Can be “payment” | “transfer”
	Type *string `form:"type"`
}

// Creates an FX Quote object
type FxQuoteCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A list of three letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be [supported currencies](https://stripe.com/docs/currencies).
	FromCurrencies []*string `form:"from_currencies"`
	// The duration that you wish the quote to be locked for. The quote will be usable for the duration specified. The default is `none`. The maximum is 1 day.
	LockDuration *string `form:"lock_duration"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	ToCurrency *string `form:"to_currency"`
	// The usage specific information for the quote.
	Usage *FxQuoteCreateUsageParams `form:"usage"`
}

// AddExpand appends a new field to expand.
func (p *FxQuoteCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieve an FX Quote object
type FxQuoteRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *FxQuoteRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type FxQuoteRatesRateDetails struct {
	// The rate for the currency pair.
	BaseRate float64 `json:"base_rate"`
	// The fee for locking the conversion rates.
	DurationPremium float64 `json:"duration_premium"`
	// The FX fee for the currency pair.
	FxFeeRate float64 `json:"fx_fee_rate"`
	// A reference rate for the currency pair provided by the reference rate provider.
	ReferenceRate float64 `json:"reference_rate"`
	// The reference rate provider.
	ReferenceRateProvider FxQuoteRatesRateDetailsReferenceRateProvider `json:"reference_rate_provider"`
}

// Information about the rates.
type FxQuoteRates struct {
	// The rate that includes the FX fee rate.
	ExchangeRate float64                  `json:"exchange_rate"`
	RateDetails  *FxQuoteRatesRateDetails `json:"rate_details"`
}

// The details required to use an FX Quote for a payment
type FxQuoteUsagePayment struct {
	// The Stripe account ID that the funds will be transferred to.
	//
	// This field should match the account ID that would be used in the PaymentIntent's transfer_data[destination] field.
	Destination string `json:"destination"`
	// The Stripe account ID that these funds are intended for.
	//
	// This field must match the account ID that would be used in the PaymentIntent's on_behalf_of field.
	OnBehalfOf string `json:"on_behalf_of"`
}

// The details required to use an FX Quote for a transfer
type FxQuoteUsageTransfer struct {
	// The Stripe account ID that the funds will be transferred to.
	//
	// This field should match the account ID that would be used in the Transfer's destination field.
	Destination string `json:"destination"`
}
type FxQuoteUsage struct {
	// The details required to use an FX Quote for a payment
	Payment *FxQuoteUsagePayment `json:"payment"`
	// The details required to use an FX Quote for a transfer
	Transfer *FxQuoteUsageTransfer `json:"transfer"`
	// The transaction type for which the FX Quote will be used.
	//
	// Can be 'payment' or 'transfer'.
	Type FxQuoteUsageType `json:"type"`
}

// The FX Quotes API provides three functions:
// - View Stripe's current exchange rate for any given currency pair.
// - Extend quoted rates for a 1-hour period or a 24-hour period, minimizing uncertainty from FX fluctuations.
// - Preview the FX fees Stripe will charge on your FX transaction, allowing you to anticipate specific settlement amounts before payment costs.
//
// [View the docs](https://docs.stripe.com/payments/currencies/localize-prices/fx-quotes-api)
type FxQuote struct {
	APIResource
	// Time at which the quote was created, measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The duration the exchange rate quote remains valid from creation time. Allowed values are none, hour, and day. Note that for the test mode API available in alpha, you can request an extended quote, but it won't be usable for any transactions.
	LockDuration FxQuoteLockDuration `json:"lock_duration"`
	// Time at which the quote will expire, measured in seconds since the Unix epoch.
	//
	// If lock_duration is set to ‘none' this field will be set to null.
	LockExpiresAt int64 `json:"lock_expires_at"`
	// Lock status of the quote. Transitions from active to expired once past the lock_expires_at timestamp.
	//
	// Can return value none, active, or expired.
	LockStatus FxQuoteLockStatus `json:"lock_status"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Information about the rates.
	Rates map[string]*FxQuoteRates `json:"rates"`
	// The currency to convert into, typically this is the currency that you want to settle to your Stripe balance. Three-letter ISO currency code, in lowercase. Must be a supported currency.
	ToCurrency Currency      `json:"to_currency"`
	Usage      *FxQuoteUsage `json:"usage"`
}

// FxQuoteList is a list of FxQuotes as retrieved from a list endpoint.
type FxQuoteList struct {
	APIResource
	ListMeta
	Data []*FxQuote `json:"data"`
}
