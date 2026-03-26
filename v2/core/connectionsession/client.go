//
//
// File generated from our OpenAPI spec
//
//

// Package connectionsession provides the connectionsession related APIs
package connectionsession

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke connectionsession related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a new ConnectionSession.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreConnectionSessionParams) (*stripe.V2CoreConnectionSession, error) {
	connectionsession := &stripe.V2CoreConnectionSession{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/connection_sessions", c.Key, params, connectionsession)
	return connectionsession, err
}

// Retrieve a ConnectionSession.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreConnectionSessionParams) (*stripe.V2CoreConnectionSession, error) {
	path := stripe.FormatURLPath("/v2/core/connection_sessions/%s", id)
	connectionsession := &stripe.V2CoreConnectionSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, connectionsession)
	return connectionsession, err
}
