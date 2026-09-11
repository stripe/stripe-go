---
title: Remove extraneous parameters from `CardUpdateParams` and `BankAccountUpdateParams`
pr_link: https://github.com/stripe/stripe-go/pull/2134
is_breaking: true
released_in_version: 83.0.0
---

- ⚠️ Removes `address_city`, `address_country`, `address_line1`, `address_line2`, `address_state`, `address_zip`, `exp_month`, `exp_year`, and `name` from `BankAccountUpdateParams`. These were not valid fields, so always received a 400 from the server if set.
- ⚠️ Removes `account_holder_name`, `account_holder_type`, `cvc`, `number`, and `owner` from `CardAccountParams`. These were not valid fields, so always received a 400 from the server if set.
