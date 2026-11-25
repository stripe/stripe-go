---
title: "`ConstructEvent`, `Client.ConstructEvent`, `Client.ParseEventNotification`, and the `webhook` package's `ConstructEvent*` functions now verify and parse a private copy of the payload"
semver_level: patch
---

- The `webhook` package's `ConstructEvent*` functions now call `stripe.ConstructEvent`, so passing them a thin event notification returns the root package's error message.
