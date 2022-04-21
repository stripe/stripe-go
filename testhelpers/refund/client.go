//
//
// File generated from our OpenAPI spec
//
//

// Package refund provides the /refunds APIs
package refund

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /refunds APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Expire is the method for the `POST /v1/test_helpers/refunds/{refund}/expire` API.
func Expire(id string, params *stripe.TestHelpersRefundExpireParams) (*stripe.Refund, error) {
	return getC().Expire(id, params)
}

// Expire is the method for the `POST /v1/test_helpers/refunds/{refund}/expire` API.
func (c Client) Expire(id string, params *stripe.TestHelpersRefundExpireParams) (*stripe.Refund, error) {
	path := stripe.FormatURLPath("/v1/test_helpers/refunds/%s/expire", id)
	refund := &stripe.Refund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, refund)
	return refund, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
