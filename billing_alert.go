//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Defines the type of the alert.
type BillingAlertAlertType string

// List of values that BillingAlertAlertType can take
const (
	BillingAlertAlertTypeCreditBalanceThreshold BillingAlertAlertType = "credit_balance_threshold"
	BillingAlertAlertTypeSpendThreshold         BillingAlertAlertType = "spend_threshold"
	BillingAlertAlertTypeUsageThreshold         BillingAlertAlertType = "usage_threshold"
)

type BillingAlertCreditBalanceThresholdFilterType string

// List of values that BillingAlertCreditBalanceThresholdFilterType can take
const (
	BillingAlertCreditBalanceThresholdFilterTypeCustomer BillingAlertCreditBalanceThresholdFilterType = "customer"
	BillingAlertCreditBalanceThresholdFilterTypeTenant   BillingAlertCreditBalanceThresholdFilterType = "tenant"
)

// The type of this balance. We currently only support `monetary` amounts.
type BillingAlertCreditBalanceThresholdLteBalanceType string

// List of values that BillingAlertCreditBalanceThresholdLteBalanceType can take
const (
	BillingAlertCreditBalanceThresholdLteBalanceTypeCustomPricingUnit BillingAlertCreditBalanceThresholdLteBalanceType = "custom_pricing_unit"
	BillingAlertCreditBalanceThresholdLteBalanceTypeMonetary          BillingAlertCreditBalanceThresholdLteBalanceType = "monetary"
)

// Defines the period over which spend is aggregated.
type BillingAlertSpendThresholdAggregationPeriod string

// List of values that BillingAlertSpendThresholdAggregationPeriod can take
const (
	BillingAlertSpendThresholdAggregationPeriodBilling BillingAlertSpendThresholdAggregationPeriod = "billing"
)

// Defines the granularity of spend aggregation.
type BillingAlertSpendThresholdGroupBy string

// List of values that BillingAlertSpendThresholdGroupBy can take
const (
	BillingAlertSpendThresholdGroupByPricingPlanSubscription BillingAlertSpendThresholdGroupBy = "pricing_plan_subscription"
)

// The type of the threshold amount.
type BillingAlertSpendThresholdGTEType string

// List of values that BillingAlertSpendThresholdGTEType can take
const (
	BillingAlertSpendThresholdGTETypeAmount            BillingAlertSpendThresholdGTEType = "amount"
	BillingAlertSpendThresholdGTETypeCustomPricingUnit BillingAlertSpendThresholdGTEType = "custom_pricing_unit"
)

// Status of the alert. This can be active, inactive or archived.
type BillingAlertStatus string

// List of values that BillingAlertStatus can take
const (
	BillingAlertStatusActive   BillingAlertStatus = "active"
	BillingAlertStatusArchived BillingAlertStatus = "archived"
	BillingAlertStatusInactive BillingAlertStatus = "inactive"
)

type BillingAlertUsageThresholdFilterType string

// List of values that BillingAlertUsageThresholdFilterType can take
const (
	BillingAlertUsageThresholdFilterTypeCustomer BillingAlertUsageThresholdFilterType = "customer"
)

// Defines how the alert will behave.
type BillingAlertUsageThresholdRecurrence string

// List of values that BillingAlertUsageThresholdRecurrence can take
const (
	BillingAlertUsageThresholdRecurrenceOneTime BillingAlertUsageThresholdRecurrence = "one_time"
)

// Lists billing active and inactive alerts
type BillingAlertListParams struct {
	ListParams `form:"*"`
	// Filter results to only include this type of alert.
	AlertType *string `form:"alert_type"`
	// Filter results to only include alerts for the given customer.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Filter results to only include alerts with the given meter.
	Meter *string `form:"meter"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
type BillingAlertCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopeBillableItemParams struct {
	// The billable item ID this credit grant should apply to.
	ID *string `form:"id"`
}

// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
type BillingAlertCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopePriceParams struct {
	// The price ID this credit grant should apply to.
	ID *string `form:"id"`
}

// Specify the scope of this applicability config.
type BillingAlertCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopeParams struct {
	// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
	BillableItems []*BillingAlertCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopeBillableItemParams `form:"billable_items"`
	// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
	Prices []*BillingAlertCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopePriceParams `form:"prices"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. Cannot be used in combination with `prices`.
	PriceType *string `form:"price_type"`
}

