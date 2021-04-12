package stripe

import "encoding/json"

// List of values that MandateStatus can take.
const (
	MandateCustomerAcceptanceTypeOffline MandateCustomerAcceptanceType = "offline"
	MandateCustomerAcceptanceTypeOnline  MandateCustomerAcceptanceType = "online"
)

// MandateCustomerAcceptanceType is the list of allowed values for the type of customer acceptance
// for a given mandate.
type MandateCustomerAcceptanceType string

// List of values that MandateStatus can take.
const (
	MandateStatusActive   MandateStatus = "active"
	MandateStatusInactive MandateStatus = "inactive"
	MandateStatusPending  MandateStatus = "pending"
)

// MandatePaymentMethodDetailsACSSDebitPaymentSchedule is the list of allowed values for an acss debit payment_schedule on payment_method_details
type MandatePaymentMethodDetailsACSSDebitPaymentSchedule string

// List of values that MandatePaymentMethodDetailsACSSDebitPaymentSchedule can take
const (
	MandatePaymentMethodDetailsACSSDebitPaymentScheduleCombined MandatePaymentMethodDetailsACSSDebitPaymentSchedule = "combined"
	MandatePaymentMethodDetailsACSSDebitPaymentScheduleInterval MandatePaymentMethodDetailsACSSDebitPaymentSchedule = "interval"
	MandatePaymentMethodDetailsACSSDebitPaymentScheduleSporadic MandatePaymentMethodDetailsACSSDebitPaymentSchedule = "sporadic"
)

// MandatePaymentMethodDetailsACSSDebitTransactionType is the list of allowed values for an acss debit transaction type
type MandatePaymentMethodDetailsACSSDebitTransactionType string

// List of values that MandatePaymentMethodDetailsACSSDebitTransactionType can take
const (
	MandatePaymentMethodDetailsACSSDebitTransactionTypeBusiness MandatePaymentMethodDetailsACSSDebitTransactionType = "business"
	MandatePaymentMethodDetailsACSSDebitTransactionTypePersonal MandatePaymentMethodDetailsACSSDebitTransactionType = "personal"
)

// MandatePaymentMethodDetailsBACSDebitNetworkStatus is the list of allowed values for the status
// with the network for a given mandate.
type MandatePaymentMethodDetailsBACSDebitNetworkStatus string

// List of values that MandateStatus can take.
const (
	MandatePaymentMethodDetailsBACSDebitNetworkStatusAccepted MandatePaymentMethodDetailsBACSDebitNetworkStatus = "accepted"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusPending  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "pending"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusRefused  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "refused"
	MandatePaymentMethodDetailsBACSDebitNetworkStatusRevoked  MandatePaymentMethodDetailsBACSDebitNetworkStatus = "revoked"
)

// MandateStatus is the list of allowed values for the mandate status.
type MandateStatus string

// List of values that MandateType can take.
const (
	MandateTypeMultiUse  MandateType = "multi_use"
	MandateTypeSingleUse MandateType = "single_use"
)

// MandateType is the list of allowed values for the mandate type.
type MandateType string

// MandateParams is the set of parameters that can be used when retrieving a mandate.
type MandateParams struct {
	Params `form:"*"`
}

// MandateCustomerAcceptanceOffline represents details about the customer acceptance of an offline
// mandate.
type MandateCustomerAcceptanceOffline struct {
}

// MandateCustomerAcceptanceOnline represents details about the customer acceptance of an online
// mandate.
type MandateCustomerAcceptanceOnline struct {
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
}

// MandateCustomerAcceptance represents details about the customer acceptance for a mandate.
type MandateCustomerAcceptance struct {
	AcceptedAt int64                             `json:"accepted_at"`
	Offline    *MandateCustomerAcceptanceOffline `json:"offline"`
	Online     *MandateCustomerAcceptanceOnline  `json:"online"`
	Type       MandateCustomerAcceptanceType     `json:"type"`
}

// MandateMultiUse represents details about a multi-use mandate.
type MandateMultiUse struct {
}

// MandatePaymentMethodDetailsACSSDebit represent details about the acss debit associated with this mandate.
type MandatePaymentMethodDetailsACSSDebit struct {
	IntervalDescription string                                              `json:"interval_description"`
	PaymentSchedule     MandatePaymentMethodDetailsACSSDebitPaymentSchedule `json:"payment_schedule"`
	TransactionType     MandatePaymentMethodDetailsACSSDebitTransactionType `json:"transaction_type"`
}

// MandatePaymentMethodDetailsAUBECSDebit represents details about the Australia BECS debit account
// associated with this mandate.
type MandatePaymentMethodDetailsAUBECSDebit struct {
	URL string `json:"url"`
}

// MandatePaymentMethodDetailsBACSDebit represents details about the BACS debit account
// associated with this mandate.
type MandatePaymentMethodDetailsBACSDebit struct {
	NetworkStatus MandatePaymentMethodDetailsBACSDebitNetworkStatus `json:"network_status"`
	Reference     string                                            `json:"reference"`
	URL           string                                            `json:"url"`
}

// MandatePaymentMethodDetailsCard represents details about the card associated with this mandate.
type MandatePaymentMethodDetailsCard struct {
}

// MandatePaymentMethodDetailsSepaDebit represents details about the SEPA debit bank account
// associated with this mandate.
type MandatePaymentMethodDetailsSepaDebit struct {
	Reference string `json:"reference"`
	URL       string `json:"url"`
}

// MandatePaymentMethodDetails represents details about the payment method associated with this
// mandate.
type MandatePaymentMethodDetails struct {
	ACSSDebit   *MandatePaymentMethodDetailsACSSDebit   `json:"acss_debit"`
	AUBECSDebit *MandatePaymentMethodDetailsAUBECSDebit `json:"au_becs_debit"`
	BACSDebit   *MandatePaymentMethodDetailsBACSDebit   `json:"bacs_debit"`
	Card        *MandatePaymentMethodDetailsCard        `json:"card"`
	SepaDebit   *MandatePaymentMethodDetailsSepaDebit   `json:"sepa_debit"`
	Type        PaymentMethodType                       `json:"type"`
}

// MandateSingleUse represents details about a single-use mandate.
type MandateSingleUse struct {
	Amount   int64    `json:"amount"`
	Currency Currency `json:"currency"`
}

// Mandate is the resource representing a Mandate.
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
func (i *Mandate) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type ma Mandate
	var v ma
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = Mandate(v)
	return nil
}
