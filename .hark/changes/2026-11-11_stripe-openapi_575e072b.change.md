---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2491
semver_level: major
is_stripe_api_change: true
---

* Add support for new resources `TaxCertificate`, `TaxFormCorrection`, `V2CoreAccountSession`, `V2ExtendObjectDefinitionVersion`, `V2ExtendObjectDefinition`, and `V2MoneyManagementSpendingMandate`
* ⚠️ Remove support for resource `V2CoreAccountSessionResource`
* Add support for `Del`, `Get`, `List`, and `New` methods on resource `TaxCertificate`
* Add support for `New` method on resource `TaxForm`
* Add support for `Cancel`, `Get`, and `List` methods on resource `V2MoneyManagementSpendingMandate`
* Add support for `Get` and `List` methods on resources `V2ExtendObjectDefinitionVersion` and `V2ExtendObjectDefinition`
* Add support for `Bills`, `FinancialAccountRewards`, `FinancialAccountsTransactions`, `FinancialAccounts`, `NetworkCostPassthroughReport`, and `RecipientsList` on `AccountSessionComponentsParams`
* Add support for `BusinessProfileManagement` on `AccountSessionComponentsParams` and `V2CoreAccountSessionResourceComponentsParams`
* Add support for `ExportTaxTransactions`, `ProductTaxCodeSelector`, and `TaxThresholdMonitoring` on `AccountSessionComponents`
* Add support for `AeonPay`, `AuPay`, `Dbarai`, and `Merpay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `FundingSources` on `ChargePaymentMethodDetailsCardWalletLink` and `ChargePaymentMethodDetailsLink`
* Add support for `Fleet` on `ChargePaymentMethodDetailsCardPresent`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsCardPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, and `PaymentRecordPaymentMethodDetailsCardPresent`
* Add support for `CardholderMessage` on `ChargePaymentMethodDetailsCardPresentReceipt`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresentReceipt`, `PaymentAttemptRecordPaymentMethodDetailsCardPresentReceipt`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresentReceipt`, and `PaymentRecordPaymentMethodDetailsCardPresentReceipt`
* Add support for `Fingerprint` on `ChargePaymentMethodDetailsPaypay`, `PaymentAttemptRecordPaymentMethodDetailsPaypay`, `PaymentRecordPaymentMethodDetailsPaypay`, and `SetupAttemptPaymentMethodDetailsPaypay`
* Add support for `CustomPaymentMethodOptions` and `PricingToken` on `CheckoutSessionParams`
* Add support for `StripeBalance` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for `SetupFutureUsage` on `CheckoutSessionPaymentMethodOptionsPaypayParams` and `CheckoutSessionPaymentMethodOptionsPaypay`
* Add support for new value `buyer_choice` on enum `CheckoutSessionPaymentMethodOptionsAcssDebitMandateOptions.TransactionType`
* Add support for new value `instant` on enums `CheckoutSessionPaymentMethodOptionsBacsDebit.VerificationMethod`, `InvoicePaymentSettingsPaymentMethodOptionsBacsDebit.VerificationMethod`, `PaymentIntentPaymentMethodOptionsBacsDebit.VerificationMethod`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBacsDebit.VerificationMethod`, `SetupIntentPaymentMethodOptionsBacsDebit.VerificationMethod`, and `SubscriptionPaymentSettingsPaymentMethodOptionsBacsDebit.VerificationMethod`
* Add support for new values `aeon_pay`, `au_pay`, `dbarai`, and `merpay` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `MarketingAttributions` on `CustomerParams`
* Add support for `ReasonDetails` and `Reason` on `CustomerCustomerTaxExemptionParams` and `CustomerTaxExemption`
* Add support for `TaxCodeGroup` on `CustomerCustomerTaxExemptionUsParams` and `CustomerTaxExemptionUs`
* Add support for `FraudulentSelfie` on `IdentityVerificationReportSignals`
* Add support for `TargetDate` on `InvoicePaymentSettingsPaymentMethodOptionsSepaDebitParams` and `InvoicePaymentSettingsPaymentMethodOptionsSepaDebit`
* Add support for `TossPay` on `MandatePaymentMethodDetails`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, and `SetupAttemptPaymentMethodDetails`
* Add support for `PaymentEvaluations` and `PaymentMethodDetails` on `PaymentAttemptRecordReportAuthenticatedParams`
* Add support for `ThreeDSecure` on `PaymentAttemptRecordReportAuthorizedPaymentMethodDetailsCardParams`, `PaymentAttemptRecordReportGuaranteedPaymentMethodDetailsCardParams`, `PaymentRecordReportPaymentAttemptGuaranteedPaymentMethodDetailsCardParams`, `PaymentRecordReportPaymentAttemptPaymentMethodDetailsCardParams`, and `PaymentRecordReportPaymentPaymentMethodDetailsCardParams`
* Change type of `PaymentAttemptRecordReportGuaranteedPaymentMethodDetailsParams.Type` and `PaymentRecordReportPaymentAttemptGuaranteedPaymentMethodDetailsParams.Type` from `literal('card')` to `enum('card'|'us_bank_account')`
* Add support for `DisputeDetails` on `PaymentAttemptRecord` and `PaymentRecord`
* Add support for `EvCharging` and `FleetData` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCardPresent`, `PaymentIntentAmountDetailsLineItemsPaymentMethodOptionsCardPresentParams`, `PaymentIntentCaptureAmountDetailsLineItemsPaymentMethodOptionsCardPresentParams`, `PaymentIntentConfirmAmountDetailsLineItemsPaymentMethodOptionsCardPresentParams`, `PaymentIntentDecrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsCardPresentParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsCardPresentParams`
* Add support for `FinancialConnections` on `PaymentIntentConfirmPaymentMethodOptionsBacsDebitParams`, `PaymentIntentPaymentMethodOptionsBacsDebitParams`, `PaymentIntentPaymentMethodOptionsBacsDebit`, `SetupIntentConfirmPaymentMethodOptionsBacsDebitParams`, `SetupIntentPaymentMethodOptionsBacsDebitParams`, and `SetupIntentPaymentMethodOptionsBacsDebit`
* Add support for new values `aeon_pay`, `au_pay`, `dbarai`, and `merpay` on enums `PaymentIntent.AllowedPaymentMethodTypes` and `SetupIntent.AllowedPaymentMethodTypes`
* Add support for new values `aeon_pay`, `au_pay`, `dbarai`, and `merpay` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `WeChatPayRedirectToMiniProgram` on `PaymentIntentNextAction`
* Add support for `PaymentDetails` on `PaymentIntentPaymentMethodOptionsCardPresent` and `PaymentIntentPaymentMethodOptionsCard`
* ⚠️ Remove support for value `on_session` from enum `PaymentIntentPaymentMethodOptionsPaypay.SetupFutureUsage`
* Add support for new value `tamara` on enum `PaymentLink.PaymentMethodTypes`
* Add support for `RakutenPay` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
* Add support for `SEPADebit` on `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsParams` and `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetails`
* Add support for new value `sepa_debit` on enum `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetails.MoneyMovementType`
* Add support for new value `continue_with_manual_review` on enum `RadarPaymentEvaluation.RecommendedAction`
* Add support for `CardTesting` on `RadarPaymentEvaluationSignals`
* Add support for `PaymentMethod` on `RefundParams`
* Add support for `PaymentRecord` on `Refund`
* Add support for `Bizum` and `Upi` on `RefundDestinationDetails`
* Add support for `TaxInvoiceMonth` on `ReportingReportRunParameters`
* Add support for `RadarOptions` on `SetupIntentConfirmParams`, `SetupIntentParams`, and `SetupIntent`
* Add support for `ReferralFeeDetails` on `SharedPaymentGrantedToken`, `SharedPaymentIssuedTokenParams`, and `SharedPaymentIssuedToken`
* Add support for `BillingGroup` on `SubscriptionListParams`, `SubscriptionParams`, and `Subscription`
* Add support for `DeviceHealthDetails` and `DeviceHealth` on `TerminalReader`
* Add support for `NetworkDetails` on `TreasuryInboundTransferFailureDetails`, `TreasuryOutboundPaymentReturnedDetails`, and `TreasuryOutboundTransferReturnedDetails`
* Add support for `Reference` on `V2MoneyManagementFinancialAddressBankAccountIban`
* Add support for new values `check_attachment_too_large`, `check_expired`, and `check_undeliverable` on enum `V2MoneyManagementOutboundPaymentStatusDetailsFailed.Reason`
* Add support for `CheckSend` on `V2MoneyManagementOutboundPaymentToPayoutMethodOptionsParams`, `V2MoneyManagementOutboundPaymentToPayoutMethodOptions`, `V2MoneyManagementOutboundPaymentTrackingDetails`, `V2MoneyManagementOutboundSetupIntentPayoutMethodDataParams`, `V2MoneyManagementPayoutIntentToPayoutMethodOptionsParams`, `V2MoneyManagementPayoutIntentToPayoutMethodOptions`, and `V2MoneyManagementPayoutMethod`
* Add support for new values `check_attachment_too_large`, `check_expired`, and `check_undeliverable` on enum `V2MoneyManagementPayoutIntentNextActionHandleFailure.FailureReason`
* Add support for new value `check_send` on enum `V2MoneyManagementPayoutMethod.Type`
* Add support for `SpendingMandate` on `V2MoneyManagementReceivedDebitBankTransfer` and `V2MoneyManagementReceivedDebitListParams`
* Add support for `AdditionalDetails` on `V2SignalsAccountSignalCustomerRisk`
* Add support for event notification `V2ExtendObjectDefinitionCreatedEvent` with related object `V2ExtendObjectDefinition`
* Add support for event notification `V2ExtendObjectDefinitionVersionCreatedEvent` with related object `V2ExtendObjectDefinitionVersion`
* Add support for event notifications `V2MoneyManagementSpendingMandateCanceledEvent`, `V2MoneyManagementSpendingMandateCreatedEvent`, `V2MoneyManagementSpendingMandateExpiredEvent`, `V2MoneyManagementSpendingMandatePendingCancellationEvent`, and `V2MoneyManagementSpendingMandateUpdatedEvent` with related object `V2MoneyManagementSpendingMandate`
* ⚠️ Add error type `MerchantNotGatedError` on `V2CoreVaultNetworkToken#CreateFromCredential`
