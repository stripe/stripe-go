---
title: "[#1424](https://github.com/stripe/stripe-go/pull/1424) API Updates"
pr_url: https://github.com/stripe/stripe-go/pull/1423
is_stripe_api_change: true
released_in_version: 72.90.0
---

* Add support for new resource `TestHelpers.TestClock`
* Add support for `TestClock` on `CustomerParams`, `Customer`, `Invoice`, `InvoiceItem`, `QuoteParams`, `Quote`, `Subscription`, and `SubscriptionSchedule`
* Add support for `PendingInvoiceItemsBehavior` on `InvoiceParams`
* Change type of `ProductUrlParams` from `string` to `emptyStringable(string)`
* Add support for `NextAction` on `Refund`
