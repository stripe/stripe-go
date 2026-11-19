//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"github.com/shopspring/decimal"
	"time"
)

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

// The maximum number of times this discount can be applied for this billing group.
type V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams struct {
	// The type of maximum applications configuration.
	Type *string `form:"type" json:"type"`
}

// Configuration for percentage off discount.
type V2BillingIntentActionApplyInvoiceDiscountRulePercentOffParams struct {
	// The maximum number of times this discount can be applied for this billing group.
	MaximumApplications *V2BillingIntentActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams `form:"maximum_applications" json:"maximum_applications"`
	// Percent that is taken off the amount. For example, a percent_off of 50.0 reduces a 100 USD amount to 50 USD.
	PercentOff *decimal.Decimal `form:"percent_off" json:"percent_off"`
}

// Details for applying a discount rule to future invoices.
type V2BillingIntentActionApplyInvoiceDiscountRuleParams struct {
	// The entity that the discount rule applies to, for example, the billing group.
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

// Controls emails for when the payment is due. Omit to use the applicable default.
type V2BillingIntentBillingGroupDataBillingSettingsEmailDeliveryPaymentDueParams struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled *bool `form:"enabled" json:"enabled,omitempty"`
	// If true the payment link to hosted invoice page would be included in email and PDF of the invoice.
	IncludePaymentLink *bool `form:"include_payment_link" json:"include_payment_link,omitempty"`
}

// Email delivery setting.
type V2BillingIntentBillingGroupDataBillingSettingsEmailDeliveryParams struct {
	// Controls emails for when the payment is due. Omit to use the applicable default.
	PaymentDue *V2BillingIntentBillingGroupDataBillingSettingsEmailDeliveryPaymentDueParams `form:"payment_due" json:"payment_due,omitempty"`
}

// Settings for invoice presentation.
type V2BillingIntentBillingGroupDataBillingSettingsInvoicePresentationParams struct {
	// The ID of the invoice rendering template to be used when generating invoices. Omit to use the applicable default.
	InvoiceRenderingTemplate *string `form:"invoice_rendering_template" json:"invoice_rendering_template,omitempty"`
}

// The amount of time until the invoice is overdue for payment.
type V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueParams struct {
	// The interval unit for the time until due.
	Interval *string `form:"interval" json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30, the invoice is due in 30 days.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Additional fields for Mandate creation.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type" json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitParams struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactParams struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage *string `form:"preferred_language" json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount" json:"amount,string,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType *string `form:"amount_type" json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description *string `form:"description" json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network *string `form:"network" json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure" json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information.
	Country *string `form:"country" json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer" json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type *string `form:"type" json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType *string `form:"funding_type" json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsKonbiniParams struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsSEPADebitParams struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []*string `form:"account_subcategories" json:"account_subcategories,omitempty"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters" json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []*string `form:"permissions" json:"permissions,omitempty"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch" json:"prefetch,omitempty"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections" json:"financial_connections"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method"`
}

// Payment Method specific configuration to be stored on the object.
type V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsParams struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitParams `form:"acss_debit" json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactParams `form:"bancontact" json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceParams `form:"customer_balance" json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsKonbiniParams `form:"konbini" json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsSEPADebitParams `form:"sepa_debit" json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// Settings for calculating tax.
type V2BillingIntentBillingGroupDataBillingSettingsTaxParams struct {
	// Determines if tax is calculated automatically based on a product tax code or manually based on rules defined by the business. Omit to use the applicable default.
	Type *string `form:"type" json:"type,omitempty"`
}

// Billing setting overrides for the group. If omitted, the applicable defaults are used.
type V2BillingIntentBillingGroupDataBillingSettingsParams struct {
	// How Stripe collects payment.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// Email delivery setting.
	EmailDelivery *V2BillingIntentBillingGroupDataBillingSettingsEmailDeliveryParams `form:"email_delivery" json:"email_delivery,omitempty"`
	// Settings for invoice presentation.
	InvoicePresentation *V2BillingIntentBillingGroupDataBillingSettingsInvoicePresentationParams `form:"invoice_presentation" json:"invoice_presentation,omitempty"`
	// The amount of time until the invoice is overdue for payment.
	InvoiceTimeUntilDue *V2BillingIntentBillingGroupDataBillingSettingsInvoiceTimeUntilDueParams `form:"invoice_time_until_due" json:"invoice_time_until_due,omitempty"`
	// The ID of the payment method configuration, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration to be stored on the object.
	PaymentMethodOptions *V2BillingIntentBillingGroupDataBillingSettingsPaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// Settings for calculating tax.
	Tax *V2BillingIntentBillingGroupDataBillingSettingsTaxParams `form:"tax" json:"tax,omitempty"`
}

// The time at which the daily grouping interval begins.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceDayTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Daily cadence details.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceDayParams struct {
	// The time at which the daily grouping interval begins.
	Time *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceDayTimeParams `form:"time" json:"time"`
}

// The time at which the monthly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceMonthTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Monthly cadence details.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceMonthParams struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The month of year, from 1 through 12, from which occurrences are calculated. Defaults to the month when the billing group was created.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// The time at which the monthly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceMonthTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the weekly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceWeekTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Weekly cadence details.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceWeekParams struct {
	// The ISO-8601 day of week, from 1 for Monday through 7 for Sunday.
	DayOfWeek *int64 `form:"day_of_week" json:"day_of_week"`
	// The time at which the weekly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceWeekTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the yearly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceYearTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Yearly cadence details.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceYearParams struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The month of year, from 1 through 12.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year"`
	// The time at which the yearly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceYearTimeParams `form:"time" json:"time,omitempty"`
}

