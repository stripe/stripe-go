//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of subscription.
type V2BillingPricingPlanSubscriptionComponentsComponentType string

// List of values that V2BillingPricingPlanSubscriptionComponentsComponentType can take
const (
	V2BillingPricingPlanSubscriptionComponentsComponentTypeLicenseFeeSubscription V2BillingPricingPlanSubscriptionComponentsComponentType = "license_fee_subscription"
	V2BillingPricingPlanSubscriptionComponentsComponentTypeRateCardSubscription   V2BillingPricingPlanSubscriptionComponentsComponentType = "rate_card_subscription"
)

// The component subscriptions of the Pricing Plan Subscription.
type V2BillingPricingPlanSubscriptionComponentsComponent struct {
	// The ID of the License Fee Subscription.
	LicenseFeeSubscription string `json:"license_fee_subscription,omitempty"`
	// The Pricing Plan Component associated with this component subscription.
	PricingPlanComponent string `json:"pricing_plan_component"`
	// The ID of the Rate Card Subscription.
	RateCardSubscription string `json:"rate_card_subscription,omitempty"`
	// The type of subscription.
	Type V2BillingPricingPlanSubscriptionComponentsComponentType `json:"type"`
}

// A set of component subscriptions for a Pricing Plan Subscription.
type V2BillingPricingPlanSubscriptionComponents struct {
	APIResource
	// The component subscriptions of the Pricing Plan Subscription.
	Components []*V2BillingPricingPlanSubscriptionComponentsComponent `json:"components"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
}
