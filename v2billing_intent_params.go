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

// The maximum number of times this discount can be applied for this cadence.
type V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams struct {
	// The type of maximum applications configuration.
	Type *string `form:"type" json:"type"`
}

// Configuration for percentage off discount.
type V2BillingIntentActionApplyInvoiceDiscountRulePercentOffParams struct {
	// The maximum number of times this discount can be applied for this cadence.
	MaximumApplications *V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams `form:"maximum_applications" json:"maximum_applications"`
	// Percent that will be taken off of the amount. For example, percent_off of 50.0 will make $100 amount $50 instead.
	PercentOff *string `form:"percent_off" json:"percent_off"`
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

// Details for an apply action.
type V2BillingIntentActionApplyParams struct {
	// Details for applying a discount rule to future invoices.
	InvoiceDiscountRule *V2BillingIntentActionApplyInvoiceDiscountRuleParams `form:"invoice_discount_rule" json:"invoice_discount_rule,omitempty"`
	// Type of the apply action details.
	Type *string `form:"type" json:"type"`
}

// When the deactivate action will take effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentActionDeactivateEffectiveAtParams struct {
	// The timestamp at which the deactivate action will take effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the deactivate action will take effect.
	Type *string `form:"type" json:"type"`
}

// Override for the license fee.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is deactivating.
	CreditProrationBehavior *string `form:"credit_proration_behavior" json:"credit_proration_behavior"`
}

// Override for the partial period behavior.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Override for the license fee.
	LicenseFee *V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// Type of the partial period behavior override.
	Type *string `form:"type" json:"type"`
}

// Allows users to override the partial period behavior.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesParams struct {
	// Override for the partial period behavior.
	PartialPeriodBehaviors []*V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams `form:"partial_period_behaviors" json:"partial_period_behaviors"`
}

// Details for deactivating a pricing plan subscription.
type V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsParams struct {
	// Allows users to override the partial period behavior.
	Overrides *V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesParams `form:"overrides" json:"overrides,omitempty"`
	// ID of the pricing plan subscription to deactivate.
	PricingPlanSubscription *string `form:"pricing_plan_subscription" json:"pricing_plan_subscription"`
}

// Details for a deactivate action.
type V2BillingIntentActionDeactivateParams struct {
	// Allows users to override the collect at behavior.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the deactivate action will take effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentActionDeactivateEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for deactivating a pricing plan subscription.
	PricingPlanSubscriptionDetails *V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsParams `form:"pricing_plan_subscription_details" json:"pricing_plan_subscription_details"`
	// Type of the action details.
	Type *string `form:"type" json:"type"`
}

// When the modify action will take effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentActionModifyEffectiveAtParams struct {
	// The timestamp at which the modify action will take effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the modify action will take effect.
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

// Override for the license fee.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is upgrading.
	CreditProrationBehavior *string `form:"credit_proration_behavior" json:"credit_proration_behavior"`
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is downgrading.
	DebitProrationBehavior *string `form:"debit_proration_behavior" json:"debit_proration_behavior"`
}

// Override for the partial period behavior.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Override for the license fee.
	LicenseFee *V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// Type of the partial period behavior override.
	Type *string `form:"type" json:"type"`
}

// Allows users to override the partial period behavior.
type V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesParams struct {
	// Override for the partial period behavior.
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
	// Allows users to override the partial period behavior.
	Overrides *V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesParams `form:"overrides" json:"overrides,omitempty"`
	// The ID of the Pricing Plan Subscription to modify.
	PricingPlanSubscription *string `form:"pricing_plan_subscription" json:"pricing_plan_subscription"`
}

// Details for a modify action.
type V2BillingIntentActionModifyParams struct {
	// Allows users to override the collect at behavior.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the modify action will take effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentActionModifyEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for modifying a pricing plan subscription.
	PricingPlanSubscriptionDetails *V2BillingIntentActionModifyPricingPlanSubscriptionDetailsParams `form:"pricing_plan_subscription_details" json:"pricing_plan_subscription_details"`
	// Type of the action details.
	Type *string `form:"type" json:"type"`
}

// Details for a remove action.
type V2BillingIntentActionRemoveParams struct {
	// The ID of the discount rule to remove for future invoices.
	InvoiceDiscountRule *string `form:"invoice_discount_rule" json:"invoice_discount_rule,omitempty"`
	// Type of the remove action.
	Type *string `form:"type" json:"type"`
}

