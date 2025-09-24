//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all BillSettingVersions by BillSetting ID.
type V2BillingBillSettingsVersionListParams struct {
	Params        `form:"*"`
	BillSettingID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a BillSettingVersion by ID.
type V2BillingBillSettingsVersionParams struct {
	Params        `form:"*"`
	BillSettingID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a BillSettingVersion by ID.
type V2BillingBillSettingsVersionRetrieveParams struct {
	Params        `form:"*"`
	BillSettingID *string `form:"-" json:"-"` // Included in URL
}
