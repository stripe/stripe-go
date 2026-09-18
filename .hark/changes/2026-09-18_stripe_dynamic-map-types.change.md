---
title: Represent dynamic objects as maps
pr_url: https://github.com/stripe/stripe-go/pull/2441
semver_level: major
---

* ⚠️ Change type of `CouponScriptParams.Configuration` from `*CouponScriptConfigurationParams` to `map[string]any`
* ⚠️ Change type of `CouponCreateScriptParams.Configuration` from `*CouponCreateScriptConfigurationParams` to `map[string]any`
* ⚠️ Change type of `CouponScript.Configuration` from `*CouponScriptConfiguration` to `map[string]any`
* ⚠️ Change type of `CryptoOnrampTransactionLimits.Limits` from `*CryptoOnrampTransactionLimitsLimits` to `map[string]any`
