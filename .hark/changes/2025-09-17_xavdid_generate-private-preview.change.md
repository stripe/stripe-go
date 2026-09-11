---
title: generate private-preview SDK w/ mid Sept changes
pr_link: https://github.com/stripe/stripe-go/pull/2118
is_stripe_api_change: true
released_in_version: 82.6.0-alpha.2
---

* Add support for `Get` method on resource `V2CoreClaimableSandbox`
* Add support for `MonthOfYear` on `V2BillingCadenceBillingCycleMonthParams` and `V2BillingCadenceBillingCycleMonth`
* Add support for `ClaimedAt`, `ExpiresAt`, `SandboxDetails`, and `Status` on `V2CoreClaimableSandbox`
* Remove support for `APIKeys` on `V2CoreClaimableSandbox`
* Add support for new value `current_billing_period_end` on enum `V2BillingIntentActionDeactivateEffectiveAt.Type`
* Add support for `WillActivateAt` and `WillCancelAt` on `V2BillingPricingPlanSubscriptionServicingStatusTransitions` and `V2BillingRateCardSubscriptionServicingStatusTransitions`
* Add support for `Category` and `Priority` on `V2BillingServiceActionCreditGrantParams`, `V2BillingServiceActionCreditGrantPerTenantParams`, `V2BillingServiceActionCreditGrantPerTenant`, and `V2BillingServiceActionCreditGrant`
* Add support for `invoices` on `EventsV2BillingCadenceBilledEvent`
* Add support for thin events `V2CoreClaimableSandboxClaimedEvent`, `V2CoreClaimableSandboxExpiredEvent`, `V2CoreClaimableSandboxExpiringEvent`, and `V2CoreClaimableSandboxSandboxDetailsOwnerAccountUpdatedEvent` with related object `V2.Core.ClaimableSandbox`
* Remove support for thin event `V2BillingCadenceErroredEvent` with related object `V2.Billing.Cadence`
