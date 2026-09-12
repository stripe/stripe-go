---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2374
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.1.0-alpha.1
---

* Add support for new resources `GiftCardOperation`, `GiftCard`, and `TaxFund`
* Add support for `Get` method on resource `GiftCardOperation`
* Add support for `Activate`, `Cashout`, `CheckBalance`, `Get`, `New`, `Reload`, and `VoidOperation` methods on resource `GiftCard`
* Add support for `Get` and `List` methods on resource `TaxFund`
* Add support for `UpdateCryptoRefundAddress` method on resource `PaymentIntent`
* Add support for `PerformanceLocationDetails` on `TaxCalculationLineItemParams`, `TaxCalculationLineItem`, and `TaxTransactionLineItem`
* ⚠️ Remove support for `MoneyServices` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, and `PaymentIntentCapturePaymentDetailsParams`
* Add support for `FRMealVoucher` on `ChargePaymentMethodDetailsCardBenefits`
* Add support for `Multicapture` on `ChargePaymentMethodDetailsCardPresent`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsCardPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, and `PaymentRecordPaymentMethodDetailsCardPresent`
* Add support for `Pix` on `CheckoutSessionCurrentAttemptPaymentMethodDetails`
* Add support for new value `jaywan` on enum `CheckoutSessionCurrentAttemptPaymentMethodDetailsCard.Brand`
* Add support for `ProvisionalCredit` on `IssuingDisputeParams` and `IssuingDispute`
* Add support for `Reason` on `PaymentAttemptRecordReportCanceledParams` and `PaymentRecordReportPaymentAttemptCanceledParams`
* Add support for `FiservValuelink`, `Givex`, and `Svs` on `PaymentAttemptRecordProcessorDetails` and `PaymentRecordProcessorDetails`
* ⚠️ Change type of `PaymentAttemptRecordProcessorDetails.Type` and `PaymentRecordProcessorDetails.Type` from `literal('custom')` to `enum('custom'|'fiserv_valuelink'|'givex'|'svs')`
* Add support for `CaptureBy` and `CaptureDelay` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardPresent`, and `PaymentIntentPaymentMethodOptionsCard`
* ⚠️ Remove support for `LiquidAsset` on `PaymentIntentConfirmPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`
* Add support for `RequestMulticapture` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, and `PaymentIntentPaymentMethodOptionsCardPresent`
* Add support for `IgnoreApplicationFee`, `IgnoreTransferData`, and `RequestPartialAuthorization` on `PaymentIntentConfirmPaymentMethodOptionsGiftCardParams` and `PaymentIntentPaymentMethodOptionsGiftCardParams`
* Add support for `LatestPaymentAttemptRecord` and `PaymentRecord` on `PaymentIntent`
* ⚠️ Remove support for `Reauthorization` and `ReauthorizeBefore` on `PaymentIntentAdvancedFeatureDetails`
* Add support for `RefundAddress` on `PaymentIntentNextActionCryptoDisplayDetailsDepositAddressesBase`, `PaymentIntentNextActionCryptoDisplayDetailsDepositAddressesSolana`, and `PaymentIntentNextActionCryptoDisplayDetailsDepositAddressesTempo`
* Add support for `Location` on `PaymentIntentPaymentDetails` and `SetupIntentSetupDetails`
* Add support for new value `transaction_verification` on enum `PaymentIntentPaymentMethodOptionsCrypto.Mode`
* Add support for `Data` on `RadarAccountEvaluationLoginInitiatedClientDeviceMetadataDetailsParams`, `RadarAccountEvaluationRegistrationInitiatedClientDeviceMetadataDetailsParams`, and `RadarCustomerEvaluationEvaluationContextClientDetailsParams`
* Add support for new value `promotion` on enum `V2CommerceProductCatalogImport.FeedType`
* ⚠️ Change type of `V2CoreFeeBatchAdjustments.TaxAdjustment` from `amount` to `an object`
* ⚠️ Change type of `V2CoreFeeBatch.Amount`, `V2CoreFeeBatchCollectionRecord.Amount`, `V2CoreFeeBatchCollectionRecordTax.Amount`, `V2CoreFeeBatchTax.Amount`, `V2CoreFeeEntry.Amount`, and `V2CoreFeeEntryTax.Amount` from `amount` to `an object`
* Add support for new value `tax_fund` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* Add support for `TaxFund` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
* Add support for new value `tax_fund` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
* Add support for error code `default_us_bank_account_cannot_be_archived` on `CannotProceedError`
