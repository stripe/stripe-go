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

// v2BillingLicensedItemService is used to invoke licenseditem related APIs.
type v2BillingLicensedItemService struct {
	B   Backend
	Key string
}

// Create a Licensed Item object.
func (c v2BillingLicensedItemService) Create(ctx context.Context, params *V2BillingLicensedItemCreateParams) (*V2BillingLicensedItem, error) {
	if params == nil {
		params = &V2BillingLicensedItemCreateParams{}
	}
	params.Context = ctx
	licenseditem := &V2BillingLicensedItem{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/licensed_items", c.Key, params, licenseditem)
	return licenseditem, err
}

// Retrieve a Licensed Item object.
func (c v2BillingLicensedItemService) Retrieve(ctx context.Context, id string, params *V2BillingLicensedItemRetrieveParams) (*V2BillingLicensedItem, error) {
	if params == nil {
		params = &V2BillingLicensedItemRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/licensed_items/%s", id)
	licenseditem := &V2BillingLicensedItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, licenseditem)
	return licenseditem, err
}

// Update a Licensed Item object. At least one of the fields is required.
func (c v2BillingLicensedItemService) Update(ctx context.Context, id string, params *V2BillingLicensedItemUpdateParams) (*V2BillingLicensedItem, error) {
	if params == nil {
		params = &V2BillingLicensedItemUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/licensed_items/%s", id)
	licenseditem := &V2BillingLicensedItem{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, licenseditem)
	return licenseditem, err
}

// List all Licensed Item objects in reverse chronological order of creation.
func (c v2BillingLicensedItemService) List(ctx context.Context, listParams *V2BillingLicensedItemListParams) Seq2[*V2BillingLicensedItem, error] {
	if listParams == nil {
		listParams = &V2BillingLicensedItemListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/licensed_items", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingLicensedItem], error) {
		page := &V2Page[*V2BillingLicensedItem]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
