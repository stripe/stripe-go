//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

type V2BillingPricingPlanVersion struct {
	APIResource
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// The timestamp when this version became inactive.
	EndDate time.Time `json:"end_date"`
	// Unique identifier for the PricingPlanVersion.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the PricingPlan this version belongs to.
	PricingPlan string `json:"pricing_plan"`
	// The timestamp when this version became active.
	StartDate time.Time `json:"start_date"`
}
