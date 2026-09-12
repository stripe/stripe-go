---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1945
is_stripe_api_change: true
released_in_version: 81.1.0-beta.1
---

* Add support for `TriggerAction` method on resource `PaymentIntent`
* Add support for `IDBankTransferPaymentsBca` and `IDBankTransferPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `BankBcaOnboarding` on `AccountSettingsParams` and `AccountSettings`
* Add support for `SendMoney` on `AccountSessionComponentsRecipientsFeaturesParams`
* Remove support for value `payout_statement_descriptor_profanity` from enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `IDBankTransfer` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `RefundDestinationDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentPaymentMethodDataParams`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for `Gopay`, `Qris`, and `Shopeepay` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
