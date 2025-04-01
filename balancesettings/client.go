//
//
// File generated from our OpenAPI spec
//
//

// Package balancesettings provides the /v1/balance_settings APIs
package balancesettings

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
)

// Client is used to invoke /v1/balance_settings APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://stripe.com/connect/authentication)
func Get(params *stripe.BalanceSettingsParams) (*stripe.BalanceSettings, error) {
	return getC().Get(params)
}

// Retrieves balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://stripe.com/connect/authentication)
func (c Client) Get(params *stripe.BalanceSettingsParams) (*stripe.BalanceSettings, error) {
	balancesettings := &stripe.BalanceSettings{}
	err := c.B.Call(
		http.MethodGet, "/v1/balance_settings", c.Key, params, balancesettings)
	return balancesettings, err
}

// Updates balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://stripe.com/connect/authentication)
func Update(params *stripe.BalanceSettingsParams) (*stripe.BalanceSettings, error) {
	return getC().Update(params)
}

// Updates balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://stripe.com/connect/authentication)
func (c Client) Update(params *stripe.BalanceSettingsParams) (*stripe.BalanceSettings, error) {
	balancesettings := &stripe.BalanceSettings{}
	err := c.B.Call(
		http.MethodPost, "/v1/balance_settings", c.Key, params, balancesettings)
	return balancesettings, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
