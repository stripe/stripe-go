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

// The type of the microdeposit sent to the customer. Used to distinguish between different verification methods.
type SetupIntentNextActionVerifyWithMicrodepositsMicrodepositType string

// List of values that SetupIntentNextActionVerifyWithMicrodepositsMicrodepositType can take
const (
	SetupIntentNextActionVerifyWithMicrodepositsMicrodepositTypeAmounts        SetupIntentNextActionVerifyWithMicrodepositsMicrodepositType = "amounts"
	SetupIntentNextActionVerifyWithMicrodepositsMicrodepositTypeDescriptorCode SetupIntentNextActionVerifyWithMicrodepositsMicrodepositType = "descriptor_code"
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

// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
type SetupIntentPaymentMethodOptionsCardMandateOptionsAmountType string

// List of values that SetupIntentPaymentMethodOptionsCardMandateOptionsAmountType can take
const (
	SetupIntentPaymentMethodOptionsCardMandateOptionsAmountTypeFixed   SetupIntentPaymentMethodOptionsCardMandateOptionsAmountType = "fixed"
	SetupIntentPaymentMethodOptionsCardMandateOptionsAmountTypeMaximum SetupIntentPaymentMethodOptionsCardMandateOptionsAmountType = "maximum"
)

// Specifies payment frequency. One of `day`, `week`, `month`, `year`, or `sporadic`.
type SetupIntentPaymentMethodOptionsCardMandateOptionsInterval string

// List of values that SetupIntentPaymentMethodOptionsCardMandateOptionsInterval can take
const (
	SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalDay      SetupIntentPaymentMethodOptionsCardMandateOptionsInterval = "day"
	SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalMonth    SetupIntentPaymentMethodOptionsCardMandateOptionsInterval = "month"
	SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalSporadic SetupIntentPaymentMethodOptionsCardMandateOptionsInterval = "sporadic"
	SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalWeek     SetupIntentPaymentMethodOptionsCardMandateOptionsInterval = "week"
	SetupIntentPaymentMethodOptionsCardMandateOptionsIntervalYear     SetupIntentPaymentMethodOptionsCardMandateOptionsInterval = "year"
)

// Specifies the type of mandates supported. Possible values are `india`.
type SetupIntentPaymentMethodOptionsCardMandateOptionsSupportedType string

// List of values that SetupIntentPaymentMethodOptionsCardMandateOptionsSupportedType can take
const (
	SetupIntentPaymentMethodOptionsCardMandateOptionsSupportedTypeIndia SetupIntentPaymentMethodOptionsCardMandateOptionsSupportedType = "india"
)

// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Permitted values include: `automatic` or `any`. If not provided, defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
type SetupIntentPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that SetupIntentPaymentMethodOptionsCardRequestThreeDSecure can take
const (
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAny           SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "any"
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureAutomatic     SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
	SetupIntentPaymentMethodOptionsCardRequestThreeDSecureChallengeOnly SetupIntentPaymentMethodOptionsCardRequestThreeDSecure = "challenge_only"
)

// Bank account verification method.
type SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod string

// List of values that SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod can take
const (
	SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethodAutomatic     SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod = "automatic"
	SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethodInstant       SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod = "instant"
	SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethodMicrodeposits SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod = "microdeposits"
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
	// The IP address from which the Mandate was accepted by the customer.
	IPAddress *string `form:"ip_address"`
	// The user agent of the browser from which the Mandate was accepted by the customer.
	UserAgent *string `form:"user_agent"`
}

// This hash contains details about the customer acceptance of the Mandate.
type SetupIntentMandateDataCustomerAcceptanceParams struct {
	// The time at which the customer accepted the Mandate.
	AcceptedAt int64 `form:"accepted_at"`
	// If this is a Mandate accepted offline, this hash contains details about the offline acceptance.
	Offline *SetupIntentMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	// If this is a Mandate accepted online, this hash contains details about the online acceptance.
	Online *SetupIntentMandateDataCustomerAcceptanceOnlineParams `form:"online"`
	// The type of customer acceptance information included with the Mandate.
	Type MandateCustomerAcceptanceType `form:"type"`
}

// This hash contains details about the Mandate to create. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
type SetupIntentMandateDataParams struct {
	// This hash contains details about the customer acceptance of the Mandate.
	CustomerAcceptance *SetupIntentMandateDataCustomerAcceptanceParams `form:"customer_acceptance"`
}

// Additional fields for Mandate creation
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// A URL for custom mandate text to render during confirmation step.
	// The URL will be rendered with additional GET parameters `payment_intent` and `payment_intent_client_secret` when confirming a Payment Intent,
	// or `setup_intent` and `setup_intent_client_secret` when confirming a Setup Intent.
	CustomMandateURL *string `form:"custom_mandate_url"`
	// List of Stripe products where this mandate can be selected automatically.
	DefaultFor []*string `form:"default_for"`
	// Description of the mandate interval. Only required if 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription *string `form:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule *string `form:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type"`
}

