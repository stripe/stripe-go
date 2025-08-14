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

// v2BillingPricingPlanService is used to invoke pricingplan related APIs.
type v2BillingPricingPlanService struct {
	B   Backend
	Key string
}

// Create a PricingPlan object.
func (c v2BillingPricingPlanService) Create(ctx context.Context, params *V2BillingPricingPlanCreateParams) (*V2BillingPricingPlan, error) {
	if params == nil {
		params = &V2BillingPricingPlanCreateParams{}
	}
	params.Context = ctx
	pricingplan := &V2BillingPricingPlan{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/pricing_plans", c.Key, params, pricingplan)
	return pricingplan, err
}

// Retrieve a PricingPlan object.
func (c v2BillingPricingPlanService) Retrieve(ctx context.Context, id string, params *V2BillingPricingPlanRetrieveParams) (*V2BillingPricingPlan, error) {
	if params == nil {
		params = &V2BillingPricingPlanRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/pricing_plans/%s", id)
	pricingplan := &V2BillingPricingPlan{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, pricingplan)
	return pricingplan, err
}

// Update a PricingPlan object.
func (c v2BillingPricingPlanService) Update(ctx context.Context, id string, params *V2BillingPricingPlanUpdateParams) (*V2BillingPricingPlan, error) {
	if params == nil {
		params = &V2BillingPricingPlanUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/pricing_plans/%s", id)
	pricingplan := &V2BillingPricingPlan{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pricingplan)
	return pricingplan, err
}

// List all PricingPlan objects.
func (c v2BillingPricingPlanService) List(ctx context.Context, listParams *V2BillingPricingPlanListParams) Seq2[*V2BillingPricingPlan, error] {
	if listParams == nil {
		listParams = &V2BillingPricingPlanListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/pricing_plans", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingPricingPlan], error) {
		page := &V2Page[*V2BillingPricingPlan]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
