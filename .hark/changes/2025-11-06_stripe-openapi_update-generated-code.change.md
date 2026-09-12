---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2200
is_stripe_api_change: true
released_in_version: 83.3.0-alpha.1
---

* Add support for new resources `TransitBalance`, `V2ReportingReportRun`, `V2ReportingReport`
* Add support for `Get` and `New` methods on resource `V2ReportingReportRun`
* Add support for `Get` method on resource `V2ReportingReport`
* Add support for `New` and `Refill` test helper methods on resource `CapitalFinancingOffer`
* Add support for `AllocatedFunds` on `Charge`, `PaymentIntentConfirmParams`, and `PaymentIntentParams`
* Add support for thin events `V2ReportingReportRunCreatedEvent`, `V2ReportingReportRunFailedEvent`, `V2ReportingReportRunSucceededEvent`, and `V2ReportingReportRunUpdatedEvent` with related object `V2ReportingReportRun`
