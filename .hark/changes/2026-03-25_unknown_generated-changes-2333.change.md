---
title: "Generated changes from [#2333](https://github.com/stripe/stripe-go/pull/2333), [#2326](https://github.com/stripe/stripe-go/pull/2326), [#2323](https://github.com/stripe/stripe-go/pull/2323), [#2286](https://github.com/stripe/stripe-go/pull/2286)"
is_breaking: true
is_stripe_api_change: true
section: ⚠️ Breaking changes due to changes in the Stripe API
released_in_version: 85.0.0
---

* Add support for `UpiPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `Upi` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `MandatePaymentMethodDetails`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodDataParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for new value `tempo` on enums `ChargePaymentMethodDetailsCrypto.Network`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.Network`, and `PaymentRecordPaymentMethodDetailsCrypto.Network`
* Add support for `IntegrationIdentifier` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `Crypto` on `CheckoutSessionPaymentMethodOptionsParams`
* Add support for `PendingInvoiceItemInterval` on `CheckoutSessionSubscriptionDataParams`
* Add support for new values `elements`, `embedded_page`, `form`, and `hosted_page` on enum `CheckoutSession.UIMode`
* Add support for new value `marine_carbon_removal` on enum `ClimateSupplier.RemovalPathway`
* Add support for new value `upi` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `Metadata` on `CreditNoteLineItem`, `CreditNoteLineParams`, `CreditNotePreviewLineParams`, and `CreditNotePreviewLinesLineParams`
* Add support for `QuantityDecimal` on `InvoiceAddLinesLineParams`, `InvoiceCreatePreviewInvoiceItemParams`, `InvoiceItemParams`, `InvoiceItem`, `InvoiceLineItemParams`, `InvoiceLineItem`, and `InvoiceUpdateLinesLineParams`
* ⚠️ Add support for `Level` on `IssuingAuthorizationRiskAssessmentCardTestingRiskParams` and `IssuingAuthorizationRiskAssessmentMerchantDisputeRiskParams`
* ⚠️ Remove support for `RiskLevel` on `IssuingAuthorizationRiskAssessmentCardTestingRiskParams` and `IssuingAuthorizationRiskAssessmentMerchantDisputeRiskParams`
* Add support for `LifecycleControls` on `IssuingCardParams` and `IssuingCard`
* Add support for `Cryptogram`, `ElectronicCommerceIndicator`, `ExemptionIndicatorApplied`, and `ExemptionIndicator` on `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure` and `PaymentRecordPaymentMethodDetailsCardThreeDSecure`
* Add support for new value `upi` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `UpiHandleRedirectOrDisplayQRCode` on `PaymentIntentNextAction` and `SetupIntentNextAction`
* Add support for new value `upi` on enum `PaymentLink.PaymentMethodTypes`
* Add support for `RecommendedAction` and `Signals` on `RadarPaymentEvaluation`
* ⚠️ Remove support for `Insights` on `RadarPaymentEvaluation`
* Add support for new value `crypto_fingerprint` on enum `RadarValueList.ItemType`
* Add support for new value `canceled_by_retention_policy` on enum `SubscriptionCancellationDetails.Reason`
* ⚠️ Change type of `V2CoreEventDestination.EventsFrom` from `enum('other_accounts'|'self')` to `string`
* Add support for error code `service_period_coupon_with_metered_tiered_item_unsupported` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
