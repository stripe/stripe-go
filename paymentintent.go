package stripe

import (
	"encoding/json"
)

// PaymentIntentCancellationReason is the list of allowed values for the cancelation reason.
type PaymentIntentCancellationReason string

// List of values that PaymentIntentCancellationReason can take.
const (
	PaymentIntentCancellationReasonAbandoned           PaymentIntentCancellationReason = "abandoned"
	PaymentIntentCancellationReasonAutomatic           PaymentIntentCancellationReason = "automatic"
	PaymentIntentCancellationReasonDuplicate           PaymentIntentCancellationReason = "duplicate"
	PaymentIntentCancellationReasonFailedInvoice       PaymentIntentCancellationReason = "failed_invoice"
	PaymentIntentCancellationReasonFraudulent          PaymentIntentCancellationReason = "fraudulent"
	PaymentIntentCancellationReasonRequestedByCustomer PaymentIntentCancellationReason = "requested_by_customer"
	PaymentIntentCancellationReasonVoidInvoice         PaymentIntentCancellationReason = "void_invoice"
)

// PaymentIntentCaptureMethod is the list of allowed values for the capture method.
type PaymentIntentCaptureMethod string

// List of values that PaymentIntentCaptureMethod can take.
const (
	PaymentIntentCaptureMethodAutomatic PaymentIntentCaptureMethod = "automatic"
	PaymentIntentCaptureMethodManual    PaymentIntentCaptureMethod = "manual"
)

// PaymentIntentConfirmationMethod is the list of allowed values for the confirmation method.
type PaymentIntentConfirmationMethod string

// List of values that PaymentIntentConfirmationMethod can take.
const (
	PaymentIntentConfirmationMethodAutomatic PaymentIntentConfirmationMethod = "automatic"
	PaymentIntentConfirmationMethodManual    PaymentIntentConfirmationMethod = "manual"
)

// PaymentIntentNextActionType is the list of allowed values for the next action's type.
type PaymentIntentNextActionType string

// List of values that PaymentIntentNextActionType can take.
const (
	PaymentIntentNextActionTypeRedirectToURL PaymentIntentNextActionType = "redirect_to_url"
)

// PaymentIntentOffSession is the list of allowed values for types of off-session.
type PaymentIntentOffSession string

// List of values that PaymentIntentOffSession can take.
const (
	PaymentIntentOffSessionOneOff    PaymentIntentOffSession = "one_off"
	PaymentIntentOffSessionRecurring PaymentIntentOffSession = "recurring"
)

// PaymentIntentPaymentMethodOptionsCardInstallmentsPlanInterval is the interval of a card installment plan.
type PaymentIntentPaymentMethodOptionsCardInstallmentsPlanInterval string

// List of values that PaymentIntentPaymentMethodOptionsCardInstallmentsPlanInterval can take.
const (
	PaymentIntentPaymentMethodOptionsCardInstallmentsPlanIntervalMonth PaymentIntentPaymentMethodOptionsCardInstallmentsPlanInterval = "month"
)

// PaymentIntentPaymentMethodOptionsCardInstallmentsPlanType is the type of a card installment plan.
type PaymentIntentPaymentMethodOptionsCardInstallmentsPlanType string

// List of values that PaymentIntentPaymentMethodOptionsCardInstallmentsPlanType can take.
const (
	PaymentIntentPaymentMethodOptionsCardInstallmentsPlanTypeFixedCount PaymentIntentPaymentMethodOptionsCardInstallmentsPlanType = "fixed_count"
)

// PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure is the list of allowed values
//controlling when to request 3D Secure on a PaymentIntent.
type PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure string

// List of values that PaymentIntentNextActionType can take.
const (
	PaymentIntentPaymentMethodOptionsCardRequestThreeDSecureAny       PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure = "any"
	PaymentIntentPaymentMethodOptionsCardRequestThreeDSecureAutomatic PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure = "automatic"
)

// PaymentIntentSetupFutureUsage is the list of allowed values for SetupFutureUsage.
type PaymentIntentSetupFutureUsage string

// List of values that PaymentIntentSetupFutureUsage can take.
const (
	PaymentIntentSetupFutureUsageOffSession PaymentIntentSetupFutureUsage = "off_session"
	PaymentIntentSetupFutureUsageOnSession  PaymentIntentSetupFutureUsage = "on_session"
)