// If this is a `acss_debit` SetupIntent, this sub-hash contains details about the ACSS Debit payment method options.
type SetupIntentPaymentMethodOptionsACSSDebitParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options"`
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type SetupIntentPaymentMethodOptionsCardMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount"`
	// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
	AmountType *string `form:"amount_type"`
	// Currency in which future payments will be charged. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// A description of the mandate or subscription that is meant to be displayed to the customer.
	Description *string `form:"description"`
	// End date of the mandate or subscription. If not provided, the mandate will be active until canceled. If provided, end date should be after start date.
	EndDate *int64 `form:"end_date"`
	// Specifies payment frequency. One of `day`, `week`, `month`, `year`, or `sporadic`.
	Interval *string `form:"interval"`
	// The number of intervals between payments. For example, `interval=month` and `interval_count=3` indicates one payment every three months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks). This parameter is optional when `interval=sporadic`.
	IntervalCount *int64 `form:"interval_count"`
	// Unique identifier for the mandate or subscription.
	Reference *string `form:"reference"`
	// Start date of the mandate or subscription. Start date should not be lesser than yesterday.
	StartDate *int64 `form:"start_date"`
	// Specifies the type of mandates supported. Possible values are `india`.
	SupportedTypes []*string `form:"supported_types"`
}

// Configuration for any card setup attempted on this SetupIntent.
type SetupIntentPaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *SetupIntentPaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options"`
	// When specified, this parameter signals that a card has been collected
	// as MOTO (Mail Order Telephone Order) and thus out of scope for SCA. This
	// parameter can only be provided during confirmation.
	MOTO *bool `form:"moto"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Permitted values include: `automatic` or `any`. If not provided, defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure"`
}

// Additional fields for Mandate creation
type SetupIntentPaymentMethodOptionsSepaDebitMandateOptionsParams struct{}

// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
type SetupIntentPaymentMethodOptionsSepaDebitParams struct {
	// Additional fields for Mandate creation
	MandateOptions *SetupIntentPaymentMethodOptionsSepaDebitMandateOptionsParams `form:"mandate_options"`
}

// If this is a `us_bank_account` SetupIntent, this sub-hash contains details about the US bank account payment method options.
type SetupIntentPaymentMethodOptionsUSBankAccountParams struct {
	// Verification method for the intent
	VerificationMethod *string `form:"verification_method"`
}

// Payment-method-specific configuration for this SetupIntent.
type SetupIntentPaymentMethodOptionsParams struct {
	// If this is a `acss_debit` SetupIntent, this sub-hash contains details about the ACSS Debit payment method options.
	ACSSDebit *SetupIntentPaymentMethodOptionsACSSDebitParams `form:"acss_debit"`
	// Configuration for any card setup attempted on this SetupIntent.
	Card *SetupIntentPaymentMethodOptionsCardParams `form:"card"`
	// If this is a `sepa_debit` SetupIntent, this sub-hash contains details about the SEPA Debit payment method options.
	SepaDebit *SetupIntentPaymentMethodOptionsSepaDebitParams `form:"sepa_debit"`
	// If this is a `us_bank_account` SetupIntent, this sub-hash contains details about the US bank account payment method options.
	USBankAccount *SetupIntentPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account"`
}

