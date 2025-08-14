//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all versions of a PricingPlan.
type V2BillingPricingPlansVersionListParams struct {
	Params `form:"*"`
	// The ID of the PricingPlan to list versions for.
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a specific version of a PricingPlan.
type V2BillingPricingPlansVersionParams struct {
	Params `form:"*"`
	// The ID of the PricingPlan the version belongs to.
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a specific version of a PricingPlan.
type V2BillingPricingPlansVersionRetrieveParams struct {
	Params `form:"*"`
	// The ID of the PricingPlan the version belongs to.
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
}
