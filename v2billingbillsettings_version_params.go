//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all BillSettingVersions by BillSetting ID.
type V2BillingBillSettingsVersionListParams struct {
	Params `form:"*"`
	// ID of the BillSettings to retrieve versions for.
	BillSettingID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a BillSettingVersion by ID.
type V2BillingBillSettingsVersionParams struct {
	Params `form:"*"`
	// The ID of the BillSetting object to retrieve the version from.
	BillSettingID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a BillSettingVersion by ID.
type V2BillingBillSettingsVersionRetrieveParams struct {
	Params `form:"*"`
	// The ID of the BillSetting object to retrieve the version from.
	BillSettingID *string `form:"-" json:"-"` // Included in URL
}
