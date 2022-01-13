//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
)

// The status of the code verification, either `pending` (awaiting verification, `attempts_remaining` should be greater than 0), `succeeded` (successful verification) or `failed` (failed verification, cannot be verified anymore as `attempts_remaining` should be 0).
type SourceCodeVerificationFlowStatus string

// List of values that SourceCodeVerificationFlowStatus can take
const (
	SourceCodeVerificationFlowStatusFailed    SourceCodeVerificationFlowStatus = "failed"
	SourceCodeVerificationFlowStatusPending   SourceCodeVerificationFlowStatus = "pending"
	SourceCodeVerificationFlowStatusSucceeded SourceCodeVerificationFlowStatus = "succeeded"
)

// The authentication `flow` of the source. `flow` is one of `redirect`, `receiver`, `code_verification`, `none`.
type SourceFlow string

// List of values that SourceFlow can take
const (
	SourceFlowCodeVerification SourceFlow = "code_verification"
	SourceFlowNone             SourceFlow = "none"
	SourceFlowReceiver         SourceFlow = "receiver"
	SourceFlowRedirect         SourceFlow = "redirect"
)

// SourceMandateAcceptanceStatus represents the possible failure reasons of a redirect flow.
type SourceMandateAcceptanceStatus string

// List of values that SourceMandateAcceptanceStatus can take.
const (
	SourceMandateAcceptanceStatusAccepted SourceMandateAcceptanceStatus = "accepted"
	SourceMandateAcceptanceStatusRefused  SourceMandateAcceptanceStatus = "refused"
)

// SourceMandateNotificationMethod represents the possible methods of notification for a mandate.
type SourceMandateNotificationMethod string

// List of values that SourceMandateNotificationMethod can take.
const (
	SourceMandateNotificationMethodEmail  SourceMandateNotificationMethod = "email"
	SourceMandateNotificationMethodManual SourceMandateNotificationMethod = "manual"
	SourceMandateNotificationMethodNone   SourceMandateNotificationMethod = "none"
)

// Type of refund attribute method, one of `email`, `manual`, or `none`.
type SourceRefundAttributesMethod string

// List of values that SourceRefundAttributesMethod can take
const (
	SourceRefundAttributesMethodEmail  SourceRefundAttributesMethod = "email"
	SourceRefundAttributesMethodManual SourceRefundAttributesMethod = "manual"
)

// Type of refund attribute status, one of `missing`, `requested`, or `available`.
type SourceRefundAttributesStatus string

// List of values that SourceRefundAttributesStatus can take
const (
	SourceRefundAttributesStatusAvailable SourceRefundAttributesStatus = "available"
	SourceRefundAttributesStatusMissing   SourceRefundAttributesStatus = "missing"
	SourceRefundAttributesStatusRequested SourceRefundAttributesStatus = "requested"
)

// The failure reason for the redirect, either `user_abort` (the customer aborted or dropped out of the redirect flow), `declined` (the authentication failed or the transaction was declined), or `processing_error` (the redirect failed due to a technical error). Present only if the redirect status is `failed`.
type SourceRedirectFlowFailureReason string

// List of values that SourceRedirectFlowFailureReason can take
const (
	SourceRedirectFlowFailureReasonDeclined        SourceRedirectFlowFailureReason = "declined"
	SourceRedirectFlowFailureReasonProcessingError SourceRedirectFlowFailureReason = "processing_error"
	SourceRedirectFlowFailureReasonUserAbort       SourceRedirectFlowFailureReason = "user_abort"
)

// The status of the redirect, either `pending` (ready to be used by your customer to authenticate the transaction), `succeeded` (succesful authentication, cannot be reused) or `not_required` (redirect should not be used) or `failed` (failed authentication, cannot be reused).
type SourceRedirectFlowStatus string

// List of values that SourceRedirectFlowStatus can take
const (
	SourceRedirectFlowStatusFailed      SourceRedirectFlowStatus = "failed"
	SourceRedirectFlowStatusNotRequired SourceRedirectFlowStatus = "not_required"
	SourceRedirectFlowStatusPending     SourceRedirectFlowStatus = "pending"
	SourceRedirectFlowStatusSucceeded   SourceRedirectFlowStatus = "succeeded"
)

// The type of this order item. Must be `sku`, `tax`, or `shipping`.
type SourceSourceOrderItemType string

// List of values that SourceSourceOrderItemType can take
const (
	SourceSourceOrderItemTypeDiscount SourceSourceOrderItemType = "discount"
	SourceSourceOrderItemTypeSKU      SourceSourceOrderItemType = "sku"
	SourceSourceOrderItemTypeShipping SourceSourceOrderItemType = "shipping"
	SourceSourceOrderItemTypeTax      SourceSourceOrderItemType = "tax"
)

