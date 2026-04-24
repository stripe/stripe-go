//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// List Billing Intents.
type V2BillingIntentListParams struct {
	Params `form:"*"`
	// Optionally set the maximum number of results per page. Defaults to 10.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Details for applying a discount.
type V2BillingIntentActionApplyDiscountParams struct {
	// The ID of the Coupon to apply.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// The ID of the PromotionCode to apply.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
	// Type of the discount.
	Type *string `form:"type" json:"type"`
}

// When the apply action takes effect. If not specified, defaults to on_reserve.
type V2BillingIntentActionApplyEffectiveAtParams struct {
	// The timestamp at which the apply action takes effect. Only present if type is timestamp. Only allowed for discount actions.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the apply action takes effect.
	Type *string `form:"type" json:"type"`
}

// The maximum number of times this discount can be applied for this cadence.
type V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams struct {
	// The type of maximum applications configuration.
	Type *string `form:"type" json:"type"`
}

// Configuration for percentage off discount.
type V2BillingIntentActionApplyInvoiceDiscountRulePercentOffParams struct {
	// The maximum number of times this discount can be applied for this cadence.
	MaximumApplications *V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams `form:"maximum_applications" json:"maximum_applications"`
	// Percent that is taken off the amount. For example, a percent_off of 50.0 reduces a 100 USD amount to 50 USD.
	PercentOff *float64 `form:"percent_off,high_precision" json:"percent_off,string"`
}

// Details for applying a discount rule to future invoices.
type V2BillingIntentActionApplyInvoiceDiscountRuleParams struct {
	// The entity that the discount rule applies to, for example, the cadence.
	AppliesTo *string `form:"applies_to" json:"applies_to"`
	// Configuration for percentage off discount.
	PercentOff *V2BillingIntentActionApplyInvoiceDiscountRulePercentOffParams `form:"percent_off" json:"percent_off,omitempty"`
	// Type of the discount rule.
	Type *string `form:"type" json:"type"`
}

// The custom pricing unit amount.
type V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendAmountCustomPricingUnitParams struct {
	// The id of the custom pricing unit.
	ID *string `form:"id" json:"id,omitempty"`
	// The value of the custom pricing unit.
	Value *string `form:"value" json:"value"`
}

// The maximum amount allowed for the billing period.
type V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendAmountParams struct {
	// The custom pricing unit amount.
	CustomPricingUnit *V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendAmountCustomPricingUnitParams `form:"custom_pricing_unit" json:"custom_pricing_unit,omitempty"`
	// The type of the amount.
	Type *string `form:"type" json:"type"`
}

// The configuration for the overage rate for the custom pricing unit.
type V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendCustomPricingUnitOverageRateParams struct {
	// ID of the custom pricing unit overage rate.
	ID *string `form:"id" json:"id"`
}

// Details for max billing period spend modifier. Only present if type is max_billing_period_spend.
type V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendParams struct {
	// The maximum amount allowed for the billing period.
	Amount *V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendAmountParams `form:"amount" json:"amount"`
	// The configuration for the overage rate for the custom pricing unit.
	CustomPricingUnitOverageRate *V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendCustomPricingUnitOverageRateParams `form:"custom_pricing_unit_overage_rate" json:"custom_pricing_unit_overage_rate"`
}

// Details for applying a spend modifier rule. Only present if type is spend_modifier_rule.
type V2BillingIntentActionApplySpendModifierRuleParams struct {
	// What the spend modifier applies to.
	AppliesTo *string `form:"applies_to" json:"applies_to"`
	// Details for max billing period spend modifier. Only present if type is max_billing_period_spend.
	MaxBillingPeriodSpend *V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendParams `form:"max_billing_period_spend" json:"max_billing_period_spend,omitempty"`
	// Type of the spend modifier.
	Type *string `form:"type" json:"type"`
}

// Details for an apply action.
type V2BillingIntentActionApplyParams struct {
	// Details for applying a discount.
	Discount *V2BillingIntentActionApplyDiscountParams `form:"discount" json:"discount,omitempty"`
	// When the apply action takes effect. If not specified, defaults to on_reserve.
	EffectiveAt *V2BillingIntentActionApplyEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for applying a discount rule to future invoices.
	InvoiceDiscountRule *V2BillingIntentActionApplyInvoiceDiscountRuleParams `form:"invoice_discount_rule" json:"invoice_discount_rule,omitempty"`
	// Details for applying a spend modifier rule. Only present if type is spend_modifier_rule.
	SpendModifierRule *V2BillingIntentActionApplySpendModifierRuleParams `form:"spend_modifier_rule" json:"spend_modifier_rule,omitempty"`
	// Type of the apply action details.
	Type *string `form:"type" json:"type"`
}

// Details about why the cancellation is being requested.
type V2BillingIntentActionDeactivateCancellationDetailsParams struct {
	// Additional comments about why the user canceled the subscription, if the subscription was canceled explicitly by the user.
	Comment *string `form:"comment" json:"comment,omitempty"`
	// The customer submitted reason for why they canceled, if the subscription was canceled explicitly by the user.
	Feedback *string `form:"feedback" json:"feedback,omitempty"`
}

// When the deactivate action takes effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentActionDeactivateEffectiveAtParams struct {
	// The timestamp at which the deactivate action takes effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the deactivate action takes effect.
	Type *string `form:"type" json:"type"`
}

// Overrides the behavior for license fee components when the action takes effect during the service period.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is deactivating. If not specified, defaults to none.
	CreditProrationBehavior *string `form:"credit_proration_behavior" json:"credit_proration_behavior"`
}

// Configurations for behaviors when the action takes effect during the service period.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Overrides the behavior for license fee components when the action takes effect during the service period.
	LicenseFee *V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// The type of behavior to override.
	Type *string `form:"type" json:"type"`
}

// Configurations for overriding behaviors related to the subscription.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesParams struct {
	// Configurations for behaviors when the action takes effect during the service period.
	PartialPeriodBehaviors []*V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams `form:"partial_period_behaviors" json:"partial_period_behaviors"`
}

