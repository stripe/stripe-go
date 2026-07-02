//
//
// File generated from our OpenAPI spec
//
//

// Package customer provides the /v1/crypto/customers APIs
package customer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke /v1/crypto/customers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Crypto Customer.
func Get(id string, params *stripe.CryptoCustomerParams) (*stripe.CryptoCustomer, error) {
	return getC().Get(id, params)
}

// Retrieves the details of a Crypto Customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.CryptoCustomerParams) (*stripe.CryptoCustomer, error) {
	path := stripe.FormatURLPath("/v1/crypto/customers/%s", id)
	customer := &stripe.CryptoCustomer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, customer)
	return customer, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
