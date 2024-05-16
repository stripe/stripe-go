//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Payment method type.
type DisputePaymentMethodDetailsType string

// List of values that DisputePaymentMethodDetailsType can take
const (
	DisputePaymentMethodDetailsTypeCard   DisputePaymentMethodDetailsType = "card"
	DisputePaymentMethodDetailsTypeKlarna DisputePaymentMethodDetailsType = "klarna"
	DisputePaymentMethodDetailsTypePaypal DisputePaymentMethodDetailsType = "paypal"
)

// Reason given by cardholder for dispute. Possible values are `bank_cannot_process`, `check_returned`, `credit_not_processed`, `customer_initiated`, `debit_not_authorized`, `duplicate`, `fraudulent`, `general`, `incorrect_account_details`, `insufficient_funds`, `product_not_received`, `product_unacceptable`, `subscription_canceled`, or `unrecognized`. Learn more about [dispute reasons](https://stripe.com/docs/disputes/categories).
type DisputeReason string

// List of values that DisputeReason can take
const (
	DisputeReasonBankCannotProcess       DisputeReason = "bank_cannot_process"
	DisputeReasonCheckReturned           DisputeReason = "check_returned"
	DisputeReasonCreditNotProcessed      DisputeReason = "credit_not_processed"
	DisputeReasonCustomerInitiated       DisputeReason = "customer_initiated"
	DisputeReasonDebitNotAuthorized      DisputeReason = "debit_not_authorized"
	DisputeReasonDuplicate               DisputeReason = "duplicate"
	DisputeReasonFraudulent              DisputeReason = "fraudulent"
	DisputeReasonGeneral                 DisputeReason = "general"
	DisputeReasonIncorrectAccountDetails DisputeReason = "incorrect_account_details"
	DisputeReasonInsufficientFunds       DisputeReason = "insufficient_funds"
	DisputeReasonProductNotReceived      DisputeReason = "product_not_received"
	DisputeReasonProductUnacceptable     DisputeReason = "product_unacceptable"
	DisputeReasonSubscriptionCanceled    DisputeReason = "subscription_canceled"
	DisputeReasonUnrecognized            DisputeReason = "unrecognized"
)

// Current status of dispute. Possible values are `warning_needs_response`, `warning_under_review`, `warning_closed`, `needs_response`, `under_review`, `won`, or `lost`.
type DisputeStatus string

// List of values that DisputeStatus can take
const (
	DisputeStatusLost                 DisputeStatus = "lost"
	DisputeStatusNeedsResponse        DisputeStatus = "needs_response"
	DisputeStatusUnderReview          DisputeStatus = "under_review"
	DisputeStatusWarningClosed        DisputeStatus = "warning_closed"
	DisputeStatusWarningNeedsResponse DisputeStatus = "warning_needs_response"
	DisputeStatusWarningUnderReview   DisputeStatus = "warning_under_review"
	DisputeStatusWon                  DisputeStatus = "won"
)

// Returns a list of your disputes.
type DisputeListParams struct {
	ListParams `form:"*"`
	// Only return disputes associated to the charge specified by this charge ID.
	Charge       *string           `form:"charge"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return disputes associated to the PaymentIntent specified by this PaymentIntent ID.
	PaymentIntent *string `form:"payment_intent"`
}

// AddExpand appends a new field to expand.
func (p *DisputeListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the dispute with the given ID.
type DisputeParams struct {
	Params `form:"*"`
	// Evidence to upload, to respond to a dispute. Updating any field in the hash will submit all fields in the hash for review. The combined character count of all fields is limited to 150,000.
	Evidence *DisputeEvidenceParams `form:"evidence"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Whether to immediately submit evidence to the bank. If `false`, evidence is staged on the dispute. Staged evidence is visible in the API and Dashboard, and can be submitted to the bank by making another request with this attribute set to `true` (the default).
	Submit *bool `form:"submit"`
}

