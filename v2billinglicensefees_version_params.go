//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all versions of a LicenseFee objects.
type V2BillingLicenseFeesVersionListParams struct {
	Params `form:"*"`
	// The ID of the LicenseFee to list versions for.
	LicenseFeeID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a LicenseFeeVersion object.
type V2BillingLicenseFeesVersionParams struct {
	Params `form:"*"`
	// The ID of the License fee object.
	LicenseFeeID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a LicenseFeeVersion object.
type V2BillingLicenseFeesVersionRetrieveParams struct {
	Params `form:"*"`
	// The ID of the License fee object.
	LicenseFeeID *string `form:"-" json:"-"` // Included in URL
}