// The applicability configuration for this credit grants filter.
type BillingAlertCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigParams struct {
	// Specify the scope of this applicability config.
	Scope *BillingAlertCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopeParams `form:"scope"`
}

// The credit grants for which to configure the credit balance alert.
type BillingAlertCreditBalanceThresholdFilterCreditGrantsParams struct {
	// The applicability configuration for this credit grants filter.
	ApplicabilityConfig *BillingAlertCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigParams `form:"applicability_config"`
}

// The filters allows limiting the scope of this credit balance alert. You must specify a customer filter at this time.
type BillingAlertCreditBalanceThresholdFilterParams struct {
	// The credit grants for which to configure the credit balance alert.
	CreditGrants *BillingAlertCreditBalanceThresholdFilterCreditGrantsParams `form:"credit_grants"`
	// Limit the scope to this credit balance alert only to this customer.
	Customer *string `form:"customer"`
	// What type of filter is being applied to this credit balance alert.
	Type *string `form:"type"`
}

// The custom pricing unit amount.
type BillingAlertCreditBalanceThresholdLteCustomPricingUnitParams struct {
	// The ID of the custom pricing unit.
	ID *string `form:"id"`
	// A positive decimal string representing the amount of the custom pricing unit threshold.
	Value *float64 `form:"value,high_precision"`
}

// The monetary amount.
type BillingAlertCreditBalanceThresholdLteMonetaryParams struct {
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the `value` parameter.
	Currency *string `form:"currency"`
	// An integer representing the amount of the threshold.
	Value *int64 `form:"value"`
}

// Defines at which value the alert will fire.
type BillingAlertCreditBalanceThresholdLteParams struct {
	// Specify the type of this balance. We currently only support `monetary` billing credits.
	BalanceType *string `form:"balance_type"`
	// The custom pricing unit amount.
	CustomPricingUnit *BillingAlertCreditBalanceThresholdLteCustomPricingUnitParams `form:"custom_pricing_unit"`
	// The monetary amount.
	Monetary *BillingAlertCreditBalanceThresholdLteMonetaryParams `form:"monetary"`
}

// The configuration of the credit balance threshold.
type BillingAlertCreditBalanceThresholdParams struct {
	// The filters allows limiting the scope of this credit balance alert. You must specify a customer filter at this time.
	Filters []*BillingAlertCreditBalanceThresholdFilterParams `form:"filters"`
	// Defines at which value the alert will fire.
	Lte *BillingAlertCreditBalanceThresholdLteParams `form:"lte"`
}

// Filters to scope the spend calculation.
type BillingAlertSpendThresholdFiltersParams struct {
	// Filter by billable item IDs. Maximum of 20 billable items.
	BillableItems []*string `form:"billable_items"`
	// Filter by billing cadence ID.
	BillingCadence *string `form:"billing_cadence"`
	// Filter by pricing plan ID.
	PricingPlan *string `form:"pricing_plan"`
	// Filter by pricing plan subscription ID.
	PricingPlanSubscription *string `form:"pricing_plan_subscription"`
}

// The monetary amount. Required when type is `amount`. The threshold is the total_before_tax, the amount consumed after all credits and discounts are applied, but before tax is applied.
type BillingAlertSpendThresholdGTEAmountParams struct {
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the `value` parameter.
	Currency *string `form:"currency"`
	// An integer representing the amount of the threshold.
	Value *int64 `form:"value"`
}

// The custom pricing unit amount. Required when type is `custom_pricing_unit`.
type BillingAlertSpendThresholdGTECustomPricingUnitParams struct {
	// The ID of the custom pricing unit.
	ID *string `form:"id"`
	// A positive decimal string representing the amount of the custom pricing unit threshold.
	Value *float64 `form:"value,high_precision"`
}

// Defines at which value the alert will fire.
type BillingAlertSpendThresholdGTEParams struct {
	// The monetary amount. Required when type is `amount`. The threshold is the total_before_tax, the amount consumed after all credits and discounts are applied, but before tax is applied.
	Amount *BillingAlertSpendThresholdGTEAmountParams `form:"amount"`
	// The custom pricing unit amount. Required when type is `custom_pricing_unit`.
	CustomPricingUnit *BillingAlertSpendThresholdGTECustomPricingUnitParams `form:"custom_pricing_unit"`
	// The type of the threshold amount.
	Type *string `form:"type"`
}

