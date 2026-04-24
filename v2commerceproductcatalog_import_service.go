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

// v2CommerceProductCatalogImportService is used to invoke import related APIs.
type v2CommerceProductCatalogImportService struct {
	B   Backend
	Key string
}

// Creates a ProductCatalogImport.
func (c v2CommerceProductCatalogImportService) Create(ctx context.Context, params *V2CommerceProductCatalogImportCreateParams) (*V2CommerceProductCatalogImport, error) {
	if params == nil {
		params = &V2CommerceProductCatalogImportCreateParams{}
	}
	params.Context = ctx
	productcatalogimport := &V2CommerceProductCatalogImport{}
	err := c.B.Call(
		http.MethodPost, "/v2/commerce/product_catalog/imports", c.Key, params, productcatalogimport)
	return productcatalogimport, err
}

// Retrieves a ProductCatalogImport by ID.
func (c v2CommerceProductCatalogImportService) Retrieve(ctx context.Context, id string, params *V2CommerceProductCatalogImportRetrieveParams) (*V2CommerceProductCatalogImport, error) {
	if params == nil {
		params = &V2CommerceProductCatalogImportRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/commerce/product_catalog/imports/%s", id)
	productcatalogimport := &V2CommerceProductCatalogImport{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, productcatalogimport)
	return productcatalogimport, err
}

// Returns a list of ProductCatalogImport objects.
func (c v2CommerceProductCatalogImportService) List(ctx context.Context, listParams *V2CommerceProductCatalogImportListParams) *V2List[*V2CommerceProductCatalogImport] {
	if listParams == nil {
		listParams = &V2CommerceProductCatalogImportListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/commerce/product_catalog/imports", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2CommerceProductCatalogImport], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2CommerceProductCatalogImport]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