// If this hash is populated, this SetupIntent will generate a single_use Mandate on success.
type SetupIntentSingleUseParams struct {
	// Amount the customer is granting permission to collect later. A positive integer representing how much to charge in the [smallest currency unit](https://stripe.com/docs/currencies#zero-decimal) (e.g., 100 cents to charge $1.00 or 100 to charge ¥100, a zero-decimal currency). The minimum amount is $0.50 US or [equivalent in charge currency](https://stripe.com/docs/currencies#minimum-and-maximum-charge-amounts). The amount value supports up to eight digits (e.g., a value of 99999999 for a USD charge of $999,999.99).
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
}

// Creates a SetupIntent object.
//
// After the SetupIntent is created, attach a payment method and [confirm](https://stripe.com/docs/api/setup_intents/confirm)
// to collect any required permissions to charge the payment method later.
type SetupIntentParams struct {
	Params `form:"*"`
	// The client secret of the SetupIntent. Required if a publishable key is used to retrieve the SetupIntent.
	ClientSecret *string `form:"client_secret"`
	// Set to `true` to attempt to confirm this SetupIntent immediately. This parameter defaults to `false`. If the payment method attached is a card, a return_url may be provided in case additional authentication is required.
	Confirm *bool `form:"confirm"`
	// ID of the Customer this SetupIntent belongs to, if one exists.
	//
	// If present, the SetupIntent's payment method will be attached to the Customer on successful setup. Payment methods attached to other Customers cannot be used with this SetupIntent.
	Customer *string `form:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// This hash contains details about the Mandate to create. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
	MandateData *SetupIntentMandateDataParams `form:"mandate_data"`
	// The Stripe account ID for which this SetupIntent is created.
	OnBehalfOf *string `form:"on_behalf_of"`
	// ID of the payment method (a PaymentMethod, Card, or saved Source object) to attach to this SetupIntent.
	PaymentMethod *string `form:"payment_method"`
	// Payment-method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	// The list of payment method types (e.g. card) that this SetupIntent is allowed to set up. If this is not provided, defaults to ["card"].
	PaymentMethodTypes []*string `form:"payment_method_types"`
	// The URL to redirect your customer back to after they authenticate or cancel their payment on the payment method's app or site. If you'd prefer to redirect to a mobile application, you can alternatively supply an application URI scheme. This parameter can only be used with [`confirm=true`](https://stripe.com/docs/api/setup_intents/create#create_setup_intent-confirm).
	ReturnURL *string `form:"return_url"`
	// If this hash is populated, this SetupIntent will generate a single_use Mandate on success.
	SingleUse *SetupIntentSingleUseParams `form:"single_use"`
	// Indicates how the payment method is intended to be used in the future. If not provided, this value defaults to `off_session`.
	Usage *string `form:"usage"`
}

// Returns a list of SetupIntents.
type SetupIntentListParams struct {
	ListParams `form:"*"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	Created *int64 `form:"created"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return SetupIntents for the customer specified by this customer ID.
	Customer *string `form:"customer"`
	// Only return SetupIntents associated with the specified payment method.
	PaymentMethod *string `form:"payment_method"`
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
	Params `form:"*"`
	// This hash contains details about the Mandate to create
	MandateData *SetupIntentMandateDataParams `form:"mandate_data"`
	// ID of the payment method (a PaymentMethod, Card, or saved Source object) to attach to this SetupIntent.
	PaymentMethod *string `form:"payment_method"`
	// Payment-method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	// The URL to redirect your customer back to after they authenticate on the payment method's app or site.
	// If you'd prefer to redirect to a mobile application, you can alternatively supply an application URI scheme.
	// This parameter is only used for cards and other redirect-based payment methods.
	ReturnURL *string `form:"return_url"`
}

// A SetupIntent object can be canceled when it is in one of these statuses: requires_payment_method, requires_confirmation, or requires_action.
//
// Once canceled, setup is abandoned and any operations on the SetupIntent will fail with an error.
type SetupIntentCancelParams struct {
	Params `form:"*"`
	// Reason for canceling this SetupIntent. Possible values are `abandoned`, `requested_by_customer`, or `duplicate`
	CancellationReason *string `form:"cancellation_reason"`
}

// Verifies microdeposits on a SetupIntent object.
type SetupIntentVerifyMicrodepositsParams struct {
	Params `form:"*"`
	// Two positive integers, in *cents*, equal to the values of the microdeposits sent to the bank account.
	Amounts []*int64 `form:"amounts"`
	// A six-character code starting with SM present in the microdeposit sent to the bank account.
	DescriptorCode *string `form:"descriptor_code"`
}
type SetupIntentNextActionRedirectToURL struct {
	// If the customer does not exit their browser while authenticating, they will be redirected to this specified URL after completion.
	ReturnURL string `json:"return_url"`
	// The URL you must redirect your customer to in order to authenticate.
	URL string `json:"url"`
}

// When confirming a SetupIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js.
type SetupIntentNextActionUseStripeSDK struct{}
type SetupIntentNextActionVerifyWithMicrodeposits struct {
	// The timestamp when the microdeposits are expected to land.
	ArrivalDate int64 `json:"arrival_date"`
	// The URL for the hosted verification page, which allows customers to verify their bank account.
	HostedVerificationURL string `json:"hosted_verification_url"`
	// The type of the microdeposit sent to the customer. Used to distinguish between different verification methods.
	MicrodepositType SetupIntentNextActionVerifyWithMicrodepositsMicrodepositType `json:"microdeposit_type"`
}

// If present, this property tells you what actions you need to take in order for your customer to continue payment setup.
type SetupIntentNextAction struct {
	RedirectToURL *SetupIntentNextActionRedirectToURL `json:"redirect_to_url"`
	// Type of the next action to perform, one of `redirect_to_url`, `use_stripe_sdk`, `alipay_handle_redirect`, `oxxo_display_details`, or `verify_with_microdeposits`.
	Type SetupIntentNextActionType `json:"type"`
	// When confirming a SetupIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js.
	UseStripeSDK            *SetupIntentNextActionUseStripeSDK            `json:"use_stripe_sdk"`
	VerifyWithMicrodeposits *SetupIntentNextActionVerifyWithMicrodeposits `json:"verify_with_microdeposits"`
}
type SetupIntentPaymentMethodOptionsACSSDebitMandateOptions struct {
	// A URL for custom mandate text
	CustomMandateURL string `json:"custom_mandate_url"`
	// List of Stripe products where this mandate can be selected automatically.
	DefaultFor []SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsDefaultFor `json:"default_for"`
	// Description of the interval. Only required if the 'payment_schedule' parameter is 'interval' or 'combined'.
	IntervalDescription string `json:"interval_description"`
	// Payment schedule for the mandate.
	PaymentSchedule SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsPaymentSchedule `json:"payment_schedule"`
	// Transaction type of the mandate.
	TransactionType SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsTransactionType `json:"transaction_type"`
}
type SetupIntentPaymentMethodOptionsACSSDebit struct {
	// See SetupIntentPaymentMethodOptionsACSSDebitCurrency for allowed values
	Currency       string                                                  `json:"currency"`
	MandateOptions *SetupIntentPaymentMethodOptionsACSSDebitMandateOptions `json:"mandate_options"`
	// Bank account verification method.
	VerificationMethod SetupIntentPaymentMethodOptionsACSSDebitVerificationMethod `json:"verification_method"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type SetupIntentPaymentMethodOptionsCardMandateOptions struct {
	// Amount to be charged for future payments.
	Amount int64 `json:"amount"`
	// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
	AmountType SetupIntentPaymentMethodOptionsCardMandateOptionsAmountType `json:"amount_type"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A description of the mandate or subscription that is meant to be displayed to the customer.
	Description string `json:"description"`
	// End date of the mandate or subscription. If not provided, the mandate will be active until canceled. If provided, end date should be after start date.
	EndDate int64 `json:"end_date"`
	// Specifies payment frequency. One of `day`, `week`, `month`, `year`, or `sporadic`.
	Interval SetupIntentPaymentMethodOptionsCardMandateOptionsInterval `json:"interval"`
	// The number of intervals between payments. For example, `interval=month` and `interval_count=3` indicates one payment every three months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks). This parameter is optional when `interval=sporadic`.
	IntervalCount int64 `json:"interval_count"`
	// Unique identifier for the mandate or subscription.
	Reference string `json:"reference"`
	// Start date of the mandate or subscription. Start date should not be lesser than yesterday.
	StartDate int64 `json:"start_date"`
	// Specifies the type of mandates supported. Possible values are `india`.
	SupportedTypes []SetupIntentPaymentMethodOptionsCardMandateOptionsSupportedType `json:"supported_types"`
}
type SetupIntentPaymentMethodOptionsCard struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *SetupIntentPaymentMethodOptionsCardMandateOptions `json:"mandate_options"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Permitted values include: `automatic` or `any`. If not provided, defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure SetupIntentPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}
type SetupIntentPaymentMethodOptionsSepaDebitMandateOptions struct{}
type SetupIntentPaymentMethodOptionsSepaDebit struct {
	MandateOptions *SetupIntentPaymentMethodOptionsSepaDebitMandateOptions `json:"mandate_options"`
}
type SetupIntentPaymentMethodOptionsUSBankAccount struct {
	// Bank account verification method.
	VerificationMethod SetupIntentPaymentMethodOptionsUSBankAccountVerificationMethod `json:"verification_method"`
}

// Payment-method-specific configuration for this SetupIntent.
type SetupIntentPaymentMethodOptions struct {
	ACSSDebit     *SetupIntentPaymentMethodOptionsACSSDebit     `json:"acss_debit"`
	Card          *SetupIntentPaymentMethodOptionsCard          `json:"card"`
	SepaDebit     *SetupIntentPaymentMethodOptionsSepaDebit     `json:"sepa_debit"`
	USBankAccount *SetupIntentPaymentMethodOptionsUSBankAccount `json:"us_bank_account"`
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
	// ID of the Connect application that created the SetupIntent.
	Application *Application `json:"application"`
	// Reason for cancellation of this SetupIntent, one of `abandoned`, `requested_by_customer`, or `duplicate`.
	CancellationReason SetupIntentCancellationReason `json:"cancellation_reason"`
	// The client secret of this SetupIntent. Used for client-side retrieval using a publishable key.
	//
	// The client secret can be used to complete payment setup from your frontend. It should not be stored, logged, or exposed to anyone other than the customer. Make sure that you have TLS enabled on any page that includes the client secret.
	ClientSecret string `json:"client_secret"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// ID of the Customer this SetupIntent belongs to, if one exists.
	//
	// If present, the SetupIntent's payment method will be attached to the Customer on successful setup. Payment methods attached to other Customers cannot be used with this SetupIntent.
	Customer *Customer `json:"customer"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The error encountered in the previous SetupIntent confirmation.
	LastSetupError *Error `json:"last_setup_error"`
	// The most recent SetupAttempt for this SetupIntent.
	LatestAttempt *SetupAttempt `json:"latest_attempt"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// ID of the multi use Mandate generated by the SetupIntent.
	Mandate *Mandate `json:"mandate"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// If present, this property tells you what actions you need to take in order for your customer to continue payment setup.
	NextAction *SetupIntentNextAction `json:"next_action"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The account (if any) for which the setup is intended.
	OnBehalfOf *Account `json:"on_behalf_of"`
	// ID of the payment method used with this SetupIntent.
	PaymentMethod *PaymentMethod `json:"payment_method"`
	// Payment-method-specific configuration for this SetupIntent.
	PaymentMethodOptions *SetupIntentPaymentMethodOptions `json:"payment_method_options"`
	// The list of payment method types (e.g. card) that this SetupIntent is allowed to set up.
	PaymentMethodTypes []string `json:"payment_method_types"`
	// ID of the single_use Mandate generated by the SetupIntent.
	SingleUseMandate *Mandate `json:"single_use_mandate"`
	// [Status](https://stripe.com/docs/payments/intents#intent-statuses) of this SetupIntent, one of `requires_payment_method`, `requires_confirmation`, `requires_action`, `processing`, `canceled`, or `succeeded`.
	Status SetupIntentStatus `json:"status"`
	// Indicates how the payment method is intended to be used in the future.
	//
	// Use `on_session` if you intend to only reuse the payment method when the customer is in your checkout flow. Use `off_session` if your customer may or may not be in your checkout flow. If not provided, this value defaults to `off_session`.
	Usage SetupIntentUsage `json:"usage"`
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
