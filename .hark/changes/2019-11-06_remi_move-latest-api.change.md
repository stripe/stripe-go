---
title: Move to the latest API version and add new changes
pr_link: https://github.com/stripe/stripe-go/pull/987
released_in_version: 67.0.0
---

* Move to API version `2019-11-05`
* Add `DefaultSettings` on `SubscritionSchedule`
* Remove `BillingThresholds`, `CollectionMethod`, `DefaultPaymentMethod` and `DefaultSource` and `invoice_settings` from `SubscriptionSchedule`
* `OffSession` on `PaymentIntent` is now always a boolean
