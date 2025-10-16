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

// v1DelegatedCheckoutRequestedSessionService is used to invoke /v1/delegated_checkout/requested_sessions APIs.
type v1DelegatedCheckoutRequestedSessionService struct {
	B   Backend
	Key string
}

// Creates a requested session
func (c v1DelegatedCheckoutRequestedSessionService) Create(ctx context.Context, params *DelegatedCheckoutRequestedSessionCreateParams) (*DelegatedCheckoutRequestedSession, error) {
	if params == nil {
		params = &DelegatedCheckoutRequestedSessionCreateParams{}
	}
	params.Context = ctx
	requestedsession := &DelegatedCheckoutRequestedSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/delegated_checkout/requested_sessions", c.Key, params, requestedsession)
	return requestedsession, err
}

// Retrieves a requested session
func (c v1DelegatedCheckoutRequestedSessionService) Retrieve(ctx context.Context, id string, params *DelegatedCheckoutRequestedSessionRetrieveParams) (*DelegatedCheckoutRequestedSession, error) {
	if params == nil {
		params = &DelegatedCheckoutRequestedSessionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/delegated_checkout/requested_sessions/%s", id)
	requestedsession := &DelegatedCheckoutRequestedSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, requestedsession)
	return requestedsession, err
}

// Updates a requested session
func (c v1DelegatedCheckoutRequestedSessionService) Update(ctx context.Context, id string, params *DelegatedCheckoutRequestedSessionUpdateParams) (*DelegatedCheckoutRequestedSession, error) {
	if params == nil {
		params = &DelegatedCheckoutRequestedSessionUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/delegated_checkout/requested_sessions/%s", id)
	requestedsession := &DelegatedCheckoutRequestedSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, requestedsession)
	return requestedsession, err
}

// Confirms a requested session
func (c v1DelegatedCheckoutRequestedSessionService) Confirm(ctx context.Context, id string, params *DelegatedCheckoutRequestedSessionConfirmParams) (*DelegatedCheckoutRequestedSession, error) {
	if params == nil {
		params = &DelegatedCheckoutRequestedSessionConfirmParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/delegated_checkout/requested_sessions/%s/confirm", id)
	requestedsession := &DelegatedCheckoutRequestedSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, requestedsession)
	return requestedsession, err
}

// Expires a requested session
func (c v1DelegatedCheckoutRequestedSessionService) Expire(ctx context.Context, id string, params *DelegatedCheckoutRequestedSessionExpireParams) (*DelegatedCheckoutRequestedSession, error) {
	if params == nil {
		params = &DelegatedCheckoutRequestedSessionExpireParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/delegated_checkout/requested_sessions/%s/expire", id)
	requestedsession := &DelegatedCheckoutRequestedSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, requestedsession)
	return requestedsession, err
}
