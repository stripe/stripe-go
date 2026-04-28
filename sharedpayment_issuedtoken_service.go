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

// v1SharedPaymentIssuedTokenService is used to invoke /v1/shared_payment/issued_tokens APIs.
type v1SharedPaymentIssuedTokenService struct {
	B   Backend
	Key string
}

// Creates a new SharedPaymentIssuedToken object
func (c v1SharedPaymentIssuedTokenService) Create(ctx context.Context, params *SharedPaymentIssuedTokenCreateParams) (*SharedPaymentIssuedToken, error) {
	if params == nil {
		params = &SharedPaymentIssuedTokenCreateParams{}
	}
	params.Context = ctx
	issuedtoken := &SharedPaymentIssuedToken{}
	err := c.B.Call(
		http.MethodPost, "/v1/shared_payment/issued_tokens", c.Key, params, issuedtoken)
	return issuedtoken, err
}

// Retrieves an existing SharedPaymentIssuedToken object
func (c v1SharedPaymentIssuedTokenService) Retrieve(ctx context.Context, id string, params *SharedPaymentIssuedTokenRetrieveParams) (*SharedPaymentIssuedToken, error) {
	if params == nil {
		params = &SharedPaymentIssuedTokenRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/shared_payment/issued_tokens/%s", id)
	issuedtoken := &SharedPaymentIssuedToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, issuedtoken)
	return issuedtoken, err
}

// Revokes a SharedPaymentIssuedToken
func (c v1SharedPaymentIssuedTokenService) Revoke(ctx context.Context, id string, params *SharedPaymentIssuedTokenRevokeParams) (*SharedPaymentIssuedToken, error) {
	if params == nil {
		params = &SharedPaymentIssuedTokenRevokeParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/shared_payment/issued_tokens/%s/revoke", id)
	issuedtoken := &SharedPaymentIssuedToken{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, issuedtoken)
	return issuedtoken, err
}
