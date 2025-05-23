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

// v1EphemeralKeyService is used to invoke /v1/ephemeral_keys APIs.
type v1EphemeralKeyService struct {
	B   Backend
	Key string
}

// Creates a short-lived API key for a given resource.
func (c v1EphemeralKeyService) Create(ctx context.Context, params *EphemeralKeyCreateParams) (*EphemeralKey, error) {
	if params == nil {
		params = &EphemeralKeyCreateParams{}
	}
	params.Context = ctx
	ephemeralkey := &EphemeralKey{}
	err := c.B.Call(
		http.MethodPost, "/v1/ephemeral_keys", c.Key, params, ephemeralkey)
	return ephemeralkey, err
}

// Invalidates a short-lived API key for a given resource.
func (c v1EphemeralKeyService) Delete(ctx context.Context, id string, params *EphemeralKeyDeleteParams) (*EphemeralKey, error) {
	if params == nil {
		params = &EphemeralKeyDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/ephemeral_keys/%s", id)
	ephemeralkey := &EphemeralKey{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, ephemeralkey)
	return ephemeralkey, err
}
