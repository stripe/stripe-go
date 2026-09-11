---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2213
is_stripe_api_change: true
released_in_version: 83.3.0-alpha.2
---

* Add support for new resource `IssuingProgram`
* Add support for `Get`, `List`, `New`, and `Update` methods on resource `IssuingProgram`
* Add support for `Schedule` on `Discount`
* Add support for `ApplicableFees` on `DelegatedCheckoutRequestedSessionTotalDetails`
* Add support for `ScheduleDetails` on `InvoiceItemParent`, `InvoiceLineItemParent`, `InvoiceParent`, and `QuotePreviewInvoiceParent`
* Add support for new value `schedule_details` on enum `InvoiceItemParent.Type`
* Add support for `BillingSchedules` on `InvoiceCreatePreviewScheduleDetailsParams`, `QuotePreviewSubscriptionSchedule`, `SubscriptionScheduleParams`, and `SubscriptionSchedule`
* Add support for new value `schedule_details` on enums `InvoiceParent.Type` and `QuotePreviewInvoiceParent.Type`
* Add support for new value `schedule_details` on enum `InvoiceLineItemParent.Type`
* Add support for `LatestInvoice` on `QuotePreviewSubscriptionSchedule` and `SubscriptionSchedule`
* Add support for `PhaseEffectiveAt` on `QuotePreviewSubscriptionScheduleDefaultSettings`, `SubscriptionScheduleDefaultSettingsParams`, and `SubscriptionScheduleDefaultSettings`
