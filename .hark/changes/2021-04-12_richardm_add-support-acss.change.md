---
title: Add support for ACSS debit payment method
pr_link: https://github.com/stripe/stripe-go/pull/1275
is_stripe_api_change: true
released_in_version: 72.42.0
---

* Add support for `acss_debit` as value for `PaymentMethodType`.
* Add support for `ACSSDebit` on `PaymentMethod`, `PaymentMethodParams`, `PaymentIntentPaymentMethodOptions`,  `PaymentIntentPaymentMethodOptionsParams`, `MandatePaymentMethodDetails`, `SetupIntentPaymentMethodOptions`, and `SetupIntentPaymentOptionsParams`.
* Add support for `ACSSDebitPayments` on `AccountCapabilities`
* Add support for `PaymentMethodOptions` on `CheckoutSession`
* Add support for `verify_with_microdeposits` and `use_stripe_sdk` on `PaymentIntentNextAction` and `SetupIntentNextAction`
