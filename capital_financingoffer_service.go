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

// v1CapitalFinancingOfferService is used to invoke /v1/capital/financing_offers APIs.
type v1CapitalFinancingOfferService struct {
	B   Backend
	Key string
}

// Get the details of the financing offer
func (c v1CapitalFinancingOfferService) Retrieve(ctx context.Context, id string, params *CapitalFinancingOfferRetrieveParams) (*CapitalFinancingOffer, error) {
	if params == nil {
		params = &CapitalFinancingOfferRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/capital/financing_offers/%s", id)
	financingoffer := &CapitalFinancingOffer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financingoffer)
	return financingoffer, err
}

// Acknowledges that platform has received and delivered the financing_offer to
// the intended merchant recipient.
func (c v1CapitalFinancingOfferService) MarkDelivered(ctx context.Context, id string, params *CapitalFinancingOfferMarkDeliveredParams) (*CapitalFinancingOffer, error) {
	if params == nil {
		params = &CapitalFinancingOfferMarkDeliveredParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/capital/financing_offers/%s/mark_delivered", id)
	financingoffer := &CapitalFinancingOffer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financingoffer)
	return financingoffer, err
}

// Retrieves the financing offers available for Connected accounts that belong to your platform.
func (c v1CapitalFinancingOfferService) List(ctx context.Context, listParams *CapitalFinancingOfferListParams) Seq2[*CapitalFinancingOffer, error] {
	if listParams == nil {
		listParams = &CapitalFinancingOfferListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*CapitalFinancingOffer], error) {
		list := &v1Page[*CapitalFinancingOffer]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/capital/financing_offers", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
