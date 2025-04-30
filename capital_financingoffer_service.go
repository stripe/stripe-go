//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1CapitalFinancingOfferService is used to invoke /v1/capital/financing_offers APIs.
type v1CapitalFinancingOfferService struct {
	B   Backend
	Key string
}

// Get the details of the financing offer
func (c v1CapitalFinancingOfferService) Retrieve(ctx context.Context, id string, params *CapitalFinancingOfferRetrieveParams) (*CapitalFinancingOffer, error) {
	path := FormatURLPath("/v1/capital/financing_offers/%s", id)
	financingoffer := &CapitalFinancingOffer{}
	if params == nil {
		params = &CapitalFinancingOfferRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, financingoffer)
	return financingoffer, err
}

// Acknowledges that platform has received and delivered the financing_offer to
// the intended merchant recipient.
func (c v1CapitalFinancingOfferService) MarkDelivered(ctx context.Context, id string, params *CapitalFinancingOfferMarkDeliveredParams) (*CapitalFinancingOffer, error) {
	path := FormatURLPath("/v1/capital/financing_offers/%s/mark_delivered", id)
	financingoffer := &CapitalFinancingOffer{}
	if params == nil {
		params = &CapitalFinancingOfferMarkDeliveredParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, financingoffer)
	return financingoffer, err
}

// Retrieves the financing offers available for Connected accounts that belong to your platform.
func (c v1CapitalFinancingOfferService) List(ctx context.Context, listParams *CapitalFinancingOfferListParams) Seq2[*CapitalFinancingOffer, error] {
	if listParams == nil {
		listParams = &CapitalFinancingOfferListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*CapitalFinancingOffer, ListContainer, error) {
		list := &CapitalFinancingOfferList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/capital/financing_offers", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
