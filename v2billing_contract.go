//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of tax calculation.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxType string

// List of values that V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxType can take
const (
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxTypeAutomatic V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxType = "automatic"
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxTypeManual    V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxType = "manual"
)

// The interval unit.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval string

// List of values that V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval can take
const (
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueIntervalDay   V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval = "day"
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueIntervalMonth V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval = "month"
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueIntervalWeek  V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval = "week"
	V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueIntervalYear  V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval = "year"
)

// The collection method.
type V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethod string

// List of values that V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethod can take
const (
	V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethodChargeAutomatically V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethod = "charge_automatically"
	V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethodSendInvoice         V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethod = "send_invoice"
)

// The type of the credit grant amount.
type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantAmountType string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantAmountType can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantAmountTypeMonetary V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantAmountType = "monetary"
)

// The price type that credit grants can apply to.
type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantApplicabilityConfigScopePriceType string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantApplicabilityConfigScopePriceType can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantApplicabilityConfigScopePriceTypeMetered V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantApplicabilityConfigScopePriceType = "metered"
)

// The category of the credit grant.
type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantCategory string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantCategory can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantCategoryPaid        V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantCategory = "paid"
	V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantCategoryPromotional V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantCategory = "promotional"
)

// The type of the expiry configuration.
type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantExpiryConfigType string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantExpiryConfigType can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantExpiryConfigTypeEndOfServicePeriod V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantExpiryConfigType = "end_of_service_period"
)

// The interval for assessing service.
type V2BillingContractContractLineDetailOverrideServiceActionAddServiceInterval string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionAddServiceInterval can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionAddServiceIntervalDay   V2BillingContractContractLineDetailOverrideServiceActionAddServiceInterval = "day"
	V2BillingContractContractLineDetailOverrideServiceActionAddServiceIntervalMonth V2BillingContractContractLineDetailOverrideServiceActionAddServiceInterval = "month"
	V2BillingContractContractLineDetailOverrideServiceActionAddServiceIntervalWeek  V2BillingContractContractLineDetailOverrideServiceActionAddServiceInterval = "week"
	V2BillingContractContractLineDetailOverrideServiceActionAddServiceIntervalYear  V2BillingContractContractLineDetailOverrideServiceActionAddServiceInterval = "year"
)

// The type of the service action.
type V2BillingContractContractLineDetailOverrideServiceActionAddType string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionAddType can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionAddTypeCreditGrant V2BillingContractContractLineDetailOverrideServiceActionAddType = "credit_grant"
)

// The type of the credit grant amount.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantAmountType string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantAmountType can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantAmountTypeMonetary V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantAmountType = "monetary"
)

// The price type that credit grants can apply to.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantApplicabilityConfigScopePriceType string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantApplicabilityConfigScopePriceType can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantApplicabilityConfigScopePriceTypeMetered V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantApplicabilityConfigScopePriceType = "metered"
)

// The category of the credit grant.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantCategory string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantCategory can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantCategoryPaid        V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantCategory = "paid"
	V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantCategoryPromotional V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantCategory = "promotional"
)

// The type of the expiry configuration.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantExpiryConfigType string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantExpiryConfigType can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantExpiryConfigTypeEndOfServicePeriod V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantExpiryConfigType = "end_of_service_period"
)

// The interval for assessing service.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceInterval string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceInterval can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceIntervalDay   V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceInterval = "day"
	V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceIntervalMonth V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceInterval = "month"
	V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceIntervalWeek  V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceInterval = "week"
	V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceIntervalYear  V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceInterval = "year"
)

// The type of the service action.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceType string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionReplaceType can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionReplaceTypeCreditGrant V2BillingContractContractLineDetailOverrideServiceActionReplaceType = "credit_grant"
)

// The type of service action override.
type V2BillingContractContractLineDetailOverrideServiceActionType string

// List of values that V2BillingContractContractLineDetailOverrideServiceActionType can take
const (
	V2BillingContractContractLineDetailOverrideServiceActionTypeAdd     V2BillingContractContractLineDetailOverrideServiceActionType = "add"
	V2BillingContractContractLineDetailOverrideServiceActionTypeReplace V2BillingContractContractLineDetailOverrideServiceActionType = "replace"
)

