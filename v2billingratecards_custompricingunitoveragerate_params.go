//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Rate Card Custom Pricing Unit Overage Rates on a Rate Card.
type V2BillingRateCardsCustomPricingUnitOverageRateListParams struct {
	Params `form:"*"`
	// The ID of the Rate Card to list overage rates for.
	RateCardID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Optionally filter by a RateCard version. If not specified, defaults to the latest version.
	RateCardVersion *string `form:"rate_card_version" json:"rate_card_version,omitempty"`
}

// Create a Rate Card Custom Pricing Unit Overage Rate on a Rate Card.
type V2BillingRateCardsCustomPricingUnitOverageRateParams struct {
	Params `form:"*"`
	// The ID of the Rate Card to create the overage rate for.
	RateCardID *string `form:"-" json:"-"` // Included in URL
	// The ID of the custom pricing unit this overage rate applies to.
	CustomPricingUnit *string `form:"custom_pricing_unit" json:"custom_pricing_unit,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The ID of the one-time item to use for overage line items.
	OneTimeItem *string `form:"one_time_item" json:"one_time_item,omitempty"`
	// The per-unit amount to charge for overages, represented as a decimal string in minor currency units with at most 12 decimal places.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardsCustomPricingUnitOverageRateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Create a Rate Card Custom Pricing Unit Overage Rate on a Rate Card.
type V2BillingRateCardsCustomPricingUnitOverageRateCreateParams struct {
	Params `form:"*"`
	// The ID of the Rate Card to create the overage rate for.
	RateCardID *string `form:"-" json:"-"` // Included in URL
	// The ID of the custom pricing unit this overage rate applies to.
	CustomPricingUnit *string `form:"custom_pricing_unit" json:"custom_pricing_unit"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The ID of the one-time item to use for overage line items.
	OneTimeItem *string `form:"one_time_item" json:"one_time_item"`
	// The per-unit amount to charge for overages, represented as a decimal string in minor currency units with at most 12 decimal places.
	UnitAmount *string `form:"unit_amount" json:"unit_amount"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardsCustomPricingUnitOverageRateCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Delete a Rate Card Custom Pricing Unit Overage Rate from a Rate Card.
type V2BillingRateCardsCustomPricingUnitOverageRateDeleteParams struct {
	Params `form:"*"`
	// The ID of the Rate Card.
	RateCardID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a Rate Card Custom Pricing Unit Overage Rate from a Rate Card.
type V2BillingRateCardsCustomPricingUnitOverageRateRetrieveParams struct {
	Params `form:"*"`
	// The ID of the Rate Card.
	RateCardID *string `form:"-" json:"-"` // Included in URL
}
