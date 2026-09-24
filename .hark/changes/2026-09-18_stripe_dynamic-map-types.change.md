---
title: Represent dynamic objects as maps
pr_url: https://github.com/stripe/stripe-go/pull/2441
semver_level: major
---

* ⚠️ Change type of `CouponScriptParams.Configuration`, `CouponCreateScriptParams.Configuration`, `CouponScript.Configuration`, and `CryptoOnrampTransactionLimits.Limits` from an empty struct to `map[string]any`
