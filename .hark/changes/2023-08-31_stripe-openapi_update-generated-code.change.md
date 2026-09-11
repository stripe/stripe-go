---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1722
is_stripe_api_change: true
released_in_version: 75.3.0
---

* Add support for new resource `AccountSession`
* Add support for `New` method on resource `AccountSession`
* Add support for new values `obligation_inbound`, `obligation_outbound`, `obligation_payout_failure`, `obligation_payout`, `obligation_reversal_inbound`, and `obligation_reversal_outbound` on enum `BalanceTransactionType`
* Change type of `EventType` from `string` to `enum`
* Add support for `Application` on `PaymentLink`
