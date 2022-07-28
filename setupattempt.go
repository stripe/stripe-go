//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Indicates the directions of money movement for which this payment method is intended to be used.
//
// Include `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.
type SetupAttemptFlowDirection string

// List of values that SetupAttemptFlowDirection can take
const (
	SetupAttemptFlowDirectionInbound  SetupAttemptFlowDirection = "inbound"
	SetupAttemptFlowDirectionOutbound SetupAttemptFlowDirection = "outbound"
)

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
	SetupAttemptPaymentMethodDetailsCardThreeDSecureResultExempted            SetupAttemptPaymentMethodDetailsCardThreeDSecureResult = "exempted"
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
	ListParams `form:"*"`
	// A filter on the list, based on the object `created` field. The value
	// can be a string with an integer Unix timestamp, or it can be a
	// dictionary with a number of different query options.
	Created *int64 `form:"created"`
	// A filter on the list, based on the object `created` field. The value
	// can be a string with an integer Unix timestamp, or it can be a
	// dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return SetupAttempts created by the SetupIntent specified by
	// this ID.
	SetupIntent *string `form:"setup_intent"`
}
type SetupAttemptPaymentMethodDetailsACSSDebit struct{}
type SetupAttemptPaymentMethodDetailsAUBECSDebit struct{}
type SetupAttemptPaymentMethodDetailsBACSDebit struct{}
type SetupAttemptPaymentMethodDetailsBancontact struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	Bic string `json:"bic"`
	// The ID of the SEPA Direct Debit PaymentMethod which was generated by this SetupAttempt.
	GeneratedSepaDebit *PaymentMethod `json:"generated_sepa_debit"`
	// The mandate for the SEPA Direct Debit PaymentMethod which was generated by this SetupAttempt.
	GeneratedSepaDebitMandate *Mandate `json:"generated_sepa_debit_mandate"`
	// Last four characters of the IBAN.
	IbanLast4 string `json:"iban_last4"`
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	// Can be one of `en`, `de`, `fr`, or `nl`
	PreferredLanguage string `json:"preferred_language"`
	// Owner's verified full name. Values are verified or provided by Bancontact directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type SetupAttemptPaymentMethodDetailsBLIK struct{}
type SetupAttemptPaymentMethodDetailsBoleto struct{}

