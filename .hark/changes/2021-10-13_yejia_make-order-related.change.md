---
title: Make order-related files codegen-able
pr_link: https://github.com/stripe/stripe-go/pull/1355
is_stripe_api_change: true
released_in_version: 72.71.0
---

* Add support for `SelectedShippingMethod` and `Status` on `OrderStatus`
* Add support for `Carrier` and `TrackingNumber` on `ShippingParams`
* Add support for `ExternalCouponCode` and `Object` on `Order`
* Add support for `Object` on `OrderItem` and `OrderReturn`
* Add support for `Deleted` and `Object` on `SKU`
