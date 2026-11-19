---
title: Honor each client's `EnableTelemetry` setting
semver_level: patch
released_in_version: 87.2.0
---

- Separate the `X-Stripe-Client-User-Agent` cache so that clients with different telemetry settings coexist in the same process.
