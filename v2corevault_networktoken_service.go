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

// v2CoreVaultNetworkTokenService is used to invoke networktoken related APIs.
type v2CoreVaultNetworkTokenService struct {
	B   Backend
	Key string
}

// Creates or returns a NetworkToken from raw card data for POST /v2/core/vault/network_tokens.
func (c v2CoreVaultNetworkTokenService) Create(ctx context.Context, params *V2CoreVaultNetworkTokenCreateParams) (*V2CoreVaultNetworkToken, error) {
	if params == nil {
		params = &V2CoreVaultNetworkTokenCreateParams{}
	}
	params.Context = ctx
	networktoken := &V2CoreVaultNetworkToken{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/network_tokens", c.Key, params, networktoken)
	return networktoken, err
}

// Retrieves the persisted NetworkToken projection for GET /v2/core/vault/network_tokens/:id.
func (c v2CoreVaultNetworkTokenService) Retrieve(ctx context.Context, id string, params *V2CoreVaultNetworkTokenRetrieveParams) (*V2CoreVaultNetworkToken, error) {
	if params == nil {
		params = &V2CoreVaultNetworkTokenRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/vault/network_tokens/%s", id)
	networktoken := &V2CoreVaultNetworkToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, networktoken)
	return networktoken, err
}

// Creates or returns a NetworkToken from an existing card reference for POST /v2/core/vault/network_tokens/create_from_credential.
func (c v2CoreVaultNetworkTokenService) CreateFromCredential(ctx context.Context, params *V2CoreVaultNetworkTokenCreateFromCredentialParams) (*V2CoreVaultNetworkToken, error) {
	if params == nil {
		params = &V2CoreVaultNetworkTokenCreateFromCredentialParams{}
	}
	params.Context = ctx
	networktoken := &V2CoreVaultNetworkToken{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/network_tokens/create_from_credential", c.Key, params, networktoken)
	return networktoken, err
}

// Generates a single-use cryptogram for POST /v2/core/vault/network_tokens/:id/generate_cryptogram.
// Every successful call generates a new cryptogram, and retrying can generate another cryptogram.
// The cryptogram is returned only in this response and is never persisted.
func (c v2CoreVaultNetworkTokenService) GenerateCryptogram(ctx context.Context, id string, params *V2CoreVaultNetworkTokenGenerateCryptogramParams) (*V2CoreVaultNetworkToken, error) {
	if params == nil {
		params = &V2CoreVaultNetworkTokenGenerateCryptogramParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/core/vault/network_tokens/%s/generate_cryptogram", id)
	networktoken := &V2CoreVaultNetworkToken{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, networktoken)
	return networktoken, err
}
