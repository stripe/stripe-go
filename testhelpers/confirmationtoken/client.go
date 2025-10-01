//
//
// File generated from our OpenAPI spec
//
//

// Package confirmationtoken provides the /v1/confirmation_tokens APIs
package confirmationtoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke /v1/confirmation_tokens APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a test mode Confirmation Token server side for your integration tests.
func New(params *stripe.TestHelpersConfirmationTokenParams) (*stripe.ConfirmationToken, error) {
	return getC().New(params)
}

// Creates a test mode Confirmation Token server side for your integration tests.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TestHelpersConfirmationTokenParams) (*stripe.ConfirmationToken, error) {
	confirmationtoken := &stripe.ConfirmationToken{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/confirmation_tokens", c.Key, params, confirmationtoken)
	return confirmationtoken, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
