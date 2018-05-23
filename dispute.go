package stripe

import (
	"encoding/json"
)

// DisputeReason is the list of allowed values for a discount's reason.
// Allowed values are "duplicate", "fraudulent", "subscription_canceled",
// "product_unacceptable", "product_not_received", "unrecognized",
// "credit_not_processed", "general".
type DisputeReason string

// DisputeStatus is the list of allowed values for a discount's status.
// Allowed values are "won", "lost", "needs_response", "under_review",
// "warning_needs_response", "warning_under_review", "charge_refunded",
// "warning_closed".
type DisputeStatus string

// DisputeParams is the set of parameters that can be used when updating a dispute.
// For more details see https://stripe.com/docs/api#update_dispute.
type DisputeParams struct {
	Params   `form:"*"`
	Evidence *DisputeEvidenceParams `form:"evidence"`
	Submit   *bool                  `form:"submit"`
}

// DisputeEvidenceParams is the set of parameters that can be used when submitting
// evidence for disputes.
type DisputeEvidenceParams struct {
	AccessActivityLog            *string `form:"access_activity_log"`
	BillingAddress               *string `form:"billing_address"`
	CancellationPolicy           *string `form:"cancellation_policy"`
	CancellationPolicyDisclosure *string `form:"cancellation_policy_disclosure"`
	CancellationRebuttal         *string `form:"cancellation_rebuttal"`
	CustomerCommunication        *string `form:"customer_communication"`
	CustomerEmailAddress         *string `form:"customer_email_address"`
	CustomerName                 *string `form:"customer_name"`
	CustomerPurchaseIP           *string `form:"customer_purchase_ip"`
	CustomerSignature            *string `form:"customer_signature"`
	DuplicateChargeDocumentation *string `form:"duplicate_charge_documentation"`
	DuplicateChargeExplanation   *string `form:"duplicate_charge_explanation"`
	DuplicateChargeID            *string `form:"duplicate_charge_id"`
	ProductDescription           *string `form:"product_description"`
	Receipt                      *string `form:"receipt"`
	RefundPolicy                 *string `form:"refund_policy"`
	RefundPolicyDisclosure       *string `form:"refund_policy_disclosure"`
	RefundRefusalExplanation     *string `form:"refund_refusal_explanation"`
	ServiceDate                  *string `form:"service_date"`
	ServiceDocumentation         *string `form:"service_documentation"`
	ShippingAddress              *string `form:"shipping_address"`
	ShippingCarrier              *string `form:"shipping_carrier"`
	ShippingDate                 *string `form:"shipping_date"`
	ShippingDocumentation        *string `form:"shipping_documentation"`
	ShippingTrackingNumber       *string `form:"shipping_tracking_number"`
	UncategorizedFile            *string `form:"uncategorized_file"`
	UncategorizedText            *string `form:"uncategorized_text"`
}

// DisputeListParams is the set of parameters that can be used when listing disputes.
// For more details see https://stripe.com/docs/api#list_disputes.
type DisputeListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// Dispute is the resource representing a Stripe dispute.
// For more details see https://stripe.com/docs/api#disputes.
type Dispute struct {
	Amount              int64                 `json:"amount"`
	BalanceTransactions []*BalanceTransaction `json:"balance_transactions"`
	Charge              *Charge               `json:"charge"`
	Created             int64                 `json:"created"`
	Currency            Currency              `json:"currency"`
	Evidence            *DisputeEvidence      `json:"evidence"`
	EvidenceDetails     *EvidenceDetails      `json:"evidence_details"`
	ID                  string                `json:"id"`
	IsChargeRefundable  bool                  `json:"is_charge_refundable"`
	Livemode            bool                  `json:"livemode"`
	Metadata            map[string]string     `json:"metadata"`
	Reason              DisputeReason         `json:"reason"`
	Status              DisputeStatus         `json:"status"`
}

// DisputeList is a list of disputes as retrieved from a list endpoint.
type DisputeList struct {
	ListMeta
	Data []*Dispute `json:"data"`
}

// EvidenceDetails is the structure representing more details about
// the dispute.
type EvidenceDetails struct {
	DueBy           int64 `json:"due_by"`
	HasEvidence     bool  `json:"has_evidence"`
	PastDue         bool  `json:"past_due"`
	SubmissionCount int64 `json:"submission_count"`
}

// DisputeEvidence is the structure that contains various details about
// the evidence submitted for the dispute.
// Almost all fields are strings since there structures (i.e. address)
// do not typically get parsed by anyone and are thus presented as-received.
type DisputeEvidence struct {
	AccessActivityLog            string `json:"access_activity_log"`
	BillingAddress               string `json:"billing_address"`
	CancellationPolicy           *File  `json:"cancellation_policy"`
	CancellationPolicyDisclosure string `json:"cancellation_policy_disclosure"`
	CancellationRebuttal         string `json:"cancellation_rebuttal"`
	CustomerCommunication        *File  `json:"customer_communication"`
	CustomerEmailAddress         string `json:"customer_email_address"`
	CustomerName                 string `json:"customer_name"`
	CustomerPurchaseIP           string `json:"customer_purchase_ip"`
	CustomerSignature            *File  `json:"customer_signature"`
	DuplicateChargeDocumentation *File  `json:"duplicate_charge_documentation"`
	DuplicateChargeExplanation   string `json:"duplicate_charge_explanation"`
	DuplicateChargeID            string `json:"duplicate_charge_id"`
	ProductDescription           string `json:"product_description"`
	Receipt                      *File  `json:"receipt"`
	RefundPolicy                 *File  `json:"refund_policy"`
	RefundPolicyDisclosure       string `json:"refund_policy_disclosure"`
	RefundRefusalExplanation     string `json:"refund_refusal_explanation"`
	ServiceDate                  string `json:"service_date"`
	ServiceDocumentation         *File  `json:"service_documentation"`
	ShippingAddress              string `json:"shipping_address"`
	ShippingCarrier              string `json:"shipping_carrier"`
	ShippingDate                 string `json:"shipping_date"`
	ShippingDocumentation        *File  `json:"shipping_documentation"`
	ShippingTrackingNumber       string `json:"shipping_tracking_number"`
	UncategorizedFile            *File  `json:"uncategorized_file"`
	UncategorizedText            string `json:"uncategorized_text"`
}

// File represents a link to downloadable content.
type File struct {
	Created  int64  `json:"created"`
	ID       string `json:"id"`
	MIMEType string `json:"mime_type"`
	Purpose  string `json:"purpose"`
	Size     int64  `json:"size"`
	URL      string `json:"url"`
}

// UnmarshalJSON handles deserialization of a Dispute.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *Dispute) UnmarshalJSON(data []byte) error {
	type dispute Dispute
	var dd dispute
	err := json.Unmarshal(data, &dd)
	if err == nil {
		*t = Dispute(dd)
	} else {
		// the id is surrounded by "\" characters, so strip them
		t.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// UnmarshalJSON handles deserialization of a File.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *File) UnmarshalJSON(data []byte) error {
	type file File
	var ff file
	err := json.Unmarshal(data, &ff)
	if err == nil {
		*f = File(ff)
	} else {
		// the id is surrounded by "\" characters, so strip them
		f.ID = string(data[1 : len(data)-1])
	}

	return nil
}
