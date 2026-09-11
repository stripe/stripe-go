---
title: A few Billing changes
pr_link: https://github.com/stripe/stripe-go/pull/922
released_in_version: 62.2.0
---

* Add `Schedule` to `Subscription`
* Add missing parameters for the Upcoming Invoice API: `Schedule`, `SubscriptionCancelAt`, `SubscriptionCancelNow`
* Add missing properties and parameters for a `SubscriptionSchedule` phase: `BillingThresholds`, `CollectionMethod`, `DefaultPaymentMethod`, `InvoiceSettings`
