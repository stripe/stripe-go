---
title: Use standardized names for URL path parameters
pr_url: https://github.com/stripe/stripe-go/pull/2444
semver_level: major
---

Renames generated parameter fields used to construct request URLs. The final
path parameter is now `ID`; earlier parameters use resource-qualified names
such as `CustomerID`, `AccountID`, and `FeeID`.
