//
//
// File generated from our OpenAPI spec
//
//

// Package loginlink provides the /accounts/{account}/login_links APIs
package loginlink

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /accounts/{account}/login_links APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new login link.
func New(params *stripe.LoginLinkParams) (*stripe.LoginLink, error) {
	return getC().New(params)
}

// New creates a new login link.
func (c Client) New(params *stripe.LoginLinkParams) (*stripe.LoginLink, error) {
	if params.Account == nil {
		return nil, fmt.Errorf("Invalid login link params: Account must be set")
	}
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/login_links",
		stripe.StringValue(params.Account),
	)
	loginlink := &stripe.LoginLink{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, loginlink)
	return loginlink, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
