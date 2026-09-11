---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1900
is_breaking: true
is_stripe_api_change: true
released_in_version: 79.8.0-beta.1
---

* Add support for `CapitalFinancingApplication` and `CapitalFinancing` on `AccountSessionComponents`
* Add support for `Payto` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for new value `custom` on enum `CheckoutSessionUiMode`
* ⚠️  Remove support for `RiskCorrelationID` on `PaymentIntentConfirmPaymentMethodOptionsRechnungParams`, `PaymentIntentPaymentMethodOptionsRechnungParams`, and `PaymentIntentPaymentMethodOptionsRechnung`
* Add support for new value `payto` on enum `PaymentLinkPaymentMethodTypes`
