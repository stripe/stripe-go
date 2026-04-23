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

// v1SharedPaymentGrantedTokenService is used to invoke /v1/shared_payment/granted_tokens APIs.
type v1SharedPaymentGrantedTokenService struct {
	B   Backend
	Key string
}

// Retrieves an existing SharedPaymentGrantedToken object
func (c v1SharedPaymentGrantedTokenService) Retrieve(ctx context.Context, id string, params *SharedPaymentGrantedTokenRetrieveParams) (*SharedPaymentGrantedToken, error) {
	if params == nil {
		params = &SharedPaymentGrantedTokenRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/shared_payment/granted_tokens/%s", id)
	grantedtoken := &SharedPaymentGrantedToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, grantedtoken)
	return grantedtoken, err
}
