---
title: Add `UnsetFields` for clearing field values in v1 and v2 API requests
pr_url: https://github.com/stripe/stripe-go/pull/2322
released_in_version: 85.0.0
---

- Added `UnsetFields` field and `AddUnsetField` method to `Params` for explicitly clearing field values in API requests. For v2 JSON requests, listed fields are sent as `"field": null`. For v1 form requests, listed fields are sent as `field=` (empty string).
- Nested params structs with emptyable fields carry their own `UnsetFields` slice, enabling clearing of nested fields (e.g. `params.CancellationDetails.AddUnsetField(...)`).
- Generated `UnsetField` string enum types provide type-safe constants for each clearable field (e.g. `SubscriptionUpdateParamsUnsetFieldDescription`).
