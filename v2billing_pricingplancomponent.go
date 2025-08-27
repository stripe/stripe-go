//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of the PricingPlanComponent.
type V2BillingPricingPlanComponentType string

// List of values that V2BillingPricingPlanComponentType can take
const (
	V2BillingPricingPlanComponentTypeLicenseFee    V2BillingPricingPlanComponentType = "license_fee"
	V2BillingPricingPlanComponentTypeRateCard      V2BillingPricingPlanComponentType = "rate_card"
	V2BillingPricingPlanComponentTypeServiceAction V2BillingPricingPlanComponentType = "service_action"
)

// Details if this component is a License Fee.
type V2BillingPricingPlanComponentLicenseFee struct {
	// The ID of the License Fee.
	ID string `json:"id"`
	// The version of the LicenseFee. Defaults to 'latest', if not specified.
	Version string `json:"version"`
}

// Details if this component is a Rate Card.
type V2BillingPricingPlanComponentRateCard struct {
	// The ID of the Rate Card.
	ID string `json:"id"`
	// The version of the RateCard. Defaults to 'latest', if not specified.
	Version string `json:"version"`
}

// Details if this component is a Service Action.
type V2BillingPricingPlanComponentServiceAction struct {
	// The ID of the service action.
	ID string `json:"id"`
}
type V2BillingPricingPlanComponent struct {
	APIResource
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Details if this component is a License Fee.
	LicenseFee *V2BillingPricingPlanComponentLicenseFee `json:"license_fee"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// An internal key you can use to search for a particular PricingPlanComponent.
	LookupKey string `json:"lookup_key"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the Pricing Plan this component belongs to.
	PricingPlan string `json:"pricing_plan"`
	// The ID of the Pricing Plan Version this component belongs to.
	PricingPlanVersion string `json:"pricing_plan_version"`
	// Details if this component is a Rate Card.
	RateCard *V2BillingPricingPlanComponentRateCard `json:"rate_card"`
	// Details if this component is a Service Action.
	ServiceAction *V2BillingPricingPlanComponentServiceAction `json:"service_action"`
	// The type of the PricingPlanComponent.
	Type V2BillingPricingPlanComponentType `json:"type"`
}
