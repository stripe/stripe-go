//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1BalanceSettingsService is used to invoke /v1/balance_settings APIs.
type v1BalanceSettingsService struct {
	B   Backend
	Key string
}

// Retrieves balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://stripe.com/connect/authentication)
func (c v1BalanceSettingsService) Retrieve(ctx context.Context, params *BalanceSettingsRetrieveParams) (*BalanceSettings, error) {
	balancesettings := &BalanceSettings{}
	if params == nil {
		params = &BalanceSettingsRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodGet, "/v1/balance_settings", c.Key, params, balancesettings)
	return balancesettings, err
}

// Updates balance settings for a given connected account.
//
//	Related guide: [Making API calls for connected accounts](https://stripe.com/connect/authentication)
func (c v1BalanceSettingsService) Update(ctx context.Context, params *BalanceSettingsUpdateParams) (*BalanceSettings, error) {
	balancesettings := &BalanceSettings{}
	if params == nil {
		params = &BalanceSettingsUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(
		http.MethodPost, "/v1/balance_settings", c.Key, params, balancesettings)
	return balancesettings, err
}
