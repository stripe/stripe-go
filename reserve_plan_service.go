//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v86/form"
)

// v1ReservePlanService is used to invoke /v1/reserve/plans APIs.
type v1ReservePlanService struct {
	B   Backend
	Key string
}

// Retrieve a ReservePlan.
func (c v1ReservePlanService) Retrieve(ctx context.Context, id string, params *ReservePlanRetrieveParams) (*ReservePlan, error) {
	if params == nil {
		params = &ReservePlanRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/reserve/plans/%s", id)
	plan := &ReservePlan{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, plan)
	return plan, err
}

// Returns a list of ReservePlans previously created. The ReservePlans are returned in sorted order, with the most recent ReservePlans appearing first.
func (c v1ReservePlanService) List(ctx context.Context, listParams *ReservePlanListParams) *V1List[*ReservePlan] {
	if listParams == nil {
		listParams = &ReservePlanListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*ReservePlan], error) {
		list := &v1Page[*ReservePlan]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/reserve/plans", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