// The type of the override.
type V2BillingContractContractLineDetailOverrideType string

// List of values that V2BillingContractContractLineDetailOverrideType can take
const (
	V2BillingContractContractLineDetailOverrideTypeServiceAction V2BillingContractContractLineDetailOverrideType = "service_action"
)

// The type of the license pricing.
type V2BillingContractLicenseQuantityLicensePricingType string

// List of values that V2BillingContractLicenseQuantityLicensePricingType can take
const (
	V2BillingContractLicenseQuantityLicensePricingTypeLicenseFee V2BillingContractLicenseQuantityLicensePricingType = "license_fee"
	V2BillingContractLicenseQuantityLicensePricingTypePrice      V2BillingContractLicenseQuantityLicensePricingType = "price"
)

// The type of the bill_at.
type V2BillingContractOneTimeFeeBillScheduleBillAtType string

// List of values that V2BillingContractOneTimeFeeBillScheduleBillAtType can take
const (
	V2BillingContractOneTimeFeeBillScheduleBillAtTypeContractStart V2BillingContractOneTimeFeeBillScheduleBillAtType = "contract_start"
	V2BillingContractOneTimeFeeBillScheduleBillAtTypeDatetime      V2BillingContractOneTimeFeeBillScheduleBillAtType = "datetime"
)

// The type of billable item that this fee references.
type V2BillingContractOneTimeFeeBillableItemType string

// List of values that V2BillingContractOneTimeFeeBillableItemType can take
const (
	V2BillingContractOneTimeFeeBillableItemTypeProduct V2BillingContractOneTimeFeeBillableItemType = "product"
)

// The type of pricing.
type V2BillingContractPricingLinePricingType string

// List of values that V2BillingContractPricingLinePricingType can take
const (
	V2BillingContractPricingLinePricingTypePrice V2BillingContractPricingLinePricingType = "price"
)

// Filter by billable item type.
type V2BillingContractPricingOverrideMultiplierCriterionBillableItemType string

// List of values that V2BillingContractPricingOverrideMultiplierCriterionBillableItemType can take
const (
	V2BillingContractPricingOverrideMultiplierCriterionBillableItemTypeLicensed V2BillingContractPricingOverrideMultiplierCriterionBillableItemType = "licensed"
	V2BillingContractPricingOverrideMultiplierCriterionBillableItemTypeMetered  V2BillingContractPricingOverrideMultiplierCriterionBillableItemType = "metered"
)

// Whether to include or exclude items matching these criteria.
type V2BillingContractPricingOverrideMultiplierCriterionType string

// List of values that V2BillingContractPricingOverrideMultiplierCriterionType can take
const (
	V2BillingContractPricingOverrideMultiplierCriterionTypeExclude V2BillingContractPricingOverrideMultiplierCriterionType = "exclude"
	V2BillingContractPricingOverrideMultiplierCriterionTypeInclude V2BillingContractPricingOverrideMultiplierCriterionType = "include"
)

// Defines whether the tiered price should be graduated or volume-based.
type V2BillingContractPricingOverrideOverwritePriceTieringMode string

// List of values that V2BillingContractPricingOverrideOverwritePriceTieringMode can take
const (
	V2BillingContractPricingOverrideOverwritePriceTieringModeGraduated V2BillingContractPricingOverrideOverwritePriceTieringMode = "graduated"
	V2BillingContractPricingOverrideOverwritePriceTieringModeVolume    V2BillingContractPricingOverrideOverwritePriceTieringMode = "volume"
)

// No upper bound to this tier.
type V2BillingContractPricingOverrideOverwritePriceTierUpToInf string

// List of values that V2BillingContractPricingOverrideOverwritePriceTierUpToInf can take
const (
	V2BillingContractPricingOverrideOverwritePriceTierUpToInfInf V2BillingContractPricingOverrideOverwritePriceTierUpToInf = "inf"
)

