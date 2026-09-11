---
title: Make paymentintent and paymentmethod codegen-able
pr_link: https://github.com/stripe/stripe-go/pull/1365
is_stripe_api_change: true
released_in_version: 72.73.0
---

* Fix `WechatPay` form name in `PaymentIntentPaymentMethodDataParams`
* Add support for `"challenge_only"` as `PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure` option
* Add support for `OffSessionOneOff` and `OffSessionRecurring` on `PaymentIntentConfirmParams`
* Add support for `BACSDebit`, `Bancontact`, `Giropay`, `InteracPresent`, `Metadata`, and `Sofort` on `PaymentIntentPaymentMethodDataParams`
* Add support for `CardPresent`, `Ideal`, `P24`, and `SepaDebit` on `PaymentIntentPaymentMethodOptionsParams` and `PaymentIntentPaymentMethodOptions`
* Add support for `ClientSecret`, `OffSessionOneOff`, and `OffSessionRecurring` on `PaymentIntentParams`
* Add support for `Object` on `PaymentIntent`
* Add support for `AmexExpressCheckout`, `ApplePay`, `GooglePay`, `Masterpass`, `SamsungPay`, and `VisaCheckout` on `PaymentMethodCardWallet`
