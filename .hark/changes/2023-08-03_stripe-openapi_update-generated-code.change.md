---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1697
is_stripe_api_change: true
released_in_version: 74.30.0-beta.1
---

* Add support for `SubmitCard` test helper method on resource `Issuing.Card`
* Add support for `AddressValidation` on `IssuingCardShippingParams` and `IssuingCardShipping`
* Add support for new value `submitted` on enum `IssuingCardShippingStatus`
* Change type of `OrderDescriptionParams`, `OrderLineItemsProductDataDescriptionParams`, `OrderLineItemsProductDataTaxCodeParams`, `OrderShippingDetailsPhoneParams`, `PaymentMethodConfigurationListApplicationParams`, and `QuoteSubscriptionDataOverridesDescriptionParams` from `string` to `emptyStringable(string)`
* Add support for `Reason` on `QuoteMarkStaleQuoteParams`
* Add support for `MarkedStale` on `QuoteStatusDetailsStaleLastReason`
