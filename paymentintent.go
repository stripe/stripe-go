package stripe

import (
	"encoding/json"
)

type PaymentIntentCaptureMethod string

const (
	PaymentIntentCaptureMethodAutomatic PaymentIntentCaptureMethod = "automatic"
	PaymentIntentCaptureMethodManual    PaymentIntentCaptureMethod = "manual"
)

type PaymentIntentConfirmationMethod string

const (
	PaymentIntentConfirmationMethodPublishable PaymentIntentConfirmationMethod = "publishable"
	PaymentIntentConfirmationMethodSecret      PaymentIntentConfirmationMethod = "secret"
)

type PaymentIntentNextActionType string

const (
	PaymentIntentNextActionAuthorizeWithURL PaymentIntentNextActionType = "authorize_with_url"
)

type PaymentIntentStatus string

const (
	PaymentIntentStatusCanceled             PaymentIntentStatus = "canceled"
	PaymentIntentStatusProcessing           PaymentIntentStatus = "processing"
	PaymentIntentStatusRequiresCapture      PaymentIntentStatus = "requires_capture"
	PaymentIntentStatusRequiresConfirmation PaymentIntentStatus = "requires_confirmation"
	PaymentIntentStatusRequiresSource       PaymentIntentStatus = "requires_source"
	PaymentIntentStatusRequiresSourceAction PaymentIntentStatus = "requires_source_action"
	PaymentIntentStatusSucceeded            PaymentIntentStatus = "succeeded"
)

type PaymentIntentTransferDataParams struct {
	Amount *int64 `form:"amount"`
}

// PaymentIntentParams is the set of parameters that can be used when handling a payment intent.
type PaymentIntentParams struct {
	Params               `form:"*"`
	AllowedSourceTypes   []*string                        `form:"allowed_source_types"`
	Amount               *int64                           `form:"amount"`
	ApplicationFee       *int64                           `form:"application_fee"`
	AttemptConfirmation  *bool                            `form:"attempt_confirmation"`
	CaptureMethod        *string                          `form:"capture_method"`
	Currency             *string                          `form:"currency"`
	Customer             *string                          `form:"customer"`
	Description          *string                          `form:"description"`
	OnBehalfOf           *string                          `form:"on_behalf_of"`
	ReceiptEmail         *string                          `form:"receipt_email"`
	ReturnURL            *string                          `form:"return_url"`
	SaveSourceToCustomer *bool                            `form:"save_source_to_customer"`
	Shipping             *ShippingDetailsParams           `form:"shipping"`
	Source               *string                          `form:"source"`
	StatementDescriptor  *string                          `form:"statement_descriptor"`
	TransferData         *PaymentIntentTransferDataParams `form:"transfer_data"`
	TransferGroup        *string                          `form:"transfer_group"`
}

// PaymentIntentListParams is the set of parameters that can be used when listing payment intents.
// For more details see https://stripe.com/docs/api#list_payouts.
type PaymentIntentListParams struct {
	ListParams `form:"*"`
}

type PaymentIntentSourceActionAuthorizeWithURL struct {
	URL string `json:"url"`
}

// PaymentIntentSourceActionValue describes the `value` property in `next_source_action`
// The `type` in the parent should indicate which object is fleshed out.
type PaymentIntentSourceActionValue struct {
	AuthorizeWithURL *PaymentIntentSourceActionAuthorizeWithURL `json:"-"`
}

type PaymentIntentSourceAction struct {
	Type  PaymentIntentNextActionType     `json:"type"`
	Value *PaymentIntentSourceActionValue `json:"-"`
}

// UnmarshalJSON handles deserialization of a PaymentIntentSourceAction.
// This custom unmarshaling is needed because the specific
// type of for `value` it refers to is specified in the `type` property
func (s *PaymentIntentSourceAction) UnmarshalJSON(data []byte) error {
	type paymentIntentSourceAction PaymentIntentSourceAction
	var v paymentIntentSourceAction
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var err error
	*s = PaymentIntentSourceAction(v)
	s.Value = &PaymentIntentSourceActionValue{}

	// Unmarshal data a second time so that we can get the raw bytes for the
	// `value` field
	var rawObject map[string]*json.RawMessage
	if err := json.Unmarshal(data, &rawObject); err != nil {
		return err
	}

	switch s.Type {
	case PaymentIntentNextActionAuthorizeWithURL:
		err = json.Unmarshal(*rawObject["value"], &s.Value.AuthorizeWithURL)
	}

	return err
}

type PaymentIntentTransferData struct {
	Amount int64 `json:"amount"`
}

// Payout is the resource representing a Stripe payout.
// For more details see https://stripe.com/docs/api#payouts.
type PaymentIntent struct {
	AllowedSourceTypes  []string                        `json:"allowed_source_types"`
	Amount              int64                           `json:"amount"`
	AmountCapturable    int64                           `json:"amount_capturable"`
	AmountReceived      int64                           `json:"amount_received"`
	Application         *Application                    `json:"application"`
	ApplicationFee      int64                           `json:"application_fee"`
	CanceledAt          int64                           `json:"canceled_at"`
	CaptureMethod       PaymentIntentCaptureMethod      `json:"capture_method"`
	Charges             *ChargeList                     `json:"charges"`
	ClientSecret        string                          `json:"client_secret"`
	ConfirmationMethod  PaymentIntentConfirmationMethod `json:"confirmation_method"`
	Created             int64                           `json:"created"`
	Currency            string                          `json:"currency"`
	Customer            *Customer                       `json:"customer"`
	Description         string                          `json:"description"`
	Livemode            bool                            `json:"livemode"`
	ID                  string                          `json:"id"`
	Metadata            map[string]string               `json:"metadata"`
	NextSourceAction    *PaymentIntentSourceAction      `json:"next_source_action"`
	OnBehalfOf          *Account                        `json:"on_behalf_of"`
	ReceiptEmail        string                          `json:"receipt_email"`
	ReturnURL           string                          `json:"return_url"`
	Shipping            ShippingDetails                 `json:"shipping"`
	Source              *PaymentSource                  `json:"source"`
	StatementDescriptor string                          `json:"statement_descriptor"`
	Status              PaymentIntentStatus             `json:"status"`
	TransferData        *PaymentIntentTransferData      `json:"transfer_data"`
	TransferGroup       string                          `json:"transfer_group"`
}

// PayoutList is a list of payouts as retrieved from a list endpoint.
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
