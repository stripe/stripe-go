---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2204
is_stripe_api_change: true
released_in_version: 83.3.0-alpha.2
---

* Remove support for resource `V2TaxAutomaticRule`
* Remove support for `Deactivate`, `Find`, `Get`, `New`, and `Update` methods on resource `V2TaxAutomaticRule`
* Add support for `SelfReportedIncome` and `SelfReportedMonthlyHousingPayment` on `AccountIndividualParams`, `AccountPersonParams`, `Person`, `TokenAccountIndividualParams`, and `TokenPersonParams`
* Add support for `BillingSchedules` and `PhaseEffectiveAt` on `QuoteSubscriptionDataOverrideParams`, `QuoteSubscriptionDataOverridesParams`, `QuoteSubscriptionDataOverrides`, `QuoteSubscriptionDataParams`, and `QuoteSubscriptionData`
* Add support for `BillFrom` on `SubscriptionBillingSchedule`
* Add support for `AmendmentEnd` and `LineEndsAt` on `SubscriptionBillingScheduleBillUntil`
* Add support for new values `amendment_end`, `line_ends_at`, `schedule_end`, and `upcoming_invoice` on enum `SubscriptionBillingScheduleBillUntil.Type`
