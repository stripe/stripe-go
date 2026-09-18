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

// v2ProvisioningCatalogServiceService is used to invoke service related APIs.
type v2ProvisioningCatalogServiceService struct {
	B   Backend
	Key string
}

// Lists services available in the catalog.
func (c v2ProvisioningCatalogServiceService) List(ctx context.Context, listParams *V2ProvisioningCatalogServiceListParams) *V2List[*V2ProvisioningProviderServiceDetail] {
	if listParams == nil {
		listParams = &V2ProvisioningCatalogServiceListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/provisioning/catalog/services", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2ProvisioningProviderServiceDetail], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2ProvisioningProviderServiceDetail]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
