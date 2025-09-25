//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all versions of a License Fee object.
type V2BillingLicenseFeesVersionListParams struct {
	Params       `form:"*"`
	LicenseFeeID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a License Fee Version object.
type V2BillingLicenseFeesVersionParams struct {
	Params       `form:"*"`
	LicenseFeeID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a License Fee Version object.
type V2BillingLicenseFeesVersionRetrieveParams struct {
	Params       `form:"*"`
	LicenseFeeID *string `form:"-" json:"-"` // Included in URL
}
