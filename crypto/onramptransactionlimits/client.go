//
//
// File generated from our OpenAPI spec
//
//

// Package onramptransactionlimits provides the /v1/crypto/onramp_transaction_limits APIs
package onramptransactionlimits

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke /v1/crypto/onramp_transaction_limits APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the remaining onramp limit for a crypto customer.
func Get(params *stripe.CryptoOnrampTransactionLimitsParams) (*stripe.CryptoOnrampTransactionLimits, error) {
	return getC().Get(params)
}

// Retrieves the remaining onramp limit for a crypto customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(params *stripe.CryptoOnrampTransactionLimitsParams) (*stripe.CryptoOnrampTransactionLimits, error) {
	onramptransactionlimits := &stripe.CryptoOnrampTransactionLimits{}
	err := c.B.Call(
		http.MethodGet, "/v1/crypto/onramp_transaction_limits", c.Key, params, onramptransactionlimits)
	return onramptransactionlimits, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
