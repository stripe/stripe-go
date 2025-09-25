//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Filter by payer. Mutually exclusive with `billing_cadence`, `pricing_plan`, and `pricing_plan_version`.
type V2BillingPricingPlanSubscriptionListPayerParams struct {
	// The ID of the Customer object. If provided, only Pricing Plan Subscriptions that are subscribed on the cadences with the specified payer will be returned.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// A string identifying the type of the payer. Currently the only supported value is `customer`.
	Type *string `form:"type" json:"type"`
}

// List all Pricing Plan Subscription objects.
type V2BillingPricingPlanSubscriptionListParams struct {
	Params `form:"*"`
	// Filter by Billing Cadence ID. Mutually exclusive with `payer`, `pricing_plan`, and `pricing_plan_version`.
	BillingCadence *string `form:"billing_cadence" json:"billing_cadence,omitempty"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by payer. Mutually exclusive with `billing_cadence`, `pricing_plan`, and `pricing_plan_version`.
	Payer *V2BillingPricingPlanSubscriptionListPayerParams `form:"payer" json:"payer,omitempty"`
	// Filter by PricingPlan ID. Mutually exlcusive with `billing_cadence`, `payer`, and `pricing_plan_version`.
	PricingPlan *string `form:"pricing_plan" json:"pricing_plan,omitempty"`
	// Filter by Pricing Plan Version ID. Mutually exlcusive with `billing_cadence`, `payer`, and `pricing_plan`.
	PricingPlanVersion *string `form:"pricing_plan_version" json:"pricing_plan_version,omitempty"`
	// Filter by servicing status.
	ServicingStatus *string `form:"servicing_status" json:"servicing_status,omitempty"`
}

// Retrieve a Pricing Plan Subscription object.
type V2BillingPricingPlanSubscriptionParams struct {
	Params `form:"*"`
	// When set to true, the `servicing_status_transition.will_cancel_at` field will be cleared.
	ClearCancelAt *bool `form:"clear_cancel_at" json:"clear_cancel_at,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingPricingPlanSubscriptionParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a Pricing Plan Subscription object.
type V2BillingPricingPlanSubscriptionRetrieveParams struct {
	Params `form:"*"`
}

// Update a Pricing Plan Subscription object.
type V2BillingPricingPlanSubscriptionUpdateParams struct {
	Params `form:"*"`
	// When set to true, the `servicing_status_transition.will_cancel_at` field will be cleared.
	ClearCancelAt *bool `form:"clear_cancel_at" json:"clear_cancel_at,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingPricingPlanSubscriptionUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
