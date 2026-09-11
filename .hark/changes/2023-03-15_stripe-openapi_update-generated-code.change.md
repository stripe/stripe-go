---
title: Update generated code (new)
pr_link: https://github.com/stripe/stripe-go/pull/1619
is_stripe_api_change: true
released_in_version: 74.12.0
---

* Add support for `CashappPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `Cashapp` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `MandatePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new value `cashapp` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Add support for `PreferredLocale` on `PaymentIntentConfirmPaymentMethodOptionsAffirmParams`, `PaymentIntentPaymentMethodOptionsAffirmParams`, and `PaymentIntentPaymentMethodOptionsAffirm`
* Add support for `CashappHandleRedirectOrDisplayQRCode` on `PaymentIntentNextAction` and `SetupIntentNextAction`
* Add support for new value `cashapp` on enum `PaymentLinkPaymentMethodTypes`
* Add support for new value `cashapp` on enum `PaymentMethodType`
