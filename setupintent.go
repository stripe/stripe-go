package stripe

import (
	"encoding/json"
)

// SetupIntentCancellationReason is the list of allowed values for the cancelation reason.
type SetupIntentCancellationReason string

// List of values that SetupIntentCancellationReason can take.
const (
	SetupIntentCancellationReasonAbandoned           SetupIntentCancellationReason = "abandoned"
	SetupIntentCancellationReasonFailedInvoice       SetupIntentCancellationReason = "failed_invoice"
	SetupIntentCancellationReasonFraudulent          SetupIntentCancellationReason = "fraudulent"
	SetupIntentCancellationReasonRequestedByCustomer SetupIntentCancellationReason = "requested_by_customer"
)

// SetupIntentNextActionType is the list of allowed values for the next action's type.
type SetupIntentNextActionType string

// List of values that SetupIntentNextActionType can take.
const (
	SetupIntentNextActionTypeRedirectToURL SetupIntentNextActionType = "redirect_to_url"
)

// SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule is the list of allowed values
// for payment_schedule on mandate_options.
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule can take.
const (
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleCombined SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "combined"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleInterval SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "interval"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleSporadic SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "sporadic"
)

// SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType is the list of allowed values for
// transaction_type on mandate_options.
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take.
const (
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod is the list of allowed values for
// verification_method on acss_debit.
type SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod can take.
const (
	SetupIntentPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	SetupIntentPaymentMethodOptionsACSSDebitVerificationMethodInstant       SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	SetupIntentPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// SetupIntentPaymentMethodOptionsCardRequestThreeDSecure is the list of allowed values controlling
// when to request 3D Secure on a SetupIntent.
type SetupIntentPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that SetupIntentNextActionType can take.
const (
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAny       SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "any"
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAutomatic SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
)

// SetupIntentStatus is the list of allowed values for the setup intent's status.
type SetupIntentStatus string

// List of values that SetupIntentStatus can take.
const (
	SetupIntentStatusCanceled              SetupIntentStatus = "canceled"
	SetupIntentStatusProcessing            SetupIntentStatus = "processing"
	SetupIntentStatusRequiresAction        SetupIntentStatus = "requires_action"
	SetupIntentStatusRequiresConfirmation  SetupIntentStatus = "requires_confirmation"
	SetupIntentStatusRequiresPaymentMethod SetupIntentStatus = "requires_payment_method"
	SetupIntentStatusSucceeded             SetupIntentStatus = "succeeded"
)

// SetupIntentUsage is the list of allowed values for the setup intent's usage.
type SetupIntentUsage string

// List of values that SetupIntentUsage can take.
const (
	SetupIntentUsageOffSession SetupIntentUsage = "off_session"
	SetupIntentUsageOnSession  SetupIntentUsage = "on_session"
)

// SetupIntentCancelParams is the set of parameters that can be used when canceling a setup intent.
type SetupIntentCancelParams struct {
	Params             `form:"*"`
	CancellationReason *string `form:"cancellation_reason"`
}

// SetupIntentConfirmParams is the set of parameters that can be used when confirming a setup intent.
type SetupIntentConfirmParams struct {
	Params               `form:"*"`
	MandateData          *SetupIntentMandateDataParams          `form:"mandate_data"`
	PaymentMethod        *string                                `form:"payment_method"`
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	ReturnURL            *string                                `form:"return_url"`
}

// SetupIntentMandateDataCustomerAcceptanceOfflineParams is the set of parameters for the customer
// acceptance of an offline mandate.
type SetupIntentMandateDataCustomerAcceptanceOfflineParams struct {
}

// SetupIntentMandateDataCustomerAcceptanceOnlineParams is the set of parameters for the customer
// acceptance of an online mandate.
type SetupIntentMandateDataCustomerAcceptanceOnlineParams struct {
	IPAddress *string `form:"ip_address"`
	UserAgent *string `form:"user_agent"`
}

// SetupIntentMandateDataCustomerAcceptanceParams is the set of parameters for the customer
// acceptance of a mandate.
type SetupIntentMandateDataCustomerAcceptanceParams struct {
	AcceptedAt int64                                                  `form:"accepted_at"`
	Offline    *SetupIntentMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	Online     *SetupIntentMandateDataCustomerAcceptanceOnlineParams  `form:"online"`
	Type       MandateCustomerAcceptanceType                          `form:"type"`
}

// SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsParams is the set of parameters for
// mandate_options on acss_debit.
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	CustomMandateURL    *string `form:"custom_mandate_url"`
	IntervalDescription *string `form:"interval_description"`
	PaymentSchedule     *string `form:"payment_schedule"`
	TransactionType     *string `form:"transaction_type"`
}

// SetupIntentPaymentMethodOptionsACSSDebitParams is the set of parameters for acss debit on
// payment_method_options.
type SetupIntentPaymentMethodOptionsACSSDebitParams struct {
	Currency           *string                                                       `form:"currency"`
	MandateOptions     *SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                       `form:"verification_method"`
}

// SetupIntentMandateDataParams is the set of parameters controlling the creation of the mandate
// associated with this SetupIntent.
type SetupIntentMandateDataParams struct {
	CustomerAcceptance *SetupIntentMandateDataCustomerAcceptanceParams `form:"customer_acceptance"`
}

// SetupIntentPaymentMethodOptionsCardParams represents the card-specific options applied to a
// SetupIntent.
type SetupIntentPaymentMethodOptionsCardParams struct {
	MOTO                *bool   `form:"moto"`
	RequestThreeDSecure *string `form:"request_three_d_secure"`
}

// SetupIntentPaymentMethodOptionsParams represents the type-specific payment method options
// applied to a SetupIntent.
type SetupIntentPaymentMethodOptionsParams struct {
	ACSSDebit *SetupIntentPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	Card      *SetupIntentPaymentMethodOptionsCardParams      `form:"card"`
}

// SetupIntentSingleUseParams represents the single-use mandate-specific parameters.
type SetupIntentSingleUseParams struct {
	Amount   *int64  `form:"amount"`
	Currency *string `form:"currency"`
}

// SetupIntentParams is the set of parameters that can be used when handling a setup intent.
type SetupIntentParams struct {
	Params               `form:"*"`
	Confirm              *bool                                  `form:"confirm"`
	Customer             *string                                `form:"customer"`
	Description          *string                                `form:"description"`
	MandateData          *SetupIntentMandateDataParams          `form:"mandate_data"`
	OnBehalfOf           *string                                `form:"on_behalf_of"`
	PaymentMethod        *string                                `form:"payment_method"`
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	PaymentMethodTypes   []*string                              `form:"payment_method_types"`
	ReturnURL            *string                                `form:"return_url"`
	SingleUse            *SetupIntentSingleUseParams            `form:"single_use"`
	Usage                *string                                `form:"usage"`
}

// SetupIntentListParams is the set of parameters that can be used when listing setup intents.
// For more details see https://stripe.com/docs/api#list_payouts.
type SetupIntentListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Customer      *string           `form:"customer"`
	PaymentMethod *string           `form:"payment_method"`
}

