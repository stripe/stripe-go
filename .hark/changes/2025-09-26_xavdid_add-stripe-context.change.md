---
title: Add `stripe.Context` object
pr_url: https://github.com/stripe/stripe-go/pull/2133
is_breaking: true
released_in_version: 83.0.0
---

- This is a new struct that helps with accessing parent and child contexts.
- Previously, you could set the stripe context only as a string via `SetStripeContext()`. You can now set it using the new struct as well via `SetStripeContextFrom()`.
- ⚠️ Change `EventNotification` (formerly known as `ThinEvent`)'s `context` property from `string` to `stripe.Context`