// The type of pricing override.
type V2BillingContractPricingOverrideType string

// List of values that V2BillingContractPricingOverrideType can take
const (
	V2BillingContractPricingOverrideTypeMultiplier     V2BillingContractPricingOverrideType = "multiplier"
	V2BillingContractPricingOverrideTypeOverwritePrice V2BillingContractPricingOverrideType = "overwrite_price"
)

// The current status of the contract.
type V2BillingContractStatus string

// List of values that V2BillingContractStatus can take
const (
	V2BillingContractStatusActive   V2BillingContractStatus = "active"
	V2BillingContractStatusCanceled V2BillingContractStatus = "canceled"
	V2BillingContractStatusDraft    V2BillingContractStatus = "draft"
	V2BillingContractStatusEnded    V2BillingContractStatus = "ended"
)

// Tax calculation settings.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTax struct {
	// The type of tax calculation.
	Type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTaxType `json:"type"`
}

// Calculation settings.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculation struct {
	// Tax calculation settings.
	Tax *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculationTax `json:"tax,omitempty"`
}

// The number of time units before the invoice is past due.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDue struct {
	// The interval unit.
	Interval V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDueInterval `json:"interval"`
	// The number of intervals.
	IntervalCount int64 `json:"interval_count"`
}

// Invoice settings.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoice struct {
	// The number of time units before the invoice is past due.
	TimeUntilDue *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoiceTimeUntilDue `json:"time_until_due,omitempty"`
}

// The bill settings details.
type V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetails struct {
	// Calculation settings.
	Calculation *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsCalculation `json:"calculation,omitempty"`
	// Invoice settings.
	Invoice *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetailsInvoice `json:"invoice,omitempty"`
}

// The billing profile details.
type V2BillingContractBillingSettingsContractBillingDetailsBillingProfileDetails struct {
	// The customer who pays for the contract invoice.
	Customer string `json:"customer"`
	// The default payment method for the contract.
	DefaultPaymentMethod string `json:"default_payment_method,omitempty"`
}

// The collection settings details.
type V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetails struct {
	// The collection method.
	CollectionMethod V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetailsCollectionMethod `json:"collection_method"`
	// The payment method configuration.
	PaymentMethodConfiguration string `json:"payment_method_configuration,omitempty"`
}

// Billing settings details for the contract.
type V2BillingContractBillingSettingsContractBillingDetails struct {
	// The billing profile details.
	BillingProfileDetails *V2BillingContractBillingSettingsContractBillingDetailsBillingProfileDetails `json:"billing_profile_details"`
	// The bill settings details.
	BillSettingsDetails *V2BillingContractBillingSettingsContractBillingDetailsBillSettingsDetails `json:"bill_settings_details,omitempty"`
	// The collection settings details.
	CollectionSettingsDetails *V2BillingContractBillingSettingsContractBillingDetailsCollectionSettingsDetails `json:"collection_settings_details"`
}

// The billing settings for the contract.
type V2BillingContractBillingSettings struct {
	// Billing settings details for the contract.
	ContractBillingDetails *V2BillingContractBillingSettingsContractBillingDetails `json:"contract_billing_details,omitempty"`
}

// The computed value details for the contract line.
type V2BillingContractContractLineDetailContractLineValueDetails struct {
	// Computed sum of all licensed fees. Represented as a decimal string in minor currency units.
	Total string `json:"total"`
}