// PaymentIntentStatus is the list of allowed values for the payment intent's status.
type PaymentIntentStatus string

// List of values that PaymentIntentStatus can take.
const (
	PaymentIntentStatusCanceled              PaymentIntentStatus = "canceled"
	PaymentIntentStatusProcessing            PaymentIntentStatus = "processing"
	PaymentIntentStatusRequiresAction        PaymentIntentStatus = "requires_action"
	PaymentIntentStatusRequiresCapture       PaymentIntentStatus = "requires_capture"
	PaymentIntentStatusRequiresConfirmation  PaymentIntentStatus = "requires_confirmation"
	PaymentIntentStatusRequiresPaymentMethod PaymentIntentStatus = "requires_payment_method"
	PaymentIntentStatusSucceeded             PaymentIntentStatus = "succeeded"
)

// PaymentIntentCancelParams is the set of parameters that can be used when canceling a payment intent.
type PaymentIntentCancelParams struct {
	Params             `form:"*"`
	CancellationReason *string `form:"cancellation_reason"`
}

// PaymentIntentCaptureParams is the set of parameters that can be used when capturing a payment intent.
type PaymentIntentCaptureParams struct {
	Params                    `form:"*"`
	AmountToCapture           *int64                           `form:"amount_to_capture"`
	ApplicationFeeAmount      *int64                           `form:"application_fee_amount"`
	StatementDescriptor       *string                          `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                          `form:"statement_descriptor_suffix"`
	TransferData              *PaymentIntentTransferDataParams `form:"transfer_data"`
}

// PaymentIntentConfirmParams is the set of parameters that can be used when confirming a payment intent.
type PaymentIntentConfirmParams struct {
	Params                `form:"*"`
	ErrorOnRequiresAction *bool                                    `form:"error_on_requires_action"`
	Mandate               *string                                  `form:"mandate"`
	MandateData           *PaymentIntentMandateDataParams          `form:"mandate_data"`
	OffSession            *bool                                    `form:"off_session"`
	PaymentMethod         *string                                  `form:"payment_method"`
	PaymentMethodOptions  *PaymentIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	PaymentMethodTypes    []*string                                `form:"payment_method_types"`
	ReceiptEmail          *string                                  `form:"receipt_email"`
	ReturnURL             *string                                  `form:"return_url"`
	SavePaymentMethod     *bool                                    `form:"save_payment_method"`
	SetupFutureUsage      *string                                  `form:"setup_future_usage"`
	Shipping              *ShippingDetailsParams                   `form:"shipping"`
	Source                *string                                  `form:"source"`
	UseStripeSDK          *bool                                    `form:"use_stripe_sdk"`
}

// PaymentIntentMandateDataCustomerAcceptanceOfflineParams is the set of parameters for the customer
// acceptance of an offline mandate.
type PaymentIntentMandateDataCustomerAcceptanceOfflineParams struct {
}

// PaymentIntentMandateDataCustomerAcceptanceOnlineParams is the set of parameters for the customer
// acceptance of an online mandate.
type PaymentIntentMandateDataCustomerAcceptanceOnlineParams struct {
	IPAddress *string `form:"ip_address"`
	UserAgent *string `form:"user_agent"`
}

// PaymentIntentMandateDataCustomerAcceptanceParams is the set of parameters for the customer
// acceptance of a mandate.
type PaymentIntentMandateDataCustomerAcceptanceParams struct {
	AcceptedAt int64                                                    `form:"accepted_at"`
	Offline    *PaymentIntentMandateDataCustomerAcceptanceOfflineParams `form:"offline"`
	Online     *PaymentIntentMandateDataCustomerAcceptanceOnlineParams  `form:"online"`
	Type       MandateCustomerAcceptanceType                            `form:"type"`
}

// PaymentIntentMandateDataParams is the set of parameters controlling the creation of the mandate
// associated with this PaymentIntent.
type PaymentIntentMandateDataParams struct {
	CustomerAcceptance *PaymentIntentMandateDataCustomerAcceptanceParams `form:"customer_acceptance"`
}

// PaymentIntentPaymentMethodOptionsCardInstallmentsPlanParams represents details about the
// installment plan chosen for this payment intent.
type PaymentIntentPaymentMethodOptionsCardInstallmentsPlanParams struct {
	Count    *int64  `form:"count"`
	Interval *string `form:"interval"`
	Type     *string `form:"type"`
}

// PaymentIntentPaymentMethodOptionsCardInstallmentsParams controls whether to enable installment
// plans for this payment intent.
type PaymentIntentPaymentMethodOptionsCardInstallmentsParams struct {
	Enabled *bool                                                        `form:"enabled"`
	Plan    *PaymentIntentPaymentMethodOptionsCardInstallmentsPlanParams `form:"plan"`
}

// PaymentIntentPaymentMethodOptionsCardParams represents the card-specific options applied to a
// PaymentIntent.
type PaymentIntentPaymentMethodOptionsCardParams struct {
	Installments        *PaymentIntentPaymentMethodOptionsCardInstallmentsParams `form:"installments"`
	MOTO                *bool                                                    `form:"moto"`
	RequestThreeDSecure *string                                                  `form:"request_three_d_secure"`
}

// PaymentIntentPaymentMethodOptionsParams represents the type-specific payment method options
// applied to a PaymentIntent.
type PaymentIntentPaymentMethodOptionsParams struct {
	Card *PaymentIntentPaymentMethodOptionsCardParams `form:"card"`
}

// PaymentIntentTransferDataParams is the set of parameters allowed for the transfer hash.
type PaymentIntentTransferDataParams struct {
	Amount      *int64  `form:"amount"`
	Destination *string `form:"destination"`
}

// PaymentIntentParams is the set of parameters that can be used when handling a payment intent.
type PaymentIntentParams struct {
	Params                    `form:"*"`
	Amount                    *int64                                   `form:"amount"`
	ApplicationFeeAmount      *int64                                   `form:"application_fee_amount"`
	CaptureMethod             *string                                  `form:"capture_method"`
	Confirm                   *bool                                    `form:"confirm"`
	ConfirmationMethod        *string                                  `form:"confirmation_method"`
	Currency                  *string                                  `form:"currency"`
	Customer                  *string                                  `form:"customer"`
	Description               *string                                  `form:"description"`
	Mandate                   *string                                  `form:"mandate"`
	MandateData               *PaymentIntentMandateDataParams          `form:"mandate_data"`
	OnBehalfOf                *string                                  `form:"on_behalf_of"`
	PaymentMethod             *string                                  `form:"payment_method"`
	PaymentMethodOptions      *PaymentIntentPaymentMethodOptionsParams `form:"payment_method_options"`
	PaymentMethodTypes        []*string                                `form:"payment_method_types"`
	ReceiptEmail              *string                                  `form:"receipt_email"`
	ReturnURL                 *string                                  `form:"return_url"`
	SavePaymentMethod         *bool                                    `form:"save_payment_method"`
	SetupFutureUsage          *string                                  `form:"setup_future_usage"`
	Shipping                  *ShippingDetailsParams                   `form:"shipping"`
	Source                    *string                                  `form:"source"`
	StatementDescriptor       *string                                  `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                                  `form:"statement_descriptor_suffix"`
	TransferData              *PaymentIntentTransferDataParams         `form:"transfer_data"`
	TransferGroup             *string                                  `form:"transfer_group"`

	// Those parameters only works if you confirm on creation.
	ErrorOnRequiresAction *bool `form:"error_on_requires_action"`
	OffSession            *bool `form:"off_session"`
	UseStripeSDK          *bool `form:"use_stripe_sdk"`
}

