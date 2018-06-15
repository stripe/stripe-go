// Package ephemeralkey provides the /ephemeral_keys APIs
package ephemeralkey

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /ephemeral_keys APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new ephemeral keys.
// For more details see https://stripe.com/docs/api#create_ephemeral_key.
func New(params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	return getC().New(params)
}

// New POSTs new ephemeral keys.
// For more details see https://stripe.com/docs/api#create_ephemeral_key.
func (c Client) New(params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	if params.StripeVersion == nil || len(stripe.StringValue(params.StripeVersion)) == 0 {
		return nil, fmt.Errorf("params.StripeVersion must be specified")
	}

	if params.Headers == nil {
		params.Headers = make(http.Header)
	}
	params.Headers.Add("Stripe-Version", stripe.StringValue(params.StripeVersion))

	ephemeralKey := &stripe.EphemeralKey{}
	err := c.B.Call(http.MethodPost, "/ephemeral_keys", c.Key, params, ephemeralKey)
	return ephemeralKey, err
}

// Del removes an ephemeral key.
// For more details see https://stripe.com/docs/api#delete_ephemeral_key.
func Del(id string, params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	return getC().Del(id, params)
}

// Del removes an ephemeral key.
// For more details see https://stripe.com/docs/api#delete_ephemeral_key.
func (c Client) Del(id string, params *stripe.EphemeralKeyParams) (*stripe.EphemeralKey, error) {
	path := stripe.FormatURLPath("/ephemeral_keys/%s", id)
	ephemeralKey := &stripe.EphemeralKey{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, ephemeralKey)
	return ephemeralKey, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
