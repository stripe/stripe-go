---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1422
is_stripe_api_change: true
released_in_version: 72.89.0
---

* Add support for `KonbiniPayments` on `AccountCapabilitiesParams`, and `AccountCapabilities`
`BillingPortalConfigurationBusinessProfileTermsOfServiceUrl` from `string` to `nullable(string)`
* Add support for `Konbini` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`,  `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for new value `konbini` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Add support for `KonbiniDisplayDetails` on `PaymentIntentNextAction`
* Add support for new value `konbini` on enum `PaymentMethodType`
