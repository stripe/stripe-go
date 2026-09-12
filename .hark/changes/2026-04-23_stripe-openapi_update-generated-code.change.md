---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2341
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.1.0
---

* Add support for `BalanceReport` and `PayoutReconciliationReport` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for `AppDistribution` and `SunbitPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for new values `fee_credit_funding`, `inbound_transfer_reversal`, and `inbound_transfer` on enum `BalanceTransaction.Type`
* Add support for `Sunbit` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new values `phantom_cash` and `usdt` on enums `ChargePaymentMethodDetailsCrypto.TokenCurrency`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.TokenCurrency`, and `PaymentRecordPaymentMethodDetailsCrypto.TokenCurrency`
* Add support for `Location` and `Reader` on `ChargePaymentMethodDetailsKlarna`, `PaymentAttemptRecordPaymentMethodDetailsKlarna`, and `PaymentRecordPaymentMethodDetailsKlarna`
* Add support for `Mandate` on `ChargePaymentMethodDetailsPix`, `PaymentAttemptRecordPaymentMethodDetailsPix`, and `PaymentRecordPaymentMethodDetailsPix`
* Add support for `ManagedPayments` on `CheckoutSessionParams`, `CheckoutSession`, `PaymentIntent`, `PaymentLinkParams`, `PaymentLink`, `SetupIntent`, and `Subscription`
* Add support for `MandateOptions` on `CheckoutSessionPaymentMethodOptionsPixParams`, `CheckoutSessionPaymentMethodOptionsPix`, `PaymentIntentConfirmPaymentMethodOptionsPixParams`, `PaymentIntentPaymentMethodOptionsPixParams`, and `PaymentIntentPaymentMethodOptionsPix`
* Change type of `CheckoutSessionPaymentMethodOptionsPixParams.SetupFutureUsage`, `PaymentIntentConfirmPaymentMethodOptionsPixParams.SetupFutureUsage`, and `PaymentIntentPaymentMethodOptionsPixParams.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
* Add support for new values `fo_vat`, `gi_tin`, `it_cf`, and `py_ruc` on enums `CheckoutSessionCustomerDetailsTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, `TaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
* ⚠️ Change type of `CheckoutSessionPaymentMethodOptionsPix.SetupFutureUsage` and `PaymentIntentPaymentMethodOptionsPix.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
* Add support for new value `sunbit` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `Pix` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `MandatePaymentMethodDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for `Upi` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for new values `pix` and `upi` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `CardPresence` on `IssuingAuthorization`
* Add support for `AllowedCardPresences` and `BlockedCardPresences` on `IssuingCardSpendingControlsParams`, `IssuingCardSpendingControls`, `IssuingCardholderSpendingControlsParams`, and `IssuingCardholderSpendingControls`
* Add support for new value `fulfillment_error` on enum `IssuingCard.CancellationReason`
* Add support for new value `fulfillment_error` on enum `IssuingCard.ReplacementReason`
* Add support for `Amount` and `Currency` on `MandateMultiUse`
* Add support for `AmountToConfirm` on `PaymentIntentConfirmParams`
* Add support for new value `sunbit` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `KlarnaDisplayQRCode` on `PaymentIntentNextAction`
* Add support for new value `sunbit` on enum `PaymentLink.PaymentMethodTypes`
* Add support for new values `low`, `not_assessed`, and `unknown` on enum `RadarPaymentEvaluationSignalsFraudulentPayment.RiskLevel`
* Add support for new value `account` on enum `RadarValueList.ItemType`
* Add support for `MOTO` on `SetupAttemptPaymentMethodDetailsCard`
* Add support for `PixDisplayQRCode` on `SetupIntentNextAction`
* Add support for error codes `action_blocked` and `approval_required` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
