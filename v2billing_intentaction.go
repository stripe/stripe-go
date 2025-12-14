//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Type of the Billing Intent Action.
type V2BillingIntentActionType string

// List of values that V2BillingIntentActionType can take
const (
	V2BillingIntentActionTypeApply      V2BillingIntentActionType = "apply"
	V2BillingIntentActionTypeDeactivate V2BillingIntentActionType = "deactivate"
	V2BillingIntentActionTypeModify     V2BillingIntentActionType = "modify"
	V2BillingIntentActionTypeRemove     V2BillingIntentActionType = "remove"
	V2BillingIntentActionTypeSubscribe  V2BillingIntentActionType = "subscribe"
)

// Type of the apply action details.
type V2BillingIntentActionApplyType string

// List of values that V2BillingIntentActionApplyType can take
const (
	V2BillingIntentActionApplyTypeInvoiceDiscountRule V2BillingIntentActionApplyType = "invoice_discount_rule"
)

// The entity that the discount rule applies to, for example, the Billing Cadence.
type V2BillingIntentActionApplyInvoiceDiscountRuleAppliesTo string

// List of values that V2BillingIntentActionApplyInvoiceDiscountRuleAppliesTo can take
const (
	V2BillingIntentActionApplyInvoiceDiscountRuleAppliesToCadence V2BillingIntentActionApplyInvoiceDiscountRuleAppliesTo = "cadence"
)

// Type of the discount rule.
type V2BillingIntentActionApplyInvoiceDiscountRuleType string

// List of values that V2BillingIntentActionApplyInvoiceDiscountRuleType can take
const (
	V2BillingIntentActionApplyInvoiceDiscountRuleTypePercentOff V2BillingIntentActionApplyInvoiceDiscountRuleType = "percent_off"
)

// The type of maximum applications configuration.
type V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsType string

// List of values that V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsType can take
const (
	V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsTypeIndefinite V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsType = "indefinite"
)

// Allows users to override the collect at behavior.
type V2BillingIntentActionDeactivateCollectAt string

// List of values that V2BillingIntentActionDeactivateCollectAt can take
const (
	V2BillingIntentActionDeactivateCollectAtNextBillingDate V2BillingIntentActionDeactivateCollectAt = "next_billing_date"
	V2BillingIntentActionDeactivateCollectAtOnEffectiveAt   V2BillingIntentActionDeactivateCollectAt = "on_effective_at"
)

// When the deactivate action will take effect.
type V2BillingIntentActionDeactivateEffectiveAtType string

// List of values that V2BillingIntentActionDeactivateEffectiveAtType can take
const (
	V2BillingIntentActionDeactivateEffectiveAtTypeCurrentBillingPeriodEnd   V2BillingIntentActionDeactivateEffectiveAtType = "current_billing_period_end"
	V2BillingIntentActionDeactivateEffectiveAtTypeCurrentBillingPeriodStart V2BillingIntentActionDeactivateEffectiveAtType = "current_billing_period_start"
	V2BillingIntentActionDeactivateEffectiveAtTypeOnReserve                 V2BillingIntentActionDeactivateEffectiveAtType = "on_reserve"
	V2BillingIntentActionDeactivateEffectiveAtTypeTimestamp                 V2BillingIntentActionDeactivateEffectiveAtType = "timestamp"
)

// Type of the action details.
type V2BillingIntentActionDeactivateType string

// List of values that V2BillingIntentActionDeactivateType can take
const (
	V2BillingIntentActionDeactivateTypePricingPlanSubscriptionDetails V2BillingIntentActionDeactivateType = "pricing_plan_subscription_details"
	V2BillingIntentActionDeactivateTypeV1SubscriptionDetails          V2BillingIntentActionDeactivateType = "v1_subscription_details"
)

// Type of the partial period behavior override.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType string

// List of values that V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType can take
const (
	V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorTypeLicenseFee V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType = "license_fee"
)

// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is deactivating.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehavior string

// List of values that V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehavior can take
const (
	V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehaviorNone     V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehavior = "none"
	V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehaviorProrated V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehavior = "prorated"
)

// Allows users to override the collect at behavior.
type V2BillingIntentActionModifyCollectAt string

// List of values that V2BillingIntentActionModifyCollectAt can take
const (
	V2BillingIntentActionModifyCollectAtNextBillingDate V2BillingIntentActionModifyCollectAt = "next_billing_date"
	V2BillingIntentActionModifyCollectAtOnEffectiveAt   V2BillingIntentActionModifyCollectAt = "on_effective_at"
)

// When the modify action will take effect.
type V2BillingIntentActionModifyEffectiveAtType string

