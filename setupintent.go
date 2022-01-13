//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Reason for cancellation of this SetupIntent, one of `abandoned`, `requested_by_customer`, or `duplicate`.
type SetupIntentCancellationReason string

// List of values that SetupIntentCancellationReason can take
const (
	SetupIntentCancellationReasonAbandoned           SetupIntentCancellationReason = "abandoned"
	SetupIntentCancellationReasonDuplicate           SetupIntentCancellationReason = "duplicate"
	SetupIntentCancellationReasonFailedInvoice       SetupIntentCancellationReason = "failed_invoice"
	SetupIntentCancellationReasonFraudulent          SetupIntentCancellationReason = "fraudulent"
	SetupIntentCancellationReasonRequestedByCustomer SetupIntentCancellationReason = "requested_by_customer"
)

// Type of the next action to perform, one of `redirect_to_url`, `use_stripe_sdk`, `alipay_handle_redirect`, `oxxo_display_details`, or `verify_with_microdeposits`.
type SetupIntentNextActionType string

// List of values that SetupIntentNextActionType can take
const (
	SetupIntentNextActionTypeRedirectToURL SetupIntentNextActionType = "redirect_to_url"
)

// Currency supported by the bank account
type SetupIntentPaymentMethodOptionsACSSDebitCurrency string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitCurrency can take
const (
	SetupIntentPaymentMethodOptionsACSSDebitCurrencyCAD SetupIntentPaymentMethodOptionsACSSDebitCurrency = "cad"
	SetupIntentPaymentMethodOptionsACSSDebitCurrencyUSD SetupIntentPaymentMethodOptionsACSSDebitCurrency = "usd"
)

// List of Stripe products where this mandate can be selected automatically.
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor can take
const (
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultForInvoice      SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor = "invoice"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultForSubscription SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor = "subscription"
)

// Payment schedule for the mandate.
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule can take
const (
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleCombined SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "combined"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleInterval SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "interval"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentScheduleSporadic SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule = "sporadic"
)

// Transaction type of the mandate.
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType can take
const (
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypeBusiness SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "business"
	SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionTypePersonal SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType = "personal"
)

// Bank account verification method.
type SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod string

// List of values that SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod can take
const (
	SetupIntentPaymentMethodOptionsACSSDebitVerificationMethodAutomatic     SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod = "automatic"
	SetupIntentPaymentMethodOptionsACSSDebitVerificationMethodInstant       SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod = "instant"
	SetupIntentPaymentMethodOptionsACSSDebitVerificationMethodMicrodeposits SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod = "microdeposits"
)

// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Permitted values include: `automatic` or `any`. If not provided, defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type SetupIntentPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that SetupIntentPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAny           SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "any"
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAutomatic     SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureChallengeOnly SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "challenge_only"
)

// [Status](https://stripe.com/docs/payments/intents#intent-statuses) of this SetupIntent, one of `requires_payment_method`, `requires_confirmation`, `requires_action`, `processing`, `canceled`, or `succeeded`.
type SetupIntentStatus string

// List of values that SetupIntentStatus can take
const (
	SetupIntentStatusCanceled              SetupIntentStatus = "canceled"
	SetupIntentStatusProcessing            SetupIntentStatus = "processing"
	SetupIntentStatusRequiresAction        SetupIntentStatus = "requires_action"
	SetupIntentStatusRequiresConfirmation  SetupIntentStatus = "requires_confirmation"
	SetupIntentStatusRequiresPaymentMethod SetupIntentStatus = "requires_payment_method"
	SetupIntentStatusSucceeded             SetupIntentStatus = "succeeded"
)

// Indicates how the payment method is intended to be used in the future.
//
// Use `on_session` if you intend to only reuse the payment method when the customer is in your checkout flow. Use `off_session` if your customer may or may not be in your checkout flow. If not provided, this value defaults to `off_session`.
type SetupIntentUsage string

// List of values that SetupIntentUsage can take
const (
	SetupIntentUsageOffSession SetupIntentUsage = "off_session"
	SetupIntentUsageOnSession  SetupIntentUsage = "on_session"
)

// If this is a Mandate accepted offline, this hash contains details about the offline acceptance.
type SetupIntentMandateDataCustomerAcceptanceOfflineParams struct{}

