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

// v2BillingRateCardsRateService is used to invoke rate related APIs.
type v2BillingRateCardsRateService struct {
	B   Backend
	Key string
}

// Set the rate for a MeteredItem on the latest version of a RateCard object. This will create a new RateCard version
// if the MeteredItem already has a rate on the RateCard.
func (c v2BillingRateCardsRateService) Create(ctx context.Context, params *V2BillingRateCardsRateCreateParams) (*V2BillingRateCardRate, error) {
	if params == nil {
		params = &V2BillingRateCardsRateCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/rates", StringValue(params.RateCardID))
	ratecardrate := &V2BillingRateCardRate{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, ratecardrate)
	return ratecardrate, err
}

// Retrieve a Rate object.
func (c v2BillingRateCardsRateService) Retrieve(ctx context.Context, id string, params *V2BillingRateCardsRateRetrieveParams) (*V2BillingRateCardRate, error) {
	if params == nil {
		params = &V2BillingRateCardsRateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/rates/%s", StringValue(params.RateCardID), id)
	ratecardrate := &V2BillingRateCardRate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, ratecardrate)
	return ratecardrate, err
}

// Remove an existing Rate from a RateCard. This will create a new RateCard version without that rate.
func (c v2BillingRateCardsRateService) Delete(ctx context.Context, id string, params *V2BillingRateCardsRateDeleteParams) (*V2BillingRateCardRate, error) {
	if params == nil {
		params = &V2BillingRateCardsRateDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/rates/%s", StringValue(params.RateCardID), id)
	ratecardrate := &V2BillingRateCardRate{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, ratecardrate)
	return ratecardrate, err
}

// List all Rates associated with a RateCard for a specific version (defaults to latest). Rates remain active for all subsequent versions until a new Rate is created for the same MeteredItem.
func (c v2BillingRateCardsRateService) List(ctx context.Context, listParams *V2BillingRateCardsRateListParams) Seq2[*V2BillingRateCardRate, error] {
	if listParams == nil {
		listParams = &V2BillingRateCardsRateListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/rates", StringValue(listParams.RateCardID))
	return NewV2List(path, listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingRateCardRate], error) {
		page := &V2Page[*V2BillingRateCardRate]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
