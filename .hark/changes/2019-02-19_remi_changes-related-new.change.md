---
title: "Changes related to the new API version `2019-02-19`:"
pr_url: https://github.com/stripe/stripe-go/pull/782
released_in_version: 57.0.0
---

* The library is now pinned to API version `2019-02-19`
* Numerous changes to the `Account` resource and APIs:
  * The `legal_entity` property on the Account API resource has been replaced with `individual`, `company`, and `business_type`
  * The `verification` hash has been replaced with a `requirements` hash
  * Multiple top-level properties were moved to the `settings` hash
  * The `keys` property on `Account` has been removed. Platforms should authenticate as their connected accounts with their own key via the `Stripe-Account` [header](https://stripe.com/docs/connect/authentication#authentication-via-the-stripe-account-header)
* The `requested_capabilities` property on `Account` creation is now required for accounts in the US
* The deprecated parameter `save_source_to_customer` on `PaymentIntent` has now been removed. Use `save_payment_method` instead
