//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Indicates whether the customer was present in your checkout flow during this payment.
type PaymentAttemptRecordCustomerPresence string

// List of values that PaymentAttemptRecordCustomerPresence can take
const (
	PaymentAttemptRecordCustomerPresenceOffSession PaymentAttemptRecordCustomerPresence = "off_session"
	PaymentAttemptRecordCustomerPresenceOnSession  PaymentAttemptRecordCustomerPresence = "on_session"
)

// Card brand. Can be `amex`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
type PaymentAttemptRecordPaymentMethodDetailsCardBrand string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardBrand can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardBrandAmex            PaymentAttemptRecordPaymentMethodDetailsCardBrand = "amex"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandCartesBancaires PaymentAttemptRecordPaymentMethodDetailsCardBrand = "cartes_bancaires"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandDiners          PaymentAttemptRecordPaymentMethodDetailsCardBrand = "diners"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandDiscover        PaymentAttemptRecordPaymentMethodDetailsCardBrand = "discover"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandEFTPOSAU        PaymentAttemptRecordPaymentMethodDetailsCardBrand = "eftpos_au"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandInterac         PaymentAttemptRecordPaymentMethodDetailsCardBrand = "interac"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandJCB             PaymentAttemptRecordPaymentMethodDetailsCardBrand = "jcb"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandLink            PaymentAttemptRecordPaymentMethodDetailsCardBrand = "link"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandMastercard      PaymentAttemptRecordPaymentMethodDetailsCardBrand = "mastercard"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandUnionpay        PaymentAttemptRecordPaymentMethodDetailsCardBrand = "unionpay"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandUnknown         PaymentAttemptRecordPaymentMethodDetailsCardBrand = "unknown"
	PaymentAttemptRecordPaymentMethodDetailsCardBrandVisa            PaymentAttemptRecordPaymentMethodDetailsCardBrand = "visa"
)

type PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1CheckFail        PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check = "fail"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1CheckPass        PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check = "pass"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1CheckUnavailable PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check = "unavailable"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1CheckUnchecked   PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check = "unchecked"
)

type PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckFail        PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "fail"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckPass        PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "pass"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckUnavailable PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "unavailable"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheckUnchecked   PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck = "unchecked"
)

type PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheckFail        PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck = "fail"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheckPass        PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck = "pass"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheckUnavailable PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck = "unavailable"
	PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheckUnchecked   PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck = "unchecked"
)

// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
type PaymentAttemptRecordPaymentMethodDetailsCardFunding string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardFunding can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardFundingCredit  PaymentAttemptRecordPaymentMethodDetailsCardFunding = "credit"
	PaymentAttemptRecordPaymentMethodDetailsCardFundingDebit   PaymentAttemptRecordPaymentMethodDetailsCardFunding = "debit"
	PaymentAttemptRecordPaymentMethodDetailsCardFundingPrepaid PaymentAttemptRecordPaymentMethodDetailsCardFunding = "prepaid"
	PaymentAttemptRecordPaymentMethodDetailsCardFundingUnknown PaymentAttemptRecordPaymentMethodDetailsCardFunding = "unknown"
)

// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
type PaymentAttemptRecordPaymentMethodDetailsCardNetwork string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardNetwork can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkAmex            PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "amex"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkCartesBancaires PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "cartes_bancaires"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkDiners          PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "diners"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkDiscover        PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "discover"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkEFTPOSAU        PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "eftpos_au"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkInterac         PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "interac"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkJCB             PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "jcb"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkLink            PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "link"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkMastercard      PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "mastercard"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkUnionpay        PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "unionpay"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkUnknown         PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "unknown"
	PaymentAttemptRecordPaymentMethodDetailsCardNetworkVisa            PaymentAttemptRecordPaymentMethodDetailsCardNetwork = "visa"
)

type PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

type PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultAuthenticated       PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultExempted            PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "exempted"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultFailed              PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "failed"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultNotSupported        PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultProcessingError     PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

type PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonBypassed            PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonCanceled            PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReasonRejected            PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

type PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion string

// List of values that PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion can take
const (
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion102 PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion = "1.0.2"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion210 PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion = "2.1.0"
	PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion220 PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion = "2.2.0"
)

// The type of Payment Method used for this payment attempt.
type PaymentAttemptRecordPaymentMethodDetailsType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsType can take
const (
	PaymentAttemptRecordPaymentMethodDetailsTypeCard   PaymentAttemptRecordPaymentMethodDetailsType = "card"
	PaymentAttemptRecordPaymentMethodDetailsTypeCustom PaymentAttemptRecordPaymentMethodDetailsType = "custom"
)

// List all the Payment Attempt Records attached to the specified Payment Record.
type PaymentAttemptRecordListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The ID of the Payment Record.
	PaymentRecord *string `form:"payment_record"`
}

// AddExpand appends a new field to expand.
func (p *PaymentAttemptRecordListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a Payment Attempt Record with the given ID
type PaymentAttemptRecordParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PaymentAttemptRecordParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmountCanceled struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmountFailed struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmountGuaranteed struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentAttemptRecordAmountRequested struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// Customer information for this payment.
type PaymentAttemptRecordCustomerDetails struct {
	// ID of the Stripe Customer associated with this payment.
	Customer string `json:"customer"`
	// The customer's email address.
	Email string `json:"email"`
	// The customer's name.
	Name string `json:"name"`
	// The customer's phone number.
	Phone string `json:"phone"`
}

// The billing details associated with the method of payment.
type PaymentAttemptRecordPaymentMethodDetailsBillingDetails struct {
	// A representation of a physical address.
	Address *Address `json:"address"`
	// The billing email associated with the method of payment.
	Email string `json:"email"`
	// The billing name associated with the method of payment.
	Name string `json:"name"`
	// The billing phone number associated with the method of payment.
	Phone string `json:"phone"`
}

// Check results by Card networks on Card address and CVC at time of payment.
type PaymentAttemptRecordPaymentMethodDetailsCardChecks struct {
	AddressLine1Check      PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressLine1Check      `json:"address_line1_check"`
	AddressPostalCodeCheck PaymentAttemptRecordPaymentMethodDetailsCardChecksAddressPostalCodeCheck `json:"address_postal_code_check"`
	CVCCheck               PaymentAttemptRecordPaymentMethodDetailsCardChecksCVCCheck               `json:"cvc_check"`
}

// If this card has network token credentials, this contains the details of the network token credentials.
type PaymentAttemptRecordPaymentMethodDetailsCardNetworkToken struct {
	Used bool `json:"used"`
}

// Populated if this transaction used 3D Secure authentication.
type PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure struct {
	AuthenticationFlow PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	Result             PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResult             `json:"result"`
	ResultReason       PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureResultReason       `json:"result_reason"`
	Version            PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecureVersion            `json:"version"`
}

// Details of the card used for this payment attempt.
type PaymentAttemptRecordPaymentMethodDetailsCard struct {
	// Card brand. Can be `amex`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Brand PaymentAttemptRecordPaymentMethodDetailsCardBrand `json:"brand"`
	// When using manual capture, a future timestamp at which the charge will be automatically refunded if uncaptured.
	CaptureBefore int64 `json:"capture_before"`
	// Check results by Card networks on Card address and CVC at time of payment.
	Checks *PaymentAttemptRecordPaymentMethodDetailsCardChecks `json:"checks"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *As of May 1, 2021, card fingerprint in India for Connect changed to allow two fingerprints for the same card---one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding PaymentAttemptRecordPaymentMethodDetailsCardFunding `json:"funding"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// True if this payment was marked as MOTO and out of scope for SCA.
	MOTO bool `json:"moto"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `interac`, `jcb`, `link`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network PaymentAttemptRecordPaymentMethodDetailsCardNetwork `json:"network"`
	// If this card has network token credentials, this contains the details of the network token credentials.
	NetworkToken *PaymentAttemptRecordPaymentMethodDetailsCardNetworkToken `json:"network_token"`
	// This is used by the financial networks to identify a transaction. Visa calls this the Transaction ID, Mastercard calls this the Trace ID, and American Express calls this the Acquirer Reference Data. This value will be present if it is returned by the financial network in the authorization response, and null otherwise.
	NetworkTransactionID string `json:"network_transaction_id"`
	// Populated if this transaction used 3D Secure authentication.
	ThreeDSecure *PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
}

// Custom Payment Methods represent Payment Method types not modeled directly in
// the Stripe API. This resource consists of details about the custom payment method
// used for this payment attempt.
type PaymentAttemptRecordPaymentMethodDetailsCustom struct {
	// Display name for the custom (user-defined) payment method type used to make this payment.
	DisplayName string `json:"display_name"`
	// The custom payment method type associated with this payment.
	Type string `json:"type"`
}

// Information about the Payment Method debited for this payment.
type PaymentAttemptRecordPaymentMethodDetails struct {
	// The billing details associated with the method of payment.
	BillingDetails *PaymentAttemptRecordPaymentMethodDetailsBillingDetails `json:"billing_details"`
	// Details of the card used for this payment attempt.
	Card *PaymentAttemptRecordPaymentMethodDetailsCard `json:"card"`
	// Custom Payment Methods represent Payment Method types not modeled directly in
	// the Stripe API. This resource consists of details about the custom payment method
	// used for this payment attempt.
	Custom *PaymentAttemptRecordPaymentMethodDetailsCustom `json:"custom"`
	// ID of the Stripe PaymentMethod used to make this payment.
	PaymentMethod string `json:"payment_method"`
	// The type of Payment Method used for this payment attempt.
	Type PaymentAttemptRecordPaymentMethodDetailsType `json:"type"`
}

// Shipping information for this payment.
type PaymentAttemptRecordShippingDetails struct {
	// A representation of a physical address.
	Address *Address `json:"address"`
	// The shipping recipient's name.
	Name string `json:"name"`
	// The shipping recipient's phone number.
	Phone string `json:"phone"`
}

// A Payment Attempt Record represents an individual attempt at making a payment, on or off Stripe.
// Each payment attempt tries to collect a fixed amount of money from a fixed customer and payment
// method. Payment Attempt Records are attached to Payment Records. Only one attempt per Payment Record
// can have guaranteed funds.
type PaymentAttemptRecord struct {
	APIResource
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountCanceled *PaymentAttemptRecordAmountCanceled `json:"amount_canceled"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountFailed *PaymentAttemptRecordAmountFailed `json:"amount_failed"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountGuaranteed *PaymentAttemptRecordAmountGuaranteed `json:"amount_guaranteed"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountRequested *PaymentAttemptRecordAmountRequested `json:"amount_requested"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Customer information for this payment.
	CustomerDetails *PaymentAttemptRecordCustomerDetails `json:"customer_details"`
	// Indicates whether the customer was present in your checkout flow during this payment.
	CustomerPresence PaymentAttemptRecordCustomerPresence `json:"customer_presence"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Information about the Payment Method debited for this payment.
	PaymentMethodDetails *PaymentAttemptRecordPaymentMethodDetails `json:"payment_method_details"`
	// ID of the Payment Record this Payment Attempt Record belongs to.
	PaymentRecord string `json:"payment_record"`
	// An opaque string for manual reconciliation of this payment, for example a check number or a payment processor ID.
	PaymentReference string `json:"payment_reference"`
	// Shipping information for this payment.
	ShippingDetails *PaymentAttemptRecordShippingDetails `json:"shipping_details"`
}

// PaymentAttemptRecordList is a list of PaymentAttemptRecords as retrieved from a list endpoint.
type PaymentAttemptRecordList struct {
	APIResource
	ListMeta
	Data []*PaymentAttemptRecord `json:"data"`
}
