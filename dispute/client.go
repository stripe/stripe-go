// Package dispute provides the dispute-related APIs
package dispute

import (
	"fmt"
	"net/url"

	stripe "github.com/stripe/stripe-go"
)

const (
	Duplicate    stripe.DisputeReason = "duplicate"
	Fraudulent   stripe.DisputeReason = "fraudulent"
	SubCanceled  stripe.DisputeReason = "subscription_canceled"
	Unacceptable stripe.DisputeReason = "product_unacceptable"
	NotReceived  stripe.DisputeReason = "product_not_received"
	Unrecognized stripe.DisputeReason = "unrecognized"
	Credit       stripe.DisputeReason = "credit_not_processed"
	General      stripe.DisputeReason = "general"

	Won             stripe.DisputeStatus = "won"
	Lost            stripe.DisputeStatus = "lost"
	Response        stripe.DisputeStatus = "needs_response"
	Review          stripe.DisputeStatus = "under_review"
	WarningResponse stripe.DisputeStatus = "warning_needs_response"
	WarningReview   stripe.DisputeStatus = "warning_under_review"
	ChargeRefunded  stripe.DisputeStatus = "charge_refunded"
)

// Client is used to invoke dispute-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Update updates a charge's dispute.
// For more details see https://stripe.com/docs/api#update_dispute.
func Update(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.DisputeParams) (*stripe.Dispute, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}

		if params.Evidence != nil {
			params.Evidence.AppendDetails(body)
		}
		params.AppendTo(body)
	}

	dispute := &stripe.Dispute{}
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/dispute", id), c.Key, body, commonParams, dispute)

	return dispute, err
}

// Close dismisses a dispute in the customer's favor.
// For more details see https://stripe.com/docs/api#close_dispute.
func Close(id string) (*stripe.Dispute, error) {
	return getC().Close(id)
}

func (c Client) Close(id string) (*stripe.Dispute, error) {
	dispute := &stripe.Dispute{}
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/dispute/close", id), c.Key, nil, nil, dispute)

	return dispute, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
