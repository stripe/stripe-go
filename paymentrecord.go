//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Indicates whether the customer was present in your checkout flow during this payment.
type PaymentRecordCustomerPresence string

// List of values that PaymentRecordCustomerPresence can take
const (
	PaymentRecordCustomerPresenceOffSession PaymentRecordCustomerPresence = "off_session"
	PaymentRecordCustomerPresenceOnSession  PaymentRecordCustomerPresence = "on_session"
)

// The type of Payment Method used for this payment attempt.
type PaymentRecordPaymentMethodDetailsType string

// List of values that PaymentRecordPaymentMethodDetailsType can take
const (
	PaymentRecordPaymentMethodDetailsTypeCustom PaymentRecordPaymentMethodDetailsType = "custom"
)

// Retrieves a Payment Record with the given ID
type PaymentRecordParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Information about the payment attempt failure.
type PaymentRecordReportPaymentAttemptFailedParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// When the reported payment failed. Measured in seconds since the Unix epoch.
	FailedAt *int64            `form:"failed_at"`
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentAttemptFailedParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentAttemptFailedParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Information about the payment attempt guarantee.
type PaymentRecordReportPaymentAttemptGuaranteedParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// When the reported payment was guaranteed. Measured in seconds since the Unix epoch.
	GuaranteedAt *int64            `form:"guaranteed_at"`
	Metadata     map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentAttemptGuaranteedParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentAttemptGuaranteedParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The billing details associated with the method of payment.
type PaymentRecordReportPaymentAttemptPaymentMethodDetailsBillingDetailsParams struct {
	// The billing address associated with the method of payment.
	Address *AddressParams `form:"address"`
	// The billing email associated with the method of payment.
	Email *string `form:"email"`
	// The billing name associated with the method of payment.
	Name *string `form:"name"`
	// The billing phone number associated with the method of payment.
	Phone *string `form:"phone"`
}

// Information about the custom (user-defined) payment method used to make this payment.
type PaymentRecordReportPaymentAttemptPaymentMethodDetailsCustomParams struct {
	// Display name for the custom (user-defined) payment method type used to make this payment.
	DisplayName *string `form:"display_name"`
	// The custom payment method type associated with this payment.
	Type *string `form:"type"`
}

// Information about the Payment Method debited for this payment.
type PaymentRecordReportPaymentAttemptPaymentMethodDetailsParams struct {
	// The billing details associated with the method of payment.
	BillingDetails *PaymentRecordReportPaymentAttemptPaymentMethodDetailsBillingDetailsParams `form:"billing_details"`
	// Information about the custom (user-defined) payment method used to make this payment.
	Custom *PaymentRecordReportPaymentAttemptPaymentMethodDetailsCustomParams `form:"custom"`
	// ID of the Stripe Payment Method used to make this payment.
	PaymentMethod *string `form:"payment_method"`
	// The type of the payment method details. An additional hash is included on the payment_method_details with a name matching this value. It contains additional information specific to the type.
	Type *string `form:"type"`
}

// Shipping information for this payment.
type PaymentRecordReportPaymentAttemptShippingDetailsParams struct {
	// The physical shipping address.
	Address *AddressParams `form:"address"`
	// The shipping recipient's name.
	Name *string `form:"name"`
	// The shipping recipient's phone number.
	Phone *string `form:"phone"`
}

// Report a new payment attempt on the specified Payment Record. A new payment
//
//	attempt can only be specified if all other payment attempts are canceled or failed.
type PaymentRecordReportPaymentAttemptParams struct {
	Params `form:"*"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Information about the payment attempt failure.
	Failed *PaymentRecordReportPaymentAttemptFailedParams `form:"failed"`
	// Information about the payment attempt guarantee.
	Guaranteed *PaymentRecordReportPaymentAttemptGuaranteedParams `form:"guaranteed"`
	// When the reported payment was initiated. Measured in seconds since the Unix epoch.
	InitiatedAt *int64 `form:"initiated_at"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The outcome of the reported payment.
	Outcome *string `form:"outcome"`
	// Information about the Payment Method debited for this payment.
	PaymentMethodDetails *PaymentRecordReportPaymentAttemptPaymentMethodDetailsParams `form:"payment_method_details"`
	// Shipping information for this payment.
	ShippingDetails *PaymentRecordReportPaymentAttemptShippingDetailsParams `form:"shipping_details"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentAttemptParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentAttemptParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Report that the most recent payment attempt on the specified Payment Record
//
//	was canceled.
type PaymentRecordReportPaymentAttemptCanceledParams struct {
	Params `form:"*"`
	// When the reported payment was canceled. Measured in seconds since the Unix epoch.
	CanceledAt *int64 `form:"canceled_at"`
	// Specifies which fields in the response should be expanded.
	Expand   []*string         `form:"expand"`
	Metadata map[string]string `form:"metadata"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentAttemptCanceledParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentAttemptCanceledParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The amount you intend to collect for this payment.
type PaymentRecordReportPaymentAmountRequestedParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value *int64 `form:"value"`
}

