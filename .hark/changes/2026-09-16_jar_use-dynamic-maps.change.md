---
title: Represent free-form API dictionaries as dynamic maps
pr_url: https://github.com/stripe/stripe-go/pull/2438
semver_level: patch
---

- Change `PaymentIntentNextAction.UseStripeSDK` from `*PaymentIntentNextActionUseStripeSDK` to `map[string]any` and remove `PaymentIntentNextActionUseStripeSDK`.  This file is inteded for use by Stripe.js.
- Change `SetupIntentNextAction.UseStripeSDK` from `*SetupIntentNextActionUseStripeSDK` to `map[string]any` and remove `SetupIntentNextActionUseStripeSDK`.  This file is inteded for use by Stripe.js.
