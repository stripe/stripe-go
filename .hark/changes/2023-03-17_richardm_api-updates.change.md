---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1622
is_stripe_api_change: true
released_in_version: 74.12.0
---

* Add support for `CashAppPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `FutureRequirements` and `Requirements` on `BankAccount`
* Add support for `CashApp` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `MandatePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Country` on `ChargePaymentMethodDetailsLink`
* Add support for new value `cashapp` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Add support for `PreferredLocale` on `PaymentIntentConfirmPaymentMethodOptionsAffirmParams`, `PaymentIntentPaymentMethodOptionsAffirmParams`, and `PaymentIntentPaymentMethodOptionsAffirm`
* Add support for new value `automatic_async` on enums `PaymentIntentCaptureMethod` and `PaymentLinkPaymentIntentDataCaptureMethod`
* Add support for `CashAppHandleRedirectOrDisplayQRCode` on `PaymentIntentNextAction` and `SetupIntentNextAction`
* Add support for new value `cashapp` on enum `PaymentLinkPaymentMethodTypes`
* Add support for new value `cashapp` on enum `PaymentMethodType`
