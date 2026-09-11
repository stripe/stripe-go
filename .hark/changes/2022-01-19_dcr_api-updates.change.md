---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1407
is_stripe_api_change: true
released_in_version: 72.84.0
---

* Change type of `ChargeStatus` from `string` to `enum('failed'|'pending'|'succeeded')`
* Add support for `BACSDebit` and `EPS` on `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, and `PaymentIntentPaymentMethodOptions`
* Add support for `ImageURLPNG` and `ImageURLSVG` on `PaymentIntentNextActionWechatPayDisplayQRCode`
