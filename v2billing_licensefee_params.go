//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all License Fee objects.
type V2BillingLicenseFeeListParams struct {
	Params `form:"*"`
	// Filter by licensed item.
	LicensedItem *string `form:"licensed_item" json:"licensed_item,omitempty"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by lookup keys.
	// You can specify up to 10 lookup keys.
	LookupKeys []*string `form:"lookup_keys" json:"lookup_keys"`
}

// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
type V2BillingLicenseFeeTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units with at most 12 decimal places.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units with at
	// most 12 decimal places.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier. Only one of `up_to_decimal` and `up_to_inf` may
	// be set.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,high_precision,omitempty"`
	// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed.
type V2BillingLicenseFeeTransformQuantityParams struct {
	// Divide usage by this number.
	DivideBy *int64 `form:"divide_by" json:"divide_by"`
	// After division, round the result up or down.
	Round *string `form:"round" json:"round"`
}

// Create a License Fee object.
type V2BillingLicenseFeeParams struct {
	Params `form:"*"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A customer-facing name for the License Fee.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// The Licensed Item that this License Fee binds to.
	LicensedItem *string `form:"licensed_item" json:"licensed_item,omitempty"`
	// Changes the version that new license fee will use. Providing `live_version = "latest"` will set the
	// license fee's `live_version` to its latest version.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// An internal key you can use to search for a particular license fee. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The interval for assessing service. For example, a monthly license fee with a rate of $1 for the first 10 "workloads"
	// and $2 thereafter means "$1 per workload up to 10 workloads during a month of service." This is similar to but
	// distinct from billing interval; the service interval deals with the rate at which the customer accumulates fees,
	// while the billing interval in Cadence deals with the rate the customer is billed.
	ServiceInterval *string `form:"service_interval" json:"service_interval,omitempty"`
	// The length of the interval for assessing service. For example, set this to 3 and `service_interval` to `"month"` in
	// order to specify quarterly service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count,omitempty"`
	// The Stripe Tax tax behavior - whether the license fee is inclusive or exclusive of tax.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior,omitempty"`
	// Defines whether the tiered price should be graduated or volume-based. In volume-based tiering, the maximum
	// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
	// grows into new tiers. Can only be set if `tiers` is set.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
	Tiers []*V2BillingLicenseFeeTierParams `form:"tiers" json:"tiers,omitempty"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed.
	TransformQuantity *V2BillingLicenseFeeTransformQuantityParams `form:"transform_quantity" json:"transform_quantity,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units with at most 12 decimal
	// places. Cannot be set if `tiers` is provided.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingLicenseFeeParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
type V2BillingLicenseFeeCreateTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units with at most 12 decimal places.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units with at
	// most 12 decimal places.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier. Only one of `up_to_decimal` and `up_to_inf` may
	// be set.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,high_precision,omitempty"`
	// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed.
type V2BillingLicenseFeeCreateTransformQuantityParams struct {
	// Divide usage by this number.
	DivideBy *int64 `form:"divide_by" json:"divide_by"`
	// After division, round the result up or down.
	Round *string `form:"round" json:"round"`
}

// Create a License Fee object.
type V2BillingLicenseFeeCreateParams struct {
	Params `form:"*"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency *string `form:"currency" json:"currency"`
	// A customer-facing name for the License Fee.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name"`
	// The Licensed Item that this License Fee binds to.
	LicensedItem *string `form:"licensed_item" json:"licensed_item"`
	// An internal key you can use to search for a particular license fee. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The interval for assessing service. For example, a monthly license fee with a rate of $1 for the first 10 "workloads"
	// and $2 thereafter means "$1 per workload up to 10 workloads during a month of service." This is similar to but
	// distinct from billing interval; the service interval deals with the rate at which the customer accumulates fees,
	// while the billing interval in Cadence deals with the rate the customer is billed.
	ServiceInterval *string `form:"service_interval" json:"service_interval"`
	// The length of the interval for assessing service. For example, set this to 3 and `service_interval` to `"month"` in
	// order to specify quarterly service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count"`
	// The Stripe Tax tax behavior - whether the license fee is inclusive or exclusive of tax.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior"`
	// Defines whether the tiered price should be graduated or volume-based. In volume-based tiering, the maximum
	// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
	// grows into new tiers. Can only be set if `tiers` is set.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
	Tiers []*V2BillingLicenseFeeCreateTierParams `form:"tiers" json:"tiers,omitempty"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed.
	TransformQuantity *V2BillingLicenseFeeCreateTransformQuantityParams `form:"transform_quantity" json:"transform_quantity,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units with at most 12 decimal
	// places. Cannot be set if `tiers` is provided.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingLicenseFeeCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve a License Fee object.
type V2BillingLicenseFeeRetrieveParams struct {
	Params `form:"*"`
}

// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
type V2BillingLicenseFeeUpdateTierParams struct {
	// Price for the entire tier, represented as a decimal string in minor currency units with at most 12 decimal places.
	FlatAmount *string `form:"flat_amount" json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units with at
	// most 12 decimal places.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier. Only one of `up_to_decimal` and `up_to_inf` may
	// be set.
	UpToDecimal *float64 `form:"up_to_decimal,high_precision" json:"up_to_decimal,high_precision,omitempty"`
	// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
	UpToInf *string `form:"up_to_inf" json:"up_to_inf,omitempty"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed.
type V2BillingLicenseFeeUpdateTransformQuantityParams struct {
	// Divide usage by this number.
	DivideBy *int64 `form:"divide_by" json:"divide_by"`
	// After division, round the result up or down.
	Round *string `form:"round" json:"round"`
}

// Update a License Fee object.
type V2BillingLicenseFeeUpdateParams struct {
	Params `form:"*"`
	// A customer-facing name for the License Fee.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name"`
	// Changes the version that new license fee will use. Providing `live_version = "latest"` will set the
	// license fee's `live_version` to its latest version.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// An internal key you can use to search for a particular license fee. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Defines whether the tiered price should be graduated or volume-based. In volume-based tiering, the maximum
	// quantity within a period determines the per-unit price. In graduated tiering, the pricing changes as the quantity
	// grows into new tiers. Can only be set if `tiers` is set.
	TieringMode *string `form:"tiering_mode" json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier. Cannot be set if `unit_amount` is provided.
	Tiers []*V2BillingLicenseFeeUpdateTierParams `form:"tiers" json:"tiers,omitempty"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed.
	TransformQuantity *V2BillingLicenseFeeUpdateTransformQuantityParams `form:"transform_quantity" json:"transform_quantity,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units with at most 12 decimal
	// places. Cannot be set if `tiers` is provided.
	UnitAmount *string `form:"unit_amount" json:"unit_amount,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingLicenseFeeUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
