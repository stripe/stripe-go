---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2354
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.2.0-alpha.6
---

* Add support for new resource `PaymentLocationCapability`
* Add support for `Get`, `List`, and `Update` methods on resource `PaymentLocationCapability`
* Add support for `Close` and `SimulateNetworkLifecycleDisputeResponse` test helper methods on resource `IssuingDispute`
* Change type of `DelegatedCheckoutRequestedSessionDiscountsParams.Codes` from `array(string)` to `emptyable(array(string))`
* ⚠️ Remove support for `CreditedItems` on `InvoiceItemProrationDetails`
* Add support for `BalanceResponse` on `IssuingAuthorization`
* Add support for `PaymentEvaluations` on `PaymentAttemptRecordReportCanceledParams`, `PaymentAttemptRecordReportFailedParams`, `PaymentRecordReportPaymentAttemptCanceledParams`, `PaymentRecordReportPaymentAttemptFailedParams`, and `PaymentRecordReportPaymentFailedParams`
* Add support for `Enabled` on `PaymentIntentConfirmPaymentDetailsBenefitFrMealVoucherParams`, `PaymentIntentPaymentDetailsBenefitFrMealVoucherParams`, `SetupIntentConfirmSetupDetailsBenefitFrMealVoucherParams`, and `SetupIntentSetupDetailsBenefitFrMealVoucherParams`
* Add support for `AdvancedFeatureDetails` and `AllowedPaymentMethodTypes` on `PaymentIntent`
* Change type of `PaymentLocationAddressParams.City` from `string` to `emptyable(string)`
* Change type of `PaymentLocationAddressParams.Line1` from `string` to `emptyable(string)`
* Change type of `PaymentLocationAddressParams.Line2` from `string` to `emptyable(string)`
* Change type of `PaymentLocationAddressParams.PostalCode` from `string` to `emptyable(string)`
* Change type of `PaymentLocationAddressParams.State` from `string` to `emptyable(string)`
* ⚠️ Remove support for `PaymentBehavior` on `SubscriptionResumeParams`
* ⚠️ Remove support for `StatusDetails` on `Subscription`
