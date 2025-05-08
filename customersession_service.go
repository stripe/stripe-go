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

// v1CustomerSessionService is used to invoke /v1/customer_sessions APIs.
type v1CustomerSessionService struct {
	B   Backend
	Key string
}

// Creates a Customer Session object that includes a single-use client secret that you can use on your front-end to grant client-side API access for certain customer resources.
func (c v1CustomerSessionService) Create(ctx context.Context, params *CustomerSessionCreateParams) (*CustomerSession, error) {
	if params == nil {
		params = &CustomerSessionCreateParams{}
	}
	params.Context = ctx
	customersession := &CustomerSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/customer_sessions", c.Key, params, customersession)
	return customersession, err
}
