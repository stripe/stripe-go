//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Pricing Plan Component objects for a Pricing Plan.
type V2BillingPricingPlansComponentListParams struct {
	Params `form:"*"`
	// The ID of the Pricing Plan to list components for.
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by lookup keys. Mutually exclusive with `pricing_plan_version`.
	// You can specify up to 10 lookup keys.
	LookupKeys []*string `form:"lookup_keys,flat_array" json:"lookup_keys,omitempty"`
	// The ID of the Pricing Plan Version to list components for. Will use the latest version if not provided.
	// Mutually exclusive with `lookup_keys`.
	PricingPlanVersion *string `form:"pricing_plan_version" json:"pricing_plan_version,omitempty"`
}

// Details if this component is a License Fee.
type V2BillingPricingPlansComponentLicenseFeeParams struct {
	// The ID of the License Fee.
	ID *string `form:"id" json:"id"`
	// The version of the LicenseFee. Defaults to 'latest', if not specified.
	Version *string `form:"version" json:"version,omitempty"`
}

// Details if this component is a Rate Card.
type V2BillingPricingPlansComponentRateCardParams struct {
	// The ID of the Rate Card.
	ID *string `form:"id" json:"id"`
	// The version of the RateCard. Defaults to 'latest', if not specified.
	Version *string `form:"version" json:"version,omitempty"`
}

// Details if this component is a Service Action.
type V2BillingPricingPlansComponentServiceActionParams struct {
	// The ID of the service action.
	ID *string `form:"id" json:"id"`
}

// Create a Pricing Plan Component object.
type V2BillingPricingPlansComponentParams struct {
	Params `form:"*"`
	// The ID of the Pricing Plan the component belongs to.
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
	// Details if this component is a License Fee.
	LicenseFee *V2BillingPricingPlansComponentLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// An identifier that can be used to find this component. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Details if this component is a Rate Card.
	RateCard *V2BillingPricingPlansComponentRateCardParams `form:"rate_card" json:"rate_card,omitempty"`
	// Details if this component is a Service Action.
	ServiceAction *V2BillingPricingPlansComponentServiceActionParams `form:"service_action" json:"service_action,omitempty"`
	// The type of the PricingPlanComponent.
	Type *string `form:"type" json:"type,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingPricingPlansComponentParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details if this component is a License Fee.
type V2BillingPricingPlansComponentCreateLicenseFeeParams struct {
	// The ID of the License Fee.
	ID *string `form:"id" json:"id"`
	// The version of the LicenseFee. Defaults to 'latest', if not specified.
	Version *string `form:"version" json:"version,omitempty"`
}

// Details if this component is a Rate Card.
type V2BillingPricingPlansComponentCreateRateCardParams struct {
	// The ID of the Rate Card.
	ID *string `form:"id" json:"id"`
	// The version of the RateCard. Defaults to 'latest', if not specified.
	Version *string `form:"version" json:"version,omitempty"`
}

// Details if this component is a Service Action.
type V2BillingPricingPlansComponentCreateServiceActionParams struct {
	// The ID of the service action.
	ID *string `form:"id" json:"id"`
}

// Create a Pricing Plan Component object.
type V2BillingPricingPlansComponentCreateParams struct {
	Params `form:"*"`
	// The ID of the PricingPlan to add the component to.
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
	// Details if this component is a License Fee.
	LicenseFee *V2BillingPricingPlansComponentCreateLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// An identifier that can be used to find this component.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Details if this component is a Rate Card.
	RateCard *V2BillingPricingPlansComponentCreateRateCardParams `form:"rate_card" json:"rate_card,omitempty"`
	// Details if this component is a Service Action.
	ServiceAction *V2BillingPricingPlansComponentCreateServiceActionParams `form:"service_action" json:"service_action,omitempty"`
	// The type of the PricingPlanComponent.
	Type *string `form:"type" json:"type"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingPricingPlansComponentCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Remove a Pricing Plan Component from the latest version of a Pricing Plan.
type V2BillingPricingPlansComponentDeleteParams struct {
	Params `form:"*"`
	// The ID of the Pricing Plan.
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a Pricing Plan Component object.
type V2BillingPricingPlansComponentRetrieveParams struct {
	Params `form:"*"`
	// The ID of the PricingPlan the component belongs to.
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
}

// Update a Pricing Plan Component object.
type V2BillingPricingPlansComponentUpdateParams struct {
	Params `form:"*"`
	// The ID of the Pricing Plan the component belongs to.
	PricingPlanID *string `form:"-" json:"-"` // Included in URL
	// An identifier that can be used to find this component. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingPricingPlansComponentUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
