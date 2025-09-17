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

// v2BillingPricingPlansComponentService is used to invoke component related APIs.
type v2BillingPricingPlansComponentService struct {
	B   Backend
	Key string
}

// Create a Pricing Plan Component object.
func (c v2BillingPricingPlansComponentService) Create(ctx context.Context, params *V2BillingPricingPlansComponentCreateParams) (*V2BillingPricingPlanComponent, error) {
	if params == nil {
		params = &V2BillingPricingPlansComponentCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/pricing_plans/%s/components", StringValue(
			params.PricingPlanID))
	pricingplancomponent := &V2BillingPricingPlanComponent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pricingplancomponent)
	return pricingplancomponent, err
}

// Retrieve a Pricing Plan Component object.
func (c v2BillingPricingPlansComponentService) Retrieve(ctx context.Context, id string, params *V2BillingPricingPlansComponentRetrieveParams) (*V2BillingPricingPlanComponent, error) {
	if params == nil {
		params = &V2BillingPricingPlansComponentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/pricing_plans/%s/components/%s", StringValue(
			params.PricingPlanID), id)
	pricingplancomponent := &V2BillingPricingPlanComponent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, pricingplancomponent)
	return pricingplancomponent, err
}

// Update a Pricing Plan Component object.
func (c v2BillingPricingPlansComponentService) Update(ctx context.Context, id string, params *V2BillingPricingPlansComponentUpdateParams) (*V2BillingPricingPlanComponent, error) {
	if params == nil {
		params = &V2BillingPricingPlansComponentUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/pricing_plans/%s/components/%s", StringValue(
			params.PricingPlanID), id)
	pricingplancomponent := &V2BillingPricingPlanComponent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pricingplancomponent)
	return pricingplancomponent, err
}

// Remove a Pricing Plan Component from the latest version of a Pricing Plan.
func (c v2BillingPricingPlansComponentService) Delete(ctx context.Context, id string, params *V2BillingPricingPlansComponentDeleteParams) (*V2DeletedObject, error) {
	if params == nil {
		params = &V2BillingPricingPlansComponentDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/pricing_plans/%s/components/%s", StringValue(
			params.PricingPlanID), id)
	deletedObj := &V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}

// List all Pricing Plan Component objects for a Pricing Plan.
func (c v2BillingPricingPlansComponentService) List(ctx context.Context, listParams *V2BillingPricingPlansComponentListParams) Seq2[*V2BillingPricingPlanComponent, error] {
	if listParams == nil {
		listParams = &V2BillingPricingPlansComponentListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/pricing_plans/%s/components", StringValue(
			listParams.PricingPlanID))
	return NewV2List(path, listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingPricingPlanComponent], error) {
		page := &V2Page[*V2BillingPricingPlanComponent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
