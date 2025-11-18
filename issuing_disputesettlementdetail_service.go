//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1IssuingDisputeSettlementDetailService is used to invoke /v1/issuing/dispute_settlement_details APIs.
type v1IssuingDisputeSettlementDetailService struct {
	B   Backend
	Key string
}

// Retrieves an Issuing DisputeSettlementDetail object.
func (c v1IssuingDisputeSettlementDetailService) Retrieve(ctx context.Context, id string, params *IssuingDisputeSettlementDetailRetrieveParams) (*IssuingDisputeSettlementDetail, error) {
	if params == nil {
		params = &IssuingDisputeSettlementDetailRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/dispute_settlement_details/%s", id)
	disputesettlementdetail := &IssuingDisputeSettlementDetail{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, disputesettlementdetail)
	return disputesettlementdetail, err
}

// Returns a list of Issuing DisputeSettlementDetail objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingDisputeSettlementDetailService) List(ctx context.Context, listParams *IssuingDisputeSettlementDetailListParams) Seq2[*IssuingDisputeSettlementDetail, error] {
	if listParams == nil {
		listParams = &IssuingDisputeSettlementDetailListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*IssuingDisputeSettlementDetail], error) {
		list := &v1Page[*IssuingDisputeSettlementDetail]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/dispute_settlement_details", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
