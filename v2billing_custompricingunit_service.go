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

// v2BillingCustomPricingUnitService is used to invoke custompricingunit related APIs.
type v2BillingCustomPricingUnitService struct {
	B   Backend
	Key string
}

// Create a CustomPricingUnit object.
func (c v2BillingCustomPricingUnitService) Create(ctx context.Context, params *V2BillingCustomPricingUnitCreateParams) (*V2BillingCustomPricingUnit, error) {
	if params == nil {
		params = &V2BillingCustomPricingUnitCreateParams{}
	}
	params.Context = ctx
	custompricingunit := &V2BillingCustomPricingUnit{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/custom_pricing_units", c.Key, params, custompricingunit)
	return custompricingunit, err
}

// Retrieve a CustomPricingUnit object.
func (c v2BillingCustomPricingUnitService) Retrieve(ctx context.Context, id string, params *V2BillingCustomPricingUnitRetrieveParams) (*V2BillingCustomPricingUnit, error) {
	if params == nil {
		params = &V2BillingCustomPricingUnitRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/custom_pricing_units/%s", id)
	custompricingunit := &V2BillingCustomPricingUnit{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, custompricingunit)
	return custompricingunit, err
}

// Update a CustomPricingUnit object.
func (c v2BillingCustomPricingUnitService) Update(ctx context.Context, id string, params *V2BillingCustomPricingUnitUpdateParams) (*V2BillingCustomPricingUnit, error) {
	if params == nil {
		params = &V2BillingCustomPricingUnitUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/custom_pricing_units/%s", id)
	custompricingunit := &V2BillingCustomPricingUnit{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, custompricingunit)
	return custompricingunit, err
}

// List all CustomPricingUnit objects.
func (c v2BillingCustomPricingUnitService) List(ctx context.Context, listParams *V2BillingCustomPricingUnitListParams) Seq2[*V2BillingCustomPricingUnit, error] {
	if listParams == nil {
		listParams = &V2BillingCustomPricingUnitListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/custom_pricing_units", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingCustomPricingUnit], error) {
		page := &V2Page[*V2BillingCustomPricingUnit]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
