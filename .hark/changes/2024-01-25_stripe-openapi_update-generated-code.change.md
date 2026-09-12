---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1803
is_stripe_api_change: true
released_in_version: 76.14.0
---

* Add support for `AnnualRevenue` and `EstimatedWorkerCount` on `AccountBusinessProfileParams` and `AccountBusinessProfile`
* Add support for new value `registered_charity` on enum `AccountCompanyStructure`
* Add support for `CollectionOptions` on `AccountLinkParams`
* Add support for `Liability` on `CheckoutSessionAutomaticTaxParams`, `CheckoutSessionAutomaticTax`, `PaymentLinkAutomaticTaxParams`, `PaymentLinkAutomaticTax`, `QuoteAutomaticTaxParams`, `QuoteAutomaticTax`, `SubscriptionScheduleDefaultSettingsAutomaticTaxParams`, `SubscriptionScheduleDefaultSettingsAutomaticTax`, `SubscriptionSchedulePhasesAutomaticTaxParams`, and `SubscriptionSchedulePhasesAutomaticTax`
* Add support for `Issuer` on `CheckoutSessionInvoiceCreationInvoiceDataParams`, `CheckoutSessionInvoiceCreationInvoiceData`, `PaymentLinkInvoiceCreationInvoiceDataParams`, `PaymentLinkInvoiceCreationInvoiceData`, `QuoteInvoiceSettingsParams`, `QuoteInvoiceSettings`, `SubscriptionScheduleDefaultSettingsInvoiceSettingsParams`, `SubscriptionScheduleDefaultSettingsInvoiceSettings`, `SubscriptionSchedulePhasesInvoiceSettingsParams`, and `SubscriptionSchedulePhasesInvoiceSettings`
* Add support for `InvoiceSettings` on `CheckoutSessionSubscriptionDataParams`, `PaymentLinkSubscriptionDataParams`, and `PaymentLinkSubscriptionData`
* Add support for `PromotionCode` on `InvoiceUpcomingDiscountsParams`, `InvoiceUpcomingInvoiceItemsDiscountsParams`, `InvoiceUpcomingLinesDiscountsParams`, and `InvoiceUpcomingLinesInvoiceItemsDiscountsParams`
* Add support for new value `challenge` on enums `InvoicePaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure` and `SubscriptionPaymentSettingsPaymentMethodOptionsCardRequestThreeDSecure`
* Add support for `AccountType` on `PaymentMethodUsBankAccountParams`