// Populated if this authorization used 3D Secure authentication.
type SetupAttemptPaymentMethodDetailsCardThreeDSecure struct {
	// For authenticated transactions: how the customer was authenticated by
	// the issuing bank.
	AuthenticationFlow SetupAttemptPaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	// Indicates the outcome of 3D Secure authentication.
	Result SetupAttemptPaymentMethodDetailsCardThreeDSecureResult `json:"result"`
	// Additional information about why 3D Secure succeeded or failed based
	// on the `result`.
	ResultReason SetupAttemptPaymentMethodDetailsCardThreeDSecureResultReason `json:"result_reason"`
	// The version of 3D Secure that was used.
	Version string `json:"version"`
}
type SetupAttemptPaymentMethodDetailsCard struct {
	// Populated if this authorization used 3D Secure authentication.
	ThreeDSecure *SetupAttemptPaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
}
type SetupAttemptPaymentMethodDetailsCardPresent struct {
	// The ID of the Card PaymentMethod which was generated by this SetupAttempt.
	GeneratedCard *PaymentMethod `json:"generated_card"`
}
type SetupAttemptPaymentMethodDetailsIdeal struct {
	// The customer's bank. Can be one of `abn_amro`, `asn_bank`, `bunq`, `handelsbanken`, `ing`, `knab`, `moneyou`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, or `van_lanschot`.
	Bank string `json:"bank"`
	// The Bank Identifier Code of the customer's bank.
	Bic string `json:"bic"`
	// The ID of the SEPA Direct Debit PaymentMethod which was generated by this SetupAttempt.
	GeneratedSepaDebit *PaymentMethod `json:"generated_sepa_debit"`
	// The mandate for the SEPA Direct Debit PaymentMethod which was generated by this SetupAttempt.
	GeneratedSepaDebitMandate *Mandate `json:"generated_sepa_debit_mandate"`
	// Last four characters of the IBAN.
	IbanLast4 string `json:"iban_last4"`
	// Owner's verified full name. Values are verified or provided by iDEAL directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type SetupAttemptPaymentMethodDetailsLink struct{}
type SetupAttemptPaymentMethodDetailsSepaDebit struct{}
type SetupAttemptPaymentMethodDetailsSofort struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Bank Identifier Code of the bank associated with the bank account.
	Bic string `json:"bic"`
	// The ID of the SEPA Direct Debit PaymentMethod which was generated by this SetupAttempt.
	GeneratedSepaDebit *PaymentMethod `json:"generated_sepa_debit"`
	// The mandate for the SEPA Direct Debit PaymentMethod which was generated by this SetupAttempt.
	GeneratedSepaDebitMandate *Mandate `json:"generated_sepa_debit_mandate"`
	// Last four characters of the IBAN.
	IbanLast4 string `json:"iban_last4"`
	// Preferred language of the Sofort authorization page that the customer is redirected to.
	// Can be one of `en`, `de`, `fr`, or `nl`
	PreferredLanguage string `json:"preferred_language"`
	// Owner's verified full name. Values are verified or provided by Sofort directly
	// (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
}
type SetupAttemptPaymentMethodDetailsUSBankAccount struct{}
type SetupAttemptPaymentMethodDetails struct {
	ACSSDebit   *SetupAttemptPaymentMethodDetailsACSSDebit   `json:"acss_debit"`
	AUBECSDebit *SetupAttemptPaymentMethodDetailsAUBECSDebit `json:"au_becs_debit"`
	BACSDebit   *SetupAttemptPaymentMethodDetailsBACSDebit   `json:"bacs_debit"`
	Bancontact  *SetupAttemptPaymentMethodDetailsBancontact  `json:"bancontact"`
	BLIK        *SetupAttemptPaymentMethodDetailsBLIK        `json:"blik"`
	Boleto      *SetupAttemptPaymentMethodDetailsBoleto      `json:"boleto"`
	Card        *SetupAttemptPaymentMethodDetailsCard        `json:"card"`
	CardPresent *SetupAttemptPaymentMethodDetailsCardPresent `json:"card_present"`
	Ideal       *SetupAttemptPaymentMethodDetailsIdeal       `json:"ideal"`
	Link        *SetupAttemptPaymentMethodDetailsLink        `json:"link"`
	SepaDebit   *SetupAttemptPaymentMethodDetailsSepaDebit   `json:"sepa_debit"`
	Sofort      *SetupAttemptPaymentMethodDetailsSofort      `json:"sofort"`
	// The type of the payment method used in the SetupIntent (e.g., `card`). An additional hash is included on `payment_method_details` with a name matching this value. It contains confirmation-specific information for the payment method.
	Type          SetupAttemptPaymentMethodDetailsType           `json:"type"`
	USBankAccount *SetupAttemptPaymentMethodDetailsUSBankAccount `json:"us_bank_account"`
}

// A SetupAttempt describes one attempted confirmation of a SetupIntent,
// whether that confirmation was successful or unsuccessful. You can use
// SetupAttempts to inspect details of a specific attempt at setting up a
// payment method using a SetupIntent.
type SetupAttempt struct {
	APIResource
	// The value of [application](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-application) on the SetupIntent at the time of this confirmation.
	Application *Application `json:"application"`
	// If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.
	//
	// It can only be used for this Stripe Account's own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.
	AttachToSelf bool `json:"attach_to_self"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The value of [customer](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-customer) on the SetupIntent at the time of this confirmation.
	Customer *Customer `json:"customer"`
	// Indicates the directions of money movement for which this payment method is intended to be used.
	//
	// Include `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.
	FlowDirections []SetupAttemptFlowDirection `json:"flow_directions"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The value of [on_behalf_of](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-on_behalf_of) on the SetupIntent at the time of this confirmation.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// ID of the payment method used with this SetupAttempt.
	PaymentMethod        *PaymentMethod                    `json:"payment_method"`
	PaymentMethodDetails *SetupAttemptPaymentMethodDetails `json:"payment_method_details"`
	// The error encountered during this attempt to confirm the SetupIntent, if any.
	SetupError *Error `json:"setup_error"`
	// ID of the SetupIntent that this attempt belongs to.
	SetupIntent *SetupIntent `json:"setup_intent"`
	// Status of this SetupAttempt, one of `requires_confirmation`, `requires_action`, `processing`, `succeeded`, `failed`, or `abandoned`.
	Status SetupAttemptStatus `json:"status"`
	// The value of [usage](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-usage) on the SetupIntent at the time of this confirmation, one of `off_session` or `on_session`.
	Usage SetupAttemptUsage `json:"usage"`
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
