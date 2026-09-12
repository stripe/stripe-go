---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1624
is_stripe_api_change: true
released_in_version: 74.13.0
---

* Add support for new resources `Tax.CalculationLineItem`, `Tax.Calculation`, `Tax.TransactionLineItem`, and `Tax.Transaction`
* Add support for `ListLineItems` and `New` methods on resource `Calculation`
* Add support for `CreateFromCalculation`, `CreateReversal`, `Get`, `ListLineItems`, and `New` methods on resource `Transaction`
* Add support for `CurrencyConversion` on `CheckoutSession`
* Add support for new value `link` on enum `PaymentLinkPaymentMethodTypes`
* Add support for `AutomaticPaymentMethods` on `SetupIntentParams` and `SetupIntent`