// The cadence configuration when type is `by_cadence`.
type V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceParams struct {
	// Timestamp that anchors the recurring grouping cadence.
	BillingCycleAnchor *time.Time `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Daily cadence details.
	Day *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceDayParams `form:"day" json:"day,omitempty"`
	// The number of interval units between groups. Defaults to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
	// Monthly cadence details.
	Month *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceMonthParams `form:"month" json:"month,omitempty"`
	// The interval unit for the cadence.
	Type *string `form:"type" json:"type"`
	// Weekly cadence details.
	Week *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceWeekParams `form:"week" json:"week,omitempty"`
	// Yearly cadence details.
	Year *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceYearParams `form:"year" json:"year,omitempty"`
}

// The strategy used to group member invoices.
type V2BillingIntentBillingGroupDataGroupingBehaviorParams struct {
	// The cadence configuration when type is `by_cadence`.
	ByCadence *V2BillingIntentBillingGroupDataGroupingBehaviorByCadenceParams `form:"by_cadence" json:"by_cadence,omitempty"`
	// The grouping strategy.
	Type *string `form:"type" json:"type"`
}

// Data for creating a new Billing Group.
type V2BillingIntentBillingGroupDataParams struct {
	// Billing setting overrides for the group. If omitted, the applicable defaults are used.
	BillingSettings *V2BillingIntentBillingGroupDataBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// The ID of the Customer object whose invoices are grouped.
	Customer *string `form:"customer" json:"customer"`
	// The ID of the PaymentMethod object used by default. If unset, the customer's default payment method is used.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// The strategy used to group member invoices.
	GroupingBehavior *V2BillingIntentBillingGroupDataGroupingBehaviorParams `form:"grouping_behavior" json:"grouping_behavior,omitempty"`
	// An internal key you can use to find a particular billing group. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
}

// Create a Billing Intent.
type V2BillingIntentParams struct {
	Params `form:"*"`
	// Actions to be performed by this Billing Intent.
	Actions []*V2BillingIntentActionParams `form:"actions" json:"actions,omitempty"`
	// ID of an existing Billing Group to use.
	BillingGroup *string `form:"billing_group" json:"billing_group,omitempty"`
	// Data for creating a new Billing Group.
	BillingGroupData *V2BillingIntentBillingGroupDataParams `form:"billing_group_data" json:"billing_group_data,omitempty"`
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

// The maximum number of times this discount can be applied for this billing group.
type V2BillingIntentCreateActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams struct {
	// The type of maximum applications configuration.
	Type *string `form:"type" json:"type"`
}

// Configuration for percentage off discount.
type V2BillingIntentCreateActionApplyInvoiceDiscountRulePercentOffParams struct {
	// The maximum number of times this discount can be applied for this billing group.
	MaximumApplications *V2BillingIntentCreateActionApplyInvoiceDiscountRulePercentOffMaximumApplicationsParams `form:"maximum_applications" json:"maximum_applications"`
	// Percent that is taken off the amount. For example, a percent_off of 50.0 reduces a 100 USD amount to 50 USD.
	PercentOff *decimal.Decimal `form:"percent_off" json:"percent_off"`
}

// Details for applying a discount rule to future invoices.
type V2BillingIntentCreateActionApplyInvoiceDiscountRuleParams struct {
	// The entity that the discount rule applies to, for example, the billing group.
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

// Controls emails for when the payment is due. Omit to use the applicable default.
type V2BillingIntentCreateBillingGroupDataBillingSettingsEmailDeliveryPaymentDueParams struct {
	// If true an email for the invoice would be generated and sent out.
	Enabled *bool `form:"enabled" json:"enabled,omitempty"`
	// If true the payment link to hosted invoice page would be included in email and PDF of the invoice.
	IncludePaymentLink *bool `form:"include_payment_link" json:"include_payment_link,omitempty"`
}

// Email delivery setting.
type V2BillingIntentCreateBillingGroupDataBillingSettingsEmailDeliveryParams struct {
	// Controls emails for when the payment is due. Omit to use the applicable default.
	PaymentDue *V2BillingIntentCreateBillingGroupDataBillingSettingsEmailDeliveryPaymentDueParams `form:"payment_due" json:"payment_due,omitempty"`
}

// Settings for invoice presentation.
type V2BillingIntentCreateBillingGroupDataBillingSettingsInvoicePresentationParams struct {
	// The ID of the invoice rendering template to be used when generating invoices. Omit to use the applicable default.
	InvoiceRenderingTemplate *string `form:"invoice_rendering_template" json:"invoice_rendering_template,omitempty"`
}

// The amount of time until the invoice is overdue for payment.
type V2BillingIntentCreateBillingGroupDataBillingSettingsInvoiceTimeUntilDueParams struct {
	// The interval unit for the time until due.
	Interval *string `form:"interval" json:"interval"`
	// The number of interval units. For example, if interval=day and interval_count=30, the invoice is due in 30 days.
	IntervalCount *int64 `form:"interval_count" json:"interval_count"`
}

// Additional fields for Mandate creation.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams struct {
	// Transaction type of the mandate.
	TransactionType *string `form:"transaction_type" json:"transaction_type,omitempty"`
}

// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitParams struct {
	// Additional fields for Mandate creation.
	MandateOptions *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method,omitempty"`
}

