---
title: Multiple API changes
pr_link: https://github.com/stripe/stripe-go/pull/1099
is_stripe_api_change: true
released_in_version: 71.12.0
---

* Add `issuing_dispute` as a `type` on `BalanceTransaction`
* Add `BalanceTransactions` as a a list of `BalanceTransaction` on Issuing `Dispute`
* Add `Fingerprint` and `TransactionId` in `ChargePaymentMethodDetailsAlipay` on `Charge`
* Add `Amount` in `InvoiceTransferData` and `InvoiceTransferDataParams` on `Invoice`
* Add `AmountPercent` in `SubscriptionTransferData` and `SubscriptionTransferDataParams` on `Subscription`
