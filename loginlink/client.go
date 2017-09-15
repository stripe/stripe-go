// Package loginlink provides the /login_links APIs
package loginlink

import (
	"errors"
	"fmt"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
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
	body := &form.Values{}

	loginLink := &stripe.LoginLink{}
	var err error

	if len(params.Account) > 0 {
		err = c.B.Call("POST", fmt.Sprintf("/accounts/%v/login_links", params.Account), c.Key, body, nil, loginLink)
	} else {
		err = errors.New("Invalid login link params: Account must be set")
	}

	return loginLink, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
