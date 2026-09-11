---
title: Fix segfault in `BackendImplementation.handleResponseBufferingErrors`
pr_link: https://github.com/stripe/stripe-go/pull/2212
released_in_version: 83.2.1
---

* Fixes [#2111](https://github.com/stripe/stripe-go/issues/2211) where a network issue during a RawRequest was causing a segfault.
