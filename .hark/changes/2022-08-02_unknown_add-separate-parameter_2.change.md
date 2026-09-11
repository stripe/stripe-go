---
title: Add separate parameter struct for Invoice `GetNext` (renamed to `Upcoming`) method (`InvoiceUpcomingParams`, and nested params `InvoiceUpcomingLinesInvoiceItemPriceDataParams`, `InvoiceUpcomingLinesInvoiceItemDiscountParams`, `InvoiceUpcomingLinesDiscountParams`, `InvoiceUpcomingLinesInvoiceItemPeriodParams`). `Upcoming`-only fields `Coupon`, `CustomerDetails`, `InvoiceItems`, `Subscription`, `SubscriptionBillingCycleAnchor`, `Schedule`, `SubscriptionBillingCycleAnchor`, `SubscriptionBillingCycleAnchorNow`, `SubscriptionBillingCycleAnchorUnchanged`, `SubscriptionCancelAt`, `SubscriptionCancelAtPeriodEnd`, `SubscriptionCancelNow`, `SubscriptionDefaultTaxRates`, `SubscriptionItems`, `SubscriptionProrationBehavior`, `SubscriptionProrationDate`, `SubscriptionStartDate`, `SubscriptionTrialEnd`, `SubscriptionTrialEndNow`, and `SubscriptionTrialFromPlan` are removed from `InvoiceParams`.
is_breaking: true
section: ⚠️ Changed
released_in_version: 73.0.0
---
