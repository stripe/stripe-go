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

// v1PaymentPlanService is used to invoke /v1/payment_plans APIs.
type v1PaymentPlanService struct {
	B   Backend
	Key string
}

// Creates a payment plan that splits a single invoice obligation into installments with their own due dates and amounts.
func (c v1PaymentPlanService) Create(ctx context.Context, params *PaymentPlanCreateParams) (*PaymentPlan, error) {
	if params == nil {
		params = &PaymentPlanCreateParams{}
	}
	params.Context = ctx
	paymentplan := &PaymentPlan{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_plans", c.Key, params, paymentplan)
	return paymentplan, err
}

// Retrieves the payment plan with the given ID.
func (c v1PaymentPlanService) Retrieve(ctx context.Context, id string, params *PaymentPlanRetrieveParams) (*PaymentPlan, error) {
	if params == nil {
		params = &PaymentPlanRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_plans/%s", id)
	paymentplan := &PaymentPlan{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentplan)
	return paymentplan, err
}

// Updates the schedule or metadata of an existing payment plan. Only unpaid installments can be updated.
func (c v1PaymentPlanService) Update(ctx context.Context, id string, params *PaymentPlanUpdateParams) (*PaymentPlan, error) {
	if params == nil {
		params = &PaymentPlanUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_plans/%s", id)
	paymentplan := &PaymentPlan{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentplan)
	return paymentplan, err
}

// Returns a list of payment plans.
func (c v1PaymentPlanService) List(ctx context.Context, listParams *PaymentPlanListParams) *V1List[*PaymentPlan] {
	if listParams == nil {
		listParams = &PaymentPlanListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*PaymentPlan], error) {
		list := &v1Page[*PaymentPlan]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_plans", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
