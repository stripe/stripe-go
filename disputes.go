package stripe

import (
	"fmt"
	"net/url"
)

type DisputeReason string
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

	Won      DisputeStatus = "won"
	Lost     DisputeStatus = "lost"
	Response DisputeStatus = "needs_response"
	Review   DisputeStatus = "under_review"
)

type DisputeParams struct {
	Evidence string
}

type Dispute struct {
	Live     bool          `json:"livemode"`
	Amount   uint64        `json:"amount"`
	Currency Currency      `json:"currency"`
	Charge   string        `json:"charge"`
	Created  int64         `json:"created"`
	Reason   DisputeReason `json:"reason"`
	Status   DisputeStatus `json:"status"`
	Balance  string        `json:"balance_transaction"`
	Evidence string        `json:"evidence"`
	Due      int64         `json:"evidence_due_by"`
}

type DisputeClient struct {
	api   Api
	token string
}

func (c *DisputeClient) Update(id string, params *DisputeParams) (*Dispute, error) {
	body := &url.Values{}

	if len(params.Evidence) > 0 {
		body.Add("evidence", params.Evidence)
	}

	dispute := &Dispute{}
	err := c.api.Call("POST", fmt.Sprintf("/charges/%v/dispute", id), c.token, body, dispute)

	return dispute, err
}

func (c *DisputeClient) Close(id string) (*Dispute, error) {
	dispute := &Dispute{}
	err := c.api.Call("POST", fmt.Sprintf("/charges/%v/dispute/close", id), c.token, nil, dispute)

	return dispute, err
}
