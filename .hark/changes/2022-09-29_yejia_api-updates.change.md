---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1549
is_stripe_api_change: true
released_in_version: 73.11.0
---

* Change type of `ChargePaymentMethodDetailsCardPresentIncrementalAuthorizationSupported` and `ChargePaymentMethodDetailsCardPresentOvercaptureSupported` from `nullable(boolean)` to `boolean`
* Add support for `Created` on `CheckoutSession`
* Add support for `SetupFutureUsage` on `PaymentIntentConfirmPaymentMethodOptionsPixParams`, `PaymentIntentPaymentMethodOptionsPixParams`, and `PaymentIntentPaymentMethodOptionsPix`
* Deprecate `CheckoutSessionSubscriptionDataTransferDataParams.items` and `CheckoutSessionSubscriptionDataItemParams` (use the `line_items` param instead). This will be removed in the next major version.
