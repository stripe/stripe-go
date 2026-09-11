---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1301
is_stripe_api_change: true
released_in_version: 72.52.0
---

* Add support for `boleto` as a `PaymentMethodType`
* Add support for `Boleto` on `ChargePaymentMethodDetails`, `PaymentMethod`, `PaymentMethodParams`, `PaymentIntentPaymentMethodOptions`, `PaymentIntentPaymentMethodDataParams`, and `PaymentIntentPaymentMethodOptionsParams`
* Add support for `BoletoDisplayDetails` on `PaymentIntentNextAction`
* Add support for `il_vat` on enums `CheckoutSessionCustomerDetailsTaxIDsType` and `TaxIDType`
