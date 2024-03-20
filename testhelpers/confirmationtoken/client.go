//
//
// File generated from our OpenAPI spec
//
//

// Package confirmationtoken provides the /confirmation_tokens APIs
package confirmationtoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /confirmation_tokens APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new confirmation token.
func New(params *stripe.TestHelpersConfirmationTokenParams) (*stripe.ConfirmationToken, error) {
	return getC().New(params)
}

// New creates a new confirmation token.
func (c Client) New(params *stripe.TestHelpersConfirmationTokenParams) (*stripe.ConfirmationToken, error) {
	confirmationtoken := &stripe.ConfirmationToken{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/test_helpers/confirmation_tokens",
		c.Key,
		params,
		confirmationtoken,
	)
	return confirmationtoken, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
