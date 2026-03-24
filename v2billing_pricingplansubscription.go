//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The customer submitted reason for why they canceled, if the subscription was canceled explicitly by the user.
type V2BillingPricingPlanSubscriptionCancellationDetailsFeedback string

// List of values that V2BillingPricingPlanSubscriptionCancellationDetailsFeedback can take
const (
	V2BillingPricingPlanSubscriptionCancellationDetailsFeedbackCustomerService V2BillingPricingPlanSubscriptionCancellationDetailsFeedback = "customer_service"
	V2BillingPricingPlanSubscriptionCancellationDetailsFeedbackLowQuality      V2BillingPricingPlanSubscriptionCancellationDetailsFeedback = "low_quality"
	V2BillingPricingPlanSubscriptionCancellationDetailsFeedbackMissingFeatures V2BillingPricingPlanSubscriptionCancellationDetailsFeedback = "missing_features"
	V2BillingPricingPlanSubscriptionCancellationDetailsFeedbackOther           V2BillingPricingPlanSubscriptionCancellationDetailsFeedback = "other"
	V2BillingPricingPlanSubscriptionCancellationDetailsFeedbackSwitchedService V2BillingPricingPlanSubscriptionCancellationDetailsFeedback = "switched_service"
	V2BillingPricingPlanSubscriptionCancellationDetailsFeedbackTooComplex      V2BillingPricingPlanSubscriptionCancellationDetailsFeedback = "too_complex"
	V2BillingPricingPlanSubscriptionCancellationDetailsFeedbackTooExpensive    V2BillingPricingPlanSubscriptionCancellationDetailsFeedback = "too_expensive"
	V2BillingPricingPlanSubscriptionCancellationDetailsFeedbackUnused          V2BillingPricingPlanSubscriptionCancellationDetailsFeedback = "unused"
)

// System-generated reason for cancellation.
type V2BillingPricingPlanSubscriptionCancellationDetailsReason string

// List of values that V2BillingPricingPlanSubscriptionCancellationDetailsReason can take
const (
	V2BillingPricingPlanSubscriptionCancellationDetailsReasonCancellationRequested V2BillingPricingPlanSubscriptionCancellationDetailsReason = "cancellation_requested"
)

// Current collection status of this subscription.
type V2BillingPricingPlanSubscriptionCollectionStatus string

// List of values that V2BillingPricingPlanSubscriptionCollectionStatus can take
const (
	V2BillingPricingPlanSubscriptionCollectionStatusAwaitingCustomerAction V2BillingPricingPlanSubscriptionCollectionStatus = "awaiting_customer_action"
	V2BillingPricingPlanSubscriptionCollectionStatusCurrent                V2BillingPricingPlanSubscriptionCollectionStatus = "current"
	V2BillingPricingPlanSubscriptionCollectionStatusPastDue                V2BillingPricingPlanSubscriptionCollectionStatus = "past_due"
	V2BillingPricingPlanSubscriptionCollectionStatusPaused                 V2BillingPricingPlanSubscriptionCollectionStatus = "paused"
	V2BillingPricingPlanSubscriptionCollectionStatusUnpaid                 V2BillingPricingPlanSubscriptionCollectionStatus = "unpaid"
)

// Type of the Discount source.
type V2BillingPricingPlanSubscriptionDiscountDetailSourceType string

// List of values that V2BillingPricingPlanSubscriptionDiscountDetailSourceType can take
const (
	V2BillingPricingPlanSubscriptionDiscountDetailSourceTypeCoupon V2BillingPricingPlanSubscriptionDiscountDetailSourceType = "coupon"
)

// The interval for assessing service.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleInterval string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleInterval can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleIntervalDay   V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleInterval = "day"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleIntervalMonth V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleInterval = "month"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleIntervalWeek  V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleInterval = "week"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleIntervalYear  V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleInterval = "year"
)

// Defines whether the tiering price should be graduated or volume-based.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTieringMode string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTieringMode can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTieringModeGraduated V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTieringMode = "graduated"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTieringModeVolume    V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTieringMode = "volume"
)

// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTierUpToInf string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTierUpToInf can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTierUpToInfInf V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTierUpToInf = "inf"
)

// After division, round the result up or down.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTransformQuantityRound string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTransformQuantityRound can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTransformQuantityRoundDown V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTransformQuantityRound = "down"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTransformQuantityRoundUp   V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTransformQuantityRound = "up"
)

