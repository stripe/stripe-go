//
//
// File generated from our OpenAPI spec
//
//

// Package ephemeralkey provides the /v1/ephemeral_keys APIs
package ephemeralkey

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/ephemeral_keys APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a short-lived API key for a given resource.
func New(params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	return getC().New(params)
}

// Creates a short-lived API key for a given resource.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	if params.StripeVersion == nil || len(stripe.StringValue(params.StripeVersion)) == 0 {
		return nil, fmt.Errorf("params.StripeVersion must be specified")
	}

	if params.Headers == nil {
		params.Headers = make(http.Header)
	}
	params.Headers.Add("Stripe-Version", stripe.StringValue(params.StripeVersion))

	ephemeralkey := &stripe.EphemeralKey{}
	err := c.B.Call(
		http.MethodPost, "/v1/ephemeral_keys", c.Key, params, ephemeralkey)
	return ephemeralkey, err
}

// Invalidates a short-lived API key for a given resource.
func Del(id string, params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	return getC().Del(id, params)
}

// Invalidates a short-lived API key for a given resource.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	path := stripe.FormatURLPath("/v1/ephemeral_keys/%s", id)
	ephemeralkey := &stripe.EphemeralKey{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, ephemeralkey)
	return ephemeralkey, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
