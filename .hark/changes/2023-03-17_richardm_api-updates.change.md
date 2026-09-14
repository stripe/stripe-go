---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1621
is_stripe_api_change: true
released_in_version: 74.13.0-beta.1
---

* Add support for `CreateFromCalculation` method on resource `Tax.Transaction`
* Change type of `InvoiceAppliesTo` from `nullable(QuotesResourceQuoteLinesAppliesTo)` to `QuotesResourceQuoteLinesAppliesTo`
* Add support for `Paypal` on `MandatePaymentMethodDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for `SetupFutureUsage` on `OrderPaymentSettingsPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypal`, `PaymentIntentConfirmPaymentMethodOptionsPaypalParams`, `PaymentIntentPaymentMethodOptionsPaypalParams`, and `PaymentIntentPaymentMethodOptionsPaypal`
* Add support for new value `automatic_async` on enums `OrderPaymentSettingsPaymentMethodOptionsAfterpayClearpayCaptureMethod` and `OrderPaymentSettingsPaymentMethodOptionsCardCaptureMethod`
* Remove support for `AppliesTo` on `QuotePreviewInvoiceLinesParams`
* Add support for `ShippingCost` on `TaxCalculationParams`, `TaxCalculation`, `TaxTransactionCreateReversalParams`, and `TaxTransaction`
* Add support for `TaxBreakdown` on `TaxCalculation`
* Remove support for `TaxSummary` on `TaxCalculation`
