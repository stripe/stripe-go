---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2398
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.3.0-alpha.1
---

* Add support for new resources `V2MoneyManagementReceivedDebitMandate`, `V2RiskInquiry`, `V2SignalsAccountActivity`, and `V2SignalsAccountEvaluation`
* Add support for `Get` and `New` methods on resource `V2SignalsAccountEvaluation`
* Add support for `Del`, `Get`, and `New` methods on resource `V2SignalsAccountActivity`
* Add support for `Get`, `List`, and `Update` methods on resource `V2RiskInquiry`
* Add support for `Cancel`, `Get`, and `List` methods on resource `V2MoneyManagementReceivedDebitMandate`
* Add support for `RateCards` on `BillingCreditBalanceSummaryFilterApplicabilityScopeParams`, `BillingCreditGrantApplicabilityConfigScopeParams`, and `BillingCreditGrantApplicabilityConfigScope`
* ⚠️ Change type of `ConfirmationTokenPaymentMethodPreviewGiftCard.Brand`, `GiftCard.Brand`, `GiftCardParams.Brand`, `PaymentMethodGiftCard.Brand`, `TerminalReaderActivateGiftCardParams.Brand`, `TerminalReaderCashoutGiftCardParams.Brand`, `TerminalReaderCheckGiftCardBalanceParams.Brand`, and `TerminalReaderReloadGiftCardParams.Brand` from `enum('fiserv_valuelink'|'givex'|'svs')` to `literal('svs')`
* Add support for new value `tempo` on enum `CryptoCustomerConsumerWallet.Network`
* Add support for new value `tempo` on enum `CryptoOnrampSessionTransactionDetails.DestinationNetwork`
* Add support for new value `tempo` on enum `CryptoOnrampSessionTransactionDetails.DestinationNetworks`
* Add support for `Tempo` on `CryptoOnrampSessionTransactionDetailsWalletAddresses`
* Add support for new value `try_again_later` on enum `GiftCardOperation.FailureCode`
* Add support for `Healthcare` on `IssuingAuthorization`
* Add support for `ProductCode` on `IssuingCardParams` and `IssuingCard`
* Add support for `ProductGraduationState` on `IssuingCard`
* Add support for `CVC` and `Number` on `RadarPaymentEvaluationPaymentDetailsPaymentMethodDetailsCardParams`
* Add support for `Card` on `RadarPaymentEvaluationPaymentDetailsPaymentMethodDetails`
* ⚠️ Change type of `TerminalReaderCollectPaymentMethodCollectConfigParams.GiftCardBrand` and `TerminalReaderProcessPaymentIntentProcessConfigParams.GiftCardBrand` from `enum('fiserv_valuelink'|'givex'|'svs')` to `literal('svs')`
* Add support for `GiftCard` on `TestHelpersTerminalReaderPresentPaymentMethodParams`
* Add support for `AmountDue` and `CustomerBalanceApplied` on `V2BillingIntentAmountDetails`
* ⚠️ Change type of `V2CoreAccountEvaluation.EvaluationsTriggered` from `literal('fraudulent_website')` to `enum('fraudulent_website'|'user_account_sharing'|'user_multi_accounting')`
* Add support for `GrossSettlement` on `V2CoreAccountConfigurationMerchantParams` and `V2CoreAccountConfigurationMerchant`
* ⚠️ Change type of `V2MoneyManagementDebitDisputeBankTransfer.Network` from `literal('ach')` to `enum('ach'|'bacs')`
* Add support for new values `beneficiary_unrecognized`, `mandate_canceled_by_stripe`, `mandate_canceled`, `no_advance_notice`, `originator_requested`, and `signature_invalid` on enum `V2MoneyManagementDebitDisputeBankTransfer.Reason`
* ⚠️ Remove support for `ManagedBy` on `V2MoneyManagementFinancialAccount`
* Add support for `PayoutIntent` on `V2MoneyManagementOutboundPayment`
* Add support for `SettlesAt` on `V2MoneyManagementReceivedDebit`
* Add support for `GBBankAccount` on `V2MoneyManagementReceivedDebitBankTransfer`
* ⚠️ Change type of `V2MoneyManagementReceivedDebitBankTransfer.OriginType` from `literal('us_bank_account')` to `enum('gb_bank_account'|'us_bank_account')`
* ⚠️ Change type of `V2MoneyManagementReceivedDebitBankTransfer.PaymentMethodType` from `literal('us_bank_account')` to `enum('gb_bank_account'|'us_bank_account')`
* Add support for new value `scheduled` on enum `V2MoneyManagementReceivedDebit.Status`
* Add support for new value `no_mandate` on enum `V2MoneyManagementReceivedDebitStatusDetailsFailed.Reason`
* Add support for `TargetDate` on `V2PaymentsOffSessionPaymentParams` and `V2PaymentsOffSessionPayment`
* Add support for `AccountEvaluation`, `FraudulentWebsite`, `PaymentDelinquencyExposure`, `UserAccountSharing`, and `UserMultiAccounting` on `V2SignalsAccountSignal`
* Add support for new values `fraudulent_website`, `user_account_sharing`, and `user_multi_accounting` on enum `V2SignalsAccountSignal.Type`
* Change type of `V2MoneyManagementFinancialAddressDebitSimulationDebitParams.Network` from `literal('ach')` to `enum('ach'|'bacs')`
* Add support for `ReceivedDebitMandate` on `V2MoneyManagementReceivedDebitListParams`
* ⚠️ Remove support for `PayoutIntent` on `V2MoneyManagementOutboundPaymentParams`
* Change type of `V2CoreAccountEvaluationParams.Signals` from `literal('fraudulent_website')` to `enum('fraudulent_website'|'user_account_sharing'|'user_multi_accounting')`
* ⚠️ Remove support for `ID` on `EventsV2SignalsAccountSignalFraudulentMerchantReadyEvent`
* Add support for event notifications `V2MoneyManagementReceivedDebitCreatedEvent` and `V2MoneyManagementReceivedDebitScheduledEvent` with related object `V2MoneyManagementReceivedDebit`
* Add support for event notifications `V2MoneyManagementReceivedDebitMandateCanceledEvent`, `V2MoneyManagementReceivedDebitMandateCreatedEvent`, `V2MoneyManagementReceivedDebitMandateExpiredEvent`, `V2MoneyManagementReceivedDebitMandatePendingCancellationEvent`, and `V2MoneyManagementReceivedDebitMandateUpdatedEvent` with related object `V2MoneyManagementReceivedDebitMandate`
* Add support for event notification `V2SignalsAccountEvaluationCompleteEvent` with related object `V2SignalsAccountEvaluation`
* Add support for event notifications `V2SignalsAccountSignalFraudulentWebsiteReadyEvent` and `V2SignalsAccountSignalPaymentDelinquencyExposureReadyEvent` with related object `V2SignalsAccountSignal`
