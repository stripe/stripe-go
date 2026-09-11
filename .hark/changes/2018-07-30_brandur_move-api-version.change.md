---
title: Move to API version `2018-07-27` (breaking)
pr_link: https://github.com/stripe/stripe-go/pull/639
released_in_version: 37.0.0
---

* Remove `SKUs` from `Product`
* Subscription creation and update can no longer take a source
* Change `PercentOff` on coupon struct and params from integer to float
