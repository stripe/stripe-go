---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1745
is_stripe_api_change: true
released_in_version: 75.11.0-beta.1
---

* Add support for new resources `AccountNotice` and `Issuing.CreditUnderwritingRecord`
* Add support for `Get`, `List`, and `Update` methods on resource `AccountNotice`
* Add support for `Correct`, `CreateFromApplication`, `CreateFromProactiveReview`, `Get`, `List`, and `ReportDecision` methods on resource `CreditUnderwritingRecord`
* Change type of `CheckoutSessionAutomaticTaxLiabilityAccount`, `CheckoutSessionInvoiceCreationInvoiceDataIssuerAccount`, `InvoiceAutomaticTaxLiabilityAccount`, `InvoiceIssuerAccount`, `QuoteAutomaticTaxLiabilityAccount`, `QuoteInvoiceSettingsIssuerAccount`, `SubscriptionAutomaticTaxLiabilityAccount`, `SubscriptionScheduleDefaultSettingsAutomaticTaxLiabilityAccount`, `SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuerAccount`, `SubscriptionSchedulePhasesAutomaticTaxLiabilityAccount`, and `SubscriptionSchedulePhasesInvoiceSettingsIssuerAccount` from `nullable(expandable($Account))` to `expandable($Account)`
* Add support for new values `account_notice.created` and `account_notice.updated` on enum `EventType`
* Add support for new values `local_amusement_tax` and `state_communications_tax` on enum `TaxRegistrationCountryOptionsUsType`
