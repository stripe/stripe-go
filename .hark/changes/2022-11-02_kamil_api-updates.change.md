---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1563
is_stripe_api_change: true
released_in_version: 73.15.0
---

* Add support for `OnBehalfOf` on `CheckoutSessionSubscriptionDataParams`, `SubscriptionParams`, `SubscriptionScheduleDefaultSettingsParams`, `SubscriptionScheduleDefaultSettings`, `SubscriptionSchedulePhasesParams`, `SubscriptionSchedulePhases`, and `Subscription`
* Add support for new values `eg_tin`, `ph_tin`, and `tr_tin` on enums `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `OrderTaxDetailsTaxIdsType`, and `TaxIdType`
* Add support for `TaxBehavior` and `TaxCode` on `InvoiceItemParams`, `InvoiceUpcomingInvoiceItemsParams`, and `InvoiceUpcomingLinesInvoiceItemsParams`