// If this is a Mandate accepted online, this hash contains details about the online acceptance.
type SetupIntentMandateDataCustomerAcceptanceOnlineParams struct {
	IPAddress *string `form:"ip_address"`
	UserAgent *string `form:"user_agent"`
}

// This hash contains details about the customer acceptance of the Mandate.
type SetupIntentMandateDataCustomerAcceptanceParams struct {
	AcceptedAt int64                                                  `form:"accepted_at"`
	Offline    *SetupIntentMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	Online     *SetupIntentMandateDataCustomerAcceptanceOnlineParams  `form:"online"`
	Type       MandateCustomerAcceptanceType                          `form:"type"`
}

// This hash contains details about the Mandate to create. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
type SetupIntentMandateDataParams struct {
	CustomerAcceptance *SetupIntentMandateDataCustomerAcceptanceParams `form:"customer_acceptance"`
}

// Additional fields for Mandate creation
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	CustomMandateURL    *string   `form:"custom_mandate_url"`
	DefaultFor          []*string `form:"default_for"`
	IntervalDescription *string   `form:"interval_description"`
	PaymentSchedule     *string   `form:"payment_schedule"`
	TransactionType     *string   `form:"transaction_type"`
}

// If this is a `acss_debit` SetupIntent, this sub-hash contains details about the ACSS Debit payment method options.
type SetupIntentPaymentMethodOptionsACSSDebitParams struct {
	Currency           *string                                                       `form:"currency"`
	MandateOptions     *SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	VerificationMethod *string                                                       `form:"verification_method"`
}

// Configuration for any card setup attempted on this SetupIntent.
type SetupIntentPaymentMethodOptionsCardParams struct {
	MOTO                *bool   `form:"moto"`
	RequestThreeDSecure *string `form:"request_three_d_secure"`
}

// Additional fields for Mandate creation
type SetupIntentPaymentMethodOptionsSepaDebitMandateOptionsParams struct{}

// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
type SetupIntentPaymentMethodOptionsSepaDebitParams struct {
	MandateOptions *SetupIntentPaymentMethodOptionsSepaDebitMandateOptionsParams `form:"mandate_options"`
}

// Payment-method-specific configuration for this SetupIntent.
type SetupIntentPaymentMethodOptionsParams struct {
	ACSSDebit *SetupIntentPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	Card      *SetupIntentPaymentMethodOptionsCardParams      `form:"card"`
	SepaDebit *SetupIntentPaymentMethodOptionsSepaDebitParams `form:"sepa_debit"`
}

// If this hash is populated, this SetupIntent will generate a single_use Mandate on success.
type SetupIntentSingleUseParams struct {
	Amount   *int64  `form:"amount"`
	Currency *string `form:"currency"`
}

// Creates a SetupIntent object.
//
// After the SetupIntent is created, attach a payment method and [confirm](https://stripe.com/docs/api/setup_intents/confirm)
// to collect any required permissions to charge the payment method later.
type SetupIntentParams struct {
	Params               `form:"*"`
	ClientSecret         *string                                `form:"client_secret"`
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

// Returns a list of SetupIntents.
type SetupIntentListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Customer      *string           `form:"customer"`
	PaymentMethod *string           `form:"payment_method"`
}

// Confirm that your customer intends to set up the current or
// provided payment method. For example, you would confirm a SetupIntent
// when a customer hits the “Save” button on a payment method management
// page on your website.
//
// If the selected payment method does not require any additional
// steps from the customer, the SetupIntent will transition to the
// succeeded status.
//
// Otherwise, it will transition to the requires_action status and
// suggest additional actions via next_action. If setup fails,
// the SetupIntent will transition to the
// requires_payment_method status.
type SetupIntentConfirmParams struct {
	Params               `form:"*"`
	MandateData          *SetupIntentMandateDataParams          `form:"mandate_data"`
	PaymentMethod        *string                                `form:"payment_method"`
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	ReturnURL            *string                                `form:"return_url"`
}

// A SetupIntent object can be canceled when it is in one of these statuses: requires_payment_method, requires_confirmation, or requires_action.
//
// Once canceled, setup is abandoned and any operations on the SetupIntent will fail with an error.
type SetupIntentCancelParams struct {
	Params             `form:"*"`
	CancellationReason *string `form:"cancellation_reason"`
}
type SetupIntentNextActionRedirectToURL struct {
	ReturnURL string `json:"return_url"`
	URL       string `json:"url"`
}

