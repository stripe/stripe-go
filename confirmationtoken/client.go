//
//
// File generated from our OpenAPI spec
//
//

// Package confirmationtoken provides the /v1/confirmation_tokens APIs
package confirmationtoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/confirmation_tokens APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an existing ConfirmationToken object
func Get(id string, params *stripe.ConfirmationTokenParams) (*stripe.ConfirmationToken, error) {
	return getC().Get(id, params)
}

// Retrieves an existing ConfirmationToken object
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ConfirmationTokenParams) (*stripe.ConfirmationToken, error) {
	path := stripe.FormatURLPath("/v1/confirmation_tokens/%s", id)
	confirmationtoken := &stripe.ConfirmationToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, confirmationtoken)
	return confirmationtoken, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