// The status of the source, one of `canceled`, `chargeable`, `consumed`, `failed`, or `pending`. Only `chargeable` sources can be used to create a charge.
type SourceStatus string

// List of values that SourceStatus can take
const (
	SourceStatusCanceled   SourceStatus = "canceled"
	SourceStatusChargeable SourceStatus = "chargeable"
	SourceStatusConsumed   SourceStatus = "consumed"
	SourceStatusFailed     SourceStatus = "failed"
	SourceStatusPending    SourceStatus = "pending"
)

// Either `reusable` or `single_use`. Whether this source should be reusable or not. Some source types may or may not be reusable by construction, while others may leave the option at creation. If an incompatible value is passed, an error will be returned.
type SourceUsage string

// List of values that SourceUsage can take
const (
	SourceUsageReusable  SourceUsage = "reusable"
	SourceUsageSingleUse SourceUsage = "single_use"
)

// Delete a specified source for a given customer.
type SourceObjectDetachParams struct {
	Params   `form:"*"`
	Customer *string `form:"-"` // Included in URL
}

// Retrieves an existing source object. Supply the unique source ID from a source creation request and Stripe will return the corresponding up-to-date source object information.
type SourceObjectParams struct {
	Params              `form:"*"`
	Amount              *int64                `form:"amount"`
	ClientSecret        *string               `form:"client_secret"`
	Currency            *string               `form:"currency"`
	Customer            *string               `form:"customer"`
	Flow                *string               `form:"flow"`
	Mandate             *SourceMandateParams  `form:"mandate"`
	OriginalSource      *string               `form:"original_source"`
	Owner               *SourceOwnerParams    `form:"owner"`
	Receiver            *SourceReceiverParams `form:"receiver"`
	Redirect            *RedirectParams       `form:"redirect"`
	SourceOrder         *SourceOrderParams    `form:"source_order"`
	StatementDescriptor *string               `form:"statement_descriptor"`
	Token               *string               `form:"token"`
	Type                *string               `form:"type"`
	TypeData            map[string]string     `form:"-"`
	Usage               *string               `form:"usage"`
}

// The parameters required to store a mandate accepted offline. Should only be set if `mandate[type]` is `offline`
type SourceMandateAcceptanceOfflineParams struct {
	ContactEmail *string `form:"contact_email"`
}

// The parameters required to store a mandate accepted online. Should only be set if `mandate[type]` is `online`
type SourceMandateAcceptanceOnlineParams struct {
	Date      *int64  `form:"date"`
	IP        *string `form:"ip"`
	UserAgent *string `form:"user_agent"`
}

// The parameters required to notify Stripe of a mandate acceptance or refusal by the customer.
type SourceMandateAcceptanceParams struct {
	Date      *int64                                `form:"date"`
	IP        *string                               `form:"ip"`
	Offline   *SourceMandateAcceptanceOfflineParams `form:"offline"`
	Online    *SourceMandateAcceptanceOnlineParams  `form:"online"`
	Status    *string                               `form:"status"`
	Type      *string                               `form:"type"`
	UserAgent *string                               `form:"user_agent"`
}

// Information about a mandate possibility attached to a source object (generally for bank debits) as well as its acceptance status.
type SourceMandateParams struct {
	Acceptance         *SourceMandateAcceptanceParams `form:"acceptance"`
	Amount             *int64                         `form:"amount"`
	Currency           *string                        `form:"currency"`
	Interval           *string                        `form:"interval"`
	NotificationMethod *string                        `form:"notification_method"`
}

// Information about the owner of the payment instrument that may be used or required by particular source types.
type SourceOwnerParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// List of items constituting the order.
type SourceOrderItemsParams struct {
	Amount      *int64  `form:"amount"`
	Currency    *string `form:"currency"`
	Description *string `form:"description"`
	Parent      *string `form:"parent"`
	Quantity    *int64  `form:"quantity"`
	Type        *string `form:"type"`
}

// Information about the items and shipping associated with the source. Required for transactional credit (for example Klarna) sources before you can charge it.
type SourceOrderParams struct {
	Items    []*SourceOrderItemsParams `form:"items"`
	Shipping *ShippingDetailsParams    `form:"shipping"`
}

// Optional parameters for the receiver flow. Can be set only if the source is a receiver (`flow` is `receiver`).
type SourceReceiverParams struct {
	RefundAttributesMethod *string `form:"refund_attributes_method"`
}

// Parameters required for the redirect flow. Required if the source is authenticated by a redirect (`flow` is `redirect`).
type RedirectParams struct {
	ReturnURL *string `form:"return_url"`
}
type CodeVerificationFlow struct {
	AttemptsRemaining int64                            `json:"attempts_remaining"`
	Status            SourceCodeVerificationFlowStatus `json:"status"`
}

