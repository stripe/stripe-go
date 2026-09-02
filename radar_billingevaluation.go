//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Describes the presence of the customer during the payment.
type RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresence string

// List of values that RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresence can take
const (
	RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresenceOffSession RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresence = "off_session"
	RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresenceOnSession  RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresence = "on_session"
)

// Describes the type of payment.
type RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType string

// List of values that RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType can take
const (
	RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentTypeOneOff         RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType = "one_off"
	RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentTypeRecurring      RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType = "recurring"
	RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentTypeSetupOneOff    RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType = "setup_one_off"
	RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentTypeSetupRecurring RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType = "setup_recurring"
)

// Describes the type of money movement. Currently only `card` is supported.
type RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsMoneyMovementType string

// List of values that RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsMoneyMovementType can take
const (
	RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsMoneyMovementTypeCard RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsMoneyMovementType = "card"
)

// Risk level.
type RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevel string

// List of values that RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevel can take
const (
	RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevelElevated    RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevel = "elevated"
	RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevelHighest     RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevel = "highest"
	RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevelLow         RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevel = "low"
	RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevelNormal      RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevel = "normal"
	RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevelNotAssessed RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevel = "not_assessed"
	RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevelUnknown     RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevel = "unknown"
)

// Details about the client device to associate with the billing evaluation.
type RadarBillingEvaluationClientDeviceMetadataDetailsParams struct {
	// ID for the Radar Session to associate with the billing evaluation. A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions about the customer behind the upcoming payment.
	RadarSession *string `form:"radar_session" json:"radar_session"`
}

