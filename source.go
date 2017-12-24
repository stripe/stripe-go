package stripe

import (
	"encoding/json"

	"github.com/stripe/stripe-go/form"
)

// SourceStatus represents the possible statuses of a source object.
type SourceStatus string

const (
	// SourceStatusCanceled we canceled the source along with any side-effect
	// it had (returned funds to customers if any were sent).
	SourceStatusCanceled SourceStatus = "canceled"

	// SourceStatusChargeable the source is ready to be charged (once if usage
	// is `single_use`, repeatedly otherwise).
	SourceStatusChargeable SourceStatus = "chargeable"

	// SourceStatusConsumed the source is `single_use` usage and has been
	// charged already.
	SourceStatusConsumed SourceStatus = "consumed"

	// SourceStatusFailed the source is no longer usable.
	SourceStatusFailed SourceStatus = "failed"

	// SourceStatusPending the source is freshly created and not yet
	// chargeable. The flow should indicate how to authenticate it with your
	// customer.
	SourceStatusPending SourceStatus = "pending"
)

// SourceFlow represents the possible flows of a source object.
type SourceFlow string

const (
	// FlowCodeVerification a verification code should be communicated by the
	// customer to authenticate the source.
	FlowCodeVerification SourceFlow = "code_verification"

	// FlowNone no particular authentication is involved the source should
	// become chargeable directly or asyncrhonously.
	FlowNone SourceFlow = "none"

	// FlowReceiver a receiver address should be communicated to the customer
	// to push funds to it.
	FlowReceiver SourceFlow = "receiver"

	// FlowRedirect a redirect is required to authenticate the source.
	FlowRedirect SourceFlow = "redirect"
)

// SourceUsage represents the possible usages of a source object.
type SourceUsage string

const (
	// UsageReusable the source can be charged multiple times for arbitrary
	// amounts.
	UsageReusable SourceUsage = "reusable"

	// UsageSingleUse the source can only be charged once for the specified
	// amount and currency.
	UsageSingleUse SourceUsage = "single_use"
)

type SourceOwnerParams struct {
	Address *AddressParams `form:"address"`
	Email   string         `form:"email"`
	Name    string         `form:"name"`
	Phone   string         `form:"phone"`
}

type RedirectParams struct {
	ReturnURL string `form:"return_url"`
}

type SourceObjectParams struct {
	Params              `form:"*"`
	Amount              uint64             `form:"amount"`
	Currency            Currency           `form:"currency"`
	Customer            string             `form:"customer"`
	Flow                SourceFlow         `form:"flow"`
	OriginalSource      string             `form:"original_source"`
	Owner               *SourceOwnerParams `form:"owner"`
	Redirect            *RedirectParams    `form:"redirect"`
	StatementDescriptor string             `form:"statement_descriptor"`
	Token               string             `form:"token"`
	Type                string             `form:"type"`
	TypeData            map[string]string  `form:"-"`
	Usage               SourceUsage        `form:"usage"`
}

// SourceObjectDetachParams is the set of parameters that can be used when detaching
// a source from a customer.
type SourceObjectDetachParams struct {
	Params   `form:"*"`
	Customer string `form:"-"`
}

