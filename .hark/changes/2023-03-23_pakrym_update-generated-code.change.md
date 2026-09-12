---
title: Update generated code for beta (new)
pr_url: https://github.com/stripe/stripe-go/pull/1623
is_stripe_api_change: true
released_in_version: 74.14.0-beta.1
---

* Add support for new resources `Tax.CalculationLineItem` and `Tax.TransactionLineItem`
* Add support for `CollectInputs` method on resource `Terminal.Reader`
* Add support for `FinancingOffer` on `CapitalFinancingSummary`
* Add support for `FxRate` on `CheckoutSessionCurrencyConversion`
* Add support for new value `link` on enum `PaymentLinkPaymentMethodTypes`
* Add support for `AutomaticPaymentMethods` on `SetupIntentParams` and `SetupIntent`
* Remove support for `Preview` on `TaxCalculationParams`
* Add support for `TaxBreakdown` on `TaxCalculation`
* Remove support for `TaxSummary` on `TaxCalculation`
* Change type of `TaxCalculationLineItems` from `$LineItem` to `$Tax.CalculationLineItem`
* Change type of `TaxTransactionLineItems` from `$LineItem` to `$Tax.TransactionLineItem`
* Add support for `CollectInputs` on `TerminalReaderAction`
* Add support for new value `collect_inputs` on enum `TerminalReaderActionType`
