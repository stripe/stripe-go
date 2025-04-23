//
//
// File generated from our OpenAPI spec
//
//

// Package refund provides the /v1/refunds APIs
package refund

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
)

// Client is used to invoke /v1/refunds APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Expire a refund with a status of requires_action.
func Expire(id string, params *stripe.TestHelpersRefundExpireParams) (*stripe.Refund, error) {
	return getC().Expire(id, params)
}

// Expire a refund with a status of requires_action.
func (c Client) Expire(id string, params *stripe.TestHelpersRefundExpireParams) (*stripe.Refund, error) {
	path := stripe.FormatURLPath("/v1/test_helpers/refunds/%s/expire", id)
	refund := &stripe.Refund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, refund)
	return refund, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
