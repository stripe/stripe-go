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

// v2NetworkBusinessProfileService is used to invoke businessprofile related APIs.
type v2NetworkBusinessProfileService struct {
	B   Backend
	Key string
}

// Retrieve a Stripe business profile by its Network ID.
func (c v2NetworkBusinessProfileService) Retrieve(ctx context.Context, id string, params *V2NetworkBusinessProfileRetrieveParams) (*V2NetworkBusinessProfile, error) {
	if params == nil {
		params = &V2NetworkBusinessProfileRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/network/business_profiles/%s", id)
	businessprofile := &V2NetworkBusinessProfile{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, businessprofile)
	return businessprofile, err
}

// Retrieve the Stripe business profile associated with the requesting merchant and livemode.
func (c v2NetworkBusinessProfileService) Me(ctx context.Context, params *V2NetworkBusinessProfileMeParams) (*V2NetworkBusinessProfile, error) {
	if params == nil {
		params = &V2NetworkBusinessProfileMeParams{}
	}
	params.Context = ctx
	businessprofile := &V2NetworkBusinessProfile{}
	err := c.B.Call(
		http.MethodGet, "/v2/network/business_profiles/me", c.Key, params, businessprofile)
	return businessprofile, err
}
