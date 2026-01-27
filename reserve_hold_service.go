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

// v1ReserveHoldService is used to invoke /v1/reserve/holds APIs.
type v1ReserveHoldService struct {
	B   Backend
	Key string
}

// Retrieve a ReserveHold.
func (c v1ReserveHoldService) Retrieve(ctx context.Context, id string, params *ReserveHoldRetrieveParams) (*ReserveHold, error) {
	if params == nil {
		params = &ReserveHoldRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/reserve/holds/%s", id)
	hold := &ReserveHold{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, hold)
	return hold, err
}

// Returns a list of ReserveHolds previously created. The ReserveHolds are returned in sorted order, with the most recent ReserveHolds appearing first.
func (c v1ReserveHoldService) List(ctx context.Context, listParams *ReserveHoldListParams) Seq2[*ReserveHold, error] {
	if listParams == nil {
		listParams = &ReserveHoldListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*ReserveHold], error) {
		list := &v1Page[*ReserveHold]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/reserve/holds", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
