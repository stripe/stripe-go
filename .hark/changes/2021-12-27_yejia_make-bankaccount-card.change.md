---
title: Make bankaccount and card codegen-able
pr_url: https://github.com/stripe/stripe-go/pull/1396
is_stripe_api_change: true
released_in_version: 72.82.0
---

* Add support for `address_city`, `address_country`, `address_line1`, `address_line2`, `address_state`, `address_zip`, `exp_month`, `exp_year`, and `name` on `BankAccountParams`
* Add support for `account_holder_name`, `account_holder_type`, and `owner` on `CardParams`
* Add support for `account` on `Card`
