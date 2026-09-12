---
title: Various API changes
pr_url: https://github.com/stripe/stripe-go/pull/972
released_in_version: 65.2.0
---

* `Requirements` on Issuing `Cardholder`
* `PaymentMethodDetails.AuBecsDebit.Mandate` on `Charge`
* `PaymentBehavior` on `Subscription` creation can now take the value `pending_if_incomplete`
* `PaymentBehavior` on `SubscriptionItem` creation is now supported
* `SubscriptionData.TrialFromPlan` is now supported on Checkout `Session` creation
* New values for `TaxIDType`
