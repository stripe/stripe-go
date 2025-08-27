//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
type V2BillingServiceActionCreditGrantAmountCustomPricingUnitParams struct {
	// The id of the custom pricing unit.
	ID *string `form:"id" json:"id"`
	// The value of the credit grant, decimal value represented as a string.
	Value *string `form:"value" json:"value"`
}

// The amount of the credit grant.
type V2BillingServiceActionCreditGrantAmountParams struct {
	// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
	CustomPricingUnit *V2BillingServiceActionCreditGrantAmountCustomPricingUnitParams `form:"custom_pricing_unit" json:"custom_pricing_unit,omitempty"`
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary *Amount `form:"monetary" json:"monetary,omitempty"`
	// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
	Type *string `form:"type" json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingServiceActionCreditGrantApplicabilityConfigScopeParams struct {
	// The billable items to apply the credit grant to.
	BillableItems []*string `form:"billable_items" json:"billable_items,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
	PriceType *string `form:"price_type" json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingServiceActionCreditGrantApplicabilityConfigParams struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingServiceActionCreditGrantApplicabilityConfigScopeParams `form:"scope" json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingServiceActionCreditGrantExpiryConfigParams struct {
	// The type of the expiry configuration. We currently support `end_of_service_period`.
	Type *string `form:"type" json:"type"`
}

// Details for the credit grant. Required if `type` is `credit_grant`.
type V2BillingServiceActionCreditGrantParams struct {
	// The amount of the credit grant.
	Amount *V2BillingServiceActionCreditGrantAmountParams `form:"amount" json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingServiceActionCreditGrantApplicabilityConfigParams `form:"applicability_config" json:"applicability_config"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingServiceActionCreditGrantExpiryConfigParams `form:"expiry_config" json:"expiry_config"`
	// A descriptive name shown in dashboard.
	Name *string `form:"name" json:"name"`
}

// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
type V2BillingServiceActionCreditGrantPerTenantAmountCustomPricingUnitParams struct {
	// The id of the custom pricing unit.
	ID *string `form:"id" json:"id"`
	// The value of the credit grant, decimal value represented as a string.
	Value *string `form:"value" json:"value"`
}

// The amount of the credit grant.
type V2BillingServiceActionCreditGrantPerTenantAmountParams struct {
	// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
	CustomPricingUnit *V2BillingServiceActionCreditGrantPerTenantAmountCustomPricingUnitParams `form:"custom_pricing_unit" json:"custom_pricing_unit,omitempty"`
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary *Amount `form:"monetary" json:"monetary,omitempty"`
	// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
	Type *string `form:"type" json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigScopeParams struct {
	// The billable items to apply the credit grant to.
	BillableItems []*string `form:"billable_items" json:"billable_items,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
	PriceType *string `form:"price_type" json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigParams struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigScopeParams `form:"scope" json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingServiceActionCreditGrantPerTenantExpiryConfigParams struct {
	// The type of the expiry configuration. We currently support `end_of_service_period`.
	Type *string `form:"type" json:"type"`
}

// Dimension-based meter segment condition.
type V2BillingServiceActionCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodMeterSegmentConditionDimensionParams struct {
	// The payload key for the dimension.
	PayloadKey *string `form:"payload_key" json:"payload_key"`
	// The value for the dimension.
	Value *string `form:"value" json:"value"`
}

// The meter segment conditions for the grant condition.
type V2BillingServiceActionCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodMeterSegmentConditionParams struct {
	// Dimension-based meter segment condition.
	Dimension *V2BillingServiceActionCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodMeterSegmentConditionDimensionParams `form:"dimension" json:"dimension,omitempty"`
	// The type of the meter segment condition. We currently support `dimension`.
	Type *string `form:"type" json:"type"`
}

// The grant condition for the meter event first per period.
type V2BillingServiceActionCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodParams struct {
	// The meter segment conditions for the grant condition.
	MeterSegmentConditions []*V2BillingServiceActionCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodMeterSegmentConditionParams `form:"meter_segment_conditions" json:"meter_segment_conditions"`
}

// The grant condition for the credit grant.
type V2BillingServiceActionCreditGrantPerTenantGrantConditionParams struct {
	// The grant condition for the meter event first per period.
	MeterEventFirstPerPeriod *V2BillingServiceActionCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodParams `form:"meter_event_first_per_period" json:"meter_event_first_per_period,omitempty"`
	// The type of the grant condition. We currently support `meter_event_first_per_period`.
	Type *string `form:"type" json:"type"`
}

// Details for the credit grant per tenant. Required if `type` is `credit_grant_per_tenant`.
type V2BillingServiceActionCreditGrantPerTenantParams struct {
	// The amount of the credit grant.
	Amount *V2BillingServiceActionCreditGrantPerTenantAmountParams `form:"amount" json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingServiceActionCreditGrantPerTenantApplicabilityConfigParams `form:"applicability_config" json:"applicability_config"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingServiceActionCreditGrantPerTenantExpiryConfigParams `form:"expiry_config" json:"expiry_config"`
	// The grant condition for the credit grant.
	GrantCondition *V2BillingServiceActionCreditGrantPerTenantGrantConditionParams `form:"grant_condition" json:"grant_condition"`
	// Customer-facing name for the credit grant.
	Name *string `form:"name" json:"name"`
}

// Create a Service Action object.
type V2BillingServiceActionParams struct {
	Params `form:"*"`
	// Details for the credit grant. Required if `type` is `credit_grant`.
	CreditGrant *V2BillingServiceActionCreditGrantParams `form:"credit_grant" json:"credit_grant,omitempty"`
	// Details for the credit grant per tenant. Required if `type` is `credit_grant_per_tenant`.
	CreditGrantPerTenant *V2BillingServiceActionCreditGrantPerTenantParams `form:"credit_grant_per_tenant" json:"credit_grant_per_tenant,omitempty"`
	// An internal key you can use to search for this service action. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// The interval for assessing service.
	ServiceInterval *string `form:"service_interval" json:"service_interval,omitempty"`
	// The length of the interval for assessing service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count,omitempty"`
	// The type of the service action.
	Type *string `form:"type" json:"type,omitempty"`
}

// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
type V2BillingServiceActionCreateCreditGrantAmountCustomPricingUnitParams struct {
	// The id of the custom pricing unit.
	ID *string `form:"id" json:"id"`
	// The value of the credit grant, decimal value represented as a string.
	Value *string `form:"value" json:"value"`
}

// The amount of the credit grant.
type V2BillingServiceActionCreateCreditGrantAmountParams struct {
	// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
	CustomPricingUnit *V2BillingServiceActionCreateCreditGrantAmountCustomPricingUnitParams `form:"custom_pricing_unit" json:"custom_pricing_unit,omitempty"`
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary *Amount `form:"monetary" json:"monetary,omitempty"`
	// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
	Type *string `form:"type" json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingServiceActionCreateCreditGrantApplicabilityConfigScopeParams struct {
	// The billable items to apply the credit grant to.
	BillableItems []*string `form:"billable_items" json:"billable_items,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
	PriceType *string `form:"price_type" json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingServiceActionCreateCreditGrantApplicabilityConfigParams struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingServiceActionCreateCreditGrantApplicabilityConfigScopeParams `form:"scope" json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingServiceActionCreateCreditGrantExpiryConfigParams struct {
	// The type of the expiry configuration. We currently support `end_of_service_period`.
	Type *string `form:"type" json:"type"`
}

// Details for the credit grant. Required if `type` is `credit_grant`.
type V2BillingServiceActionCreateCreditGrantParams struct {
	// The amount of the credit grant.
	Amount *V2BillingServiceActionCreateCreditGrantAmountParams `form:"amount" json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingServiceActionCreateCreditGrantApplicabilityConfigParams `form:"applicability_config" json:"applicability_config"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingServiceActionCreateCreditGrantExpiryConfigParams `form:"expiry_config" json:"expiry_config"`
	// A descriptive name shown in dashboard.
	Name *string `form:"name" json:"name"`
}

// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
type V2BillingServiceActionCreateCreditGrantPerTenantAmountCustomPricingUnitParams struct {
	// The id of the custom pricing unit.
	ID *string `form:"id" json:"id"`
	// The value of the credit grant, decimal value represented as a string.
	Value *string `form:"value" json:"value"`
}

// The amount of the credit grant.
type V2BillingServiceActionCreateCreditGrantPerTenantAmountParams struct {
	// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
	CustomPricingUnit *V2BillingServiceActionCreateCreditGrantPerTenantAmountCustomPricingUnitParams `form:"custom_pricing_unit" json:"custom_pricing_unit,omitempty"`
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary *Amount `form:"monetary" json:"monetary,omitempty"`
	// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
	Type *string `form:"type" json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingServiceActionCreateCreditGrantPerTenantApplicabilityConfigScopeParams struct {
	// The billable items to apply the credit grant to.
	BillableItems []*string `form:"billable_items" json:"billable_items,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
	PriceType *string `form:"price_type" json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingServiceActionCreateCreditGrantPerTenantApplicabilityConfigParams struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingServiceActionCreateCreditGrantPerTenantApplicabilityConfigScopeParams `form:"scope" json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingServiceActionCreateCreditGrantPerTenantExpiryConfigParams struct {
	// The type of the expiry configuration. We currently support `end_of_service_period`.
	Type *string `form:"type" json:"type"`
}

// Dimension-based meter segment condition.
type V2BillingServiceActionCreateCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodMeterSegmentConditionDimensionParams struct {
	// The payload key for the dimension.
	PayloadKey *string `form:"payload_key" json:"payload_key"`
	// The value for the dimension.
	Value *string `form:"value" json:"value"`
}

// The meter segment conditions for the grant condition.
type V2BillingServiceActionCreateCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodMeterSegmentConditionParams struct {
	// Dimension-based meter segment condition.
	Dimension *V2BillingServiceActionCreateCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodMeterSegmentConditionDimensionParams `form:"dimension" json:"dimension,omitempty"`
	// The type of the meter segment condition. We currently support `dimension`.
	Type *string `form:"type" json:"type"`
}

// The grant condition for the meter event first per period.
type V2BillingServiceActionCreateCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodParams struct {
	// The meter segment conditions for the grant condition.
	MeterSegmentConditions []*V2BillingServiceActionCreateCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodMeterSegmentConditionParams `form:"meter_segment_conditions" json:"meter_segment_conditions"`
}

// The grant condition for the credit grant.
type V2BillingServiceActionCreateCreditGrantPerTenantGrantConditionParams struct {
	// The grant condition for the meter event first per period.
	MeterEventFirstPerPeriod *V2BillingServiceActionCreateCreditGrantPerTenantGrantConditionMeterEventFirstPerPeriodParams `form:"meter_event_first_per_period" json:"meter_event_first_per_period,omitempty"`
	// The type of the grant condition. We currently support `meter_event_first_per_period`.
	Type *string `form:"type" json:"type"`
}

// Details for the credit grant per tenant. Required if `type` is `credit_grant_per_tenant`.
type V2BillingServiceActionCreateCreditGrantPerTenantParams struct {
	// The amount of the credit grant.
	Amount *V2BillingServiceActionCreateCreditGrantPerTenantAmountParams `form:"amount" json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingServiceActionCreateCreditGrantPerTenantApplicabilityConfigParams `form:"applicability_config" json:"applicability_config"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingServiceActionCreateCreditGrantPerTenantExpiryConfigParams `form:"expiry_config" json:"expiry_config"`
	// The grant condition for the credit grant.
	GrantCondition *V2BillingServiceActionCreateCreditGrantPerTenantGrantConditionParams `form:"grant_condition" json:"grant_condition"`
	// Customer-facing name for the credit grant.
	Name *string `form:"name" json:"name"`
}

// Create a Service Action object.
type V2BillingServiceActionCreateParams struct {
	Params `form:"*"`
	// Details for the credit grant. Required if `type` is `credit_grant`.
	CreditGrant *V2BillingServiceActionCreateCreditGrantParams `form:"credit_grant" json:"credit_grant,omitempty"`
	// Details for the credit grant per tenant. Required if `type` is `credit_grant_per_tenant`.
	CreditGrantPerTenant *V2BillingServiceActionCreateCreditGrantPerTenantParams `form:"credit_grant_per_tenant" json:"credit_grant_per_tenant,omitempty"`
	// An internal key you can use to search for this service action. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// The interval for assessing service.
	ServiceInterval *string `form:"service_interval" json:"service_interval"`
	// The length of the interval for assessing service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count"`
	// The type of the service action.
	Type *string `form:"type" json:"type"`
}

// Retrieve a Service Action object.
type V2BillingServiceActionRetrieveParams struct {
	Params `form:"*"`
}
