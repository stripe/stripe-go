---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1453
is_stripe_api_change: true
released_in_version: 72.104.0
---

* Add support for new resource `CashBalance`
* Change type of `BillingPortalConfigurationApplication` from `$Application` to `deletable($Application)`
* Add support for `Alipay` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for new value `eu_oss_vat` on enums `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, and `TaxIdType`
* Add support for `CashBalance` on `Customer`
* Add support for `Application` on `Invoice`, `Quote`, `SubscriptionSchedule`, and `Subscription`
