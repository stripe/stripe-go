---
title: Represent dynamic objects as maps
pr_url: https://github.com/stripe/stripe-go/pull/2441
semver_level: major
released_in_version: 87.1.0-alpha.1
---

* Change type of `CouponScriptParams.Configuration`, `CouponCreateScriptParams.Configuration`, `CouponScript.Configuration`, and `CryptoOnrampTransactionLimits.Limits` from an empty struct to `map[string]any`