// When the subscribe action will take effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentActionSubscribeEffectiveAtParams struct {
	// The timestamp at which the subscribe action will take effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the subscribe action will take effect.
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

// Override for the license fee.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is subscribing.
	DebitProrationBehavior *string `form:"debit_proration_behavior" json:"debit_proration_behavior"`
}

// Override for the partial period behavior.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Override for the license fee.
	LicenseFee *V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// Type of the partial period behavior override.
	Type *string `form:"type" json:"type"`
}

// Allows users to override the partial period behavior.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesParams struct {
	// Override for the partial period behavior.
	PartialPeriodBehaviors []*V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams `form:"partial_period_behaviors" json:"partial_period_behaviors"`
}

// Details for subscribing to a pricing plan.
type V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsParams struct {
	// Configurations for the components of the pricing plan.
	ComponentConfigurations []*V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsComponentConfigurationParams `form:"component_configurations" json:"component_configurations,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Allows users to override the partial period behavior.
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
	// Quantity for this item. If not provided, will default to 1.
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
	// Allows users to override the collect at behavior.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the subscribe action will take effect. If not specified, the default behavior is on_reserve.
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

// Create a Billing Intent.
type V2BillingIntentParams struct {
	Params `form:"*"`
	// Actions to be performed by this Billing Intent.
	Actions []*V2BillingIntentActionParams `form:"actions" json:"actions,omitempty"`
	// ID of an existing Cadence to use.
	Cadence *string `form:"cadence" json:"cadence,omitempty"`
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

// The maximum number of times this discount can be applied for this cadence.
type V2BillingIntentCreateActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams struct {
	// The type of maximum applications configuration.
	Type *string `form:"type" json:"type"`
}

// Configuration for percentage off discount.
type V2BillingIntentCreateActionApplyInvoiceDiscountRulePercentOffParams struct {
	// The maximum number of times this discount can be applied for this cadence.
	MaximumApplications *V2BillingIntentCreateActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams `form:"maximum_applications" json:"maximum_applications"`
	// Percent that will be taken off of the amount. For example, percent_off of 50.0 will make $100 amount $50 instead.
	PercentOff *string `form:"percent_off" json:"percent_off"`
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

// Details for an apply action.
type V2BillingIntentCreateActionApplyParams struct {
	// Details for applying a discount rule to future invoices.
	InvoiceDiscountRule *V2BillingIntentCreateActionApplyInvoiceDiscountRuleParams `form:"invoice_discount_rule" json:"invoice_discount_rule,omitempty"`
	// Type of the apply action details.
	Type *string `form:"type" json:"type"`
}

// When the deactivate action will take effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentCreateActionDeactivateEffectiveAtParams struct {
	// The timestamp at which the deactivate action will take effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the deactivate action will take effect.
	Type *string `form:"type" json:"type"`
}

// Override for the license fee.
type V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is deactivating.
	CreditProrationBehavior *string `form:"credit_proration_behavior" json:"credit_proration_behavior"`
}

// Override for the partial period behavior.
type V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Override for the license fee.
	LicenseFee *V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// Type of the partial period behavior override.
	Type *string `form:"type" json:"type"`
}

// Allows users to override the partial period behavior.
type V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesParams struct {
	// Override for the partial period behavior.
	PartialPeriodBehaviors []*V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams `form:"partial_period_behaviors" json:"partial_period_behaviors"`
}

// Details for deactivating a pricing plan subscription.
type V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsParams struct {
	// Allows users to override the partial period behavior.
	Overrides *V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsOverridesParams `form:"overrides" json:"overrides,omitempty"`
	// ID of the pricing plan subscription to deactivate.
	PricingPlanSubscription *string `form:"pricing_plan_subscription" json:"pricing_plan_subscription"`
}

// Details for a deactivate action.
type V2BillingIntentCreateActionDeactivateParams struct {
	// Allows users to override the collect at behavior.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the deactivate action will take effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentCreateActionDeactivateEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for deactivating a pricing plan subscription.
	PricingPlanSubscriptionDetails *V2BillingIntentCreateActionDeactivatePricingPlanSubscriptionDetailsParams `form:"pricing_plan_subscription_details" json:"pricing_plan_subscription_details"`
	// Type of the action details.
	Type *string `form:"type" json:"type"`
}

// When the modify action will take effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentCreateActionModifyEffectiveAtParams struct {
	// The timestamp at which the modify action will take effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the modify action will take effect.
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

// Override for the license fee.
type V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is upgrading.
	CreditProrationBehavior *string `form:"credit_proration_behavior" json:"credit_proration_behavior"`
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is downgrading.
	DebitProrationBehavior *string `form:"debit_proration_behavior" json:"debit_proration_behavior"`
}

// Override for the partial period behavior.
type V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Override for the license fee.
	LicenseFee *V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// Type of the partial period behavior override.
	Type *string `form:"type" json:"type"`
}

// Allows users to override the partial period behavior.
type V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesParams struct {
	// Override for the partial period behavior.
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
	// Allows users to override the partial period behavior.
	Overrides *V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsOverridesParams `form:"overrides" json:"overrides,omitempty"`
	// The ID of the Pricing Plan Subscription to modify.
	PricingPlanSubscription *string `form:"pricing_plan_subscription" json:"pricing_plan_subscription"`
}

// Details for a modify action.
type V2BillingIntentCreateActionModifyParams struct {
	// Allows users to override the collect at behavior.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the modify action will take effect. If not specified, the default behavior is on_reserve.
	EffectiveAt *V2BillingIntentCreateActionModifyEffectiveAtParams `form:"effective_at" json:"effective_at,omitempty"`
	// Details for modifying a pricing plan subscription.
	PricingPlanSubscriptionDetails *V2BillingIntentCreateActionModifyPricingPlanSubscriptionDetailsParams `form:"pricing_plan_subscription_details" json:"pricing_plan_subscription_details"`
	// Type of the action details.
	Type *string `form:"type" json:"type"`
}

// Details for a remove action.
type V2BillingIntentCreateActionRemoveParams struct {
	// The ID of the discount rule to remove for future invoices.
	InvoiceDiscountRule *string `form:"invoice_discount_rule" json:"invoice_discount_rule,omitempty"`
	// Type of the remove action.
	Type *string `form:"type" json:"type"`
}

// When the subscribe action will take effect. If not specified, the default behavior is on_reserve.
type V2BillingIntentCreateActionSubscribeEffectiveAtParams struct {
	// The timestamp at which the subscribe action will take effect. Only present if type is timestamp.
	Timestamp *time.Time `form:"timestamp" json:"timestamp,omitempty"`
	// When the subscribe action will take effect.
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

// Override for the license fee.
type V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams struct {
	// The proration behavior for the partial servicing period. Defines how we prorate the license fee when the user is subscribing.
	DebitProrationBehavior *string `form:"debit_proration_behavior" json:"debit_proration_behavior"`
}

// Override for the partial period behavior.
type V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams struct {
	// Override for the license fee.
	LicenseFee *V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorLicenseFeeParams `form:"license_fee" json:"license_fee,omitempty"`
	// Type of the partial period behavior override.
	Type *string `form:"type" json:"type"`
}

// Allows users to override the partial period behavior.
type V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesParams struct {
	// Override for the partial period behavior.
	PartialPeriodBehaviors []*V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams `form:"partial_period_behaviors" json:"partial_period_behaviors"`
}

// Details for subscribing to a pricing plan.
type V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsParams struct {
	// Configurations for the components of the pricing plan.
	ComponentConfigurations []*V2BillingIntentCreateActionSubscribePricingPlanSubscriptionDetailsComponentConfigurationParams `form:"component_configurations" json:"component_configurations,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Allows users to override the partial period behavior.
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
	// Quantity for this item. If not provided, will default to 1.
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
	// Allows users to override the collect at behavior.
	CollectAt *string `form:"collect_at" json:"collect_at,omitempty"`
	// When the subscribe action will take effect. If not specified, the default behavior is on_reserve.
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

// Create a Billing Intent.
type V2BillingIntentCreateParams struct {
	Params `form:"*"`
	// Actions to be performed by this Billing Intent.
	Actions []*V2BillingIntentCreateActionParams `form:"actions" json:"actions"`
	// ID of an existing Cadence to use.
	Cadence *string `form:"cadence" json:"cadence,omitempty"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency *string `form:"currency" json:"currency"`
}

// Retrieve a Billing Intent.
type V2BillingIntentRetrieveParams struct {
	Params `form:"*"`
}
