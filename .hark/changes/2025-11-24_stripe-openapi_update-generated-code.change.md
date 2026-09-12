---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2224
is_stripe_api_change: true
released_in_version: 84.1.0-alpha.3
---

* Add support for new resource `ProductCatalogTrialOffer`
* Add support for `New` method on resource `ProductCatalogTrialOffer`
* Remove support for `AmountSubtotalAfterDiscount` on `DelegatedCheckoutRequestedSessionLineItemDetail` and `DelegatedCheckoutRequestedSessionTotalDetails`
* Remove support for `AmountTotal`, `UnitAmountAfterDiscount`, and `UnitDiscount` on `DelegatedCheckoutRequestedSessionLineItemDetail`
* Add support for `AmountCartDiscount` and `AmountItemsDiscount` on `DelegatedCheckoutRequestedSessionTotalDetails`
* Remove support for `AmountDiscount` on `DelegatedCheckoutRequestedSessionTotalDetails`
* Add support for `PaymentsOrchestration` on `PaymentIntentParams` and `PaymentIntent`