// SourceMandateAcceptance describes a source mandate acceptance state.
type SourceMandateAcceptance struct {
	Date      int64                         `json:"date"`
	IP        string                        `json:"ip"`
	Status    SourceMandateAcceptanceStatus `json:"status"`
	UserAgent string                        `json:"user_agent"`
}

// SourceMandate describes a source mandate.
type SourceMandate struct {
	Acceptance         *SourceMandateAcceptance        `json:"acceptance"`
	NotificationMethod SourceMandateNotificationMethod `json:"notification_method"`
	Reference          string                          `json:"reference"`
	URL                string                          `json:"url"`
}

// Information about the owner of the payment instrument that may be used or required by particular source types.
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
type ReceiverFlow struct {
	Address                string                       `json:"address"`
	AmountCharged          int64                        `json:"amount_charged"`
	AmountReceived         int64                        `json:"amount_received"`
	AmountReturned         int64                        `json:"amount_returned"`
	RefundAttributesMethod SourceRefundAttributesMethod `json:"refund_attributes_method"`
	RefundAttributesStatus SourceRefundAttributesStatus `json:"refund_attributes_status"`
}
type RedirectFlow struct {
	FailureReason SourceRedirectFlowFailureReason `json:"failure_reason"`
	ReturnURL     string                          `json:"return_url"`
	Status        SourceRedirectFlowStatus        `json:"status"`
	URL           string                          `json:"url"`
}

// List of items constituting the order.
type SourceSourceOrderItems struct {
	Amount      int64                     `json:"amount"`
	Currency    Currency                  `json:"currency"`
	Description string                    `json:"description"`
	Parent      string                    `json:"parent"`
	Quantity    int64                     `json:"quantity"`
	Type        SourceSourceOrderItemType `json:"type"`
}
type SourceSourceOrder struct {
	Amount   int64                     `json:"amount"`
	Currency Currency                  `json:"currency"`
	Email    string                    `json:"email"`
	Items    *[]SourceSourceOrderItems `json:"items"`
	Shipping *ShippingDetails          `json:"shipping"`
}

// `Source` objects allow you to accept a variety of payment methods. They
// represent a customer's payment instrument, and can be used with the Stripe API
// just like a `Card` object: once chargeable, they can be charged, or can be
// attached to customers.
//
// Related guides: [Sources API](https://stripe.com/docs/sources) and [Sources & Customers](https://stripe.com/docs/sources/customers).
type Source struct {
	APIResource
	Amount              int64                 `json:"amount"`
	ClientSecret        string                `json:"client_secret"`
	CodeVerification    *CodeVerificationFlow `json:"code_verification,omitempty"`
	Created             int64                 `json:"created"`
	Currency            Currency              `json:"currency"`
	Customer            string                `json:"customer"`
	Flow                SourceFlow            `json:"flow"`
	ID                  string                `json:"id"`
	Livemode            bool                  `json:"livemode"`
	Mandate             *SourceMandate        `json:"mandate"`
	Metadata            map[string]string     `json:"metadata"`
	Object              string                `json:"object"`
	Owner               *SourceOwner          `json:"owner"`
	Receiver            *ReceiverFlow         `json:"receiver,omitempty"`
	Redirect            *RedirectFlow         `json:"redirect,omitempty"`
	SourceOrder         *SourceSourceOrder    `json:"source_order"`
	StatementDescriptor string                `json:"statement_descriptor"`
	Status              SourceStatus          `json:"status"`
	Type                string                `json:"type"`
	TypeData            map[string]interface{}
	Usage               SourceUsage `json:"usage"`
}

// AppendTo implements custom encoding logic for SourceObjectParams so that the special
// "TypeData" value for is sent as the correct parameter based on the Source type
func (p *SourceObjectParams) AppendTo(body *form.Values, keyParts []string) {
	if len(p.TypeData) > 0 && p.Type == nil {
		panic("You can not fill TypeData if you don't explicitly set Type")
	}

	for k, vs := range p.TypeData {
		body.Add(form.FormatKey(append(keyParts, StringValue(p.Type), k)), vs)
	}
}

// UnmarshalJSON handles deserialization of a Source. This custom unmarshaling
// is needed to extract the type specific data (accessible under `TypeData`)
// but stored in JSON under a hash named after the `type` of the source.
func (s *Source) UnmarshalJSON(data []byte) error {
	type source Source
	var v source
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = Source(v)

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if d, ok := raw[s.Type]; ok {
		if m, ok := d.(map[string]interface{}); ok {
			s.TypeData = m
		}
	}

	return nil
}
