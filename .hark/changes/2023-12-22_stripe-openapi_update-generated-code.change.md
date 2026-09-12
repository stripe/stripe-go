---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1791
is_stripe_api_change: true
released_in_version: 76.11.0-beta.1
---

* Add support for `CapitalFinancingPromotion` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for new value `shipping_address_invalid` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Change type of `InvoiceIssuer` and `SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer` from `nullable(ConnectAccountReference)` to `ConnectAccountReference`
* Change type of `PaymentLinkSubscriptionDataInvoiceSettings` from `nullable(PaymentLinksResourceSubscriptionDataInvoiceSettings)` to `PaymentLinksResourceSubscriptionDataInvoiceSettings`
* Add support for `ShipFromDetails` on `TaxCalculationParams`, `TaxCalculation`, and `TaxTransaction`
