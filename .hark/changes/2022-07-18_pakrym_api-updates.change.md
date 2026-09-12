---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1497
is_stripe_api_change: true
released_in_version: 72.120.0
---

* Add support for `BLIKPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `BLIK` on `ChargePaymentMethodDetails`, `MandatePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodDataParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Change type of `CheckoutSessionConsentCollectionPromotionsParams`, `CheckoutSessionConsentCollectionPromotions`, `PaymentLinkConsentCollectionPromotionsParams`, and `PaymentLinkConsentCollectionPromotions` from `literal('auto')` to `enum('auto'|'none')`
* Add support for new value `blik` on enum `PaymentLinkPaymentMethodTypes`
* Add support for new value `blik` on enum `PaymentMethodType`
