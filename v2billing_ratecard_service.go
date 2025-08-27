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

// v2BillingRateCardService is used to invoke ratecard related APIs.
type v2BillingRateCardService struct {
	B   Backend
	Key string
}

// Create a Rate Card object.
func (c v2BillingRateCardService) Create(ctx context.Context, params *V2BillingRateCardCreateParams) (*V2BillingRateCard, error) {
	if params == nil {
		params = &V2BillingRateCardCreateParams{}
	}
	params.Context = ctx
	ratecard := &V2BillingRateCard{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/rate_cards", c.Key, params, ratecard)
	return ratecard, err
}

// Retrieve the latest version of a Rate Card object.
func (c v2BillingRateCardService) Retrieve(ctx context.Context, id string, params *V2BillingRateCardRetrieveParams) (*V2BillingRateCard, error) {
	if params == nil {
		params = &V2BillingRateCardRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/rate_cards/%s", id)
	ratecard := &V2BillingRateCard{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, ratecard)
	return ratecard, err
}

// Update a Rate Card object.
func (c v2BillingRateCardService) Update(ctx context.Context, id string, params *V2BillingRateCardUpdateParams) (*V2BillingRateCard, error) {
	if params == nil {
		params = &V2BillingRateCardUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/rate_cards/%s", id)
	ratecard := &V2BillingRateCard{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, ratecard)
	return ratecard, err
}

// List all Rate Card objects.
func (c v2BillingRateCardService) List(ctx context.Context, listParams *V2BillingRateCardListParams) Seq2[*V2BillingRateCard, error] {
	if listParams == nil {
		listParams = &V2BillingRateCardListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/rate_cards", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingRateCard], error) {
		page := &V2Page[*V2BillingRateCard]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
