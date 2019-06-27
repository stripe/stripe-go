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
	Params        `form:"*"`
	PaymentMethod *string `form:"payment_method"`
	ReturnURL     *string `form:"return_url"`
}

// SetupIntentParams is the set of parameters that can be used when handling a setup intent.
type SetupIntentParams struct {
	Params             `form:"*"`
	Customer           *string   `form:"customer"`
	Description        *string   `form:"description"`
	OnBehalfOf         *string   `form:"on_behalf_of"`
	PaymentMethod      *string   `form:"payment_method"`
	PaymentMethodTypes []*string `form:"payment_method_types"`
	Usage              *string   `form:"usage"`
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

// SetupIntent is the resource representing a Stripe payout.
// For more details see https://stripe.com/docs/api#payment_intents.
type SetupIntent struct {
	Application        *Application                  `json:"application"`
	CancellationReason SetupIntentCancellationReason `json:"cancellation_reason"`
	ClientSecret       string                        `json:"client_secret"`
	Created            int64                         `json:"created"`
	Customer           *Customer                     `json:"customer"`
	Description        string                        `json:"description"`
	ID                 string                        `json:"id"`
	LastSetupError     *Error                        `json:"last_setup_error"`
	Livemode           bool                          `json:"livemode"`
	Metadata           map[string]string             `json:"metadata"`
	NextAction         *PaymentIntentNextAction      `json:"next_action"`
	Object             string                        `json:"object"`
	OnBehalfOf         *Account                      `json:"on_behalf_of"`
	PaymentMethod      *PaymentMethod                `json:"payment_method"`
	PaymentMethodTypes []string                      `json:"payment_method_types"`
	Status             SetupIntentStatus             `json:"status"`
	Usage              SetupIntentUsage              `json:"usage"`
}

// SetupIntentList is a list of setup intents as retrieved from a list endpoint.
type SetupIntentList struct {
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
