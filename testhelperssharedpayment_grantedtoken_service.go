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

// v1TestHelpersSharedPaymentGrantedTokenService is used to invoke /v1/shared_payment/granted_tokens APIs.
type v1TestHelpersSharedPaymentGrantedTokenService struct {
	B   Backend
	Key string
}

// Creates a new test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to create SharedPaymentGrantedTokens for testing their integration
func (c v1TestHelpersSharedPaymentGrantedTokenService) Create(ctx context.Context, params *TestHelpersSharedPaymentGrantedTokenCreateParams) (*SharedPaymentGrantedToken, error) {
	if params == nil {
		params = &TestHelpersSharedPaymentGrantedTokenCreateParams{}
	}
	params.Context = ctx
	grantedtoken := &SharedPaymentGrantedToken{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/shared_payment/granted_tokens", c.Key, params, grantedtoken)
	return grantedtoken, err
}

// Revokes a test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to revoke SharedPaymentGrantedTokens for testing their integration
func (c v1TestHelpersSharedPaymentGrantedTokenService) Revoke(ctx context.Context, id string, params *TestHelpersSharedPaymentGrantedTokenRevokeParams) (*SharedPaymentGrantedToken, error) {
	if params == nil {
		params = &TestHelpersSharedPaymentGrantedTokenRevokeParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/shared_payment/granted_tokens/%s/revoke", id)
	grantedtoken := &SharedPaymentGrantedToken{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, grantedtoken)
	return grantedtoken, err
}
