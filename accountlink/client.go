//
//
// File generated from our OpenAPI spec
//
//

// Package accountlink provides the /account_links APIs
package accountlink

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /account_links APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new account link.
func New(params *stripe.AccountLinkParams) (*stripe.AccountLink, error) {
	return getC().New(params)
}

// New creates a new account link.
func (c Client) New(params *stripe.AccountLinkParams) (*stripe.AccountLink, error) {
	accountlink := &stripe.AccountLink{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/account_links",
		c.Key,
		params,
		accountlink,
	)
	return accountlink, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
