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

// Set the Rate for a Metered Item on the latest version of a Rate Card object. This will create a new Rate Card version
// if the Metered Item already has a rate on the Rate Card.
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

// Remove an existing Rate from a Rate Card. This will create a new Rate Card Version without that Rate.
func (c v2BillingRateCardsRateService) Delete(ctx context.Context, id string, params *V2BillingRateCardsRateDeleteParams) (*V2DeletedObject, error) {
	if params == nil {
		params = &V2BillingRateCardsRateDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/rates/%s", StringValue(params.RateCardID), id)
	deletedObj := &V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}

// List all Rates associated with a Rate Card for a specific version (defaults to latest). Rates remain active for all subsequent versions until a new rate is created for the same Metered Item.
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
