---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1662
is_stripe_api_change: true
released_in_version: 74.20.0
---

* Add support for `ZipPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `Zip` on `ChargePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new value `zip` on enum `PaymentMethodType`
