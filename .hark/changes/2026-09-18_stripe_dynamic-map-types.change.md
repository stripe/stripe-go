---
title: Represent script configuration as dynamic maps
pr_url: https://github.com/stripe/stripe-go/pull/2440
semver_level: major
---

* ⚠️ Change type of `CouponScriptParams.Configuration`, `CouponCreateScriptParams.Configuration`, and `CouponScript.Configuration` from an empty struct to `map[string]any`
