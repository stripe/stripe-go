//
//
// File generated from our OpenAPI spec
//
//

// Package persontoken provides the persontoken related APIs
package persontoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke persontoken related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a single-use token that represents the details for a person. Use this when you create or update persons associated with an Account v2. Learn more about [account tokens](https://docs.stripe.com/connect/account-tokens).
// You can only create person tokens with your application's publishable key and in live mode. You can use your application's secret key to create person tokens only in test mode.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreAccountsPersonTokenParams) (*stripe.V2CoreAccountPersonToken, error) {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/person_tokens", stripe.StringValue(params.AccountID))
	accountpersontoken := &stripe.V2CoreAccountPersonToken{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountpersontoken)
	return accountpersontoken, err
}

// Retrieves a Person Token associated with an Account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreAccountsPersonTokenParams) (*stripe.V2CoreAccountPersonToken, error) {
	path := stripe.FormatURLPath(
		"/v2/core/accounts/%s/person_tokens/%s", stripe.StringValue(
			params.AccountID), id)
	accountpersontoken := &stripe.V2CoreAccountPersonToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountpersontoken)
	return accountpersontoken, err
}
