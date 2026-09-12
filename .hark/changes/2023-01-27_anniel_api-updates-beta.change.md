---
title: API Updates for beta branch
pr_url: https://github.com/stripe/stripe-go/pull/1598
is_stripe_api_change: true
released_in_version: 74.7.0-beta.2
---

* Updated stable APIs to the latest version
* Add support for `ListTransactions` method on resource `Tax.Transaction`
* Add support for `BillingAgreementID` on `CheckoutSessionPaymentMethodOptionsPaypalParams`
* Change type of `QuoteSubscriptionDataOverridesParams` from `array(create_specs)` to `emptyStringable(array(update_specs))`