// AddExpand appends a new field to expand.
func (p *DisputeParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *DisputeParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Evidence to upload, to respond to a dispute. Updating any field in the hash will submit all fields in the hash for review. The combined character count of all fields is limited to 150,000.
type DisputeEvidenceParams struct {
	// Any server or activity logs showing proof that the customer accessed or downloaded the purchased digital product. This information should include IP addresses, corresponding timestamps, and any detailed recorded activity. Has a maximum character count of 20,000.
	AccessActivityLog *string `form:"access_activity_log"`
	// The billing address provided by the customer.
	BillingAddress *string `form:"billing_address"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Your subscription cancellation policy, as shown to the customer.
	CancellationPolicy *string `form:"cancellation_policy"`
	// An explanation of how and when the customer was shown your refund policy prior to purchase. Has a maximum character count of 20,000.
	CancellationPolicyDisclosure *string `form:"cancellation_policy_disclosure"`
	// A justification for why the customer's subscription was not canceled. Has a maximum character count of 20,000.
	CancellationRebuttal *string `form:"cancellation_rebuttal"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Any communication with the customer that you feel is relevant to your case. Examples include emails proving that the customer received the product or service, or demonstrating their use of or satisfaction with the product or service.
	CustomerCommunication *string `form:"customer_communication"`
	// The email address of the customer.
	CustomerEmailAddress *string `form:"customer_email_address"`
	// The name of the customer.
	CustomerName *string `form:"customer_name"`
	// The IP address that the customer used when making the purchase.
	CustomerPurchaseIP *string `form:"customer_purchase_ip"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) A relevant document or contract showing the customer's signature.
	CustomerSignature *string `form:"customer_signature"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Documentation for the prior charge that can uniquely identify the charge, such as a receipt, shipping label, work order, etc. This document should be paired with a similar document from the disputed payment that proves the two payments are separate.
	DuplicateChargeDocumentation *string `form:"duplicate_charge_documentation"`
	// An explanation of the difference between the disputed charge versus the prior charge that appears to be a duplicate. Has a maximum character count of 20,000.
	DuplicateChargeExplanation *string `form:"duplicate_charge_explanation"`
	// The Stripe ID for the prior charge which appears to be a duplicate of the disputed charge.
	DuplicateChargeID *string `form:"duplicate_charge_id"`
	// A description of the product or service that was sold. Has a maximum character count of 20,000.
	ProductDescription *string `form:"product_description"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Any receipt or message sent to the customer notifying them of the charge.
	Receipt *string `form:"receipt"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Your refund policy, as shown to the customer.
	RefundPolicy *string `form:"refund_policy"`
	// Documentation demonstrating that the customer was shown your refund policy prior to purchase. Has a maximum character count of 20,000.
	RefundPolicyDisclosure *string `form:"refund_policy_disclosure"`
	// A justification for why the customer is not entitled to a refund. Has a maximum character count of 20,000.
	RefundRefusalExplanation *string `form:"refund_refusal_explanation"`
	// The date on which the customer received or began receiving the purchased service, in a clear human-readable format.
	ServiceDate *string `form:"service_date"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Documentation showing proof that a service was provided to the customer. This could include a copy of a signed contract, work order, or other form of written agreement.
	ServiceDocumentation *string `form:"service_documentation"`
	// The address to which a physical product was shipped. You should try to include as complete address information as possible.
	ShippingAddress *string `form:"shipping_address"`
	// The delivery service that shipped a physical product, such as Fedex, UPS, USPS, etc. If multiple carriers were used for this purchase, please separate them with commas.
	ShippingCarrier *string `form:"shipping_carrier"`
	// The date on which a physical product began its route to the shipping address, in a clear human-readable format.
	ShippingDate *string `form:"shipping_date"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Documentation showing proof that a product was shipped to the customer at the same address the customer provided to you. This could include a copy of the shipment receipt, shipping label, etc. It should show the customer's full shipping address, if possible.
	ShippingDocumentation *string `form:"shipping_documentation"`
	// The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.
	ShippingTrackingNumber *string `form:"shipping_tracking_number"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Any additional evidence or statements.
	UncategorizedFile *string `form:"uncategorized_file"`
	// Any additional evidence or statements. Has a maximum character count of 20,000.
	UncategorizedText *string `form:"uncategorized_text"`
}
type DisputeEvidence struct {
	// Any server or activity logs showing proof that the customer accessed or downloaded the purchased digital product. This information should include IP addresses, corresponding timestamps, and any detailed recorded activity.
	AccessActivityLog string `json:"access_activity_log"`
	// The billing address provided by the customer.
	BillingAddress string `json:"billing_address"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Your subscription cancellation policy, as shown to the customer.
	CancellationPolicy *File `json:"cancellation_policy"`
	// An explanation of how and when the customer was shown your refund policy prior to purchase.
	CancellationPolicyDisclosure string `json:"cancellation_policy_disclosure"`
	// A justification for why the customer's subscription was not canceled.
	CancellationRebuttal string `json:"cancellation_rebuttal"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Any communication with the customer that you feel is relevant to your case. Examples include emails proving that the customer received the product or service, or demonstrating their use of or satisfaction with the product or service.
	CustomerCommunication *File `json:"customer_communication"`
	// The email address of the customer.
	CustomerEmailAddress string `json:"customer_email_address"`
	// The name of the customer.
	CustomerName string `json:"customer_name"`
	// The IP address that the customer used when making the purchase.
	CustomerPurchaseIP string `json:"customer_purchase_ip"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) A relevant document or contract showing the customer's signature.
	CustomerSignature *File `json:"customer_signature"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Documentation for the prior charge that can uniquely identify the charge, such as a receipt, shipping label, work order, etc. This document should be paired with a similar document from the disputed payment that proves the two payments are separate.
	DuplicateChargeDocumentation *File `json:"duplicate_charge_documentation"`
	// An explanation of the difference between the disputed charge versus the prior charge that appears to be a duplicate.
	DuplicateChargeExplanation string `json:"duplicate_charge_explanation"`
	// The Stripe ID for the prior charge which appears to be a duplicate of the disputed charge.
	DuplicateChargeID string `json:"duplicate_charge_id"`
	// A description of the product or service that was sold.
	ProductDescription string `json:"product_description"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Any receipt or message sent to the customer notifying them of the charge.
	Receipt *File `json:"receipt"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Your refund policy, as shown to the customer.
	RefundPolicy *File `json:"refund_policy"`
	// Documentation demonstrating that the customer was shown your refund policy prior to purchase.
	RefundPolicyDisclosure string `json:"refund_policy_disclosure"`
	// A justification for why the customer is not entitled to a refund.
	RefundRefusalExplanation string `json:"refund_refusal_explanation"`
	// The date on which the customer received or began receiving the purchased service, in a clear human-readable format.
	ServiceDate string `json:"service_date"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Documentation showing proof that a service was provided to the customer. This could include a copy of a signed contract, work order, or other form of written agreement.
	ServiceDocumentation *File `json:"service_documentation"`
	// The address to which a physical product was shipped. You should try to include as complete address information as possible.
	ShippingAddress string `json:"shipping_address"`
	// The delivery service that shipped a physical product, such as Fedex, UPS, USPS, etc. If multiple carriers were used for this purchase, please separate them with commas.
	ShippingCarrier string `json:"shipping_carrier"`
	// The date on which a physical product began its route to the shipping address, in a clear human-readable format.
	ShippingDate string `json:"shipping_date"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Documentation showing proof that a product was shipped to the customer at the same address the customer provided to you. This could include a copy of the shipment receipt, shipping label, etc. It should show the customer's full shipping address, if possible.
	ShippingDocumentation *File `json:"shipping_documentation"`
	// The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.
	ShippingTrackingNumber string `json:"shipping_tracking_number"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Any additional evidence or statements.
	UncategorizedFile *File `json:"uncategorized_file"`
	// Any additional evidence or statements.
	UncategorizedText string `json:"uncategorized_text"`
}
type DisputeEvidenceDetails struct {
	// Date by which evidence must be submitted in order to successfully challenge dispute. Will be 0 if the customer's bank or credit card company doesn't allow a response for this particular dispute.
	DueBy int64 `json:"due_by"`
	// Whether evidence has been staged for this dispute.
	HasEvidence bool `json:"has_evidence"`
	// Whether the last evidence submission was submitted past the due date. Defaults to `false` if no evidence submissions have occurred. If `true`, then delivery of the latest evidence is *not* guaranteed.
	PastDue bool `json:"past_due"`
	// The number of times evidence has been submitted. Typically, you may only submit evidence once.
	SubmissionCount int64 `json:"submission_count"`
}
type DisputePaymentMethodDetailsCard struct {
	// Card brand. Can be `amex`, `diners`, `discover`, `eftpos_au`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Brand string `json:"brand"`
	// The card network's specific dispute reason code, which maps to one of Stripe's primary dispute categories to simplify response guidance. The [Network code map](https://stripe.com/docs/disputes/categories#network-code-map) lists all available dispute reason codes by network.
	NetworkReasonCode string `json:"network_reason_code"`
}
type DisputePaymentMethodDetailsKlarna struct {
	// The reason for the dispute as defined by Klarna
	ReasonCode string `json:"reason_code"`
}
type DisputePaymentMethodDetailsPaypal struct {
	// The ID of the dispute in PayPal.
	CaseID string `json:"case_id"`
	// The reason for the dispute as defined by PayPal
	ReasonCode string `json:"reason_code"`
}
type DisputePaymentMethodDetails struct {
	Card   *DisputePaymentMethodDetailsCard   `json:"card"`
	Klarna *DisputePaymentMethodDetailsKlarna `json:"klarna"`
	Paypal *DisputePaymentMethodDetailsPaypal `json:"paypal"`
	// Payment method type.
	Type DisputePaymentMethodDetailsType `json:"type"`
}

