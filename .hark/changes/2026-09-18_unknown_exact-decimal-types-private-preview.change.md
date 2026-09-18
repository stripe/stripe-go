---
title: Use exact decimal types for private-preview decimal API fields
semver_level: major
---

* Change all decimal API request fields from `*float64` to `*decimal.Decimal`.
* Change all decimal API response fields from `float64` to `decimal.Decimal`.
* Add `github.com/shopspring/decimal` as a public dependency.
