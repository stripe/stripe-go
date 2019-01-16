// Package accountlink provides API functions related to account links.
//
// For more details, see: https://stripe.com/docs/api/go#account_links.
package accountlink

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke APIs related to account links.
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
	link := &stripe.AccountLink{}
	err := c.B.Call(http.MethodPost, "/v1/account_links", c.Key, params, link)
	return link, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
