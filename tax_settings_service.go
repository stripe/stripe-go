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

// v1TaxSettingsService is used to invoke /v1/tax/settings APIs.
type v1TaxSettingsService struct {
	B   Backend
	Key string
}

// Retrieves Tax Settings for a merchant.
func (c v1TaxSettingsService) Retrieve(ctx context.Context, params *TaxSettingsRetrieveParams) (*TaxSettings, error) {
	if params == nil {
		params = &TaxSettingsRetrieveParams{}
	}
	params.Context = ctx
	settings := &TaxSettings{}
	err := c.B.Call(http.MethodGet, "/v1/tax/settings", c.Key, params, settings)
	return settings, err
}

// Updates Tax Settings parameters used in tax calculations. All parameters are editable but none can be removed once set.
func (c v1TaxSettingsService) Update(ctx context.Context, params *TaxSettingsUpdateParams) (*TaxSettings, error) {
	if params == nil {
		params = &TaxSettingsUpdateParams{}
	}
	params.Context = ctx
	settings := &TaxSettings{}
	err := c.B.Call(http.MethodPost, "/v1/tax/settings", c.Key, params, settings)
	return settings, err
}
