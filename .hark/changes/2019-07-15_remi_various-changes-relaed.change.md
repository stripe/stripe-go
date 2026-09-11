---
title: Various changes relaed to SCA for Billing
pr_link: https://github.com/stripe/stripe-go/pull/891
is_stripe_api_change: true
released_in_version: 61.14.0
---

* Add support for `PendingSetupIntent` on `Subscription`
* Add support for `PaymentBehavior` on `Subscription` creation and update
* Add support for `PaymentBehavior` on `SubscriptionItem` update
* Add support for `OffSession` when paying an `Invoice`
* Add support for `OffSession` on `Subscription` creation and update
