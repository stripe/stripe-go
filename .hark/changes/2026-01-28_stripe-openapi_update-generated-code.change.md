---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2259
is_stripe_api_change: true
released_in_version: 84.4.0-alpha.1
---

* Add support for new resources `FRMealVouchersOnboarding`, `ReserveHold`, `ReservePlan`, and `ReserveRelease`
* Add support for `Get`, `List`, `New`, and `Update` methods on resource `FRMealVouchersOnboarding`
* Add support for `Get` and `List` methods on resources `ReserveHold` and `ReserveRelease`
* Add support for `Get` method on resource `ReservePlan`
* Add support for `Pause` method on resource `Subscription`
* Add support for `ServicePeriodDetails` on `Discount`
* Add support for `AgenticCommerceSettings` on `AccountSessionComponents`
* Add support for new value `risk_reserved` on enum `BalanceTransaction.BalanceType`
* Add support for `ServicePeriod` on `CouponParams` and `Coupon`
* Add support for new value `service_period` on enum `Coupon.Duration`
* Change type of `InvoiceItemPricingPriceDetails.Price` and `InvoiceLineItemPricingPriceDetails.Price` from `string` to `expandable($Price)`
* Add support for `Settings` on `InvoiceCreatePreviewDiscountsParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentDiscountActionAddParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentDiscountActionSetParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionAddDiscountParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionSetDiscountParams`, `InvoiceCreatePreviewScheduleDetailsPhaseDiscountsParams`, `InvoiceCreatePreviewScheduleDetailsPhaseItemDiscountsParams`, `InvoiceCreatePreviewSubscriptionDetailsItemDiscountsParams`, `QuoteLineActionAddDiscountParams`, `QuoteLineActionAddItemDiscountParams`, `QuoteLineActionSetDiscountParams`, `QuoteLineActionSetItemDiscountParams`, `SubscriptionDiscountsParams`, `SubscriptionItemDiscountsParams`, `SubscriptionScheduleAmendAmendmentDiscountActionAddParams`, `SubscriptionScheduleAmendAmendmentDiscountActionSetParams`, `SubscriptionScheduleAmendAmendmentItemActionAddDiscountParams`, `SubscriptionScheduleAmendAmendmentItemActionSetDiscountParams`, `SubscriptionSchedulePhaseDiscountsParams`, and `SubscriptionSchedulePhaseItemDiscountsParams`
* Add support for `Subtotal` on `InvoiceLineItem`
* Add support for `BillingCadence` on `SubscriptionListParams`
