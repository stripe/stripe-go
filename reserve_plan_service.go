//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
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
