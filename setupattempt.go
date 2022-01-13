//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// For authenticated transactions: how the customer was authenticated by
// the issuing bank.
type SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take
const (
	SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// Indicates the outcome of 3D Secure authentication.
type SetupAttemptPaymentMethodDetailsCardThreeDSecureResult string

// List of values that SetupAttemptPaymentMethodDetailsCardThreeDSecureResult can take
const (
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultAuthenticated       SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultFailed              SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "failed"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultNotSupported        SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultProcessingError     SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// Additional information about why 3D Secure succeeded or failed based
// on the `result`.
type SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason can take
const (
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonBypassed            SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonCanceled            SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReasonRejected            SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

// The type of the payment method used in the SetupIntent (e.g., `card`). An additional hash is included on `payment_method_details` with a name matching this value. It contains confirmation-specific information for the payment method.
type SetupAttemptPaymentMethodDetailsType string

// List of values that SetupAttemptPaymentMethodDetailsType can take
const (
	SetupAttemptPaymentMethodDetailsTypeCard SetupAttemptPaymentMethodDetailsType = "card"
)

// Status of this SetupAttempt, one of `requires_confirmation`, `requires_action`, `processing`, `succeeded`, `failed`, or `abandoned`.
type SetupAttemptStatus string

// List of values that SetupAttemptStatus can take
const (
	SetupAttemptStatusAbandoned            SetupAttemptStatus = "abandoned"
	SetupAttemptStatusFailed               SetupAttemptStatus = "failed"
	SetupAttemptStatusProcessing           SetupAttemptStatus = "processing"
	SetupAttemptStatusRequiresAction       SetupAttemptStatus = "requires_action"
	SetupAttemptStatusRequiresConfirmation SetupAttemptStatus = "requires_confirmation"
	SetupAttemptStatusSucceeded            SetupAttemptStatus = "succeeded"
)

// The value of [usage](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-usage) on the SetupIntent at the time of this confirmation, one of `off_session` or `on_session`.
type SetupAttemptUsage string

// List of values that SetupAttemptUsage can take
const (
	SetupAttemptUsageOffSession SetupAttemptUsage = "off_session"
	SetupAttemptUsageOnSession  SetupAttemptUsage = "on_session"
)

// Returns a list of SetupAttempts associated with a provided SetupIntent.
type SetupAttemptListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	SetupIntent  *string           `form:"setup_intent"`
}
type SetupAttemptPaymentMethodDetailsACSSDebit struct{}
type SetupAttemptPaymentMethodDetailsAUBECSDebit struct{}
type SetupAttemptPaymentMethodDetailsBACSDebit struct{}
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
type SetupAttemptPaymentMethodDetailsBoleto struct{}

// Populated if this authorization used 3D Secure authentication.
type SetupAttemptPaymentMethodDetailsCardThreeDSecure struct {
	AuthenticationFlow SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	Result             SetupAttemptPaymentMethodDetailsCardThreeDSecureResult             `json:"result"`
	ResultReason       SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason       `json:"result_reason"`
	Version            string                                                             `json:"version"`
}
type SetupAttemptPaymentMethodDetailsCard struct {
	ThreeDSecure *SetupAttemptPaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
}
type SetupAttemptPaymentMethodDetailsCardPresent struct {
	GeneratedCard *PaymentMethod `json:"generated_card"`
}
type SetupAttemptPaymentMethodDetailsIdeal struct {
	Bank                      string         `json:"bank"`
	Bic                       string         `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string         `json:"iban_last4"`
	VerifiedName              string         `json:"verified_name"`
}
type SetupAttemptPaymentMethodDetailsSepaDebit struct{}
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
type SetupAttemptPaymentMethodDetails struct {
	ACSSDebit   *SetupAttemptPaymentMethodDetailsACSSDebit   `json:"acss_debit"`
	AUBECSDebit *SetupAttemptPaymentMethodDetailsAUBECSDebit `json:"au_becs_debit"`
	BACSDebit   *SetupAttemptPaymentMethodDetailsBACSDebit   `json:"bacs_debit"`
	Bancontact  *SetupAttemptPaymentMethodDetailsBancontact  `json:"bancontact"`
	Boleto      *SetupAttemptPaymentMethodDetailsBoleto      `json:"boleto"`
	Card        *SetupAttemptPaymentMethodDetailsCard        `json:"card"`
	CardPresent *SetupAttemptPaymentMethodDetailsCardPresent `json:"card_present"`
	Ideal       *SetupAttemptPaymentMethodDetailsIdeal       `json:"ideal"`
	SepaDebit   *SetupAttemptPaymentMethodDetailsSepaDebit   `json:"sepa_debit"`
	Sofort      *SetupAttemptPaymentMethodDetailsSofort      `json:"sofort"`
	Type        SetupAttemptPaymentMethodDetailsType         `json:"type"`
}

// A SetupAttempt describes one attempted confirmation of a SetupIntent,
// whether that confirmation was successful or unsuccessful. You can use
// SetupAttempts to inspect details of a specific attempt at setting up a
// payment method using a SetupIntent.
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
	SetupIntent          *SetupIntent                      `json:"setup_intent"`
	Status               SetupAttemptStatus                `json:"status"`
	Usage                SetupAttemptUsage                 `json:"usage"`
}

// SetupAttemptList is a list of SetupAttempts as retrieved from a list endpoint.
type SetupAttemptList struct {
	APIResource
	ListMeta
	Data []*SetupAttempt `json:"data"`
}

// UnmarshalJSON handles deserialization of a SetupAttempt.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SetupAttempt) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type setupAttempt SetupAttempt
	var v setupAttempt
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SetupAttempt(v)
	return nil
}
