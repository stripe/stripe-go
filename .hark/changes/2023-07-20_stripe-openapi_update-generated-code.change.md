---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1691
is_stripe_api_change: true
released_in_version: 74.27.0
---

* Add support for new value `ro_tin` on enums `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, and `TaxTransactionCustomerDetailsTaxIdsType`
* Remove support for values `excluded_territory`, `jurisdiction_unsupported`, and `vat_exempt` from enums `CheckoutSessionShippingCostTaxesTaxabilityReason`, `CheckoutSessionTotalDetailsBreakdownTaxesTaxabilityReason`, `CreditNoteShippingCostTaxesTaxabilityReason`, `InvoiceShippingCostTaxesTaxabilityReason`, `LineItemTaxesTaxabilityReason`, `QuoteComputedRecurringTotalDetailsBreakdownTaxesTaxabilityReason`, `QuoteComputedUpfrontTotalDetailsBreakdownTaxesTaxabilityReason`, and `QuoteTotalDetailsBreakdownTaxesTaxabilityReason`
* Add support for `UseStripeSDK` on `SetupIntentConfirmParams` and `SetupIntentParams`
* Add support for new value `service_tax` on enum `TaxRateTaxType`
