---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1397
is_stripe_api_change: true
released_in_version: 72.81.0
---

* Add support for `AUBECSDebit` on `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, and `PaymentIntentPaymentMethodOptions`
* Change type of `PaymentIntentProcessingType` from `string` to `literal('card')`. This is not considered a breaking change as the field was added in the same release.