// List of values that V2BillingIntentActionModifyEffectiveAtType can take
const (
	V2BillingIntentActionModifyEffectiveAtTypeCurrentBillingPeriodStart V2BillingIntentActionModifyEffectiveAtType = "current_billing_period_start"
	V2BillingIntentActionModifyEffectiveAtTypeOnReserve                 V2BillingIntentActionModifyEffectiveAtType = "on_reserve"
	V2BillingIntentActionModifyEffectiveAtTypeTimestamp                 V2BillingIntentActionModifyEffectiveAtType = "timestamp"
)

// Type of the action details.
type V2BillingIntentActionModifyType string

// List of values that V2BillingIntentActionModifyType can take
const (
	V2BillingIntentActionModifyTypePricingPlanSubscriptionDetails V2BillingIntentActionModifyType = "pricing_plan_subscription_details"
	V2BillingIntentActionModifyTypeV1SubscriptionDetails          V2BillingIntentActionModifyType = "v1_subscription_details"
)

// Type of the partial period behavior override.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType string

// List of values that V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType can take
const (
	V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorTypeLicenseFee V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType = "license_fee"
)

// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is upgrading.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehavior string

// List of values that V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehavior can take
const (
	V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehaviorNone     V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehavior = "none"
	V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehaviorProrated V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehavior = "prorated"
)

// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is downgrading.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehavior string

// List of values that V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehavior can take
const (
	V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehaviorNone     V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehavior = "none"
	V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehaviorProrated V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehavior = "prorated"
)

// Type of the remove action.
type V2BillingIntentActionRemoveType string

// List of values that V2BillingIntentActionRemoveType can take
const (
	V2BillingIntentActionRemoveTypeInvoiceDiscountRule V2BillingIntentActionRemoveType = "invoice_discount_rule"
)

// Allows users to override the collect at behavior.
type V2BillingIntentActionSubscribeCollectAt string

// List of values that V2BillingIntentActionSubscribeCollectAt can take
const (
	V2BillingIntentActionSubscribeCollectAtNextBillingDate V2BillingIntentActionSubscribeCollectAt = "next_billing_date"
	V2BillingIntentActionSubscribeCollectAtOnEffectiveAt   V2BillingIntentActionSubscribeCollectAt = "on_effective_at"
)

// When the subscribe action will take effect.
type V2BillingIntentActionSubscribeEffectiveAtType string

// List of values that V2BillingIntentActionSubscribeEffectiveAtType can take
const (
	V2BillingIntentActionSubscribeEffectiveAtTypeCurrentBillingPeriodStart V2BillingIntentActionSubscribeEffectiveAtType = "current_billing_period_start"
	V2BillingIntentActionSubscribeEffectiveAtTypeOnReserve                 V2BillingIntentActionSubscribeEffectiveAtType = "on_reserve"
	V2BillingIntentActionSubscribeEffectiveAtTypeTimestamp                 V2BillingIntentActionSubscribeEffectiveAtType = "timestamp"
)

// Type of the action details.
type V2BillingIntentActionSubscribeType string

// List of values that V2BillingIntentActionSubscribeType can take
const (
	V2BillingIntentActionSubscribeTypePricingPlanSubscriptionDetails V2BillingIntentActionSubscribeType = "pricing_plan_subscription_details"
	V2BillingIntentActionSubscribeTypeV1SubscriptionDetails          V2BillingIntentActionSubscribeType = "v1_subscription_details"
)

// Type of the partial period behavior override.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType string

// List of values that V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType can take
const (
	V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorTypeLicenseFee V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType = "license_fee"
)

// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is subscribing.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehavior string

// List of values that V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehavior can take
const (
	V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehaviorNone     V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehavior = "none"
	V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehaviorProrated V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehavior = "prorated"
)

// The maximum number of times this discount can be applied for this Billing Cadence.
type V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplications struct {
	// The type of maximum applications configuration.
	Type V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsType `json:"type"`
}

// Configuration for percentage off discount.
type V2BillingIntentActionApplyInvoiceDiscountRulePercentOff struct {
	// The maximum number of times this discount can be applied for this Billing Cadence.
	MaximumApplications *V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplications `json:"maximum_applications"`
	// Percent that will be taken off of the amount. For example, percent_off of 50.0 will make $100 amount $50 instead.
	PercentOff string `json:"percent_off"`
}

// Details for applying a discount rule to future invoices.
type V2BillingIntentActionApplyInvoiceDiscountRule struct {
	// The entity that the discount rule applies to, for example, the Billing Cadence.
	AppliesTo V2BillingIntentActionApplyInvoiceDiscountRuleAppliesTo `json:"applies_to"`
	// The ID of the created discount rule. This is only present once the Billing Intent is committed and the discount rule is created.
	InvoiceDiscountRule string `json:"invoice_discount_rule,omitempty"`
	// Configuration for percentage off discount.
	PercentOff *V2BillingIntentActionApplyInvoiceDiscountRulePercentOff `json:"percent_off,omitempty"`
	// Type of the discount rule.
	Type V2BillingIntentActionApplyInvoiceDiscountRuleType `json:"type"`
}

