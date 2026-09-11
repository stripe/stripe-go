---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1487
is_stripe_api_change: true
released_in_version: 72.117.0
---

* Add support for `DeliverCard`, `FailCard`, `ReturnCard`, and `ShipCard` test helper methods on resource `Issuing.Card`
* Change type of `PaymentLinkPaymentMethodTypesParams` and `PaymentLinkPaymentMethodTypes` from `literal('card')` to `enum`
* Add support for `HostedRegulatoryReceiptURL` on `TreasuryReceivedCredit` and `TreasuryReceivedDebit`
