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
	Params `form:"*"`
	// Amount associated with the source. This is the amount for which the source will be chargeable once ready. Required for `single_use` sources. Not supported for `receiver` type sources, where charge amount may not be specified until funds land.
	Amount *int64 `form:"amount"`
	// The client secret of the source. Required if a publishable key is used to retrieve the source.
	ClientSecret *string `form:"client_secret"`
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) associated with the source. This is the currency for which the source will be chargeable once ready.
	Currency *string `form:"currency"`
	// The `Customer` to whom the original source is attached to. Must be set when the original source is not a `Source` (e.g., `Card`).
	Customer *string `form:"customer"`
	// The authentication `flow` of the source to create. `flow` is one of `redirect`, `receiver`, `code_verification`, `none`. It is generally inferred unless a type supports multiple flows.
	Flow *string `form:"flow"`
	// Information about a mandate possibility attached to a source object (generally for bank debits) as well as its acceptance status.
	Mandate *SourceMandateParams `form:"mandate"`
	// The source to share.
	OriginalSource *string `form:"original_source"`
	// Information about the owner of the payment instrument that may be used or required by particular source types.
	Owner *SourceOwnerParams `form:"owner"`
	// Optional parameters for the receiver flow. Can be set only if the source is a receiver (`flow` is `receiver`).
	Receiver *SourceReceiverParams `form:"receiver"`
	// Parameters required for the redirect flow. Required if the source is authenticated by a redirect (`flow` is `redirect`).
	Redirect *RedirectParams `form:"redirect"`
	// Information about the items and shipping associated with the source. Required for transactional credit (for example Klarna) sources before you can charge it.
	SourceOrder *SourceOrderParams `form:"source_order"`
	// An arbitrary string to be displayed on your customer's statement. As an example, if your website is `RunClub` and the item you're charging for is a race ticket, you may want to specify a `statement_descriptor` of `RunClub 5K race ticket.` While many payment types will display this information, some may not display it at all.
	StatementDescriptor *string `form:"statement_descriptor"`
	// An optional token used to create the source. When passed, token properties will override source parameters.
	Token *string `form:"token"`
	// The `type` of the source to create. Required unless `customer` and `original_source` are specified (see the [Cloning card Sources](https://stripe.com/docs/sources/connect#cloning-card-sources) guide)
	Type     *string           `form:"type"`
	TypeData map[string]string `form:"-"`
	Usage    *string           `form:"usage"`
}

// The parameters required to store a mandate accepted offline. Should only be set if `mandate[type]` is `offline`
type SourceMandateAcceptanceOfflineParams struct {
	// An email to contact you with if a copy of the mandate is requested, required if `type` is `offline`.
	ContactEmail *string `form:"contact_email"`
}

// The parameters required to store a mandate accepted online. Should only be set if `mandate[type]` is `online`
type SourceMandateAcceptanceOnlineParams struct {
	// The Unix timestamp (in seconds) when the mandate was accepted or refused by the customer.
	Date *int64 `form:"date"`
	// The IP address from which the mandate was accepted or refused by the customer.
	IP *string `form:"ip"`
	// The user agent of the browser from which the mandate was accepted or refused by the customer.
	UserAgent *string `form:"user_agent"`
}

// The parameters required to notify Stripe of a mandate acceptance or refusal by the customer.
type SourceMandateAcceptanceParams struct {
	// The Unix timestamp (in seconds) when the mandate was accepted or refused by the customer.
	Date *int64 `form:"date"`
	// The IP address from which the mandate was accepted or refused by the customer.
	IP *string `form:"ip"`
	// The parameters required to store a mandate accepted offline. Should only be set if `mandate[type]` is `offline`
	Offline *SourceMandateAcceptanceOfflineParams `form:"offline"`
	// The parameters required to store a mandate accepted online. Should only be set if `mandate[type]` is `online`
	Online *SourceMandateAcceptanceOnlineParams `form:"online"`
	// The status of the mandate acceptance. Either `accepted` (the mandate was accepted) or `refused` (the mandate was refused).
	Status *string `form:"status"`
	// The type of acceptance information included with the mandate. Either `online` or `offline`
	Type *string `form:"type"`
	// The user agent of the browser from which the mandate was accepted or refused by the customer.
	UserAgent *string `form:"user_agent"`
}

