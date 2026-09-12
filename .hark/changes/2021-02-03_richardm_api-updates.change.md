---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1245
is_stripe_api_change: true
released_in_version: 72.32.0
---

* Add `nationality` to `Person` and `PersonParams`
  - (TokenParams includes PersonParams, so this also allows it to be specified on token.Create)
* Add `gb_vat` as a member of `TaxIDType` and `CheckoutSessionCustomerDetailsTaxIDsType`
