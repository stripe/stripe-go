---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1291
is_stripe_api_change: true
released_in_version: 72.48.0
---

* Add new resource `TaxCode`.
* Add support for `AutomaticTax` on `CheckoutSession`, `Invoice`, `Subscription`, and `SubscriptionScheduleDefaultSettings`.
* Add support for `CustomerUpdate` on `CheckoutSessionCustomerUpdateParams`
* Add support for `Tax` on `Customer` and `CustomerParams`
* Add support for `CustomerDetails` on `InvoiceParams`
* Add support for `TaxBehavior` on `Price`, `PriceParams`, `CheckoutSessionLineItemPriceDataParams`,  `PriceParams`, `SubscriptionItemPriceDataParams`, `SubscriptionSchedulePhaseAutomaticTaxParams`,`SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams`, and `InvoiceItemPriceDataParams`
* Add support for `TaxCode` on `CheckoutSessionLineItemPriceDataProductParams`, `Product`, `ProductParams`, `PlanProductParams` and `PriceProductDataParams`