// Information about a mandate possibility attached to a source object (generally for bank debits) as well as its acceptance status.
type SourceMandateParams struct {
	// The parameters required to notify Stripe of a mandate acceptance or refusal by the customer.
	Acceptance *SourceMandateAcceptanceParams `form:"acceptance"`
	// The amount specified by the mandate. (Leave null for a mandate covering all amounts)
	Amount *int64 `form:"amount"`
	// The currency specified by the mandate. (Must match `currency` of the source)
	Currency *string `form:"currency"`
	// The interval of debits permitted by the mandate. Either `one_time` (just permitting a single debit), `scheduled` (with debits on an agreed schedule or for clearly-defined events), or `variable`(for debits with any frequency)
	Interval *string `form:"interval"`
	// The method Stripe should use to notify the customer of upcoming debit instructions and/or mandate confirmation as required by the underlying debit network. Either `email` (an email is sent directly to the customer), `manual` (a `source.mandate_notification` event is sent to your webhooks endpoint and you should handle the notification) or `none` (the underlying debit network does not require any notification).
	NotificationMethod *string `form:"notification_method"`
}

// Information about the owner of the payment instrument that may be used or required by particular source types.
type SourceOwnerParams struct {
	// Owner's address.
	Address *AddressParams `form:"address"`
	// Owner's email address.
	Email *string `form:"email"`
	// Owner's full name.
	Name *string `form:"name"`
	// Owner's phone number.
	Phone *string `form:"phone"`
}

// List of items constituting the order.
type SourceOrderItemsParams struct {
	Amount      *int64  `form:"amount"`
	Currency    *string `form:"currency"`
	Description *string `form:"description"`
	// The ID of the SKU being ordered.
	Parent *string `form:"parent"`
	// The quantity of this order item. When type is `sku`, this is the number of instances of the SKU to be ordered.
	Quantity *int64  `form:"quantity"`
	Type     *string `form:"type"`
}

// Information about the items and shipping associated with the source. Required for transactional credit (for example Klarna) sources before you can charge it.
type SourceOrderParams struct {
	// List of items constituting the order.
	Items []*SourceOrderItemsParams `form:"items"`
	// Shipping address for the order. Required if any of the SKUs are for products that have `shippable` set to true.
	Shipping *ShippingDetailsParams `form:"shipping"`
}

// Optional parameters for the receiver flow. Can be set only if the source is a receiver (`flow` is `receiver`).
type SourceReceiverParams struct {
	// The method Stripe should use to request information needed to process a refund or mispayment. Either `email` (an email is sent directly to the customer) or `manual` (a `source.refund_attributes_required` event is sent to your webhooks endpoint). Refer to each payment method's documentation to learn which refund attributes may be required.
	RefundAttributesMethod *string `form:"refund_attributes_method"`
}

// Parameters required for the redirect flow. Required if the source is authenticated by a redirect (`flow` is `redirect`).
type RedirectParams struct {
	// The URL you provide to redirect the customer back to you after they authenticated their payment. It can use your application URI scheme in the context of a mobile application.
	ReturnURL *string `form:"return_url"`
}
type CodeVerificationFlow struct {
	// The number of attempts remaining to authenticate the source object with a verification code.
	AttemptsRemaining int64 `json:"attempts_remaining"`
	// The status of the code verification, either `pending` (awaiting verification, `attempts_remaining` should be greater than 0), `succeeded` (successful verification) or `failed` (failed verification, cannot be verified anymore as `attempts_remaining` should be 0).
	Status SourceCodeVerificationFlowStatus `json:"status"`
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
	// Owner's address.
	Address *Address `json:"address,omitempty"`
	// Owner's email address.
	Email string `json:"email"`
	// Owner's full name.
	Name string `json:"name"`
	// Owner's phone number (including extension).
	Phone string `json:"phone"`
	// Verified owner's address. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedAddress *Address `json:"verified_address,omitempty"`
	// Verified owner's email address. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedEmail string `json:"verified_email"`
	// Verified owner's full name. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name"`
	// Verified owner's phone number (including extension). Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedPhone string `json:"verified_phone"`
}
type ReceiverFlow struct {
	// The address of the receiver source. This is the value that should be communicated to the customer to send their funds to.
	Address string `json:"address"`
	// The total amount that was moved to your balance. This is almost always equal to the amount charged. In rare cases when customers deposit excess funds and we are unable to refund those, those funds get moved to your balance and show up in amount_charged as well. The amount charged is expressed in the source's currency.
	AmountCharged int64 `json:"amount_charged"`
	// The total amount received by the receiver source. `amount_received = amount_returned + amount_charged` should be true for consumed sources unless customers deposit excess funds. The amount received is expressed in the source's currency.
	AmountReceived int64 `json:"amount_received"`
	// The total amount that was returned to the customer. The amount returned is expressed in the source's currency.
	AmountReturned int64 `json:"amount_returned"`
	// Type of refund attribute method, one of `email`, `manual`, or `none`.
	RefundAttributesMethod SourceRefundAttributesMethod `json:"refund_attributes_method"`
	// Type of refund attribute status, one of `missing`, `requested`, or `available`.
	RefundAttributesStatus SourceRefundAttributesStatus `json:"refund_attributes_status"`
}
type RedirectFlow struct {
	// The failure reason for the redirect, either `user_abort` (the customer aborted or dropped out of the redirect flow), `declined` (the authentication failed or the transaction was declined), or `processing_error` (the redirect failed due to a technical error). Present only if the redirect status is `failed`.
	FailureReason SourceRedirectFlowFailureReason `json:"failure_reason"`
	// The URL you provide to redirect the customer to after they authenticated their payment.
	ReturnURL string `json:"return_url"`
	// The status of the redirect, either `pending` (ready to be used by your customer to authenticate the transaction), `succeeded` (succesful authentication, cannot be reused) or `not_required` (redirect should not be used) or `failed` (failed authentication, cannot be reused).
	Status SourceRedirectFlowStatus `json:"status"`
	// The URL provided to you to redirect a customer to as part of a `redirect` authentication flow.
	URL string `json:"url"`
}

