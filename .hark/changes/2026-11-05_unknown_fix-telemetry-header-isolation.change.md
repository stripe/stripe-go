---
title: Honor each client's `EnableTelemetry` setting
semver_level: patch
---

- Separate the `X-Stripe-Client-User-Agent` cache so that clients with different telemetry settings coexist in the same process.
