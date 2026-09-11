---
title: "Multiple breaking changes:"
pr_link: https://github.com/stripe/stripe-go/pull/1000
released_in_version: 68.0.0
---

* Pin to API version `2019-12-03`
* Rename `InvoiceBillingStatus` to `InvoiceStatus` for consistency
* Remove typo-ed field `OutOfBankdAmount` on `CreditNote`
* Remove deprecated `PaymentIntentPaymentMethodOptionsCardRequestThreeDSecureChallengeOnly` and `SetupIntentPaymentMethodOptionsCardRequestThreeDSecureChallengeOnly` from `PaymentIntent` and `SetupIntent`.
* Remove `OperatorAccount` on `TerminalLocationListParams`
