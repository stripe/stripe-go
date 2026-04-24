//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Rate Card objects.
type V2BillingRateCardListParams struct {
	Params `form:"*"`
	// Optionally filter to active/inactive RateCards.
	Active *bool `form:"active" json:"active,omitempty"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by lookup keys.
	// You can specify up to 10 lookup keys.
	LookupKeys []*string `form:"lookup_keys" json:"lookup_keys,omitempty"`
}

// Create a Rate Card object.
type V2BillingRateCardParams struct {
	Params `form:"*"`
	// Sets whether the RateCard is active. Inactive RateCards cannot be used in new activations or have new rates added.
	Active *bool `form:"active" json:"active,omitempty"`
	// The currency of this RateCard.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A customer-facing name for the RateCard.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Changes the version that new RateCard activations use. Providing `live_version = "latest"` sets the
	// RateCard's `live_version` to its latest version.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// An internal key you can use to search for a particular RateCard. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
	// The interval for assessing service. For example, a monthly RateCard with a rate of 1 USD for the first 10 "workloads"
	// and 2 USD thereafter means "1 USD per workload up to 10 workloads during a month of service." This is similar to but
	// distinct from billing interval; the service interval deals with the rate at which the customer accumulates fees,
	// while the billing interval in Cadence deals with the rate the customer is billed.
	ServiceInterval *string `form:"service_interval" json:"service_interval,omitempty"`
	// The length of the interval for assessing service. For example, set this to 3 and `service_interval` to `"month"`
	// to specify quarterly service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count,omitempty"`
	// The tax behavior for Stripe Tax — whether the rate card price includes or excludes tax.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}

// Optional array of Meter segments to filter event dimension keys for billing.
type V2BillingRateCardModifyRatesRatesToCreateMeteredItemDataMeterSegmentConditionParams struct {
	// A Meter dimension.
	Dimension *string `form:"dimension" json:"dimension"`
	// To count usage towards this metered item, the dimension must have this value.
	Value *string `form:"value" json:"value"`
}

// The data to create a metered item that binds to this rate. Cannot be set if `metered_item` is provided, and must be set if it isn't.
type V2BillingRateCardModifyRatesRatesToCreateMeteredItemDataParams struct {
	// Description that customers see in the invoice line item.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name"`
	// An internal key you can use to search for a particular metered item.
	// Must be unique among metered items.
	// Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// ID of the Meter that measures usage for this Metered Item.
	Meter *string `form:"meter" json:"meter"`
	// Optional array of Meter segments to filter event dimension keys for billing.
	MeterSegmentConditions []*V2BillingRateCardModifyRatesRatesToCreateMeteredItemDataMeterSegmentConditionParams `form:"meter_segment_conditions" json:"meter_segment_conditions"`
	// The unit to use when displaying prices for this billable item in places like Checkout. For example, set this field
	// to "CPU-hour" for Checkout to display "(price) per CPU-hour", or "1 million events" to display "(price) per 1
	// million events".
	// Maximum length of 100 characters.
	UnitLabel *string `form:"unit_label" json:"unit_label,omitempty"`
}

// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
type V2BillingRateCardModifyRatesRatesToCreateTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units with at most 12 decimal places.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units with at
	// most 12 decimal places.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity is contained in the tier. Only one of `up_to_decimal` and `up_to_inf` may
	// be set.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed.
type V2BillingRateCardModifyRatesRatesToCreateTransformQuantityParams struct {
	// Divide usage by this number.
	DivideBy *int64 `form:"divide_by" json:"divide_by,string"`
	// After division, round the result up or down.
	Round *string `form:"round" json:"round"`
}

// The list of RateCard rates to create or update. Maximum of 100 rates.
type V2BillingRateCardModifyRatesRatesToCreateParams struct {
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The Metered Item that this rate binds to. Cannot be set if `metered_item_data` is provided, and must be set if it isn't.
	MeteredItem *string `form:"metered_item" json:"metered_item,omitempty"`
	// The data to create a metered item that binds to this rate. Cannot be set if `metered_item` is provided, and must be set if it isn't.
	MeteredItemData *V2BillingRateCardModifyRatesRatesToCreateMeteredItemDataParams `form:"metered_item_data" json:"metered_item_data,omitempty"`
	// Defines whether the tiered price is graduated or volume-based. In volume-based tiering, the maximum
	// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
	// grows into new tiers. Can only be set if `tiers` is set.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
	Tiers []*V2BillingRateCardModifyRatesRatesToCreateTierParams `form:"tiers" json:"tiers,omitempty"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed.
	TransformQuantity *V2BillingRateCardModifyRatesRatesToCreateTransformQuantityParams `form:"transform_quantity" json:"transform_quantity,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units with at most 12 decimal
	// places. Cannot be set if `tiers` is provided.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardModifyRatesRatesToCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The list of RateCard rates to delete. Maximum of 100 rates.
type V2BillingRateCardModifyRatesRatesToDeleteParams struct {
	// The ID of the RateCard rate to delete.
	ID *string `form:"id" json:"id"`
}

// Creates, updates, and/or deletes multiple Rates on a Rate Card atomically.
type V2BillingRateCardModifyRatesParams struct {
	Params `form:"*"`
	// The list of RateCard rates to create or update. Maximum of 100 rates.
	RatesToCreate []*V2BillingRateCardModifyRatesRatesToCreateParams `form:"rates_to_create" json:"rates_to_create"`
	// The list of RateCard rates to delete. Maximum of 100 rates.
	RatesToDelete []*V2BillingRateCardModifyRatesRatesToDeleteParams `form:"rates_to_delete" json:"rates_to_delete"`
}

// Create a Rate Card object.
type V2BillingRateCardCreateParams struct {
	Params `form:"*"`
	// The currency of this RateCard.
	Currency *string `form:"currency" json:"currency"`
	// A customer-facing name for the RateCard.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name"`
	// An internal key you can use to search for a particular RateCard. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The interval for assessing service. For example, a monthly RateCard with a rate of 1 USD for the first 10 "workloads"
	// and 2 USD thereafter means "1 USD per workload up to 10 workloads during a month of service." This is similar to but
	// distinct from billing interval; the service interval deals with the rate at which the customer accumulates fees,
	// while the billing interval in Cadence deals with the rate the customer is billed.
	ServiceInterval *string `form:"service_interval" json:"service_interval"`
	// The length of the interval for assessing service. For example, set this to 3 and `service_interval` to `"month"`
	// to specify quarterly service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count"`
	// The tax behavior for Stripe Tax — whether the rate card price includes or excludes tax.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve the latest version of a Rate Card object.
type V2BillingRateCardRetrieveParams struct {
	Params `form:"*"`
}

// Update a Rate Card object.
type V2BillingRateCardUpdateParams struct {
	Params `form:"*"`
	// Sets whether the RateCard is active. Inactive RateCards cannot be used in new activations or have new rates added.
	Active *bool `form:"active" json:"active,omitempty"`
	// A customer-facing name for the RateCard.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Changes the version that new RateCard activations use. Providing `live_version = "latest"` sets the
	// RateCard's `live_version` to its latest version.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// An internal key you can use to search for a particular RateCard. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]*string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardUpdateParams) AddMetadata(key string, value *string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]*string)
	}

	p.Metadata[key] = value
}
