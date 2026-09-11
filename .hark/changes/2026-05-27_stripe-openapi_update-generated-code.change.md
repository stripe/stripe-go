---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2358
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.3.0-alpha.1
---

* Change type of `BillingAlertSpendThresholdParams.GroupBy` from `literal('pricing_plan_subscription')` to `enum('billing_cadence'|'pricing_plan_subscription')`
* ⚠️ Change type of `BillingAlertSpendThreshold.GroupBy` from `literal('pricing_plan_subscription')` to `enum('billing_cadence'|'pricing_plan_subscription')`
* Add support for new value `institution_requirement` on enum `FinancialConnectionsAccountStatusDetailsInactive.Cause`
* Add support for `WeChatPay` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for `GiftCard` on `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptionsParams`, and `PaymentIntentPaymentMethodOptions`
* Add support for `PaymentDetails` on `PaymentIntentPaymentsOrchestrationParams`
* Add support for `Enabled` on `PaymentIntentPaymentDetailsBenefitFrMealVoucher` and `SetupIntentSetupDetailsBenefitFrMealVoucher`
* ⚠️ Remove support for `LoginFailed`, `RegistrationFailed`, `RegistrationSuccess`, and `Type` on `RadarCustomerEvaluationParams`
* ⚠️ Remove support for `LatestVersion` on `V2BillingLicenseFee`, `V2BillingPricingPlan`, and `V2BillingRateCard`
* ⚠️ Remove support for `ServiceIntervalCount` and `ServiceInterval` on `V2BillingLicenseFee` and `V2BillingRateCard`
* Add support for `DebitAgreement` on `V2MoneyManagementReceivedCreditStripeBalancePayment`
* Add support for `CanonicalPath` on `EventsV2CoreHealthTrafficVolumeDropFiringEventImpact` and `EventsV2CoreHealthTrafficVolumeDropResolvedEventImpact`
* Add support for snapshot event `EventTypePaymentIntentExpired` with resource `PaymentIntent`
* Add support for event notifications `V2CoreHealthElementsErrorFiringEvent`, `V2CoreHealthElementsErrorResolvedEvent`, `V2CoreHealthInvoiceCountDroppedFiringEvent`, and `V2CoreHealthInvoiceCountDroppedResolvedEvent`
