//
//
// File generated from our OpenAPI spec
//
//

// Package balancesettings provides the /v1/balance_settings APIs
package balancesettings

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke /v1/balance_settings APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
func Get(params *stripe.BalanceSettingsParams) (*stripe.BalanceSettings, error) {
	return getC().Get(params)
}

// Retrieves balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(params *stripe.BalanceSettingsParams) (*stripe.BalanceSettings, error) {
	balancesettings := &stripe.BalanceSettings{}
	err := c.B.Call(
		http.MethodGet, "/v1/balance_settings", c.Key, params, balancesettings)
	return balancesettings, err
}

// Updates balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
func Update(params *stripe.BalanceSettingsParams) (*stripe.BalanceSettings, error) {
	return getC().Update(params)
}

// Updates balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://docs.stripe.com/connect/authentication)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(params *stripe.BalanceSettingsParams) (*stripe.BalanceSettings, error) {
	balancesettings := &stripe.BalanceSettings{}
	err := c.B.Call(
		http.MethodPost, "/v1/balance_settings", c.Key, params, balancesettings)
	return balancesettings, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
