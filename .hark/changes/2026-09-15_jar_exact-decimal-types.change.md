---
title: Use exact decimal types for decimal API fields
semver_level: major
---

* Change decimal API fields from `float64`/`*float64` to `decimal.Decimal`/`*decimal.Decimal`.
* Add `github.com/shopspring/decimal` as a public dependency.
