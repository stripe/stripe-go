//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The interval for assessing service.
type V2BillingServiceActionServiceInterval string

// List of values that V2BillingServiceActionServiceInterval can take
const (
	V2BillingServiceActionServiceIntervalDay   V2BillingServiceActionServiceInterval = "day"
	V2BillingServiceActionServiceIntervalMonth V2BillingServiceActionServiceInterval = "month"
	V2BillingServiceActionServiceIntervalWeek  V2BillingServiceActionServiceInterval = "week"
	V2BillingServiceActionServiceIntervalYear  V2BillingServiceActionServiceInterval = "year"
)

// The type of the service action.
type V2BillingServiceActionType string

// List of values that V2BillingServiceActionType can take
const (
	V2BillingServiceActionTypeCreditGrant          V2BillingServiceActionType = "credit_grant"
	V2BillingServiceActionTypeCreditGrantPerTenant V2BillingServiceActionType = "credit_grant_per_tenant"
)

// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
type V2BillingServiceActionCreditGrantAmountType string

// List of values that V2BillingServiceActionCreditGrantAmountType can take
const (
	V2BillingServiceActionCreditGrantAmountTypeCustomPricingUnit V2BillingServiceActionCreditGrantAmountType = "custom_pricing_unit"
	V2BillingServiceActionCreditGrantAmountTypeMonetary          V2BillingServiceActionCreditGrantAmountType = "monetary"
)

// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
type V2BillingServiceActionCreditGrantApplicabilityConfigScopePriceType string

// List of values that V2BillingServiceActionCreditGrantApplicabilityConfigScopePriceType can take
const (
	V2BillingServiceActionCreditGrantApplicabilityConfigScopePriceTypeMetered V2BillingServiceActionCreditGrantApplicabilityConfigScopePriceType = "metered"
)

// The type of the expiry configuration. We currently support `end_of_service_period`.
type V2BillingServiceActionCreditGrantExpiryConfigType string

// List of values that V2BillingServiceActionCreditGrantExpiryConfigType can take
const (
	V2BillingServiceActionCreditGrantExpiryConfigTypeEndOfServicePeriod V2BillingServiceActionCreditGrantExpiryConfigType = "end_of_service_period"
)

// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
type V2BillingServiceActionCreditGrantPerTenantAmountType string

// List of values that V2BillingServiceActionCreditGrantPerTenantAmountType can take
const (
	V2BillingServiceActionCreditGrantPerTenantAmountTypeCustomPricingUnit V2BillingServiceActionCreditGrantPerTenantAmountType = "custom_pricing_unit"
	V2BillingServiceActionCreditGrantPerTenantAmountTypeMonetary          V2BillingServiceActionCreditGrantPerTenantAmountType = "monetary"
)

// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
type V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigScopePriceType string

// List of values that V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigScopePriceType can take
const (
	V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigScopePriceTypeMetered V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigScopePriceType = "metered"
)

// The type of the expiry configuration. We currently support `end_of_service_period`.
type V2BillingServiceActionCreditGrantPerTenantExpiryConfigType string

// List of values that V2BillingServiceActionCreditGrantPerTenantExpiryConfigType can take
const (
	V2BillingServiceActionCreditGrantPerTenantExpiryConfigTypeEndOfServicePeriod V2BillingServiceActionCreditGrantPerTenantExpiryConfigType = "end_of_service_period"
)

// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
type V2BillingServiceActionCreditGrantAmountCustomPricingUnit struct {
	// The id of the custom pricing unit.
	ID string `json:"id"`
	// The value of the credit grant, decimal value represented as a string.
	Value float64 `json:"value,string"`
}

// The amount of the credit grant.
type V2BillingServiceActionCreditGrantAmount struct {
	// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
	CustomPricingUnit *V2BillingServiceActionCreditGrantAmountCustomPricingUnit `json:"custom_pricing_unit,omitempty"`
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary Amount `json:"monetary,omitempty"`
	// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
	Type V2BillingServiceActionCreditGrantAmountType `json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingServiceActionCreditGrantApplicabilityConfigScope struct {
	// The billable items to apply the credit grant to.
	BillableItems []string `json:"billable_items,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
	PriceType V2BillingServiceActionCreditGrantApplicabilityConfigScopePriceType `json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingServiceActionCreditGrantApplicabilityConfig struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingServiceActionCreditGrantApplicabilityConfigScope `json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingServiceActionCreditGrantExpiryConfig struct {
	// The type of the expiry configuration. We currently support `end_of_service_period`.
	Type V2BillingServiceActionCreditGrantExpiryConfigType `json:"type"`
}

// Details for the credit grant. Provided only if `type` is "credit_grant".
type V2BillingServiceActionCreditGrant struct {
	// The amount of the credit grant.
	Amount *V2BillingServiceActionCreditGrantAmount `json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingServiceActionCreditGrantApplicabilityConfig `json:"applicability_config"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingServiceActionCreditGrantExpiryConfig `json:"expiry_config"`
	// A descriptive name shown in dashboard.
	Name string `json:"name"`
}

// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
type V2BillingServiceActionCreditGrantPerTenantAmountCustomPricingUnit struct {
	// The id of the custom pricing unit.
	ID string `json:"id"`
	// The value of the credit grant, decimal value represented as a string.
	Value float64 `json:"value,string"`
}

// The amount of the credit grant.
type V2BillingServiceActionCreditGrantPerTenantAmount struct {
	// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
	CustomPricingUnit *V2BillingServiceActionCreditGrantPerTenantAmountCustomPricingUnit `json:"custom_pricing_unit,omitempty"`
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary Amount `json:"monetary,omitempty"`
	// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
	Type V2BillingServiceActionCreditGrantPerTenantAmountType `json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigScope struct {
	// The billable items to apply the credit grant to.
	BillableItems []string `json:"billable_items,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
	PriceType V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigScopePriceType `json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingServiceActionCreditGrantPerTenantApplicabilityConfig struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigScope `json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingServiceActionCreditGrantPerTenantExpiryConfig struct {
	// The type of the expiry configuration. We currently support `end_of_service_period`.
	Type V2BillingServiceActionCreditGrantPerTenantExpiryConfigType `json:"type"`
}

// Details for the credit grant per tenant. Provided only if `type` is "credit_grant_per_tenant".
type V2BillingServiceActionCreditGrantPerTenant struct {
	// The amount of the credit grant.
	Amount *V2BillingServiceActionCreditGrantPerTenantAmount `json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingServiceActionCreditGrantPerTenantApplicabilityConfig `json:"applicability_config"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingServiceActionCreditGrantPerTenantExpiryConfig `json:"expiry_config"`
	// Customer-facing name for the credit grant.
	Name string `json:"name"`
}
type V2BillingServiceAction struct {
	APIResource
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Details for the credit grant. Provided only if `type` is "credit_grant".
	CreditGrant *V2BillingServiceActionCreditGrant `json:"credit_grant,omitempty"`
	// Details for the credit grant per tenant. Provided only if `type` is "credit_grant_per_tenant".
	CreditGrantPerTenant *V2BillingServiceActionCreditGrantPerTenant `json:"credit_grant_per_tenant,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// An internal key you can use to search for this service action.
	LookupKey string `json:"lookup_key,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The interval for assessing service.
	ServiceInterval V2BillingServiceActionServiceInterval `json:"service_interval"`
	// The length of the interval for assessing service.
	ServiceIntervalCount int64 `json:"service_interval_count"`
	// The type of the service action.
	Type V2BillingServiceActionType `json:"type"`
}
