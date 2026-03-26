//
//
// File generated from our OpenAPI spec
//
//

// Package paymentattempt provides the /v1/orchestration/payment_attempts APIs
package paymentattempt

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke /v1/orchestration/payment_attempts APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves orchestration information for the given payment attempt record (e.g. return url).
func Get(id string, params *stripe.OrchestrationPaymentAttemptParams) (*stripe.OrchestrationPaymentAttempt, error) {
	return getC().Get(id, params)
}

// Retrieves orchestration information for the given payment attempt record (e.g. return url).
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.OrchestrationPaymentAttemptParams) (*stripe.OrchestrationPaymentAttempt, error) {
	path := stripe.FormatURLPath("/v1/orchestration/payment_attempts/%s", id)
	paymentattempt := &stripe.OrchestrationPaymentAttempt{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentattempt)
	return paymentattempt, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
