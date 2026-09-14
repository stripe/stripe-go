---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1806
is_stripe_api_change: true
released_in_version: 76.16.0-beta.1
---

* Add support for new resources `Entitlements.Event` and `Entitlements.Feature`
* Add support for `New` method on resource `Event`
* Add support for `List` and `New` methods on resource `Feature`
* Add support for `Swish` on `ConfirmationTokenPaymentMethodPreview`
* Add support for new value `swish` on enum `ConfirmationTokenPaymentMethodPreviewType`
* Add support for new value `customer.entitlement_summary.updated` on enum `EventType`
* Add support for `AccountTaxIDs` on `InvoiceCreatePreviewScheduleDetailsPhasesInvoiceSettingsParams`, `InvoiceUpcomingLinesScheduleDetailsPhasesInvoiceSettingsParams`, and `InvoiceUpcomingScheduleDetailsPhasesInvoiceSettingsParams`
* Add support for `Feature` on `ProductFeaturesParams` and `ProductFeatures`
