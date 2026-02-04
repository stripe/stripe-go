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

// v2CoreConnectionSessionService is used to invoke connectionsession related APIs.
type v2CoreConnectionSessionService struct {
	B   Backend
	Key string
}

// Create a new ConnectionSession.
func (c v2CoreConnectionSessionService) Create(ctx context.Context, params *V2CoreConnectionSessionCreateParams) (*V2CoreConnectionSession, error) {
	if params == nil {
		params = &V2CoreConnectionSessionCreateParams{}
	}
	params.Context = ctx
	connectionsession := &V2CoreConnectionSession{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/connection_sessions", c.Key, params, connectionsession)
	return connectionsession, err
}

// Retrieve a ConnectionSession.
func (c v2CoreConnectionSessionService) Retrieve(ctx context.Context, id string, params *V2CoreConnectionSessionRetrieveParams) (*V2CoreConnectionSession, error) {
	if params == nil {
		params = &V2CoreConnectionSessionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/connection_sessions/%s", id)
	connectionsession := &V2CoreConnectionSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, connectionsession)
	return connectionsession, err
}