// The interval for assessing service.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleInterval string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleInterval can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleIntervalDay   V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleInterval = "day"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleIntervalMonth V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleInterval = "month"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleIntervalWeek  V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleInterval = "week"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleIntervalYear  V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleInterval = "year"
)

// Whether the rates are inclusive or exclusive of tax.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsTaxBehavior string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsTaxBehavior can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsTaxBehaviorExclusive V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsTaxBehavior = "exclusive"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsTaxBehaviorInclusive V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsTaxBehavior = "inclusive"
)

// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmountType string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmountType can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmountTypeCustomPricingUnit V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmountType = "custom_pricing_unit"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmountTypeMonetary          V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmountType = "monetary"
)

// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsApplicabilityConfigScopePriceType string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsApplicabilityConfigScopePriceType can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsApplicabilityConfigScopePriceTypeMetered V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsApplicabilityConfigScopePriceType = "metered"
)

// The type of the expiry configuration. We currently support `end_of_service_period`.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsExpiryConfigType string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsExpiryConfigType can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsExpiryConfigTypeEndOfServicePeriod V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsExpiryConfigType = "end_of_service_period"
)

// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmountType string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmountType can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmountTypeCustomPricingUnit V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmountType = "custom_pricing_unit"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmountTypeMonetary          V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmountType = "monetary"
)

// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsApplicabilityConfigScopePriceType string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsApplicabilityConfigScopePriceType can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsApplicabilityConfigScopePriceTypeMetered V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsApplicabilityConfigScopePriceType = "metered"
)

// The type of the expiry configuration. We currently support `end_of_service_period`.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsExpiryConfigType string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsExpiryConfigType can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsExpiryConfigTypeEndOfServicePeriod V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsExpiryConfigType = "end_of_service_period"
)

// The interval for assessing service.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleInterval string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleInterval can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleIntervalDay   V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleInterval = "day"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleIntervalMonth V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleInterval = "month"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleIntervalWeek  V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleInterval = "week"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleIntervalYear  V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleInterval = "year"
)

// The type of the recurring credit grant.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsType string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsType can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsTypeCreditGrant          V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsType = "credit_grant"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsTypeCreditGrantPerTenant V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsType = "credit_grant_per_tenant"
)

// The type of component details included.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailType string

// List of values that V2BillingPricingPlanSubscriptionPricingPlanComponentDetailType can take
const (
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailTypeLicenseFeeDetails           V2BillingPricingPlanSubscriptionPricingPlanComponentDetailType = "license_fee_details"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailTypeRateCardDetails             V2BillingPricingPlanSubscriptionPricingPlanComponentDetailType = "rate_card_details"
	V2BillingPricingPlanSubscriptionPricingPlanComponentDetailTypeRecurringCreditGrantDetails V2BillingPricingPlanSubscriptionPricingPlanComponentDetailType = "recurring_credit_grant_details"
)

// Current servicing status of this subscription.
type V2BillingPricingPlanSubscriptionServicingStatus string

// List of values that V2BillingPricingPlanSubscriptionServicingStatus can take
const (
	V2BillingPricingPlanSubscriptionServicingStatusActive   V2BillingPricingPlanSubscriptionServicingStatus = "active"
	V2BillingPricingPlanSubscriptionServicingStatusCanceled V2BillingPricingPlanSubscriptionServicingStatus = "canceled"
	V2BillingPricingPlanSubscriptionServicingStatusPaused   V2BillingPricingPlanSubscriptionServicingStatus = "paused"
	V2BillingPricingPlanSubscriptionServicingStatusPending  V2BillingPricingPlanSubscriptionServicingStatus = "pending"
)

// Details about why the subscription was canceled, if applicable. Includes system-generated reason.
type V2BillingPricingPlanSubscriptionCancellationDetails struct {
	// Additional comments about why the user canceled the subscription, if the subscription was canceled explicitly by the user.
	Comment string `json:"comment,omitempty"`
	// The customer submitted reason for why they canceled, if the subscription was canceled explicitly by the user.
	Feedback V2BillingPricingPlanSubscriptionCancellationDetailsFeedback `json:"feedback,omitempty"`
	// System-generated reason for cancellation.
	Reason V2BillingPricingPlanSubscriptionCancellationDetailsReason `json:"reason,omitempty"`
}

