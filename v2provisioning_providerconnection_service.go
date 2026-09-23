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

// v2ProvisioningProviderConnectionService is used to invoke providerconnection related APIs.
type v2ProvisioningProviderConnectionService struct {
	B   Backend
	Key string
}

// Unlinks a provider connection so it can no longer be used to create resources.
func (c v2ProvisioningProviderConnectionService) Unlink(ctx context.Context, id string, params *V2ProvisioningProviderConnectionUnlinkParams) (*V2ProvisioningProviderConnection, error) {
	if params == nil {
		params = &V2ProvisioningProviderConnectionUnlinkParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/provisioning/provider_connections/%s/unlink", id)
	providerconnection := &V2ProvisioningProviderConnection{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, providerconnection)
	return providerconnection, err
}

// Lists the provider connections for the account.
func (c v2ProvisioningProviderConnectionService) List(ctx context.Context, listParams *V2ProvisioningProviderConnectionListParams) *V2List[*V2ProvisioningProviderConnection] {
	if listParams == nil {
		listParams = &V2ProvisioningProviderConnectionListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/provisioning/provider_connections", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2ProvisioningProviderConnection], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2ProvisioningProviderConnection]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
