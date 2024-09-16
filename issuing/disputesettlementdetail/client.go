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
	"github.com/stripe/stripe-go/v79/form"
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

// Returns a list of Issuing DisputeSettlementDetail objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func List(params *stripe.IssuingDisputeSettlementDetailListParams) *Iter {
	return getC().List(params)
}

// Returns a list of Issuing DisputeSettlementDetail objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c Client) List(listParams *stripe.IssuingDisputeSettlementDetailListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IssuingDisputeSettlementDetailList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/issuing/dispute_settlement_details", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for issuing dispute settlement details.
type Iter struct {
	*stripe.Iter
}

// IssuingDisputeSettlementDetail returns the issuing dispute settlement detail which the iterator is currently pointing to.
func (i *Iter) IssuingDisputeSettlementDetail() *stripe.IssuingDisputeSettlementDetail {
	return i.Current().(*stripe.IssuingDisputeSettlementDetail)
}

// IssuingDisputeSettlementDetailList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IssuingDisputeSettlementDetailList() *stripe.IssuingDisputeSettlementDetailList {
	return i.List().(*stripe.IssuingDisputeSettlementDetailList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