// Timestamps for collection status transitions.
type V2BillingPricingPlanSubscriptionCollectionStatusTransitions struct {
	// When the collection status transitioned to awaiting customer action.
	AwaitingCustomerActionAt string `json:"awaiting_customer_action_at,omitempty"`
	// When the collection status transitioned to current.
	CurrentAt string `json:"current_at,omitempty"`
	// When the collection status transitioned to past due.
	PastDueAt string `json:"past_due_at,omitempty"`
	// When the collection status transitioned to paused.
	PausedAt string `json:"paused_at,omitempty"`
	// When the collection status transitioned to unpaid.
	UnpaidAt string `json:"unpaid_at,omitempty"`
}

// The source of the Discount.
type V2BillingPricingPlanSubscriptionDiscountDetailSource struct {
	// The ID of the Coupon.
	Coupon string `json:"coupon,omitempty"`
	// Type of the Discount source.
	Type V2BillingPricingPlanSubscriptionDiscountDetailSourceType `json:"type"`
}

// Details about Discounts applied to this subscription.
type V2BillingPricingPlanSubscriptionDiscountDetail struct {
	// The ID of the Discount.
	Discount string `json:"discount"`
	// The time at which the Discount ends, if applicable.
	End time.Time `json:"end,omitempty"`
	// The ID of the PromotionCode, if applicable.
	PromotionCode string `json:"promotion_code,omitempty"`
	// The source of the Discount.
	Source *V2BillingPricingPlanSubscriptionDiscountDetailSource `json:"source"`
	// The time at which the Discount starts.
	Start time.Time `json:"start"`
}

// The service cycle configuration.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycle struct {
	// The interval for assessing service.
	Interval V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycleInterval `json:"interval"`
	// The length of the interval for assessing service. For example, set this to 3 and `interval` to `"month"` in
	// order to specify quarterly service.
	IntervalCount int64 `json:"interval_count"`
}

// Each element represents a pricing tier.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTier struct {
	// Price for the entire tier, represented as a decimal string in minor currency units with at most 12 decimal places.
	FlatAmount string `json:"flat_amount,omitempty"`
	// Per-unit price for units included in this tier, represented as a decimal string in minor currency units with at
	// most 12 decimal places.
	UnitAmount string `json:"unit_amount,omitempty"`
	// Up to and including this quantity will be contained in the tier. Only one of `up_to_decimal` and `up_to_inf` may
	// be set.
	UpToDecimal float64 `json:"up_to_decimal,string,omitempty"`
	// No upper bound to this tier. Only one of `up_to_decimal` and `up_to_inf` may be set.
	UpToInf V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTierUpToInf `json:"up_to_inf,omitempty"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTransformQuantity struct {
	// Divide usage by this number.
	DivideBy int64 `json:"divide_by,string"`
	// After division, round the result up or down.
	Round V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTransformQuantityRound `json:"round"`
}

// License fee details, present when type is license_fee_details.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetails struct {
	// Three-letter ISO currency code, in lowercase.
	Currency Currency `json:"currency"`
	// A customer-facing name for the license fee.
	DisplayName string `json:"display_name"`
	// The ID of the License Fee.
	LicenseFee string `json:"license_fee"`
	// The ID of the License Fee Version.
	LicenseFeeVersion string `json:"license_fee_version"`
	// Quantity of the license fee on the subscription.
	Quantity int64 `json:"quantity"`
	// The service cycle configuration.
	ServiceCycle *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsServiceCycle `json:"service_cycle"`
	// Defines whether the tiering price should be graduated or volume-based.
	TieringMode V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTieringMode `json:"tiering_mode,omitempty"`
	// Each element represents a pricing tier.
	Tiers []*V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTier `json:"tiers"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed.
	TransformQuantity *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetailsTransformQuantity `json:"transform_quantity,omitempty"`
	// The per-unit amount to be charged, represented as a decimal string in minor currency units with at most 12 decimal places.
	UnitAmount string `json:"unit_amount,omitempty"`
	// The unit label from the licensed item, used for display purposes (e.g. "seat", "environment").
	UnitLabel string `json:"unit_label,omitempty"`
}

// The service cycle configuration.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycle struct {
	// The interval for assessing service.
	Interval V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycleInterval `json:"interval"`
	// The length of the interval for assessing service. For example, set this to 3 and `interval` to `"month"` in
	// order to specify quarterly service.
	IntervalCount int64 `json:"interval_count"`
}

