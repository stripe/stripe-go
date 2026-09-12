---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1886
is_breaking: true
is_stripe_api_change: true
released_in_version: 79.3.0
---

* ⚠️ Remove support for values `billing_policy_remote_function_response_invalid`, `billing_policy_remote_function_timeout`, `billing_policy_remote_function_unexpected_status_code`, and `billing_policy_remote_function_unreachable` from enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`.
* ⚠️ Remove support for value `payment_intent_fx_quote_invalid` from enum `StripeErrorCode`. The was mistakenly released last week.
* Add support for `PaymentMethodOptions` on `ConfirmationToken`
* Add support for `PaymentElement` on `CustomerSessionComponentsParams` and `CustomerSessionComponents`
* Add support for `AddressValidation` on `IssuingCardShippingParams` and `IssuingCardShipping`
* Add support for `Shipping` on `IssuingCardParams`
