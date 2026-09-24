---
title: Multiple API changes
pr_url: https://github.com/stripe/stripe-go/pull/1170
is_stripe_api_change: true
released_in_version: 72.0.0
---

* Move to latest API version `2020-08-27`
* Remove `Prorate` across Billing APIs in favor of `ProrationBehavior`
* Remove `TaxPercent` across Billing APIs in favor of `TaxRate`-related parameters and properties
* Remove `DisplayItems` on Checkout `Session` in favor of `LineItems`
* Remove `FailureURL` and `SuccessURL` on `AccountLink` in favor of `RefreshURL` and `ReturnURL`
* Remove `AccountLinkTypeCustomAccountUpdate ` and `AccountLinkTypeCustomAccountVerification ` on `AccountLink` in favor of `AccountLinkTypeAccountOnboarding ` and `AccountLinkTypeAccountUpdate `
* Remove `Authenticated` and `Succeeded` on `ChargePaymentMethodDetailsCardThreeDSecure`
* Remove `Plan`, `Quantity`, `TaxPercent` and `TrialEnd` from `Customer` creation or update in favor of the Subscription API
* Rename `Plans` to `Items` on `SubscriptionSchedule`
