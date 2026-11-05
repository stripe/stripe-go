---
title: Honor each client's telemetry setting in cached headers
semver_level: patch
---

- Honor each client's `EnableTelemetry` setting in `X-Stripe-Client-User-Agent` when clients with different telemetry settings coexist in the same process.
