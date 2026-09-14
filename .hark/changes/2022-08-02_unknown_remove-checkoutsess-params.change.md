---
title: Remove `CheckoutSessionPaymentIntentDataParams.Params`, `CheckoutSessionSetupIntentDataParams.Params`, `CheckoutSessionSubscriptionDataParams.Params`. `Params` should only be embedded in root method struct, and has extraneous fields not applicable to child/sub structs.
semver_level: major
section: ⚠️ Removed
released_in_version: 73.0.0
---
