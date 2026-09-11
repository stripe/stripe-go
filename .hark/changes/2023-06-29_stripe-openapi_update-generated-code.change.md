---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1682
is_stripe_api_change: true
released_in_version: 74.24.0
---

* Add support for new value `application_fees_not_allowed` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for new values `ad_nrt`, `ar_cuit`, `bo_tin`, `cn_tin`, `co_nit`, `cr_tin`, `do_rcn`, `ec_ruc`, `pe_ruc`, `rs_pib`, `sv_nit`, `uy_ruc`, `ve_rif`, and `vn_tin` on enums `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, and `TaxTransactionCustomerDetailsTaxIdsType`
* Add support for `EffectiveAt` on `CreditNoteParams`, `CreditNotePreviewLinesParams`, `CreditNotePreviewParams`, `CreditNote`, `InvoiceParams`, and `Invoice`
