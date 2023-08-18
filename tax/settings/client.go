//
//
// File generated from our OpenAPI spec
//
//

// Package settings provides the /tax/settings APIs
package settings

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
)

// Client is used to invoke /tax/settings APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a tax settings.
func Get(params *stripe.TaxSettingsParams) (*stripe.TaxSettings, error) {
	return getC().Get(params)
}

// Get returns the details of a tax settings.
func (c Client) Get(params *stripe.TaxSettingsParams) (*stripe.TaxSettings, error) {
	settings := &stripe.TaxSettings{}
	err := c.B.Call(http.MethodGet, "/v1/tax/settings", c.Key, params, settings)
	return settings, err
}

// Update updates a tax settings's properties.
func Update(params *stripe.TaxSettingsParams) (*stripe.TaxSettings, error) {
	return getC().Update(params)
}

// Update updates a tax settings's properties.
func (c Client) Update(params *stripe.TaxSettingsParams) (*stripe.TaxSettings, error) {
	settings := &stripe.TaxSettings{}
	err := c.B.Call(http.MethodPost, "/v1/tax/settings", c.Key, params, settings)
	return settings, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
