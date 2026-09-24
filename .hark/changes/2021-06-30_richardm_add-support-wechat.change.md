---
title: Add support for Wechat Pay
pr_url: https://github.com/stripe/stripe-go/pull/1304
is_stripe_api_change: true
released_in_version: 72.54.0
---

* Add support for `WechatPay` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, and `PaymentMethod`
* Add support for new value `wechat_pay` on enums `InvoicePaymentSettingsPaymentMethodType` and `PaymentMethodType`
* Add support for `WechatPayDisplayQRCode`, `WechatPayRedirectToAndroidApp`, and `WechatPayRedirectToIOSApp` on `PaymentIntentNextAction`
