---
title: Move `V2.Event` API resources to `V2.Core.Events`
pr_url: https://github.com/stripe/stripe-go/pull/2138
semver_level: major
released_in_version: 83.0.0
---

- ⚠️ Rename all types starting with `V2Event` to start with `V2CoreEvent`. For example
   - `V2EventNotification` -> `V2CoreEventNotification`
   - `V2EventReason` -> `V2CoreEventReason`
   - `V2Event` -> `V2CoreEvent`
   - `V2RawEvent` -> `V2CoreRawEvent`
   - `V2EventDestination` -> V`2CoreEventDestination`