// Details for an apply action.
type V2BillingIntentActionApply struct {
	// Details for applying a discount rule to future invoices.
	InvoiceDiscountRule *V2BillingIntentActionApplyInvoiceDiscountRule `json:"invoice_discount_rule,omitempty"`
	// Type of the apply action details.
	Type V2BillingIntentActionApplyType `json:"type"`
}

// When the deactivate action will take effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentActionDeactivateEffectiveAt struct {
	// The timestamp at which the deactivate action will take effect. Only present if type is timestamp.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// When the deactivate action will take effect.
	Type V2BillingIntentActionDeactivateEffectiveAtType `json:"type"`
}

// Override for the license fee.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFee struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is deactivating.
	CreditProrationBehavior V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehavior `json:"credit_proration_behavior"`
}

// Override for the partial period behavior.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior struct {
	// Override for the license fee.
	LicenseFee *V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFee `json:"license_fee,omitempty"`
	// Type of the partial period behavior override.
	Type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType `json:"type"`
}

// Allows users to override the partial period behavior.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverrides struct {
	// Override for the partial period behavior.
	PartialPeriodBehaviors []*V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior `json:"partial_period_behaviors"`
}

// Details for deactivating a Pricing Plan Subscription.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetails struct {
	// Allows users to override the partial period behavior.
	Overrides *V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverrides `json:"overrides,omitempty"`
	// ID of the Pricing Plan Subscription to deactivate.
	PricingPlanSubscription string `json:"pricing_plan_subscription"`
}

// Details for a deactivate action.
type V2BillingIntentActionDeactivate struct {
	// Allows users to override the collect at behavior.
	CollectAt V2BillingIntentActionDeactivateCollectAt `json:"collect_at"`
	// When the deactivate action will take effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentActionDeactivateEffectiveAt `json:"effective_at"`
	// Details for deactivating a Pricing Plan Subscription.
	PricingPlanSubscriptionDetails *V2BillingIntentActionDeactivatePricingPlanSubscriptionDetails `json:"pricing_plan_subscription_details,omitempty"`
	// Type of the action details.
	Type V2BillingIntentActionDeactivateType `json:"type"`
}

// When the modify action will take effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentActionModifyEffectiveAt struct {
	// The timestamp at which the modify action will take effect. Only present if type is timestamp.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// When the modify action will take effect.
	Type V2BillingIntentActionModifyEffectiveAtType `json:"type"`
}

// New configurations for the components of the Pricing Plan.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsComponentConfiguration struct {
	// Lookup key for the pricing plan component.
	LookupKey string `json:"lookup_key,omitempty"`
	// ID of the pricing plan component.
	PricingPlanComponent string `json:"pricing_plan_component,omitempty"`
	// Quantity of the component to be used.
	Quantity int64 `json:"quantity,omitempty"`
}

// Override for the license fee.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFee struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is upgrading.
	CreditProrationBehavior V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeCreditProrationBehavior `json:"credit_proration_behavior"`
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is downgrading.
	DebitProrationBehavior V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehavior `json:"debit_proration_behavior"`
}

// Override for the partial period behavior.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior struct {
	// Override for the license fee.
	LicenseFee *V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFee `json:"license_fee,omitempty"`
	// Type of the partial period behavior override.
	Type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType `json:"type"`
}

// Allows users to override the partial period behavior.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverrides struct {
	// Override for the partial period behavior.
	PartialPeriodBehaviors []*V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior `json:"partial_period_behaviors"`
}

// Details for modifying a Pricing Plan Subscription.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetails struct {
	// New configurations for the components of the Pricing Plan.
	ComponentConfigurations []*V2BillingIntentActionModifyPricingPlanSubscriptionDetailsComponentConfiguration `json:"component_configurations"`
	// ID of the new Pricing Plan.
	NewPricingPlan string `json:"new_pricing_plan"`
	// Version of the Pricing Plan to use.
	NewPricingPlanVersion string `json:"new_pricing_plan_version"`
	// Allows users to override the partial period behavior.
	Overrides *V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverrides `json:"overrides,omitempty"`
	// ID of the Pricing Plan Subscription to modify.
	PricingPlanSubscription string `json:"pricing_plan_subscription"`
}

// Details for a modify action.
type V2BillingIntentActionModify struct {
	// Allows users to override the collect at behavior.
	CollectAt V2BillingIntentActionModifyCollectAt `json:"collect_at"`
	// When the modify action will take effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentActionModifyEffectiveAt `json:"effective_at"`
	// Details for modifying a Pricing Plan Subscription.
	PricingPlanSubscriptionDetails *V2BillingIntentActionModifyPricingPlanSubscriptionDetails `json:"pricing_plan_subscription_details,omitempty"`
	// Type of the action details.
	Type V2BillingIntentActionModifyType `json:"type"`
}

