//
//
// File generated from our OpenAPI spec
//
//

// Package connectiontoken provides the /v1/terminal/connection_tokens APIs
package connectiontoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke /v1/terminal/connection_tokens APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// To connect to a reader the Stripe Terminal SDK needs to retrieve a short-lived connection token from Stripe, proxied through your server. On your backend, add an endpoint that creates and returns a connection token.
func New(params *stripe.TerminalConnectionTokenParams) (*stripe.TerminalConnectionToken, error) {
	return getC().New(params)
}

// To connect to a reader the Stripe Terminal SDK needs to retrieve a short-lived connection token from Stripe, proxied through your server. On your backend, add an endpoint that creates and returns a connection token.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TerminalConnectionTokenParams) (*stripe.TerminalConnectionToken, error) {
	connectiontoken := &stripe.TerminalConnectionToken{}
	err := c.B.Call(
		http.MethodPost, "/v1/terminal/connection_tokens", c.Key, params, connectiontoken)
	return connectiontoken, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
