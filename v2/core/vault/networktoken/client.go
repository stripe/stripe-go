//
//
// File generated from our OpenAPI spec
//
//

// Package networktoken provides the networktoken related APIs
package networktoken

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke networktoken related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create or Return a Network Token Using Raw Card Data.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreVaultNetworkTokenParams) (*stripe.V2CoreVaultNetworkToken, error) {
	networktoken := &stripe.V2CoreVaultNetworkToken{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/network_tokens", c.Key, params, networktoken)
	return networktoken, err
}

// Retrieves an existing network token.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreVaultNetworkTokenParams) (*stripe.V2CoreVaultNetworkToken, error) {
	path := stripe.FormatURLPath("/v2/core/vault/network_tokens/%s", id)
	networktoken := &stripe.V2CoreVaultNetworkToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, networktoken)
	return networktoken, err
}

// Creates or returns a Network Token from an existing card reference.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) CreateFromCredential(params *stripe.V2CoreVaultNetworkTokenCreateFromCredentialParams) (*stripe.V2CoreVaultNetworkToken, error) {
	networktoken := &stripe.V2CoreVaultNetworkToken{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/vault/network_tokens/create_from_credential", c.Key, params, networktoken)
	return networktoken, err
}

// Every successful call generates a new cryptogram, and retrying can generate another cryptogram.
// The cryptogram is returned only in this response and is never persisted.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) GenerateCryptogram(id string, params *stripe.V2CoreVaultNetworkTokenGenerateCryptogramParams) (*stripe.V2CoreVaultNetworkToken, error) {
	path := stripe.FormatURLPath(
		"/v2/core/vault/network_tokens/%s/generate_cryptogram", id)
	networktoken := &stripe.V2CoreVaultNetworkToken{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, networktoken)
	return networktoken, err
}
