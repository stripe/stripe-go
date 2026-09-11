---
title: "codegen: 14 more files"
pr_link: https://github.com/stripe/stripe-go/pull/1312
released_in_version: 72.57.0
---

* Add support for `BillingAddressCollection` to `CheckoutSession`
* Add support for `NetworkReasonCode` to `DisputeReason`
* Add support for `Object` to `EphemeralKey`, `ApplicationFee`, and `DisputeReason`
* Add support for `Description` to `Refund`
* Add const definition for value `blocked` on enum `IssuingCardholderStatus`
* Bugfix: add support for `Rate` on `CheckoutSessionTotalDetailsBreakdownTax` -- the existing field `TaxRate` has the wrong json annotation and should be deprecated.