// A dispute occurs when a customer questions your charge with their card issuer.
// When this happens, you have the opportunity to respond to the dispute with
// evidence that shows that the charge is legitimate.
//
// Related guide: [Disputes and fraud](https://stripe.com/docs/disputes)
type Dispute struct {
	APIResource
	// Disputed amount. Usually the amount of the charge, but it can differ (usually because of currency fluctuation or because only part of the order is disputed).
	Amount int64 `json:"amount"`
	// List of zero, one, or two balance transactions that show funds withdrawn and reinstated to your Stripe account as a result of this dispute.
	BalanceTransactions []*BalanceTransaction `json:"balance_transactions"`
	// ID of the charge that's disputed.
	Charge *Charge `json:"charge"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency        Currency                `json:"currency"`
	Evidence        *DisputeEvidence        `json:"evidence"`
	EvidenceDetails *DisputeEvidenceDetails `json:"evidence_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If true, it's still possible to refund the disputed payment. After the payment has been fully refunded, no further funds are withdrawn from your Stripe account as a result of this dispute.
	IsChargeRefundable bool `json:"is_charge_refundable"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Network-dependent reason code for the dispute.
	NetworkReasonCode string `json:"network_reason_code"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// ID of the PaymentIntent that's disputed.
	PaymentIntent        *PaymentIntent               `json:"payment_intent"`
	PaymentMethodDetails *DisputePaymentMethodDetails `json:"payment_method_details"`
	// Reason given by cardholder for dispute. Possible values are `bank_cannot_process`, `check_returned`, `credit_not_processed`, `customer_initiated`, `debit_not_authorized`, `duplicate`, `fraudulent`, `general`, `incorrect_account_details`, `insufficient_funds`, `product_not_received`, `product_unacceptable`, `subscription_canceled`, or `unrecognized`. Learn more about [dispute reasons](https://stripe.com/docs/disputes/categories).
	Reason DisputeReason `json:"reason"`
	// Current status of dispute. Possible values are `warning_needs_response`, `warning_under_review`, `warning_closed`, `needs_response`, `under_review`, `won`, or `lost`.
	Status DisputeStatus `json:"status"`
}

// DisputeList is a list of Disputes as retrieved from a list endpoint.
type DisputeList struct {
	APIResource
	ListMeta
	Data []*Dispute `json:"data"`
}

// UnmarshalJSON handles deserialization of a Dispute.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (d *Dispute) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		d.ID = id
		return nil
	}

	type dispute Dispute
	var v dispute
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*d = Dispute(v)
	return nil
}
