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

// v2BillingPricingPlanSubscriptionsComponentService is used to invoke component related APIs.
type v2BillingPricingPlanSubscriptionsComponentService struct {
	B   Backend
	Key string
}

// Retrieve a Pricing Plan Subscription's components.
func (c v2BillingPricingPlanSubscriptionsComponentService) Retrieve(ctx context.Context, id string, params *V2BillingPricingPlanSubscriptionsComponentRetrieveParams) (*V2BillingPricingPlanSubscriptionComponents, error) {
	if params == nil {
		params = &V2BillingPricingPlanSubscriptionsComponentRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/pricing_plan_subscriptions/%s/components", id)
	pricingplansubscriptioncomponents := &V2BillingPricingPlanSubscriptionComponents{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, pricingplansubscriptioncomponents)
	return pricingplansubscriptioncomponents, err
}
