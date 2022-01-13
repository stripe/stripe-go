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
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
}
type MandateCustomerAcceptance struct {
	AcceptedAt int64                             `json:"accepted_at"`
	Offline    *MandateCustomerAcceptanceOffline `json:"offline"`
	Online     *MandateCustomerAcceptanceOnline  `json:"online"`
	Type       MandateCustomerAcceptanceType     `json:"type"`
}
type MandateMultiUse struct{}
type MandatePaymentMethodDetailsACSSDebit struct {
	DefaultFor          []MandatePaymentMethodDetailsACSSDebitDefaultFor    `json:"default_for"`
	IntervalDescription string                                              `json:"interval_description"`
	PaymentSchedule     MandatePaymentMethodDetailsACSSDebitPaymentSchedule `json:"payment_schedule"`
	TransactionType     MandatePaymentMethodDetailsACSSDebitTransactionType `json:"transaction_type"`
}
type MandatePaymentMethodDetailsAUBECSDebit struct {
	URL string `json:"url"`
}
type MandatePaymentMethodDetailsBACSDebit struct {
	NetworkStatus MandatePaymentMethodDetailsBACSDebitNetworkStatus `json:"network_status"`
	Reference     string                                            `json:"reference"`
	URL           string                                            `json:"url"`
}
type MandatePaymentMethodDetailsCard struct{}
type MandatePaymentMethodDetailsSepaDebit struct {
	Reference string `json:"reference"`
	URL       string `json:"url"`
}
type MandatePaymentMethodDetails struct {
	ACSSDebit   *MandatePaymentMethodDetailsACSSDebit   `json:"acss_debit"`
	AUBECSDebit *MandatePaymentMethodDetailsAUBECSDebit `json:"au_becs_debit"`
	BACSDebit   *MandatePaymentMethodDetailsBACSDebit   `json:"bacs_debit"`
	Card        *MandatePaymentMethodDetailsCard        `json:"card"`
	SepaDebit   *MandatePaymentMethodDetailsSepaDebit   `json:"sepa_debit"`
	Type        PaymentMethodType                       `json:"type"`
}
type MandateSingleUse struct {
	Amount   int64    `json:"amount"`
	Currency Currency `json:"currency"`
}

// A Mandate is a record of the permission a customer has given you to debit their payment method.
type Mandate struct {
	APIResource
	CustomerAcceptance   *MandateCustomerAcceptance   `json:"customer_acceptance"`
	ID                   string                       `json:"id"`
	Livemode             bool                         `json:"livemode"`
	MultiUse             *MandateMultiUse             `json:"multi_use"`
	Object               string                       `json:"object"`
	PaymentMethod        *PaymentMethod               `json:"payment_method"`
	PaymentMethodDetails *MandatePaymentMethodDetails `json:"payment_method_details"`
	SingleUse            *MandateSingleUse            `json:"single_use"`
	Status               MandateStatus                `json:"status"`
	Type                 MandateType                  `json:"type"`
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