// Details for a remove action.
type V2BillingIntentActionRemove struct {
	// The ID of the discount rule to remove for future invoices.
	InvoiceDiscountRule string `json:"invoice_discount_rule,omitempty"`
	// Type of the remove action.
	Type V2BillingIntentActionRemoveType `json:"type"`
}

// When the subscribe action will take effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentActionSubscribeEffectiveAt struct {
	// The timestamp at which the subscribe action will take effect. Only present if type is timestamp.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// When the subscribe action will take effect.
	Type V2BillingIntentActionSubscribeEffectiveAtType `json:"type"`
}

// Configurations for the components of the Pricing Plan.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsComponentConfiguration struct {
	// Lookup key for the pricing plan component.
	LookupKey string `json:"lookup_key,omitempty"`
	// ID of the pricing plan component.
	PricingPlanComponent string `json:"pricing_plan_component,omitempty"`
	// Quantity of the component to be used.
	Quantity int64 `json:"quantity,omitempty"`
}

// Override for the license fee.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFee struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is subscribing.
	DebitProrationBehavior V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeDebitProrationBehavior `json:"debit_proration_behavior"`
}

// Override for the partial period behavior.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior struct {
	// Override for the license fee.
	LicenseFee *V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFee `json:"license_fee,omitempty"`
	// Type of the partial period behavior override.
	Type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorType `json:"type"`
}

// Allows users to override the partial period behavior.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverrides struct {
	// Override for the partial period behavior.
	PartialPeriodBehaviors []*V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior `json:"partial_period_behaviors"`
}

// Details for subscribing to a Pricing Plan.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetails struct {
	// Configurations for the components of the Pricing Plan.
	ComponentConfigurations []*V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsComponentConfiguration `json:"component_configurations"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Allows users to override the partial period behavior.
	Overrides *V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverrides `json:"overrides,omitempty"`
	// ID of the Pricing Plan to subscribe to.
	PricingPlan string `json:"pricing_plan"`
	// ID of the created Pricing Plan Subscription. This is only present once the Billing Intent is committed and the Pricing Plan Subscription is created.
	PricingPlanSubscription string `json:"pricing_plan_subscription,omitempty"`
	// Version of the Pricing Plan to use.
	PricingPlanVersion string `json:"pricing_plan_version"`
}

// A list of up to 20 subscription items, each with an attached price.
type V2BillingIntentActionSubscribeV1SubscriptionDetailsItem struct {
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The ID of the price object.
	Price string `json:"price"`
	// Quantity for this item. If not provided, will default to 1.
	Quantity int64 `json:"quantity,omitempty"`
}

// Details for subscribing to a V1 subscription.
type V2BillingIntentActionSubscribeV1SubscriptionDetails struct {
	// The subscription's description, meant to be displayable to the customer.
	// Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description string `json:"description,omitempty"`
	// A list of up to 20 subscription items, each with an attached price.
	Items []*V2BillingIntentActionSubscribeV1SubscriptionDetailsItem `json:"items"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
}

// Details for a subscribe action.
type V2BillingIntentActionSubscribe struct {
	// Allows users to override the collect at behavior.
	CollectAt V2BillingIntentActionSubscribeCollectAt `json:"collect_at"`
	// When the subscribe action will take effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentActionSubscribeEffectiveAt `json:"effective_at"`
	// Details for subscribing to a Pricing Plan.
	PricingPlanSubscriptionDetails *V2BillingIntentActionSubscribePricingPlanSubscriptionDetails `json:"pricing_plan_subscription_details,omitempty"`
	// Type of the action details.
	Type V2BillingIntentActionSubscribeType `json:"type"`
	// Details for subscribing to a V1 subscription.
	V1SubscriptionDetails *V2BillingIntentActionSubscribeV1SubscriptionDetails `json:"v1_subscription_details,omitempty"`
}
type V2BillingIntentAction struct {
	APIResource
	// Details for an apply action.
	Apply *V2BillingIntentActionApply `json:"apply,omitempty"`
	// Time at which the object was created.
	Created time.Time `json:"created"`
	// Details for a deactivate action.
	Deactivate *V2BillingIntentActionDeactivate `json:"deactivate,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Details for a modify action.
	Modify *V2BillingIntentActionModify `json:"modify,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Details for a remove action.
	Remove *V2BillingIntentActionRemove `json:"remove,omitempty"`
	// Details for a subscribe action.
	Subscribe *V2BillingIntentActionSubscribe `json:"subscribe,omitempty"`
	// Type of the Billing Intent Action.
	Type V2BillingIntentActionType `json:"type"`
}