// Attributes of the customer being evaluated. Supply these when the customer isn't represented by a Customer or an Account. If `customer` or `customer_account` is also supplied, the attributes on that object are used and these are ignored.
type RadarBillingEvaluationCustomerDetailsDataParams struct {
	// The email address of the customer being evaluated.
	Email *string `form:"email" json:"email,omitempty"`
	// The full name or business name of the customer being evaluated.
	Name *string `form:"name" json:"name,omitempty"`
	// The phone number of the customer being evaluated.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Details about the customer whose upcoming payment is being evaluated.
type RadarBillingEvaluationCustomerDetailsParams struct {
	// The ID of the customer whose upcoming payment is being evaluated.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The ID of the Account representing the customer whose upcoming payment is being evaluated.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// Attributes of the customer being evaluated. Supply these when the customer isn't represented by a Customer or an Account. If `customer` or `customer_account` is also supplied, the attributes on that object are used and these are ignored.
	Data *RadarBillingEvaluationCustomerDetailsDataParams `form:"data" json:"data,omitempty"`
}

// Describes card money movement details.
type RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardParams struct {
	// Describes the presence of the customer during the payment.
	CustomerPresence *string `form:"customer_presence" json:"customer_presence,omitempty"`
	// Describes the type of payment.
	PaymentType *string `form:"payment_type" json:"payment_type,omitempty"`
}

// Details about how the money for the upcoming payment moves.
type RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsParams struct {
	// Describes card money movement details.
	Card *RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardParams `form:"card" json:"card,omitempty"`
	// Describes the type of money movement. Currently only `card` is supported.
	MoneyMovementType *string `form:"money_movement_type" json:"money_movement_type"`
}

// Billing information associated with the payment method used for the upcoming payment.
type RadarBillingEvaluationPaymentDetailsPaymentMethodDetailsBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Details about the payment method that the upcoming payment is charged to.
type RadarBillingEvaluationPaymentDetailsPaymentMethodDetailsParams struct {
	// Billing information associated with the payment method used for the upcoming payment.
	BillingDetails *RadarBillingEvaluationPaymentDetailsPaymentMethodDetailsBillingDetailsParams `form:"billing_details" json:"billing_details,omitempty"`
	// ID of the payment method that the upcoming payment is charged to.
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
}

// Shipping details for the goods or services covered by the upcoming payment.
type RadarBillingEvaluationPaymentDetailsShippingDetailsParams struct {
	// Shipping address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Shipping name.
	Name *string `form:"name" json:"name,omitempty"`
	// Shipping phone number.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Details about the upcoming payment being evaluated.
type RadarBillingEvaluationPaymentDetailsParams struct {
	// The amount that the upcoming payment collects. A positive integer representing how much is charged in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (for example, 100 cents to charge 1.00 USD or 100 to charge 100 Yen, a zero-decimal currency).
	Amount *int64 `form:"amount" json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
	// An arbitrary description of the upcoming payment.
	Description *string `form:"description" json:"description,omitempty"`
	// Details about how the money for the upcoming payment moves.
	MoneyMovementDetails *RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsParams `form:"money_movement_details" json:"money_movement_details,omitempty"`
	// Details about the payment method that the upcoming payment is charged to.
	PaymentMethodDetails *RadarBillingEvaluationPaymentDetailsPaymentMethodDetailsParams `form:"payment_method_details" json:"payment_method_details"`
	// Shipping details for the goods or services covered by the upcoming payment.
	ShippingDetails *RadarBillingEvaluationPaymentDetailsShippingDetailsParams `form:"shipping_details" json:"shipping_details,omitempty"`
	// The statement descriptor that appears on the customer's statement for the upcoming payment.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
}

// Request Stripe Radar's assessment of the non-payment abuse risk of an upcoming charge, before the payment is attempted.
type RadarBillingEvaluationParams struct {
	Params `form:"*"`
	// Details about the client device to associate with the billing evaluation.
	ClientDeviceMetadataDetails *RadarBillingEvaluationClientDeviceMetadataDetailsParams `form:"client_device_metadata_details" json:"client_device_metadata_details,omitempty"`
	// Details about the customer whose upcoming payment is being evaluated.
	CustomerDetails *RadarBillingEvaluationCustomerDetailsParams `form:"customer_details" json:"customer_details"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Details about the upcoming payment being evaluated.
	PaymentDetails *RadarBillingEvaluationPaymentDetailsParams `form:"payment_details" json:"payment_details"`
}

// AddExpand appends a new field to expand.
func (p *RadarBillingEvaluationParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *RadarBillingEvaluationParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details about the client device to associate with the billing evaluation.
type RadarBillingEvaluationCreateClientDeviceMetadataDetailsParams struct {
	// ID for the Radar Session to associate with the billing evaluation. A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions about the customer behind the upcoming payment.
	RadarSession *string `form:"radar_session" json:"radar_session"`
}

// Attributes of the customer being evaluated. Supply these when the customer isn't represented by a Customer or an Account. If `customer` or `customer_account` is also supplied, the attributes on that object are used and these are ignored.
type RadarBillingEvaluationCreateCustomerDetailsDataParams struct {
	// The email address of the customer being evaluated.
	Email *string `form:"email" json:"email,omitempty"`
	// The full name or business name of the customer being evaluated.
	Name *string `form:"name" json:"name,omitempty"`
	// The phone number of the customer being evaluated.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Details about the customer whose upcoming payment is being evaluated.
type RadarBillingEvaluationCreateCustomerDetailsParams struct {
	// The ID of the customer whose upcoming payment is being evaluated.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// The ID of the Account representing the customer whose upcoming payment is being evaluated.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// Attributes of the customer being evaluated. Supply these when the customer isn't represented by a Customer or an Account. If `customer` or `customer_account` is also supplied, the attributes on that object are used and these are ignored.
	Data *RadarBillingEvaluationCreateCustomerDetailsDataParams `form:"data" json:"data,omitempty"`
}

// Describes card money movement details.
type RadarBillingEvaluationCreatePaymentDetailsMoneyMovementDetailsCardParams struct {
	// Describes the presence of the customer during the payment.
	CustomerPresence *string `form:"customer_presence" json:"customer_presence,omitempty"`
	// Describes the type of payment.
	PaymentType *string `form:"payment_type" json:"payment_type,omitempty"`
}

// Details about how the money for the upcoming payment moves.
type RadarBillingEvaluationCreatePaymentDetailsMoneyMovementDetailsParams struct {
	// Describes card money movement details.
	Card *RadarBillingEvaluationCreatePaymentDetailsMoneyMovementDetailsCardParams `form:"card" json:"card,omitempty"`
	// Describes the type of money movement. Currently only `card` is supported.
	MoneyMovementType *string `form:"money_movement_type" json:"money_movement_type"`
}

// Billing information associated with the payment method used for the upcoming payment.
type RadarBillingEvaluationCreatePaymentDetailsPaymentMethodDetailsBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Details about the payment method that the upcoming payment is charged to.
type RadarBillingEvaluationCreatePaymentDetailsPaymentMethodDetailsParams struct {
	// Billing information associated with the payment method used for the upcoming payment.
	BillingDetails *RadarBillingEvaluationCreatePaymentDetailsPaymentMethodDetailsBillingDetailsParams `form:"billing_details" json:"billing_details,omitempty"`
	// ID of the payment method that the upcoming payment is charged to.
	PaymentMethod *string `form:"payment_method" json:"payment_method"`
}

// Shipping details for the goods or services covered by the upcoming payment.
type RadarBillingEvaluationCreatePaymentDetailsShippingDetailsParams struct {
	// Shipping address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Shipping name.
	Name *string `form:"name" json:"name,omitempty"`
	// Shipping phone number.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Details about the upcoming payment being evaluated.
type RadarBillingEvaluationCreatePaymentDetailsParams struct {
	// The amount that the upcoming payment collects. A positive integer representing how much is charged in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (for example, 100 cents to charge 1.00 USD or 100 to charge 100 Yen, a zero-decimal currency).
	Amount *int64 `form:"amount" json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency" json:"currency"`
	// An arbitrary description of the upcoming payment.
	Description *string `form:"description" json:"description,omitempty"`
	// Details about how the money for the upcoming payment moves.
	MoneyMovementDetails *RadarBillingEvaluationCreatePaymentDetailsMoneyMovementDetailsParams `form:"money_movement_details" json:"money_movement_details,omitempty"`
	// Details about the payment method that the upcoming payment is charged to.
	PaymentMethodDetails *RadarBillingEvaluationCreatePaymentDetailsPaymentMethodDetailsParams `form:"payment_method_details" json:"payment_method_details"`
	// Shipping details for the goods or services covered by the upcoming payment.
	ShippingDetails *RadarBillingEvaluationCreatePaymentDetailsShippingDetailsParams `form:"shipping_details" json:"shipping_details,omitempty"`
	// The statement descriptor that appears on the customer's statement for the upcoming payment.
	StatementDescriptor *string `form:"statement_descriptor" json:"statement_descriptor,omitempty"`
}

// Request Stripe Radar's assessment of the non-payment abuse risk of an upcoming charge, before the payment is attempted.
type RadarBillingEvaluationCreateParams struct {
	Params `form:"*"`
	// Details about the client device to associate with the billing evaluation.
	ClientDeviceMetadataDetails *RadarBillingEvaluationCreateClientDeviceMetadataDetailsParams `form:"client_device_metadata_details" json:"client_device_metadata_details,omitempty"`
	// Details about the customer whose upcoming payment is being evaluated.
	CustomerDetails *RadarBillingEvaluationCreateCustomerDetailsParams `form:"customer_details" json:"customer_details"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Details about the upcoming payment being evaluated.
	PaymentDetails *RadarBillingEvaluationCreatePaymentDetailsParams `form:"payment_details" json:"payment_details"`
}

// AddExpand appends a new field to expand.
func (p *RadarBillingEvaluationCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *RadarBillingEvaluationCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Client device metadata attached to this billing evaluation.
type RadarBillingEvaluationClientDeviceMetadataDetails struct {
	// ID for the Radar Session associated with the billing evaluation. A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	RadarSession string `json:"radar_session"`
}

// Attributes of the customer being evaluated. These are populated from the `customer` or `customer_account` object when one was supplied, and from the request otherwise.
type RadarBillingEvaluationCustomerDetailsData struct {
	// The customer's email address.
	Email string `json:"email"`
	// The customer's full name or business name.
	Name string `json:"name"`
	// The customer's phone number.
	Phone string `json:"phone"`
}

// Details of the customer this billing evaluation assesses.
type RadarBillingEvaluationCustomerDetails struct {
	// The ID of the customer whose upcoming payment was evaluated.
	Customer string `json:"customer"`
	// The ID of the Account representing the customer whose upcoming payment was evaluated.
	CustomerAccount string `json:"customer_account"`
	// Attributes of the customer being evaluated. These are populated from the `customer` or `customer_account` object when one was supplied, and from the request otherwise.
	Data *RadarBillingEvaluationCustomerDetailsData `json:"data"`
}

// Describes card money movement details.
type RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCard struct {
	// Describes the presence of the customer during the payment.
	CustomerPresence RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardCustomerPresence `json:"customer_presence"`
	// Describes the type of payment.
	PaymentType RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCardPaymentType `json:"payment_type"`
}

// Details about the payment's customer presence and type.
type RadarBillingEvaluationPaymentDetailsMoneyMovementDetails struct {
	// Describes card money movement details.
	Card *RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsCard `json:"card"`
	// Describes the type of money movement. Currently only `card` is supported.
	MoneyMovementType RadarBillingEvaluationPaymentDetailsMoneyMovementDetailsMoneyMovementType `json:"money_movement_type"`
}

// Billing information associated with the billing evaluation.
type RadarBillingEvaluationPaymentDetailsPaymentMethodDetailsBillingDetails struct {
	// Address data.
	Address *Address `json:"address"`
	// Email address.
	Email string `json:"email"`
	// Full name.
	Name string `json:"name"`
	// Billing phone number (including extension).
	Phone string `json:"phone"`
}

// Details about the payment method that will be charged.
type RadarBillingEvaluationPaymentDetailsPaymentMethodDetails struct {
	// Billing information associated with the billing evaluation.
	BillingDetails *RadarBillingEvaluationPaymentDetailsPaymentMethodDetailsBillingDetails `json:"billing_details"`
	// The payment method that will be charged.
	PaymentMethod string `json:"payment_method"`
}

// Shipping details for the billing evaluation.
type RadarBillingEvaluationPaymentDetailsShippingDetails struct {
	// Address data.
	Address *Address `json:"address"`
	// Shipping name.
	Name string `json:"name"`
	// Shipping phone number.
	Phone string `json:"phone"`
}

// Payment details for the upcoming charge this billing evaluation assesses.
type RadarBillingEvaluationPaymentDetails struct {
	// Amount intended to be collected by this payment. A positive integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://docs.stripe.com/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Details about the payment's customer presence and type.
	MoneyMovementDetails *RadarBillingEvaluationPaymentDetailsMoneyMovementDetails `json:"money_movement_details"`
	// Details about the payment method that will be charged.
	PaymentMethodDetails *RadarBillingEvaluationPaymentDetailsPaymentMethodDetails `json:"payment_method_details"`
	// Shipping details for the billing evaluation.
	ShippingDetails *RadarBillingEvaluationPaymentDetailsShippingDetails `json:"shipping_details"`
	// Payment statement descriptor.
	StatementDescriptor string `json:"statement_descriptor"`
}

// Stripe Radar's assessment of the likelihood that the upcoming charge results in non-payment abuse.
type RadarBillingEvaluationSignalsNonPaymentAbuse struct {
	// The time when this signal was evaluated.
	EvaluatedAt int64 `json:"evaluated_at"`
	// Risk level.
	RiskLevel RadarBillingEvaluationSignalsNonPaymentAbuseRiskLevel `json:"risk_level"`
}

// Stripe Radar's signals for the upcoming charge this billing evaluation assesses.
type RadarBillingEvaluationSignals struct {
	// Stripe Radar's assessment of the likelihood that the upcoming charge results in non-payment abuse.
	NonPaymentAbuse *RadarBillingEvaluationSignalsNonPaymentAbuse `json:"non_payment_abuse"`
}

// Billing Evaluations represent Stripe Radar's assessment of the non-payment abuse risk of an upcoming charge. Unlike a [Payment Evaluation](https://docs.stripe.com/api/radar/payment-evaluation), a billing evaluation is created before the payment is attempted and returns the `non_payment_abuse` signal only.
type RadarBillingEvaluation struct {
	APIResource
	// Client device metadata attached to this billing evaluation.
	ClientDeviceMetadataDetails *RadarBillingEvaluationClientDeviceMetadataDetails `json:"client_device_metadata_details,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	CreatedAt int64 `json:"created_at"`
	// Details of the customer this billing evaluation assesses.
	CustomerDetails *RadarBillingEvaluationCustomerDetails `json:"customer_details,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Payment details for the upcoming charge this billing evaluation assesses.
	PaymentDetails *RadarBillingEvaluationPaymentDetails `json:"payment_details,omitempty"`
	// Stripe Radar's signals for the upcoming charge this billing evaluation assesses.
	Signals *RadarBillingEvaluationSignals `json:"signals"`
}
