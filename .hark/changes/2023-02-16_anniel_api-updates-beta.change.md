---
title: API Updates for beta branch
pr_link: https://github.com/stripe/stripe-go/pull/1608
is_stripe_api_change: true
released_in_version: 74.9.0-beta.1
---

* Updated stable APIs to the latest version
* Add support for `CurrencyConversion` on `CheckoutSession`
* Add support for `Limits` on `FinancialConnectionsSessionParams` and `FinancialConnectionsSession`
* Remove support for `Enabled` on `FinancialConnectionsSessionManualEntryParams`
* Change type of `QuoteStatusDetailsCanceled` from `nullable(QuotesResourceStatusDetailsCanceledStatusDetails)` to `QuotesResourceStatusDetailsCanceledStatusDetails`
* Change type of `QuoteStatusDetailsStale` from `nullable(QuotesResourceStatusDetailsStaleStatusDetails)` to `QuotesResourceStatusDetailsStaleStatusDetails`
* Remove support for `Reference` on `TaxCalculationParams` and `TaxCalculation`
* Add support for `Reference` on `TaxTransactionParams`
