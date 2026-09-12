---
title: Do not allow setting API Version directly
pr_url: https://github.com/stripe/stripe-go/pull/1941
released_in_version: 81.1.0-beta.1
---

* `stripe.APIVersion` is no longer settable. If you were using this to set the beta headers, use the helper method `stripe.AddBetaVersion()` instead.