// SetupIntentNextActionRedirectToURL represents the resource for the next action of type
// "redirect_to_url".
type SetupIntentNextActionRedirectToURL struct {
	ReturnURL string `json:"return_url"`
	URL       string `json:"url"`
}

// SetupIntentNextActionUseStripeSDK represents the resource for the next action of type
// "use_stripe_sdk".
type SetupIntentNextActionUseStripeSDK struct{}

// SetupIntentNextActionVerifyWithMicrodeposits represents the resource for the next action
// of type "verify_with_microdeposits".
type SetupIntentNextActionVerifyWithMicrodeposits struct {
	ArrivalDate           int64  `json:"arrival_date"`
	HostedVerificationURL string `json:"hosted_verification_url"`
}

// SetupIntentNextAction represents the type of action to take on a setup intent.
type SetupIntentNextAction struct {
	RedirectToURL           *SetupIntentNextActionRedirectToURL           `json:"redirect_to_url"`
	Type                    SetupIntentNextActionType                     `json:"type"`
	UseStripeSDK            *SetupIntentNextActionUseStripeSDK            `json:"use_stripe_sdk"`
	VerifyWithMicrodeposits *SetupIntentNextActionVerifyWithMicrodeposits `json:"verify_with_microdeposits"`
}

// SetupIntentPaymentMethodOptionsACSSDebitMandateOptions describe the mandate options for
// acss debit associated with a setup intent.
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptions struct {
	CustomMandateURL    string                                                                `json:"custom_mandate_url"`
	IntervalDescription string                                                                `json:"interval_description"`
	PaymentSchedule     SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule `json:"payment_schedule"`
	TransactionType     SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}

