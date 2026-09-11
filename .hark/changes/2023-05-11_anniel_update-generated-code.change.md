---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1654
is_stripe_api_change: true
released_in_version: 74.19.0-beta.1
---

* Add support for `PayerEmail`, `PayerName`, and `SellerProtection` on `ChargePaymentMethodDetailsPaypal`
* Add support for `CaptureMethod`, `PreferredLocale`, `ReferenceID`, and `SetupFutureUsage` on `CheckoutSessionPaymentMethodOptionsPaypalParams`
* Add support for `Reference` on `CheckoutSessionPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypal`, `PaymentIntentConfirmPaymentMethodOptionsPaypalParams`, `PaymentIntentPaymentMethodOptionsPaypalParams`, and `PaymentIntentPaymentMethodOptionsPaypal`
* Add support for `RiskCorrelationID` on `CheckoutSessionPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypalParams`, `PaymentIntentConfirmPaymentMethodOptionsPaypalParams`, and `PaymentIntentPaymentMethodOptionsPaypalParams`
* Remove support for `BillingAgreementID` and `Currency` on `CheckoutSessionPaymentMethodOptionsPaypalParams`
* Add support for `Fingerprint`, `PayerID`, and `VerifiedEmail` on `MandatePaymentMethodDetailsPaypal` and `PaymentMethodPaypal`
* Add support for `TaxabilityReason` and `TaxableAmount` on `OrderShippingCostTaxes`, `OrderTotalDetailsBreakdownTaxes`, and `QuotePhaseTotalDetailsBreakdownTaxes`
* Add support for `HeadOffice` on `TaxSettingsParams` and `TaxSettings`
