---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/2336
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.2.0-beta.1
---

* Add support for new resources `SharedPaymentGrantedToken` and `SharedPaymentIssuedToken`
* Add support for `Get` method on resource `SharedPaymentGrantedToken`
* Add support for `New` and `Revoke` test helper methods on resource `SharedPaymentGrantedToken`
* Add support for `Get`, `New`, and `Revoke` methods on resource `SharedPaymentIssuedToken`
* Add support for `BLIK` on `CheckoutSessionPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for new values `fo_vat`, `gi_tin`, `it_cf`, and `py_ruc` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, and `QuotePreviewInvoiceCustomerTaxIds.Type`
* Add support for `SharedPaymentGrantedToken` on `ConfirmationTokenPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `ValidationErrors` on `PrivacyRedactionJob`
* Add support for `TaxDetails` on `Product`
* Add support for new value `blik` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
* ⚠️ Change type of `QuotePreviewInvoiceTotalTaxesTaxRateDetails.TaxRate` from `string` to `expandable($TaxRate)`
* Add support for `AdmissionsTax`, `AttendanceTax`, `EntertainmentTax`, `GrossReceiptsTax`, `HospitalityTax`, `LuxuryTax`, `ResortTax`, and `TourismTax` on `TaxRegistrationCountryOptionsUsParams`
* Add support for `Purpose` on `TreasuryOutboundPaymentParams` and `TreasuryOutboundPayment`
* Add support for error codes `action_blocked` and `approval_required` on `QuotePreviewInvoiceLastFinalizationError`
