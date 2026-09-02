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

// v2TaxIntegrationConfigurationService is used to invoke integrationconfiguration related APIs.
type v2TaxIntegrationConfigurationService struct {
	B   Backend
	Key string
}

// Retrieve the tax integration configuration for this account.
func (c v2TaxIntegrationConfigurationService) Retrieve(ctx context.Context, params *V2TaxIntegrationConfigurationRetrieveParams) (*V2TaxIntegrationConfiguration, error) {
	if params == nil {
		params = &V2TaxIntegrationConfigurationRetrieveParams{}
	}
	params.Context = ctx
	integrationconfiguration := &V2TaxIntegrationConfiguration{}
	err := c.B.Call(
		http.MethodGet, "/v2/tax/integration_configurations", c.Key, params, integrationconfiguration)
	return integrationconfiguration, err
}

// Update the tax integration configuration for this account.
func (c v2TaxIntegrationConfigurationService) Update(ctx context.Context, params *V2TaxIntegrationConfigurationUpdateParams) (*V2TaxIntegrationConfiguration, error) {
	if params == nil {
		params = &V2TaxIntegrationConfigurationUpdateParams{}
	}
	params.Context = ctx
	integrationconfiguration := &V2TaxIntegrationConfiguration{}
	err := c.B.Call(
		http.MethodPost, "/v2/tax/integration_configurations", c.Key, params, integrationconfiguration)
	return integrationconfiguration, err
}