// Timestamp to indicate when the contract line ends.
type V2BillingContractContractLineDetailEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// Timestamp to indicate when the override ends.
type V2BillingContractContractLineDetailOverrideEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// The amount of the credit grant.
type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantAmount struct {
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary Amount `json:"monetary,omitempty"`
	// The type of the credit grant amount.
	Type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantAmountType `json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantApplicabilityConfigScope struct {
	// The billable items to apply the credit grant to.
	BillableItems []string `json:"billable_items,omitempty"`
	// The price type that credit grants can apply to.
	PriceType V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantApplicabilityConfigScopePriceType `json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantApplicabilityConfig struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantApplicabilityConfigScope `json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantExpiryConfig struct {
	// The type of the expiry configuration.
	Type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantExpiryConfigType `json:"type"`
}

// Details for the credit grant. Required if `type` is `credit_grant`.
type V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrant struct {
	// The amount of the credit grant.
	Amount *V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantAmount `json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantApplicabilityConfig `json:"applicability_config"`
	// The category of the credit grant.
	Category V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantCategory `json:"category,omitempty"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrantExpiryConfig `json:"expiry_config"`
	// A descriptive name.
	Name string `json:"name"`
	// The desired priority for applying this credit grant. The highest priority is 0 and the lowest is 100.
	Priority int64 `json:"priority,omitempty"`
}

// Parameters for adding a new service action.
type V2BillingContractContractLineDetailOverrideServiceActionAdd struct {
	// Details for the credit grant. Required if `type` is `credit_grant`.
	CreditGrant *V2BillingContractContractLineDetailOverrideServiceActionAddCreditGrant `json:"credit_grant,omitempty"`
	// The interval for assessing service.
	ServiceInterval V2BillingContractContractLineDetailOverrideServiceActionAddServiceInterval `json:"service_interval"`
	// The length of the interval for assessing service.
	ServiceIntervalCount int64 `json:"service_interval_count"`
	// The type of the service action.
	Type V2BillingContractContractLineDetailOverrideServiceActionAddType `json:"type"`
}

// The amount of the credit grant.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantAmount struct {
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary Amount `json:"monetary,omitempty"`
	// The type of the credit grant amount.
	Type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantAmountType `json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantApplicabilityConfigScope struct {
	// The billable items to apply the credit grant to.
	BillableItems []string `json:"billable_items,omitempty"`
	// The price type that credit grants can apply to.
	PriceType V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantApplicabilityConfigScopePriceType `json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantApplicabilityConfig struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantApplicabilityConfigScope `json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantExpiryConfig struct {
	// The type of the expiry configuration.
	Type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantExpiryConfigType `json:"type"`
}

// Details for the credit grant. Required if `type` is `credit_grant`.
type V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrant struct {
	// The amount of the credit grant.
	Amount *V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantAmount `json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantApplicabilityConfig `json:"applicability_config"`
	// The category of the credit grant.
	Category V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantCategory `json:"category,omitempty"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrantExpiryConfig `json:"expiry_config"`
	// A descriptive name.
	Name string `json:"name"`
	// The desired priority for applying this credit grant. The highest priority is 0 and the lowest is 100.
	Priority int64 `json:"priority,omitempty"`
}

// Parameters for replacing an existing service action.
type V2BillingContractContractLineDetailOverrideServiceActionReplace struct {
	// Details for the credit grant. Required if `type` is `credit_grant`.
	CreditGrant *V2BillingContractContractLineDetailOverrideServiceActionReplaceCreditGrant `json:"credit_grant,omitempty"`
	// The ID of the service action to replace.
	ID string `json:"id,omitempty"`
	// The lookup key for the service action to replace.
	LookupKey string `json:"lookup_key,omitempty"`
	// The interval for assessing service.
	ServiceInterval V2BillingContractContractLineDetailOverrideServiceActionReplaceServiceInterval `json:"service_interval"`
	// The length of the interval for assessing service.
	ServiceIntervalCount int64 `json:"service_interval_count"`
	// The type of the service action.
	Type V2BillingContractContractLineDetailOverrideServiceActionReplaceType `json:"type"`
}

// Service action override details.
type V2BillingContractContractLineDetailOverrideServiceAction struct {
	// Parameters for adding a new service action.
	Add *V2BillingContractContractLineDetailOverrideServiceActionAdd `json:"add,omitempty"`
	// Parameters for replacing an existing service action.
	Replace *V2BillingContractContractLineDetailOverrideServiceActionReplace `json:"replace,omitempty"`
	// The type of service action override.
	Type V2BillingContractContractLineDetailOverrideServiceActionType `json:"type"`
}

// Timestamp to indicate when the override starts.
type V2BillingContractContractLineDetailOverrideStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// List of overrides applied to the contract line.
type V2BillingContractContractLineDetailOverride struct {
	// Timestamp to indicate when the override ends.
	EndsAt *V2BillingContractContractLineDetailOverrideEndsAt `json:"ends_at"`
	// Service action override details.
	ServiceAction *V2BillingContractContractLineDetailOverrideServiceAction `json:"service_action,omitempty"`
	// Timestamp to indicate when the override starts.
	StartsAt *V2BillingContractContractLineDetailOverrideStartsAt `json:"starts_at"`
	// The type of the override.
	Type V2BillingContractContractLineDetailOverrideType `json:"type"`
}

// The pricing configuration for the contract line.
type V2BillingContractContractLineDetailPricing struct{}

// Timestamp to indicate when the contract line starts.
type V2BillingContractContractLineDetailStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// The contract line details of the contract. Only populated when `contract_line_details` is passed in the `include` parameter.
type V2BillingContractContractLineDetail struct {
	// The ID of the contract line.
	ContractLine string `json:"contract_line"`
	// The computed value details for the contract line.
	ContractLineValueDetails *V2BillingContractContractLineDetailContractLineValueDetails `json:"contract_line_value_details"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// Timestamp to indicate when the contract line ends.
	EndsAt *V2BillingContractContractLineDetailEndsAt `json:"ends_at"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `json:"metadata,omitempty"`
	// List of overrides applied to the contract line.
	Overrides []*V2BillingContractContractLineDetailOverride `json:"overrides"`
	// The pricing configuration for the contract line.
	Pricing *V2BillingContractContractLineDetailPricing `json:"pricing"`
	// Timestamp to indicate when the contract line starts.
	StartsAt *V2BillingContractContractLineDetailStartsAt `json:"starts_at"`
}

// The computed total value of all contract lines.
type V2BillingContractContractValueDetails struct {
	// The total value represented as a decimal string in minor currency units.
	Total string `json:"total"`
}

// The license quantities of the contract. Only populated when `license_quantities` is passed in the `include` parameter.
type V2BillingContractLicenseQuantity struct {
	// The ID of the license pricing.
	LicensePricingID string `json:"license_pricing_id"`
	// The type of the license pricing.
	LicensePricingType V2BillingContractLicenseQuantityLicensePricingType `json:"license_pricing_type"`
	// The ID of the pricing line associated with this license quantity.
	PricingLine string `json:"pricing_line"`
	// The current quantity of the license.
	Quantity int64 `json:"quantity"`
}

// When this entry will be billed.
type V2BillingContractOneTimeFeeBillScheduleBillAt struct {
	// The datetime at which the entry will be billed. Set when `type` is `datetime`.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// The type of the bill_at.
	Type V2BillingContractOneTimeFeeBillScheduleBillAtType `json:"type"`
}

// The resolved bill schedule for the fee.
type V2BillingContractOneTimeFeeBillSchedule struct {
	// When this entry will be billed.
	BillAt *V2BillingContractOneTimeFeeBillScheduleBillAt `json:"bill_at"`
	// The amount to bill for this entry, in the smallest currency unit.
	Value int64 `json:"value,string"`
}

// Details for a product billable target. Set when `billable_item_type` is `product`.
type V2BillingContractOneTimeFeeProductDetails struct {
	// The ID of the v1 Product.
	Product string `json:"product"`
}

// The one-time fees of the contract. Only populated when `one_time_fees` is passed in the `include` parameter.
type V2BillingContractOneTimeFee struct {
	// The type of billable item that this fee references.
	BillableItemType V2BillingContractOneTimeFeeBillableItemType `json:"billable_item_type"`
	// The resolved bill schedule for the fee.
	BillSchedule []*V2BillingContractOneTimeFeeBillSchedule `json:"bill_schedule"`
	// Details for a product billable target. Set when `billable_item_type` is `product`.
	ProductDetails *V2BillingContractOneTimeFeeProductDetails `json:"product_details,omitempty"`
}

// Resolved timestamp when the pricing line ends.
type V2BillingContractPricingLineEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// V1 price details. Present when `type` is `price`.
type V2BillingContractPricingLinePricingPriceDetails struct {
	// The ID of the V1 price.
	Price string `json:"price"`
}

// The pricing configuration for the pricing line.
type V2BillingContractPricingLinePricing struct {
	// V1 price details. Present when `type` is `price`.
	PriceDetails *V2BillingContractPricingLinePricingPriceDetails `json:"price_details,omitempty"`
	// The type of pricing.
	Type V2BillingContractPricingLinePricingType `json:"type"`
}

// Resolved timestamp when the pricing line starts.
type V2BillingContractPricingLineStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// The pricing lines of the contract. Only populated when `pricing_lines` is passed in the `include` parameter.
type V2BillingContractPricingLine struct {
	// Resolved timestamp when the pricing line ends.
	EndsAt *V2BillingContractPricingLineEndsAt `json:"ends_at"`
	// The user-provided lookup key for the pricing line.
	LookupKey string `json:"lookup_key,omitempty"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The pricing configuration for the pricing line.
	Pricing *V2BillingContractPricingLinePricing `json:"pricing"`
	// The ID of the pricing line.
	PricingLine string `json:"pricing_line"`
	// Resolved timestamp when the pricing line starts.
	StartsAt *V2BillingContractPricingLineStartsAt `json:"starts_at"`
}

// Resolved timestamp when the pricing override ends.
type V2BillingContractPricingOverrideEndsAt struct {
	// The timestamp when the item ends.
	Timestamp time.Time `json:"timestamp"`
}

// All of these key-value conditions must match.
type V2BillingContractPricingOverrideMultiplierCriterionMetadataConditionAllOf struct {
	// The metadata key.
	Key string `json:"key"`
	// The metadata value.
	Value string `json:"value"`
}

// Filter by metadata conditions.
type V2BillingContractPricingOverrideMultiplierCriterionMetadataCondition struct {
	// All of these key-value conditions must match.
	AllOf []*V2BillingContractPricingOverrideMultiplierCriterionMetadataConditionAllOf `json:"all_of"`
}

// Criteria determining which rates the multiplier applies to.
type V2BillingContractPricingOverrideMultiplierCriterion struct {
	// Filter by billable item IDs.
	BillableItemIDs []string `json:"billable_item_ids"`
	// Filter by billable item lookup keys.
	BillableItemLookupKeys []string `json:"billable_item_lookup_keys"`
	// Filter by billable item type.
	BillableItemTypes []V2BillingContractPricingOverrideMultiplierCriterionBillableItemType `json:"billable_item_types"`
	// Filter by metadata conditions.
	MetadataConditions []*V2BillingContractPricingOverrideMultiplierCriterionMetadataCondition `json:"metadata_conditions"`
	// Filter by rate card IDs. Only applicable for `multiplier` overrides.
	RateCardIDs []string `json:"rate_card_ids"`
	// Whether to include or exclude items matching these criteria.
	Type V2BillingContractPricingOverrideMultiplierCriterionType `json:"type"`
}

// Details for a multiplier override.
type V2BillingContractPricingOverrideMultiplier struct {
	// Criteria determining which rates the multiplier applies to.
	Criteria []*V2BillingContractPricingOverrideMultiplierCriterion `json:"criteria"`
	// The multiplier factor, represented as a decimal string. e.g. "0.8" for a 20% reduction.
	Factor string `json:"factor"`
}

// Each element represents a pricing tier.
type V2BillingContractPricingOverrideOverwritePriceTier struct {
	// Price for the entire tier, represented as a decimal string in minor currency units.
	FlatAmount string `json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units.
	UnitAmount string `json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier.
	UpToDecimal float64 `json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier.
	UpToInf V2BillingContractPricingOverrideOverwritePriceTierUpToInf `json:"up_to_inf,omitempty"`
}

// Details for an overwrite_price override.
type V2BillingContractPricingOverrideOverwritePrice struct {
	// The ID of the V1 price to overwrite.
	Price string `json:"price"`
	// Defines whether the tiered price should be graduated or volume-based.
	TieringMode V2BillingContractPricingOverrideOverwritePriceTieringMode `json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingContractPricingOverrideOverwritePriceTier `json:"tiers"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units.
	UnitAmount string `json:"unit_amount,omitempty"`
}

// Resolved timestamp when the pricing override starts.
type V2BillingContractPricingOverrideStartsAt struct {
	// The timestamp when the item starts.
	Timestamp time.Time `json:"timestamp"`
}

// The pricing overrides of the contract. Only populated when `pricing_overrides` is passed in the `include` parameter.
type V2BillingContractPricingOverride struct {
	// Resolved timestamp when the pricing override ends.
	EndsAt *V2BillingContractPricingOverrideEndsAt `json:"ends_at"`
	// The user-provided lookup key for the pricing override.
	LookupKey string `json:"lookup_key,omitempty"`
	// Details for a multiplier override.
	Multiplier *V2BillingContractPricingOverrideMultiplier `json:"multiplier,omitempty"`
	// Details for an overwrite_price override.
	OverwritePrice *V2BillingContractPricingOverrideOverwritePrice `json:"overwrite_price,omitempty"`
	// The ID of the pricing override.
	PricingOverride string `json:"pricing_override"`
	// The priority of this override relative to others. Lower number = higher priority.
	Priority int64 `json:"priority"`
	// Resolved timestamp when the pricing override starts.
	StartsAt *V2BillingContractPricingOverrideStartsAt `json:"starts_at"`
	// The type of pricing override.
	Type V2BillingContractPricingOverrideType `json:"type"`
}

// Details of the active contract status.
type V2BillingContractStatusDetailsActive struct {
	// The timestamp when the contract was activated.
	ActivatedAt time.Time `json:"activated_at"`
}

// Details of the canceled contract status.
type V2BillingContractStatusDetailsCanceled struct {
	// The timestamp when the contract was canceled.
	CanceledAt time.Time `json:"canceled_at"`
}

// Information about the contract status transitions.
type V2BillingContractStatusDetails struct {
	// Details of the active contract status.
	Active *V2BillingContractStatusDetailsActive `json:"active,omitempty"`
	// Details of the canceled contract status.
	Canceled *V2BillingContractStatusDetailsCanceled `json:"canceled,omitempty"`
}

// Main Contract resource representing a comprehensive billing agreement
type V2BillingContract struct {
	APIResource
	// The billing settings for the contract.
	BillingSettings *V2BillingContractBillingSettings `json:"billing_settings,omitempty"`
	// The contract line details of the contract. Only populated when `contract_line_details` is passed in the `include` parameter.
	ContractLineDetails []*V2BillingContractContractLineDetail `json:"contract_line_details"`
	// A unique user-provided contract number e.g. C-2026-0001.
	ContractNumber string `json:"contract_number"`
	// The computed total value of all contract lines.
	ContractValueDetails *V2BillingContractContractValueDetails `json:"contract_value_details"`
	// Timestamp of when the object was created.
	Created time.Time `json:"created"`
	// The currency of the contract.
	Currency Currency `json:"currency"`
	// The ID of the customer associated with the contract.
	Customer string `json:"customer"`
	// The ID of the contract object.
	ID string `json:"id"`
	// The license quantities of the contract. Only populated when `license_quantities` is passed in the `include` parameter.
	LicenseQuantities []*V2BillingContractLicenseQuantity `json:"license_quantities"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of key-value pairs that you can attach to an object.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The one-time fees of the contract. Only populated when `one_time_fees` is passed in the `include` parameter.
	OneTimeFees []*V2BillingContractOneTimeFee `json:"one_time_fees,omitempty"`
	// The pricing lines of the contract. Only populated when `pricing_lines` is passed in the `include` parameter.
	PricingLines []*V2BillingContractPricingLine `json:"pricing_lines"`
	// The pricing overrides of the contract. Only populated when `pricing_overrides` is passed in the `include` parameter.
	PricingOverrides []*V2BillingContractPricingOverride `json:"pricing_overrides"`
	// The current status of the contract.
	Status V2BillingContractStatus `json:"status"`
	// Information about the contract status transitions.
	StatusDetails *V2BillingContractStatusDetails `json:"status_details"`
}
