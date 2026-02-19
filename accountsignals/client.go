//
//
// File generated from our OpenAPI spec
//
//

// Package accountsignals provides the /v1/accounts/{account_id}/signals APIs
package accountsignals

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/accounts/{account_id}/signals APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the account's Signal objects
func Get(id string, params *stripe.AccountSignalsParams) (*stripe.AccountSignals, error) {
	return getC().Get(id, params)
}

// Retrieves the account's Signal objects
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.AccountSignalsParams) (*stripe.AccountSignals, error) {
	path := stripe.FormatURLPath("/v1/accounts/%s/signals", id)
	accountsignals := &stripe.AccountSignals{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountsignals)
	return accountsignals, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
