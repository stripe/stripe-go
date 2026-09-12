---
title: Unexport GetBaseEvent and V2ErrorCode
pr_url: https://github.com/stripe/stripe-go/pull/2054
is_breaking: true
released_in_version: 82.2.0-beta.1
---

* ⚠️ Unexported `GetBaseEvent` --> `getBaseEvent`. This function should not be called. Instead, type-cast the `V2Event` to a concrete Event struct.
* ⚠️ Removed the `V2ErrorCode` type. Error codes should not be handled programmatically for V2 errors.
