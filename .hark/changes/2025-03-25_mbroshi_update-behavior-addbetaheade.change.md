---
title: Update behavior of AddBetaHeader
pr_url: https://github.com/stripe/stripe-go/pull/2002
section: Changes
released_in_version: 82.1.0-beta.1
---

- AddBetaVersion will use the highest version number used for a beta feature instead of return an `error` on a conflict as it had done previously.
