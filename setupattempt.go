package stripe

import (
	"encoding/json"
)

// SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow indicates the type of 3D Secure authentication performed.
type SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take.
const (
	SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// SetupAttemptPaymentMethodDetailsCardThreeDSecureResult indicates the outcome of 3D Secure authentication.
type SetupAttemptPaymentMethodDetailsCardThreeDSecureResult string

// List of values that SetupAttemptPaymentMethodDetailsCardThreeDSecureResult can take.
const (
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultAuthenticated       SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultFailed              SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "failed"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultNotSupported        SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultProcessingError     SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason represents dditional information about why 3D Secure succeeded or failed
type SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason can take.
const (
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonBypassed            SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonCanceled            SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonRejected            SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

// SetupAttemptPaymentMethodDetailsType is the type of the payment method associated with the setup attempt's payment method details.
type SetupAttemptPaymentMethodDetailsType string

// List of values that SetupAttemptPaymentMethodDetailsType can take.
const (
	SetupAttemptPaymentMethodDetailsTypeCard SetupAttemptPaymentMethodDetailsType = "card"
)

// SetupAttemptUsage is the list of allowed values for usage.
type SetupAttemptUsage string

// List of values that SetupAttemptUsage can take.
const (
	SetupAttemptUsageOffSession SetupAttemptUsage = "off_session"
	SetupAttemptUsageOnSession  SetupAttemptUsage = "on_session"
)

// SetupAttemptStatus is the list of allowed values for the setup attempt's status.
type SetupAttemptStatus string

// List of values that SetupAttemptStatus can take.
const (
	SetupAttemptStatusAbandoned            SetupAttemptStatus = "abandoned"
	SetupAttemptStatusFailed               SetupAttemptStatus = "failed"
	SetupAttemptStatusProcessing           SetupAttemptStatus = "processing"
	SetupAttemptStatusRequiresAction       SetupAttemptStatus = "requires_action"
	SetupAttemptStatusRequiresConfirmation SetupAttemptStatus = "requires_confirmation"
	SetupAttemptStatusSucceeded            SetupAttemptStatus = "succeeded"
)

// SetupAttemptListParams is the set of parameters that can be used when listing setup attempts.
type SetupAttemptListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	SetupIntent  *string           `form:"setup_intent"`
}

// SetupAttemptPaymentMethodDetailsCardThreeDSecure represents details about 3DS associated with the setup attempt's payment method.
type SetupAttemptPaymentMethodDetailsCardThreeDSecure struct {
	AuthenticationFlow SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	Result             SetupAttemptPaymentMethodDetailsCardThreeDSecureResult             `json:"result"`
	ResultReason       SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason       `json:"result_reason"`
	Version            string                                                             `json:"version"`
}

// SetupAttemptPaymentMethodDetailsBancontact represents details about the Bancontact PaymentMethod.
type SetupAttemptPaymentMethodDetailsBancontact struct {
	BankCode                  string         `json:"bank_code"`
	BankName                  string         `json:"bank_name"`
	Bic                       string         `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string         `json:"iban_last4"`
	PreferredLanguage         string         `json:"preferred_language"`
	VerifiedName              string         `json:"verified_name"`
}

// SetupAttemptPaymentMethodDetailsCard represents details about the Card PaymentMethod.
type SetupAttemptPaymentMethodDetailsCard struct {
	ThreeDSecure *SetupAttemptPaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
}

// SetupAttemptPaymentMethodDetailsIdeal represents details about the Bancontact PaymentMethod.
type SetupAttemptPaymentMethodDetailsIdeal struct {
	Bank                      string         `json:"bank"`
	Bic                       string         `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string         `json:"iban_last4"`
	VerifiedName              string         `json:"verified_name"`
}

// SetupAttemptPaymentMethodDetailsSofort represents details about the Bancontact PaymentMethod.
type SetupAttemptPaymentMethodDetailsSofort struct {
	BankCode                  string         `json:"bank_code"`
	BankName                  string         `json:"bank_name"`
	Bic                       string         `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string         `json:"iban_last4"`
	PreferredLanguage         string         `json:"preferred_language"`
	VerifiedName              string         `json:"verified_name"`
}

// SetupAttemptPaymentMethodDetails represents the details about the PaymentMethod associated with the setup attempt.
type SetupAttemptPaymentMethodDetails struct {
	Bancontact *SetupAttemptPaymentMethodDetailsBancontact `json:"bancontact"`
	Card       *SetupAttemptPaymentMethodDetailsCard       `json:"card"`
	Ideal      *SetupAttemptPaymentMethodDetailsIdeal      `json:"ideal"`
	Sofort     *SetupAttemptPaymentMethodDetailsSofort     `json:"sofort"`
	Type       SetupAttemptPaymentMethodDetailsType        `json:"type"`
}

// SetupAttempt is the resource representing a Stripe setup attempt.
type SetupAttempt struct {
	APIResource
	Application          *Application                      `json:"application"`
	Created              int64                             `json:"created"`
	Customer             *Customer                         `json:"customer"`
	ID                   string                            `json:"id"`
	Livemode             bool                              `json:"livemode"`
	Object               string                            `json:"object"`
	OnBehalfOf           *Account                          `json:"on_behalf_of"`
	PaymentMethod        *PaymentMethod                    `json:"payment_method"`
	PaymentMethodDetails *SetupAttemptPaymentMethodDetails `json:"payment_method_details"`
	SetupError           *Error                            `json:"setup_error"`
	Status               SetupAttemptStatus                `json:"status"`
	Usage                SetupAttemptUsage                 `json:"usage"`
}

// SetupAttemptList is a list of setup attempts as retrieved from a list endpoint.
type SetupAttemptList struct {
	APIResource
	ListMeta
	Data []*SetupAttempt `json:"data"`
}

// UnmarshalJSON handles deserialization of a SetupAttempt.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *SetupAttempt) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type setupAttempt SetupAttempt
	var v setupAttempt
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = SetupAttempt(v)
	return nil
}
