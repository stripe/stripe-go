//
//
// File generated from our OpenAPI spec
//
//

// Package disputesettlementdetail provides the /issuing/dispute_settlement_details APIs
package disputesettlementdetail

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v79"
)

// Client is used to invoke /issuing/dispute_settlement_details APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an Issuing DisputeSettlementDetail object.
func Get(id string, params *stripe.IssuingDisputeSettlementDetailParams) (*stripe.IssuingDisputeSettlementDetail, error) {
	return getC().Get(id, params)
}

// Retrieves an Issuing DisputeSettlementDetail object.
func (c Client) Get(id string, params *stripe.IssuingDisputeSettlementDetailParams) (*stripe.IssuingDisputeSettlementDetail, error) {
	path := stripe.FormatURLPath("/v1/issuing/dispute_settlement_details/%s", id)
	disputesettlementdetail := &stripe.IssuingDisputeSettlementDetail{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, disputesettlementdetail)
	return disputesettlementdetail, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
