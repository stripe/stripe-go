//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List license quantity changes for a contract given a license pricing ID.
type V2BillingContractsLicensePricingQuantityChangeListQuantityChangesParams struct {
	Params `form:"*"`
	// The ID of the license pricing.
	LicensePricingID *string `form:"-" json:"-"` // Included in URL
	// The ID of the contract.
	ContractID *string `form:"-" json:"-"` // Included in URL
	// The limit for the number of results per page.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}
