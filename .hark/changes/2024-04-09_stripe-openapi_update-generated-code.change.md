---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1844
is_stripe_api_change: true
released_in_version: 76.25.0
---

* Add support for new resources `Entitlements.ActiveEntitlement` and `Entitlements.Feature`
* Add support for `Get` and `List` methods on resource `ActiveEntitlement`
* Add support for `Get`, `List`, `New`, and `Update` methods on resource `Feature`
* Add support for `Controller` on `AccountParams`
* Add support for `Fees`, `Losses`, `RequirementCollection`, and `StripeDashboard` on `AccountController`
* Add support for new value `none` on enum `AccountType`
* Add support for `EventName` on `BillingMeterEventAdjustmentParams` and `BillingMeterEventAdjustment`
* Add support for `Cancel` and `Type` on `BillingMeterEventAdjustment`
