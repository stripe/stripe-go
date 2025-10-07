//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all CollectionSettingVersions by CollectionSetting ID.
type V2BillingCollectionSettingsVersionListParams struct {
	Params `form:"*"`
	// ID of the CollectionSettings to retrieve versions for.
	CollectionSettingID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a CollectionSetting Version by ID.
type V2BillingCollectionSettingsVersionParams struct {
	Params `form:"*"`
	// The ID of the CollectionSetting that has the version.
	CollectionSettingID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a CollectionSetting Version by ID.
type V2BillingCollectionSettingsVersionRetrieveParams struct {
	Params `form:"*"`
	// The ID of the CollectionSetting that has the version.
	CollectionSettingID *string `form:"-" json:"-"` // Included in URL
}
