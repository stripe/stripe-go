---
title: Multiple breaking API changes
pr_url: https://github.com/stripe/stripe-go/pull/1068
released_in_version: 71.0.0
---

* `PaymentIntent` is now expandable on `Charge`
* `Percentage` was removed as a filter when listing `TaxRate`
* Removed `RenewalInterval` on `SubscriptionSchedule`
* Removed `Country` and `RoutingNumber` from `ChargePaymentMethodDetailsAcssDebit`
