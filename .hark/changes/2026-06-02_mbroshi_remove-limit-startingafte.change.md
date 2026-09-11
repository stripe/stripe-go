---
title: Remove `Limit`, `StartingAfter`, and `EndingBefore` fields for `List` methods that do not accept those fields
pr_link: https://github.com/stripe/stripe-go/pull/2368
released_in_version: 85.3.0-alpha.2
---

<!-- Include any links or additional information that help explain this change. -->
- Fixes a bug where `Limit`, `StartingAfter`, and `EndingBefore` were embedded in `CapabilityListParams`, `PaymentLocationCapabilityListParams`, and `ReportingReportTypeListParams` even though they are not valid parameters and would have always resulted in a 400 from the Stripe API if set. If you were including them before in any of those 3 structs, you can safely remove them.