// Details for deactivating a pricing plan subscription.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsParams struct {
	// Configurations for overriding behaviors related to the subscription.
	Overrides *V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesParams `form:"overrides" json:"overrides,omitempty"`
	// ID of the pricing plan subscription to deactivate.
	PricingPlanSubscription *string `form:"pricing_plan_subscription" json:"pricing_plan_subscription"`
}

// Details for a deactivate action.
type V2BillingIntentActionDeactivateParams struct {
	// Details about why the cancellation is being requested.
	CancellationDetails *V2BillingIntentActionDeactivateCancellationDetailsParams `form:"cancellation_details" json:"cancellation_details,omitempty"`
	// When the invoice is collected. If not specified, the default behavior is on_effective_at.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the deactivate action takes effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentActionDeactivateEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for deactivating a pricing plan subscription.
	PricingPlanSubscriptionDetails *V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsParams `form:"pricing_plan_subscription_details" json:"pricing_plan_subscription_details"`
	// Type of the action details.
	Type *string `form:"type" json:"type"`
}

// When the modify action takes effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentActionModifyEffectiveAtParams struct {
	// The timestamp at which the modify action takes effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the modify action takes effect.
	Type *string `form:"type" json:"type"`
}

// New configurations for the components of the pricing plan.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsComponentConfigurationParams struct {
	// Lookup key for the pricing plan component.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// ID of the pricing plan component.
	PricingPlanComponent *string `form:"pricing_plan_component" json:"pricing_plan_component,omitempty"`
	// Quantity of the component to be used.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
}

// Overrides the behavior for license fee components when the action takes effect during the service period.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user modifies the subscription. If not specified, defaults to prorated.
	CreditProrationBehavior *string `form:"credit_proration_behavior" json:"credit_proration_behavior"`
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user modifies the subscription. If not specified, defaults to prorated.
	DebitProrationBehavior *string `form:"debit_proration_behavior" json:"debit_proration_behavior"`
}

// Overrides the behavior for recurring credit grant components when the action takes effect during the service period.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorRecurringCreditGrantParams struct {
	// Controls credit grant creation behavior during partial periods. If not specified, defaults to full_credits.
	CreateBehavior *string `form:"create_behavior" json:"create_behavior"`
}

// Configurations for behaviors when the action takes effect during the service period.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Overrides the behavior for license fee components when the action takes effect during the service period.
	LicenseFee *V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// Overrides the behavior for recurring credit grant components when the action takes effect during the service period.
	RecurringCreditGrant *V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorRecurringCreditGrantParams `form:"recurring_credit_grant" json:"recurring_credit_grant,omitempty"`
	// The type of behavior to override.
	Type *string `form:"type" json:"type"`
}

// Configurations for overriding behaviors related to the subscription.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesParams struct {
	// Configurations for behaviors when the action takes effect during the service period.
	PartialPeriodBehaviors []*V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams `form:"partial_period_behaviors" json:"partial_period_behaviors"`
}

// Details for modifying a pricing plan subscription.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsParams struct {
	// New configurations for the components of the pricing plan.
	ComponentConfigurations []*V2BillingIntentActionModifyPricingPlanSubscriptionDetailsComponentConfigurationParams `form:"component_configurations" json:"component_configurations,omitempty"`
	// The ID of the new Pricing Plan, if changing plans.
	NewPricingPlan *string `form:"new_pricing_plan" json:"new_pricing_plan,omitempty"`
	// The ID of the new Pricing Plan Version to use.
	NewPricingPlanVersion *string `form:"new_pricing_plan_version" json:"new_pricing_plan_version,omitempty"`
	// Configurations for overriding behaviors related to the subscription.
	Overrides *V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesParams `form:"overrides" json:"overrides,omitempty"`
	// The ID of the Pricing Plan Subscription to modify.
	PricingPlanSubscription *string `form:"pricing_plan_subscription" json:"pricing_plan_subscription"`
}

// Details for a modify action.
type V2BillingIntentActionModifyParams struct {
	// When the invoice is collected. If not specified, the default behavior is next_billing_date.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the modify action takes effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentActionModifyEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for modifying a pricing plan subscription.
	PricingPlanSubscriptionDetails *V2BillingIntentActionModifyPricingPlanSubscriptionDetailsParams `form:"pricing_plan_subscription_details" json:"pricing_plan_subscription_details"`
	// Type of the action details.
	Type *string `form:"type" json:"type"`
}

// When the remove action takes effect. If not specified, defaults to on_reserve.
type V2BillingIntentActionRemoveEffectiveAtParams struct {
	// When the remove action takes effect.
	Type *string `form:"type" json:"type"`
}

// Details for a remove action.
type V2BillingIntentActionRemoveParams struct {
	// When the remove action takes effect. If not specified, defaults to on_reserve.
	EffectiveAt *V2BillingIntentActionRemoveEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// The ID of the discount rule to remove for future invoices.
	InvoiceDiscountRule *string `form:"invoice_discount_rule" json:"invoice_discount_rule,omitempty"`
	// The ID of the spend modifier rule to remove.
	SpendModifierRule *string `form:"spend_modifier_rule" json:"spend_modifier_rule,omitempty"`
	// Type of the remove action.
	Type *string `form:"type" json:"type"`
}

// When the subscribe action takes effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentActionSubscribeEffectiveAtParams struct {
	// The timestamp at which the subscribe action takes effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the subscribe action takes effect.
	Type *string `form:"type" json:"type"`
}

// Configurations for the components of the pricing plan.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsComponentConfigurationParams struct {
	// Lookup key for the pricing plan component.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// ID of the pricing plan component.
	PricingPlanComponent *string `form:"pricing_plan_component" json:"pricing_plan_component,omitempty"`
	// Quantity of the component to be used.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
}

// Overrides the behavior for license fee components when the action takes effect during the service period.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is subscribing. If not specified, defaults to prorated.
	DebitProrationBehavior *string `form:"debit_proration_behavior" json:"debit_proration_behavior"`
}

// Overrides the behavior for recurring credit grant components when the action takes effect during the service period.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorRecurringCreditGrantParams struct {
	// Controls credit grant creation behavior during partial periods. If not specified, defaults to full_credits.
	CreateBehavior *string `form:"create_behavior" json:"create_behavior"`
}

