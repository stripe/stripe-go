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

// v2IamAPIKeyService is used to invoke apikey related APIs.
type v2IamAPIKeyService struct {
	B   Backend
	Key string
}

// Create an API key.
func (c v2IamAPIKeyService) Create(ctx context.Context, params *V2IamAPIKeyCreateParams) (*V2IamAPIKey, error) {
	if params == nil {
		params = &V2IamAPIKeyCreateParams{}
	}
	params.Context = ctx
	apikey := &V2IamAPIKey{}
	err := c.B.Call(http.MethodPost, "/v2/iam/api_keys", c.Key, params, apikey)
	return apikey, err
}

// Retrieve an API key.
func (c v2IamAPIKeyService) Retrieve(ctx context.Context, id string, params *V2IamAPIKeyRetrieveParams) (*V2IamAPIKey, error) {
	if params == nil {
		params = &V2IamAPIKeyRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/iam/api_keys/%s", id)
	apikey := &V2IamAPIKey{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, apikey)
	return apikey, err
}

// Update an API key.
func (c v2IamAPIKeyService) Update(ctx context.Context, id string, params *V2IamAPIKeyUpdateParams) (*V2IamAPIKey, error) {
	if params == nil {
		params = &V2IamAPIKeyUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/iam/api_keys/%s", id)
	apikey := &V2IamAPIKey{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, apikey)
	return apikey, err
}

// Expire an API key.
func (c v2IamAPIKeyService) Expire(ctx context.Context, id string, params *V2IamAPIKeyExpireParams) (*V2IamAPIKey, error) {
	if params == nil {
		params = &V2IamAPIKeyExpireParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/iam/api_keys/%s/expire", id)
	apikey := &V2IamAPIKey{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, apikey)
	return apikey, err
}

// Rotate an API key.
func (c v2IamAPIKeyService) Rotate(ctx context.Context, id string, params *V2IamAPIKeyRotateParams) (*V2IamAPIKey, error) {
	if params == nil {
		params = &V2IamAPIKeyRotateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/iam/api_keys/%s/rotate", id)
	apikey := &V2IamAPIKey{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, apikey)
	return apikey, err
}

// List all API keys of an account.
func (c v2IamAPIKeyService) List(ctx context.Context, listParams *V2IamAPIKeyListParams) Seq2[*V2IamAPIKey, error] {
	if listParams == nil {
		listParams = &V2IamAPIKeyListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/iam/api_keys", listParams, func(path string, p ParamsContainer) (*V2Page[*V2IamAPIKey], error) {
		page := &V2Page[*V2IamAPIKey]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
