//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The Stripe Tax tax behavior - whether the PricingPlan is inclusive or exclusive of tax.
type V2BillingPricingPlanTaxBehavior string

// List of values that V2BillingPricingPlanTaxBehavior can take
const (
	V2BillingPricingPlanTaxBehaviorExclusive V2BillingPricingPlanTaxBehavior = "exclusive"
	V2BillingPricingPlanTaxBehaviorInclusive V2BillingPricingPlanTaxBehavior = "inclusive"
)

type V2BillingPricingPlan struct {
	APIResource
	// Whether the PricingPlan is active.
	Active bool `json:"active"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// The currency of the PricingPlan.
	Currency Currency `json:"currency"`
	// A description for pricing plan subscription.
	// Maximum length of 500 characters.
	Description string `json:"description"`
	// Display name of the PricingPlan.
	DisplayName string `json:"display_name"`
	// Unique identifier for the PricingPlan.
	ID string `json:"id"`
	// The ID of the latest version of the PricingPlan.
	LatestVersion string `json:"latest_version"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The ID of the live version of the PricingPlan.
	LiveVersion string `json:"live_version"`
	// An internal key you can use to search for a particular PricingPlan. Maximum length of 200 characters.
	LookupKey string `json:"lookup_key"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The Stripe Tax tax behavior - whether the PricingPlan is inclusive or exclusive of tax.
	TaxBehavior V2BillingPricingPlanTaxBehavior `json:"tax_behavior"`
}
