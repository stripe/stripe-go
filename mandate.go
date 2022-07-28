//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The type of customer acceptance information included with the Mandate. One of `online` or `offline`.
type MandateCustomerAcceptanceType string

// List of values that MandateCustomerAcceptanceType can take
const (
	MandateCustomerAcceptanceTypeOffline MandateCustomerAcceptanceType = "offline"
	MandateCustomerAcceptanceTypeOnline  MandateCustomerAcceptanceType = "online"
)

// List of Stripe products where this mandate can be selected automatically.
type MandatePaymentMethodDetailsACSSDebitDefaultFor string

// List of values that MandatePaymentMethodDetailsACSSDebitDefaultFor can take
const (
	MandatePaymentMethodDetailsACSSDebitDefaultForInvoice      MandatePaymentMethodDetailsACSSDebitDefaultFor = "invoice"
	MandatePaymentMethodDetailsACSSDebitDefaultForSubscription MandatePaymentMethodDetailsACSSDebitDefaultFor = "subscription"
)

// Payment schedule for the mandate.
type MandatePaymentMethodDetailsACSSDebitPaymentSchedule string

// List of values that MandatePaymentMethodDetailsACSSDebitPaymentSchedule can take
const (
	MandatePaymentMethodDetailsACSSDebitPaymentScheduleCombined MandatePaymentMethodDetailsACSSDebitPaymentSchedule = "combined"
	MandatePaymentMethodDetailsACSSDebitPaymentScheduleInterval MandatePaymentMethodDetailsACSSDebitPaymentSchedule = "interval"
	MandatePaymentMethodDetailsACSSDebitPaymentScheduleSporadic MandatePaymentMethodDetailsACSSDebitPaymentSchedule = "sporadic"
)

// Transaction type of the mandate.
type MandatePaymentMethodDetailsACSSDebitTransactionType string

// List of values that MandatePaymentMethodDetailsACSSDebitTransactionType can take
const (
	MandatePaymentMethodDetailsACSSDebitTransactionTypeBusiness MandatePaymentMethodDetailsACSSDebitTransactionType = "business"
	MandatePaymentMethodDetailsACSSDebitTransactionTypePersonal MandatePaymentMethodDetailsACSSDebitTransactionType = "personal"
)

// The status of the mandate on the Bacs network. Can be one of `pending`, `revoked`, `refused`, or `accepted`.
type MandatePaymentMethodDetailsBACSDebitNetworkStatus string

// List of values that MandatePaymentMethodDetailsBACSDebitNetworkStatus can take
const (
	MandatePaymentMethodDetailsBACSDebitNetworkStatusAccepted MandatePaymentMethodDetailsBACSDebitNetworkStatus = "accepted"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusPending  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "pending"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusRefused  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "refused"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusRevoked  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "revoked"
)

// Frequency interval of each recurring payment.
type MandatePaymentMethodDetailsBLIKOffSessionInterval string

// List of values that MandatePaymentMethodDetailsBLIKOffSessionInterval can take
const (
	MandatePaymentMethodDetailsBLIKOffSessionIntervalDay   MandatePaymentMethodDetailsBLIKOffSessionInterval = "day"
	MandatePaymentMethodDetailsBLIKOffSessionIntervalMonth MandatePaymentMethodDetailsBLIKOffSessionInterval = "month"
	MandatePaymentMethodDetailsBLIKOffSessionIntervalWeek  MandatePaymentMethodDetailsBLIKOffSessionInterval = "week"
	MandatePaymentMethodDetailsBLIKOffSessionIntervalYear  MandatePaymentMethodDetailsBLIKOffSessionInterval = "year"
)

// Type of the mandate.
type MandatePaymentMethodDetailsBLIKType string

// List of values that MandatePaymentMethodDetailsBLIKType can take
const (
	MandatePaymentMethodDetailsBLIKTypeOffSession MandatePaymentMethodDetailsBLIKType = "off_session"
	MandatePaymentMethodDetailsBLIKTypeOnSession  MandatePaymentMethodDetailsBLIKType = "on_session"
)

// The status of the mandate, which indicates whether it can be used to initiate a payment.
type MandateStatus string

// List of values that MandateStatus can take
const (
	MandateStatusActive   MandateStatus = "active"
	MandateStatusInactive MandateStatus = "inactive"
	MandateStatusPending  MandateStatus = "pending"
)

// The type of the mandate.
type MandateType string

// List of values that MandateType can take
const (
	MandateTypeMultiUse  MandateType = "multi_use"
	MandateTypeSingleUse MandateType = "single_use"
)

