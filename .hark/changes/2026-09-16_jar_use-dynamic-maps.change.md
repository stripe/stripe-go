---
title: Represent free-form API dictionaries as dynamic maps
semver_level: major
section: ⚠️ Changed
---

- ⚠️ Change `PaymentIntentNextAction.UseStripeSDK` from `*PaymentIntentNextActionUseStripeSDK` to `map[string]any` and remove `PaymentIntentNextActionUseStripeSDK`.
- ⚠️ Change `SetupIntentNextAction.UseStripeSDK` from `*SetupIntentNextActionUseStripeSDK` to `map[string]any` and remove `SetupIntentNextActionUseStripeSDK`.

These fields can now retain and expose arbitrary keys returned by the Stripe API.