// SetupIntentPaymentMethodOptionsACSSDebit represents the ACSS debit-specific options applied
// to a SetupIntent.
type SetupIntentPaymentMethodOptionsACSSDebit struct {
	Currency           string                                                     `json:"currency"`
	MandateOptions     *SetupIntentPaymentMethodOptionsACSSDebitMandateOptions    `json:"mandate_options"`
	VerificationMethod SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// SetupIntentPaymentMethodOptionsCard represents the card-specific options applied to a
// SetupIntent.
type SetupIntentPaymentMethodOptionsCard struct {
	RequestThreeDSecure SetupIntentPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}

// SetupIntentPaymentMethodOptions represents the type-specific payment method options applied to a
// SetupIntent.
type SetupIntentPaymentMethodOptions struct {
	ACSSDebit *SetupIntentPaymentMethodOptionsACSSDebit `json:"acss_debit"`
	Card      *SetupIntentPaymentMethodOptionsCard      `json:"card"`
}

// SetupIntent is the resource representing a Stripe payout.
// For more details see https://stripe.com/docs/api#payment_intents.
type SetupIntent struct {
	APIResource
	Application          *Application                     `json:"application"`
	CancellationReason   SetupIntentCancellationReason    `json:"cancellation_reason"`
	ClientSecret         string                           `json:"client_secret"`
	Created              int64                            `json:"created"`
	Customer             *Customer                        `json:"customer"`
	Description          string                           `json:"description"`
	ID                   string                           `json:"id"`
	LastSetupError       *Error                           `json:"last_setup_error"`
	LatestAttempt        *SetupAttempt                    `json:"latest_attempt"`
	Livemode             bool                             `json:"livemode"`
	Mandate              *Mandate                         `json:"mandate"`
	Metadata             map[string]string                `json:"metadata"`
	NextAction           *SetupIntentNextAction           `json:"next_action"`
	Object               string                           `json:"object"`
	OnBehalfOf           *Account                         `json:"on_behalf_of"`
	PaymentMethod        *PaymentMethod                   `json:"payment_method"`
	PaymentMethodOptions *SetupIntentPaymentMethodOptions `json:"payment_method_options"`
	PaymentMethodTypes   []string                         `json:"payment_method_types"`
	SingleUseMandate     *Mandate                         `json:"single_use_mandate"`
	Status               SetupIntentStatus                `json:"status"`
	Usage                SetupIntentUsage                 `json:"usage"`
}

// SetupIntentList is a list of setup intents as retrieved from a list endpoint.
type SetupIntentList struct {
	APIResource
	ListMeta
	Data []*SetupIntent `json:"data"`
}

// UnmarshalJSON handles deserialization of a SetupIntent.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *SetupIntent) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type setupIntent SetupIntent
	var v setupIntent
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = SetupIntent(v)
	return nil
}
