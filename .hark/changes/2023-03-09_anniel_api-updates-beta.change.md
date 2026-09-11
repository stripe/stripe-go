---
title: API Updates for beta branch
pr_link: https://github.com/stripe/stripe-go/pull/1617
is_stripe_api_change: true
released_in_version: 74.12.0-beta.1
---

* Updated stable APIs to the latest version
* Remove support for `ListTransactions` method on resource `Tax.Transaction`
* Add support for `UpdateBehavior` on `SubscriptionPrebillingParams`, `SubscriptionPrebilling`, `SubscriptionSchedulePrebillingParams`, and `SubscriptionSchedulePrebilling`
* Add support for `Prebilling` on `SubscriptionScheduleAmendParams`
* Change type of `SubscriptionScheduleAppliesTo` from `nullable(QuotesResourceQuoteLinesAppliesTo)` to `QuotesResourceQuoteLinesAppliesTo`
* Add support for `TaxabilityOverride` on `TaxCalculationCustomerDetailsParams`, `TaxCalculationCustomerDetails`, and `TaxTransactionCustomerDetails`
* Add support for `TaxSummary` on `TaxCalculation`
* Remove support for `TaxBreakdown` on `TaxCalculation`
* Add support for `TaxBehavior` on `TaxSettingsDefaultsParams` and `TaxSettingsDefaults`
