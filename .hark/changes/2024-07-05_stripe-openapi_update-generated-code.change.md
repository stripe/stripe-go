---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1881
is_stripe_api_change: true
released_in_version: 79.2.0
---

* Add support for `AddLines`, `RemoveLines`, and `UpdateLines` methods on resource `Invoice`
* Add support for new value `payment_intent_fx_quote_invalid` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `PostedAt` on `TaxTransactionCreateFromCalculationParams` and `TaxTransaction`
