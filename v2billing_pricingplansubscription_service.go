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

// v2BillingPricingPlanSubscriptionService is used to invoke pricingplansubscription related APIs.
type v2BillingPricingPlanSubscriptionService struct {
	B   Backend
	Key string
}

// Retrieve a Pricing Plan Subscription object.
func (c v2BillingPricingPlanSubscriptionService) Retrieve(ctx context.Context, id string, params *V2BillingPricingPlanSubscriptionRetrieveParams) (*V2BillingPricingPlanSubscription, error) {
	if params == nil {
		params = &V2BillingPricingPlanSubscriptionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/pricing_plan_subscriptions/%s", id)
	pricingplansubscription := &V2BillingPricingPlanSubscription{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, pricingplansubscription)
	return pricingplansubscription, err
}

// Update a Pricing Plan Subscription object.
func (c v2BillingPricingPlanSubscriptionService) Update(ctx context.Context, id string, params *V2BillingPricingPlanSubscriptionUpdateParams) (*V2BillingPricingPlanSubscription, error) {
	if params == nil {
		params = &V2BillingPricingPlanSubscriptionUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/pricing_plan_subscriptions/%s", id)
	pricingplansubscription := &V2BillingPricingPlanSubscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pricingplansubscription)
	return pricingplansubscription, err
}

// Remove Discounts from a Pricing Plan Subscription.
func (c v2BillingPricingPlanSubscriptionService) RemoveDiscounts(ctx context.Context, id string, params *V2BillingPricingPlanSubscriptionRemoveDiscountsParams) (*V2BillingPricingPlanSubscription, error) {
	if params == nil {
		params = &V2BillingPricingPlanSubscriptionRemoveDiscountsParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/pricing_plan_subscriptions/%s/remove_discounts", id)
	pricingplansubscription := &V2BillingPricingPlanSubscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, pricingplansubscription)
	return pricingplansubscription, err
}

// List all Pricing Plan Subscription objects.
func (c v2BillingPricingPlanSubscriptionService) List(ctx context.Context, listParams *V2BillingPricingPlanSubscriptionListParams) *V2List[*V2BillingPricingPlanSubscription] {
	if listParams == nil {
		listParams = &V2BillingPricingPlanSubscriptionListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/billing/pricing_plan_subscriptions", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2BillingPricingPlanSubscription], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2BillingPricingPlanSubscription]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