// This sub-hash contains details about the Bancontact payment method.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactParams struct {
	// Preferred language of the Bancontact authorization page that the customer is redirected to.
	PreferredLanguage *string `form:"preferred_language" json:"preferred_language,omitempty"`
}

// Configuration options for setting up an eMandate for cards issued in India.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount" json:"amount,string,omitempty"`
	// The AmountType for the mandate. One of `fixed` or `maximum`.
	AmountType *string `form:"amount_type" json:"amount_type,omitempty"`
	// A description of the mandate that is meant to be displayed to the customer.
	Description *string `form:"description" json:"description,omitempty"`
}

// This sub-hash contains details about the Card payment method options.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsCardParams struct {
	// Configuration options for setting up an eMandate for cards issued in India.
	MandateOptions *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsCardMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
	// Selected network to process the payment on. Depends on the available networks of the card.
	Network *string `form:"network" json:"network,omitempty"`
	// An advanced option 3D Secure. We strongly recommend that you rely on our SCA Engine to automatically prompt your customers
	// for authentication based on risk level and [other requirements](https://docs.stripe.com/strong-customer-authentication).
	// However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option.
	// Read our guide on [manually requesting 3D Secure](https://docs.stripe.com/payments/3d-secure/authentication-flow#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure *string `form:"request_three_d_secure" json:"request_three_d_secure,omitempty"`
}

// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams struct {
	// The desired country code of the bank account information.
	Country *string `form:"country" json:"country"`
}

// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams struct {
	// Configuration for `eu_bank_transfer` funding type. Required if `type` is `eu_bank_transfer`.
	EUBankTransfer *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferEUBankTransferParams `form:"eu_bank_transfer" json:"eu_bank_transfer,omitempty"`
	// The bank transfer type that can be used for funding.
	Type *string `form:"type" json:"type,omitempty"`
}

// This sub-hash contains details about the Bank transfer payment method options.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceParams struct {
	// Configuration for the bank transfer funding type, if the `funding_type` is set to `bank_transfer`.
	BankTransfer *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceBankTransferParams `form:"bank_transfer" json:"bank_transfer,omitempty"`
	// The funding method type to be used when there are not enough funds in the customer balance. Currently the only supported value is `bank_transfer`.
	FundingType *string `form:"funding_type" json:"funding_type,omitempty"`
}

// This sub-hash contains details about the Konbini payment method options.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsKonbiniParams struct{}