// Rate card details, present when type is rate_card_details.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetails struct {
	// Three-letter ISO currency code, in lowercase.
	Currency Currency `json:"currency"`
	// A customer-facing name for the rate card.
	DisplayName string `json:"display_name"`
	// The ID of the Rate Card.
	RateCard string `json:"rate_card"`
	// The ID of the Rate Card Version.
	RateCardVersion string `json:"rate_card_version"`
	// The service cycle configuration.
	ServiceCycle *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsServiceCycle `json:"service_cycle"`
	// Whether the rates are inclusive or exclusive of tax.
	TaxBehavior V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetailsTaxBehavior `json:"tax_behavior"`
}

// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmountCustomPricingUnit struct {
	// The Custom Pricing Unit object.
	CustomPricingUnitDetails *V2BillingCustomPricingUnit `json:"custom_pricing_unit_details,omitempty"`
	// The id of the custom pricing unit.
	ID string `json:"id"`
	// The value of the credit grant, decimal value represented as a string.
	Value float64 `json:"value,string"`
}

// The amount of the credit grant.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmount struct {
	// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
	CustomPricingUnit *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmountCustomPricingUnit `json:"custom_pricing_unit,omitempty"`
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary Amount `json:"monetary,omitempty"`
	// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
	Type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmountType `json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsApplicabilityConfigScope struct {
	// The billable items to apply the credit grant to.
	BillableItems []string `json:"billable_items,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
	PriceType V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsApplicabilityConfigScopePriceType `json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsApplicabilityConfig struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsApplicabilityConfigScope `json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsExpiryConfig struct {
	// The type of the expiry configuration. We currently support `end_of_service_period`.
	Type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsExpiryConfigType `json:"type"`
}

// Credit grant details, present when type is CREDIT_GRANT.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetails struct {
	// The amount of the credit grant.
	Amount *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsAmount `json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsApplicabilityConfig `json:"applicability_config"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetailsExpiryConfig `json:"expiry_config"`
}

// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmountCustomPricingUnit struct {
	// The Custom Pricing Unit object.
	CustomPricingUnitDetails *V2BillingCustomPricingUnit `json:"custom_pricing_unit_details,omitempty"`
	// The id of the custom pricing unit.
	ID string `json:"id"`
	// The value of the credit grant, decimal value represented as a string.
	Value float64 `json:"value,string"`
}

// The amount of the credit grant.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmount struct {
	// The custom pricing unit amount of the credit grant. Required if `type` is `custom_pricing_unit`.
	CustomPricingUnit *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmountCustomPricingUnit `json:"custom_pricing_unit,omitempty"`
	// The monetary amount of the credit grant. Required if `type` is `monetary`.
	Monetary Amount `json:"monetary,omitempty"`
	// The type of the credit grant amount. We currently support `monetary` and `custom_pricing_unit` billing credits.
	Type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmountType `json:"type"`
}

