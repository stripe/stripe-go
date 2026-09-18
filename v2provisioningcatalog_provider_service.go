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

// v2ProvisioningCatalogProviderService is used to invoke provider related APIs.
type v2ProvisioningCatalogProviderService struct {
	B   Backend
	Key string
}

// Lists providers available in the catalog.
func (c v2ProvisioningCatalogProviderService) List(ctx context.Context, listParams *V2ProvisioningCatalogProviderListParams) *V2List[*V2ProvisioningProvider] {
	if listParams == nil {
		listParams = &V2ProvisioningCatalogProviderListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/provisioning/catalog/providers", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2ProvisioningProvider], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2ProvisioningProvider]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
