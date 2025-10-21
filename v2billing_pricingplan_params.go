//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Pricing Plan objects.
type V2BillingPricingPlanListParams struct {
	Params `form:"*"`
	// Filter for active/inactive PricingPlans. Mutually exclusive with `lookup_keys`.
	Active *bool `form:"active" json:"active,omitempty"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by lookup keys. Mutually exclusive with `active`.
	// You can specify up to 10 lookup keys.
	LookupKeys []*string `form:"lookup_keys,flat_array" json:"lookup_keys,omitempty"`
}

// Create a Pricing Plan object.
type V2BillingPricingPlanParams struct {
	Params `form:"*"`
	// Whether the PricingPlan is active.
	Active *bool `form:"active" json:"active,omitempty"`
	// The currency of the PricingPlan.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// Description of pricing plan subscription.
	Description *string `form:"description" json:"description,omitempty"`
	// Display name of the PricingPlan. Maximum 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// The ID of the live version of the PricingPlan.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// An internal key you can use to search for a particular PricingPlan. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The Stripe Tax tax behavior - whether the PricingPlan is inclusive or exclusive of tax.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingPricingPlanParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Create a Pricing Plan object.
type V2BillingPricingPlanCreateParams struct {
	Params `form:"*"`
	// The currency of the PricingPlan.
	Currency *string `form:"currency" json:"currency"`
	// Description of pricing plan subscription.
	Description *string `form:"description" json:"description,omitempty"`
	// Display name of the PricingPlan. Maximum 250 characters.
	DisplayName *string `form:"display_name" json:"display_name"`
	// An internal key you can use to search for a particular PricingPlan. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The Stripe Tax tax behavior - whether the PricingPlan is inclusive or exclusive of tax.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingPricingPlanCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a Pricing Plan object.
type V2BillingPricingPlanRetrieveParams struct {
	Params `form:"*"`
}

// Update a Pricing Plan object.
type V2BillingPricingPlanUpdateParams struct {
	Params `form:"*"`
	// Whether the PricingPlan is active.
	Active *bool `form:"active" json:"active,omitempty"`
	// Description of pricing plan subscription.
	Description *string `form:"description" json:"description,omitempty"`
	// Display name of the PricingPlan. Maximum 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// The ID of the live version of the PricingPlan.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// An internal key you can use to search for a particular PricingPlan. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingPricingPlanUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
