//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Rates associated with a RateCard for a specific version (defaults to latest). Rates remain active for all subsequent versions until a new Rate is created for the same MeteredItem.
type V2BillingRateCardsRateListParams struct {
	Params `form:"*"`
	// The ID of the RateCard to retrieve rates for.
	RateCardID *string `form:"-" json:"-"` // Included in URL
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Optionally filter by a MeteredItem.
	MeteredItem *string `form:"metered_item" json:"metered_item,omitempty"`
	// Optionally filter by a RateCard version. If not specified, defaults to the latest version.
	RateCardVersion *string `form:"rate_card_version" json:"rate_card_version,omitempty"`
}

// The custom pricing unit that this rate binds to.
type V2BillingRateCardsRateCustomPricingUnitAmountParams struct {
	// The id of the custom pricing unit.
	ID *string `form:"id" json:"id"`
	// The unit value for the custom pricing unit, as a string.
	Value *string `form:"value" json:"value"`
}

// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
type V2BillingRateCardsRateTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units with at most 12 decimal places.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units with at
	// most 12 decimal places.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier. Only one of `up_to_decimal` and `up_to_inf` may
	// be set.
	UpToDecimal *string `form:"up_to_decimal" json:"up_to_decimal,omitempty"`
	// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed.
type V2BillingRateCardsRateTransformQuantityParams struct {
	// Divide usage by this number.
	DivideBy *int64 `form:"divide_by" json:"divide_by"`
	// After division, round the result up or down.
	Round *string `form:"round" json:"round"`
}

// Set the rate for a MeteredItem on the latest version of a RateCard object. This will create a new RateCard version
// if the MeteredItem already has a rate on the RateCard.
type V2BillingRateCardsRateParams struct {
	Params `form:"*"`
	// The ID of the RateCard to create a new rate for.
	RateCardID *string `form:"-" json:"-"` // Included in URL
	// The custom pricing unit that this rate binds to.
	CustomPricingUnitAmount *V2BillingRateCardsRateCustomPricingUnitAmountParams `form:"custom_pricing_unit_amount" json:"custom_pricing_unit_amount,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The MeteredItem that this rate binds to.
	MeteredItem *string `form:"metered_item" json:"metered_item,omitempty"`
	// The ID of the Price object to take price information from. The Price must have the same interval as the RateCard.
	// Updates to the Price will not be reflected in the RateCard or its rates.
	Price *string `form:"price" json:"price,omitempty"`
	// Defines whether the tiered price should be graduated or volume-based. In volume-based tiering, the maximum
	// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
	// grows into new tiers. Can only be set if `tiers` is set.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
	Tiers []*V2BillingRateCardsRateTierParams `form:"tiers" json:"tiers,omitempty"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed.
	TransformQuantity *V2BillingRateCardsRateTransformQuantityParams `form:"transform_quantity" json:"transform_quantity,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units with at most 12 decimal
	// places. Cannot be set if `tiers` is provided.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardsRateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The custom pricing unit that this rate binds to.
type V2BillingRateCardsRateCreateCustomPricingUnitAmountParams struct {
	// The id of the custom pricing unit.
	ID *string `form:"id" json:"id"`
	// The unit value for the custom pricing unit, as a string.
	Value *string `form:"value" json:"value"`
}

// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
type V2BillingRateCardsRateCreateTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units with at most 12 decimal places.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units with at
	// most 12 decimal places.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier. Only one of `up_to_decimal` and `up_to_inf` may
	// be set.
	UpToDecimal *string `form:"up_to_decimal" json:"up_to_decimal,omitempty"`
	// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed.
type V2BillingRateCardsRateCreateTransformQuantityParams struct {
	// Divide usage by this number.
	DivideBy *int64 `form:"divide_by" json:"divide_by"`
	// After division, round the result up or down.
	Round *string `form:"round" json:"round"`
}

// Set the rate for a MeteredItem on the latest version of a RateCard object. This will create a new RateCard version
// if the MeteredItem already has a rate on the RateCard.
type V2BillingRateCardsRateCreateParams struct {
	Params `form:"*"`
	// The ID of the RateCard to create a new rate for.
	RateCardID *string `form:"-" json:"-"` // Included in URL
	// The custom pricing unit that this rate binds to.
	CustomPricingUnitAmount *V2BillingRateCardsRateCreateCustomPricingUnitAmountParams `form:"custom_pricing_unit_amount" json:"custom_pricing_unit_amount,omitempty"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The MeteredItem that this rate binds to.
	MeteredItem *string `form:"metered_item" json:"metered_item,omitempty"`
	// The ID of the Price object to take price information from. The Price must have the same interval as the RateCard.
	// Updates to the Price will not be reflected in the RateCard or its rates.
	Price *string `form:"price" json:"price,omitempty"`
	// Defines whether the tiered price should be graduated or volume-based. In volume-based tiering, the maximum
	// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
	// grows into new tiers. Can only be set if `tiers` is set.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
	Tiers []*V2BillingRateCardsRateCreateTierParams `form:"tiers" json:"tiers,omitempty"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed.
	TransformQuantity *V2BillingRateCardsRateCreateTransformQuantityParams `form:"transform_quantity" json:"transform_quantity,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units with at most 12 decimal
	// places. Cannot be set if `tiers` is provided.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardsRateCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Remove an existing Rate from a RateCard. This will create a new RateCard version without that rate.
type V2BillingRateCardsRateDeleteParams struct {
	Params `form:"*"`
	// The ID of the RateCard.
	RateCardID *string `form:"-" json:"-"` // Included in URL
}

// Retrieve a Rate object.
type V2BillingRateCardsRateRetrieveParams struct {
	Params `form:"*"`
	// The ID of the RateCard.
	RateCardID *string `form:"-" json:"-"` // Included in URL
}
