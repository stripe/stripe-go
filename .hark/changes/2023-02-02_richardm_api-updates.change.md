---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1600
is_stripe_api_change: true
released_in_version: 74.7.0
---

* Add support for `Resume` method on resource `Subscription`
* Add support for `PaymentLink` on `CheckoutSessionListParams`
* Add support for `TrialSettings` on `CheckoutSessionSubscriptionDataParams`, `SubscriptionParams`, and `Subscription`
* Add support for new value `BE` on enums `CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferEuBankTransferCountry`, `InvoicePaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEuBankTransferCountry`, `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferEuBankTransferCountry`, and `SubscriptionPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferEuBankTransferCountry`
* Add support for `ShippingCost` on `CreditNoteParams`, `CreditNotePreviewLinesParams`, `CreditNotePreviewParams`, `CreditNote`, `InvoiceParams`, and `Invoice`
* Add support for `AmountShipping` on `CreditNote` and `Invoice`
* Add support for `ShippingDetails` on `InvoiceParams` and `Invoice`
* Add support for `SubscriptionResumeAt` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
* Add support for `InvoiceCreation` on `PaymentLinkParams` and `PaymentLink`
* Add support for new value `paused` on enum `SubscriptionStatus`
* Add support for new value `funding_reversed` on enum `CustomerCashBalanceTransactionType`
