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

// v1ConfirmationTokenService is used to invoke /v1/confirmation_tokens APIs.
type v1ConfirmationTokenService struct {
	B   Backend
	Key string
}

// Retrieves an existing ConfirmationToken object
func (c v1ConfirmationTokenService) Retrieve(ctx context.Context, id string, params *ConfirmationTokenRetrieveParams) (*ConfirmationToken, error) {
	path := FormatURLPath("/v1/confirmation_tokens/%s", id)
	confirmationtoken := &ConfirmationToken{}
	if params == nil {
		params = &ConfirmationTokenRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, confirmationtoken)
	return confirmationtoken, err
}
