---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1854
is_stripe_api_change: true
released_in_version: 78.6.0-beta.1
---

* Add support for `RechnungPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `Rechnung` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Multibanco` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for new value `rechnung` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
