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

// v2CoreAccountTokenService is used to invoke accounttoken related APIs.
type v2CoreAccountTokenService struct {
	B   Backend
	Key string
}

// Creates an Account Token.
func (c v2CoreAccountTokenService) Create(ctx context.Context, params *V2CoreAccountTokenCreateParams) (*V2CoreAccountToken, error) {
	if params == nil {
		params = &V2CoreAccountTokenCreateParams{}
	}
	params.Context = ctx
	accounttoken := &V2CoreAccountToken{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/account_tokens", c.Key, params, accounttoken)
	return accounttoken, err
}

// Retrieves an Account Token.
func (c v2CoreAccountTokenService) Retrieve(ctx context.Context, id string, params *V2CoreAccountTokenRetrieveParams) (*V2CoreAccountToken, error) {
	if params == nil {
		params = &V2CoreAccountTokenRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/account_tokens/%s", id)
	accounttoken := &V2CoreAccountToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accounttoken)
	return accounttoken, err
}
