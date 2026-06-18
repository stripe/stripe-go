//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List quantity changes for a pricing line on a contract.
type V2BillingContractsPricingLinesQuantityChangeListContractPricingLineQuantityChangesParams struct {
	Params `form:"*"`
	// The ID of the pricing line.
	PricingLineID *string `form:"-" json:"-"` // Included in URL
	// The ID of the contract.
	ContractID *string `form:"-" json:"-"` // Included in URL
	// The limit for the number of results per page.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}
