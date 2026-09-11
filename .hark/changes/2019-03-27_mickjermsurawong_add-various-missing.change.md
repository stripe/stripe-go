---
title: Add various missing parameters
pr_link: https://github.com/stripe/stripe-go/pull/820
released_in_version: 60.0.0
---

* On `PIIParams` the previous `PersonalIDNumber` is fixed to `IDNumber` which we're releasing as a minor breaking change even though the old version probably didn't work correctly
