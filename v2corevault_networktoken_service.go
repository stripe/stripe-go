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

// Create or Return a Network Token Using Raw Card Data.
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

// Retrieves an existing network token.
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

// Creates or returns a Network Token from an existing card reference.
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
