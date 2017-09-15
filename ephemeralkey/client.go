// Package ephemeralkey provides the /ephemeral_keys APIs
package ephemeralkey

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
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
	if params.StripeVersion == "" {
		return nil, fmt.Errorf("params.StripeVersion must be specified")
	}

	body := &form.Values{}
	form.AppendTo(body, params)

	if len(params.StripeVersion) > 0 {
		if params.Headers == nil {
			params.Headers = make(http.Header)
		}
		params.Headers.Add("Stripe-Version", params.StripeVersion)
	}

	ephemeralKey := &stripe.EphemeralKey{}
	err := c.B.Call("POST", "ephemeral_keys", c.Key, body, &params.Params, ephemeralKey)

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
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	ephemeralKey := &stripe.EphemeralKey{}
	err := c.B.Call("DELETE", "/ephemeral_keys/"+id, c.Key, body, commonParams, ephemeralKey)

	return ephemeralKey, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