// When confirming a SetupIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js.
type SetupIntentNextActionUseStripeSDK struct{}
type SetupIntentNextActionVerifyWithMicrodeposits struct {
	ArrivalDate           int64  `json:"arrival_date"`
	HostedVerificationURL string `json:"hosted_verification_url"`
}

// If present, this property tells you what actions you need to take in order for your customer to continue payment setup.
type SetupIntentNextAction struct {
	RedirectToURL           *SetupIntentNextActionRedirectToURL           `json:"redirect_to_url"`
	Type                    SetupIntentNextActionType                     `json:"type"`
	UseStripeSDK            *SetupIntentNextActionUseStripeSDK            `json:"use_stripe_sdk"`
	VerifyWithMicrodeposits *SetupIntentNextActionVerifyWithMicrodeposits `json:"verify_with_microdeposits"`
}
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptions struct {
	CustomMandateURL    string                                                                `json:"custom_mandate_url"`
	DefaultFor          []SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor    `json:"default_for"`
	IntervalDescription string                                                                `json:"interval_description"`
	PaymentSchedule     SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule `json:"payment_schedule"`
	TransactionType     SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}
type SetupIntentPaymentMethodOptionsACSSDebit struct {
	// See SetupIntentPaymentMethodOptionsACSSDebitCurrency for allowed values
	Currency           string                                                     `json:"currency"`
	MandateOptions     *SetupIntentPaymentMethodOptionsACSSDebitMandateOptions    `json:"mandate_options"`
	VerificationMethod SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}
type SetupIntentPaymentMethodOptionsCard struct {
	RequestThreeDSecure SetupIntentPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}
type SetupIntentPaymentMethodOptionsSepaDebitMandateOptions struct{}
type SetupIntentPaymentMethodOptionsSepaDebit struct {
	MandateOptions *SetupIntentPaymentMethodOptionsSepaDebitMandateOptions `json:"mandate_options"`
}

// Payment-method-specific configuration for this SetupIntent.
type SetupIntentPaymentMethodOptions struct {
	ACSSDebit *SetupIntentPaymentMethodOptionsACSSDebit `json:"acss_debit"`
	Card      *SetupIntentPaymentMethodOptionsCard      `json:"card"`
	SepaDebit *SetupIntentPaymentMethodOptionsSepaDebit `json:"sepa_debit"`
}

// A SetupIntent guides you through the process of setting up and saving a customer's payment credentials for future payments.
// For example, you could use a SetupIntent to set up and save your customer's card without immediately collecting a payment.
// Later, you can use [PaymentIntents](https://stripe.com/docs/api#payment_intents) to drive the payment flow.
//
// Create a SetupIntent as soon as you're ready to collect your customer's payment credentials.
// Do not maintain long-lived, unconfirmed SetupIntents as they may no longer be valid.
// The SetupIntent then transitions through multiple [statuses](https://stripe.com/docs/payments/intents#intent-statuses) as it guides
// you through the setup process.
//
// Successful SetupIntents result in payment credentials that are optimized for future payments.
// For example, cardholders in [certain regions](https://stripe.com/guides/strong-customer-authentication) may need to be run through
// [Strong Customer Authentication](https://stripe.com/docs/strong-customer-authentication) at the time of payment method collection
// in order to streamline later [off-session payments](https://stripe.com/docs/payments/setup-intents).
// If the SetupIntent is used with a [Customer](https://stripe.com/docs/api#setup_intent_object-customer), upon success,
// it will automatically attach the resulting payment method to that Customer.
// We recommend using SetupIntents or [setup_future_usage](https://stripe.com/docs/api#payment_intent_object-setup_future_usage) on
// PaymentIntents to save payment methods in order to prevent saving invalid or unoptimized payment methods.
//
// By using SetupIntents, you ensure that your customers experience the minimum set of required friction,
// even as regulations change over time.
//
// Related guide: [Setup Intents API](https://stripe.com/docs/payments/setup-intents).
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

// SetupIntentList is a list of SetupIntents as retrieved from a list endpoint.
type SetupIntentList struct {
	APIResource
	ListMeta
	Data []*SetupIntent `json:"data"`
}

// UnmarshalJSON handles deserialization of a SetupIntent.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SetupIntent) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type setupIntent SetupIntent
	var v setupIntent
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SetupIntent(v)
	return nil
}
