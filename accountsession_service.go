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

// v1AccountSessionService is used to invoke /v1/account_sessions APIs.
type v1AccountSessionService struct {
	B   Backend
	Key string
}

// Creates a AccountSession object that includes a single-use token that the platform can use on their front-end to grant client-side API access.
func (c v1AccountSessionService) Create(ctx context.Context, params *AccountSessionCreateParams) (*AccountSession, error) {
	if params == nil {
		params = &AccountSessionCreateParams{}
	}
	params.Context = ctx
	accountsession := &AccountSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/account_sessions", c.Key, params, accountsession)
	return accountsession, err
}
