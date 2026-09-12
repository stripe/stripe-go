---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1737
is_stripe_api_change: true
released_in_version: 75.8.0-beta.1
---

* Remove support for `Customer` on `ConfirmationToken`
* Add support for `Issuer` on `InvoiceParams`, `InvoiceUpcomingLinesParams`, `InvoiceUpcomingParams`, `Invoice`, `QuoteInvoiceSettingsParams`, `QuoteInvoiceSettings`, `SubscriptionScheduleDefaultSettingsInvoiceSettingsParams`, `SubscriptionScheduleDefaultSettingsInvoiceSettings`, `SubscriptionSchedulePhasesInvoiceSettingsParams`, and `SubscriptionSchedulePhasesInvoiceSettings`
* Add support for `OnBehalfOf` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
* Add support for `Liability` on `InvoiceAutomaticTaxParams`, `InvoiceAutomaticTax`, `InvoiceUpcomingAutomaticTaxParams`, `InvoiceUpcomingLinesAutomaticTaxParams`, `QuoteAutomaticTaxParams`, `QuoteAutomaticTax`, `SubscriptionAutomaticTaxParams`, `SubscriptionAutomaticTax`, `SubscriptionScheduleDefaultSettingsAutomaticTaxParams`, `SubscriptionScheduleDefaultSettingsAutomaticTax`, `SubscriptionSchedulePhasesAutomaticTaxParams`, and `SubscriptionSchedulePhasesAutomaticTax`
* Add support for `InvoiceSettings` on `SubscriptionParams`
