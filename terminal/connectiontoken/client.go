//
//
// File generated from our OpenAPI spec
//
//

// Package connectiontoken provides the /terminal/connection_tokens APIs
package connectiontoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /terminal/connection_tokens APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new terminal connection token.
func New(params *stripe.TerminalConnectionTokenParams) (*stripe.TerminalConnectionToken, error) {
	return getC().New(params)
}

// New creates a new terminal connection token.
func (c Client) New(params *stripe.TerminalConnectionTokenParams) (*stripe.TerminalConnectionToken, error) {
	connectiontoken := &stripe.TerminalConnectionToken{}
	var err error
	sr := stripe.NewStripeRequest(
		http.MethodPost,
		"/v1/terminal/connection_tokens",
		c.Key,
	)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, connectiontoken)
	return connectiontoken, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
