---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1442
is_stripe_api_change: true
released_in_version: 72.99.0
---

* Add support for `BankTransferPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `CaptureBefore` on `ChargePaymentMethodDetailsCardPresent`
* Add support for `Address` and `Name` on `CheckoutSessionCustomerDetails`
* Add support for `CustomerBalance` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for new value `customer_balance` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Add support for `RequestExtendedAuthorization` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, and `PaymentIntentPaymentMethodOptionsCardPresent`
