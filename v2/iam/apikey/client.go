//
//
// File generated from our OpenAPI spec
//
//

// Package apikey provides the apikey related APIs
package apikey

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke apikey related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create an API key. To create a secret key in livemode, a public key for encryption must be provided.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2IamAPIKeyParams) (*stripe.V2IamAPIKey, error) {
	apikey := &stripe.V2IamAPIKey{}
	err := c.B.Call(http.MethodPost, "/v2/iam/api_keys", c.Key, params, apikey)
	return apikey, err
}

// Retrieve an API key. For livemode secret keys, secret tokens are only returned on creation, and never returned here.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2IamAPIKeyParams) (*stripe.V2IamAPIKey, error) {
	path := stripe.FormatURLPath("/v2/iam/api_keys/%s", id)
	apikey := &stripe.V2IamAPIKey{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, apikey)
	return apikey, err
}

// Update an API key. Only parameters that are specified in the request will be updated.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2IamAPIKeyParams) (*stripe.V2IamAPIKey, error) {
	path := stripe.FormatURLPath("/v2/iam/api_keys/%s", id)
	apikey := &stripe.V2IamAPIKey{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, apikey)
	return apikey, err
}

// Expire an API key. The specified key becomes invalid immediately.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Expire(id string, params *stripe.V2IamAPIKeyExpireParams) (*stripe.V2IamAPIKey, error) {
	path := stripe.FormatURLPath("/v2/iam/api_keys/%s/expire", id)
	apikey := &stripe.V2IamAPIKey{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, apikey)
	return apikey, err
}

// Rotate an API key. A new key with the same properties is created and returned. The existing key is expired immediately, unless an expiry time is specified.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Rotate(id string, params *stripe.V2IamAPIKeyRotateParams) (*stripe.V2IamAPIKey, error) {
	path := stripe.FormatURLPath("/v2/iam/api_keys/%s/rotate", id)
	apikey := &stripe.V2IamAPIKey{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, apikey)
	return apikey, err
}

// List all API keys of an account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2IamAPIKeyListParams) stripe.Seq2[*stripe.V2IamAPIKey, error] {
	return stripe.NewV2List("/v2/iam/api_keys", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2IamAPIKey], error) {
		page := &stripe.V2Page[*stripe.V2IamAPIKey]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
