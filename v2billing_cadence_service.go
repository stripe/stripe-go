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

// v2BillingCadenceService is used to invoke cadence related APIs.
type v2BillingCadenceService struct {
	B   Backend
	Key string
}

// Create a Billing Cadence object.
func (c v2BillingCadenceService) Create(ctx context.Context, params *V2BillingCadenceCreateParams) (*V2BillingCadence, error) {
	if params == nil {
		params = &V2BillingCadenceCreateParams{}
	}
	params.Context = ctx
	cadence := &V2BillingCadence{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/cadences", c.Key, params, cadence)
	return cadence, err
}

// Retrieve a Billing Cadence object.
func (c v2BillingCadenceService) Retrieve(ctx context.Context, id string, params *V2BillingCadenceRetrieveParams) (*V2BillingCadence, error) {
	if params == nil {
		params = &V2BillingCadenceRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/cadences/%s", id)
	cadence := &V2BillingCadence{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, cadence)
	return cadence, err
}

// Update a Billing Cadence object.
func (c v2BillingCadenceService) Update(ctx context.Context, id string, params *V2BillingCadenceUpdateParams) (*V2BillingCadence, error) {
	if params == nil {
		params = &V2BillingCadenceUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/cadences/%s", id)
	cadence := &V2BillingCadence{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cadence)
	return cadence, err
}

// Cancel the Billing Cadence.
func (c v2BillingCadenceService) Cancel(ctx context.Context, id string, params *V2BillingCadenceCancelParams) (*V2BillingCadence, error) {
	if params == nil {
		params = &V2BillingCadenceCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/cadences/%s/cancel", id)
	cadence := &V2BillingCadence{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, cadence)
	return cadence, err
}

// List Billing Cadences.
func (c v2BillingCadenceService) List(ctx context.Context, listParams *V2BillingCadenceListParams) Seq2[*V2BillingCadence, error] {
	if listParams == nil {
		listParams = &V2BillingCadenceListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/cadences", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingCadence], error) {
		page := &V2Page[*V2BillingCadence]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
