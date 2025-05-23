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

// v1TerminalConnectionTokenService is used to invoke /v1/terminal/connection_tokens APIs.
type v1TerminalConnectionTokenService struct {
	B   Backend
	Key string
}

// To connect to a reader the Stripe Terminal SDK needs to retrieve a short-lived connection token from Stripe, proxied through your server. On your backend, add an endpoint that creates and returns a connection token.
func (c v1TerminalConnectionTokenService) Create(ctx context.Context, params *TerminalConnectionTokenCreateParams) (*TerminalConnectionToken, error) {
	if params == nil {
		params = &TerminalConnectionTokenCreateParams{}
	}
	params.Context = ctx
	connectiontoken := &TerminalConnectionToken{}
	err := c.B.Call(
		http.MethodPost, "/v1/terminal/connection_tokens", c.Key, params, connectiontoken)
	return connectiontoken, err
}