// Customer information for this payment.
type PaymentRecordReportPaymentCustomerDetailsParams struct {
	// The customer who made the payment.
	Customer *string `form:"customer"`
	// The customer's phone number.
	Email *string `form:"email"`
	// The customer's name.
	Name *string `form:"name"`
	// The customer's phone number.
	Phone *string `form:"phone"`
}

// Information about the payment attempt failure.
type PaymentRecordReportPaymentFailedParams struct {
	// When the reported payment failed. Measured in seconds since the Unix epoch.
	FailedAt *int64 `form:"failed_at"`
}

// Information about the payment attempt guarantee.
type PaymentRecordReportPaymentGuaranteedParams struct {
	// When the reported payment was guaranteed. Measured in seconds since the Unix epoch.
	GuaranteedAt *int64 `form:"guaranteed_at"`
}

// The billing details associated with the method of payment.
type PaymentRecordReportPaymentPaymentMethodDetailsBillingDetailsParams struct {
	// The billing address associated with the method of payment.
	Address *AddressParams `form:"address"`
	// The billing email associated with the method of payment.
	Email *string `form:"email"`
	// The billing name associated with the method of payment.
	Name *string `form:"name"`
	// The billing phone number associated with the method of payment.
	Phone *string `form:"phone"`
}

// Information about the custom (user-defined) payment method used to make this payment.
type PaymentRecordReportPaymentPaymentMethodDetailsCustomParams struct {
	// Display name for the custom (user-defined) payment method type used to make this payment.
	DisplayName *string `form:"display_name"`
	// The custom payment method type associated with this payment.
	Type *string `form:"type"`
}

// Information about the Payment Method debited for this payment.
type PaymentRecordReportPaymentPaymentMethodDetailsParams struct {
	// The billing details associated with the method of payment.
	BillingDetails *PaymentRecordReportPaymentPaymentMethodDetailsBillingDetailsParams `form:"billing_details"`
	// Information about the custom (user-defined) payment method used to make this payment.
	Custom *PaymentRecordReportPaymentPaymentMethodDetailsCustomParams `form:"custom"`
	// ID of the Stripe Payment Method used to make this payment.
	PaymentMethod *string `form:"payment_method"`
	// The type of the payment method details. An additional hash is included on the payment_method_details with a name matching this value. It contains additional information specific to the type.
	Type *string `form:"type"`
}

// Shipping information for this payment.
type PaymentRecordReportPaymentShippingDetailsParams struct {
	// The physical shipping address.
	Address *AddressParams `form:"address"`
	// The shipping recipient's name.
	Name *string `form:"name"`
	// The shipping recipient's phone number.
	Phone *string `form:"phone"`
}

// Report a new Payment Record. You may report a Payment Record as it is
//
//	initialized and later report updates through the other report_* methods, or report Payment
//	Records in a terminal state directly, through this method.
type PaymentRecordReportPaymentParams struct {
	Params `form:"*"`
	// The amount you intend to collect for this payment.
	AmountRequested *PaymentRecordReportPaymentAmountRequestedParams `form:"amount_requested"`
	// Customer information for this payment.
	CustomerDetails *PaymentRecordReportPaymentCustomerDetailsParams `form:"customer_details"`
	// Indicates whether the customer was present in your checkout flow during this payment.
	CustomerPresence *string `form:"customer_presence"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Information about the payment attempt failure.
	Failed *PaymentRecordReportPaymentFailedParams `form:"failed"`
	// Information about the payment attempt guarantee.
	Guaranteed *PaymentRecordReportPaymentGuaranteedParams `form:"guaranteed"`
	// When the reported payment was initiated. Measured in seconds since the Unix epoch.
	InitiatedAt *int64 `form:"initiated_at"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The outcome of the reported payment.
	Outcome *string `form:"outcome"`
	// Information about the Payment Method debited for this payment.
	PaymentMethodDetails *PaymentRecordReportPaymentPaymentMethodDetailsParams `form:"payment_method_details"`
	// An opaque string for manual reconciliation of this payment, for example a check number or a payment processor ID.
	PaymentReference *string `form:"payment_reference"`
	// Shipping information for this payment.
	ShippingDetails *PaymentRecordReportPaymentShippingDetailsParams `form:"shipping_details"`
}

