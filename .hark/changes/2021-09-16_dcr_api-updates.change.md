---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1342
is_stripe_api_change: true
released_in_version: 72.65.0
---

* Add support for `Livemode` on `ReportingReportType`.
* Add support for `DefaultFor` on `CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptionsParams`, `CheckoutSessionPaymentMethodOptionsACSSDebitMandateOptions`, `MandatePaymentMethodDetailsACSSDebit`, `SetupIntentPaymentMethodOptionsACSSDebitMandateOptionsParams`, and `SetupIntentPaymentMethodOptionsACSSDebitMandateOptions`.
* Add support for `ACSSDebit` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`.
* Add support for new value `acss_debit` on enums `InvoicePaymentSettingsPaymentMethodType` and `SubscriptionPaymentSettingsPaymentMethodType`.
* Add support for `FullNameAliases` on `PersonParams` and `Person`.
