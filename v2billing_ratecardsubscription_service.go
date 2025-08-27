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

// v2BillingRateCardSubscriptionService is used to invoke ratecardsubscription related APIs.
type v2BillingRateCardSubscriptionService struct {
	B   Backend
	Key string
}

// Create a Rate Card Subscription to bill a Rate Card on a specified Billing Cadence.
func (c v2BillingRateCardSubscriptionService) Create(ctx context.Context, params *V2BillingRateCardSubscriptionCreateParams) (*V2BillingRateCardSubscription, error) {
	if params == nil {
		params = &V2BillingRateCardSubscriptionCreateParams{}
	}
	params.Context = ctx
	ratecardsubscription := &V2BillingRateCardSubscription{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/rate_card_subscriptions", c.Key, params, ratecardsubscription)
	return ratecardsubscription, err
}

// Retrieve a Rate Card Subscription by ID.
func (c v2BillingRateCardSubscriptionService) Retrieve(ctx context.Context, id string, params *V2BillingRateCardSubscriptionRetrieveParams) (*V2BillingRateCardSubscription, error) {
	if params == nil {
		params = &V2BillingRateCardSubscriptionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/rate_card_subscriptions/%s", id)
	ratecardsubscription := &V2BillingRateCardSubscription{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, ratecardsubscription)
	return ratecardsubscription, err
}

// Update fields on an existing, active Rate Card Subscription.
func (c v2BillingRateCardSubscriptionService) Update(ctx context.Context, id string, params *V2BillingRateCardSubscriptionUpdateParams) (*V2BillingRateCardSubscription, error) {
	if params == nil {
		params = &V2BillingRateCardSubscriptionUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/rate_card_subscriptions/%s", id)
	ratecardsubscription := &V2BillingRateCardSubscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, ratecardsubscription)
	return ratecardsubscription, err
}

// Cancel an existing, active Rate Card Subscription.
func (c v2BillingRateCardSubscriptionService) Cancel(ctx context.Context, id string, params *V2BillingRateCardSubscriptionCancelParams) (*V2BillingRateCardSubscription, error) {
	if params == nil {
		params = &V2BillingRateCardSubscriptionCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/rate_card_subscriptions/%s/cancel", id)
	ratecardsubscription := &V2BillingRateCardSubscription{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, ratecardsubscription)
	return ratecardsubscription, err
}

// List all Rate Card Subscription objects.
func (c v2BillingRateCardSubscriptionService) List(ctx context.Context, listParams *V2BillingRateCardSubscriptionListParams) Seq2[*V2BillingRateCardSubscription, error] {
	if listParams == nil {
		listParams = &V2BillingRateCardSubscriptionListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/rate_card_subscriptions", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingRateCardSubscription], error) {
		page := &V2Page[*V2BillingRateCardSubscription]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
