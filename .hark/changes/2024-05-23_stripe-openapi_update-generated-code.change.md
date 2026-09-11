---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1864
is_stripe_api_change: true
released_in_version: 78.8.0
---

* Add support for `ExternalAccountCollection` on `AccountSessionComponentsBalancesFeaturesParams`, `AccountSessionComponentsBalancesFeatures`, `AccountSessionComponentsPayoutsFeaturesParams`, and `AccountSessionComponentsPayoutsFeatures`
* Add support for new value `terminal_reader_invalid_location_for_payment` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `PaymentMethodRemove` on `CheckoutSessionSavedPaymentMethodOptions`
