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

// v1TopupService is used to invoke /v1/topups APIs.
type v1TopupService struct {
	B   Backend
	Key string
}

// Top up the balance of an account
func (c v1TopupService) Create(ctx context.Context, params *TopupCreateParams) (*Topup, error) {
	if params == nil {
		params = &TopupCreateParams{}
	}
	params.Context = ctx
	topup := &Topup{}
	err := c.B.Call(http.MethodPost, "/v1/topups", c.Key, params, topup)
	return topup, err
}

// Retrieves the details of a top-up that has previously been created. Supply the unique top-up ID that was returned from your previous request, and Stripe will return the corresponding top-up information.
func (c v1TopupService) Retrieve(ctx context.Context, id string, params *TopupRetrieveParams) (*Topup, error) {
	if params == nil {
		params = &TopupRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/topups/%s", id)
	topup := &Topup{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, topup)
	return topup, err
}

// Updates the metadata of a top-up. Other top-up details are not editable by design.
func (c v1TopupService) Update(ctx context.Context, id string, params *TopupUpdateParams) (*Topup, error) {
	if params == nil {
		params = &TopupUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/topups/%s", id)
	topup := &Topup{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, topup)
	return topup, err
}

// Cancels a top-up. Only pending top-ups can be canceled.
func (c v1TopupService) Cancel(ctx context.Context, id string, params *TopupCancelParams) (*Topup, error) {
	path := FormatURLPath("/v1/topups/%s/cancel", id)
	topup := &Topup{}
	if params == nil {
		params = &TopupCancelParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, topup)
	return topup, err
}

// Returns a list of top-ups.
func (c v1TopupService) List(ctx context.Context, listParams *TopupListParams) Seq2[*Topup, error] {
	if listParams == nil {
		listParams = &TopupListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Topup, ListContainer, error) {
		list := &TopupList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/topups", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