// The applicability scope of the credit grant.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsApplicabilityConfigScope struct {
	// The billable items to apply the credit grant to.
	BillableItems []string `json:"billable_items,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. This will apply to metered prices and rate cards. Cannot be used in combination with `billable_items`.
	PriceType V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsApplicabilityConfigScopePriceType `json:"price_type,omitempty"`
}

// Defines the scope where the credit grant is applicable.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsApplicabilityConfig struct {
	// The applicability scope of the credit grant.
	Scope *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsApplicabilityConfigScope `json:"scope"`
}

// The expiry configuration for the credit grant.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsExpiryConfig struct {
	// The type of the expiry configuration. We currently support `end_of_service_period`.
	Type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsExpiryConfigType `json:"type"`
}

// Credit grant per tenant details, present when type is CREDIT_GRANT_PER_TENANT.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetails struct {
	// The amount of the credit grant.
	Amount *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsAmount `json:"amount"`
	// Defines the scope where the credit grant is applicable.
	ApplicabilityConfig *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsApplicabilityConfig `json:"applicability_config"`
	// The expiry configuration for the credit grant.
	ExpiryConfig *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetailsExpiryConfig `json:"expiry_config"`
}

// The service cycle configuration.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycle struct {
	// The interval for assessing service.
	Interval V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycleInterval `json:"interval"`
	// The length of the interval for assessing service. For example, set this to 3 and `interval` to `"month"` in
	// order to specify quarterly service.
	IntervalCount int64 `json:"interval_count"`
}

// Recurring credit grant details, present when type is recurring_credit_grant_details.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetails struct {
	// Credit grant details, present when type is CREDIT_GRANT.
	CreditGrantDetails *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantDetails `json:"credit_grant_details,omitempty"`
	// Credit grant per tenant details, present when type is CREDIT_GRANT_PER_TENANT.
	CreditGrantPerTenantDetails *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsCreditGrantPerTenantDetails `json:"credit_grant_per_tenant_details,omitempty"`
	// The ID of the Recurring Credit Grant.
	RecurringCreditGrant string `json:"recurring_credit_grant"`
	// The service cycle configuration.
	ServiceCycle *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsServiceCycle `json:"service_cycle"`
	// The type of the recurring credit grant.
	Type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetailsType `json:"type"`
}

// Pricing plan component details for the subscription, populated when pricing_plan_component_details is included.
type V2BillingPricingPlanSubscriptionPricingPlanComponentDetail struct {
	// License fee details, present when type is license_fee_details.
	LicenseFeeDetails *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailLicenseFeeDetails `json:"license_fee_details,omitempty"`
	// The ID of the Pricing Plan Component.
	PricingPlanComponent string `json:"pricing_plan_component"`
	// Rate card details, present when type is rate_card_details.
	RateCardDetails *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRateCardDetails `json:"rate_card_details,omitempty"`
	// Recurring credit grant details, present when type is recurring_credit_grant_details.
	RecurringCreditGrantDetails *V2BillingPricingPlanSubscriptionPricingPlanComponentDetailRecurringCreditGrantDetails `json:"recurring_credit_grant_details,omitempty"`
	// The type of component details included.
	Type V2BillingPricingPlanSubscriptionPricingPlanComponentDetailType `json:"type"`
}

// Timestamps for servicing status transitions.
type V2BillingPricingPlanSubscriptionServicingStatusTransitions struct {
	// When the servicing status transitioned to activated.
	ActivatedAt string `json:"activated_at,omitempty"`
	// When the servicing status transitioned to canceled.
	CanceledAt string `json:"canceled_at,omitempty"`
	// When the servicing status transitioned to paused.
	PausedAt string `json:"paused_at,omitempty"`
	// When the servicing is scheduled to transition to activate.
	WillActivateAt string `json:"will_activate_at,omitempty"`
	// When the servicing is scheduled to cancel.
	WillCancelAt string `json:"will_cancel_at,omitempty"`
}

// A Pricing Plan Subscription represents a customer's active subscription to a Pricing Plan. It tracks both the servicing
// status (whether the customer is receiving service) and collection status (whether payments are current). Subscriptions
// are created through Billing Intents and bill according to the associated Billing Cadence.
type V2BillingPricingPlanSubscription struct {
	APIResource
	// The ID of the Billing Cadence this subscription is billed on.
	BillingCadence string `json:"billing_cadence"`
	// Details about why the subscription was canceled, if applicable. Includes system-generated reason.
	CancellationDetails *V2BillingPricingPlanSubscriptionCancellationDetails `json:"cancellation_details,omitempty"`
	// Current collection status of this subscription.
	CollectionStatus V2BillingPricingPlanSubscriptionCollectionStatus `json:"collection_status"`
	// Timestamps for collection status transitions.
	CollectionStatusTransitions *V2BillingPricingPlanSubscriptionCollectionStatusTransitions `json:"collection_status_transitions"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Details about Discounts applied to this subscription.
	DiscountDetails []*V2BillingPricingPlanSubscriptionDiscountDetail `json:"discount_details,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the Pricing Plan for this subscription.
	PricingPlan string `json:"pricing_plan"`
	// Pricing plan component details for the subscription, populated when pricing_plan_component_details is included.
	PricingPlanComponentDetails []*V2BillingPricingPlanSubscriptionPricingPlanComponentDetail `json:"pricing_plan_component_details,omitempty"`
	// The ID of the Pricing Plan Version for this subscription.
	PricingPlanVersion string `json:"pricing_plan_version"`
	// Current servicing status of this subscription.
	ServicingStatus V2BillingPricingPlanSubscriptionServicingStatus `json:"servicing_status"`
	// Timestamps for servicing status transitions.
	ServicingStatusTransitions *V2BillingPricingPlanSubscriptionServicingStatusTransitions `json:"servicing_status_transitions"`
	// The ID of the Test Clock of the associated Billing Cadence, if any.
	TestClock string `json:"test_clock,omitempty"`
}
