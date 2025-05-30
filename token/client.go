//
//
// File generated from our OpenAPI spec
//
//

// Package token provides the /v1/tokens APIs
package token

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke /v1/tokens APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a single-use token that represents a bank account's details.
// You can use this token with any v1 API method in place of a bank account dictionary. You can only use this token once. To do so, attach it to a [connected account](https://docs.stripe.com/api#accounts) where [controller.requirement_collection](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection) is application, which includes Custom accounts.
func New(params *stripe.TokenParams) (*stripe.Token, error) {
	return getC().New(params)
}

// Creates a single-use token that represents a bank account's details.
// You can use this token with any v1 API method in place of a bank account dictionary. You can only use this token once. To do so, attach it to a [connected account](https://docs.stripe.com/api#accounts) where [controller.requirement_collection](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection) is application, which includes Custom accounts.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TokenParams) (*stripe.Token, error) {
	token := &stripe.Token{}
	err := c.B.Call(http.MethodPost, "/v1/tokens", c.Key, params, token)
	return token, err
}

// Retrieves the token with the given ID.
func Get(id string, params *stripe.TokenParams) (*stripe.Token, error) {
	return getC().Get(id, params)
}

// Retrieves the token with the given ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TokenParams) (*stripe.Token, error) {
	path := stripe.FormatURLPath("/v1/tokens/%s", id)
	token := &stripe.Token{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, token)
	return token, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