// The configuration of the spend threshold. An event fires when the amount consumed exceeds the threshold, after all credits and discounts are applied but before tax is applied.
type BillingAlertSpendThresholdParams struct {
	// Defines the period over which spend is aggregated.
	AggregationPeriod *string `form:"aggregation_period"`
	// Filters to scope the spend calculation.
	Filters *BillingAlertSpendThresholdFiltersParams `form:"filters"`
	// Defines the granularity of spend aggregation. Defaults to `pricing_plan_subscription`.
	GroupBy *string `form:"group_by"`
	// Defines at which value the alert will fire.
	GTE *BillingAlertSpendThresholdGTEParams `form:"gte"`
}

// The filters allows limiting the scope of this usage alert. You can only specify up to one filter at this time.
type BillingAlertUsageThresholdFilterParams struct {
	// Limit the scope to this usage alert only to this customer.
	Customer *string `form:"customer"`
	// What type of filter is being applied to this usage alert.
	Type *string `form:"type"`
}

// The configuration of the usage threshold.
type BillingAlertUsageThresholdParams struct {
	// The filters allows limiting the scope of this usage alert. You can only specify up to one filter at this time.
	Filters []*BillingAlertUsageThresholdFilterParams `form:"filters"`
	// Defines the threshold value that triggers the alert.
	GTE *int64 `form:"gte"`
	// The [Billing Meter](https://docs.stripe.com/api/billing/meter) ID whose usage is monitored.
	Meter *string `form:"meter"`
	// Defines how the alert will behave.
	Recurrence *string `form:"recurrence"`
}