// This sub-hash contains details about the SEPA Direct Debit payment method options.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsSEPADebitParams struct{}

// Provide filters for the linked accounts that the customer can select for the payment method.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams struct {
	// The account subcategories to use to filter for selectable accounts.
	AccountSubcategories []*string `form:"account_subcategories" json:"account_subcategories,omitempty"`
}

// Additional fields for Financial Connections Session creation.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams struct {
	// Provide filters for the linked accounts that the customer can select for the payment method.
	Filters *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsFiltersParams `form:"filters" json:"filters,omitempty"`
	// The list of permissions to request. If this parameter is passed, the `payment_method` permission must be included.
	Permissions []*string `form:"permissions" json:"permissions,omitempty"`
	// List of data features that you would like to retrieve upon account creation.
	Prefetch []*string `form:"prefetch" json:"prefetch,omitempty"`
}

// This sub-hash contains details about the ACH direct debit payment method options.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountParams struct {
	// Additional fields for Financial Connections Session creation.
	FinancialConnections *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountFinancialConnectionsParams `form:"financial_connections" json:"financial_connections"`
	// Verification method.
	VerificationMethod *string `form:"verification_method" json:"verification_method"`
}

// Payment Method specific configuration to be stored on the object.
type V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsParams struct {
	// This sub-hash contains details about the Canadian pre-authorized debit payment method options.
	ACSSDebit *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsACSSDebitParams `form:"acss_debit" json:"acss_debit,omitempty"`
	// This sub-hash contains details about the Bancontact payment method.
	Bancontact *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsBancontactParams `form:"bancontact" json:"bancontact,omitempty"`
	// This sub-hash contains details about the Card payment method options.
	Card *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
	// This sub-hash contains details about the Bank transfer payment method options.
	CustomerBalance *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsCustomerBalanceParams `form:"customer_balance" json:"customer_balance,omitempty"`
	// This sub-hash contains details about the Konbini payment method options.
	Konbini *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsKonbiniParams `form:"konbini" json:"konbini,omitempty"`
	// This sub-hash contains details about the SEPA Direct Debit payment method options.
	SEPADebit *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsSEPADebitParams `form:"sepa_debit" json:"sepa_debit,omitempty"`
	// This sub-hash contains details about the ACH direct debit payment method options.
	USBankAccount *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
}

// Settings for calculating tax.
type V2BillingIntentCreateBillingGroupDataBillingSettingsTaxParams struct {
	// Determines if tax is calculated automatically based on a product tax code or manually based on rules defined by the business. Omit to use the applicable default.
	Type *string `form:"type" json:"type,omitempty"`
}

// Billing setting overrides for the group. If omitted, the applicable defaults are used.
type V2BillingIntentCreateBillingGroupDataBillingSettingsParams struct {
	// How Stripe collects payment.
	CollectionMethod *string `form:"collection_method" json:"collection_method,omitempty"`
	// Email delivery setting.
	EmailDelivery *V2BillingIntentCreateBillingGroupDataBillingSettingsEmailDeliveryParams `form:"email_delivery" json:"email_delivery,omitempty"`
	// Settings for invoice presentation.
	InvoicePresentation *V2BillingIntentCreateBillingGroupDataBillingSettingsInvoicePresentationParams `form:"invoice_presentation" json:"invoice_presentation,omitempty"`
	// The amount of time until the invoice is overdue for payment.
	InvoiceTimeUntilDue *V2BillingIntentCreateBillingGroupDataBillingSettingsInvoiceTimeUntilDueParams `form:"invoice_time_until_due" json:"invoice_time_until_due,omitempty"`
	// The ID of the payment method configuration, which controls which payment methods are displayed to your customers.
	PaymentMethodConfiguration *string `form:"payment_method_configuration" json:"payment_method_configuration,omitempty"`
	// Payment Method specific configuration to be stored on the object.
	PaymentMethodOptions *V2BillingIntentCreateBillingGroupDataBillingSettingsPaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// Settings for calculating tax.
	Tax *V2BillingIntentCreateBillingGroupDataBillingSettingsTaxParams `form:"tax" json:"tax,omitempty"`
}

// The time at which the daily grouping interval begins.
type V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceDayTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Daily cadence details.
type V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceDayParams struct {
	// The time at which the daily grouping interval begins.
	Time *V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceDayTimeParams `form:"time" json:"time"`
}

