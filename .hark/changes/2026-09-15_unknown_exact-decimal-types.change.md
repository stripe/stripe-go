---
title: Use exact decimal types for decimal API fields
semver_level: major
---

* ⚠️ Change decimal API fields from `float64`/`*float64` to `decimal.Decimal`/`*decimal.Decimal`. See the [v87 migration guide](https://github.com/stripe/stripe-go/wiki/Migration-guide-for-v87) for complete upgrade guidance.
* Add `github.com/shopspring/decimal` as a public dependency.