// PaymentIntentListParams is the set of parameters that can be used when listing payment intents.
// For more details see https://stripe.com/docs/api#list_payouts.
type PaymentIntentListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Customer     *string           `form:"customer"`
}

// PaymentIntentNextActionRedirectToURL represents the resource for the next action of type
// "redirect_to_url".
type PaymentIntentNextActionRedirectToURL struct {
	ReturnURL string `json:"return_url"`
	URL       string `json:"url"`
}

// PaymentIntentNextAction represents the type of action to take on a payment intent.
type PaymentIntentNextAction struct {
	RedirectToURL *PaymentIntentNextActionRedirectToURL `json:"redirect_to_url"`
	Type          PaymentIntentNextActionType           `json:"type"`
}

// PaymentIntentPaymentMethodOptionsCardInstallmentsPlan describe a specific card installment plan.
type PaymentIntentPaymentMethodOptionsCardInstallmentsPlan struct {
	Count    int64                                                         `json:"count"`
	Interval PaymentIntentPaymentMethodOptionsCardInstallmentsPlanInterval `json:"interval"`
	Type     PaymentIntentPaymentMethodOptionsCardInstallmentsPlanType     `json:"type"`
}

// PaymentIntentPaymentMethodOptionsCardInstallments describe the installment options available for
// a card associated with that payment intent.
type PaymentIntentPaymentMethodOptionsCardInstallments struct {
	AvailablePlans []*PaymentIntentPaymentMethodOptionsCardInstallmentsPlan `json:"available_plans"`
	Enabled        bool                                                     `json:"enabled"`
	Plan           *PaymentIntentPaymentMethodOptionsCardInstallmentsPlan   `json:"plan"`
}

