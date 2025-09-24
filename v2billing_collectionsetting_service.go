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

// v2BillingCollectionSettingService is used to invoke collectionsetting related APIs.
type v2BillingCollectionSettingService struct {
	B   Backend
	Key string
}

// Create a CollectionSetting object.
func (c v2BillingCollectionSettingService) Create(ctx context.Context, params *V2BillingCollectionSettingCreateParams) (*V2BillingCollectionSetting, error) {
	if params == nil {
		params = &V2BillingCollectionSettingCreateParams{}
	}
	params.Context = ctx
	collectionsetting := &V2BillingCollectionSetting{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/collection_settings", c.Key, params, collectionsetting)
	return collectionsetting, err
}

// Retrieve a CollectionSetting by ID.
func (c v2BillingCollectionSettingService) Retrieve(ctx context.Context, id string, params *V2BillingCollectionSettingRetrieveParams) (*V2BillingCollectionSetting, error) {
	if params == nil {
		params = &V2BillingCollectionSettingRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/collection_settings/%s", id)
	collectionsetting := &V2BillingCollectionSetting{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, collectionsetting)
	return collectionsetting, err
}

// Update fields on an existing CollectionSetting.
func (c v2BillingCollectionSettingService) Update(ctx context.Context, id string, params *V2BillingCollectionSettingUpdateParams) (*V2BillingCollectionSetting, error) {
	if params == nil {
		params = &V2BillingCollectionSettingUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/collection_settings/%s", id)
	collectionsetting := &V2BillingCollectionSetting{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, collectionsetting)
	return collectionsetting, err
}

// List all CollectionSetting objects.
func (c v2BillingCollectionSettingService) List(ctx context.Context, listParams *V2BillingCollectionSettingListParams) Seq2[*V2BillingCollectionSetting, error] {
	if listParams == nil {
		listParams = &V2BillingCollectionSettingListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/billing/collection_settings", listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingCollectionSetting], error) {
		page := &V2Page[*V2BillingCollectionSetting]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
