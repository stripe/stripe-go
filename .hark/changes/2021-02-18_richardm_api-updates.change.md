---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1252
is_stripe_api_change: true
released_in_version: 72.34.0
---

* Add support for `afterpay_clearpay` on `PaymentMethod`, `PaymentMethodParams`, `PaymentIntentPaymentMethodDataParams`, and `ChargePaymentMethodDetails`
* Add `afterpay_clearpay` as an enum member on `PaymentMethodType`
* Add support for `adjustable_quantity` on `CheckoutSessionLineItemParams`
* Add support for `on_behalf_of` on `InvoiceParams` and `Invoice`
