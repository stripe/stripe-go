---
title: API Updates for beta branch
pr_link: https://github.com/stripe/stripe-go/pull/1498
is_stripe_api_change: true
released_in_version: 72.121.0-beta.1
---

- Updated stable APIs to the latest version
- Add `Price.MigrateTo` property
- Add `SubscriptionSchedule.Amend` method.
- Add `Discount.SubscriptionItem` property.
- Add `Quote.SubscriptionData.BillingBehavior`, `BillingCycleAnchor`, `EndBehavior`, `FromSchedule`, `FromSubscription`, `Prebilling`, `ProrationBehavior` properties.
- Add `Phases` parameter to `Quote.Create`
- Add `Subscription.Discounts`, `Prebilling` properties.
