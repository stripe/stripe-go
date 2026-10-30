---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2478
semver_level: major
is_stripe_api_change: true
released_in_version: 87.2.0-alpha.1
---

* Add support for new resources `V2BillingContractTermQuantityChange`, `V2BillingContractTerm`, `V2CoreAccountSession`, `V2MoneyManagementFundingIntent`, `V2MoneyManagementInboundFunding`, `V2MoneyManagementRewardProgramEnrollment`, `V2MoneyManagementRewardProgramRule`, `V2MoneyManagementRewardProgram`, and `V2RiskForm`
* ⚠️ Remove support for resource `V2RiskInquiry`
* Add support for `Get` and `Update` methods on resource `V2RiskForm`
* Add support for `Activate`, `Deactivate`, `Get`, `List`, and `New` methods on resource `V2MoneyManagementRewardProgramEnrollment`
* Add support for `Activate`, `Deactivate`, `Get`, `List`, `New`, and `Update` methods on resources `V2MoneyManagementRewardProgramRule` and `V2MoneyManagementRewardProgram`
* Add support for `Get` and `List` methods on resource `V2MoneyManagementInboundFunding`
* Add support for `Cancel`, `Confirm`, `Get`, `List`, `New`, and `Update` methods on resource `V2MoneyManagementFundingIntent`
* Add support for `New` method on resource `V2CoreAccountSession`
* Add support for `ListContractTerms` and `RetrieveContractTerm` methods on resource `V2BillingContractTerm`
* ⚠️ Remove support for `Get`, `List`, and `Update` methods on resource `V2RiskInquiry`
* ⚠️ Remove support for `FxQuote` method on resource `V2MoneyManagementPayoutIntent`
* Add support for `Reject` and `Unreject` methods on resource `V2CoreAccount`
* Add support for `BillingGroup` on `BillingAlertNotification`, `BillingAlertSpendThresholdFiltersParams`, `BillingAlertSpendThresholdFilters`, `BillingAlertTriggeredGroupBy`, and `BillingBillingAlertAlertNotificationListParams`
* Add support for `TossPay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Mandate` on `ChargePaymentMethodDetailsPaypay`, `PaymentAttemptRecordPaymentMethodDetailsPaypay`, `PaymentRecordPaymentMethodDetailsPaypay`, `SubscriptionPaymentSettingsParams`, and `SubscriptionPaymentSettings`
* Add support for `AutomaticPassThroughFee` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `Paypay` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for `AdjustableQuantity` on `CheckoutSessionItemSubscriptionItem`
* Add support for `MarketingAttributions` on `CustomerParams` and `Customer`
* Add support for `MerchantOfRecord` on `PaymentIntentConfirmPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentDetailsMoneyServicesAccountFunding`
* ⚠️ Change type of `RadarPaymentEvaluationSignals.EarlyFraudWarning` from `Payment Evaluation Signal` to `nullable(Payment Evaluation Signal)`
* Add support for `PaymentMethodData` and `PaymentMethodOptions` on `RefundParams`
* Add support for `DebitBehavior` on `SubscriptionPaymentSettingsPaymentMethodOptionsSepaDebitParams` and `SubscriptionPaymentSettingsPaymentMethodOptionsSepaDebit`
* Add support for `Terms` on `V2BillingContractCancelParams`, `V2BillingContractParams`, and `V2BillingContract`
* ⚠️ Add support for new value `user_affinity` on enum `V2CoreAccountEvaluation.EvaluationsTriggered`
* ⚠️ Remove support for value `user_merchant_fit` from enum `V2CoreAccountEvaluation.EvaluationsTriggered`
* Add support for `Rejection` on `V2CoreAccount`
* Add support for `AlipayPayments`, `ScalapayPayments`, and `UpiPayments` on `V2CoreAccountConfigurationMerchantCapabilitiesParams` and `V2CoreAccountConfigurationMerchantCapabilities`
* Add support for `InboundFundings` on `V2CoreAccountConfigurationMoneyManagerCapabilitiesParams` and `V2CoreAccountConfigurationMoneyManagerCapabilities`
* ⚠️ Remove support for `Instant` and `SEPACredit` on `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsParams` and `V2CoreAccountConfigurationRecipientCapabilitiesBankAccounts`
* Add support for `Compression` on `V2DataQueryRunResultOptionsParams`, `V2DataQueryRunResultOptions`, `V2DataReportRunResultOptionsParams`, and `V2DataReportRunResultOptions`
* ⚠️ Remove support for `CompressFile` on `V2DataQueryRunResultOptionsParams`, `V2DataQueryRunResultOptions`, `V2DataReportRunResultOptionsParams`, and `V2DataReportRunResultOptions`
* Add support for `QueryErrors` on `V2DataQueryRunStatusDetails` and `V2DataReportRunStatusDetails`
* Add support for new values `execution_failed` and `resource_limit_exceeded` on enums `V2DataQueryRunStatusDetails.Code` and `V2DataReportRunStatusDetails.Code`
* Add support for `Format` on `V2DataReportRun`
* Add support for `RefreshedAt` on `V2DataReport`
* Add support for `SchemaName` on `V2DataSchemaColumnForeignKeysFrom` and `V2DataSchemaColumnForeignKeysTo`
* ⚠️ Remove support for `Authentication` and `SessionFingerprint` on `V2IamActivityLogDetailsUserAccess`
* Add support for `UserAgent` on `V2IamActivityLogDetailsUserAccessDashboardClient`
* ⚠️ Remove support for `BrowserVersion`, `Browser`, `DeviceType`, and `Os` on `V2IamActivityLogDetailsUserAccessDashboardClient`
* ⚠️ Remove support for `Level` on `V2IamActivityLogDetailsUserAccessRisk`
* Add support for `Reference` on `V2MoneyManagementFinancialAddressBankAccountIban`
* Add support for new value `sepa` on enums `V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccount.PreferredNetworks`, `V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccount.PreferredNetworks`, `V2MoneyManagementOutboundTransferToPayoutMethodOptionsBankAccount.PreferredNetworks`, and `V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccount.PreferredNetworks`
* Add support for `BankAccount`, `Card`, `CryptoWallet`, `Status`, and `Type` on `V2MoneyManagementOutboundPaymentTrackingDetails`
* Add support for `RecipientVerification` on `V2MoneyManagementPayoutIntent`
* Add support for `FieldsToReview` on `V2MoneyManagementPayoutIntentNextActionConfirm`
* ⚠️ Add support for new values `payout_method_invalid` and `recipient_capability_not_active` on enum `V2MoneyManagementPayoutIntentNextActionHandleFailure.FailureReason`
* ⚠️ Remove support for value `currency_required` from enum `V2MoneyManagementPayoutIntentNextActionHandleFailure.FailureReason`
* ⚠️ Remove support for `CaBankAccount`, `EUBankAccount`, and `MXBankAccount` on `V2MoneyManagementReceivedCreditBankTransfer`
* ⚠️ Remove support for `Network` on `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountAba`, `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountBreB`, `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountClabe`, `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountCpa`, `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountIban`, `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountNip`, `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountPix`, and `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCode`
* Add support for `PayoutMethod`, `Recipient`, and `VerificationType` on `V2MoneyManagementRecipientVerification`
* ⚠️ Change type of `V2MoneyManagementRecipientVerification.ConsumedBy` from `string` to `an object`
* Add support for `MatchedBusinessType` and `ProvidedBusinessType` on `V2MoneyManagementRecipientVerificationMatchResultDetails`
* ⚠️ Remove support for value `requires_capture` from enum `V2PaymentsOffSessionPayment.Status`
* Add support for `Configuration` on `V2ProvisioningResource`
* Add support for `CustomMarketingAction` and `TriggeringSignal` on `V2SignalsAccountActivityParams` and `V2SignalsAccountActivity`
* Add support for `MfaChallenges` on `V2SignalsAccountActivityLoginDecisionParams` and `V2SignalsAccountActivityLoginDecision`
* Add support for new value `custom_marketing_action` on enum `V2SignalsAccountActivity.Type`
* Add support for `UserAffinity` on `V2SignalsAccountEvaluationEvaluatedSignals` and `V2SignalsAccountSignal`
* ⚠️ Remove support for `UserMerchantFit` on `V2SignalsAccountEvaluationEvaluatedSignals` and `V2SignalsAccountSignal`
* ⚠️ Add support for new value `user_affinity` on enum `V2SignalsAccountEvaluation.PendingSignals`
* ⚠️ Remove support for value `user_merchant_fit` from enum `V2SignalsAccountEvaluation.PendingSignals`
* ⚠️ Add support for new value `user_affinity` on enum `V2SignalsAccountEvaluation.RequestedSignals`
* ⚠️ Remove support for value `user_merchant_fit` from enum `V2SignalsAccountEvaluation.RequestedSignals`
* Add support for `AdditionalDetails` on `V2SignalsAccountSignalCustomerRisk`
* ⚠️ Add support for new value `user_affinity` on enum `V2SignalsAccountSignal.Type`
* ⚠️ Remove support for value `user_merchant_fit` from enum `V2SignalsAccountSignal.Type`
* Add support for `Subscriptions` on `V2TaxIntegrationConfigurationParams` and `V2TaxIntegrationConfiguration`
* Add support for `Include` on `V2ProvisioningResourceParams`
* ⚠️ Remove support for `Amount` and `From` on `V2MoneyManagementPayoutIntentParams`
* Add support for `Parameters` on `V2DataQueryRunParams`
* Add support for `Actions` on `V2BillingContractParams`
* Add support for new values `alipay_payments`, `scalapay_payments`, and `upi_payments` on enum `EventsV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for new value `inbound_fundings.paper_checks` on enum `EventsV2CoreAccountIncludingConfigurationMoneyManagerCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for `Severity` on `EventsV2CoreHealthApiErrorFiringEvent`, `EventsV2CoreHealthApiErrorResolvedEvent`, `EventsV2CoreHealthApiLatencyFiringEvent`, `EventsV2CoreHealthApiLatencyResolvedEvent`, `EventsV2CoreHealthAuthorizationRateDropFiringEvent`, `EventsV2CoreHealthAuthorizationRateDropResolvedEvent`, `EventsV2CoreHealthElementsErrorFiringEvent`, `EventsV2CoreHealthElementsErrorResolvedEvent`, `EventsV2CoreHealthFraudRateIncreasedEvent`, `EventsV2CoreHealthInvoiceCountDroppedFiringEvent`, `EventsV2CoreHealthInvoiceCountDroppedResolvedEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestErrorsFiringEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestErrorsResolvedEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutFiringEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEvent`, `EventsV2CoreHealthMeterEventSummariesDelayedFiringEvent`, `EventsV2CoreHealthMeterEventSummariesDelayedResolvedEvent`, `EventsV2CoreHealthMetronomeNotificationLatencyFiringEvent`, `EventsV2CoreHealthMetronomeNotificationLatencyResolvedEvent`, `EventsV2CoreHealthPaymentMethodErrorFiringEvent`, `EventsV2CoreHealthPaymentMethodErrorResolvedEvent`, `EventsV2CoreHealthSepaDebitDelayedFiringEvent`, `EventsV2CoreHealthSepaDebitDelayedResolvedEvent`, `EventsV2CoreHealthTrafficVolumeDropFiringEvent`, `EventsV2CoreHealthTrafficVolumeDropResolvedEvent`, `EventsV2CoreHealthWebhookLatencyFiringEvent`, and `EventsV2CoreHealthWebhookLatencyResolvedEvent`
* Add support for event notifications `V2BillingContractTermCreatedEvent`, `V2BillingContractTermDeletedEvent`, and `V2BillingContractTermUpdatedEvent` with related object `V2BillingContractTerm`
* Add support for event notifications `V2MoneyManagementFundingIntentCanceledEvent`, `V2MoneyManagementFundingIntentCreatedEvent`, `V2MoneyManagementFundingIntentProcessingEvent`, `V2MoneyManagementFundingIntentRequiresActionEvent`, `V2MoneyManagementFundingIntentRequiresConfirmationEvent`, `V2MoneyManagementFundingIntentSucceededEvent`, `V2MoneyManagementFundingIntentUpdatedEvent`, and `V2MoneyManagementFundingIntentValidatingEvent` with related object `V2MoneyManagementFundingIntent`
* Add support for event notifications `V2MoneyManagementInboundFundingAvailableEvent`, `V2MoneyManagementInboundFundingCheckFailedEvent`, `V2MoneyManagementInboundFundingCheckProcessingEvent`, `V2MoneyManagementInboundFundingCheckQueuedEvent`, `V2MoneyManagementInboundFundingCheckReturnedEvent`, `V2MoneyManagementInboundFundingCheckSucceededEvent`, `V2MoneyManagementInboundFundingCreatedEvent`, and `V2MoneyManagementInboundFundingUpdatedEvent` with related object `V2MoneyManagementInboundFunding`
* Add support for event notifications `V2MoneyManagementRewardProgramRewardEarnedEvent` and `V2MoneyManagementRewardProgramRewardReversedEvent` with related object `V2MoneyManagementRewardProgram`
* ⚠️ Remove support for event notifications `V2PaymentsOffSessionPaymentAuthorizationAttemptFailedEvent` and `V2PaymentsOffSessionPaymentAuthorizationAttemptStartedEvent` with related object `V2PaymentsOffSessionPayment`
* Add support for error codes `check_funding_intent_confirmation_method_cannot_be_updated`, `check_funding_intent_details_cannot_be_updated`, `check_funding_intent_too_large`, `funding_intent_cannot_be_canceled`, and `funding_intent_cannot_be_confirmed` on `CannotProceedError`
* Add support for error code `funding_intent_settlement_timing_not_available_for_account_type` on `FeatureNotEnabledError`