// PaymentIntentPaymentMethodOptionsCard is the set of card-specific options associated with that
// payment intent.
type PaymentIntentPaymentMethodOptionsCard struct {
	Installments        *PaymentIntentPaymentMethodOptionsCardInstallments       `json:"installments"`
	RequestThreeDSecure PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure `json:"request_three_d_secure"`
}

// PaymentIntentPaymentMethodOptions is the set of payment method-specific options associated with
// that payment intent.
type PaymentIntentPaymentMethodOptions struct {
	Card *PaymentIntentPaymentMethodOptionsCard `json:"card"`
}

// PaymentIntentTransferData represents the information for the transfer associated with a payment intent.
type PaymentIntentTransferData struct {
	Amount      int64    `json:"amount"`
	Destination *Account `json:"destination"`
}

// PaymentIntent is the resource representing a Stripe payout.
// For more details see https://stripe.com/docs/api#payment_intents.
type PaymentIntent struct {
	Amount                    int64                              `json:"amount"`
	AmountCapturable          int64                              `json:"amount_capturable"`
	AmountReceived            int64                              `json:"amount_received"`
	Application               *Application                       `json:"application"`
	ApplicationFeeAmount      int64                              `json:"application_fee_amount"`
	CanceledAt                int64                              `json:"canceled_at"`
	CancellationReason        PaymentIntentCancellationReason    `json:"cancellation_reason"`
	CaptureMethod             PaymentIntentCaptureMethod         `json:"capture_method"`
	Charges                   *ChargeList                        `json:"charges"`
	ClientSecret              string                             `json:"client_secret"`
	ConfirmationMethod        PaymentIntentConfirmationMethod    `json:"confirmation_method"`
	Created                   int64                              `json:"created"`
	Currency                  string                             `json:"currency"`
	Customer                  *Customer                          `json:"customer"`
	Description               string                             `json:"description"`
	Invoice                   *Invoice                           `json:"invoice"`
	LastPaymentError          *Error                             `json:"last_payment_error"`
	Livemode                  bool                               `json:"livemode"`
	ID                        string                             `json:"id"`
	Metadata                  map[string]string                  `json:"metadata"`
	NextAction                *PaymentIntentNextAction           `json:"next_action"`
	OnBehalfOf                *Account                           `json:"on_behalf_of"`
	PaymentMethod             *PaymentMethod                     `json:"payment_method"`
	PaymentMethodOptions      *PaymentIntentPaymentMethodOptions `json:"payment_method_options"`
	PaymentMethodTypes        []string                           `json:"payment_method_types"`
	ReceiptEmail              string                             `json:"receipt_email"`
	Review                    *Review                            `json:"review"`
	SetupFutureUsage          PaymentIntentSetupFutureUsage      `json:"setup_future_usage"`
	Shipping                  ShippingDetails                    `json:"shipping"`
	Source                    *PaymentSource                     `json:"source"`
	StatementDescriptor       string                             `json:"statement_descriptor"`
	StatementDescriptorSuffix string                             `json:"statement_descriptor_suffix"`
	Status                    PaymentIntentStatus                `json:"status"`
	TransferData              *PaymentIntentTransferData         `json:"transfer_data"`
	TransferGroup             string                             `json:"transfer_group"`
}

// PaymentIntentList is a list of payment intents as retrieved from a list endpoint.
type PaymentIntentList struct {
	ListMeta
	Data []*PaymentIntent `json:"data"`
}

// UnmarshalJSON handles deserialization of a Payment Intent.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PaymentIntent) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type paymentintent PaymentIntent
	var v paymentintent
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PaymentIntent(v)
	return nil
}
