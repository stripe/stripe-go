---
title: Fix MinimumBalanceByCurrency map value type to support clearing
pr_url: https://github.com/stripe/stripe-go/pull/2324
is_breaking: true
released_in_version: 85.0.0
---

* `MinimumBalanceByCurrency` changed from `map[string]int64` to `map[string]*int64` on `BalanceSettingsPaymentsPayoutsParams` and `BalanceSettingsUpdatePaymentsPayoutsParams`. This field now supports clearing a value in the map by assigning null to the map key.
