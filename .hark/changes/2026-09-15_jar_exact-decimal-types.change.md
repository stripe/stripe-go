---
title: Use exact decimal types for decimal API fields
semver_level: major
released_in_version: 87.0.0
---

* Change decimal API fields from `float64`/`*float64` to `decimal.Decimal`/`*decimal.Decimal`.
* Add `github.com/shopspring/decimal` as a public dependency.
