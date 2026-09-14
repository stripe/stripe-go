---
title: Remove `Limit`, `StartingAfter`, and `EndingBefore` fields for `List` methods that do not accept those fields
pr_url: https://github.com/stripe/stripe-go/pull/2367
released_in_version: 86.0.0
---

<!-- Include any links or additional information that help explain this change. -->
- Fixes a bug where `Limit`, `StartingAfter`, and `EndingBefore` were embedded in `CapabilityListParams` and `ReportingReportTypeListParams` even though they are not valid parameters and would have always resulted in a 400 from the Stripe API if set. If you were including them before in either of those 2 structs, you can safely remove them.
