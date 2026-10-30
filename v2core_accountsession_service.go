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

// v2CoreAccountSessionService is used to invoke accountsession related APIs.
type v2CoreAccountSessionService struct {
	B   Backend
	Key string
}

// Create an Account Session.
func (c v2CoreAccountSessionService) Create(ctx context.Context, params *V2CoreAccountSessionCreateParams) (*V2CoreAccountSession, error) {
	if params == nil {
		params = &V2CoreAccountSessionCreateParams{}
	}
	params.Context = ctx
	accountsession := &V2CoreAccountSession{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/account_sessions", c.Key, params, accountsession)
	return accountsession, err
}
