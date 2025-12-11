//
//
// File generated from our OpenAPI spec
//
//

// Package grantedtoken provides the /v1/shared_payment/granted_tokens APIs
package grantedtoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/shared_payment/granted_tokens APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to create SharedPaymentGrantedTokens for testing their integration
func New(params *stripe.TestHelpersSharedPaymentGrantedTokenParams) (*stripe.SharedPaymentGrantedToken, error) {
	return getC().New(params)
}

// Creates a new test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to create SharedPaymentGrantedTokens for testing their integration
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TestHelpersSharedPaymentGrantedTokenParams) (*stripe.SharedPaymentGrantedToken, error) {
	grantedtoken := &stripe.SharedPaymentGrantedToken{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/shared_payment/granted_tokens", c.Key, params, grantedtoken)
	return grantedtoken, err
}

// Revokes a test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to revoke SharedPaymentGrantedTokens for testing their integration
func Update(id string, params *stripe.TestHelpersSharedPaymentGrantedTokenParams) (*stripe.SharedPaymentGrantedToken, error) {
	return getC().Update(id, params)
}

// Revokes a test SharedPaymentGrantedToken object. This endpoint is only available in test mode and allows sellers to revoke SharedPaymentGrantedTokens for testing their integration
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.TestHelpersSharedPaymentGrantedTokenParams) (*stripe.SharedPaymentGrantedToken, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/shared_payment/granted_tokens/%s/revoke", id)
	grantedtoken := &stripe.SharedPaymentGrantedToken{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, grantedtoken)
	return grantedtoken, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
