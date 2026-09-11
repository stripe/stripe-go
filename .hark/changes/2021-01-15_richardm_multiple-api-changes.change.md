---
title: Multiple API Changes
pr_link: https://github.com/stripe/stripe-go/pull/1241
is_stripe_api_change: true
released_in_version: 72.30.0
---

* Added support for `dynamic_tax_rates` on `CheckoutSessionParams.line_items`
* Added support for `customer_details` on `CheckoutSession`
* Added support for `type` on `IssuingTransactionListParams`
* Added support for `country` and `state` on `TaxRateParams` and `TaxRate`
