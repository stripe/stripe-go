---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2348
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.2.0-alpha.2
---

* Add support for new resource `V2DataAnalyticsMetricQueryResult`
* Add support for `Get`, `New`, and `Revoke` methods on resource `SharedPaymentIssuedToken`
* Add support for `New` method on resource `V2DataAnalyticsMetricQueryResult`
* Add support for `BalanceReport` and `PayoutReconciliationReport` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for `AppDistribution` and `SunbitPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for new values `fee_credit_funding`, `inbound_transfer_reversal`, and `inbound_transfer` on enum `BalanceTransaction.Type`
* Add support for `Sunbit` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new values `phantom_cash` and `usdt` on enums `ChargePaymentMethodDetailsCrypto.TokenCurrency`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.TokenCurrency`, and `PaymentRecordPaymentMethodDetailsCrypto.TokenCurrency`
* Add support for `Last4` on `ChargePaymentMethodDetailsGiftCard`, `PaymentAttemptRecordPaymentMethodDetailsGiftCard`, and `PaymentRecordPaymentMethodDetailsGiftCard`
* Add support for `Location` and `Reader` on `ChargePaymentMethodDetailsKlarna`, `PaymentAttemptRecordPaymentMethodDetailsKlarna`, and `PaymentRecordPaymentMethodDetailsKlarna`
* Add support for `BLIK` on `CheckoutSessionPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for new values `fo_vat`, `gi_tin`, `it_cf`, and `py_ruc` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `CheckoutSessionCustomerDetailsTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, `QuotePreviewInvoiceCustomerTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, `TaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
* Add support for `SharedPaymentGrantedToken` on `ConfirmationTokenPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new value `sunbit` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* ⚠️ Change type of `CreditNoteLineItemTaxesTaxRateDetails.TaxRate`, `CreditNoteTotalTaxesTaxRateDetails.TaxRate`, `InvoiceLineItemTaxesTaxRateDetails.TaxRate`, `InvoiceTotalTaxesTaxRateDetails.TaxRate`, and `QuotePreviewInvoiceTotalTaxesTaxRateDetails.TaxRate` from `string` to `expandable($TaxRate)`
* Add support for `BuyerConsents` on `DelegatedCheckoutRequestedSessionConfirmParams`
* Add support for `Consents` on `DelegatedCheckoutRequestedSessionBuyerConsentsMarketing`
* Add support for new value `blik` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `PaymentFacilitatorID` and `SubMerchantID` on `IssuingAuthorizationMerchantDataParams`, `TestHelpersIssuingTransactionCreateForceCaptureMerchantDataParams`, and `TestHelpersIssuingTransactionCreateUnlinkedRefundMerchantDataParams`
* Add support for `CardPresence` on `IssuingAuthorization`
* Add support for `AllowedCardPresences` and `BlockedCardPresences` on `IssuingCardSpendingControlsParams`, `IssuingCardSpendingControls`, `IssuingCardholderSpendingControlsParams`, and `IssuingCardholderSpendingControls`
* Add support for new value `fulfillment_error` on enum `IssuingCard.CancellationReason`
* Add support for new value `fulfillment_error` on enum `IssuingCard.ReplacementReason`
* ⚠️ Change type of `PaymentAttemptRecordPaymentMethodDetailsGiftCard.Balance` and `PaymentRecordPaymentMethodDetailsGiftCard.Balance` from `PaymentFlowsPrivatePaymentMethodsGiftCardDeprecatedDetailsResourceBalanceAmount` to `nullable(PaymentsPrimitivesPaymentRecordsResourcePaymentMethodGiftCardDetailsResourceBalance)`
* Add support for `AmountToConfirm` on `PaymentIntentConfirmParams`
* Add support for new value `sunbit` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `KlarnaDisplayQRCode` on `PaymentIntentNextAction`
* Add support for new value `sunbit` on enum `PaymentLink.PaymentMethodTypes`
* Add support for `ValidationErrors` on `PrivacyRedactionJob`
* Add support for `TaxDetails` on `Product`
* Add support for new values `low`, `not_assessed`, and `unknown` on enum `RadarPaymentEvaluationSignalsFraudulentPayment.RiskLevel`
* Add support for new value `account` on enum `RadarValueList.ItemType`
* Add support for `MOTO` on `SetupAttemptPaymentMethodDetailsCard`
* Add support for `AdmissionsTax`, `AttendanceTax`, `EntertainmentTax`, `GrossReceiptsTax`, `HospitalityTax`, `LuxuryTax`, `ResortTax`, and `TourismTax` on `TaxRegistrationCountryOptionsUsParams`
* Add support for `Purpose` on `TreasuryOutboundPaymentParams` and `TreasuryOutboundPayment`
* Add support for `CryptoWallet` on `V2MoneyManagementFinancialAddressCredentials`
* Add support for `MXBankAccount` on `V2MoneyManagementFinancialAddressCredentials` and `V2MoneyManagementReceivedCreditBankTransfer`
* Add support for new values `crypto_wallet` and `mx_bank_account` on enum `V2MoneyManagementFinancialAddressCredentials.Type`
* Add support for `CryptoWalletTransfer` on `V2MoneyManagementReceivedCredit`
* Add support for `EUBankAccount` on `V2MoneyManagementReceivedCreditBankTransfer`
* Add support for new values `crypto_wallet`, `eu_bank_account`, and `mx_bank_account` on enum `V2MoneyManagementReceivedCreditBankTransfer.OriginType`
* Add support for new value `crypto_wallet_transfer` on enum `V2MoneyManagementReceivedCredit.Type`
* Add support for `CryptoProperties` and `SettlementCurrency` on `V2MoneyManagementFinancialAddressParams`
* Add support for event notifications `V2CoreApprovalRequestCreatedEvent` and `V2CoreApprovalRequestExpiredEvent` with related object `V2CoreApprovalRequest`
* Add support for event notification `V2ExtendExtensionRunFailedEvent`
* Add support for error codes `action_blocked` and `approval_required` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `QuotePreviewInvoiceLastFinalizationError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
