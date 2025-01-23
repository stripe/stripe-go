//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

// Reason the notice is being sent. The reason determines what copy the notice must contain. See the [regulated customer notices](https://stripe.com/docs/issuing/compliance-us/issuing-regulated-customer-notices) guide. All reasons might not apply to your integration, and Stripe might add new reasons in the future, so we recommend an internal warning when you receive an unknown reason.
type AccountNoticeReason string

// List of values that AccountNoticeReason can take
const (
	AccountNoticeReasonIssuingAccountClosedForInactivity                             AccountNoticeReason = "issuing.account_closed_for_inactivity"
	AccountNoticeReasonIssuingAccountClosedForNotProvidingBusinessModelClarification AccountNoticeReason = "issuing.account_closed_for_not_providing_business_model_clarification"
	AccountNoticeReasonIssuingAccountClosedForNotProvidingURLClarification           AccountNoticeReason = "issuing.account_closed_for_not_providing_url_clarification"
	AccountNoticeReasonIssuingAccountClosedForNotProvidingUseCaseClarification       AccountNoticeReason = "issuing.account_closed_for_not_providing_use_case_clarification"
	AccountNoticeReasonIssuingAccountClosedForTermsOfServiceViolation                AccountNoticeReason = "issuing.account_closed_for_terms_of_service_violation"
	AccountNoticeReasonIssuingApplicationRejectedForFailureToVerify                  AccountNoticeReason = "issuing.application_rejected_for_failure_to_verify"
	AccountNoticeReasonIssuingCreditApplicationRejected                              AccountNoticeReason = "issuing.credit_application_rejected"
	AccountNoticeReasonIssuingCreditIncreaseApplicationRejected                      AccountNoticeReason = "issuing.credit_increase_application_rejected"
	AccountNoticeReasonIssuingCreditLimitDecreased                                   AccountNoticeReason = "issuing.credit_limit_decreased"
	AccountNoticeReasonIssuingCreditLineClosed                                       AccountNoticeReason = "issuing.credit_line_closed"
	AccountNoticeReasonIssuingDisputeLost                                            AccountNoticeReason = "issuing.dispute_lost"
	AccountNoticeReasonIssuingDisputeSubmitted                                       AccountNoticeReason = "issuing.dispute_submitted"
	AccountNoticeReasonIssuingDisputeWon                                             AccountNoticeReason = "issuing.dispute_won"
)

// Retrieves a list of AccountNotice objects. The objects are sorted in descending order by creation date, with the most-recently-created object appearing first.
type AccountNoticeListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set to false to only return unsent AccountNotices.
	Sent *bool `form:"sent"`
}

// AddExpand appends a new field to expand.
func (p *AccountNoticeListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an AccountNotice object.
type AccountNoticeParams struct {
	Params `form:"*"`
	// Information about the email you sent.
	Email *AccountNoticeEmailParams `form:"email"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// Date when you sent the notice.
	SentAt *time.Time `form:"sent_at"`
}

// AddExpand appends a new field to expand.
func (p *AccountNoticeParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *AccountNoticeParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Information about the email you sent.
type AccountNoticeEmailParams struct {
	// Content of the email in plain text. The copy must match exactly the language that Stripe Compliance has approved for use.
	PlainText *string `form:"plain_text"`
	// Email address of the recipient.
	Recipient *string `form:"recipient"`
	// Subject of the email.
	Subject *string `form:"subject"`
}

// Information about the email when sent.
type AccountNoticeEmail struct {
	// Content of the email in plain text. The copy must match exactly the language that Stripe Compliance has approved for use.
	PlainText string `json:"plain_text"`
	// Email address of the recipient.
	Recipient string `json:"recipient"`
	// Subject of the email.
	Subject string `json:"subject"`
}

// Information about objects related to the notice.
type AccountNoticeLinkedObjects struct {
	// Associated [Capability](https://stripe.com/docs/api/capabilities)
	Capability string `json:"capability"`
	// Associated [Credit Underwriting Record](https://stripe.com/docs/api/issuing/credit_underwriting_record)
	IssuingCreditUnderwritingRecord string `json:"issuing_credit_underwriting_record"`
	// Associated [Issuing Dispute](https://stripe.com/docs/api/issuing/disputes)
	IssuingDispute string `json:"issuing_dispute"`
}

// A notice to a Connected account. Notice can be sent by Stripe on your behalf or you can opt to send the notices yourself.
//
// See the [guide to send notices](https://stripe.com/docs/issuing/compliance-us/issuing-regulated-customer-notices) to your connected accounts.
type AccountNotice struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created time.Time `json:"created"`
	// When present, the deadline for sending the notice to meet the relevant regulations.
	Deadline time.Time `json:"deadline"`
	// Information about the email when sent.
	Email *AccountNoticeEmail `json:"email"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Information about objects related to the notice.
	LinkedObjects *AccountNoticeLinkedObjects `json:"linked_objects"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Reason the notice is being sent. The reason determines what copy the notice must contain. See the [regulated customer notices](https://stripe.com/docs/issuing/compliance-us/issuing-regulated-customer-notices) guide. All reasons might not apply to your integration, and Stripe might add new reasons in the future, so we recommend an internal warning when you receive an unknown reason.
	Reason AccountNoticeReason `json:"reason"`
	// Date when the notice was sent. When absent, you must send the notice, update the content of the email and date when it was sent.
	SentAt time.Time `json:"sent_at"`
}

// AccountNoticeList is a list of AccountNotices as retrieved from a list endpoint.
type AccountNoticeList struct {
	APIResource
	ListMeta
	Data []*AccountNotice `json:"data"`
}

// UnmarshalJSON handles deserialization of an AccountNotice.
// This custom unmarshaling is needed to handle the time fields correctly.
func (a *AccountNotice) UnmarshalJSON(data []byte) error {
	type accountNotice AccountNotice
	v := struct {
		Created  int64 `json:"created"`
		Deadline int64 `json:"deadline"`
		SentAt   int64 `json:"sent_at"`
		*accountNotice
	}{
		accountNotice: (*accountNotice)(a),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	a.Created = time.Unix(v.Created, 0)
	a.Deadline = time.Unix(v.Deadline, 0)
	a.SentAt = time.Unix(v.SentAt, 0)
	return nil
}

// MarshalJSON handles serialization of an AccountNotice.
// This custom marshaling is needed to handle the time fields correctly.
func (a AccountNotice) MarshalJSON() ([]byte, error) {
	type accountNotice AccountNotice
	v := struct {
		Created  int64 `json:"created"`
		Deadline int64 `json:"deadline"`
		SentAt   int64 `json:"sent_at"`
		accountNotice
	}{
		accountNotice: (accountNotice)(a),
		Created:       a.Created.Unix(),
		Deadline:      a.Deadline.Unix(),
		SentAt:        a.SentAt.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
