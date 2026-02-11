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

// v2BillingRateCardsCustomPricingUnitOverageRateService is used to invoke custompricingunitoveragerate related APIs.
type v2BillingRateCardsCustomPricingUnitOverageRateService struct {
	B   Backend
	Key string
}

// Create a Rate Card Custom Pricing Unit Overage Rate on a Rate Card.
func (c v2BillingRateCardsCustomPricingUnitOverageRateService) Create(ctx context.Context, params *V2BillingRateCardsCustomPricingUnitOverageRateCreateParams) (*V2BillingRateCardCustomPricingUnitOverageRate, error) {
	if params == nil {
		params = &V2BillingRateCardsCustomPricingUnitOverageRateCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/custom_pricing_unit_overage_rates", StringValue(
			params.RateCardID))
	ratecardcustompricingunitoveragerate := &V2BillingRateCardCustomPricingUnitOverageRate{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, ratecardcustompricingunitoveragerate)
	return ratecardcustompricingunitoveragerate, err
}

// Retrieve a Rate Card Custom Pricing Unit Overage Rate from a Rate Card.
func (c v2BillingRateCardsCustomPricingUnitOverageRateService) Retrieve(ctx context.Context, id string, params *V2BillingRateCardsCustomPricingUnitOverageRateRetrieveParams) (*V2BillingRateCardCustomPricingUnitOverageRate, error) {
	if params == nil {
		params = &V2BillingRateCardsCustomPricingUnitOverageRateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/custom_pricing_unit_overage_rates/%s", StringValue(
			params.RateCardID), id)
	ratecardcustompricingunitoveragerate := &V2BillingRateCardCustomPricingUnitOverageRate{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, ratecardcustompricingunitoveragerate)
	return ratecardcustompricingunitoveragerate, err
}

// Delete a Rate Card Custom Pricing Unit Overage Rate from a Rate Card.
func (c v2BillingRateCardsCustomPricingUnitOverageRateService) Delete(ctx context.Context, id string, params *V2BillingRateCardsCustomPricingUnitOverageRateDeleteParams) (*V2DeletedObject, error) {
	if params == nil {
		params = &V2BillingRateCardsCustomPricingUnitOverageRateDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/custom_pricing_unit_overage_rates/%s", StringValue(
			params.RateCardID), id)
	deletedObj := &V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}

// List all Rate Card Custom Pricing Unit Overage Rates on a Rate Card.
func (c v2BillingRateCardsCustomPricingUnitOverageRateService) List(ctx context.Context, listParams *V2BillingRateCardsCustomPricingUnitOverageRateListParams) Seq2[*V2BillingRateCardCustomPricingUnitOverageRate, error] {
	if listParams == nil {
		listParams = &V2BillingRateCardsCustomPricingUnitOverageRateListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/rate_cards/%s/custom_pricing_unit_overage_rates", StringValue(
			listParams.RateCardID))
	return NewV2List(path, listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingRateCardCustomPricingUnitOverageRate], error) {
		page := &V2Page[*V2BillingRateCardCustomPricingUnitOverageRate]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
