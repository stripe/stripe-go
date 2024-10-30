//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Whether the customer was present during the transaction.
type PaymentRecordCustomerPresence string

// List of values that PaymentRecordCustomerPresence can take
const (
	PaymentRecordCustomerPresenceOffSession PaymentRecordCustomerPresence = "off_session"
	PaymentRecordCustomerPresenceOnSession  PaymentRecordCustomerPresence = "on_session"
)

// The type of payment method used for this payment attempt.
type PaymentRecordPaymentMethodDetailsType string

// List of values that PaymentRecordPaymentMethodDetailsType can take
const (
	PaymentRecordPaymentMethodDetailsTypeCustom PaymentRecordPaymentMethodDetailsType = "custom"
)

// Amount object
type PaymentRecordAmountCanceled struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// Amount object
type PaymentRecordAmountFailed struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// Amount object
type PaymentRecordAmountGuaranteed struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// Amount object
type PaymentRecordAmountRefunded struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// Amount object
type PaymentRecordAmountRequested struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) e.g., 100 cents for $1.00 or 100 for ¥100, a zero-decimal currency).
	Value int64 `json:"value"`
}

// Customer information for this payment.
type PaymentRecordCustomerDetails struct {
	// ID of the Stripe customer for this payment.
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
	// Address data
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
	// The Custom Payment Method Type associated with this payment.
	Type string `json:"type"`
}

// Information about the method used to make this payment.
type PaymentRecordPaymentMethodDetails struct {
	// The billing details associated with the method of payment.
	BillingDetails *PaymentRecordPaymentMethodDetailsBillingDetails `json:"billing_details"`
	// Information about the custom (user-defined) payment method used to make this payment.
	Custom *PaymentRecordPaymentMethodDetailsCustom `json:"custom"`
	// ID of the Stripe PaymentMethod used to make this payment.
	PaymentMethod string `json:"payment_method"`
	// The type of payment method used for this payment attempt.
	Type PaymentRecordPaymentMethodDetailsType `json:"type"`
}

// Shipping information for this payment.
type PaymentRecordShippingDetails struct {
	// Address data
	Address *Address `json:"address"`
	// The recipient's name.
	Name string `json:"name"`
	// The recipient's phone number.
	Phone string `json:"phone"`
}

// A PaymentRecord represents a payment that happened on or off Stripe.
type PaymentRecord struct {
	// Amount object
	AmountCanceled *PaymentRecordAmountCanceled `json:"amount_canceled"`
	// Amount object
	AmountFailed *PaymentRecordAmountFailed `json:"amount_failed"`
	// Amount object
	AmountGuaranteed *PaymentRecordAmountGuaranteed `json:"amount_guaranteed"`
	// Amount object
	AmountRefunded *PaymentRecordAmountRefunded `json:"amount_refunded"`
	// Amount object
	AmountRequested *PaymentRecordAmountRequested `json:"amount_requested"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Customer information for this payment.
	CustomerDetails *PaymentRecordCustomerDetails `json:"customer_details"`
	// Whether the customer was present during the transaction.
	CustomerPresence PaymentRecordCustomerPresence `json:"customer_presence"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// ID of the latest PaymentAttemptRecord attached to this PaymentRecord.
	LatestPaymentAttemptRecord string `json:"latest_payment_attempt_record"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Information about the method used to make this payment.
	PaymentMethodDetails *PaymentRecordPaymentMethodDetails `json:"payment_method_details"`
	// An opaque string for manual reconciliation of this payment, for example a check number.
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
