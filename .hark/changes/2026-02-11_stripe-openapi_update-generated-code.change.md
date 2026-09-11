---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2264
is_stripe_api_change: true
released_in_version: 84.4.0-alpha.3
---

* Add support for new resources `V2BillingCadenceSpendModifier`, `V2BillingOneTimeItem`, and `V2BillingRateCardCustomPricingUnitOverageRate`
* Add support for `Del`, `Get`, `List`, and `New` methods on resource `V2BillingRateCardCustomPricingUnitOverageRate`
* Add support for `Get`, `List`, `New`, and `Update` methods on resource `V2BillingOneTimeItem`
* Add support for `Get` method on resource `V2BillingCadenceSpendModifier`
* Add support for `SettlementType` on `ApplicationFee`
* Add support for `RateCardCustomPricingUnitOverageRateDetails` on `InvoiceItemPricing` and `InvoiceLineItemPricing`
* Add support for new value `rate_card_custom_pricing_unit_overage_rate_details` on enums `InvoiceItemPricing.Type` and `InvoiceLineItemPricing.Type`
* Add support for `DefaultSettings` on `InvoiceCreatePreviewScheduleDetailsParams`
* Add support for `PaymentBehavior` on `SubscriptionResumeParams`
* Add support for `EffectiveAt` and `SpendModifierRule` on `V2BillingIntentActionApplyParams`, `V2BillingIntentActionApply`, `V2BillingIntentActionRemoveParams`, and `V2BillingIntentActionRemove`
* Change type of `V2BillingIntentActionApply.Type`, `V2BillingIntentActionApplyParams.Type`, `V2BillingIntentActionRemove.Type`, and `V2BillingIntentActionRemoveParams.Type` from `literal('invoice_discount_rule')` to `enum('invoice_discount_rule'|'spend_modifier_rule')`
