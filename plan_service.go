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

// v1PlanService is used to invoke /v1/plans APIs.
type v1PlanService struct {
	B   Backend
	Key string
}

// You can now model subscriptions more flexibly using the [Prices API](https://docs.stripe.com/api#prices). It replaces the Plans API and is backwards compatible to simplify your migration.
func (c v1PlanService) Create(ctx context.Context, params *PlanCreateParams) (*Plan, error) {
	if params == nil {
		params = &PlanCreateParams{}
	}
	params.Context = ctx
	plan := &Plan{}
	err := c.B.Call(http.MethodPost, "/v1/plans", c.Key, params, plan)
	return plan, err
}

// Retrieves the plan with the given ID.
func (c v1PlanService) Retrieve(ctx context.Context, id string, params *PlanRetrieveParams) (*Plan, error) {
	if params == nil {
		params = &PlanRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/plans/%s", id)
	plan := &Plan{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, plan)
	return plan, err
}

// Updates the specified plan by setting the values of the parameters passed. Any parameters not provided are left unchanged. By design, you cannot change a plan's ID, amount, currency, or billing cycle.
func (c v1PlanService) Update(ctx context.Context, id string, params *PlanUpdateParams) (*Plan, error) {
	if params == nil {
		params = &PlanUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/plans/%s", id)
	plan := &Plan{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, plan)
	return plan, err
}

// Deleting plans means new subscribers can't be added. Existing subscribers aren't affected.
func (c v1PlanService) Delete(ctx context.Context, id string, params *PlanDeleteParams) (*Plan, error) {
	if params == nil {
		params = &PlanDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/plans/%s", id)
	plan := &Plan{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, plan)
	return plan, err
}

// Returns a list of your plans.
func (c v1PlanService) List(ctx context.Context, listParams *PlanListParams) Seq2[*Plan, error] {
	if listParams == nil {
		listParams = &PlanListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Plan, ListContainer, error) {
		list := &PlanList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/plans", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
