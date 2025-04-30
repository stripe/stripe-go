//
//
// File generated from our OpenAPI spec
//
//

// Package accountsession provides the /v1/account_sessions APIs
package accountsession

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/account_sessions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a AccountSession object that includes a single-use token that the platform can use on their front-end to grant client-side API access.
func New(params *stripe.AccountSessionParams) (*stripe.AccountSession, error) {
	return getC().New(params)
}

// Creates a AccountSession object that includes a single-use token that the platform can use on their front-end to grant client-side API access.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.AccountSessionParams) (*stripe.AccountSession, error) {
	accountsession := &stripe.AccountSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/account_sessions", c.Key, params, accountsession)
	return accountsession, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