// List of items constituting the order.
type SourceSourceOrderItems struct {
	// The amount (price) for this order item.
	Amount int64 `json:"amount"`
	// This currency of this order item. Required when `amount` is present.
	Currency Currency `json:"currency"`
	// Human-readable description for this order item.
	Description string `json:"description"`
	// The ID of the associated object for this line item. Expandable if not null (e.g., expandable to a SKU).
	Parent string `json:"parent"`
	// The quantity of this order item. When type is `sku`, this is the number of instances of the SKU to be ordered.
	Quantity int64 `json:"quantity"`
	// The type of this order item. Must be `sku`, `tax`, or `shipping`.
	Type SourceSourceOrderItemType `json:"type"`
}
type SourceSourceOrder struct {
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount for the order.
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The email address of the customer placing the order.
	Email string `json:"email"`
	// List of items constituting the order.
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
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount associated with the source. This is the amount for which the source will be chargeable once ready. Required for `single_use` sources.
	Amount int64 `json:"amount"`
	// The client secret of the source. Used for client-side retrieval using a publishable key.
	ClientSecret     string                `json:"client_secret"`
	CodeVerification *CodeVerificationFlow `json:"code_verification,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) associated with the source. This is the currency for which the source will be chargeable once ready. Required for `single_use` sources.
	Currency Currency `json:"currency"`
	// The ID of the customer to which this source is attached. This will not be present when the source has not been attached to a customer.
	Customer string `json:"customer"`
	// The authentication `flow` of the source. `flow` is one of `redirect`, `receiver`, `code_verification`, `none`.
	Flow SourceFlow `json:"flow"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool           `json:"livemode"`
	Mandate  *SourceMandate `json:"mandate"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Information about the owner of the payment instrument that may be used or required by particular source types.
	Owner       *SourceOwner       `json:"owner"`
	Receiver    *ReceiverFlow      `json:"receiver,omitempty"`
	Redirect    *RedirectFlow      `json:"redirect,omitempty"`
	SourceOrder *SourceSourceOrder `json:"source_order"`
	// Extra information about a source. This will appear on your customer's statement every time you charge the source.
	StatementDescriptor string `json:"statement_descriptor"`
	// The status of the source, one of `canceled`, `chargeable`, `consumed`, `failed`, or `pending`. Only `chargeable` sources can be used to create a charge.
	Status SourceStatus `json:"status"`
	// The `type` of the source. The `type` is a payment method, one of `ach_credit_transfer`, `ach_debit`, `alipay`, `bancontact`, `card`, `card_present`, `eps`, `giropay`, `ideal`, `multibanco`, `klarna`, `p24`, `sepa_debit`, `sofort`, `three_d_secure`, or `wechat`. An additional hash is included on the source with a name matching this value. It contains additional information specific to the [payment method](https://stripe.com/docs/sources) used.
	Type     string `json:"type"`
	TypeData map[string]interface{}
	// Either `reusable` or `single_use`. Whether this source should be reusable or not. Some source types may or may not be reusable by construction, while others may leave the option at creation. If an incompatible value is passed, an error will be returned.
	Usage SourceUsage `json:"usage"`
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
