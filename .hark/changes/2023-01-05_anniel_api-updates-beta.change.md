---
title: API Updates for beta branch
pr_url: https://github.com/stripe/stripe-go/pull/1589
is_stripe_api_change: true
released_in_version: 74.6.0-beta.1
---

* Updated stable APIs to the latest version
* Add support for `MarkStaleQuote` method on resource `Quote`
* Add support for `Duration` and `LineEndsAt` on `QuoteSubscriptionDataBillOnAcceptanceBillUntilParams` and `QuoteSubscriptionDataOverridesBillOnAcceptanceBillUntilParams`
* Remove support for `LineStartsAt` on `QuoteSubscriptionDataBillOnAcceptanceBillUntilParams` and `QuoteSubscriptionDataOverridesBillOnAcceptanceBillUntilParams`
* Add support for `Metadata` on `TerminalReaderActionRefundPayment` and `TerminalReaderRefundPaymentParams`
