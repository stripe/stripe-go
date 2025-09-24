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

// v2BillingCollectionSettingsVersionService is used to invoke version related APIs.
type v2BillingCollectionSettingsVersionService struct {
	B   Backend
	Key string
}

// Retrieve a CollectionSetting Version by ID.
func (c v2BillingCollectionSettingsVersionService) Retrieve(ctx context.Context, id string, params *V2BillingCollectionSettingsVersionRetrieveParams) (*V2BillingCollectionSettingVersion, error) {
	if params == nil {
		params = &V2BillingCollectionSettingsVersionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/billing/collection_settings/%s/versions/%s", StringValue(
			params.CollectionSettingID), id)
	collectionsettingversion := &V2BillingCollectionSettingVersion{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, collectionsettingversion)
	return collectionsettingversion, err
}

// List all CollectionSettingVersions by CollectionSetting ID.
func (c v2BillingCollectionSettingsVersionService) List(ctx context.Context, listParams *V2BillingCollectionSettingsVersionListParams) Seq2[*V2BillingCollectionSettingVersion, error] {
	if listParams == nil {
		listParams = &V2BillingCollectionSettingsVersionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/billing/collection_settings/%s/versions", StringValue(
			listParams.CollectionSettingID))
	return NewV2List(path, listParams, func(path string, p ParamsContainer) (*V2Page[*V2BillingCollectionSettingVersion], error) {
		page := &V2Page[*V2BillingCollectionSettingVersion]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