// The time at which the monthly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceMonthTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Monthly cadence details.
type V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceMonthParams struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The month of year, from 1 through 12, from which occurrences are calculated. Defaults to the month when the billing group was created.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year,omitempty"`
	// The time at which the monthly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceMonthTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the weekly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceWeekTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Weekly cadence details.
type V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceWeekParams struct {
	// The ISO-8601 day of week, from 1 for Monday through 7 for Sunday.
	DayOfWeek *int64 `form:"day_of_week" json:"day_of_week"`
	// The time at which the weekly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceWeekTimeParams `form:"time" json:"time,omitempty"`
}

// The time at which the yearly grouping interval begins. Defaults to the time the billing group is created, in UTC.
type V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceYearTimeParams struct {
	// The hour, from 0 through 23.
	Hour *int64 `form:"hour" json:"hour"`
	// The minute, from 0 through 59.
	Minute *int64 `form:"minute" json:"minute"`
	// The second, from 0 through 59.
	Second *int64 `form:"second" json:"second"`
}

// Yearly cadence details.
type V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceYearParams struct {
	// The day of month, from 1 through 31. For shorter months, the cadence anchors to the month's last day.
	DayOfMonth *int64 `form:"day_of_month" json:"day_of_month"`
	// The month of year, from 1 through 12.
	MonthOfYear *int64 `form:"month_of_year" json:"month_of_year"`
	// The time at which the yearly grouping interval begins. Defaults to the time the billing group is created, in UTC.
	Time *V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceYearTimeParams `form:"time" json:"time,omitempty"`
}

// The cadence configuration when type is `by_cadence`.
type V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceParams struct {
	// Timestamp that anchors the recurring grouping cadence.
	BillingCycleAnchor *time.Time `form:"billing_cycle_anchor" json:"billing_cycle_anchor,omitempty"`
	// Daily cadence details.
	Day *V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceDayParams `form:"day" json:"day,omitempty"`
	// The number of interval units between groups. Defaults to 1.
	IntervalCount *int64 `form:"interval_count" json:"interval_count,omitempty"`
	// Monthly cadence details.
	Month *V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceMonthParams `form:"month" json:"month,omitempty"`
	// The interval unit for the cadence.
	Type *string `form:"type" json:"type"`
	// Weekly cadence details.
	Week *V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceWeekParams `form:"week" json:"week,omitempty"`
	// Yearly cadence details.
	Year *V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceYearParams `form:"year" json:"year,omitempty"`
}

// The strategy used to group member invoices.
type V2BillingIntentCreateBillingGroupDataGroupingBehaviorParams struct {
	// The cadence configuration when type is `by_cadence`.
	ByCadence *V2BillingIntentCreateBillingGroupDataGroupingBehaviorByCadenceParams `form:"by_cadence" json:"by_cadence,omitempty"`
	// The grouping strategy.
	Type *string `form:"type" json:"type"`
}

// Data for creating a new Billing Group.
type V2BillingIntentCreateBillingGroupDataParams struct {
	// Billing setting overrides for the group. If omitted, the applicable defaults are used.
	BillingSettings *V2BillingIntentCreateBillingGroupDataBillingSettingsParams `form:"billing_settings" json:"billing_settings,omitempty"`
	// The ID of the Customer object whose invoices are grouped.
	Customer *string `form:"customer" json:"customer"`
	// The ID of the PaymentMethod object used by default. If unset, the customer's default payment method is used.
	DefaultPaymentMethod *string `form:"default_payment_method" json:"default_payment_method,omitempty"`
	// The strategy used to group member invoices.
	GroupingBehavior *V2BillingIntentCreateBillingGroupDataGroupingBehaviorParams `form:"grouping_behavior" json:"grouping_behavior,omitempty"`
	// An internal key you can use to find a particular billing group. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
}

// Create a Billing Intent.
type V2BillingIntentCreateParams struct {
	Params `form:"*"`
	// Actions to be performed by this Billing Intent.
	Actions []*V2BillingIntentCreateActionParams `form:"actions" json:"actions"`
	// ID of an existing Billing Group to use.
	BillingGroup *string `form:"billing_group" json:"billing_group,omitempty"`
	// Data for creating a new Billing Group.
	BillingGroupData *V2BillingIntentCreateBillingGroupDataParams `form:"billing_group_data" json:"billing_group_data,omitempty"`
	// Three-letter ISO currency code, in lowercase. Must be a supported currency.
	Currency *string `form:"currency" json:"currency"`
}

// Retrieve a Billing Intent.
type V2BillingIntentRetrieveParams struct {
	Params `form:"*"`
}
