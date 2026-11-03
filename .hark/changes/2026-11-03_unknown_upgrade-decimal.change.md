---
title: Upgrade decimal dependency to v1.5.0
semver_level: patch
---

- Upgrade `github.com/shopspring/decimal` to v1.5.0. Decimal values with excessively large positive or negative exponents now return decoding errors, retaining the upstream default exponent limit of 10,000.
- See [stripe-go issue #2484](https://github.com/stripe/stripe-go/issues/2484) and the [upstream fix](https://github.com/shopspring/decimal/pull/440).