// AddExpand appends a new field to expand.
func (p *PaymentRecordReportPaymentParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *PaymentRecordReportPaymentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmountCanceled struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmountFailed struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmountGuaranteed struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// A representation of an amount of money, consisting of an amount and a currency.
type PaymentRecordAmountRequested struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// Customer information for this payment.
type PaymentRecordCustomerDetails struct {
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
type PaymentRecordPaymentMethodDetailsBillingDetails struct {
	// A representation of a physical address.
	Address *Address `json:"address"`
	// The billing email associated with the method of payment.
	Email string `json:"email"`
	// The billing name associated with the method of payment.
	Name string `json:"name"`
	// The billing phone number associated with the method of payment.
	Phone string `json:"phone"`
}

// Information about the custom (user-defined) payment method used to make this payment.
type PaymentRecordPaymentMethodDetailsCustom struct {
	// Display name for the custom (user-defined) payment method type used to make this payment.
	DisplayName string `json:"display_name"`
	// The custom payment method type associated with this payment.
	Type string `json:"type"`
}

// Information about the Payment Method debited for this payment.
type PaymentRecordPaymentMethodDetails struct {
	// The billing details associated with the method of payment.
	BillingDetails *PaymentRecordPaymentMethodDetailsBillingDetails `json:"billing_details"`
	// Information about the custom (user-defined) payment method used to make this payment.
	Custom *PaymentRecordPaymentMethodDetailsCustom `json:"custom"`
	// ID of the Stripe PaymentMethod used to make this payment.
	PaymentMethod string `json:"payment_method"`
	// The type of Payment Method used for this payment attempt.
	Type PaymentRecordPaymentMethodDetailsType `json:"type"`
}

// Shipping information for this payment.
type PaymentRecordShippingDetails struct {
	// A representation of a physical address.
	Address *Address `json:"address"`
	// The shipping recipient's name.
	Name string `json:"name"`
	// The shipping recipient's phone number.
	Phone string `json:"phone"`
}

// A Payment Record is a resource that allows you to represent payments that occur on- or off-Stripe.
// For example, you can create a Payment Record to model a payment made on a different payment processor,
// in order to mark an Invoice as paid and a Subscription as active. Payment Records consist of one or
// more Payment Attempt Records, which represent individual attempts made on a payment network.
type PaymentRecord struct {
	APIResource
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountCanceled *PaymentRecordAmountCanceled `json:"amount_canceled"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountFailed *PaymentRecordAmountFailed `json:"amount_failed"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountGuaranteed *PaymentRecordAmountGuaranteed `json:"amount_guaranteed"`
	// A representation of an amount of money, consisting of an amount and a currency.
	AmountRequested *PaymentRecordAmountRequested `json:"amount_requested"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Customer information for this payment.
	CustomerDetails *PaymentRecordCustomerDetails `json:"customer_details"`
	// Indicates whether the customer was present in your checkout flow during this payment.
	CustomerPresence PaymentRecordCustomerPresence `json:"customer_presence"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// ID of the latest Payment Attempt Record attached to this Payment Record.
	LatestPaymentAttemptRecord string `json:"latest_payment_attempt_record"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Information about the Payment Method debited for this payment.
	PaymentMethodDetails *PaymentRecordPaymentMethodDetails `json:"payment_method_details"`
	// An opaque string for manual reconciliation of this payment, for example a check number or a payment processor ID.
	PaymentReference string `json:"payment_reference"`
	// Shipping information for this payment.
	ShippingDetails *PaymentRecordShippingDetails `json:"shipping_details"`
}

// UnmarshalJSON handles deserialization of a PaymentRecord.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PaymentRecord) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type paymentRecord PaymentRecord
	var v paymentRecord
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PaymentRecord(v)
	return nil
}