// Creates a billing alert
type BillingAlertParams struct {
	Params `form:"*"`
	// The type of alert to create.
	AlertType *string `form:"alert_type"`
	// The configuration of the credit balance threshold.
	CreditBalanceThreshold *BillingAlertCreditBalanceThresholdParams `form:"credit_balance_threshold"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The configuration of the spend threshold. An event fires when the amount consumed exceeds the threshold, after all credits and discounts are applied but before tax is applied.
	SpendThreshold *BillingAlertSpendThresholdParams `form:"spend_threshold"`
	// The title of the alert.
	Title *string `form:"title"`
	// The configuration of the usage threshold.
	UsageThreshold *BillingAlertUsageThresholdParams `form:"usage_threshold"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Reactivates this alert, allowing it to trigger again.
type BillingAlertActivateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertActivateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Archives this alert, removing it from the list view and APIs. This is non-reversible.
type BillingAlertArchiveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertArchiveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Deactivates this alert, preventing it from triggering.
type BillingAlertDeactivateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertDeactivateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
type BillingAlertCreateCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopeBillableItemParams struct {
	// The billable item ID this credit grant should apply to.
	ID *string `form:"id"`
}

// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
type BillingAlertCreateCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopePriceParams struct {
	// The price ID this credit grant should apply to.
	ID *string `form:"id"`
}

// Specify the scope of this applicability config.
type BillingAlertCreateCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopeParams struct {
	// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
	BillableItems []*BillingAlertCreateCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopeBillableItemParams `form:"billable_items"`
	// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
	Prices []*BillingAlertCreateCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopePriceParams `form:"prices"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. Cannot be used in combination with `prices`.
	PriceType *string `form:"price_type"`
}

// The applicability configuration for this credit grants filter.
type BillingAlertCreateCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigParams struct {
	// Specify the scope of this applicability config.
	Scope *BillingAlertCreateCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigScopeParams `form:"scope"`
}

// The credit grants for which to configure the credit balance alert.
type BillingAlertCreateCreditBalanceThresholdFilterCreditGrantsParams struct {
	// The applicability configuration for this credit grants filter.
	ApplicabilityConfig *BillingAlertCreateCreditBalanceThresholdFilterCreditGrantsApplicabilityConfigParams `form:"applicability_config"`
}

// The filters allows limiting the scope of this credit balance alert. You must specify a customer filter at this time.
type BillingAlertCreateCreditBalanceThresholdFilterParams struct {
	// The credit grants for which to configure the credit balance alert.
	CreditGrants *BillingAlertCreateCreditBalanceThresholdFilterCreditGrantsParams `form:"credit_grants"`
	// Limit the scope to this credit balance alert only to this customer.
	Customer *string `form:"customer"`
	// What type of filter is being applied to this credit balance alert.
	Type *string `form:"type"`
}

// The custom pricing unit amount.
type BillingAlertCreateCreditBalanceThresholdLteCustomPricingUnitParams struct {
	// The ID of the custom pricing unit.
	ID *string `form:"id"`
	// A positive decimal string representing the amount of the custom pricing unit threshold.
	Value *float64 `form:"value,high_precision"`
}

// The monetary amount.
type BillingAlertCreateCreditBalanceThresholdLteMonetaryParams struct {
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the `value` parameter.
	Currency *string `form:"currency"`
	// An integer representing the amount of the threshold.
	Value *int64 `form:"value"`
}

// Defines at which value the alert will fire.
type BillingAlertCreateCreditBalanceThresholdLteParams struct {
	// Specify the type of this balance. We currently only support `monetary` billing credits.
	BalanceType *string `form:"balance_type"`
	// The custom pricing unit amount.
	CustomPricingUnit *BillingAlertCreateCreditBalanceThresholdLteCustomPricingUnitParams `form:"custom_pricing_unit"`
	// The monetary amount.
	Monetary *BillingAlertCreateCreditBalanceThresholdLteMonetaryParams `form:"monetary"`
}

// The configuration of the credit balance threshold.
type BillingAlertCreateCreditBalanceThresholdParams struct {
	// The filters allows limiting the scope of this credit balance alert. You must specify a customer filter at this time.
	Filters []*BillingAlertCreateCreditBalanceThresholdFilterParams `form:"filters"`
	// Defines at which value the alert will fire.
	Lte *BillingAlertCreateCreditBalanceThresholdLteParams `form:"lte"`
}

// Filters to scope the spend calculation.
type BillingAlertCreateSpendThresholdFiltersParams struct {
	// Filter by billable item IDs. Maximum of 20 billable items.
	BillableItems []*string `form:"billable_items"`
	// Filter by billing cadence ID.
	BillingCadence *string `form:"billing_cadence"`
	// Filter by pricing plan ID.
	PricingPlan *string `form:"pricing_plan"`
	// Filter by pricing plan subscription ID.
	PricingPlanSubscription *string `form:"pricing_plan_subscription"`
}

// The monetary amount. Required when type is `amount`. The threshold is the total_before_tax, the amount consumed after all credits and discounts are applied, but before tax is applied.
type BillingAlertCreateSpendThresholdGTEAmountParams struct {
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the `value` parameter.
	Currency *string `form:"currency"`
	// An integer representing the amount of the threshold.
	Value *int64 `form:"value"`
}

// The custom pricing unit amount. Required when type is `custom_pricing_unit`.
type BillingAlertCreateSpendThresholdGTECustomPricingUnitParams struct {
	// The ID of the custom pricing unit.
	ID *string `form:"id"`
	// A positive decimal string representing the amount of the custom pricing unit threshold.
	Value *float64 `form:"value,high_precision"`
}

// Defines at which value the alert will fire.
type BillingAlertCreateSpendThresholdGTEParams struct {
	// The monetary amount. Required when type is `amount`. The threshold is the total_before_tax, the amount consumed after all credits and discounts are applied, but before tax is applied.
	Amount *BillingAlertCreateSpendThresholdGTEAmountParams `form:"amount"`
	// The custom pricing unit amount. Required when type is `custom_pricing_unit`.
	CustomPricingUnit *BillingAlertCreateSpendThresholdGTECustomPricingUnitParams `form:"custom_pricing_unit"`
	// The type of the threshold amount.
	Type *string `form:"type"`
}

// The configuration of the spend threshold. An event fires when the amount consumed exceeds the threshold, after all credits and discounts are applied but before tax is applied.
type BillingAlertCreateSpendThresholdParams struct {
	// Defines the period over which spend is aggregated.
	AggregationPeriod *string `form:"aggregation_period"`
	// Filters to scope the spend calculation.
	Filters *BillingAlertCreateSpendThresholdFiltersParams `form:"filters"`
	// Defines the granularity of spend aggregation. Defaults to `pricing_plan_subscription`.
	GroupBy *string `form:"group_by"`
	// Defines at which value the alert will fire.
	GTE *BillingAlertCreateSpendThresholdGTEParams `form:"gte"`
}

// The filters allows limiting the scope of this usage alert. You can only specify up to one filter at this time.
type BillingAlertCreateUsageThresholdFilterParams struct {
	// Limit the scope to this usage alert only to this customer.
	Customer *string `form:"customer"`
	// What type of filter is being applied to this usage alert.
	Type *string `form:"type"`
}

// The configuration of the usage threshold.
type BillingAlertCreateUsageThresholdParams struct {
	// The filters allows limiting the scope of this usage alert. You can only specify up to one filter at this time.
	Filters []*BillingAlertCreateUsageThresholdFilterParams `form:"filters"`
	// Defines the threshold value that triggers the alert.
	GTE *int64 `form:"gte"`
	// The [Billing Meter](https://docs.stripe.com/api/billing/meter) ID whose usage is monitored.
	Meter *string `form:"meter"`
	// Defines how the alert will behave.
	Recurrence *string `form:"recurrence"`
}

// Creates a billing alert
type BillingAlertCreateParams struct {
	Params `form:"*"`
	// The type of alert to create.
	AlertType *string `form:"alert_type"`
	// The configuration of the credit balance threshold.
	CreditBalanceThreshold *BillingAlertCreateCreditBalanceThresholdParams `form:"credit_balance_threshold"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The configuration of the spend threshold. An event fires when the amount consumed exceeds the threshold, after all credits and discounts are applied but before tax is applied.
	SpendThreshold *BillingAlertCreateSpendThresholdParams `form:"spend_threshold"`
	// The title of the alert.
	Title *string `form:"title"`
	// The configuration of the usage threshold.
	UsageThreshold *BillingAlertCreateUsageThresholdParams `form:"usage_threshold"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a billing alert given an ID
type BillingAlertRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *BillingAlertRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The filters allow limiting the scope of this credit balance alert. You must specify only a customer filter at this time.
type BillingAlertCreditBalanceThresholdFilter struct {
	// Limit the scope of the alert to this customer ID
	Customer *Customer                                    `json:"customer"`
	Type     BillingAlertCreditBalanceThresholdFilterType `json:"type"`
}

// The custom pricing unit object.
type BillingAlertCreditBalanceThresholdLteCustomPricingUnitCustomPricingUnitDetails struct {
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The name of the custom pricing unit.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A lookup key for the custom pricing unit.
	LookupKey string `json:"lookup_key"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The status of the custom pricing unit.
	Status string `json:"status"`
}

// The custom pricing unit amount.
type BillingAlertCreditBalanceThresholdLteCustomPricingUnit struct {
	// The custom pricing unit object.
	CustomPricingUnitDetails *BillingAlertCreditBalanceThresholdLteCustomPricingUnitCustomPricingUnitDetails `json:"custom_pricing_unit_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A positive decimal string representing the amount.
	Value float64 `json:"value,string"`
}

// The monetary amount.
type BillingAlertCreditBalanceThresholdLteMonetary struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount.
	Value int64 `json:"value"`
}
type BillingAlertCreditBalanceThresholdLte struct {
	// The type of this balance. We currently only support `monetary` amounts.
	BalanceType BillingAlertCreditBalanceThresholdLteBalanceType `json:"balance_type"`
	// The custom pricing unit amount.
	CustomPricingUnit *BillingAlertCreditBalanceThresholdLteCustomPricingUnit `json:"custom_pricing_unit"`
	// The monetary amount.
	Monetary *BillingAlertCreditBalanceThresholdLteMonetary `json:"monetary"`
}

// Encapsulates configuration of the alert to monitor billing credit balance.
type BillingAlertCreditBalanceThreshold struct {
	// The filters allow limiting the scope of this credit balance alert. You must specify only a customer filter at this time.
	Filters []*BillingAlertCreditBalanceThresholdFilter `json:"filters"`
	Lte     *BillingAlertCreditBalanceThresholdLte      `json:"lte"`
}

// Filters to scope the spend calculation.
type BillingAlertSpendThresholdFilters struct {
	// Filter by billable item IDs.
	BillableItems []string `json:"billable_items"`
	// Filter by billing cadence ID.
	BillingCadence string `json:"billing_cadence"`
	// Filter by pricing plan ID.
	PricingPlan string `json:"pricing_plan"`
	// Filter by pricing plan subscription ID.
	PricingPlanSubscription string `json:"pricing_plan_subscription"`
}

// The monetary amount. Present when type is `amount`.
type BillingAlertSpendThresholdGTEAmount struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount.
	Value int64 `json:"value"`
}

// The custom pricing unit object.
type BillingAlertSpendThresholdGTECustomPricingUnitCustomPricingUnitDetails struct {
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The name of the custom pricing unit.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A lookup key for the custom pricing unit.
	LookupKey string `json:"lookup_key"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The status of the custom pricing unit.
	Status string `json:"status"`
}

// The custom pricing unit amount. Present when type is `custom_pricing_unit`.
type BillingAlertSpendThresholdGTECustomPricingUnit struct {
	// The custom pricing unit object.
	CustomPricingUnitDetails *BillingAlertSpendThresholdGTECustomPricingUnitCustomPricingUnitDetails `json:"custom_pricing_unit_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A positive decimal string representing the amount.
	Value float64 `json:"value,string"`
}

// The threshold value configuration for a spend threshold alert.
type BillingAlertSpendThresholdGTE struct {
	// The monetary amount. Present when type is `amount`.
	Amount *BillingAlertSpendThresholdGTEAmount `json:"amount"`
	// The custom pricing unit amount. Present when type is `custom_pricing_unit`.
	CustomPricingUnit *BillingAlertSpendThresholdGTECustomPricingUnit `json:"custom_pricing_unit"`
	// The type of the threshold amount.
	Type BillingAlertSpendThresholdGTEType `json:"type"`
}

// Encapsulates the alert's configuration to monitor spend on pricing plan subscriptions.
type BillingAlertSpendThreshold struct {
	// Defines the period over which spend is aggregated.
	AggregationPeriod BillingAlertSpendThresholdAggregationPeriod `json:"aggregation_period"`
	// Filters to scope the spend calculation.
	Filters *BillingAlertSpendThresholdFilters `json:"filters"`
	// Defines the granularity of spend aggregation.
	GroupBy BillingAlertSpendThresholdGroupBy `json:"group_by"`
	// The threshold value configuration for a spend threshold alert.
	GTE *BillingAlertSpendThresholdGTE `json:"gte"`
}

// The filters allow limiting the scope of this usage alert. You can only specify up to one filter at this time.
type BillingAlertUsageThresholdFilter struct {
	// Limit the scope of the alert to this customer ID
	Customer *Customer                            `json:"customer"`
	Type     BillingAlertUsageThresholdFilterType `json:"type"`
}

// Encapsulates configuration of the alert to monitor usage on a specific [Billing Meter](https://docs.stripe.com/api/billing/meter).
type BillingAlertUsageThreshold struct {
	// The filters allow limiting the scope of this usage alert. You can only specify up to one filter at this time.
	Filters []*BillingAlertUsageThresholdFilter `json:"filters"`
	// The value at which this alert will trigger.
	GTE int64 `json:"gte"`
	// The [Billing Meter](https://docs.stripe.com/api/billing/meter) ID whose usage is monitored.
	Meter *BillingMeter `json:"meter"`
	// Defines how the alert will behave.
	Recurrence BillingAlertUsageThresholdRecurrence `json:"recurrence"`
}

// A billing alert is a resource that notifies you when a certain usage threshold on a meter is crossed. For example, you might create a billing alert to notify you when a certain user made 100 API requests.
type BillingAlert struct {
	APIResource
	// Defines the type of the alert.
	AlertType BillingAlertAlertType `json:"alert_type"`
	// Encapsulates configuration of the alert to monitor billing credit balance.
	CreditBalanceThreshold *BillingAlertCreditBalanceThreshold `json:"credit_balance_threshold"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Encapsulates the alert's configuration to monitor spend on pricing plan subscriptions.
	SpendThreshold *BillingAlertSpendThreshold `json:"spend_threshold"`
	// Status of the alert. This can be active, inactive or archived.
	Status BillingAlertStatus `json:"status"`
	// Title of the alert.
	Title string `json:"title"`
	// Encapsulates configuration of the alert to monitor usage on a specific [Billing Meter](https://docs.stripe.com/api/billing/meter).
	UsageThreshold *BillingAlertUsageThreshold `json:"usage_threshold"`
}

// BillingAlertList is a list of Alerts as retrieved from a list endpoint.
type BillingAlertList struct {
	APIResource
	ListMeta
	Data []*BillingAlert `json:"data"`
}
