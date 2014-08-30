package stripe

import (
	"fmt"
	"net/url"
)

// DisputeReason is the list of allowed values for a discount's reason.
// Allowed values are "duplicate", "fraudulent", "subscription_canceled",
// "product_unacceptable", "product_not_received", "unrecognized",
// "credit_not_processed", "general".
type DisputeReason string

// DisputeStatus is the list of allowed values for a discount's status.
// Allowed values are "won", "lost", "needs_ressponse", "under_review",
// "warning_needs_response", "warning_under_review", "charge_refunded".
type DisputeStatus string

const (
	Duplicate    DisputeReason = "duplicate"
	Fraudulent   DisputeReason = "fraudulent"
	SubCanceled  DisputeReason = "subscription_canceled"
	Unacceptable DisputeReason = "product_unacceptable"
	NotReceived  DisputeReason = "product_not_received"
	Unrecognized DisputeReason = "unrecognized"
	Credit       DisputeReason = "credit_not_processed"
	General      DisputeReason = "general"

	Won             DisputeStatus = "won"
	Lost            DisputeStatus = "lost"
	Response        DisputeStatus = "needs_response"
	Review          DisputeStatus = "under_review"
	WarningResponse DisputeStatus = "warning_needs_response"
	WarningReview   DisputeStatus = "warning_under_review"
	ChargeRefunded  DisputeStatus = "charge_refunded"
)

// DisputeParams is the set of parameters that can be used when updating a dispute.
// For more details see https://stripe.com/docs/api#update_dispute.
type DisputeParams struct {
	Params
	Evidence string
}

// Dispute is the resource representing a Stripe dispute.
// For more details see https://stripe.com/docs/api#disputes.
type Dispute struct {
	Live         bool              `json:"livemode"`
	Amount       uint64            `json:"amount"`
	Currency     Currency          `json:"currency"`
	Charge       string            `json:"charge"`
	Created      int64             `json:"created"`
	Reason       DisputeReason     `json:"reason"`
	Status       DisputeStatus     `json:"status"`
	Transactions []*Transaction    `json:"balance_transactions"`
	Evidence     string            `json:"evidence"`
	DueDate      int64             `json:"evidence_due_by"`
	Meta         map[string]string `json:"metadata"`
}

// DisputeClient is the client used to invoke dispute-related APIs.
type DisputeClient struct {
	api   Api
	token string
}

// Update updates a charge's dispute.
// For more details see https://stripe.com/docs/api#update_dispute.
func (c *DisputeClient) Update(id string, params *DisputeParams) (*Dispute, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Evidence) > 0 {
			body.Add("evidence", params.Evidence)
		}

		for k, v := range params.Meta {
			body.Add(fmt.Sprintf("metadata[%v]", k), v)
		}
	}

	dispute := &Dispute{}
	err := c.api.Call("POST", fmt.Sprintf("/charges/%v/dispute", id), c.token, body, dispute)

	return dispute, err
}

// Close dismisses a dispute in the customer's favor.
// For more details see https://stripe.com/docs/api#close_dispute.
func (c *DisputeClient) Close(id string) (*Dispute, error) {
	dispute := &Dispute{}
	err := c.api.Call("POST", fmt.Sprintf("/charges/%v/dispute/close", id), c.token, nil, dispute)

	return dispute, err
}
