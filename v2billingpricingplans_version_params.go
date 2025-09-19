//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Pricing Plan Versions of a Pricing Plan.
type V2BillingPricingPlansVersionListParams struct {
	Params        `form:"*"`
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieve a specific Pricing Plan Version of a Pricing Plan.
type V2BillingPricingPlansVersionParams struct {
	Params        `form:"*"`
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a specific Pricing Plan Version of a Pricing Plan.
type V2BillingPricingPlansVersionRetrieveParams struct {
	Params        `form:"*"`
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
}
