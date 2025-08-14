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

// Details if this component is a LicenseFee.
type V2BillingPricingPlanComponentLicenseFee struct {
	// The ID of the LicenseFee.
	ID string `json:"id"`
	// The version of the LicenseFee.
	Version string `json:"version"`
}

// Details if this component is a RateCard.
type V2BillingPricingPlanComponentRateCard struct {
	// The ID of the RateCard.
	ID string `json:"id"`
	// The version of the RateCard.
	Version string `json:"version"`
}

// Details if this component is a ServiceAction.
type V2BillingPricingPlanComponentServiceAction struct {
	// The ID of the ServiceAction.
	ID string `json:"id"`
	// The version of the ServiceAction.
	Version string `json:"version"`
}
type V2BillingPricingPlanComponent struct {
	APIResource
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Unique identifier for the PricingPlanComponent.
	ID string `json:"id"`
	// Details if this component is a LicenseFee.
	LicenseFee *V2BillingPricingPlanComponentLicenseFee `json:"license_fee"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// An internal key you can use to search for a particular PricingPlanComponent.
	LookupKey string `json:"lookup_key"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the PricingPlan this component belongs to.
	PricingPlan string `json:"pricing_plan"`
	// The ID of the PricingPlanVersion this component belongs to.
	PricingPlanVersion string `json:"pricing_plan_version"`
	// Details if this component is a RateCard.
	RateCard *V2BillingPricingPlanComponentRateCard `json:"rate_card"`
	// Details if this component is a ServiceAction.
	ServiceAction *V2BillingPricingPlanComponentServiceAction `json:"service_action"`
	// The type of the PricingPlanComponent.
	Type V2BillingPricingPlanComponentType `json:"type"`
}
