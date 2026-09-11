---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1455
is_stripe_api_change: true
released_in_version: 72.105.0
---

* Add support for new resources `FinancialConnections.AccountOwner`, `FinancialConnections.AccountOwnership`, `FinancialConnections.Account`, and `FinancialConnections.Session`
* Add support for `FinancialConnections` on `CheckoutSessionPaymentMethodOptionsUsBankAccountParams`, `CheckoutSessionPaymentMethodOptionsUsBankAccount`, `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountParams`, `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccount`, `PaymentIntentConfirmPaymentMethodOptionsUsBankAccountParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountParams`, `PaymentIntentPaymentMethodOptionsUsBankAccount`, `SetupIntentConfirmPaymentMethodOptionsUsBankAccountParams`, `SetupIntentPaymentMethodOptionsUsBankAccountParams`, `SetupIntentPaymentMethodOptionsUsBankAccount`, `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccount`
* Add support for `FinancialConnectionsAccount` on `PaymentIntentConfirmPaymentMethodDataUsBankAccountParams`, `PaymentIntentPaymentMethodDataUsBankAccountParams`, `PaymentMethodUsBankAccountParams`, `PaymentMethodUsBankAccount`, `SetupIntentConfirmPaymentMethodDataUsBankAccountParams`, and `SetupIntentPaymentMethodDataUsBankAccountParams`
