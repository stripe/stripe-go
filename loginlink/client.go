// Package loginlink provides the /login_links APIs
package loginlink

import (
	"errors"
	"net/http"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /login_links APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs to create a login_link for an account
// For more details see https://stripe.com/docs/api#create_login_link.
func New(params *stripe.LoginLinkParams) (*stripe.LoginLink, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.LoginLinkParams) (*stripe.LoginLink, error) {
	if params.Account == nil {
		return nil, errors.New("Invalid login link params: Account must be set")
	}

	path := stripe.FormatURLPath("/accounts/%s/login_links", stripe.StringValue(params.Account))
	loginLink := &stripe.LoginLink{}
	err := c.B.Call(http.MethodPost, path, c.Key, nil, loginLink)
	return loginLink, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
