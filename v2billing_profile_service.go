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

// v2BillingProfileService is used to invoke profile related APIs.
type v2BillingProfileService struct {
	B   Backend
	Key string
}

// Create a BillingProfile object.
func (c v2BillingProfileService) Create(ctx context.Context, params *V2BillingProfileCreateParams) (*V2BillingProfile, error) {
	if params == nil {
		params = &V2BillingProfileCreateParams{}
	}
	params.Context = ctx
	profile := &V2BillingProfile{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/profiles", c.Key, params, profile)
	return profile, err
}

// Retrieve a BillingProfile object.
func (c v2BillingProfileService) Retrieve(ctx context.Context, id string, params *V2BillingProfileRetrieveParams) (*V2BillingProfile, error) {
	if params == nil {
		params = &V2BillingProfileRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/profiles/%s", id)
	profile := &V2BillingProfile{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, profile)
	return profile, err
}

// Update a BillingProfile object.
func (c v2BillingProfileService) Update(ctx context.Context, id string, params *V2BillingProfileUpdateParams) (*V2BillingProfile, error) {
	if params == nil {
		params = &V2BillingProfileUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/profiles/%s", id)
	profile := &V2BillingProfile{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, profile)
	return profile, err
}

// List Billing Profiles.
func (c v2BillingProfileService) List(ctx context.Context, listParams *V2BillingProfileListParams) Seq2[*V2BillingProfile, error] {
	if listParams == nil {
		listParams = &V2BillingProfileListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/profiles", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingProfile], error) {
		page := &V2Page[*V2BillingProfile]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
