---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1431
is_stripe_api_change: true
released_in_version: 72.93.0
---

* Add support for `Mandate` on `ChargePaymentMethodDetailsCard`
* Add support for `MandateOptions` on `SetupIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCard`, SetupIntentConfirmPaymentMethodOptionsCardParams`, and `SetupIntentPaymentMethodOptionsCard`
* Add support for `CardAwaitNotification` on `PaymentIntentNextAction`
* Add support for `CustomerNotification` on `PaymentIntentProcessingCard`
