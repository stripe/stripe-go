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

// v2BillingPricingPlansVersionService is used to invoke version related APIs.
type v2BillingPricingPlansVersionService struct {
	B   Backend
	Key string
}

// Retrieve a specific version of a PricingPlan.
func (c v2BillingPricingPlansVersionService) Retrieve(ctx context.Context, id string, params *V2BillingPricingPlansVersionRetrieveParams) (*V2BillingPricingPlanVersion, error) {
	if params == nil {
		params = &V2BillingPricingPlansVersionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/pricing_plans/%s/versions/%s", StringValue(
			params.PricingPlanID), id)
	pricingplanversion := &V2BillingPricingPlanVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, pricingplanversion)
	return pricingplanversion, err
}

// List all versions of a PricingPlan.
func (c v2BillingPricingPlansVersionService) List(ctx context.Context, listParams *V2BillingPricingPlansVersionListParams) Seq2[*V2BillingPricingPlanVersion, error] {
	if listParams == nil {
		listParams = &V2BillingPricingPlansVersionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/pricing_plans/%s/versions", StringValue(
			listParams.PricingPlanID))
	return NewV2List(path, listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingPricingPlanVersion], error) {
		page := &V2Page[*V2BillingPricingPlanVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