// Configurations for behaviors when the action takes effect during the service period.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Overrides the behavior for license fee components when the action takes effect during the service period.
	LicenseFee *V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// Overrides the behavior for recurring credit grant components when the action takes effect during the service period.
	RecurringCreditGrant *V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorRecurringCreditGrantParams `form:"recurring_credit_grant" json:"recurring_credit_grant,omitempty"`
	// The type of behavior to override.
	Type *string `form:"type" json:"type"`
}

// Configurations for overriding behaviors related to the subscription.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesParams struct {
	// Configurations for behaviors when the action takes effect during the service period.
	PartialPeriodBehaviors []*V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams `form:"partial_period_behaviors" json:"partial_period_behaviors"`
}

// Details for subscribing to a pricing plan.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsParams struct {
	// Configurations for the components of the pricing plan.
	ComponentConfigurations []*V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsComponentConfigurationParams `form:"component_configurations" json:"component_configurations,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Configurations for overriding behaviors related to the subscription.
	Overrides *V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesParams `form:"overrides" json:"overrides,omitempty"`
	// ID of the Pricing Plan to subscribe to.
	PricingPlan *string `form:"pricing_plan" json:"pricing_plan"`
	// Version of the Pricing Plan to use.
	PricingPlanVersion *string `form:"pricing_plan_version" json:"pricing_plan_version"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A list of up to 20 subscription items, each with an attached price.
