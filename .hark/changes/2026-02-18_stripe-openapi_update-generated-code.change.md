---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2268
is_breaking: true
is_stripe_api_change: true
released_in_version: 84.4.0-alpha.4
---

* Add support for `SpendThreshold` on `BillingAlertParams` and `BillingAlert`
* ⚠️ Add support for new value `spend_threshold` on enum `BillingAlert.AlertType`
* Add support for `InvoiceItem`, `ProrationDetails`, `Proration`, and `Subscription` on `InvoiceLineItemParentScheduleDetails`
* Add support for `Custom` on `PaymentMethodParams`
* Add support for `PaymentMethodReference` and `Usage` on `PaymentMethodCustom`
* ⚠️ Change type of `QuoteSubscriptionDataOverridesParams.BillingSchedules` from `emptyable(array(billing_schedules_update_specs))` to `array(billing_schedules_update_specs)`
* Add support for `OutstandingUsageThrough` and `UnusedTimeFrom` on `SubscriptionPauseBillForParams`
* ⚠️ Remove support for `OutstandingUsage` and `UnusedTime` on `SubscriptionPauseBillForParams`
* ⚠️ Remove support for `PaymentBehavior` on `SubscriptionResumeParams`
