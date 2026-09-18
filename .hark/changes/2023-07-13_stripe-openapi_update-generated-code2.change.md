---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1688
is_stripe_api_change: true
released_in_version: 74.26.0
---

* Add support for new resource `Tax.Settings`
* Add support for `Get` and `Update` methods on resource `Settings`
* Add support for new value `invalid_tax_location` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `OrderID` on `ChargePaymentMethodDetailsAfterpayClearpay`
* Add support for `AllowRedirects` on `PaymentIntentAutomaticPaymentMethodsParams`, `PaymentIntentAutomaticPaymentMethods`, `SetupIntentAutomaticPaymentMethodsParams`, and `SetupIntentAutomaticPaymentMethods`
* Add support for new values `amusement_tax` and `communications_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationTaxBreakdownTaxRateDetailsTaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetailsTaxType`
* Add support for `Product` on `TaxTransactionLineItem`
