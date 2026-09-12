---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2428
is_stripe_api_change: true
released_in_version: 86.5.0-alpha.3
---

* Add support for `CustomerTaxExemption` on `TaxCalculationLineItemTaxBreakdown`, `TaxCalculationShippingCostTaxBreakdown`, and `TaxTransactionShippingCostTaxBreakdown`
* Add support for new value `data_share_only` on enums `ChargePaymentMethodDetailsCardThreeDSecure.Result`, `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure.Result`, `PaymentRecordPaymentMethodDetailsCardThreeDSecure.Result`, and `SetupAttemptPaymentMethodDetailsCardThreeDSecure.Result`
* Add support for `BackdateStartDate` on `CheckoutSessionItemSubscriptionParams` and `CheckoutSessionItemSubscription`
* Add support for `Signals` on `IdentityVerificationReport`
* Add support for `NetworkResponseCode` on `IssuingAuthorizationRequestHistory`
* Add support for `UnitCostPrecision` on `PaymentIntentAmountDetailsLineItem`, `PaymentIntentAmountDetailsLineItemsParams`, `PaymentIntentCaptureAmountDetailsLineItemsParams`, `PaymentIntentConfirmAmountDetailsLineItemsParams`, `PaymentIntentDecrementAuthorizationAmountDetailsLineItemsParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsParams`
* Add support for `Active` on `ProductCatalogTrialOfferListParams`
* Add support for new value `rtp` on enum `TreasuryFinancialAccountFinancialAddress.SupportedNetworks`
* Add support for new value `blik_recurring_payments` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
