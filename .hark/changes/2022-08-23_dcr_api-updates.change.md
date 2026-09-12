---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1532
is_stripe_api_change: true
released_in_version: 73.4.0
---

* Change type of `TreasuryOutboundTransferDestinationPaymentMethod` from `string` to `nullable(string)`
* Change return type of `FundCashBalance` method on `Customer` from `Customer` to `CustomerCashBalanceTransaction`
  * This is technically a breaking change, but this return type was actually incorrect and so the result of this method did not deserialize correctly.
* Change return type of `RetrieveFeatures` and `UpdateFeatures` methods on `TreasuryFinancialAccount` from `TreasuryFinancialAccount` to `TreasuryFinancialAccountFeatures`
  * This is technically a breaking change, but this return type was actually incorrect and so the result of this method did not deserialize correctly.
