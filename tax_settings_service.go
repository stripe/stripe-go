//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"encoding/json"
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

// Serializes a Settings update request into a batch job JSONL line.
func (c v1TaxSettingsService) MarshalBatchUpdate(params *TaxSettingsUpdateParams) (string, error) {
	itemID, err := newUUID4()
	if err != nil {
		return "", err
	}

	item := struct {
		ID            string            `json:"id"`
		Context       string            `json:"context,omitempty"`
		StripeVersion string            `json:"stripe_version,omitempty"`
		PathParams    map[string]string `json:"path_params,omitempty"`
		Params        interface{}       `json:"params"`
	}{
		ID:            itemID,
		PathParams:    nil,
		StripeVersion: APIVersion,
	}
	if params != nil {
		item.Params = params
		if params.StripeContext != nil {
			item.Context = *params.StripeContext
		}
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
