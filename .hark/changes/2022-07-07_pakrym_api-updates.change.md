---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1492
is_stripe_api_change: true
released_in_version: 72.118.0
---

* Add support for `Currency` on `CheckoutSessionParams`, `InvoiceUpcomingLinesParams`, `InvoiceUpcomingParams`, `PaymentLinkParams`, `SubscriptionParams`, `SubscriptionSchedulePhasesParams`, `SubscriptionSchedulePhases`, and `Subscription`
* Add support for `CurrencyOptions` on `CheckoutSessionShippingOptionsShippingRateDataFixedAmountParams`, `CouponParams`, `Coupon`, `OrderShippingCostShippingRateDataFixedAmountParams`, `PriceParams`, `Price`, `ProductDefaultPriceDataParams`, `PromotionCodeRestrictionsParams`, `PromotionCodeRestrictions`, `ShippingRateFixedAmountParams`, and `ShippingRateFixedAmount`
* Add support for `Restrictions` on `PromotionCodeParams`
* Add support for `FixedAmount` and `TaxBehavior` on `ShippingRateParams`
