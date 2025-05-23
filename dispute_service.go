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

// v1DisputeService is used to invoke /v1/disputes APIs.
type v1DisputeService struct {
	B   Backend
	Key string
}

// Retrieves the dispute with the given ID.
func (c v1DisputeService) Retrieve(ctx context.Context, id string, params *DisputeRetrieveParams) (*Dispute, error) {
	if params == nil {
		params = &DisputeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/disputes/%s", id)
	dispute := &Dispute{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, dispute)
	return dispute, err
}

// When you get a dispute, contacting your customer is always the best first step. If that doesn't work, you can submit evidence to help us resolve the dispute in your favor. You can do this in your [dashboard](https://dashboard.stripe.com/disputes), but if you prefer, you can use the API to submit evidence programmatically.
//
// Depending on your dispute type, different evidence fields will give you a better chance of winning your dispute. To figure out which evidence fields to provide, see our [guide to dispute types](https://docs.stripe.com/docs/disputes/categories).
func (c v1DisputeService) Update(ctx context.Context, id string, params *DisputeUpdateParams) (*Dispute, error) {
	if params == nil {
		params = &DisputeUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/disputes/%s", id)
	dispute := &Dispute{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, dispute)
	return dispute, err
}

// Closing the dispute for a charge indicates that you do not have any evidence to submit and are essentially dismissing the dispute, acknowledging it as lost.
//
// The status of the dispute will change from needs_response to lost. Closing a dispute is irreversible.
func (c v1DisputeService) Close(ctx context.Context, id string, params *DisputeCloseParams) (*Dispute, error) {
	path := FormatURLPath("/v1/disputes/%s/close", id)
	dispute := &Dispute{}
	if params == nil {
		params = &DisputeCloseParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, dispute)
	return dispute, err
}

// Returns a list of your disputes.
func (c v1DisputeService) List(ctx context.Context, listParams *DisputeListParams) Seq2[*Dispute, error] {
	if listParams == nil {
		listParams = &DisputeListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Dispute, ListContainer, error) {
		list := &DisputeList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/disputes", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
