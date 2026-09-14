---
title: API Updates for beta branch
pr_url: https://github.com/stripe/stripe-go/pull/1591
is_stripe_api_change: true
released_in_version: 74.6.0-beta.2
---

* Updated stable APIs to the latest version
* Add support for new resource `Tax.Registration`
* Add support for `List`, `New`, and `Update` methods on resource `Registration`
* Add support for `Controller` on `AccountParams`
* Add support for `Application` and `Dashboard` on `AccountController`
* Remove support for `Timestamp` on `QuoteLineActionsAddDiscountDiscountEnd`
* Change type of `QuoteLineActionsAddDiscountDiscountEndType` from `literal('timestamp')` to `literal('line_ends_at')`
* Remove support for `Index` on `QuoteLineActionsAddItemDiscounts`, `QuoteLineActionsRemoveDiscount`, `QuoteLineActionsSetDiscounts`, `QuoteLineActionsSetItemsDiscounts`, `SubscriptionSchedulePhasesAddInvoiceItemsDiscounts`, `SubscriptionSchedulePhasesDiscounts`, and `SubscriptionSchedulePhasesItemsDiscounts`
