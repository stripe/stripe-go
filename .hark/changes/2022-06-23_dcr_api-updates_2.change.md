---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1481
is_stripe_api_change: true
released_in_version: 72.116.0
---

* Add support for `PromptPayPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `PromptPay` on `ChargePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `SubtotalExcludingTax` on `CreditNote` and `Invoice`
* Add support for `AmountExcludingTax` and `UnitAmountExcludingTax` on `CreditNoteLineItem` and `InvoiceLineItem`
* Add support for `RenderingOptions` on `InvoiceParams`
* Add support for `TotalExcludingTax` on `Invoice`
* Add support for new value `promptpay` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Add support for `AutomaticPaymentMethods` on `OrderPaymentSettings`
* Add support for `PromptPayDisplayQRCode` on `PaymentIntentNextAction`
* Add support for new value `promptpay` on enum `PaymentMethodType`
