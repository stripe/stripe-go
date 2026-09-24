---
title: Remove multiple deprecated APIs
pr_url: https://github.com/stripe/stripe-go/pull/1171
is_stripe_api_change: true
released_in_version: 72.0.0
---

* Remove support for the `Recipient` API
* Remove support for the `RecipientTransfer` API
* Remove support for the `BitcoinReceiver` API
* Remove support for the `ThreeDSecure` API which has been replaced by PaymentIntent and PaymentMethod
* Remove support for the `ExchangeRate` API which has never shipped publicly and is being reworked