// Retrieves a Mandate object.
type MandateParams struct {
	Params `form:"*"`
}
type MandateCustomerAcceptanceOffline struct{}
type MandateCustomerAcceptanceOnline struct {
	// The IP address from which the Mandate was accepted by the customer.
	IPAddress string `json:"ip_address"`
	// The user agent of the browser from which the Mandate was accepted by the customer.
	UserAgent string `json:"user_agent"`
}
type MandateCustomerAcceptance struct {
	// The time at which the customer accepted the Mandate.
	AcceptedAt int64                             `json:"accepted_at"`
	Offline    *MandateCustomerAcceptanceOffline `json:"offline"`
	Online     *MandateCustomerAcceptanceOnline  `json:"online"`
	// The type of customer acceptance information included with the Mandate. One of `online` or `offline`.
	Type MandateCustomerAcceptanceType `json:"type"`
}
type MandateMultiUse struct{}
type MandatePaymentMethodDetailsACSSDebit struct {
	// List of Stripe products where this mandate can be selected automatically.
	DefaultFor []MandatePaymentMethodDetailsACSSDebitDefaultFor `json:"default_for"`
	// Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription string `json:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule MandatePaymentMethodDetailsACSSDebitPaymentSchedule `json:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType MandatePaymentMethodDetailsACSSDebitTransactionType `json:"transaction_type"`
}
type MandatePaymentMethodDetailsAUBECSDebit struct {
	// The URL of the mandate. This URL generally contains sensitive information about the customer and should be shared with them exclusively.
	URL string `json:"url"`
}
type MandatePaymentMethodDetailsBACSDebit struct {
	// The status of the mandate on the Bacs network. Can be one of `pending`, `revoked`, `refused`, or `accepted`.
	NetworkStatus MandatePaymentMethodDetailsBACSDebitNetworkStatus `json:"network_status"`
	// The unique reference identifying the mandate on the Bacs network.
	Reference string `json:"reference"`
	// The URL that will contain the mandate that the customer has signed.
	URL string `json:"url"`
}
type MandatePaymentMethodDetailsBLIKOffSession struct {
	// Amount of each recurring payment.
	Amount int64 `json:"amount"`
	// Currency of each recurring payment.
	Currency Currency `json:"currency"`
	// Frequency interval of each recurring payment.
	Interval MandatePaymentMethodDetailsBLIKOffSessionInterval `json:"interval"`
	// Frequency indicator of each recurring payment.
	IntervalCount int64 `json:"interval_count"`
}
type MandatePaymentMethodDetailsBLIK struct {
	// Date at which the mandate expires.
	ExpiresAfter int64                                      `json:"expires_after"`
	OffSession   *MandatePaymentMethodDetailsBLIKOffSession `json:"off_session"`
	// Type of the mandate.
	Type MandatePaymentMethodDetailsBLIKType `json:"type"`
}
type MandatePaymentMethodDetailsCard struct{}
type MandatePaymentMethodDetailsLink struct{}
type MandatePaymentMethodDetailsSepaDebit struct {
	// The unique reference of the mandate.
	Reference string `json:"reference"`
	// The URL of the mandate. This URL generally contains sensitive information about the customer and should be shared with them exclusively.
	URL string `json:"url"`
}
type MandatePaymentMethodDetailsUSBankAccount struct{}
type MandatePaymentMethodDetails struct {
	ACSSDebit   *MandatePaymentMethodDetailsACSSDebit   `json:"acss_debit"`
	AUBECSDebit *MandatePaymentMethodDetailsAUBECSDebit `json:"au_becs_debit"`
	BACSDebit   *MandatePaymentMethodDetailsBACSDebit   `json:"bacs_debit"`
	BLIK        *MandatePaymentMethodDetailsBLIK        `json:"blik"`
	Card        *MandatePaymentMethodDetailsCard        `json:"card"`
	Link        *MandatePaymentMethodDetailsLink        `json:"link"`
	SepaDebit   *MandatePaymentMethodDetailsSepaDebit   `json:"sepa_debit"`
	// The type of the payment method associated with this mandate. An additional hash is included on `payment_method_details` with a name matching this value. It contains mandate information specific to the payment method.
	Type          PaymentMethodType                         `json:"type"`
	USBankAccount *MandatePaymentMethodDetailsUSBankAccount `json:"us_bank_account"`
}
type MandateSingleUse struct {
	// On a single use mandate, the amount of the payment.
	Amount int64 `json:"amount"`
	// On a single use mandate, the currency of the payment.
	Currency Currency `json:"currency"`
}

// A Mandate is a record of the permission a customer has given you to debit their payment method.
type Mandate struct {
	APIResource
	CustomerAcceptance *MandateCustomerAcceptance `json:"customer_acceptance"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool             `json:"livemode"`
	MultiUse *MandateMultiUse `json:"multi_use"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// ID of the payment method associated with this mandate.
	PaymentMethod        *PaymentMethod               `json:"payment_method"`
	PaymentMethodDetails *MandatePaymentMethodDetails `json:"payment_method_details"`
	SingleUse            *MandateSingleUse            `json:"single_use"`
	// The status of the mandate, which indicates whether it can be used to initiate a payment.
	Status MandateStatus `json:"status"`
	// The type of the mandate.
	Type MandateType `json:"type"`
}

// UnmarshalJSON handles deserialization of a Mandate.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (m *Mandate) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		m.ID = id
		return nil
	}

	type mandate Mandate
	var v mandate
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*m = Mandate(v)
	return nil
}
