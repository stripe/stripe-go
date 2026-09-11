---
title: Move to API version [`2019-08-14`](https://docs.stripe.com/changelog/2019-08-14) and other changes
pr_link: https://github.com/stripe/stripe-go/pull/915
released_in_version: 62.0.0
---

* Pin to API version `2019-08-14`
* Rename `AccountCapabilityPlatformPayments` to `AccountCapabilityTransfers`
* Add `Executive` in `PersonRelationship`
* Remove `PayentMethodOptions` as there was a typo which was fixed
* Make `OffSession` only support booleans on `PaymentIntent`
* Remove `PaymentIntentLastPaymentError` and use `Error` instead
* Move `DeclineCode` on `Error` to the `DeclineCode` type instead of `string`
