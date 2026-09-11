---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2296
is_breaking: true
is_stripe_api_change: true
released_in_version: 84.5.0-alpha.4
---

* Add support for new resources `OrchestrationPaymentAttempt` and `RadarCustomerEvaluation`
* Add support for `Get` method on resource `OrchestrationPaymentAttempt`
* Add support for `New` and `Update` methods on resource `RadarCustomerEvaluation`
* Add support for `Approve` method on resource `CheckoutSession`
* Add support for `ReportAuthenticated`, `ReportCanceled`, `ReportFailed`, `ReportGuaranteed`, `ReportInformational`, and `ReportRefund` methods on resource `PaymentAttemptRecord`
* Add support for `SimulateCryptoDeposit` test helper method on resource `PaymentIntent`
* Add support for `CreateUSPaperCheckOnApplication` on `AccountSessionComponentsCheckScanningFeaturesParams`
* Add support for `ApprovalMethod` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `CurrentAttempt` on `CheckoutSession`
* Add support for `SelectedFulfillmentOptionOverrides` on `DelegatedCheckoutRequestedSessionFulfillmentDetailsParams`
* Add support for `PricingPlanSubscriptionDetails` on `InvoiceItemParent` and `InvoiceLineItemParent`
* ⚠️ Remove support for `LicenseFeeSubscriptionDetails` on `InvoiceItemParent` and `InvoiceLineItemParent`
* ⚠️ Remove support for `PricingPlanSubscription` and `PricingPlanVersion` on `InvoiceItemParentRateCardSubscriptionDetails` and `InvoiceLineItemParentRateCardSubscriptionDetails`
* Add support for new value `pricing_plan_subscription_details` on enum `InvoiceItemParent.Type`
* ⚠️ Remove support for value `license_fee_subscription_details` from enum `InvoiceItemParent.Type`
* Add support for new value `discounts` on enum `InvoiceItem.FrozenFields`
* Add support for new value `pricing_plan_subscription_details` on enum `InvoiceLineItemParent.Type`
* ⚠️ Remove support for value `license_fee_subscription_details` from enum `InvoiceLineItemParent.Type`
* Add support for `TokenDetails` on `IssuingAuthorization`
* Add support for `DepositOptions` and `Mode` on `PaymentIntentConfirmPaymentMethodOptionsCryptoParams`, `PaymentIntentPaymentMethodOptionsCryptoParams`, and `PaymentIntentPaymentMethodOptionsCrypto`
* Add support for `CryptoDisplayDetails` on `PaymentIntentNextAction`
* Add support for `FailureCode` on `PaymentRecordReportPaymentAttemptFailedParams` and `PaymentRecordReportPaymentFailedParams`
* Add support for `RecurringInterval` on `SharedPaymentGrantedTokenUsageLimitsParams`
* Add support for `HomeRuleTax` on `TaxRegistrationCountryOptionsUsParams` and `TaxRegistrationCountryOptionsUs`
* Add support for new value `home_rule_tax` on enum `TaxRegistrationCountryOptionsUs.Type`
