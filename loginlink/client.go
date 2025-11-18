//
//
// File generated from our OpenAPI spec
//
//

// Package loginlink provides the /v1/accounts/{account}/login_links APIs
package loginlink

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/accounts/{account}/login_links APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a login link for a connected account to access the Express Dashboard.
//
// You can only create login links for accounts that use the [Express Dashboard](https://docs.stripe.com/connect/express-dashboard) and are connected to your platform.
func New(params *stripe.LoginLinkParams) (*stripe.LoginLink, error) {
	return getC().New(params)
}

// Creates a login link for a connected account to access the Express Dashboard.
//
// You can only create login links for accounts that use the [Express Dashboard](https://docs.stripe.com/connect/express-dashboard) and are connected to your platform.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.LoginLinkParams) (*stripe.LoginLink, error) {
	if params.Account == nil {
		return nil, fmt.Errorf("Invalid login link params: Account must be set")
	}
	path := stripe.FormatURLPath(
		"/v1/accounts/%s/login_links", stripe.StringValue(params.Account))
	loginlink := &stripe.LoginLink{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, loginlink)
	return loginlink, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
