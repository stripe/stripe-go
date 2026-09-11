---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1744
is_stripe_api_change: true
released_in_version: 75.10.0-beta.1
---

* Add support for `MarkDraft` and `MarkStale` methods on resource `Quote`
* Remove support for `DraftQuote` and `MarkStaleQuote` methods on resource `Quote`
* Add support for `Liability` on `CheckoutSessionAutomaticTaxParams` and `CheckoutSessionAutomaticTax`
* Add support for `Issuer` on `CheckoutSessionInvoiceCreationInvoiceDataParams` and `CheckoutSessionInvoiceCreationInvoiceData`
* Add support for `InvoiceSettings` on `CheckoutSessionSubscriptionDataParams`
* Add support for `PersonalizationDesign` on `IssuingCardListParams`
* Add support for `AllowBackdatedLines` on `QuoteParams` and `Quote`
