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
	NoSubmit bool                   `form:"submit,invert"`
}

// DisputeEvidenceParams is the set of parameters that can be used when submitting
// evidence for disputes.
type DisputeEvidenceParams struct {
	ActivityLog string `form:"access_activity_log"`

	BillingAddress string `form:"billing_address"`

	CancellationPolicy           string `form:"cancellation_policy"`
	CancellationPolicyDisclsoure string `form:"cancellation_policy_disclosure"`
	CancellationRebuttal         string `form:"cancellation_rebuttal"`

	CustomerName  string `form:"customer_name"`
	CustomerEmail string `form:"customer_email_address"`
	CustomerIP    string `form:"customer_purchase_ip"`
	CustomerSig   string `form:"customer_signature"`
	CustomerComm  string `form:"customer_communication"`

	DuplicateCharge       string `form:"duplicate_charge_id"`
	DuplicateChargeReason string `form:"duplicate_charge_explanation"`
	DuplicateChargeDoc    string `form:"duplicate_charge_documentation"`

	ProductDesc string `form:"product_description"`

	Receipt string `form:"receipt"`

	RefundPolicy           string `form:"refund_policy"`
	RefundPolicyDisclosure string `form:"refund_policy_disclosure"`
	RefundRefusalReason    string `form:"refund_refusal_explanation"`

	ServiceDate string `form:"service_date"`
	ServiceDoc  string `form:"service_documentation"`

	ShippingAddress  string `form:"shipping_address"`
	ShippingCarrier  string `form:"shipping_carrier"`
	ShippingDate     string `form:"shipping_date"`
	ShippingTracking string `form:"shipping_tracking_number"`
	ShippingDoc      string `form:"shipping_documentation"`

	UncategorizedText string `form:"uncategorized_text"`
	UncategorizedFile string `form:"uncategorized_file"`
}

// DisputeListParams is the set of parameters that can be used when listing disputes.
// For more details see https://stripe.com/docs/api#list_disputes.
type DisputeListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// Dispute is the resource representing a Stripe dispute.
// For more details see https://stripe.com/docs/api#disputes.
type Dispute struct {
	ID              string            `json:"id"`
	Live            bool              `json:"livemode"`
	Amount          uint64            `json:"amount"`
	Currency        Currency          `json:"currency"`
	Charge          string            `json:"charge"`
	Created         int64             `json:"created"`
	Refundable      bool              `json:"is_charge_refundable"`
	Reason          DisputeReason     `json:"reason"`
	Status          DisputeStatus     `json:"status"`
	Transactions    []*Transaction    `json:"balance_transactions"`
	Evidence        *DisputeEvidence  `json:"evidence"`
	EvidenceDetails *EvidenceDetails  `json:"evidence_details"`
	Meta            map[string]string `json:"metadata"`
}

// DisputeList is a list of disputes as retrieved from a list endpoint.
type DisputeList struct {
	ListMeta
	Values []*Dispute `json:"data"`
}

// EvidenceDetails is the structure representing more details about
// the dispute.
type EvidenceDetails struct {
	Count       int   `json:"submission_count"`
	DueDate     int64 `json:"due_by"`
	HasEvidence bool  `json:"has_evidence"`
	PastDue     bool  `json:"past_due"`
}

// DisputeEvidence is the structure that contains various details about
// the evidence submitted for the dispute.
// Almost all fields are strings since there structures (i.e. address)
// do not typically get parsed by anyone and are thus presented as-received.
type DisputeEvidence struct {
	ProductDesc                  string `json:"product_description"`
	CustomerName                 string `json:"customer_name"`
	CustomerEmail                string `json:"customer_email_address"`
	CustomerIP                   string `json:"customer_purchase_ip"`
	CustomerSig                  *File  `json:"customer_signature"`
	BillingAddress               string `json:"billing_address"`
	Receipt                      *File  `json:"receipt"`
	ShippingAddress              string `json:"shipping_address"`
	ShippingDate                 string `json:"shipping_date"`
	ShippingTracking             string `json:"shipping_tracking_number"`
	ShippingCarrier              string `json:"shipping_carrier"`
	ShippingDoc                  *File  `json:"shipping_documentation"`
	RefundPolicy                 *File  `json:"refund_policy"`
	RefundPolicyDisclosure       string `json:"refund_policy_disclosure"`
	RefundRefusalReason          string `json:"refund_refusal_explanation"`
	CancellationPolicy           *File  `json:"cancellation_policy"`
	CancellationPolicyDisclosure string `json:"cancellation_policy_disclosure"`
	CancellationRebuttal         string `json:"cancellation_rebuttal"`
	ActivityLog                  string `json:"access_activity_log"`
	ServiceDate                  string `json:"service_date"`
	ServiceDoc                   *File  `json:"service_documentation"`
	DuplicateCharge              string `json:"duplicate_charge_id"`
	DuplicateChargeReason        string `json:"duplicate_charge_explanation"`
	DuplicateChargeDoc           *File  `json:"duplicate_charge_documentation"`
	CustomerComm                 *File  `json:"customer_communication"`
	UncategorizedText            string `json:"uncategorized_text"`
	UncategorizedFile            *File  `json:"uncategorized_file"`
}

// File represents a link to downloadable content.
type File struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	Size    int    `json:"size"`
	Purpose string `json:"purpose"`
	URL     string `json:"url"`
	Mime    string `json:"mime_type"`
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
