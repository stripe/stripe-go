---
title: Make paymentsource.go and client codegen-able
pr_url: https://github.com/stripe/stripe-go/pull/1400
is_stripe_api_change: true
released_in_version: 72.82.0
---

* Add support for `account_holder_name`, `account_holder_type`, `address_city`, `address_country`, `address_line1`, `address_line2`, `address_state`, `address_zip`, `exp_month`, `exp_year`, `name`, `owner` on `CustomerSourceParams`
* Add support for `PaymentSourceOwnerParams`
* Add support for `Object` on `SourceListParams`
