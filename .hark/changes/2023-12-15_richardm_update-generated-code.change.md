---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1785
is_stripe_api_change: true
released_in_version: 76.10.0-beta.1
---

* Add support for `PreviewMode` and `SubscriptionDetails` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
* Remove support for `SubscriptionTrialFromPlan` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
* Add support for `BillingBehavior`, `EndBehavior`, and `ProrationBehavior` on `InvoiceUpcomingLinesScheduleDetailsParams` and `InvoiceUpcomingScheduleDetailsParams`