type V2BillingIntentActionSubscribeV1SubscriptionDetailsItemParams struct {
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The ID of the price object.
	Price *string `form:"price" json:"price"`
	// Quantity for this item. If not provided, defaults to 1.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingIntentActionSubscribeV1SubscriptionDetailsItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details for subscribing to a v1 subscription.
type V2BillingIntentActionSubscribeV1SubscriptionDetailsParams struct {
	// The subscription's description, meant to be displayable to the customer.
	// Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description" json:"description,omitempty"`
	// A list of up to 20 subscription items, each with an attached price.
	Items []*V2BillingIntentActionSubscribeV1SubscriptionDetailsItemParams `form:"items" json:"items"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingIntentActionSubscribeV1SubscriptionDetailsParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details for a subscribe action.
type V2BillingIntentActionSubscribeParams struct {
	// When the invoice is collected. If not specified, defaults to on_effective_at.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the subscribe action takes effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentActionSubscribeEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for subscribing to a pricing plan.
	PricingPlanSubscriptionDetails *V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsParams `form:"pricing_plan_subscription_details" json:"pricing_plan_subscription_details,omitempty"`
	// Type of the action details.
	Type *string `form:"type" json:"type"`
	// Details for subscribing to a v1 subscription.
	V1SubscriptionDetails *V2BillingIntentActionSubscribeV1SubscriptionDetailsParams `form:"v1_subscription_details" json:"v1_subscription_details,omitempty"`
}

// Actions to be performed by this Billing Intent.
type V2BillingIntentActionParams struct {
	// Details for an apply action.
	Apply *V2BillingIntentActionApplyParams `form:"apply" json:"apply,omitempty"`
	// Details for a deactivate action.
	Deactivate *V2BillingIntentActionDeactivateParams `form:"deactivate" json:"deactivate,omitempty"`
	// Details for a modify action.
	Modify *V2BillingIntentActionModifyParams `form:"modify" json:"modify,omitempty"`
	// Details for a remove action.
	Remove *V2BillingIntentActionRemoveParams `form:"remove" json:"remove,omitempty"`
	// Details for a subscribe action.
	Subscribe *V2BillingIntentActionSubscribeParams `form:"subscribe" json:"subscribe,omitempty"`
	// Type of the Billing Intent action.
	Type *string `form:"type" json:"type"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it defaults to
// the time at which the cadence was created in UTC timezone.
type V2BillingIntentCadenceDataBillingCycleDayTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute *int64 `form:"minute" json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second *int64 `form:"second" json:"second"`
}

// Specific configuration for determining billing dates when type=day.
type V2BillingIntentCadenceDataBillingCycleDayParams struct {
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it defaults to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingIntentCadenceDataBillingCycleDayTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it defaults to
// the time at which the cadence was created in UTC timezone.
type V2BillingIntentCadenceDataBillingCycleMonthTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute *int64 `form:"minute" json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second *int64 `form:"second" json:"second"`
}

// Specific configuration for determining billing dates when type=month.
type V2BillingIntentCadenceDataBillingCycleMonthParams struct {
	// The day to anchor the billing on for a type="month" billing cycle from
	// 1-31. If this number is greater than the number of days in the month being
	// billed, this anchors to the last day of the month. If not provided,
	// this defaults to the day the cadence was created.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The month to anchor the billing on for a type="month" billing cycle from
	// 1-12. If not provided, this defaults to the month the cadence was created.
	// This setting can only be used for monthly billing cycles with `interval_count` of 2, 3, 4 or 6.
	// All occurrences are calculated from the month provided.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it defaults to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingIntentCadenceDataBillingCycleMonthTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it defaults to
// the time at which the cadence was created in UTC timezone.
type V2BillingIntentCadenceDataBillingCycleWeekTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute *int64 `form:"minute" json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second *int64 `form:"second" json:"second"`
}

// Specific configuration for determining billing dates when type=week.
type V2BillingIntentCadenceDataBillingCycleWeekParams struct {
	// The day of the week to bill the type=week billing cycle on.
	// Numbered from 1-7 for Monday to Sunday respectively, based on the ISO-8601
	// week day numbering. If not provided, this defaults to the day the
	// cadence was created.
	DayOfWeek *int64 `form:"day_of_week" json:"day_of_week"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it defaults to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingIntentCadenceDataBillingCycleWeekTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it defaults to
// the time at which the cadence was created in UTC timezone.
type V2BillingIntentCadenceDataBillingCycleYearTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute *int64 `form:"minute" json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second *int64 `form:"second" json:"second"`
}

// Specific configuration for determining billing dates when type=year.
type V2BillingIntentCadenceDataBillingCycleYearParams struct {
	// The day to anchor the billing on for a type="month" billing cycle from
	// 1-31. If this number is greater than the number of days in the month being
	// billed, this anchors to the last day of the month. If not provided,
	// this defaults to the day the cadence was created.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month,omitempty"`
	// The month to bill on from 1-12. If not provided, this defaults to the
	// month the cadence was created.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it defaults to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingIntentCadenceDataBillingCycleYearTimeParams `form:"time" json:"time,omitempty"`
}

// The billing cycle configuration for this Cadence.
type V2BillingIntentCadenceDataBillingCycleParams struct {
	// Specific configuration for determining billing dates when type=day.
	Day *V2BillingIntentCadenceDataBillingCycleDayParams `form:"day" json:"day,omitempty"`
	// The number of intervals (specified in the interval attribute) between
	// cadence billings. For example, type=month and interval_count=3 bills every
	// 3 months. If not provided, this defaults to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
	// Specific configuration for determining billing dates when type=month.
	Month *V2BillingIntentCadenceDataBillingCycleMonthParams `form:"month" json:"month,omitempty"`
	// The frequency at which a cadence bills.
	Type *string `form:"type" json:"type"`
	// Specific configuration for determining billing dates when type=week.
	Week *V2BillingIntentCadenceDataBillingCycleWeekParams `form:"week" json:"week,omitempty"`
	// Specific configuration for determining billing dates when type=year.
	Year *V2BillingIntentCadenceDataBillingCycleYearParams `form:"year" json:"year,omitempty"`
}

// Data for creating a new profile.
type V2BillingIntentCadenceDataPayerBillingProfileDataParams struct {
	// The customer to associate with the profile.
	Customer *string `form:"customer" json:"customer"`
	// The default payment method to use when billing this profile.
	// If left blank, the `PaymentMethod` from the `PaymentIntent` provided
	// on commit is used to create the profile.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
}

// Information about the payer for this Cadence.
type V2BillingIntentCadenceDataPayerParams struct {
	// The ID of the Billing Profile object which determines how a bill is paid.
	BillingProfile *string `form:"billing_profile" json:"billing_profile,omitempty"`
	// Data for creating a new profile.
	BillingProfileData *V2BillingIntentCadenceDataPayerBillingProfileDataParams `form:"billing_profile_data" json:"billing_profile_data,omitempty"`
}

// Settings that configure bill generation, which includes calculating totals, tax, and presenting invoices.
// If no setting is provided here, the settings from the customer referenced on the payer will be used.
// If no customer settings are present, the merchant default settings will be used.
type V2BillingIntentCadenceDataSettingsBillParams struct {
	// The ID of the referenced settings object.
	ID *string `form:"id" json:"id"`
	// An optional field to specify the version of the Settings to use.
	// If not provided, this defaults to the live version any time the settings are used.
	Version *string `form:"version" json:"version,omitempty"`
}

// Settings that configure and manage the behavior of collecting payments.
// If no setting is provided here, the settings from the customer referenced from the payer will be used if they exist.
// If no customer settings are present, the merchant default settings will be used.
type V2BillingIntentCadenceDataSettingsCollectionParams struct {
	// The ID of the referenced settings object.
	ID *string `form:"id" json:"id"`
	// An optional field to specify the version of the Settings to use.
	// If not provided, this defaults to the live version any time the settings are used.
	Version *string `form:"version" json:"version,omitempty"`
}

// Settings for creating the Cadence.
type V2BillingIntentCadenceDataSettingsParams struct {
	// Settings that configure bill generation, which includes calculating totals, tax, and presenting invoices.
	// If no setting is provided here, the settings from the customer referenced on the payer will be used.
	// If no customer settings are present, the merchant default settings will be used.
	Bill *V2BillingIntentCadenceDataSettingsBillParams `form:"bill" json:"bill,omitempty"`
	// Settings that configure and manage the behavior of collecting payments.
	// If no setting is provided here, the settings from the customer referenced from the payer will be used if they exist.
	// If no customer settings are present, the merchant default settings will be used.
	Collection *V2BillingIntentCadenceDataSettingsCollectionParams `form:"collection" json:"collection,omitempty"`
}

// Data for creating a new Cadence.
type V2BillingIntentCadenceDataParams struct {
	// The billing cycle configuration for this Cadence.
	BillingCycle *V2BillingIntentCadenceDataBillingCycleParams `form:"billing_cycle" json:"billing_cycle"`
	// Information about the payer for this Cadence.
	Payer *V2BillingIntentCadenceDataPayerParams `form:"payer" json:"payer"`
	// Settings for creating the Cadence.
	Settings *V2BillingIntentCadenceDataSettingsParams `form:"settings" json:"settings,omitempty"`
}

// Create a Billing Intent.
type V2BillingIntentParams struct {
	Params `form:"*"`
	// Actions to be performed by this Billing Intent.
	Actions []*V2BillingIntentActionParams `form:"actions" json:"actions,omitempty"`
	// ID of an existing Cadence to use.
	Cadence *string `form:"cadence" json:"cadence,omitempty"`
	// Data for creating a new Cadence.
	CadenceData *V2BillingIntentCadenceDataParams `form:"cadence_data" json:"cadence_data,omitempty"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency *string `form:"currency" json:"currency,omitempty"`
}

// Cancel a Billing Intent.
type V2BillingIntentCancelParams struct {
	Params `form:"*"`
}

// Commit a Billing Intent.
type V2BillingIntentCommitParams struct {
	Params `form:"*"`
	// ID of the PaymentIntent associated with this commit.
	PaymentIntent *string `form:"payment_intent" json:"payment_intent,omitempty"`
}

// Release a Billing Intent.
type V2BillingIntentReleaseReservationParams struct {
	Params `form:"*"`
}

// Reserve a Billing Intent.
type V2BillingIntentReserveParams struct {
	Params `form:"*"`
}

// Details for applying a discount.
type V2BillingIntentCreateActionApplyDiscountParams struct {
	// The ID of the Coupon to apply.
	Coupon *string `form:"coupon" json:"coupon,omitempty"`
	// The ID of the PromotionCode to apply.
	PromotionCode *string `form:"promotion_code" json:"promotion_code,omitempty"`
	// Type of the discount.
	Type *string `form:"type" json:"type"`
}

// When the apply action takes effect. If not specified, defaults to on_reserve.
type V2BillingIntentCreateActionApplyEffectiveAtParams struct {
	// The timestamp at which the apply action takes effect. Only present if type is timestamp. Only allowed for discount actions.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the apply action takes effect.
	Type *string `form:"type" json:"type"`
}

// The maximum number of times this discount can be applied for this cadence.
type V2BillingIntentCreateActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams struct {
	// The type of maximum applications configuration.
	Type *string `form:"type" json:"type"`
}

// Configuration for percentage off discount.
type V2BillingIntentCreateActionApplyInvoiceDiscountRulePercentOffParams struct {
	// The maximum number of times this discount can be applied for this cadence.
	MaximumApplications *V2BillingIntentCreateActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams `form:"maximum_applications" json:"maximum_applications"`
	// Percent that is taken off the amount. For example, a percent_off of 50.0 reduces a 100 USD amount to 50 USD.
	PercentOff *float64 `form:"percent_off,high_precision" json:"percent_off,string"`
}

// Details for applying a discount rule to future invoices.
type V2BillingIntentCreateActionApplyInvoiceDiscountRuleParams struct {
	// The entity that the discount rule applies to, for example, the cadence.
	AppliesTo *string `form:"applies_to" json:"applies_to"`
	// Configuration for percentage off discount.
	PercentOff *V2BillingIntentCreateActionApplyInvoiceDiscountRulePercentOffParams `form:"percent_off" json:"percent_off,omitempty"`
	// Type of the discount rule.
	Type *string `form:"type" json:"type"`
}

// The custom pricing unit amount.
type V2BillingIntentCreateActionApplySpendModifierRuleMaxBillingPeriodSpendAmountCustomPricingUnitParams struct {
	// The id of the custom pricing unit.
	ID *string `form:"id" json:"id,omitempty"`
	// The value of the custom pricing unit.
	Value *string `form:"value" json:"value"`
}

// The maximum amount allowed for the billing period.
type V2BillingIntentCreateActionApplySpendModifierRuleMaxBillingPeriodSpendAmountParams struct {
	// The custom pricing unit amount.
	CustomPricingUnit *V2BillingIntentCreateActionApplySpendModifierRuleMaxBillingPeriodSpendAmountCustomPricingUnitParams `form:"custom_pricing_unit" json:"custom_pricing_unit,omitempty"`
	// The type of the amount.
	Type *string `form:"type" json:"type"`
}

// The configuration for the overage rate for the custom pricing unit.
type V2BillingIntentCreateActionApplySpendModifierRuleMaxBillingPeriodSpendCustomPricingUnitOverageRateParams struct {
	// ID of the custom pricing unit overage rate.
	ID *string `form:"id" json:"id"`
}

// Details for max billing period spend modifier. Only present if type is max_billing_period_spend.
type V2BillingIntentCreateActionApplySpendModifierRuleMaxBillingPeriodSpendParams struct {
	// The maximum amount allowed for the billing period.
	Amount *V2BillingIntentCreateActionApplySpendModifierRuleMaxBillingPeriodSpendAmountParams `form:"amount" json:"amount"`
	// The configuration for the overage rate for the custom pricing unit.
	CustomPricingUnitOverageRate *V2BillingIntentCreateActionApplySpendModifierRuleMaxBillingPeriodSpendCustomPricingUnitOverageRateParams `form:"custom_pricing_unit_overage_rate" json:"custom_pricing_unit_overage_rate"`
}

// Details for applying a spend modifier rule. Only present if type is spend_modifier_rule.
type V2BillingIntentCreateActionApplySpendModifierRuleParams struct {
	// What the spend modifier applies to.
	AppliesTo *string `form:"applies_to" json:"applies_to"`
	// Details for max billing period spend modifier. Only present if type is max_billing_period_spend.
	MaxBillingPeriodSpend *V2BillingIntentCreateActionApplySpendModifierRuleMaxBillingPeriodSpendParams `form:"max_billing_period_spend" json:"max_billing_period_spend,omitempty"`
	// Type of the spend modifier.
	Type *string `form:"type" json:"type"`
}

// Details for an apply action.
type V2BillingIntentCreateActionApplyParams struct {
	// Details for applying a discount.
	Discount *V2BillingIntentCreateActionApplyDiscountParams `form:"discount" json:"discount,omitempty"`
	// When the apply action takes effect. If not specified, defaults to on_reserve.
	EffectiveAt *V2BillingIntentCreateActionApplyEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for applying a discount rule to future invoices.
	InvoiceDiscountRule *V2BillingIntentCreateActionApplyInvoiceDiscountRuleParams `form:"invoice_discount_rule" json:"invoice_discount_rule,omitempty"`
	// Details for applying a spend modifier rule. Only present if type is spend_modifier_rule.
	SpendModifierRule *V2BillingIntentCreateActionApplySpendModifierRuleParams `form:"spend_modifier_rule" json:"spend_modifier_rule,omitempty"`
	// Type of the apply action details.
	Type *string `form:"type" json:"type"`
}

// Details about why the cancellation is being requested.
type V2BillingIntentCreateActionDeactivateCancellationDetailsParams struct {
	// Additional comments about why the user canceled the subscription, if the subscription was canceled explicitly by the user.
	Comment *string `form:"comment" json:"comment,omitempty"`
	// The customer submitted reason for why they canceled, if the subscription was canceled explicitly by the user.
	Feedback *string `form:"feedback" json:"feedback,omitempty"`
}

// When the deactivate action takes effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentCreateActionDeactivateEffectiveAtParams struct {
	// The timestamp at which the deactivate action takes effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the deactivate action takes effect.
	Type *string `form:"type" json:"type"`
}

// Overrides the behavior for license fee components when the action takes effect during the service period.
type V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is deactivating. If not specified, defaults to none.
	CreditProrationBehavior *string `form:"credit_proration_behavior" json:"credit_proration_behavior"`
}

// Configurations for behaviors when the action takes effect during the service period.
type V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Overrides the behavior for license fee components when the action takes effect during the service period.
	LicenseFee *V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// The type of behavior to override.
	Type *string `form:"type" json:"type"`
}

// Configurations for overriding behaviors related to the subscription.
type V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesParams struct {
	// Configurations for behaviors when the action takes effect during the service period.
	PartialPeriodBehaviors []*V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams `form:"partial_period_behaviors" json:"partial_period_behaviors"`
}

// Details for deactivating a pricing plan subscription.
type V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsParams struct {
	// Configurations for overriding behaviors related to the subscription.
	Overrides *V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesParams `form:"overrides" json:"overrides,omitempty"`
	// ID of the pricing plan subscription to deactivate.
	PricingPlanSubscription *string `form:"pricing_plan_subscription" json:"pricing_plan_subscription"`
}

// Details for a deactivate action.
type V2BillingIntentCreateActionDeactivateParams struct {
	// Details about why the cancellation is being requested.
	CancellationDetails *V2BillingIntentCreateActionDeactivateCancellationDetailsParams `form:"cancellation_details" json:"cancellation_details,omitempty"`
	// When the invoice is collected. If not specified, the default behavior is on_effective_at.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the deactivate action takes effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentCreateActionDeactivateEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for deactivating a pricing plan subscription.
	PricingPlanSubscriptionDetails *V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsParams `form:"pricing_plan_subscription_details" json:"pricing_plan_subscription_details"`
	// Type of the action details.
	Type *string `form:"type" json:"type"`
}

// When the modify action takes effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentCreateActionModifyEffectiveAtParams struct {
	// The timestamp at which the modify action takes effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the modify action takes effect.
	Type *string `form:"type" json:"type"`
}

// New configurations for the components of the pricing plan.
type V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsComponentConfigurationParams struct {
	// Lookup key for the pricing plan component.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// ID of the pricing plan component.
	PricingPlanComponent *string `form:"pricing_plan_component" json:"pricing_plan_component,omitempty"`
	// Quantity of the component to be used.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
}

// Overrides the behavior for license fee components when the action takes effect during the service period.
type V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user modifies the subscription. If not specified, defaults to prorated.
	CreditProrationBehavior *string `form:"credit_proration_behavior" json:"credit_proration_behavior"`
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user modifies the subscription. If not specified, defaults to prorated.
	DebitProrationBehavior *string `form:"debit_proration_behavior" json:"debit_proration_behavior"`
}

// Overrides the behavior for recurring credit grant components when the action takes effect during the service period.
type V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorRecurringCreditGrantParams struct {
	// Controls credit grant creation behavior during partial periods. If not specified, defaults to full_credits.
	CreateBehavior *string `form:"create_behavior" json:"create_behavior"`
}

// Configurations for behaviors when the action takes effect during the service period.
type V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Overrides the behavior for license fee components when the action takes effect during the service period.
	LicenseFee *V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// Overrides the behavior for recurring credit grant components when the action takes effect during the service period.
	RecurringCreditGrant *V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorRecurringCreditGrantParams `form:"recurring_credit_grant" json:"recurring_credit_grant,omitempty"`
	// The type of behavior to override.
	Type *string `form:"type" json:"type"`
}

// Configurations for overriding behaviors related to the subscription.
type V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesParams struct {
	// Configurations for behaviors when the action takes effect during the service period.
	PartialPeriodBehaviors []*V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams `form:"partial_period_behaviors" json:"partial_period_behaviors"`
}

// Details for modifying a pricing plan subscription.
type V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsParams struct {
	// New configurations for the components of the pricing plan.
	ComponentConfigurations []*V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsComponentConfigurationParams `form:"component_configurations" json:"component_configurations,omitempty"`
	// The ID of the new Pricing Plan, if changing plans.
	NewPricingPlan *string `form:"new_pricing_plan" json:"new_pricing_plan,omitempty"`
	// The ID of the new Pricing Plan Version to use.
	NewPricingPlanVersion *string `form:"new_pricing_plan_version" json:"new_pricing_plan_version,omitempty"`
	// Configurations for overriding behaviors related to the subscription.
	Overrides *V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesParams `form:"overrides" json:"overrides,omitempty"`
	// The ID of the Pricing Plan Subscription to modify.
	PricingPlanSubscription *string `form:"pricing_plan_subscription" json:"pricing_plan_subscription"`
}

// Details for a modify action.
type V2BillingIntentCreateActionModifyParams struct {
	// When the invoice is collected. If not specified, the default behavior is next_billing_date.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the modify action takes effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentCreateActionModifyEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for modifying a pricing plan subscription.
	PricingPlanSubscriptionDetails *V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsParams `form:"pricing_plan_subscription_details" json:"pricing_plan_subscription_details"`
	// Type of the action details.
	Type *string `form:"type" json:"type"`
}

// When the remove action takes effect. If not specified, defaults to on_reserve.
type V2BillingIntentCreateActionRemoveEffectiveAtParams struct {
	// When the remove action takes effect.
	Type *string `form:"type" json:"type"`
}

// Details for a remove action.
type V2BillingIntentCreateActionRemoveParams struct {
	// When the remove action takes effect. If not specified, defaults to on_reserve.
	EffectiveAt *V2BillingIntentCreateActionRemoveEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// The ID of the discount rule to remove for future invoices.
	InvoiceDiscountRule *string `form:"invoice_discount_rule" json:"invoice_discount_rule,omitempty"`
	// The ID of the spend modifier rule to remove.
	SpendModifierRule *string `form:"spend_modifier_rule" json:"spend_modifier_rule,omitempty"`
	// Type of the remove action.
	Type *string `form:"type" json:"type"`
}

// When the subscribe action takes effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentCreateActionSubscribeEffectiveAtParams struct {
	// The timestamp at which the subscribe action takes effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the subscribe action takes effect.
	Type *string `form:"type" json:"type"`
}

// Configurations for the components of the pricing plan.
type V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsComponentConfigurationParams struct {
	// Lookup key for the pricing plan component.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// ID of the pricing plan component.
	PricingPlanComponent *string `form:"pricing_plan_component" json:"pricing_plan_component,omitempty"`
	// Quantity of the component to be used.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
}

// Overrides the behavior for license fee components when the action takes effect during the service period.
type V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is subscribing. If not specified, defaults to prorated.
	DebitProrationBehavior *string `form:"debit_proration_behavior" json:"debit_proration_behavior"`
}

// Overrides the behavior for recurring credit grant components when the action takes effect during the service period.
type V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorRecurringCreditGrantParams struct {
	// Controls credit grant creation behavior during partial periods. If not specified, defaults to full_credits.
	CreateBehavior *string `form:"create_behavior" json:"create_behavior"`
}

// Configurations for behaviors when the action takes effect during the service period.
type V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Overrides the behavior for license fee components when the action takes effect during the service period.
	LicenseFee *V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// Overrides the behavior for recurring credit grant components when the action takes effect during the service period.
	RecurringCreditGrant *V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorRecurringCreditGrantParams `form:"recurring_credit_grant" json:"recurring_credit_grant,omitempty"`
	// The type of behavior to override.
	Type *string `form:"type" json:"type"`
}

// Configurations for overriding behaviors related to the subscription.
type V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesParams struct {
	// Configurations for behaviors when the action takes effect during the service period.
	PartialPeriodBehaviors []*V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams `form:"partial_period_behaviors" json:"partial_period_behaviors"`
}

// Details for subscribing to a pricing plan.
type V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsParams struct {
	// Configurations for the components of the pricing plan.
	ComponentConfigurations []*V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsComponentConfigurationParams `form:"component_configurations" json:"component_configurations,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Configurations for overriding behaviors related to the subscription.
	Overrides *V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesParams `form:"overrides" json:"overrides,omitempty"`
	// ID of the Pricing Plan to subscribe to.
	PricingPlan *string `form:"pricing_plan" json:"pricing_plan"`
	// Version of the Pricing Plan to use.
	PricingPlanVersion *string `form:"pricing_plan_version" json:"pricing_plan_version"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A list of up to 20 subscription items, each with an attached price.
type V2BillingIntentCreateActionSubscribeV1SubscriptionDetailsItemParams struct {
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The ID of the price object.
	Price *string `form:"price" json:"price"`
	// Quantity for this item. If not provided, defaults to 1.
	Quantity *int64 `form:"quantity" json:"quantity,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingIntentCreateActionSubscribeV1SubscriptionDetailsItemParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details for subscribing to a v1 subscription.
type V2BillingIntentCreateActionSubscribeV1SubscriptionDetailsParams struct {
	// The subscription's description, meant to be displayable to the customer.
	// Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces and certain local payment methods UIs.
	Description *string `form:"description" json:"description,omitempty"`
	// A list of up to 20 subscription items, each with an attached price.
	Items []*V2BillingIntentCreateActionSubscribeV1SubscriptionDetailsItemParams `form:"items" json:"items"`
	// Set of key-value pairs that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingIntentCreateActionSubscribeV1SubscriptionDetailsParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Details for a subscribe action.
type V2BillingIntentCreateActionSubscribeParams struct {
	// When the invoice is collected. If not specified, defaults to on_effective_at.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the subscribe action takes effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentCreateActionSubscribeEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for subscribing to a pricing plan.
	PricingPlanSubscriptionDetails *V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsParams `form:"pricing_plan_subscription_details" json:"pricing_plan_subscription_details,omitempty"`
	// Type of the action details.
	Type *string `form:"type" json:"type"`
	// Details for subscribing to a v1 subscription.
	V1SubscriptionDetails *V2BillingIntentCreateActionSubscribeV1SubscriptionDetailsParams `form:"v1_subscription_details" json:"v1_subscription_details,omitempty"`
}

// Actions to be performed by this Billing Intent.
type V2BillingIntentCreateActionParams struct {
	// Details for an apply action.
	Apply *V2BillingIntentCreateActionApplyParams `form:"apply" json:"apply,omitempty"`
	// Details for a deactivate action.
	Deactivate *V2BillingIntentCreateActionDeactivateParams `form:"deactivate" json:"deactivate,omitempty"`
	// Details for a modify action.
	Modify *V2BillingIntentCreateActionModifyParams `form:"modify" json:"modify,omitempty"`
	// Details for a remove action.
	Remove *V2BillingIntentCreateActionRemoveParams `form:"remove" json:"remove,omitempty"`
	// Details for a subscribe action.
	Subscribe *V2BillingIntentCreateActionSubscribeParams `form:"subscribe" json:"subscribe,omitempty"`
	// Type of the Billing Intent action.
	Type *string `form:"type" json:"type"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it defaults to
// the time at which the cadence was created in UTC timezone.
type V2BillingIntentCreateCadenceDataBillingCycleDayTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute *int64 `form:"minute" json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second *int64 `form:"second" json:"second"`
}

// Specific configuration for determining billing dates when type=day.
type V2BillingIntentCreateCadenceDataBillingCycleDayParams struct {
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it defaults to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingIntentCreateCadenceDataBillingCycleDayTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it defaults to
// the time at which the cadence was created in UTC timezone.
type V2BillingIntentCreateCadenceDataBillingCycleMonthTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute *int64 `form:"minute" json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second *int64 `form:"second" json:"second"`
}

// Specific configuration for determining billing dates when type=month.
type V2BillingIntentCreateCadenceDataBillingCycleMonthParams struct {
	// The day to anchor the billing on for a type="month" billing cycle from
	// 1-31. If this number is greater than the number of days in the month being
	// billed, this anchors to the last day of the month. If not provided,
	// this defaults to the day the cadence was created.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The month to anchor the billing on for a type="month" billing cycle from
	// 1-12. If not provided, this defaults to the month the cadence was created.
	// This setting can only be used for monthly billing cycles with `interval_count` of 2, 3, 4 or 6.
	// All occurrences are calculated from the month provided.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it defaults to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingIntentCreateCadenceDataBillingCycleMonthTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it defaults to
// the time at which the cadence was created in UTC timezone.
type V2BillingIntentCreateCadenceDataBillingCycleWeekTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute *int64 `form:"minute" json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second *int64 `form:"second" json:"second"`
}

// Specific configuration for determining billing dates when type=week.
type V2BillingIntentCreateCadenceDataBillingCycleWeekParams struct {
	// The day of the week to bill the type=week billing cycle on.
	// Numbered from 1-7 for Monday to Sunday respectively, based on the ISO-8601
	// week day numbering. If not provided, this defaults to the day the
	// cadence was created.
	DayOfWeek *int64 `form:"day_of_week" json:"day_of_week"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it defaults to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingIntentCreateCadenceDataBillingCycleWeekTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the billing cycle ends.
// This field is optional, and if not provided, it defaults to
// the time at which the cadence was created in UTC timezone.
type V2BillingIntentCreateCadenceDataBillingCycleYearTimeParams struct {
	// The hour at which the billing cycle ends.
	// This must be an integer between 0 and 23, inclusive.
	// 0 represents midnight, and 23 represents 11 PM.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Minute *int64 `form:"minute" json:"minute"`
	// The second at which the billing cycle ends.
	// Must be an integer between 0 and 59, inclusive.
	Second *int64 `form:"second" json:"second"`
}

// Specific configuration for determining billing dates when type=year.
type V2BillingIntentCreateCadenceDataBillingCycleYearParams struct {
	// The day to anchor the billing on for a type="month" billing cycle from
	// 1-31. If this number is greater than the number of days in the month being
	// billed, this anchors to the last day of the month. If not provided,
	// this defaults to the day the cadence was created.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month,omitempty"`
	// The month to bill on from 1-12. If not provided, this defaults to the
	// month the cadence was created.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// The time at which the billing cycle ends.
	// This field is optional, and if not provided, it defaults to
	// the time at which the cadence was created in UTC timezone.
	Time *V2BillingIntentCreateCadenceDataBillingCycleYearTimeParams `form:"time" json:"time,omitempty"`
}

// The billing cycle configuration for this Cadence.
type V2BillingIntentCreateCadenceDataBillingCycleParams struct {
	// Specific configuration for determining billing dates when type=day.
	Day *V2BillingIntentCreateCadenceDataBillingCycleDayParams `form:"day" json:"day,omitempty"`
	// The number of intervals (specified in the interval attribute) between
	// cadence billings. For example, type=month and interval_count=3 bills every
	// 3 months. If not provided, this defaults to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
	// Specific configuration for determining billing dates when type=month.
	Month *V2BillingIntentCreateCadenceDataBillingCycleMonthParams `form:"month" json:"month,omitempty"`
	// The frequency at which a cadence bills.
	Type *string `form:"type" json:"type"`
	// Specific configuration for determining billing dates when type=week.
	Week *V2BillingIntentCreateCadenceDataBillingCycleWeekParams `form:"week" json:"week,omitempty"`
	// Specific configuration for determining billing dates when type=year.
	Year *V2BillingIntentCreateCadenceDataBillingCycleYearParams `form:"year" json:"year,omitempty"`
}

// Data for creating a new profile.
type V2BillingIntentCreateCadenceDataPayerBillingProfileDataParams struct {
	// The customer to associate with the profile.
	Customer *string `form:"customer" json:"customer"`
	// The default payment method to use when billing this profile.
	// If left blank, the `PaymentMethod` from the `PaymentIntent` provided
	// on commit is used to create the profile.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
}

// Information about the payer for this Cadence.
type V2BillingIntentCreateCadenceDataPayerParams struct {
	// The ID of the Billing Profile object which determines how a bill is paid.
	BillingProfile *string `form:"billing_profile" json:"billing_profile,omitempty"`
	// Data for creating a new profile.
	BillingProfileData *V2BillingIntentCreateCadenceDataPayerBillingProfileDataParams `form:"billing_profile_data" json:"billing_profile_data,omitempty"`
}

// Settings that configure bill generation, which includes calculating totals, tax, and presenting invoices.
// If no setting is provided here, the settings from the customer referenced on the payer will be used.
// If no customer settings are present, the merchant default settings will be used.
type V2BillingIntentCreateCadenceDataSettingsBillParams struct {
	// The ID of the referenced settings object.
	ID *string `form:"id" json:"id"`
	// An optional field to specify the version of the Settings to use.
	// If not provided, this defaults to the live version any time the settings are used.
	Version *string `form:"version" json:"version,omitempty"`
}

// Settings that configure and manage the behavior of collecting payments.
// If no setting is provided here, the settings from the customer referenced from the payer will be used if they exist.
// If no customer settings are present, the merchant default settings will be used.
type V2BillingIntentCreateCadenceDataSettingsCollectionParams struct {
	// The ID of the referenced settings object.
	ID *string `form:"id" json:"id"`
	// An optional field to specify the version of the Settings to use.
	// If not provided, this defaults to the live version any time the settings are used.
	Version *string `form:"version" json:"version,omitempty"`
}

// Settings for creating the Cadence.
type V2BillingIntentCreateCadenceDataSettingsParams struct {
	// Settings that configure bill generation, which includes calculating totals, tax, and presenting invoices.
	// If no setting is provided here, the settings from the customer referenced on the payer will be used.
	// If no customer settings are present, the merchant default settings will be used.
	Bill *V2BillingIntentCreateCadenceDataSettingsBillParams `form:"bill" json:"bill,omitempty"`
	// Settings that configure and manage the behavior of collecting payments.
	// If no setting is provided here, the settings from the customer referenced from the payer will be used if they exist.
	// If no customer settings are present, the merchant default settings will be used.
	Collection *V2BillingIntentCreateCadenceDataSettingsCollectionParams `form:"collection" json:"collection,omitempty"`
}

// Data for creating a new Cadence.
type V2BillingIntentCreateCadenceDataParams struct {
	// The billing cycle configuration for this Cadence.
	BillingCycle *V2BillingIntentCreateCadenceDataBillingCycleParams `form:"billing_cycle" json:"billing_cycle"`
	// Information about the payer for this Cadence.
	Payer *V2BillingIntentCreateCadenceDataPayerParams `form:"payer" json:"payer"`
	// Settings for creating the Cadence.
	Settings *V2BillingIntentCreateCadenceDataSettingsParams `form:"settings" json:"settings,omitempty"`
}

// Create a Billing Intent.
type V2BillingIntentCreateParams struct {
	Params `form:"*"`
	// Actions to be performed by this Billing Intent.
	Actions []*V2BillingIntentCreateActionParams `form:"actions" json:"actions"`
	// ID of an existing Cadence to use.
	Cadence *string `form:"cadence" json:"cadence,omitempty"`
	// Data for creating a new Cadence.
	CadenceData *V2BillingIntentCreateCadenceDataParams `form:"cadence_data" json:"cadence_data,omitempty"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency *string `form:"currency" json:"currency"`
}

// Retrieve a Billing Intent.
type V2BillingIntentRetrieveParams struct {
	Params `form:"*"`
}
