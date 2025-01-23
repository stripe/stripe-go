//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

// Indicates whether the customer was present in your checkout flow during this payment.
type PaymentAttemptRecordCustomerPresence string

// List of values that PaymentAttemptRecordCustomerPresence can take
const (
	PaymentAttemptRecordCustomerPresenceOffSession PaymentAttemptRecordCustomerPresence = "off_session"
	PaymentAttemptRecordCustomerPresenceOnSession  PaymentAttemptRecordCustomerPresence = "on_session"
)

// The type of Payment Method used for this payment attempt.
type PaymentAttemptRecordPaymentMethodDetailsType string

// List of values that PaymentAttemptRecordPaymentMethodDetailsType can take
const (
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

// Information about the custom (user-defined) payment method used to make this payment.
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
	// Information about the custom (user-defined) payment method used to make this payment.
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
	Created time.Time `json:"created"`
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

// UnmarshalJSON handles deserialization of a PaymentAttemptRecord.
// This custom unmarshaling is needed to handle the time fields correctly.
func (p *PaymentAttemptRecord) UnmarshalJSON(data []byte) error {
	type paymentAttemptRecord PaymentAttemptRecord
	v := struct {
		Created int64 `json:"created"`
		*paymentAttemptRecord
	}{
		paymentAttemptRecord: (*paymentAttemptRecord)(p),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	p.Created = time.Unix(v.Created, 0)
	return nil
}

// MarshalJSON handles serialization of a PaymentAttemptRecord.
// This custom marshaling is needed to handle the time fields correctly.
func (p PaymentAttemptRecord) MarshalJSON() ([]byte, error) {
	type paymentAttemptRecord PaymentAttemptRecord
	v := struct {
		Created int64 `json:"created"`
		paymentAttemptRecord
	}{
		paymentAttemptRecord: (paymentAttemptRecord)(p),
		Created:              p.Created.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