type SourceOwner struct {
	Address         *Address `json:"address,omitempty"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	Phone           string   `json:"phone"`
	VerifiedAddress *Address `json:"verified_address,omitempty"`
	VerifiedEmail   string   `json:"verified_email"`
	VerifiedName    string   `json:"verified_name"`
	VerifiedPhone   string   `json:"verified_phone"`
}

// RedirectFlowFailureReason represents the possible failure reasons of a redirect flow.
type RedirectFlowFailureReason string

const (
	RedirectFlowFailureReasonDeclined        RedirectFlowFailureReason = "declined"
	RedirectFlowFailureReasonProcessingError RedirectFlowFailureReason = "processing_error"
	RedirectFlowFailureReasonUserAbort       RedirectFlowFailureReason = "user_abort"
)

// RedirectFlowStatus represents the possible statuses of a redirect flow.
type RedirectFlowStatus string

const (
	RedirectFlowStatusFailed    RedirectFlowStatus = "failed"
	RedirectFlowStatusPending   RedirectFlowStatus = "pending"
	RedirectFlowStatusSucceeded RedirectFlowStatus = "succeeded"
)

// ReceiverFlow informs of the state of a redirect authentication flow.
type RedirectFlow struct {
	FailureReason RedirectFlowFailureReason `json:"failure_reason"`
	ReturnURL     string                    `json:"return_url"`
	Status        RedirectFlowStatus        `json:"status"`
	URL           string                    `json:"url"`
}

// RefundAttributesStatus are the possible status of a receiver's refund
// attributes.
type RefundAttributesStatus string

const (
	// RefundAttributesAvailable the refund attributes are available
	RefundAttributesAvailable RefundAttributesStatus = "available"

	// RefundAttributesMissing the refund attributes are missing
	RefundAttributesMissing RefundAttributesStatus = "missing"

	// RefundAttributesRequested the refund attributes have been requested
	RefundAttributesRequested RefundAttributesStatus = "requested"
)

// RefundAttributesMethod are the possible method to retrieve a receiver's
// refund attributes.
type RefundAttributesMethod string

const (
	// RefundAttributesEmail the refund attributes are automatically collected over email
	RefundAttributesEmail RefundAttributesMethod = "email"

	// RefundAttributesManual the refund attributes should be collected by the user
	RefundAttributesManual RefundAttributesMethod = "manual"
)

// ReceiverFlow informs of the state of a receiver authentication flow.
type ReceiverFlow struct {
	Address                string                 `json:"address"`
	AmountCharged          int64                  `json:"amount_charged"`
	AmountReceived         int64                  `json:"amount_received"`
	AmountReturned         int64                  `json:"amount_returned"`
	RefundAttributesMethod RefundAttributesMethod `json:"refund_attributes_method"`
	RefundAttributesStatus RefundAttributesStatus `json:"refund_attributes_status"`
}

// CodeVerificationFlowStatus represents the possible statuses of a code verification
// flow.
type CodeVerificationFlowStatus string

const (
	CodeVerificationFlowStatusFailed    CodeVerificationFlowStatus = "failed"
	CodeVerificationFlowStatusPending   CodeVerificationFlowStatus = "pending"
	CodeVerificationFlowStatusSucceeded CodeVerificationFlowStatus = "succeeded"
)

// CodeVerificationFlow informs of the state of a verification authentication flow.
type CodeVerificationFlow struct {
	AttemptsRemaining uint64                     `json:"attempts_remaining"`
	Status            CodeVerificationFlowStatus `json:"status"`
}

type Source struct {
	Amount              int64                 `json:"amount"`
	ClientSecret        string                `json:"client_secret"`
	CodeVerification    *CodeVerificationFlow `json:"code_verification,omitempty"`
	Created             int64                 `json:"created"`
	Currency            Currency              `json:"currency"`
	Flow                SourceFlow            `json:"flow"`
	ID                  string                `json:"id"`
	Live                bool                  `json:"livemode"`
	Meta                map[string]string     `json:"metadata"`
	Owner               SourceOwner           `json:"owner"`
	Receiver            *ReceiverFlow         `json:"receiver,omitempty"`
	Redirect            *RedirectFlow         `json:"redirect,omitempty"`
	StatementDescriptor string                `json:"statement_descriptor"`
	Status              SourceStatus          `json:"status"`
	Type                string                `json:"type"`
	TypeData            map[string]interface{}
	Usage               SourceUsage `json:"usage"`
}

// AppendTo implements custom encoding logic for SourceObjectParams so that the special
// "TypeData" value for is sent as the correct parameter based on the Source type
func (p *SourceObjectParams) AppendTo(body *form.Values, keyParts []string) {
	if len(p.TypeData) > 0 && len(p.Type) == 0 {
		panic("You can not fill TypeData if you don't explicitly set Type")
	}

	for k, vs := range p.TypeData {
		body.Add(form.FormatKey(append(keyParts, p.Type, k)), vs)
	}
}

// UnmarshalJSON handles deserialization of an Source. This custom unmarshaling
// is needed to extract the type specific data (accessible under `TypeData`)
// but stored in JSON under a hash named after the `type` of the source.
func (s *Source) UnmarshalJSON(data []byte) error {
	type source Source
	var ss source
	err := json.Unmarshal(data, &ss)
	if err != nil {
		return err
	}
	*s = Source(ss)

	var raw map[string]interface{}
	err = json.Unmarshal(data, &raw)
	if err != nil {
		return err
	}
	if d, ok := raw[s.Type]; ok {
		if m, ok := d.(map[string]interface{}); ok {
			s.TypeData = m
		}
	}

	return nil
}
