<!--
THIS IS A GENERATED FILE. Any changes you make to it directly will be blown away.
Instead, edit a corresponding `.change.md` file and run `hark build`.
-->

# Changelog

> This changelog only covers the **private preview** releases. Each release builds on the most recent GA release; see those notes in [the GA changelog](https://github.com/stripe/stripe-go/blob/master/CHANGELOG.md).

## 86.5.0-alpha.3 - 2026-09-09
* [#2428](https://github.com/stripe/stripe-go/pull/2428) Update generated code for private-preview
  * Add support for `CustomerTaxExemption` on `TaxCalculationLineItemTaxBreakdown`, `TaxCalculationShippingCostTaxBreakdown`, and `TaxTransactionShippingCostTaxBreakdown`
  * Add support for new value `data_share_only` on enums `ChargePaymentMethodDetailsCardThreeDSecure.Result`, `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure.Result`, `PaymentRecordPaymentMethodDetailsCardThreeDSecure.Result`, and `SetupAttemptPaymentMethodDetailsCardThreeDSecure.Result`
  * Add support for `BackdateStartDate` on `CheckoutSessionItemSubscriptionParams` and `CheckoutSessionItemSubscription`
  * Add support for `Signals` on `IdentityVerificationReport`
  * Add support for `NetworkResponseCode` on `IssuingAuthorizationRequestHistory`
  * Add support for `UnitCostPrecision` on `PaymentIntentAmountDetailsLineItem`, `PaymentIntentAmountDetailsLineItemsParams`, `PaymentIntentCaptureAmountDetailsLineItemsParams`, `PaymentIntentConfirmAmountDetailsLineItemsParams`, `PaymentIntentDecrementAuthorizationAmountDetailsLineItemsParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsParams`
  * Add support for `Active` on `ProductCatalogTrialOfferListParams`
  * Add support for new value `rtp` on enum `TreasuryFinancialAccountFinancialAddress.SupportedNetworks`
  * Add support for new value `blik_recurring_payments` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`

## 86.5.0-alpha.2 - 2026-09-02
* [#2427](https://github.com/stripe/stripe-go/pull/2427) Use an origin-relative path in the beta-header tests
* ⚠️ [#2420](https://github.com/stripe/stripe-go/pull/2420) Update generated code for private-preview
  * Add support for new resources `RadarBillingEvaluation`, `V2SignalsPaymentRetryEvaluation`, `V2SignalsPaymentRetrySignal`, and `V2TaxIntegrationConfiguration`
  * Add support for `New` method on resource `RadarBillingEvaluation`
  * Add support for `Deactivate`, `Get`, `List`, `New`, and `Update` methods on resource `BillingFeedbackOption`
  * Add support for `Get` and `Update` methods on resource `V2TaxIntegrationConfiguration`
  * Add support for `Get` method on resource `V2SignalsPaymentRetrySignal`
  * Add support for `Cancel`, `Get`, `New`, and `Update` methods on resource `V2SignalsPaymentRetryEvaluation`
  * Add support for `Disable` method on resource `V2MoneyManagementPayoutMethod`
  * Add support for `Update` method on resource `V2CoreApprovalRequest`
  * ⚠️ Remove support for `Execute` and `Submit` methods on resource `V2CoreApprovalRequest`
  * Add support for `DisableStripeUserAuthentication` on `AccountSessionComponentsPaymentMethodSettingsFeaturesParams`
  * Add support for `CapitalFinancingManualPayment` on `AccountSessionComponents`
  * Add support for `SequraPayments` on `AccountCapabilities`
  * Add support for `FeedbackOptions` on `BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonParams`
  * Add support for new value `fundbox_ca_financing` on enum `CapitalFinancingSummaryDetails.DisclaimerVariant`
  * Add support for `Sequra` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentPaymentMethodOptions`, and `PaymentRecordPaymentMethodDetails`
  * ⚠️ Remove support for value `data_share_only` from enums `ChargePaymentMethodDetailsCardThreeDSecure.Result`, `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure.Result`, `PaymentRecordPaymentMethodDetailsCardThreeDSecure.Result`, and `SetupAttemptPaymentMethodDetailsCardThreeDSecure.Result`
  * Add support for `FundingTypesBlocked` on `CheckoutSessionPaymentMethodOptionsCardRestrictionsParams`
  * Add support for `PaymentIntentData` on `CheckoutSessionParams`
  * ⚠️ Change type of `CheckoutSessionPaymentMethodOptionsBancontact.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
  * Add support for `Metadata` on `ConfirmationToken`, `V2SignalsAccountActivityParams`, and `V2SignalsAccountActivity`
  * Add support for new value `sequra` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
  * Add support for `ActiveEntitlements` and `CustomerPortal` on `CustomerSessionComponentsParams`
  * Add support for `AddressMatchConfidence` and `NameMatchConfidence` on `IdentityVerificationReportEmail` and `IdentityVerificationReportPhone`
  * Add support for `DomainCountry`, `EmailExistsConfidence`, `ObservedDomainTenureDays`, `ObservedEmailTenureDays`, and `PhoneMatchConfidence` on `IdentityVerificationReportEmail`
  * Add support for new values `email_address_mismatch`, `email_name_mismatch`, `email_ownership_unverified`, `email_phone_mismatch`, and `email_short_tenure` on enum `IdentityVerificationReportEmailError.Code`
  * Add support for `Carrier`, `LineType`, and `ObservedPhoneTenureDays` on `IdentityVerificationReportPhone`
  * Add support for new values `phone_address_mismatch`, `phone_invalid_line_type`, `phone_invalid`, `phone_name_mismatch`, `phone_ownership_unverified`, `phone_short_tenure`, and `phone_unsupported_country` on enum `IdentityVerificationReportPhoneError.Code`
  * Add support for new values `email_address_mismatch`, `email_name_mismatch`, `email_ownership_unverified`, `email_phone_mismatch`, `email_short_tenure`, `phone_address_mismatch`, `phone_invalid_line_type`, `phone_invalid`, `phone_name_mismatch`, `phone_ownership_unverified`, `phone_short_tenure`, and `phone_unsupported_country` on enum `IdentityVerificationSessionLastError.Code`
  * Add support for new value `truemoney` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
  * ⚠️ Remove support for `PaymentMethodTypes` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, and `SetupIntentParams`
  * Add support for `VerificationMethod` on `PaymentIntentConfirmPaymentMethodOptionsBacsDebitParams`, `PaymentIntentPaymentMethodOptionsBacsDebitParams`, `PaymentIntentPaymentMethodOptionsBacsDebit`, `SetupIntentConfirmPaymentMethodOptionsBacsDebitParams`, `SetupIntentPaymentMethodOptionsBacsDebitParams`, and `SetupIntentPaymentMethodOptionsBacsDebit`
  * Add support for new value `touch_n_go` on enums `PaymentIntent.AllowedPaymentMethodTypes` and `SetupIntent.AllowedPaymentMethodTypes`
  * Add support for new value `sequra` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
  * Add support for `ApplicationFeeAmount`, `ApplicationFeePercent`, `OnBehalfOf`, and `TransferData` on `PaymentLinkParams`
  * Add support for `Canceled` on `PaymentRecordReportPaymentAttemptParams` and `PaymentRecordReportPaymentParams`
  * ⚠️ Change type of `ProductCatalogTrialOffer.Price` from `$Price` to `deletable($Price)`
  * Add support for `Recurring` on `SharedPaymentGrantedTokenUsageLimitsParams`, `SharedPaymentGrantedTokenUsageLimits`, `SharedPaymentIssuedTokenUsageLimitsParams`, and `SharedPaymentIssuedTokenUsageLimits`
  * Add support for `FeedbackOption` on `SubscriptionCancelCancellationDetailsParams` and `SubscriptionCancellationDetailsParams`
  * Add support for `PricingToken` on `SubscriptionParams`
  * Add support for `Igic` on `TaxRegistrationCountryOptionsAtParams`, `TaxRegistrationCountryOptionsBeParams`, `TaxRegistrationCountryOptionsBgParams`, `TaxRegistrationCountryOptionsCyParams`, `TaxRegistrationCountryOptionsCzParams`, `TaxRegistrationCountryOptionsDeParams`, `TaxRegistrationCountryOptionsDkParams`, `TaxRegistrationCountryOptionsEeParams`, `TaxRegistrationCountryOptionsEsParams`, `TaxRegistrationCountryOptionsFiParams`, `TaxRegistrationCountryOptionsFrParams`, `TaxRegistrationCountryOptionsGrParams`, `TaxRegistrationCountryOptionsHrParams`, `TaxRegistrationCountryOptionsHuParams`, `TaxRegistrationCountryOptionsIeParams`, `TaxRegistrationCountryOptionsItParams`, `TaxRegistrationCountryOptionsLtParams`, `TaxRegistrationCountryOptionsLuParams`, `TaxRegistrationCountryOptionsLvParams`, `TaxRegistrationCountryOptionsMtParams`, `TaxRegistrationCountryOptionsNlParams`, `TaxRegistrationCountryOptionsPlParams`, `TaxRegistrationCountryOptionsPtParams`, `TaxRegistrationCountryOptionsRoParams`, `TaxRegistrationCountryOptionsSeParams`, `TaxRegistrationCountryOptionsSiParams`, and `TaxRegistrationCountryOptionsSkParams`
  * Add support for `OneTimeFees` on `V2BillingContractParams` and `V2BillingContract`
  * ⚠️ Remove support for `PaymentMethodCollection` on `V2CoreAccountConfigurationMerchantGrossSettlementParams` and `V2CoreAccountConfigurationMerchantGrossSettlement`
  * Add support for `PayoutMethods` on `V2CoreAccountDefaultsParams` and `V2CoreAccountDefaults`
  * Add support for `Reason` on `V2CoreApprovalRequest`
  * ⚠️ Remove support for `Description` on `V2CoreApprovalRequest`
  * Add support for `APIKey`, `Type`, and `User` on `V2CoreApprovalRequestRequestedBy` and `V2CoreApprovalRequestReviewReviewedBy`
  * ⚠️ Remove support for `ID` and `Name` on `V2CoreApprovalRequestRequestedBy` and `V2CoreApprovalRequestReviewReviewedBy`
  * Add support for `ApprovedAt` on `V2CoreApprovalRequestStatusTransitions`
  * ⚠️ Remove support for `RequiresExecutionAt` on `V2CoreApprovalRequestStatusTransitions`
  * Add support for `CryptoTransaction` on `V2CoreFeeBatchCollectionRecord`
  * Add support for new value `crypto_transaction` on enum `V2CoreFeeBatchCollectionRecord.Type`
  * Add support for `Restricted` on `V2CoreVaultGbBankAccount` and `V2CoreVaultUsBankAccount`
  * Add support for `Savings` on `V2MoneyManagementFinancialAccountParams` and `V2MoneyManagementFinancialAccount`
  * Add support for new value `savings` on enum `V2MoneyManagementFinancialAccount.Type`
  * Add support for `EnabledDeliverySchemes` on `V2MoneyManagementPayoutMethodBankAccount`
  * ⚠️ Remove support for `EnabledDeliveryOptions` on `V2MoneyManagementPayoutMethodBankAccount`
  * Add support for new value `disabled` on enum `V2MoneyManagementPayoutMethodUsageStatus.Payments`
  * Add support for new value `disabled` on enum `V2MoneyManagementPayoutMethodUsageStatus.Transfers`
  * Add support for `ToAccount` on `V2MoneyManagementReceivedDebitBalanceTransfer`
  * Add support for `AccountRestricted` and `AccountSuspended` on `V2SignalsAccountActivityParams` and `V2SignalsAccountActivity`
  * Add support for new values `account_restricted` and `account_suspended` on enum `V2SignalsAccountActivity.Type`
  * ⚠️ Remove support for value `not_assessed` from enums `V2SignalsAccountEvaluationEvaluatedSignalsFraudulentWebsite.RiskLevel`, `V2SignalsAccountEvaluationEvaluatedSignalsUserAccountSharing.RiskLevel`, `V2SignalsAccountEvaluationEvaluatedSignalsUserMultiAccounting.RiskLevel`, `V2SignalsAccountSignalFraudulentMerchant.RiskLevel`, `V2SignalsAccountSignalFraudulentWebsite.RiskLevel`, `V2SignalsAccountSignalMerchantDelinquency.RiskLevel`, `V2SignalsAccountSignalUserAccountSharing.RiskLevel`, and `V2SignalsAccountSignalUserMultiAccounting.RiskLevel`
  * Add support for `AdditionalDetails` on `V2SignalsAccountSignalFraudulentMerchant` and `V2SignalsAccountSignalMerchantDelinquency`
  * ⚠️ Remove support for `Indicators` on `V2SignalsAccountSignalFraudulentMerchant` and `V2SignalsAccountSignalMerchantDelinquency`
  * Add support for `Action`, `Created`, and `Status` on `V2CoreApprovalRequestListParams`
  * Add support for `OneTimeFeeActions` on `V2BillingContractParams`
  * Add support for event notifications `V2CoreHealthMetronomeNotificationLatencyFiringEvent`, `V2CoreHealthMetronomeNotificationLatencyResolvedEvent`, and `V2SignalsPaymentRetryEvaluationsRetryRecommendedEvent`
  * Add support for event notifications `V2MoneyManagementPayoutIntentCanceledEvent`, `V2MoneyManagementPayoutIntentCreatedEvent`, `V2MoneyManagementPayoutIntentPostedEvent`, `V2MoneyManagementPayoutIntentProcessingEvent`, and `V2MoneyManagementPayoutIntentRequiresActionEvent` with related object `V2MoneyManagementPayoutIntent`
  * Add support for error codes `authentication_failure`, `capability_not_active`, `expired_payment_method`, `incorrect_postal_code`, `invalid_canceled_subscription_fields`, and `payment_method_restricted` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `QuotePreviewInvoiceLastFinalizationError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, `StripeError`, and `TerminalReaderActionApiError`
  * Add support for error code `contract_number_already_exists` on `AlreadyExistsError`
  * Add support for error codes `default_payout_method_cannot_be_disabled`, `evaluation_not_monitoring`, `missing_payment_data_for_evaluation`, `one_time_fee_already_billed`, `payment_not_eligible`, and `webhook_endpoint_not_configured` on `CannotProceedError`

## 86.5.0-alpha.1 - 2026-08-26
This release changes the pinned API version to `2026-08-26.preview`.

* [#2409](https://github.com/stripe/stripe-go/pull/2409) Add non-verified manged handlers
* ⚠️ [#2414](https://github.com/stripe/stripe-go/pull/2414) Update generated code for private-preview
  * Add support for new resource `CustomerTaxExemption`
  * Add support for `Del`, `Get`, `List`, and `New` methods on resource `CustomerTaxExemption`
  * Add support for `Details` on `AccountFutureRequirementsErrors`, `AccountRequirementsErrors`, `BankAccountFutureRequirementsErrors`, `BankAccountRequirementsErrors`, `CapabilityFutureRequirementsError`, and `PersonFutureRequirementsError`
  * ⚠️ Remove support for `SequraPayments` on `AccountCapabilities`
  * Add support for `SubscriptionPause` on `BillingPortalSessionFlowDataParams`
  * ⚠️ Remove support for `Sequra` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentPaymentMethodOptions`, and `PaymentRecordPaymentMethodDetails`
  * Add support for `EnablementDetails` on `CheckoutSessionAutomaticTax`
  * ⚠️ Remove support for value `sequra` from enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
  * Add support for `Credit` on `FinancialConnectionsTransactionClassifications`
  * Change type of `FinancialConnectionsTransactionClassifications.MoneyMovement` from `nullable(BankConnectionsResourceTransactionResourceClassificationsLabels)` to `BankConnectionsResourceTransactionResourceClassificationsLabels`
  * Change type of `FinancialConnectionsTransactionClassifications.PersonalFinance` from `nullable(BankConnectionsResourceTransactionResourceClassificationsLabels)` to `BankConnectionsResourceTransactionResourceClassificationsLabels`
  * Add support for `UserConsent` on `IdentityVerificationSessionParams`
  * Add support for `CompanyDetails` on `InvoicePaymentSettingsPaymentMethodOptionsBillie`, `PaymentIntentConfirmPaymentMethodOptionsBillieParams`, `PaymentIntentPaymentMethodOptionsBillieParams`, `PaymentIntentPaymentMethodOptionsBillie`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBillie`, and `SubscriptionPaymentSettingsPaymentMethodOptionsBillie`
  * Add support for `Reference` on `InvoicePaymentSettingsPaymentMethodOptionsBillie`, `PaymentIntentConfirmPaymentMethodOptionsBillieParams`, `PaymentIntentPaymentMethodOptionsBillieParams`, `PaymentIntentPaymentMethodOptionsBillie`, and `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBillie`
  * Add support for `PosCondition` on `IssuingAuthorizationParams` and `IssuingAuthorization`
  * Add support for `CryptoWallet` on `IssuingCardParams` and `IssuingCard`
  * Add support for `PaymentEvaluations` and `PaymentMethodDetails` on `PaymentAttemptRecordReportAuthorizedParams`
  * Add support for `AadeData` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams` and `PaymentIntentPaymentMethodOptionsCardPresentParams`
  * ⚠️ Remove support for `CancelAtPeriodEnd` on `SubscriptionPendingUpdate`
  * Add support for `BLIKRecurringPayments` on `V2CoreAccountConfigurationMerchantCapabilitiesParams` and `V2CoreAccountConfigurationMerchantCapabilities`
  * Add support for `UserAccess` on `V2IamActivityLogDetails`
  * Add support for new value `user_access` on enum `V2IamActivityLogDetails.Type`
  * Add support for new value `user_access_started` on enum `V2IamActivityLog.Type`
  * Add support for new value `blik_recurring_payments` on enum `EventsV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent.UpdatedCapability`

## 86.4.0-alpha.2 - 2026-08-19
* ⚠️ [#2407](https://github.com/stripe/stripe-go/pull/2407) Update generated code for private-preview
  * Add support for new resources `BillingFeedbackOption` and `PaymentPlan`
  * ⚠️ Remove support for resource `BillingFeedbackOptions`
  * Add support for `Get`, `List`, `New`, and `Update` methods on resource `PaymentPlan`
  * Add support for `Update` method on resource `V2MoneyManagementTransaction`
  * Add support for `WeChatPayPayments` on `AccountSettingsParams` and `AccountSettings`
  * ⚠️ Change type of `BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReason.FeedbackOptions` from `$Billing.FeedbackOptions` to `$Billing.FeedbackOption`
  * Add support for `SubscriptionPause` on `BillingPortalSessionFlow`
  * Add support for new value `subscription_pause` on enum `BillingPortalSessionFlow.Type`
  * Add support for new value `usdt` on enum `CryptoOnrampSessionTransactionDetails.DestinationCurrencies`
  * Add support for new value `usdt` on enum `CryptoOnrampSessionTransactionDetails.DestinationCurrency`
  * Add support for `ActiveEntitlements` on `CustomerSessionComponents`
  * Add support for `SharedPaymentIssuedToken` on `DelegatedCheckoutRequestedSessionConfirmParams`
  * Add support for new values `payment_plan.created`, `payment_plan.installment_due`, `payment_plan.installment_paid`, `payment_plan.installment_will_be_due`, and `payment_plan.updated` on enum `Event.Type`
  * Add support for `ManagedPayments` on `InvoiceItemParams`, `InvoiceItem`, `InvoiceParams`, `Invoice`, and `QuotePreviewInvoice`
  * Add support for `PaymentPlan` on `Invoice`
  * Add support for `EstimatedFeeDetails` and `EstimatedFee` on `IssuingAuthorizationPendingRequestHoldAmountDetails` and `IssuingAuthorizationRequestHistoryHoldAmountDetails`
  * ⚠️ Remove support for `Cryptogram` on `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure` and `PaymentRecordPaymentMethodDetailsCardThreeDSecure`
  * ⚠️ Change type of `ProductCatalogTrialOfferEndBehaviorTransition.Price` from `$Price` to `deletable($Price)`
  * ⚠️ Change type of `SubscriptionCancellationDetails.FeedbackOption` from `$Billing.FeedbackOptions` to `$Billing.FeedbackOption`
  * Add support for `CancelAtPeriodEnd` on `SubscriptionPendingUpdate`
  * Add support for `Igic` on `TaxRegistrationCountryOptionsAt`, `TaxRegistrationCountryOptionsBe`, `TaxRegistrationCountryOptionsBg`, `TaxRegistrationCountryOptionsCy`, `TaxRegistrationCountryOptionsCz`, `TaxRegistrationCountryOptionsDe`, `TaxRegistrationCountryOptionsDk`, `TaxRegistrationCountryOptionsEe`, `TaxRegistrationCountryOptionsEs`, `TaxRegistrationCountryOptionsFi`, `TaxRegistrationCountryOptionsFr`, `TaxRegistrationCountryOptionsGr`, `TaxRegistrationCountryOptionsHr`, `TaxRegistrationCountryOptionsHu`, `TaxRegistrationCountryOptionsIe`, `TaxRegistrationCountryOptionsIt`, `TaxRegistrationCountryOptionsLt`, `TaxRegistrationCountryOptionsLu`, `TaxRegistrationCountryOptionsLv`, `TaxRegistrationCountryOptionsMt`, `TaxRegistrationCountryOptionsNl`, `TaxRegistrationCountryOptionsPl`, `TaxRegistrationCountryOptionsPt`, `TaxRegistrationCountryOptionsRo`, `TaxRegistrationCountryOptionsSe`, `TaxRegistrationCountryOptionsSi`, and `TaxRegistrationCountryOptionsSk`
  * Add support for `Metadata` on `V2BillingContractParams`, `V2BillingContractPricingLineActionUpdateParams`, `V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesData`, `V2BillingContractPricingOverrideActionAddParams`, `V2BillingContractPricingOverrideActionUpdateParams`, `V2BillingContractPricingOverrideParams`, `V2BillingContractPricingOverridesData`, and `V2MoneyManagementTransaction`
  * Add support for `TaxAmount` on `V2MoneyManagementOutboundPaymentQuoteEstimatedFee`
  * Add support for `PayoutMethodOptions` on `V2MoneyManagementOutboundPaymentQuoteToParams` and `V2MoneyManagementOutboundPaymentQuoteTo`
  * Change type of `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateParams.Metadata` from `string` to `emptyable(string)`
  * Add support for snapshot events `EventTypePaymentPlanCreated`, `EventTypePaymentPlanInstallmentDue`, `EventTypePaymentPlanInstallmentPaid`, `EventTypePaymentPlanInstallmentWillBeDue`, and `EventTypePaymentPlanUpdated` with resource `PaymentPlan`

## 86.4.0-alpha.1 - 2026-08-12
This release changes the pinned API version to `2026-08-12.preview`.

* ⚠️ [#2404](https://github.com/stripe/stripe-go/pull/2404) Update generated code for private-preview
  * Add support for new resource `V2TaxOperationsResolveAddressResult`
  * Add support for `ResolveAddress` method on resource `V2TaxOperationsResolveAddressResult`
  * Add support for `Confirm` and `FxQuote` methods on resource `V2MoneyManagementPayoutIntent`
  * ⚠️ Add support for new value `partner_disabled` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
  * ⚠️ Remove support for values `partner_disabled_dispute_rate`, `partner_disabled_responsibilities`, `partner_disabled_restricted_business`, and `partner_disabled_suspected_fraud` from enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
  * Add support for `CustomerUpdate` on `BillingPortalSessionFlow`
  * Add support for new value `customer_update` on enum `BillingPortalSessionFlow.Type`
  * Add support for `FundingSourceGroup` on `ChargePaymentMethodDetailsLink`
  * ⚠️ Remove support for `PricingGroup` on `ChargePaymentMethodDetailsLink`
  * Add support for new value `celo` on enum `CryptoCustomerConsumerWallet.Network`
  * Add support for new value `celo` on enum `CryptoOnrampSessionTransactionDetails.DestinationNetwork`
  * Add support for new value `celo` on enum `CryptoOnrampSessionTransactionDetails.DestinationNetworks`
  * Add support for `Celo` on `CryptoOnrampSessionTransactionDetailsWalletAddresses`
  * Add support for `CustomerPortal` on `CustomerSessionComponents`
  * Add support for `AppliedToInvoice` and `Type` on `CustomerCustomerBalanceTransactionParams`
  * Add support for `ClassificationState` and `EnrichmentState` on `FinancialConnectionsAccount`
  * Add support for `Country` on `FinancialConnectionsSessionFilters`
  * Add support for `Classifications` and `Enrichments` on `FinancialConnectionsTransaction`
  * Add support for `CustomerBalance` on `Invoice` and `QuotePreviewInvoice`
  * Add support for `Billie` on `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
  * Add support for new values `billie`, `paypay`, and `vipps` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
  * Add support for `HoldAmountDetails` and `HoldAmount` on `IssuingAuthorizationPendingRequest` and `IssuingAuthorizationRequestHistory`
  * Add support for new values `hu` and `ro` on enum `IssuingCardholder.PreferredLocales`
  * Add support for `NetworkDeclineCode` on `PaymentAttemptRecordReportFailedPaymentMethodDetailsCardParams` and `PaymentRecordReportPaymentAttemptFailedPaymentMethodDetailsCardParams`
  * Add support for `SetupFutureUsage` on `PaymentIntentPaymentMethodOptionsSequra`
  * Add support for `Status` on `QuotePreviewSubscriptionSchedulePauseSchedulePause`, `QuotePreviewSubscriptionSchedulePauseScheduleResume`, `SubscriptionSchedulePauseSchedulePause`, and `SubscriptionSchedulePauseScheduleResume`
  * Change type of `SubscriptionSchedulePauseSchedulesParams.Resume` from `pause_schedule_update_resume_params` to `emptyable(pause_schedule_update_resume_params)`
  * Add support for `Acquirer` on `EventsV2CoreHealthAuthorizationRateDropFiringEventImpactDimension`, `EventsV2CoreHealthAuthorizationRateDropResolvedEventImpactDimension`, `V2CoreHealthAlertAuthorizationRateDropDimension`, and `V2CoreHealthAlertHistoryEntryAuthorizationRateDropDimension`
  * ⚠️ Change type of `EventsV2CoreHealthAuthorizationRateDropFiringEventImpactDimension.Type`, `EventsV2CoreHealthAuthorizationRateDropResolvedEventImpactDimension.Type`, `V2CoreHealthAlertAuthorizationRateDropDimension.Type`, and `V2CoreHealthAlertHistoryEntryAuthorizationRateDropDimension.Type` from `literal('issuer')` to `enum('acquirer'|'issuer')`
  * Add support for `ConfirmationMethod` on `V2MoneyManagementPayoutIntentParams` and `V2MoneyManagementPayoutIntent`
  * Add support for `EstimatedFees` and `FxQuote` on `V2MoneyManagementPayoutIntent`
  * Add support for `Debited` on `V2MoneyManagementPayoutIntentFrom`
  * Add support for `Confirm` on `V2MoneyManagementPayoutIntentNextAction`
  * ⚠️ Change type of `V2MoneyManagementPayoutIntentNextAction.Type` from `literal('handle_failure')` to `enum('confirm'|'handle_failure')`
  * Add support for `Credited` on `V2MoneyManagementPayoutIntentTo`
  * Add support for event notification `V1BalanceSettingsUpdatedEvent` with related object `BalanceSettings`
  * Add support for event notification `V1BillingCreditBalanceTransactionCreatedEvent` with related object `BillingCreditBalanceTransaction`
  * Add support for event notifications `V1BillingCreditGrantCreatedEvent` and `V1BillingCreditGrantUpdatedEvent` with related object `BillingCreditGrant`
  * Add support for event notifications `V1BillingMeterCreatedEvent`, `V1BillingMeterDeactivatedEvent`, `V1BillingMeterReactivatedEvent`, and `V1BillingMeterUpdatedEvent` with related object `BillingMeter`
  * Add support for event notifications `V1FinancialConnectionsAccountAccountNumbersUpdatedEvent`, `V1FinancialConnectionsAccountExpectedDeactivationDateUpdatedEvent`, `V1FinancialConnectionsAccountSupportedPaymentMethodTypesUpdatedEvent`, `V1FinancialConnectionsAccountUpcomingAccountNumberExpiryEvent`, and `V1FinancialConnectionsAccountUpcomingDeactivationEvent` with related object `FinancialConnectionsAccount`
  * Add support for event notification `V1InvoicePaymentAttemptRequiredEvent` with related object `Invoice`
  * Add support for error type `FxQuoteNeedsRefreshError`

## 86.3.0-alpha.2 - 2026-08-05
* [#2402](https://github.com/stripe/stripe-go/pull/2402) Update generated code for private-preview
  * Add support for new resource `BillingFeedbackOptions`
  * Add support for `SequraPayments` on `AccountCapabilities`
  * Add support for `FeedbackOptions` on `BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReason`
  * Add support for `Sequra` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentPaymentMethodOptions`, and `PaymentRecordPaymentMethodDetails`
  * Add support for `RetrievalReferenceNumber` on `ChargePaymentMethodDetailsCardPresent`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsCardPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, and `PaymentRecordPaymentMethodDetailsCardPresent`
  * Add support for `PricingGroup` on `ChargePaymentMethodDetailsLink`
  * Add support for `TaxRates` on `CheckoutSessionShippingOptionParams`, `CheckoutSessionShippingOption`, and `CheckoutSessionShippingOptionsParams`
  * Add support for new value `daikin` on enums `CheckoutSessionAutomaticSurcharge.Provider` and `PaymentLinkAutomaticSurcharge.Provider`
  * Add support for `FundingTypesBlocked` on `CheckoutSessionPaymentMethodOptionsCardRestrictions`
  * Add support for new value `sequra` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
  * Add support for `Healthcare` on `IssuingAuthorizationParams`, `TestHelpersIssuingAuthorizationCapturePurchaseDetailsParams`, `TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsParams`, and `TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsParams`
  * Add support for `IsAnomalous` on `PaymentAttemptRecordReportGuaranteedParams`
  * Add support for new value `sequra` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
  * Add support for `AadeData` on `PaymentIntentPaymentMethodOptionsCardPresent`
  * Add support for `FeedbackOption` on `SubscriptionCancellationDetails`
  * Add support for `Application` on `V2PaymentsOffSessionPayment`
  * Add support for `Status` on `V2MoneyManagementFinancialAccountStatementListParams`

## 86.3.0-alpha.1 - 2026-07-29
This release changes the pinned API version to `2026-07-29.preview`.

* ⚠️ [#2398](https://github.com/stripe/stripe-go/pull/2398) Update generated code for private-preview
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

## 86.2.0-alpha.5 - 2026-07-22
* ⚠️ [#2396](https://github.com/stripe/stripe-go/pull/2396) Update generated code for private-preview
  * Add support for new resources `BillingAlertNotification` and `CryptoDepositAddress`
  * Add support for `Get`, `List`, and `New` methods on resource `CryptoDepositAddress`
  * Add support for `List` method on resource `BillingAlertNotification`
  * Add support for new values `partner_disabled_dispute_rate`, `partner_disabled_responsibilities`, `partner_disabled_restricted_business`, and `partner_disabled_suspected_fraud` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
  * Add support for new value `data_share_only` on enums `ChargePaymentMethodDetailsCardThreeDSecure.Result`, `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure.Result`, `PaymentRecordPaymentMethodDetailsCardThreeDSecure.Result`, and `SetupAttemptPaymentMethodDetailsCardThreeDSecure.Result`
  * Add support for `Vipps` on `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for new value `vipps` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
  * Add support for new value `sui` on enum `CryptoCustomerConsumerWallet.Network`
  * Add support for new value `sui` on enum `CryptoOnrampSessionTransactionDetails.DestinationNetwork`
  * Add support for new value `sui` on enum `CryptoOnrampSessionTransactionDetails.DestinationNetworks`
  * Add support for `Sui` on `CryptoOnrampSessionTransactionDetailsWalletAddresses`
  * Add support for `UseStripeSDK` on `DelegatedCheckoutRequestedSessionConfirmParams`
  * Add support for new value `mb_way` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
  * Add support for `EvCharging` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCard`, `PaymentIntentAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentCaptureAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentConfirmAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentDecrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsCardParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsCardParams`
  * ⚠️ Change type of `PaymentIntent.AllowedPaymentMethodTypes` from `string` to `enum`
  * Add support for new value `vipps` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
  * Add support for `TaxItems` on `PaymentIntentPaymentDetailsCarRentalDatumTotalTax`, `PaymentIntentPaymentDetailsFlightDatumTotalTax`, and `PaymentIntentPaymentDetailsLodgingDatumTotalTax`
  * ⚠️ Remove support for `Taxes` on `PaymentIntentPaymentDetailsCarRentalDatumTotalTax`, `PaymentIntentPaymentDetailsFlightDatumTotalTax`, and `PaymentIntentPaymentDetailsLodgingDatumTotalTax`
  * Add support for `Card` on `RadarPaymentEvaluationPaymentDetailsPaymentMethodDetailsParams`
  * ⚠️ Remove support for `ACSSDebit`, `AUBECSDebit`, `AfterpayClearpay`, `Alipay`, `Alma`, `AmazonPay`, `BACSDebit`, `BLIK`, `Bancontact`, `Billie`, `Bizum`, `Boleto`, `CardPresent`, `CashApp`, `Crypto`, `CustomerBalance`, `EPS`, `FPX`, `GiftCard`, `Giropay`, `Gopay`, `Grabpay`, `IDBankTransfer`, `IDEAL`, `InteracPresent`, `KakaoPay`, `Konbini`, `KrCard`, `MbWay`, `Mobilepay`, `Multibanco`, `NaverPay`, `NzBankAccount`, `OXXO`, `P24`, `PayByBank`, `PayNow`, `Payco`, `Paypal`, `Paypay`, `Payto`, `Pix`, `PromptPay`, `Qris`, `Rechnung`, `RevolutPay`, `SEPADebit`, `SamsungPay`, `Satispay`, `Scalapay`, `Shopeepay`, `Sofort`, `StripeBalance`, `Sunbit`, `Swish`, `TWINT`, `Tamara`, `USBankAccount`, `Upi`, `WeChatPay`, and `Zip` on `SharedPaymentGrantedTokenPaymentMethodDetails`
  * ⚠️ Add support for new value `shop_pay` on enum `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
  * ⚠️ Remove support for values `acss_debit`, `afterpay_clearpay`, `alipay`, `alma`, `amazon_pay`, `au_becs_debit`, `bacs_debit`, `bancontact`, `billie`, `bizum`, `blik`, `boleto`, `card_present`, `cashapp`, `crypto`, `custom`, `customer_balance`, `eps`, `fpx`, `gift_card`, `giropay`, `gopay`, `grabpay`, `id_bank_transfer`, `ideal`, `interac_present`, `kakao_pay`, `konbini`, `kr_card`, `mb_way`, `mobilepay`, `multibanco`, `naver_pay`, `nz_bank_account`, `oxxo`, `p24`, `pay_by_bank`, `payco`, `paynow`, `paypal`, `paypay`, `payto`, `pix`, `promptpay`, `qris`, `rechnung`, `revolut_pay`, `samsung_pay`, `satispay`, `scalapay`, `sepa_debit`, `shopeepay`, `sofort`, `stripe_balance`, `sunbit`, `swish`, `tamara`, `twint`, `upi`, `us_bank_account`, `wechat_pay`, and `zip` from enum `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
  * Add support for `SpendCard` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeParams` and `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripe`
  * Add support for new value `commercial.stripe.spend_card` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for new value `payment_delinquency_exposure` on enum `V2SignalsAccountSignal.Type`
  * Add support for new value `commercial.stripe.spend_card` on enum `EventsV2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEvent.UpdatedCapability`

## 86.2.0-alpha.4 - 2026-07-16
* ⚠️ [#2388](https://github.com/stripe/stripe-go/pull/2388) Update generated code for private-preview
  * ⚠️ Remove support for resource `FRMealVouchersOnboarding`
  * ⚠️ Remove support for `Get`, `List`, `New`, and `Update` methods on resource `FRMealVouchersOnboarding`
  * Add support for `New` method on resource `PaymentRecord`
  * Add support for new value `chaps` on enums `FundingInstructionsBankTransferFinancialAddress.SupportedNetworks` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddress.SupportedNetworks`
  * ⚠️ Remove support for `FinancialAccountsTransactions`, `FinancialAccounts`, and `RecipientsList` on `AccountSessionComponentsParams`
  * Add support for `SmartDisputesManagement` on `AccountSessionComponentsDisputesListFeatures`, `AccountSessionComponentsPaymentDetailsFeatures`, `AccountSessionComponentsPaymentDisputesFeatures`, and `AccountSessionComponentsPaymentsFeatures`
  * Add support for new value `ic_nif` on enums `CheckoutSessionCollectedInformationTaxId.Type`, `CheckoutSessionCustomerDetailsTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, `QuotePreviewInvoiceCustomerTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
  * Add support for new values `financial_connections.account.expected_deactivation_date_updated`, `financial_connections.account.supported_payment_method_types_updated`, `financial_connections.account.upcoming_deactivation`, `financial_connections.authorization.expected_deactivation_date_updated`, and `financial_connections.authorization.upcoming_deactivation` on enum `Event.Type`
  * Add support for `Mode` on `FinancialConnectionsSessionManualEntry`
  * Add support for new values `alipay` and `sequra` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
  * Add support for new value `stripe_internal_error` on enum `IssuingAuthorizationRequestHistory.Reason`
  * Add support for `BusinessName` on `IssuingCardShipping`
  * Add support for new value `correos` on enum `IssuingCardShipping.Carrier`
  * ⚠️ Change type of `IssuingTransactionNetworkData.TraceID` from `IssuingTransactionTraceId` to `nullable(IssuingTransactionTraceId)`
  * Add support for `PauseSchedules` on `QuotePreviewSubscriptionSchedule`, `SubscriptionScheduleParams`, and `SubscriptionSchedule`
  * Add support for `Trial` on `QuotePreviewSubscriptionSchedulePhase` and `SubscriptionSchedulePhase`
  * Add support for `PaymentRecord` on `RefundParams`
  * Add support for `RedirectToURL` on `SharedPaymentIssuedTokenNextAction`
  * ⚠️ Change type of `SharedPaymentIssuedTokenNextAction.Type` from `literal('use_stripe_sdk')` to `enum('redirect_to_url'|'use_stripe_sdk')`
  * Add support for snapshot events `EventTypeFinancialConnectionsAccountExpectedDeactivationDateUpdated`, `EventTypeFinancialConnectionsAccountSupportedPaymentMethodTypesUpdated`, and `EventTypeFinancialConnectionsAccountUpcomingDeactivation` with resource `FinancialConnectionsAccount`
  * Add support for snapshot events `EventTypeFinancialConnectionsAuthorizationExpectedDeactivationDateUpdated` and `EventTypeFinancialConnectionsAuthorizationUpcomingDeactivation` with resource `FinancialConnectionsAuthorization`

## 86.2.0-alpha.3 - 2026-07-08
This release changes the pinned API version to `2026-07-08.preview`.

* ⚠️ [#2383](https://github.com/stripe/stripe-go/pull/2383) Update generated code for private-preview
  * Add support for `ActivateGiftCard`, `CashoutGiftCard`, `CheckGiftCardBalance`, and `ReloadGiftCard` methods on resource `TerminalReader`
  * Add support for `AggregationPeriod` on `BillingAlertRecovered`
  * Add support for new values `mass_transit_parking_tax` and `parking_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationTaxBreakdownTaxRateDetails.TaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetails.TaxType`
  * Add support for `AdministrativeAddress` and `PrincipalPlaceOfBusiness` on `AccountCompany`
  * Add support for `AddressCollectionPrecision` on `CheckoutSessionAutomaticTax`
  * Add support for `TaxID` on `CheckoutSessionCollectedInformation`
  * ⚠️ Remove support for `TaxIDs` on `CheckoutSessionCollectedInformation`
  * Add support for `SetupFutureUsage` on `CheckoutSessionPaymentMethodOptionsPayco`, `CheckoutSessionPaymentMethodOptionsSamsungPay`, `PaymentIntentConfirmPaymentMethodOptionsPaypayParams`, `PaymentIntentPaymentMethodOptionsPayco`, `PaymentIntentPaymentMethodOptionsPaypayParams`, `PaymentIntentPaymentMethodOptionsPaypay`, and `PaymentIntentPaymentMethodOptionsSamsungPay`
  * Add support for new values `bnp_paribas`, `citibank`, and `mbsb_bank` on enums `ConfirmationTokenPaymentMethodPreviewFpx.Bank`, `PaymentAttemptRecordPaymentMethodDetailsFpx.Bank`, `PaymentRecordPaymentMethodDetailsFpx.Bank`, and `SharedPaymentGrantedTokenPaymentMethodDetailsFpx.Bank`
  * Add support for `Network` on `DisputePaymentMethodDetailsCard`
  * Add support for `RequirePaymentMethodSupport` on `FinancialConnectionsSessionFilters`
  * Add support for `NetworkData` on `IssuingAuthorizationRequestHistory`
  * Add support for `AcquiringInstitutionCountry`, `AcquiringInstitutionID`, `RetrievalReferenceNumber`, `RoutedNetwork`, and `TraceID` on `IssuingTransactionNetworkData`
  * Add support for `CustomFields`, `Description`, and `Footer` on `QuoteInvoiceSettings`, `QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettings`, `QuotePreviewSubscriptionSchedulePhaseInvoiceSettings`, `SubscriptionScheduleDefaultSettingsInvoiceSettings`, and `SubscriptionSchedulePhaseInvoiceSettings`
  * Add support for `Paypay` on `SetupAttemptPaymentMethodDetails`
  * Add support for `MassTransitParkingTax` and `ParkingTax` on `TaxRegistrationCountryOptionsUs`
  * Add support for new values `mass_transit_parking_tax` and `parking_tax` on enum `TaxRegistrationCountryOptionsUs.Type`
  * Add support for `GiftCardBrand` on `TerminalReaderCollectPaymentMethodCollectConfigParams` and `TerminalReaderProcessPaymentIntentProcessConfigParams`
  * Add support for `ActivateGiftCard`, `CashoutGiftCard`, `CheckGiftCardBalance`, `DeactivateGiftCard`, and `ReloadGiftCard` on `TerminalReaderAction`
  * Add support for new values `activate_gift_card`, `cashout_gift_card`, `check_gift_card_balance`, `deactivate_gift_card`, and `reload_gift_card` on enum `TerminalReaderAction.Type`
  * Add support for `StatusTransitions` on `V2BillingContract`
  * ⚠️ Remove support for `OneTimeFees` on `V2BillingContractParams` and `V2BillingContract`
  * ⚠️ Remove support for `StatusDetails` on `V2BillingContract`
  * Add support for `ID` and `Priority` on `V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesData`
  * ⚠️ Remove support for `PricingOverride` on `V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesData`
  * ⚠️ Remove support for `TieringMode` and `Tiers` on `V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideOverwritePriceParams`, `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddOverwritePriceParams`, `V2BillingContractPricingLinePricingPriceDetailsPricingOverrideOverwritePriceParams`, `V2BillingContractPricingLinesDataPricingPriceDetailsPricingOverridesDataOverwritePrice`, and `V2BillingContractPricingOverrideActionAddOverwritePriceParams`
  * Add support for `MultiplyPricing` on `V2BillingContractPricingOverrideActionAddParams`, `V2BillingContractPricingOverrideParams`, and `V2BillingContractPricingOverridesData`
  * ⚠️ Remove support for `Multiplier` on `V2BillingContractPricingOverrideActionAddParams`, `V2BillingContractPricingOverrideParams`, and `V2BillingContractPricingOverridesData`
  * ⚠️ Change type of `V2BillingContractPricingOverrideActionAddParams.Type`, `V2BillingContractPricingOverrideParams.Type`, and `V2BillingContractPricingOverridesData.Type` from `literal('multiplier')` to `literal('multiply_pricing')`
  * Add support for `RelatedNetworkObject` on `V2CoreAccountListParams` and `V2CoreAccount`
  * Add support for new value `network_business_profile_wallet` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * Add support for `NetworkBusinessProfileWallet` on `V2MoneyManagementPayoutMethod`
  * Add support for new value `network_business_profile_wallet` on enum `V2MoneyManagementPayoutMethod.Type`
  * Add support for `StripeNetworkTransfer` on `V2MoneyManagementReceivedCredit`
  * Add support for new value `stripe_network_transfer` on enum `V2MoneyManagementReceivedCredit.Type`
  * ⚠️ Change type of `V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideEndsAtParams.Type`, `V2BillingContractPricingLineEndsAtParams.Type`, `V2BillingContractPricingLinePricingPriceDetailsPricingOverrideEndsAtParams.Type`, and `V2BillingContractPricingOverrideEndsAtParams.Type` from `enum('contract_end'|'timestamp')` to `literal('timestamp')`
  * ⚠️ Change type of `V2BillingContractPricingLineActionAddPricingPriceDetailsPricingOverrideStartsAtParams.Type`, `V2BillingContractPricingLinePricingPriceDetailsPricingOverrideStartsAtParams.Type`, `V2BillingContractPricingLineStartsAtParams.Type`, and `V2BillingContractPricingOverrideStartsAtParams.Type` from `enum('contract_start'|'timestamp')` to `literal('timestamp')`
  * ⚠️ Change type of `V2BillingContractPricingLineActionAddEndsAtParams.Type`, `V2BillingContractPricingLineActionUpdateEndsAtParams.Type`, `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddEndsAtParams.Type`, `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateEndsAtParams.Type`, `V2BillingContractPricingOverrideActionAddEndsAtParams.Type`, and `V2BillingContractPricingOverrideActionUpdateEndsAtParams.Type` from `enum('billing_period_end'|'timestamp')` to `literal('timestamp')`
  * ⚠️ Change type of `V2BillingContractPricingLineActionAddStartsAtParams.Type`, `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionAddStartsAtParams.Type`, `V2BillingContractPricingLineActionUpdatePricingPriceDetailsPricingOverrideActionUpdateStartsAtParams.Type`, `V2BillingContractPricingLineActionUpdateStartsAtParams.Type`, `V2BillingContractPricingOverrideActionAddStartsAtParams.Type`, and `V2BillingContractPricingOverrideActionUpdateStartsAtParams.Type` from `enum('billing_period_start'|'timestamp')` to `literal('timestamp')`
  * Add support for event notifications `V2BillingContractActivatedEvent`, `V2BillingContractCanceledEvent`, `V2BillingContractCreatedEvent`, `V2BillingContractEndedEvent`, and `V2BillingContractUpdatedEvent` with related object `V2BillingContract`

## 86.2.0-alpha.2 - 2026-07-01
This release changes the pinned API version to `2026-07-01.preview`.

* ⚠️ [#2381](https://github.com/stripe/stripe-go/pull/2381) Update generated code for private-preview
  * Add support for new resources `CryptoCustomerConsumerWallet`, `CryptoCustomerPaymentToken`, `CryptoCustomer`, `CryptoOnrampSession`, and `CryptoOnrampTransactionLimits`
  * Add support for `Get` and `List` methods on resource `CryptoCustomer`
  * Add support for `Checkout`, `Get`, `List`, `New`, and `Quote` methods on resource `CryptoOnrampSession`
  * Add support for `Get` method on resource `CryptoOnrampTransactionLimits`
  * Add support for `ElectronicCommerceIndicator` on `ChargePaymentMethodDetailsCard`
  * Add support for `AmountReceived` and `AmountRequested` on `ChargePaymentMethodDetailsCrypto`, `PaymentAttemptRecordPaymentMethodDetailsCrypto`, and `PaymentRecordPaymentMethodDetailsCrypto`
  * Add support for `Fingerprint` on `ChargePaymentMethodDetailsGiftCard`, `PaymentAttemptRecordPaymentMethodDetailsGiftCard`, and `PaymentRecordPaymentMethodDetailsGiftCard`
  * Add support for `AddressCollectionPrecision` on `CheckoutSessionAutomaticTaxParams`
  * Add support for `Subscription` on `CheckoutSessionItem`
  * ⚠️  Remove support for `Deactivation` on `GiftCardOperation`
  * ⚠️  Remove support for value `deactivation` from enum `GiftCardOperation.Type`
  * Add support for `MerchantAmountExchangeRate` on `IssuingAuthorization` and `IssuingTransaction`
  * Add support for `DeviceID` on `IssuingAuthorizationTokenDetailsNetworkDataDevice` and `IssuingTokenNetworkDataDevice`
  * Add support for `Program` on `IssuingCard`
  * Add support for `PaymentMethodDetails` on `PaymentAttemptRecordReportFailedParams` and `PaymentRecordReportPaymentAttemptFailedParams`
  * Add support for `Reason` on `PaymentAttemptRecordReportRefundParams` and `PaymentRecordReportRefundParams`
  * Add support for `AmountReconciliation` on `PaymentIntentConfirmPaymentMethodOptionsCryptoParams`, `PaymentIntentPaymentMethodOptionsCryptoParams`, and `PaymentIntentPaymentMethodOptionsCrypto`
  * Add support for `ConnectPermissions` and `Permissions` on `V2IamApiKeyParams` and `V2IamApiKey`
  * Add support for `Credit` on `V2MoneyManagementFinancialAccount`
  * Add support for new value `credit` on enum `V2MoneyManagementFinancialAccount.Type`
  * Add support for new value `currency_required` on enum `V2MoneyManagementPayoutIntentNextActionHandleFailure.FailureReason`
  * Add support for new values `issuing_authorization`, `issuing_transaction`, and `platform_funded_credit_transaction` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
  * Add support for `Account`, `IssuingAuthorization`, `IssuingDispute`, and `IssuingTransaction` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
  * Add support for new values `issuing_authorization`, `issuing_dispute`, and `issuing_transaction` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
  * Change type of `V2MoneyManagementFinancialAccountParams.Type` from `literal('storage')` to `enum('credit'|'storage')`
  * Add support for `ExpiresAt` on `V2IamApiKeyParams`

## 86.2.0-alpha.1 - 2026-06-24
This release changes the pinned API version to `2026-06-24.preview`.

* ⚠️ [#2378](https://github.com/stripe/stripe-go/pull/2378) Update generated code for private-preview
  * Add support for new resources `V2BillingContractPricingLineQuantityChange`, `V2CoreHealthAlertHistoryEntry`, `V2CoreHealthAlert`, `V2MoneyManagementFinancialAddressDebitSimulation`, and `V2MoneyManagementPayoutIntent`
  * ⚠️ Remove support for resource `V2BillingContractLicensePricingQuantityChange`
  * Add support for `ReportOfferAcceptance` method on resource `IssuingCreditUnderwritingRecord`
  * Add support for `ProvisionalCredit` test helper method on resource `IssuingDispute`
  * Add support for `ReportEarlyFraudWarning` method on resource `PaymentAttemptRecord`
  * Add support for `Search` method on resource `PaymentRecord`
  * Add support for `Debit` method on resource `V2MoneyManagementFinancialAddressDebitSimulation`
  * Add support for `Cancel`, `Get`, `List`, `New`, and `Update` methods on resource `V2MoneyManagementPayoutIntent`
  * Add support for `Get` and `List` methods on resource `V2CoreHealthAlert`
  * Add support for `Del` method on resource `V2BillingContract`
  * ⚠️ Remove support for `PerformanceLocationDetails` on `TaxTransactionLineItem`
  * Add support for `FinancialAccountsTransactions`, `FinancialAccounts`, and `RecipientsList` on `AccountSessionComponentsParams` and `AccountSessionComponents`
  * Add support for `Location` and `Reader` on `ChargePaymentMethodDetailsGiftCard`, `GiftCardOperation`, `PaymentAttemptRecordPaymentMethodDetailsGiftCard`, and `PaymentRecordPaymentMethodDetailsGiftCard`
  * Add support for `Subscription` on `CheckoutSessionItemParams`
  * Add support for `Items` on `CheckoutSession`
  * Add support for `Brand` on `CheckoutSessionCurrentAttemptPaymentMethodDetailsCard`
  * Add support for `NetworkData` on `TestHelpersIssuingAuthorizationCaptureParams` and `TestHelpersIssuingTransactionCreateForceCaptureParams`
  * Add support for `EnrichedMerchantData` on `IssuingAuthorization`
  * Add support for `AvailableBalance` and `CurrentBalance` on `IssuingAuthorizationBalanceResponse`
  * ⚠️ Remove support for `Amount` on `IssuingAuthorizationBalanceResponse`
  * Add support for `DecisionDeadlineUpdatedAt` on `IssuingCreditUnderwritingRecord`
  * Add support for `AcquirerReferenceNumber` on `IssuingTransactionNetworkData`
  * Add support for `Tip` on `PaymentIntentAmountDetailsParams`, `PaymentIntentCaptureAmountDetailsParams`, `PaymentIntentConfirmAmountDetailsParams`, `PaymentIntentDecrementAuthorizationAmountDetailsParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsParams`
  * Add support for `BillingCycleAnchor` on `V2BillingContractParams` and `V2BillingContract`
  * ⚠️ Remove support for `ContractLineDetails`, `ContractValueDetails`, and `LicenseQuantities` on `V2BillingContract`
  * Add support for `BillSettingsDetails` on `V2BillingContractBillingSettingsParams` and `V2BillingContractBillingSettings`
  * Add support for `BillingProfileDetails` and `CollectionSettingsDetails` on `V2BillingContractBillingSettings`
  * ⚠️ Remove support for `ContractBillingDetails` on `V2BillingContractBillingSettingsParams` and `V2BillingContractBillingSettings`
  * ⚠️ Change type of `V2BillingContract.OneTimeFees` from `array(an object)` to `an object`
  * ⚠️ Change type of `V2BillingContract.PricingLines` from `array(an object)` to `an object`
  * ⚠️ Change type of `V2BillingContract.PricingOverrides` from `array(an object)` to `an object`
  * Add support for `Mode` on `V2CommerceProductCatalogImport`
  * Add support for new value `money_manager` on enums `EventsV2CoreAccountLinkReturnedEvent.Configurations`, `V2CoreAccountLinkUseCaseAccountOnboarding.Configurations`, and `V2CoreAccountLinkUseCaseAccountUpdate.Configurations`
  * ⚠️ Add support for new value `money_manager` on enum `V2CoreAccount.AppliedConfigurations`
  * ⚠️ Remove support for value `storer` from enum `V2CoreAccount.AppliedConfigurations`
  * Add support for `MoneyManager` on `V2CoreAccountConfigurationParams`, `V2CoreAccountConfiguration`, `V2CoreAccountIdentityAttestationsTermsOfServiceParams`, `V2CoreAccountIdentityAttestationsTermsOfService`, and `V2CoreAccountTokenIdentityAttestationsTermsOfServiceParams`
  * ⚠️ Remove support for `Storer` on `V2CoreAccountConfigurationParams`, `V2CoreAccountConfiguration`, `V2CoreAccountIdentityAttestationsTermsOfServiceParams`, `V2CoreAccountIdentityAttestationsTermsOfService`, and `V2CoreAccountTokenIdentityAttestationsTermsOfServiceParams`
  * Add support for `SunbitPayments` on `V2CoreAccountConfigurationMerchantCapabilitiesParams` and `V2CoreAccountConfigurationMerchantCapabilities`
  * Add support for `ACH`, `BECS`, `Eft`, `FPS`, `Fedwire`, `Npp`, `RTP`, `SEPACredit`, `SEPAInstant`, and `Swift` on `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsParams` and `V2CoreAccountConfigurationRecipientCapabilitiesBankAccounts`
  * Add support for new values `bank_accounts.ach`, `bank_accounts.becs`, `bank_accounts.eft`, `bank_accounts.fedwire`, `bank_accounts.fps`, `bank_accounts.npp`, `bank_accounts.rtp`, `bank_accounts.sepa_credit`, `bank_accounts.sepa_instant`, `bank_accounts.swift`, `business_storage.inbound.eur`, `business_storage.inbound.gbp`, `business_storage.inbound.usd`, `business_storage.outbound.eur`, `business_storage.outbound.gbp`, `business_storage.outbound.usd`, `consumer_storage.inbound.usd`, `consumer_storage.outbound.usd`, `received_credits.bank_accounts`, and `received_debits.bank_accounts` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for new value `money_manager` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Configuration` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Configuration`
  * Add support for `ConsumerMoneyManager` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams` and `V2CoreAccountIdentityAttestationsTermsOfService`
  * Add support for `CryptoMoneyManager` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams`, `V2CoreAccountIdentityAttestationsTermsOfService`, and `V2CoreAccountTokenIdentityAttestationsTermsOfServiceParams`
  * ⚠️ Remove support for `ConsumerStorer` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams` and `V2CoreAccountIdentityAttestationsTermsOfService`
  * ⚠️ Remove support for `CryptoStorer` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams`, `V2CoreAccountIdentityAttestationsTermsOfService`, and `V2CoreAccountTokenIdentityAttestationsTermsOfServiceParams`
  * ⚠️ Remove support for `MaximumRps` on `V2CoreBatchJobParams` and `V2CoreBatchJob`
  * Add support for `BIC` on `V2MoneyManagementFinancialAddressCredentialsUsBankAccount`
  * ⚠️ Remove support for `SwiftCode` on `V2MoneyManagementFinancialAddressCredentialsUsBankAccount`
  * Add support for `Attachment` on `V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheckParams` and `V2MoneyManagementOutboundPaymentDeliveryOptionsPaperCheck`
  * Add support for `Processing` on `V2MoneyManagementOutboundPaymentStatusDetails` and `V2MoneyManagementOutboundTransferStatusDetails`
  * Add support for new values `fx_rate_drift_exceeded_after_review` and `review_rejected` on enum `V2MoneyManagementOutboundPaymentStatusDetailsFailed.Reason`
  * Add support for `PayoutMethodOptions` on `V2MoneyManagementOutboundPaymentToParams`, `V2MoneyManagementOutboundPaymentTo`, `V2MoneyManagementOutboundTransferToParams`, and `V2MoneyManagementOutboundTransferTo`
  * Add support for new values `fx_rate_drift_exceeded_after_review` and `review_rejected` on enum `V2MoneyManagementOutboundTransferStatusDetailsFailed.Reason`
  * Add support for `AccountHolderName` on `V2MoneyManagementReceivedCreditBankTransferUsBankAccount`
  * Add support for `Returned` on `V2MoneyManagementReceivedDebitStatusDetails`
  * Add support for new value `capability_inactive` on enum `V2MoneyManagementReceivedDebitStatusDetailsFailed.Reason`
  * Add support for `ReturnedAt` on `V2MoneyManagementReceivedDebitStatusTransitions`
  * Add support for `PayoutIntent` on `V2MoneyManagementOutboundPaymentParams`
  * Add support for `Statuses` on `V2MoneyManagementFinancialAccountListParams`
  * ⚠️ Remove support for `Status` on `V2MoneyManagementFinancialAccountListParams`
  * Add support for `Include` on `V2BillingContractListParams`
  * ⚠️ Remove support for `ContractLines` on `V2BillingContractParams`
  * ⚠️ Remove support for `LicenseQuantityActions` on `V2BillingContractParams`
  * ⚠️ Add support for `BillingProfileDetails` and `CollectionSettingsDetails` on `V2BillingContractBillingSettingsParams`
  * ⚠️ Add support for `Amount`, `BillAt`, and `Product` on `V2BillingContractOneTimeFeeParams`
  * Add support for `LookupKey` on `V2BillingContractOneTimeFeeParams`
  * ⚠️ Remove support for `BillSchedule`, `BillableItemType`, and `ProductDetails` on `V2BillingContractOneTimeFeeParams`
  * Add support for `PricingOverrides` and `QuantityChanges` on `V2BillingContractPricingLineActionAddPricingPriceDetailsParams` and `V2BillingContractPricingLinePricingPriceDetailsParams`
  * ⚠️ Remove support for `Quantity` on `V2BillingContractPricingLineActionAddPricingPriceDetailsParams` and `V2BillingContractPricingLinePricingPriceDetailsParams`
  * ⚠️ Remove support for `OverwritePrice` on `V2BillingContractPricingOverrideParams`
  * Add support for `PricingLineIDs` and `PricingLineLookupKeys` on `V2BillingContractPricingOverrideActionAddMultiplierCriterionParams` and `V2BillingContractPricingOverrideMultiplierCriterionParams`
  * ⚠️ Remove support for `BillableItemIDs`, `BillableItemLookupKeys`, `BillableItemTypes`, `MetadataConditions`, and `RateCardIDs` on `V2BillingContractPricingOverrideActionAddMultiplierCriterionParams` and `V2BillingContractPricingOverrideMultiplierCriterionParams`
  * ⚠️ Change type of `V2BillingContractPricingOverrideActionAddParams.Type` and `V2BillingContractPricingOverrideParams.Type` from `enum('multiplier'|'overwrite_price')` to `literal('multiplier')`
  * Add support for `Pricing` on `V2BillingContractPricingLineActionUpdateParams`
  * ⚠️ Remove support for `Price` on `V2BillingContractPricingOverrideActionAddOverwritePriceParams`
  * Add support for `CancelPricingLines` and `ProrationBehavior` on `V2BillingContractCancelParams`
  * Add support for new value `sunbit_payments` on enum `EventsV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent.UpdatedCapability`
  * Add support for new values `bank_accounts.ach`, `bank_accounts.becs`, `bank_accounts.eft`, `bank_accounts.fedwire`, `bank_accounts.fps`, `bank_accounts.npp`, `bank_accounts.rtp`, `bank_accounts.sepa_credit`, `bank_accounts.sepa_instant`, and `bank_accounts.swift` on enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`
  * Add support for event notifications `V2CoreAccountIncludingConfigurationMoneyManagerCapabilityStatusUpdatedEvent` and `V2CoreAccountIncludingConfigurationMoneyManagerUpdatedEvent` with related object `V2CoreAccount`
  * Add support for event notifications `V2MoneyManagementDebitDisputeFailedEvent`, `V2MoneyManagementDebitDisputeSubmittedEvent`, and `V2MoneyManagementDebitDisputeSucceededEvent` with related object `V2MoneyManagementDebitDispute`
  * Add support for event notification `V2MoneyManagementOutboundPaymentUnderReviewEvent` with related object `V2MoneyManagementOutboundPayment`
  * Add support for event notification `V2MoneyManagementOutboundTransferUnderReviewEvent` with related object `V2MoneyManagementOutboundTransfer`
  * ⚠️ Remove support for event notifications `V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent` and `V2CoreAccountIncludingConfigurationStorerUpdatedEvent` with related object `V2CoreAccount`
  * Add support for error codes `us_bank_account_microdeposits_cannot_be_confirmed` and `us_bank_account_microdeposits_cannot_be_sent` on `ControlledByAlternateResourceError`
  * Add support for error code `payout_intent_not_cancelable` on `NotCancelableError`

## 86.1.0-alpha.2 - 2026-06-17
* ⚠️ [#2376](https://github.com/stripe/stripe-go/pull/2376) Update generated code for private-preview
  * Add support for `Get` method on resource `RadarCustomerEvaluation`
  * Add support for `DisableStripeUserAuthentication` on `AccountSessionComponentsBillsFeatures`
  * Add support for `Tamara` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentPaymentMethodDataParams`, and `SharedPaymentGrantedTokenPaymentMethodDetails`
  * Add support for `Status` on `ChargePaymentMethodDetailsCardAccountFunding`
  * ⚠️ Remove support for `ProcessedTransactionType` on `ChargePaymentMethodDetailsCardAccountFunding`
  * Add support for `Items` on `CheckoutSessionParams`
  * ⚠️ Remove support for `Brand` on `CheckoutSessionCurrentAttemptPaymentMethodDetailsCard`
  * ⚠️ Remove support for `First6` on `ConfirmationTokenPaymentMethodPreviewGiftCard`, `PaymentMethodGiftCard`, and `SharedPaymentGrantedTokenPaymentMethodDetailsGiftCard`
  * Add support for new value `tamara` on enums `ConfirmationTokenPaymentMethodPreview.Type`, `PaymentMethod.Type`, and `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
  * Add support for `Fingerprint` on `GiftCard`
  * Add support for `BLIK` on `MandatePaymentMethodDetails`
  * Add support for `BuyerID` on `OrderPaymentSettingsPaymentMethodOptionsWechatPayParams`, `OrderPaymentSettingsPaymentMethodOptionsWechatPay`, `PaymentIntentConfirmPaymentMethodOptionsWechatPayParams`, `PaymentIntentPaymentMethodOptionsWechatPayParams`, and `PaymentIntentPaymentMethodOptionsWechatPay`
  * Add support for new value `mini_program` on enums `OrderPaymentSettingsPaymentMethodOptionsWechatPay.Client` and `PaymentIntentPaymentMethodOptionsWechatPay.Client`
  * Add support for `PaymentMethodDetails` on `PaymentAttemptRecordReportGuaranteedParams` and `PaymentRecordReportPaymentAttemptGuaranteedParams`
  * Add support for `Failed` and `RefundGroup` on `PaymentAttemptRecordReportRefundParams` and `PaymentRecordReportRefundParams`
  * Change type of `PaymentAttemptRecordReportRefundParams.Outcome` and `PaymentRecordReportRefundParams.Outcome` from `literal('refunded')` to `enum('failed'|'refunded')`
  * Add support for `BeneficiaryDetails` on `PaymentIntentConfirmPaymentDetailsMoneyServicesParams`, `PaymentIntentPaymentDetailsMoneyServicesParams`, and `PaymentIntentPaymentDetailsMoneyServices`
  * ⚠️ Remove support for `BeneficiaryAccount` and `BeneficiaryDetails` on `PaymentIntentConfirmPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentDetailsMoneyServicesAccountFunding`
  * ⚠️ Remove support for `SenderAccount` on `PaymentIntentConfirmPaymentDetailsMoneyServicesAccountFundingParams` and `PaymentIntentPaymentDetailsMoneyServicesAccountFundingParams`
  * Add support for `GivenName` and `Surname` on `PaymentIntentConfirmPaymentDetailsMoneyServicesAccountFundingSenderDetailsParams`, `PaymentIntentPaymentDetailsMoneyServicesAccountFundingSenderDetailsParams`, and `PaymentIntentPaymentDetailsMoneyServicesAccountFundingSenderDetails`
  * ⚠️ Remove support for `Name` on `PaymentIntentConfirmPaymentDetailsMoneyServicesAccountFundingSenderDetailsParams`, `PaymentIntentPaymentDetailsMoneyServicesAccountFundingSenderDetailsParams`, and `PaymentIntentPaymentDetailsMoneyServicesAccountFundingSenderDetails`
  * Change type of `PaymentIntentConfirmPaymentMethodOptionsCardParams.CaptureMethod` and `PaymentIntentPaymentMethodOptionsCardParams.CaptureMethod` from `literal('manual')` to `enum('automatic_delayed'|'manual')`
  * ⚠️ Remove support for `Wallet` on `PaymentIntentConfirmPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`
  * Add support for `TransactionVerificationOptions` on `PaymentIntentConfirmPaymentMethodOptionsCryptoParams`, `PaymentIntentPaymentMethodOptionsCryptoParams`, and `PaymentIntentPaymentMethodOptionsCrypto`
  * Change type of `TestHelpersPaymentIntentSimulateCryptoDepositParams.TokenCurrency` from `literal('usdc')` to `enum('usdc'|'usdg'|'usdp')`
  * Add support for `ForcedCapture` on `PaymentIntentAdvancedFeatureDetails`
  * Add support for new value `tamara` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
  * Add support for `WeChatPayHandleAppRedirect` on `PaymentIntentNextAction` and `SetupIntentNextAction`
  * Add support for `Ethereum` and `Polygon` on `PaymentIntentNextActionCryptoDisplayDetailsDepositAddresses`
  * ⚠️ Change type of `PaymentIntentNextActionCryptoDisplayDetailsDepositAddressesBaseSupportedToken.TokenCurrency`, `PaymentIntentNextActionCryptoDisplayDetailsDepositAddressesSolanaSupportedToken.TokenCurrency`, and `PaymentIntentNextActionCryptoDisplayDetailsDepositAddressesTempoSupportedToken.TokenCurrency` from `literal('usdc')` to `enum('usdc'|'usdg'|'usdp')`
  * Add support for `BeneficiaryAccount` on `PaymentIntentPaymentDetailsMoneyServices`
  * ⚠️ Change type of `PaymentIntentPaymentMethodOptionsCard.CaptureMethod` from `literal('manual')` to `enum('automatic_delayed'|'manual')`
  * Add support for new value `automatic_delayed` on enum `PaymentIntentPaymentMethodOptionsCardPresent.CaptureMethod`
  * Add support for new values `ethereum` and `polygon` on enum `PaymentIntentPaymentMethodOptionsCryptoDepositOptions.Networks`
  * Change type of `PaymentLocationBusinessRegistrationParams.Siret` from `string` to `emptyable(string)`
  * Add support for `Card` on `PaymentRecordReportPaymentAttemptPaymentMethodDetailsParams` and `PaymentRecordReportPaymentPaymentMethodDetailsParams`
  * Change type of `PaymentRecordReportPaymentAttemptPaymentMethodDetailsParams.Type` and `PaymentRecordReportPaymentPaymentMethodDetailsParams.Type` from `literal('custom')` to `enum('card'|'custom')`
  * Add support for `ManagedPayments` on `Product`
  * Add support for `PaymentAttemptRecord` on `RefundListParams` and `RefundParams`
  * Add support for `PaymentRecord` on `RefundListParams`
  * Add support for `Protections` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankPrepaidCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialFifthThirdChargeCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCelticRevolvingCreditCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerCrossRiverBankPrepaidCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadDebitCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadDebitCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCardParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadPrepaidCard`, `V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTaxParams`, `V2CoreAccountConfigurationCustomerCapabilitiesAutomaticIndirectTax`, `V2CoreAccountConfigurationMerchantCapabilitiesAchDebitPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesAchDebitPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAcssDebitPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesAcssDebitPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAffirmPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesAffirmPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesAfterpayClearpayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAlmaPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesAlmaPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesAmazonPayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesAuBecsDebitPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesAuBecsDebitPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesBacsDebitPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesBacsDebitPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesBancontactPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesBancontactPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesBlikPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesBlikPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesBoletoPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesBoletoPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesCardPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesCardPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesCartesBancairesPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesCashappPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesCashappPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesEpsPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesEpsPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesFpxPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesFpxPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesGbBankTransferPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesGbBankTransferPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesGrabpayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesIdealPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesIdealPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesJcbPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesJcbPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesJpBankTransferPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesJpBankTransferPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesKakaoPayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesKlarnaPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesKonbiniPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesKrCardPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesKrCardPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesLinkPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesLinkPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesMobilepayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesMultibancoPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesMxBankTransferPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesMxBankTransferPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesNaverPayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesOxxoPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesOxxoPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesP24PaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesP24Payments`, `V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesPayByBankPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesPaycoPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesPaycoPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesPaynowPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesPaynowPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesPromptpayPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesPromptpayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesRevolutPayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesSamsungPayPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesSepaBankTransferPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesSepaBankTransferPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesSepaDebitPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesSepaDebitPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesStripeBalancePayouts`, `V2CoreAccountConfigurationMerchantCapabilitiesSwishPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesSwishPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesTwintPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesTwintPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesUsBankTransferPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesUsBankTransferPayments`, `V2CoreAccountConfigurationMerchantCapabilitiesZipPaymentsParams`, `V2CoreAccountConfigurationMerchantCapabilitiesZipPayments`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsInstantParams`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsInstant`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocalParams`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsLocal`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWireParams`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsWire`, `V2CoreAccountConfigurationRecipientCapabilitiesCardsParams`, `V2CoreAccountConfigurationRecipientCapabilitiesCards`, `V2CoreAccountConfigurationRecipientCapabilitiesCryptoWalletsParams`, `V2CoreAccountConfigurationRecipientCapabilitiesCryptoWallets`, `V2CoreAccountConfigurationRecipientCapabilitiesPaperChecksParams`, `V2CoreAccountConfigurationRecipientCapabilitiesPaperChecks`, `V2CoreAccountConfigurationRecipientCapabilitiesStripeBalancePayouts`, `V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfersParams`, `V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceStripeTransfers`, `V2CoreAccountConfigurationStorerCapabilitiesConsumerHoldsCurrenciesUsdParams`, `V2CoreAccountConfigurationStorerCapabilitiesConsumerHoldsCurrenciesUsd`, `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesBankAccountsParams`, `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesBankAccounts`, `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesCryptoWalletsParams`, `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesCryptoWallets`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesEurParams`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesEur`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesGbpParams`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesGbp`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUsdParams`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUsd`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUsdcParams`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUsdc`, `V2CoreAccountConfigurationStorerCapabilitiesInboundTransfersBankAccountsParams`, `V2CoreAccountConfigurationStorerCapabilitiesInboundTransfersBankAccounts`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsBankAccountsParams`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsBankAccounts`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCardsParams`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCards`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCryptoWalletsParams`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCryptoWallets`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsFinancialAccountsParams`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsFinancialAccounts`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsPaperChecksParams`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsPaperChecks`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersBankAccountsParams`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersBankAccounts`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersCryptoWalletsParams`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersCryptoWallets`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersFinancialAccountsParams`, and `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersFinancialAccounts`

## 86.1.0-alpha.1 - 2026-06-10
This release changes the pinned API version to `2026-06-10.preview`.

* ⚠️ [#2374](https://github.com/stripe/stripe-go/pull/2374) Update generated code for private-preview
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

## 85.3.0-alpha.2 - 2026-06-03
This release changes the pinned API version to `2026-06-03.preview`.

* [#2368](https://github.com/stripe/stripe-go/pull/2368) Remove `Limit`, `StartingAfter`, and `EndingBefore` fields for `List` methods that do not accept those fields
  <!-- Include any links or additional information that help explain this change. -->
  - Fixes a bug where `Limit`, `StartingAfter`, and `EndingBefore` were embedded in `CapabilityListParams`, `PaymentLocationCapabilityListParams`, and `ReportingReportTypeListParams` even though they are not valid parameters and would have always resulted in a 400 from the Stripe API if set. If you were including them before in any of those 3 structs, you can safely remove them.
* ⚠️ [#2365](https://github.com/stripe/stripe-go/pull/2365) Update generated code for private-preview
  * Add support for new resources `DelegatedCheckoutOrderEvent`, `DelegatedCheckoutOrder`, `V2BillingContractLicensePricingQuantityChange`, `V2BillingContract`, and `V2SignalsAccountSignal`
  * Add support for `Get` method on resource `DelegatedCheckoutOrder`
  * Add support for `ListOrders` method on resource `DelegatedCheckoutRequestedSession`
  * Add support for `Get` and `List` methods on resource `V2SignalsAccountSignal`
  * Add support for `Activate`, `Cancel`, `Get`, `List`, `New`, and `Update` methods on resource `V2BillingContract`
  * Add support for `BirthAddress` on `AccountIndividualParams`, `AccountPersonParams`, `Person`, `TokenAccountIndividualParams`, and `TokenPersonParams`
  * Change type of `ChargeCapturePaymentDetailsMoneyServicesParams.TransactionType`, `ChargePaymentDetailsMoneyServicesParams.TransactionType`, `PaymentIntentCapturePaymentDetailsMoneyServicesParams.TransactionType`, `PaymentIntentConfirmPaymentDetailsMoneyServicesParams.TransactionType`, and `PaymentIntentPaymentDetailsMoneyServicesParams.TransactionType` from `literal('account_funding')` to `enum('account_funding'|'debt_repayment')`
  * Add support for new value `proserv` on enums `CheckoutSessionAutomaticSurcharge.Provider` and `PaymentLinkAutomaticSurcharge.Provider`
  * Add support for `ProvisioningDecision` and `TokenType` on `IssuingAuthorizationTokenDetails` and `IssuingToken`
  * Add support for `TokenDecisionRecommendation` on `IssuingAuthorizationTokenDetailsNetworkDataVisa` and `IssuingTokenNetworkDataVisa`
  * Add support for `Language` on `IssuingTokenNetworkDataDevice`
  * Add support for `DigitalAssetCategory` on `PaymentIntentConfirmPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`
  * Add support for `StaticAddress` on `PaymentIntentConfirmPaymentMethodOptionsCryptoDepositOptionsParams`, `PaymentIntentPaymentMethodOptionsCryptoDepositOptionsParams`, and `PaymentIntentPaymentMethodOptionsCryptoDepositOptions`
  * Add support for `PaymentReference` on `PaymentIntentPaymentsOrchestrationParams`
  * ⚠️ Remove support for `PaymentDetails` on `PaymentIntentPaymentsOrchestrationParams`
  * ⚠️ Change type of `PaymentIntentPaymentDetailsMoneyServices.TransactionType` from `literal('account_funding')` to `enum('account_funding'|'debt_repayment')`
  * ⚠️ Add support for `EndingBefore`, `Limit`, and `StartingAfter` on `PaymentLocationListParams`
  * Add support for `Schema` on `V2DataReportingQueryRunResultFile` and `V2ReportingReportRunResultFile`
  * Add support for new value `payout_method_amount_limit_exceeded` on enum `V2MoneyManagementOutboundPaymentStatusDetailsFailed.Reason`
  * Add support for `Include` on `V2DataReportingQueryRunParams` and `V2ReportingReportRunParams`
  * Add support for `RequirementsCollector` on `V2CoreAccountDefaultsResponsibilitiesParams`
  * Add support for event notification `V2SignalsAccountSignalMerchantDelinquencyReadyEvent` with related object `V2SignalsAccountSignal`

## 85.3.0-alpha.1 - 2026-05-27
This release changes the pinned API version to `2026-05-27.preview`.

* ⚠️ [#2358](https://github.com/stripe/stripe-go/pull/2358) Update generated code for private-preview
  * Change type of `BillingAlertSpendThresholdParams.GroupBy` from `literal('pricing_plan_subscription')` to `enum('billing_cadence'|'pricing_plan_subscription')`
  * ⚠️ Change type of `BillingAlertSpendThreshold.GroupBy` from `literal('pricing_plan_subscription')` to `enum('billing_cadence'|'pricing_plan_subscription')`
  * Add support for new value `institution_requirement` on enum `FinancialConnectionsAccountStatusDetailsInactive.Cause`
  * Add support for `WeChatPay` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
  * Add support for `GiftCard` on `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptionsParams`, and `PaymentIntentPaymentMethodOptions`
  * Add support for `PaymentDetails` on `PaymentIntentPaymentsOrchestrationParams`
  * Add support for `Enabled` on `PaymentIntentPaymentDetailsBenefitFrMealVoucher` and `SetupIntentSetupDetailsBenefitFrMealVoucher`
  * ⚠️ Remove support for `LoginFailed`, `RegistrationFailed`, `RegistrationSuccess`, and `Type` on `RadarCustomerEvaluationParams`
  * ⚠️ Remove support for `LatestVersion` on `V2BillingLicenseFee`, `V2BillingPricingPlan`, and `V2BillingRateCard`
  * ⚠️ Remove support for `ServiceIntervalCount` and `ServiceInterval` on `V2BillingLicenseFee` and `V2BillingRateCard`
  * Add support for `DebitAgreement` on `V2MoneyManagementReceivedCreditStripeBalancePayment`
  * Add support for `CanonicalPath` on `EventsV2CoreHealthTrafficVolumeDropFiringEventImpact` and `EventsV2CoreHealthTrafficVolumeDropResolvedEventImpact`
  * Add support for snapshot event `EventTypePaymentIntentExpired` with resource `PaymentIntent`
  * Add support for event notifications `V2CoreHealthElementsErrorFiringEvent`, `V2CoreHealthElementsErrorResolvedEvent`, `V2CoreHealthInvoiceCountDroppedFiringEvent`, and `V2CoreHealthInvoiceCountDroppedResolvedEvent`

## 85.2.0-alpha.6 - 2026-05-20
* ⚠️ [#2354](https://github.com/stripe/stripe-go/pull/2354) Update generated code for private-preview
  * Add support for new resource `PaymentLocationCapability`
  * Add support for `Get`, `List`, and `Update` methods on resource `PaymentLocationCapability`
  * Add support for `Close` and `SimulateNetworkLifecycleDisputeResponse` test helper methods on resource `IssuingDispute`
  * Change type of `DelegatedCheckoutRequestedSessionDiscountsParams.Codes` from `array(string)` to `emptyable(array(string))`
  * ⚠️ Remove support for `CreditedItems` on `InvoiceItemProrationDetails`
  * Add support for `BalanceResponse` on `IssuingAuthorization`
  * Add support for `PaymentEvaluations` on `PaymentAttemptRecordReportCanceledParams`, `PaymentAttemptRecordReportFailedParams`, `PaymentRecordReportPaymentAttemptCanceledParams`, `PaymentRecordReportPaymentAttemptFailedParams`, and `PaymentRecordReportPaymentFailedParams`
  * Add support for `Enabled` on `PaymentIntentConfirmPaymentDetailsBenefitFrMealVoucherParams`, `PaymentIntentPaymentDetailsBenefitFrMealVoucherParams`, `SetupIntentConfirmSetupDetailsBenefitFrMealVoucherParams`, and `SetupIntentSetupDetailsBenefitFrMealVoucherParams`
  * Add support for `AdvancedFeatureDetails` and `AllowedPaymentMethodTypes` on `PaymentIntent`
  * Change type of `PaymentLocationAddressParams.City` from `string` to `emptyable(string)`
  * Change type of `PaymentLocationAddressParams.Line1` from `string` to `emptyable(string)`
  * Change type of `PaymentLocationAddressParams.Line2` from `string` to `emptyable(string)`
  * Change type of `PaymentLocationAddressParams.PostalCode` from `string` to `emptyable(string)`
  * Change type of `PaymentLocationAddressParams.State` from `string` to `emptyable(string)`
  * ⚠️ Remove support for `PaymentBehavior` on `SubscriptionResumeParams`
  * ⚠️ Remove support for `StatusDetails` on `Subscription`

## 85.2.0-alpha.5 - 2026-05-13
* ⚠️ [#2353](https://github.com/stripe/stripe-go/pull/2353) Update generated code for private-preview
  * Add support for new resources `V2CoreFeeBatch`, `V2CoreFeeEntry`, `V2MoneyManagementDebitDispute`, and `V2MoneyManagementFinancialAccountStatement`
  * Add support for `SimulateNetworkLifecyclePreArbitrationResponse` and `SimulateNetworkLifecyclePreArbitrationSubmission` test helper methods on resource `IssuingDispute`
  * Add support for `List` method on resource `PaymentLocation`
  * Add support for `Get` and `List` methods on resources `V2CoreFeeBatch`, `V2CoreFeeEntry`, and `V2MoneyManagementFinancialAccountStatement`
  * Add support for `Get`, `List`, and `New` methods on resource `V2MoneyManagementDebitDispute`
  * Add support for `Discounts` on `DelegatedCheckoutRequestedSessionParams` and `DelegatedCheckoutRequestedSession`
  * Add support for `AmountSale` on `DelegatedCheckoutRequestedSessionLineItemDetail` and `DelegatedCheckoutRequestedSessionTotalDetails`
  * Add support for `AmountDiscount` and `Breakdown` on `DelegatedCheckoutRequestedSessionTotalDetails`
  * ⚠️ Remove support for `CheckDepositAddress` on `InvoicePaymentSettingsPaymentMethodOptionsCheckScanParams`, `InvoicePaymentSettingsPaymentMethodOptionsCheckScan`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCheckScan`, `SubscriptionPaymentSettingsPaymentMethodOptionsCheckScanParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsCheckScan`
  * Add support for `PaymentEvaluations` on `PaymentAttemptRecordReportGuaranteedParams`, `PaymentRecordReportPaymentAttemptGuaranteedParams`, and `PaymentRecordReportPaymentGuaranteedParams`
  * Add support for `Location` on `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentPaymentDetailsParams`, `SetupIntentConfirmSetupDetailsParams`, and `SetupIntentSetupDetailsParams`
  * Add support for `OnboardingDataUpdateAcknowledged` on `PaymentLocationParams`
  * Add support for `Customer` on `RadarCustomerEvaluationParams`
  * Add support for `Status` on `RadarCustomerEvaluationParams` and `RadarCustomerEvaluation`
  * Add support for `PaymentBehavior` on `SubscriptionResumeParams`
  * Add support for `DisputeDetails` on `V2MoneyManagementReceivedDebit`
  * Add support for new value `debit_dispute` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
  * Add support for `DebitDispute` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
  * Add support for new value `debit_dispute` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
  * Add support for `PaymentAttemptRecord` on `EventsV2PaymentsOffSessionPaymentAttemptFailedEvent` and `EventsV2PaymentsOffSessionPaymentFailedEvent`
  * Add support for event notifications `V2MoneyManagementFinancialAccountStatementCreatedEvent` and `V2MoneyManagementFinancialAccountStatementRestatedEvent` with related object `V2MoneyManagementFinancialAccountStatement`

## 85.2.0-alpha.4 - 2026-05-06
* [#2352](https://github.com/stripe/stripe-go/pull/2352) Add EventNotificationHandler (private preview)
* [#2351](https://github.com/stripe/stripe-go/pull/2351) Update generated code for private-preview
  * Add support for new resource `PaymentLocation`
  * Add support for `Del`, `Get`, `New`, and `Update` methods on resource `PaymentLocation`
  * Add support for `Protections` on `AccountCapabilitiesCardPaymentsParams` and `Capability`
  * Add support for `GiftCard` on `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentPaymentMethodDataParams`, and `SharedPaymentGrantedTokenPaymentMethodDetails`
  * Add support for new value `gift_card` on enums `ConfirmationTokenPaymentMethodPreview.Type`, `PaymentMethod.Type`, and `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
  * Add support for `Metadata` on `DelegatedCheckoutRequestedSessionConfirmParams`
  * Add support for `CreditedItems` on `InvoiceItemProrationDetails`
  * Add support for `NetworkLifecycle` on `IssuingDispute`
  * Add support for new value `gift_card` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
  * Add support for `StatusDetails` on `Subscription`

## 85.2.0-alpha.3 - 2026-04-28
* [#2349](https://github.com/stripe/stripe-go/pull/2349) Update generated code for private-preview
  * Add support for `DebitCard` on `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLeadParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesConsumerLead`, `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLeadParams`, and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorConsumerLead`
  * Add support for new value `consumer.lead.debit_card` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for new value `consumer.lead.debit_card` on enum `EventsV2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEvent.UpdatedCapability`

## 85.2.0-alpha.2 - 2026-04-28
* ⚠️ [#2348](https://github.com/stripe/stripe-go/pull/2348) Update generated code for private-preview
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

## 85.2.0-alpha.1 - 2026-04-24
This release changes the pinned API version to `2026-04-22.preview`.

* ⚠️ [#2345](https://github.com/stripe/stripe-go/pull/2345) Update generated code for private-preview
  * Add support for new resources `V2CommerceProductCatalogImport`, `V2CoreApprovalRequest`, `V2ExtendWorkflowRun`, `V2ExtendWorkflow`, `V2IamActivityLog`, `V2NetworkBusinessProfile`, and `V2OrchestratedCommerceAgreement`
  * ⚠️ Remove support for resources `V2CoreWorkflowRun` and `V2CoreWorkflow`
  * Add support for `Confirm`, `Get`, `List`, `New`, and `Terminate` methods on resource `V2OrchestratedCommerceAgreement`
  * Add support for `Get` and `Me` methods on resource `V2NetworkBusinessProfile`
  * Add support for `List` method on resource `V2IamActivityLog`
  * Add support for `Get` and `List` methods on resource `V2ExtendWorkflowRun`
  * Add support for `Get`, `Invoke`, and `List` methods on resource `V2ExtendWorkflow`
  * Add support for `Cancel`, `Execute`, `Get`, `List`, and `Submit` methods on resource `V2CoreApprovalRequest`
  * Add support for `Get` and `New` methods on resource `V2CommerceProductCatalogImport`
  * ⚠️ Remove support for `Get` and `List` methods on resource `V2CoreWorkflowRun`
  * ⚠️ Remove support for `Get`, `Invoke`, and `List` methods on resource `V2CoreWorkflow`
  * Add support for `RenewOnboardingLink` method on resource `V2CoreClaimableSandbox`
  * ⚠️ Remove support for `Customer` on `SharedPaymentIssuedToken`
  * Add support for `BillManagement` and `SendMoney` on `AccountSessionComponentsBillsFeatures`
  * Add support for `GiftCard` on `ChargePaymentMethodDetails`, `PaymentAttemptRecordPaymentMethodDetails`, and `PaymentRecordPaymentMethodDetails`
  * Add support for `CustomPaymentMethodTypes` on `CheckoutSessionParams` and `CheckoutSession`
  * Add support for `PaymentRecord` on `CheckoutSession`
  * ⚠️ Remove support for `SharedPaymentGrantedToken` on `ConfirmationTokenPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntent`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for `PaymentMethod` on `ConfirmationTokenPaymentMethodPreviewSepaDebitGeneratedFrom`, `PaymentMethodSepaDebitGeneratedFrom`, and `SharedPaymentGrantedTokenPaymentMethodDetailsSepaDebitGeneratedFrom`
  * Add support for `ReturnURL` on `DelegatedCheckoutRequestedSessionConfirmParams`
  * Add support for `BuyerConsents` on `DelegatedCheckoutRequestedSession`
  * ⚠️ Change type of `DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptions.Type`, `DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOption.Type`, and `DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionOverrides.Type` from `string` to `enum('digital'|'shipping')`
  * Add support for `CryptoTransactions` on `IssuingAuthorization`, `IssuingDispute`, and `IssuingTransaction`
  * Add support for `PaymentFacilitatorID` and `SubMerchantID` on `IssuingAuthorizationMerchantData`
  * Add support for `Identifiers` on `OrderLineItemProductDataParams`, `ProductParams`, and `Product`
  * Add support for `AgentDetails` on `PaymentIntent`
  * Add support for `ExternalReference` on `PriceParams`
  * Add support for `LoginSucceeded` and `RegistrationSucceeded` on `RadarAccountEvaluationEvents` and `RadarAccountEvaluationParams`
  * Add support for `PrintContent` on `TerminalReaderAction`
  * Add support for new value `print_content` on enum `TerminalReaderAction.Type`
  * Add support for new values `cn_bank_account` and `jp_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * Add support for new values `bm_crn`, `bo_tin`, `bt_tpn`, `co_nit`, `ec_ruc`, `eg_tin`, `gh_tin`, `gy_tin`, `hn_rtn`, `jm_trn`, `jo_crn`, `ke_pin`, `ky_crn`, `lk_tin`, `mo_tin`, `mv_tin`, `ng_tin`, `pa_ruc`, `ph_tin`, `py_ruc`, `sl_tin`, `sv_nit`, `uy_ruc`, `vg_cn`, and `za_tin` on enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
  * Add support for new values `bm_pp`, `bo_ci`, `bt_cid`, `eg_tin`, `gh_pin`, `gy_tin`, `hn_rtn`, `jm_trn`, `jo_pin`, `ky_pp`, `lk_nic`, `mo_bir`, `mt_nic`, `mv_tin`, `pa_ruc`, `ph_tin`, `py_ruc`, `si_pin`, `sv_nit`, and `vg_pp` on enums `V2CoreAccountIdentityIndividualIdNumber.Type` and `V2CoreAccountPersonIdNumber.Type`
  * Add support for `AppChannel` on `V2CoreClaimableSandboxParams` and `V2CoreClaimableSandbox`
  * Add support for `OnboardingLinkDetails` and `OwnerDetails` on `V2CoreClaimableSandbox`
  * ⚠️ Remove support for `ClaimURL` on `V2CoreClaimableSandbox`
  * ⚠️ Remove support for `OwnerAccount` on `V2CoreClaimableSandboxSandboxDetails`
  * Add support for new value `live` on enum `V2CoreClaimableSandbox.Status`
  * Add support for `SnapshotEvent` on `V2CoreEvent`
  * Add support for new values `futsu` and `toza` on enums `V2CoreVaultGbBankAccount.BankAccountType` and `V2MoneyManagementPayoutMethodBankAccount.BankAccountType`
  * Add support for `MultiprocessorSettlement` on `V2MoneyManagementFinancialAccount`
  * Add support for new value `multiprocessor_settlement` on enum `V2MoneyManagementFinancialAccount.Type`
  * Add support for `CaBankAccount` on `V2MoneyManagementFinancialAddressCredentials` and `V2MoneyManagementReceivedCreditBankTransfer`
  * Add support for new value `ca_bank_account` on enum `V2MoneyManagementFinancialAddressCredentials.Type`
  * Add support for new value `tempo` on enum `V2MoneyManagementPayoutMethodCryptoWallet.Network`
  * Add support for new value `ca_bank_account` on enum `V2MoneyManagementReceivedCreditBankTransfer.OriginType`
  * ⚠️ Remove support for value `return` from enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
  * Add support for `AmountDetails` and `PaymentDetails` on `V2PaymentsOffSessionPaymentCaptureParams`, `V2PaymentsOffSessionPaymentParams`, and `V2PaymentsOffSessionPayment`
  * Add support for `Description` on `V2PaymentsOffSessionPaymentParams` and `V2PaymentsOffSessionPayment`
  * Add support for `MCC` on `V2PaymentsOffSessionPaymentPaymentMethodOptionsCardParams`
  * Add support for `Storage` on `V2MoneyManagementFinancialAccountParams`
  * Add support for `FxQuote` on `V2MoneyManagementCurrencyConversionParams`
  * ⚠️ Add support for `OnboardingLinkDetails` on `V2CoreClaimableSandboxParams`
  * Change type of `V2CoreBatchJobEndpointParams.HTTPMethod` from `literal('post')` to `enum('delete'|'post')`
  * Add support for new value `meter_event_value_too_many_digits` on enums `EventsV1BillingMeterErrorReportTriggeredEventReasonErrorType.Code` and `EventsV1BillingMeterNoMeterFoundEventReasonErrorType.Code`
  * Add support for `TreasuryTransaction` on `EventsV2MoneyManagementTransactionCreatedEvent`
  * Add support for event notifications `V1AccountApplicationAuthorizedEvent`, `V1AccountApplicationDeauthorizedEvent`, `V1AccountExternalAccountCreatedEvent`, `V1AccountExternalAccountDeletedEvent`, `V1AccountExternalAccountUpdatedEvent`, `V1BillingPortalSessionCreatedEvent`, `V1EntitlementsActiveEntitlementSummaryUpdatedEvent`, `V2CoreHealthMeterEventSummariesDelayedFiringEvent`, and `V2CoreHealthMeterEventSummariesDelayedResolvedEvent`
  * Add support for event notification `V1AccountUpdatedEvent` with related object `Account`
  * Add support for event notifications `V1ApplicationFeeCreatedEvent` and `V1ApplicationFeeRefundedEvent` with related object `ApplicationFee`
  * Add support for event notification `V1ApplicationFeeRefundUpdatedEvent` with related object `FeeRefund`
  * Add support for event notification `V1BalanceAvailableEvent` with related object `Balance`
  * Add support for event notification `V1BillingAlertTriggeredEvent` with related object `BillingAlert`
  * Add support for event notifications `V1BillingPortalConfigurationCreatedEvent` and `V1BillingPortalConfigurationUpdatedEvent` with related object `BillingPortalConfiguration`
  * Add support for event notification `V1CapabilityUpdatedEvent` with related object `Capability`
  * Add support for event notification `V1CashBalanceFundsAvailableEvent` with related object `CashBalance`
  * Add support for event notifications `V1ChargeCapturedEvent`, `V1ChargeExpiredEvent`, `V1ChargeFailedEvent`, `V1ChargePendingEvent`, `V1ChargeRefundedEvent`, `V1ChargeSucceededEvent`, and `V1ChargeUpdatedEvent` with related object `Charge`
  * Add support for event notifications `V1ChargeDisputeClosedEvent`, `V1ChargeDisputeCreatedEvent`, `V1ChargeDisputeFundsReinstatedEvent`, `V1ChargeDisputeFundsWithdrawnEvent`, and `V1ChargeDisputeUpdatedEvent` with related object `Dispute`
  * Add support for event notifications `V1ChargeRefundUpdatedEvent`, `V1RefundCreatedEvent`, `V1RefundFailedEvent`, and `V1RefundUpdatedEvent` with related object `Refund`
  * Add support for event notifications `V1CheckoutSessionAsyncPaymentFailedEvent`, `V1CheckoutSessionAsyncPaymentSucceededEvent`, `V1CheckoutSessionCompletedEvent`, and `V1CheckoutSessionExpiredEvent` with related object `CheckoutSession`
  * Add support for event notifications `V1ClimateOrderCanceledEvent`, `V1ClimateOrderCreatedEvent`, `V1ClimateOrderDelayedEvent`, `V1ClimateOrderDeliveredEvent`, and `V1ClimateOrderProductSubstitutedEvent` with related object `ClimateOrder`
  * Add support for event notifications `V1ClimateProductCreatedEvent` and `V1ClimateProductPricingUpdatedEvent` with related object `ClimateProduct`
  * Add support for event notifications `V1CouponCreatedEvent`, `V1CouponDeletedEvent`, and `V1CouponUpdatedEvent` with related object `Coupon`
  * Add support for event notifications `V1CreditNoteCreatedEvent`, `V1CreditNoteUpdatedEvent`, and `V1CreditNoteVoidedEvent` with related object `CreditNote`
  * Add support for event notifications `V1CustomerCreatedEvent`, `V1CustomerDeletedEvent`, and `V1CustomerUpdatedEvent` with related object `Customer`
  * Add support for event notifications `V1CustomerSubscriptionCreatedEvent`, `V1CustomerSubscriptionDeletedEvent`, `V1CustomerSubscriptionPausedEvent`, `V1CustomerSubscriptionPendingUpdateAppliedEvent`, `V1CustomerSubscriptionPendingUpdateExpiredEvent`, `V1CustomerSubscriptionResumedEvent`, `V1CustomerSubscriptionTrialWillEndEvent`, and `V1CustomerSubscriptionUpdatedEvent` with related object `Subscription`
  * Add support for event notifications `V1CustomerTaxIdCreatedEvent`, `V1CustomerTaxIdDeletedEvent`, and `V1CustomerTaxIdUpdatedEvent` with related object `TaxID`
  * Add support for event notification `V1CustomerCashBalanceTransactionCreatedEvent` with related object `CustomerCashBalanceTransaction`
  * Add support for event notification `V1FileCreatedEvent` with related object `File`
  * Add support for event notifications `V1FinancialConnectionsAccountCreatedEvent`, `V1FinancialConnectionsAccountDeactivatedEvent`, `V1FinancialConnectionsAccountDisconnectedEvent`, `V1FinancialConnectionsAccountReactivatedEvent`, `V1FinancialConnectionsAccountRefreshedBalanceEvent`, `V1FinancialConnectionsAccountRefreshedOwnershipEvent`, and `V1FinancialConnectionsAccountRefreshedTransactionsEvent` with related object `FinancialConnectionsAccount`
  * Add support for event notifications `V1IdentityVerificationSessionCanceledEvent`, `V1IdentityVerificationSessionCreatedEvent`, `V1IdentityVerificationSessionProcessingEvent`, `V1IdentityVerificationSessionRedactedEvent`, `V1IdentityVerificationSessionRequiresInputEvent`, and `V1IdentityVerificationSessionVerifiedEvent` with related object `IdentityVerificationSession`
  * Add support for event notifications `V1InvoiceCreatedEvent`, `V1InvoiceDeletedEvent`, `V1InvoiceFinalizationFailedEvent`, `V1InvoiceFinalizedEvent`, `V1InvoiceMarkedUncollectibleEvent`, `V1InvoiceOverdueEvent`, `V1InvoiceOverpaidEvent`, `V1InvoicePaidEvent`, `V1InvoicePaymentActionRequiredEvent`, `V1InvoicePaymentFailedEvent`, `V1InvoicePaymentSucceededEvent`, `V1InvoiceSentEvent`, `V1InvoiceUpcomingEvent`, `V1InvoiceUpdatedEvent`, `V1InvoiceVoidedEvent`, and `V1InvoiceWillBeDueEvent` with related object `Invoice`
  * Add support for event notification `V1InvoicePaymentPaidEvent` with related object `InvoicePayment`
  * Add support for event notifications `V1InvoiceitemCreatedEvent` and `V1InvoiceitemDeletedEvent` with related object `InvoiceItem`
  * Add support for event notifications `V1IssuingAuthorizationCreatedEvent`, `V1IssuingAuthorizationRequestEvent`, and `V1IssuingAuthorizationUpdatedEvent` with related object `IssuingAuthorization`
  * Add support for event notifications `V1IssuingCardCreatedEvent` and `V1IssuingCardUpdatedEvent` with related object `IssuingCard`
  * Add support for event notifications `V1IssuingCardholderCreatedEvent` and `V1IssuingCardholderUpdatedEvent` with related object `IssuingCardholder`
  * Add support for event notifications `V1IssuingDisputeClosedEvent`, `V1IssuingDisputeCreatedEvent`, `V1IssuingDisputeFundsReinstatedEvent`, `V1IssuingDisputeFundsRescindedEvent`, `V1IssuingDisputeSubmittedEvent`, and `V1IssuingDisputeUpdatedEvent` with related object `IssuingDispute`
  * Add support for event notifications `V1IssuingPersonalizationDesignActivatedEvent`, `V1IssuingPersonalizationDesignDeactivatedEvent`, `V1IssuingPersonalizationDesignRejectedEvent`, and `V1IssuingPersonalizationDesignUpdatedEvent` with related object `IssuingPersonalizationDesign`
  * Add support for event notifications `V1IssuingTokenCreatedEvent` and `V1IssuingTokenUpdatedEvent` with related object `IssuingToken`
  * Add support for event notifications `V1IssuingTransactionCreatedEvent`, `V1IssuingTransactionPurchaseDetailsReceiptUpdatedEvent`, and `V1IssuingTransactionUpdatedEvent` with related object `IssuingTransaction`
  * Add support for event notification `V1MandateUpdatedEvent` with related object `Mandate`
  * Add support for event notifications `V1PaymentIntentAmountCapturableUpdatedEvent`, `V1PaymentIntentCanceledEvent`, `V1PaymentIntentCreatedEvent`, `V1PaymentIntentPartiallyFundedEvent`, `V1PaymentIntentPaymentFailedEvent`, `V1PaymentIntentProcessingEvent`, `V1PaymentIntentRequiresActionEvent`, and `V1PaymentIntentSucceededEvent` with related object `PaymentIntent`
  * Add support for event notifications `V1PaymentLinkCreatedEvent` and `V1PaymentLinkUpdatedEvent` with related object `PaymentLink`
  * Add support for event notifications `V1PaymentMethodAttachedEvent`, `V1PaymentMethodAutomaticallyUpdatedEvent`, `V1PaymentMethodDetachedEvent`, and `V1PaymentMethodUpdatedEvent` with related object `PaymentMethod`
  * Add support for event notifications `V1PayoutCanceledEvent`, `V1PayoutCreatedEvent`, `V1PayoutFailedEvent`, `V1PayoutPaidEvent`, `V1PayoutReconciliationCompletedEvent`, and `V1PayoutUpdatedEvent` with related object `Payout`
  * Add support for event notifications `V1PersonCreatedEvent`, `V1PersonDeletedEvent`, and `V1PersonUpdatedEvent` with related object `Person`
  * Add support for event notifications `V1PlanCreatedEvent`, `V1PlanDeletedEvent`, and `V1PlanUpdatedEvent` with related object `Plan`
  * Add support for event notifications `V1PriceCreatedEvent`, `V1PriceDeletedEvent`, and `V1PriceUpdatedEvent` with related object `Price`
  * Add support for event notifications `V1ProductCreatedEvent`, `V1ProductDeletedEvent`, and `V1ProductUpdatedEvent` with related object `Product`
  * Add support for event notifications `V1PromotionCodeCreatedEvent` and `V1PromotionCodeUpdatedEvent` with related object `PromotionCode`
  * Add support for event notifications `V1QuoteAcceptedEvent`, `V1QuoteCanceledEvent`, `V1QuoteCreatedEvent`, and `V1QuoteFinalizedEvent` with related object `Quote`
  * Add support for event notifications `V1RadarEarlyFraudWarningCreatedEvent` and `V1RadarEarlyFraudWarningUpdatedEvent` with related object `RadarEarlyFraudWarning`
  * Add support for event notifications `V1ReviewClosedEvent` and `V1ReviewOpenedEvent` with related object `Review`
  * Add support for event notifications `V1SetupIntentCanceledEvent`, `V1SetupIntentCreatedEvent`, `V1SetupIntentRequiresActionEvent`, `V1SetupIntentSetupFailedEvent`, and `V1SetupIntentSucceededEvent` with related object `SetupIntent`
  * Add support for event notification `V1SigmaScheduledQueryRunCreatedEvent` with related object `SigmaScheduledQueryRun`
  * Add support for event notifications `V1SourceCanceledEvent`, `V1SourceChargeableEvent`, `V1SourceFailedEvent`, and `V1SourceRefundAttributesRequiredEvent` with related object `Source`
  * Add support for event notifications `V1SubscriptionScheduleAbortedEvent`, `V1SubscriptionScheduleCanceledEvent`, `V1SubscriptionScheduleCompletedEvent`, `V1SubscriptionScheduleCreatedEvent`, `V1SubscriptionScheduleExpiringEvent`, `V1SubscriptionScheduleReleasedEvent`, and `V1SubscriptionScheduleUpdatedEvent` with related object `SubscriptionSchedule`
  * Add support for event notification `V1TaxSettingsUpdatedEvent` with related object `TaxSettings`
  * Add support for event notifications `V1TaxRateCreatedEvent` and `V1TaxRateUpdatedEvent` with related object `TaxRate`
  * Add support for event notifications `V1TerminalReaderActionFailedEvent`, `V1TerminalReaderActionSucceededEvent`, and `V1TerminalReaderActionUpdatedEvent` with related object `TerminalReader`
  * Add support for event notifications `V1TestHelpersTestClockAdvancingEvent`, `V1TestHelpersTestClockCreatedEvent`, `V1TestHelpersTestClockDeletedEvent`, `V1TestHelpersTestClockInternalFailureEvent`, and `V1TestHelpersTestClockReadyEvent` with related object `TestHelpersTestClock`
  * Add support for event notifications `V1TopupCanceledEvent`, `V1TopupCreatedEvent`, `V1TopupFailedEvent`, `V1TopupReversedEvent`, and `V1TopupSucceededEvent` with related object `Topup`
  * Add support for event notifications `V1TransferCreatedEvent`, `V1TransferReversedEvent`, and `V1TransferUpdatedEvent` with related object `Transfer`
  * Add support for event notifications `V2CommerceProductCatalogImportsFailedEvent`, `V2CommerceProductCatalogImportsProcessingEvent`, `V2CommerceProductCatalogImportsSucceededEvent`, and `V2CommerceProductCatalogImportsSucceededWithErrorsEvent` with related object `V2CommerceProductCatalogImport`
  * Add support for event notifications `V2CoreApprovalRequestApprovedEvent`, `V2CoreApprovalRequestCanceledEvent`, `V2CoreApprovalRequestFailedEvent`, `V2CoreApprovalRequestRejectedEvent`, and `V2CoreApprovalRequestSucceededEvent` with related object `V2CoreApprovalRequest`
  * Add support for event notification `V2CoreClaimableSandboxUpdatedEvent` with related object `V2CoreClaimableSandbox`
  * Add support for event notifications `V2ExtendWorkflowRunFailedEvent`, `V2ExtendWorkflowRunStartedEvent`, and `V2ExtendWorkflowRunSucceededEvent` with related object `V2ExtendWorkflowRun`
  * Add support for event notifications `V2OrchestratedCommerceAgreementConfirmedEvent`, `V2OrchestratedCommerceAgreementCreatedEvent`, `V2OrchestratedCommerceAgreementPartiallyConfirmedEvent`, and `V2OrchestratedCommerceAgreementTerminatedEvent` with related object `V2OrchestratedCommerceAgreement`
  * ⚠️ Remove support for event notification `V2CoreClaimableSandboxSandboxDetailsOwnerAccountUpdatedEvent` with related object `V2CoreClaimableSandbox`
  * Add support for error type `FxQuoteExpiredError`
  * Add support for error codes `invalid_workflow_input_parameters` and `workflow_not_invokable` on `CannotProceedError`

## 85.1.0-alpha.4 - 2026-04-15
* [#2344](https://github.com/stripe/stripe-go/pull/2344) Update generated code for private-preview
  * Add support for `LatestVersion` on `V2BillingLicenseFee`, `V2BillingPricingPlan`, and `V2BillingRateCard`
  * Add support for `ServiceIntervalCount` and `ServiceInterval` on `V2BillingLicenseFee` and `V2BillingRateCard`
* ⚠️ [#2343](https://github.com/stripe/stripe-go/pull/2343) Update generated code for private-preview
  * Add support for new resources `V2CoreWorkflowRun` and `V2CoreWorkflow`
  * Add support for `ReportAuthorized` method on resource `PaymentAttemptRecord`
  * Add support for `Get` and `List` methods on resource `V2CoreWorkflowRun`
  * Add support for `Get`, `Invoke`, and `List` methods on resource `V2CoreWorkflow`
  * Add support for `NextAction` and `Status` on `SharedPaymentIssuedToken`
  * ⚠️ Remove support for `NetworkID` on `SharedPaymentIssuedTokenSellerDetails`
  * Add support for `Bills` on `AccountSessionComponents`
  * Add support for `SettlementCurrencies` on `BalanceSettingsPaymentsParams` and `BalanceSettingsPayments`
  * Add support for `DefaultSettlementCurrency` on `BalanceSettingsPayments`
  * Add support for `AccountFunding` on `ChargePaymentMethodDetailsCard`
  * Add support for `AutomaticSurcharge` on `CheckoutSessionParams`, `CheckoutSession`, `PaymentLinkParams`, and `PaymentLink`
  * Add support for `Bizum` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
  * Add support for `SurchargeCost` on `CheckoutSession`
  * Add support for `AmountSurcharge` on `CheckoutSessionTotalDetails`
  * Add support for `SharedPaymentGrantedToken` on `ConfirmationTokenPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
  * Add support for `Details` on `IdentityVerificationReportEmail`
  * Add support for new value `email` on enums `IdentityVerificationReport.Type` and `IdentityVerificationSession.Type`
  * Add support for `Confirm` on `IdentityVerificationSessionParams`
  * Add support for `Subscription` on `InvoiceItemParentScheduleDetails`
  * ⚠️ Remove support for `SharedPaymentGrantedToken` on `PaymentIntentConfirmParams` and `PaymentIntentParams`
  * Add support for `MoneyServices` on `PaymentIntentPaymentDetails`
  * ⚠️ Remove support for `ExternalReference` on `Plan`

## 85.1.0-alpha.3 - 2026-04-08
This release changes the pinned API version to `2026-04-08.preview`.

* ⚠️ [#2339](https://github.com/stripe/stripe-go/pull/2339) Update generated code for private-preview
  * Add support for `PaymentRecord` on `ApplicationFeeFeeSource`
  * Add support for `FleetData` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCard`, `PaymentIntentAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentCaptureAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentDecrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentPaymentDetailsParams`, and `PaymentIntentPaymentDetails`
  * Add support for `BeneficiaryAccount`, `BeneficiaryDetails`, `SenderAccount`, and `SenderDetails` on `ChargeCapturePaymentDetailsMoneyServicesAccountFundingParams`, `ChargePaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentCapturePaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentConfirmPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentDetailsMoneyServicesAccountFundingParams`
  * Change type of `ChargeCapturePaymentDetailsMoneyServicesParams.TransactionType`, `ChargePaymentDetailsMoneyServicesParams.TransactionType`, `PaymentIntentCapturePaymentDetailsMoneyServicesParams.TransactionType`, `PaymentIntentConfirmPaymentDetailsMoneyServicesParams.TransactionType`, and `PaymentIntentPaymentDetailsMoneyServicesParams.TransactionType` from `literal('account_funding')` to `emptyable(literal('account_funding'))`
  * Add support for new value `requires_action` on enum `DelegatedCheckoutRequestedSession.Status`
  * Add support for `Bizum` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
  * Add support for new value `bizum` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
  * Add support for `QuantityPrecision` on `PaymentIntentAmountDetailsLineItem`, `PaymentIntentAmountDetailsLineItemsParams`, `PaymentIntentCaptureAmountDetailsLineItemsParams`, `PaymentIntentConfirmAmountDetailsLineItemsParams`, `PaymentIntentDecrementAuthorizationAmountDetailsLineItemsParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsParams`
  * Add support for `LiquidAsset` and `Wallet` on `PaymentIntentConfirmPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`
  * Add support for `SharedPaymentGrantedToken` on `PaymentMethod`
  * ⚠️ Change type of `RadarCustomerEvaluation.EventType` from `string` to `enum('login'|'registration')`
  * ⚠️ Change type of `RadarCustomerEvaluationSignalsAccountSharing.RiskLevel` and `RadarCustomerEvaluationSignalsMultiAccounting.RiskLevel` from `string` to `enum`
  * Add support for `Data` on `RadarPaymentEvaluationClientDeviceMetadataDetailsParams` and `RadarPaymentEvaluationClientDeviceMetadataDetails`
  * Add support for `Sunbit` on `SharedPaymentGrantedTokenPaymentMethodDetails`
  * Add support for new value `sunbit` on enum `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
  * ⚠️ Remove support for values `bm_crn`, `bo_tin`, `bt_tpn`, `co_nit`, `ec_ruc`, `eg_tin`, `gh_tin`, `gy_tin`, `hn_rtn`, `jm_trn`, `jo_crn`, `ke_pin`, `ky_crn`, `lk_tin`, `mo_tin`, `mv_tin`, `ng_tin`, `pa_ruc`, `ph_tin`, `py_ruc`, `sl_tin`, `sv_nit`, `uy_ruc`, `vg_cn`, and `za_tin` from enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
  * ⚠️ Remove support for values `bm_pp`, `bo_ci`, `bt_cid`, `eg_tin`, `gh_pin`, `gy_tin`, `hn_rtn`, `jm_trn`, `jo_pin`, `ky_pp`, `lk_nic`, `mo_bir`, `mt_nic`, `mv_tin`, `pa_ruc`, `ph_tin`, `py_ruc`, `si_pin`, `sv_nit`, and `vg_pp` from enums `V2CoreAccountIdentityIndividualIdNumber.Type` and `V2CoreAccountPersonIdNumber.Type`
  * Add support for error type `CannotProceedError`

## 85.1.0-alpha.2 - 2026-04-01
This release changes the pinned API version to `2026-04-01.preview`.

* ⚠️ [#2335](https://github.com/stripe/stripe-go/pull/2335) Update generated code for private-preview
  * Add support for new resources `SharedPaymentIssuedToken` and `V2DataReportingQueryRun`
  * Add support for `Get` and `New` methods on resource `V2DataReportingQueryRun`
  * Add support for `Pause` and `Resume` methods on resource `V2PaymentsOffSessionPayment`
  * Add support for `TenantKeys`, `TenantOperator`, and `TenantValues` on `BillingBillingMeterMeterEventSummaryListParams`
  * Add support for `MoneyServices` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmPaymentDetailsParams`, and `PaymentIntentPaymentDetailsParams`
  * Add support for `PaymentMethodOptions` on `DelegatedCheckoutRequestedSessionParams` and `DelegatedCheckoutRequestedSession`
  * ⚠️ Remove support for `PaymentMethodData` on `DelegatedCheckoutRequestedSessionConfirmParams` and `DelegatedCheckoutRequestedSessionParams`
  * Add support for `CardBrands` and `PaymentMethodTypes` on `DelegatedCheckoutRequestedSessionSellerDetails`
  * ⚠️ Change type of `DelegatedCheckoutRequestedSession.SharedPaymentIssuedToken` from `string` to `expandable($SharedPayment.IssuedToken)`
  * Add support for `CheckScan` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
  * Add support for new value `check_scan` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
  * Add support for `ProcessorDetails` on `PaymentAttemptRecordReportFailedParams`, `PaymentAttemptRecordReportGuaranteedParams`, `PaymentRecordReportPaymentAttemptFailedParams`, `PaymentRecordReportPaymentAttemptGuaranteedParams`, `PaymentRecordReportPaymentFailedParams`, and `PaymentRecordReportPaymentGuaranteedParams`
  * Add support for `PaymentDetails` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardParams`, and `PaymentIntentPaymentMethodOptionsCardPresentParams`
  * ⚠️ Remove support for `BillFrom` on `QuotePreviewSubscriptionScheduleBillingSchedule`, `SubscriptionBillingSchedule`, and `SubscriptionScheduleBillingSchedule`
  * Add support for `AgentDetails`, `PaymentMethodDetails`, and `RiskDetails` on `SharedPaymentGrantedToken`
  * Add support for `PaperChecks` on `V2AccountConfigurationRecipientDataFeaturesParams`, `V2AccountConfigurationRecipientDataFeatures`, `V2CoreAccountConfigurationRecipientCapabilitiesParams`, `V2CoreAccountConfigurationRecipientCapabilities`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsParams`, and `V2CoreAccountConfigurationStorerCapabilitiesOutboundPayments`
  * Add support for new value `paper_checks` on enum `V2AccountConfigurationSupportableFeatures.RecipientData`
  * Add support for new value `paper_checks` on enum `V2AccountRequirementImpact.RequiredForFeatures`
  * ⚠️ Change type of `V2BillingCadenceSettingsDataCollectionPaymentMethodOptions.Konbini`, `V2BillingCollectionSettingPaymentMethodOptions.Konbini`, `V2BillingCollectionSettingPaymentMethodOptionsParams.Konbini`, and `V2BillingCollectionSettingVersionPaymentMethodOptions.Konbini` from `map(string: dynamic)` to `an object`
  * ⚠️ Change type of `V2BillingCadenceSettingsDataCollectionPaymentMethodOptions.SEPADebit`, `V2BillingCollectionSettingPaymentMethodOptions.SEPADebit`, `V2BillingCollectionSettingPaymentMethodOptionsParams.SEPADebit`, and `V2BillingCollectionSettingVersionPaymentMethodOptions.SEPADebit` from `map(string: dynamic)` to `an object`
  * Add support for `ID` on `V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmountCustomPricingUnit`, `V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendAmountCustomPricingUnitParams`, and `V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendAmountCustomPricingUnit`
  * Add support for new values `outbound_payments.paper_checks` and `paper_checks` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for new values `bm_crn`, `bo_tin`, `bt_tpn`, `co_nit`, `ec_ruc`, `eg_tin`, `gh_tin`, `gy_tin`, `hn_rtn`, `jm_trn`, `jo_crn`, `ke_pin`, `ky_crn`, `lk_tin`, `mo_tin`, `mv_tin`, `ng_tin`, `pa_ruc`, `ph_tin`, `py_ruc`, `sl_tin`, `sv_nit`, `uy_ruc`, `vg_cn`, and `za_tin` on enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
  * Add support for new values `bm_pp`, `bo_ci`, `bt_cid`, `eg_tin`, `gh_pin`, `gy_tin`, `hn_rtn`, `jm_trn`, `jo_pin`, `ky_pp`, `lk_nic`, `mo_bir`, `mt_nic`, `mv_tin`, `pa_ruc`, `ph_tin`, `py_ruc`, `si_pin`, `sv_nit`, and `vg_pp` on enums `V2CoreAccountIdentityIndividualIdNumber.Type` and `V2CoreAccountPersonIdNumber.Type`
  * ⚠️ Change type of `V2CoreEventReasonRequestClient.StripeAction` from `map(string: dynamic)` to `an object`
  * ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitProcessing` from `map(string: dynamic)` to `an object`
  * ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitQueued` from `map(string: dynamic)` to `an object`
  * ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitSucceeded` from `map(string: dynamic)` to `an object`
  * Add support for new values `paper_check_attachment_too_large`, `paper_check_expired`, and `paper_check_undeliverable` on enum `V2MoneyManagementOutboundPaymentStatusDetailsFailed.Reason`
  * ⚠️ Remove support for `Town` on `V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckMailingAddress`
  * Add support for new value `payout_method_amount_limit_exceeded` on enum `V2MoneyManagementOutboundTransferStatusDetailsFailed.Reason`
  * Add support for `ApplicationFeeAmountRequested` on `V2PaymentsOffSessionPayment`
  * ⚠️ Remove support for `CompartmentID` on `V2PaymentsOffSessionPayment`
  * Add support for new value `exceeded_retry_window` on enum `V2PaymentsOffSessionPayment.FailureReason`
  * Add support for `RetryUntil` on `V2PaymentsOffSessionPaymentRetryDetails`
  * Add support for new value `paused` on enum `V2PaymentsOffSessionPayment.Status`
  * Add support for `ApplicationFeeAmount` on `V2PaymentsOffSessionPaymentCaptureParams` and `V2PaymentsOffSessionPaymentParams`
  * Add support for new value `paper_checks` on enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`
  * Add support for new value `outbound_payments.paper_checks` on enum `EventsV2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent.UpdatedCapability`
  * Add support for `AlertID` on `EventsV2CoreHealthApiErrorResolvedEvent`, `EventsV2CoreHealthApiLatencyResolvedEvent`, `EventsV2CoreHealthAuthorizationRateDropResolvedEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestErrorsFiringEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestErrorsResolvedEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEvent`, `EventsV2CoreHealthPaymentMethodErrorResolvedEvent`, `EventsV2CoreHealthSepaDebitDelayedFiringEvent`, `EventsV2CoreHealthSepaDebitDelayedResolvedEvent`, `EventsV2CoreHealthTrafficVolumeDropResolvedEvent`, and `EventsV2CoreHealthWebhookLatencyResolvedEvent`
  * Add support for `APIKey` on `EventsV2IamApiKeyCreatedEvent`, `EventsV2IamApiKeyDefaultSecretRevealedEvent`, `EventsV2IamApiKeyExpiredEvent`, `EventsV2IamApiKeyPermissionsUpdatedEvent`, `EventsV2IamApiKeyRotatedEvent`, and `EventsV2IamApiKeyUpdatedEvent`
  * Add support for `StripeAccessGrant` on `EventsV2IamStripeAccessGrantApprovedEvent`, `EventsV2IamStripeAccessGrantCanceledEvent`, `EventsV2IamStripeAccessGrantDeniedEvent`, `EventsV2IamStripeAccessGrantRemovedEvent`, `EventsV2IamStripeAccessGrantRequestedEvent`, and `EventsV2IamStripeAccessGrantUpdatedEvent`
  * Add support for event notifications `V2DataReportingQueryRunCreatedEvent`, `V2DataReportingQueryRunFailedEvent`, `V2DataReportingQueryRunSucceededEvent`, and `V2DataReportingQueryRunUpdatedEvent` with related object `V2DataReportingQueryRun`
  * Add support for event notifications `V2PaymentsOffSessionPaymentPausedEvent` and `V2PaymentsOffSessionPaymentResumedEvent` with related object `V2PaymentsOffSessionPayment`

## 85.1.0-alpha.1 - 2026-03-25
This release changes the pinned API version to `2026-03-25.preview`.

This release contains additional breaking changes. See the [GA changelog](https://github.com/stripe/stripe-go/blob/master/CHANGELOG.md#8500---2026-03-25) for more information.

* [#2319](https://github.com/stripe/stripe-go/pull/2319) Commented out a failing test
* ⚠️ [#2305](https://github.com/stripe/stripe-go/pull/2305) Update generated code for private-preview
  * Add support for new resource `V2CoreAccountEvaluation`
  * ⚠️ Remove support for resources `V2BillingLicenseFeeSubscription` and `V2BillingPricingPlanSubscriptionComponents`
  * Add support for `New` method on resource `V2CoreAccountEvaluation`
  * ⚠️ Remove support for `Get` method on resources `V2BillingLicenseFeeSubscription` and `V2BillingPricingPlanSubscriptionComponents`
  * Add support for `ModifyRates` method on resource `V2BillingRateCard`
  * Add support for `RemoveDiscounts` method on resource `V2BillingPricingPlanSubscription`
  * Add support for new value `eg_bank_account` on enum `V2AccountConfigurationRecipientDataDefaultOutboundDestination.Type`
  * Add support for `InvoiceResources` on `V2BillingIntent`
  * Add support for `AmountDue` and `CustomerBalanceApplied` on `V2BillingIntentAmountDetails`
  * Add support for `ExpiresAt` on `V2BillingIntentStatusTransitions`
  * Add support for `Discount` on `V2BillingIntentActionApplyParams` and `V2BillingIntentActionApply`
  * Add support for `Timestamp` on `V2BillingIntentActionApplyEffectiveAtParams` and `V2BillingIntentActionApplyEffectiveAt`
  * Add support for new values `current_billing_period_start` and `timestamp` on enum `V2BillingIntentActionApplyEffectiveAt.Type`
  * Add support for new value `discount` on enum `V2BillingIntentActionApply.Type`
  * ⚠️ Change type of `V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior.Type`, `V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams.Type`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior.Type`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams.Type`, `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior.Type`, and `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams.Type` from `literal('license_fee')` to `enum('license_fee'|'recurring_credit_grant')`
  * Add support for `ServiceCycle` on `V2BillingLicenseFee` and `V2BillingRateCard`
  * ⚠️ Remove support for `LatestVersion` on `V2BillingLicenseFee`, `V2BillingPricingPlan`, and `V2BillingRateCard`
  * ⚠️ Remove support for `ServiceIntervalCount` and `ServiceInterval` on `V2BillingLicenseFee` and `V2BillingRateCard`
  * ⚠️ Change type of `V2BillingLicenseFeeTransformQuantity.DivideBy`, `V2BillingLicenseFeeTransformQuantityParams.DivideBy`, `V2BillingLicenseFeeVersionTransformQuantity.DivideBy`, `V2BillingRateCardRateTransformQuantity.DivideBy`, and `V2BillingRateCardRateTransformQuantityParams.DivideBy` from `longInteger` to `int64_string`
  * Add support for `DiscountDetails` and `PricingPlanComponentDetails` on `V2BillingPricingPlanSubscription`
  * Add support for new value `crypto_wallets` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * ⚠️ Remove support for value `crypto` from enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for `BalanceByFundsType` on `V2MoneyManagementFinancialAccountPayments`
  * Add support for new value `next_day_payout_fee` on enum `V2MoneyManagementOutboundPaymentQuoteEstimatedFee.Type`
  * Add support for `TreasuryTransactionEntry` on `V2MoneyManagementTransactionEntry`
  * Add support for `TreasuryCreditReversal`, `TreasuryDebitReversal`, `TreasuryInboundTransfer`, `TreasuryIssuingAuthorization`, `TreasuryOutboundPayment`, `TreasuryOutboundTransfer`, `TreasuryReceivedCredit`, and `TreasuryReceivedDebit` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
  * Add support for new values `treasury_credit_reversal`, `treasury_debit_reversal`, `treasury_inbound_transfer`, `treasury_issuing_authorization`, `treasury_other`, `treasury_outbound_payment`, `treasury_outbound_transfer`, `treasury_received_credit`, and `treasury_received_debit` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
  * Add support for `TreasuryTransaction` on `V2MoneyManagementTransaction`
  * Add support for new value `no_valid_payment_method` on enum `V2PaymentsOffSessionPayment.FailureReason`
  * Add support for `Metadata` on `V2PaymentsSettlementAllocationIntentSplit`
  * ⚠️ Change type of `V2ReportingReportRunResultFile.Size` from `longInteger` to `int64_string`
  * Add support for `StatementDescriptor` on `V2MoneyManagementOutboundPaymentParams` and `V2MoneyManagementOutboundTransferParams`
  * Add support for `Include` on `V2BillingIntentParams`, `V2BillingIntentReserveParams`, `V2BillingPricingPlanSubscriptionListParams`, `V2BillingPricingPlanSubscriptionParams`, `V2MoneyManagementFinancialAccountListParams`, and `V2MoneyManagementFinancialAccountParams`
  * Add support for event notifications `V1AccountSignalsIncludingDelinquencyCreatedEvent`, `V2CoreAccountSignalsFraudulentWebsiteReadyEvent`, and `V2SignalsAccountSignalFraudulentMerchantReadyEvent`
* ⚠️ [#2334](https://github.com/stripe/stripe-go/pull/2334) Update generated code for private-preview
  * Add support for new resource `RiskSignals`
  * Add support for `FinancialAccountRewards` and `NestingDemo` on `AccountSessionComponents`
  * Add support for `UpiPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
  * Add support for `RiskSignals` on `Account`
  * Add support for `FraudIntent` on `AccountSignals`
  * Add support for new value `related_accounts` on enum `AccountSignalsDelinquencyIndicators.Indicator`
  * Add support for `RiskReserved` on `Balance`
  * ⚠️ Remove support for `BillableItems` on `BillingAlertSpendThresholdFilters`
  * Add support for `Upi` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `MandatePaymentMethodDetails`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodDataParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
  * Add support for new value `tempo` on enums `ChargePaymentMethodDetailsCrypto.Network`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.Network`, and `PaymentRecordPaymentMethodDetailsCrypto.Network`
  * ⚠️ Remove support for `SourceType` on `ChargePaymentMethodDetailsStripeBalance`, `ConfirmationTokenPaymentMethodDataStripeBalanceParams`, `ConfirmationTokenPaymentMethodPreviewStripeBalance`, `PaymentAttemptRecordPaymentMethodDetailsStripeBalance`, `PaymentIntentConfirmPaymentMethodDataStripeBalanceParams`, `PaymentIntentPaymentMethodDataStripeBalanceParams`, `PaymentMethodStripeBalanceParams`, `PaymentMethodStripeBalance`, `PaymentRecordPaymentMethodDetailsStripeBalance`, `SetupIntentConfirmPaymentMethodDataStripeBalanceParams`, and `SetupIntentPaymentMethodDataStripeBalanceParams`
  * Add support for `IntegrationIdentifier` on `CheckoutSessionParams` and `CheckoutSession`
  * Change type of `CheckoutSessionLineItemPriceDataProductDataTaxDetailsParams.TaxCode`, `InvoiceAddLinesLinePriceDataProductDataTaxDetailsParams.TaxCode`, `InvoiceLineItemPriceDataProductDataTaxDetailsParams.TaxCode`, `InvoiceUpdateLinesLinePriceDataProductDataTaxDetailsParams.TaxCode`, `PaymentLinkLineItemPriceDataProductDataTaxDetailsParams.TaxCode`, `PlanProductTaxDetailsParams.TaxCode`, `PriceProductDataTaxDetailsParams.TaxCode`, and `ProductTaxDetailsParams.TaxCode` from `string` to `emptyable(string)`
  * Add support for `Crypto` on `CheckoutSessionPaymentMethodOptionsParams`
  * Add support for `PendingInvoiceItemInterval` on `CheckoutSessionSubscriptionDataParams`
  * Add support for new value `application` on enums `CheckoutSessionAutomaticTaxLiability.Type`, `CheckoutSessionInvoiceCreationInvoiceDataIssuer.Type`, `InvoiceAutomaticTaxLiability.Type`, `InvoiceIssuer.Type`, `PaymentLinkAutomaticTaxLiability.Type`, `PaymentLinkInvoiceCreationInvoiceDataIssuer.Type`, `PaymentLinkSubscriptionDataInvoiceSettingsIssuer.Type`, `QuoteAutomaticTaxLiability.Type`, `QuoteInvoiceSettingsIssuer.Type`, `QuotePreviewInvoiceAutomaticTaxLiability.Type`, `QuotePreviewInvoiceIssuer.Type`, `QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTaxLiability.Type`, `QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer.Type`, `QuotePreviewSubscriptionSchedulePhaseAutomaticTaxLiability.Type`, `QuotePreviewSubscriptionSchedulePhaseInvoiceSettingsIssuer.Type`, `SubscriptionAutomaticTaxLiability.Type`, `SubscriptionInvoiceSettingsIssuer.Type`, `SubscriptionScheduleDefaultSettingsInvoiceSettingsIssuer.Type`, and `SubscriptionSchedulePhaseInvoiceSettingsIssuer.Type`
  * Add support for `AUBECSDebit`, `BACSDebit`, `Boleto`, `Link`, `SEPADebit`, and `USBankAccount` on `CheckoutSessionCurrentAttemptPaymentMethodDetails`
  * Add support for new values `elements`, `embedded_page`, `form`, and `hosted_page` on enum `CheckoutSession.UIMode`
  * ⚠️ Remove support for values `custom`, `embedded`, and `hosted` from enum `CheckoutSession.UIMode`
  * Add support for new value `marine_carbon_removal` on enum `ClimateSupplier.RemovalPathway`
  * Add support for new value `upi` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
  * Add support for `Metadata` on `CreditNoteLineItem`, `CreditNoteLineParams`, `CreditNotePreviewLineParams`, and `CreditNotePreviewLinesLineParams`
  * Add support for `SelectedFulfillmentOptionOverrides` on `DelegatedCheckoutRequestedSessionFulfillmentDetails`
  * Add support for `LineItemKeys` on `DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionsDigitalDigitalOptions` and `DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptionsShippingShippingOptions`
  * Add support for `QuantityDecimal` on `InvoiceAddLinesLineParams`, `InvoiceCreatePreviewInvoiceItemParams`, `InvoiceItemParams`, `InvoiceItem`, `InvoiceLineItemParams`, `InvoiceLineItem`, and `InvoiceUpdateLinesLineParams`
  * Add support for `ExpiresAfterSeconds` on `InvoicePaymentSettingsPaymentMethodOptionsPixParams`, `InvoicePaymentSettingsPaymentMethodOptionsPix`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsPix`, `SubscriptionPaymentSettingsPaymentMethodOptionsPixParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsPix`
  * ⚠️ Add support for `Level` on `IssuingAuthorizationRiskAssessmentCardTestingRiskParams` and `IssuingAuthorizationRiskAssessmentMerchantDisputeRiskParams`
  * ⚠️ Remove support for `RiskLevel` on `IssuingAuthorizationRiskAssessmentCardTestingRiskParams` and `IssuingAuthorizationRiskAssessmentMerchantDisputeRiskParams`
  * Add support for new values `da`, `pl`, and `sv` on enum `IssuingCardholder.PreferredLocales`
  * Add support for `LifecycleControls` on `IssuingCardParams` and `IssuingCard`
  * Add support for `Cryptogram`, `ElectronicCommerceIndicator`, `ExemptionIndicatorApplied`, and `ExemptionIndicator` on `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure` and `PaymentRecordPaymentMethodDetailsCardThreeDSecure`
  * Add support for `Surcharge` on `PaymentIntentAmountDetailsParams`, `PaymentIntentAmountDetails`, `PaymentIntentCaptureAmountDetailsParams`, `PaymentIntentConfirmAmountDetailsParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsParams`
  * Add support for `MandateOptions` on `PaymentIntentConfirmPaymentMethodOptionsStripeBalanceParams`, `PaymentIntentPaymentMethodOptionsStripeBalanceParams`, and `PaymentIntentPaymentMethodOptionsStripeBalance`
  * Add support for `AmountDetails` and `PaymentDetails` on `PaymentIntentDecrementAuthorizationParams`
  * Add support for new value `upi` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
  * Add support for `UpiHandleRedirectOrDisplayQRCode` on `PaymentIntentNextAction` and `SetupIntentNextAction`
  * Add support for `ManagedPayments` on `PaymentLinkParams` and `PaymentLink`
  * Add support for new value `upi` on enum `PaymentLink.PaymentMethodTypes`
  * Add support for `RecommendedAction` and `Signals` on `RadarPaymentEvaluation`
  * ⚠️ Remove support for `Insights` on `RadarPaymentEvaluation`
  * Add support for new value `crypto_fingerprint` on enum `RadarValueList.ItemType`
  * Add support for `StripeBalance` on `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
  * Add support for new value `resolved` on enum `SharedPaymentGrantedToken.DeactivatedReason`
  * Add support for `RecurringInterval` on `SharedPaymentGrantedTokenUsageLimits`
  * Add support for `PresentmentDetails` on `Subscription`
  * Add support for new value `canceled_by_retention_policy` on enum `SubscriptionCancellationDetails.Reason`
  * ⚠️ Remove support for `InvoiceResources` on `V2BillingIntent`
  * ⚠️ Remove support for `AmountDue` and `CustomerBalanceApplied` on `V2BillingIntentAmountDetails`
  * Add support for `RecurringCreditGrant` on `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior`, `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams`, and `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior`
  * Add support for `ConsumerPrivacyDisclosures` and `ConsumerStorer` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams` and `V2CoreAccountIdentityAttestationsTermsOfService`
  * ⚠️ Remove support for `Include` on `V2BillingIntentParams` and `V2BillingIntentReserveParams`
  * Add support for error code `service_period_coupon_with_metered_tiered_item_unsupported` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `QuotePreviewInvoiceLastFinalizationError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
* ⚠️ [#2332](https://github.com/stripe/stripe-go/pull/2332) Update generated code for private-preview
  * Add support for new resource `V2CoreAccountEvaluation`
  * ⚠️ Remove support for resources `V2BillingLicenseFeeSubscription` and `V2BillingPricingPlanSubscriptionComponents`
  * Add support for `New` method on resource `V2CoreAccountEvaluation`
  * ⚠️ Remove support for `Get` method on resources `V2BillingLicenseFeeSubscription` and `V2BillingPricingPlanSubscriptionComponents`
  * Add support for `ModifyRates` method on resource `V2BillingRateCard`
  * Add support for `RemoveDiscounts` method on resource `V2BillingPricingPlanSubscription`
  * Add support for new value `eg_bank_account` on enum `V2AccountConfigurationRecipientDataDefaultOutboundDestination.Type`
  * Add support for `InvoiceResources` on `V2BillingIntent`
  * Add support for `AmountDue` and `CustomerBalanceApplied` on `V2BillingIntentAmountDetails`
  * Add support for `ExpiresAt` on `V2BillingIntentStatusTransitions`
  * Add support for `Discount` on `V2BillingIntentActionApplyParams` and `V2BillingIntentActionApply`
  * Add support for `Timestamp` on `V2BillingIntentActionApplyEffectiveAtParams` and `V2BillingIntentActionApplyEffectiveAt`
  * Add support for new values `current_billing_period_start` and `timestamp` on enum `V2BillingIntentActionApplyEffectiveAt.Type`
  * Add support for new value `discount` on enum `V2BillingIntentActionApply.Type`
  * ⚠️ Change type of `V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior.Type`, `V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams.Type`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior.Type`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams.Type`, `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehavior.Type`, and `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsOverridesPartialPeriodBehaviorParams.Type` from `literal('license_fee')` to `enum('license_fee'|'recurring_credit_grant')`
  * Add support for `ServiceCycle` on `V2BillingLicenseFee` and `V2BillingRateCard`
  * ⚠️ Remove support for `LatestVersion` on `V2BillingLicenseFee`, `V2BillingPricingPlan`, and `V2BillingRateCard`
  * ⚠️ Remove support for `ServiceIntervalCount` and `ServiceInterval` on `V2BillingLicenseFee` and `V2BillingRateCard`
  * ⚠️ Change type of `V2BillingLicenseFeeTransformQuantity.DivideBy`, `V2BillingLicenseFeeTransformQuantityParams.DivideBy`, `V2BillingLicenseFeeVersionTransformQuantity.DivideBy`, `V2BillingRateCardRateTransformQuantity.DivideBy`, and `V2BillingRateCardRateTransformQuantityParams.DivideBy` from `longInteger` to `int64_string`
  * Add support for `DiscountDetails` and `PricingPlanComponentDetails` on `V2BillingPricingPlanSubscription`
  * Add support for new value `crypto_wallets` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * ⚠️ Remove support for value `crypto` from enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for `BalanceByFundsType` on `V2MoneyManagementFinancialAccountPayments`
  * Add support for new value `next_day_payout_fee` on enum `V2MoneyManagementOutboundPaymentQuoteEstimatedFee.Type`
  * Add support for `TreasuryTransactionEntry` on `V2MoneyManagementTransactionEntry`
  * Add support for `TreasuryCreditReversal`, `TreasuryDebitReversal`, `TreasuryInboundTransfer`, `TreasuryIssuingAuthorization`, `TreasuryOutboundPayment`, `TreasuryOutboundTransfer`, `TreasuryReceivedCredit`, and `TreasuryReceivedDebit` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
  * Add support for new values `treasury_credit_reversal`, `treasury_debit_reversal`, `treasury_inbound_transfer`, `treasury_issuing_authorization`, `treasury_other`, `treasury_outbound_payment`, `treasury_outbound_transfer`, `treasury_received_credit`, and `treasury_received_debit` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
  * Add support for `TreasuryTransaction` on `V2MoneyManagementTransaction`
  * Add support for new value `no_valid_payment_method` on enum `V2PaymentsOffSessionPayment.FailureReason`
  * Add support for `Metadata` on `V2PaymentsSettlementAllocationIntentSplit`
  * ⚠️ Change type of `V2ReportingReportRunResultFile.Size` from `longInteger` to `int64_string`
  * Add support for `StatementDescriptor` on `V2MoneyManagementOutboundPaymentParams` and `V2MoneyManagementOutboundTransferParams`
  * Add support for `Include` on `V2BillingIntentParams`, `V2BillingIntentReserveParams`, `V2BillingPricingPlanSubscriptionListParams`, `V2BillingPricingPlanSubscriptionParams`, `V2MoneyManagementFinancialAccountListParams`, and `V2MoneyManagementFinancialAccountParams`
  * Add support for event notifications `V1AccountSignalsIncludingDelinquencyCreatedEvent`, `V2CoreAccountSignalsFraudulentWebsiteReadyEvent`, and `V2SignalsAccountSignalFraudulentMerchantReadyEvent`

## 84.5.0-alpha.4 - 2026-03-18
* ⚠️ [#2296](https://github.com/stripe/stripe-go/pull/2296) Update generated code for private-preview
  * Add support for new resources `OrchestrationPaymentAttempt` and `RadarCustomerEvaluation`
  * Add support for `Get` method on resource `OrchestrationPaymentAttempt`
  * Add support for `New` and `Update` methods on resource `RadarCustomerEvaluation`
  * Add support for `Approve` method on resource `CheckoutSession`
  * Add support for `ReportAuthenticated`, `ReportCanceled`, `ReportFailed`, `ReportGuaranteed`, `ReportInformational`, and `ReportRefund` methods on resource `PaymentAttemptRecord`
  * Add support for `SimulateCryptoDeposit` test helper method on resource `PaymentIntent`
  * Add support for `CreateUSPaperCheckOnApplication` on `AccountSessionComponentsCheckScanningFeaturesParams`
  * Add support for `ApprovalMethod` on `CheckoutSessionParams` and `CheckoutSession`
  * Add support for `CurrentAttempt` on `CheckoutSession`
  * Add support for `SelectedFulfillmentOptionOverrides` on `DelegatedCheckoutRequestedSessionFulfillmentDetailsParams`
  * Add support for `PricingPlanSubscriptionDetails` on `InvoiceItemParent` and `InvoiceLineItemParent`
  * ⚠️ Remove support for `LicenseFeeSubscriptionDetails` on `InvoiceItemParent` and `InvoiceLineItemParent`
  * ⚠️ Remove support for `PricingPlanSubscription` and `PricingPlanVersion` on `InvoiceItemParentRateCardSubscriptionDetails` and `InvoiceLineItemParentRateCardSubscriptionDetails`
  * Add support for new value `pricing_plan_subscription_details` on enum `InvoiceItemParent.Type`
  * ⚠️ Remove support for value `license_fee_subscription_details` from enum `InvoiceItemParent.Type`
  * Add support for new value `discounts` on enum `InvoiceItem.FrozenFields`
  * Add support for new value `pricing_plan_subscription_details` on enum `InvoiceLineItemParent.Type`
  * ⚠️ Remove support for value `license_fee_subscription_details` from enum `InvoiceLineItemParent.Type`
  * Add support for `TokenDetails` on `IssuingAuthorization`
  * Add support for `DepositOptions` and `Mode` on `PaymentIntentConfirmPaymentMethodOptionsCryptoParams`, `PaymentIntentPaymentMethodOptionsCryptoParams`, and `PaymentIntentPaymentMethodOptionsCrypto`
  * Add support for `CryptoDisplayDetails` on `PaymentIntentNextAction`
  * Add support for `FailureCode` on `PaymentRecordReportPaymentAttemptFailedParams` and `PaymentRecordReportPaymentFailedParams`
  * Add support for `RecurringInterval` on `SharedPaymentGrantedTokenUsageLimitsParams`
  * Add support for `HomeRuleTax` on `TaxRegistrationCountryOptionsUsParams` and `TaxRegistrationCountryOptionsUs`
  * Add support for new value `home_rule_tax` on enum `TaxRegistrationCountryOptionsUs.Type`

## 84.5.0-alpha.3 - 2026-03-11
* ⚠️ [#2289](https://github.com/stripe/stripe-go/pull/2289) Update generated code for private-preview
  * Add support for new resource `RadarIssuingAuthorizationEvaluation`
  * Add support for `New` method on resource `RadarIssuingAuthorizationEvaluation`
  * Add support for new value `fee_credits` on enum `BalanceTransaction.BalanceType`
  * ⚠️ Rename `AffiliateAttributions` to `AffiliateAttribution` on `DelegatedCheckoutRequestedSessionConfirmParams` and `DelegatedCheckoutRequestedSessionParams`
  * Add support for `AmountToCounter` on `Dispute`
  * Add support for `FrozenFields` on `InvoiceItem`
  * Add support for new value `next_billing_period_start` on enum `V2BillingIntentActionApplyEffectiveAt.Type`
  * Add support for `Consumer` on `V2CoreAccountConfigurationCardCreatorCapabilitiesParams`, `V2CoreAccountConfigurationCardCreatorCapabilities`, `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorParams`, and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreator`
  * Add support for `FifthThird` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercial`, `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialParams`, and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercial`
  * Add support for `PrepaidCard` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBank`, `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBankParams`, and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialCrossRiverBank`
  * Add support for new values `commercial.cross_river_bank.prepaid_card`, `commercial.fifth_third.charge_card`, `consumer.celtic.revolving_credit_card`, `consumer.cross_river_bank.prepaid_card`, and `consumer.lead.prepaid_card` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for `PaymentMethodData` on `V2PaymentsOffSessionPaymentParams`
  * Add support for new values `commercial.cross_river_bank.prepaid_card`, `commercial.fifth_third.charge_card`, `consumer.celtic.revolving_credit_card`, `consumer.cross_river_bank.prepaid_card`, and `consumer.lead.prepaid_card` on enum `EventsV2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEvent.UpdatedCapability`

## 84.5.0-alpha.2 - 2026-03-04
This release changes the pinned API version to `2026-03-04.preview`.

* ⚠️ [#2283](https://github.com/stripe/stripe-go/pull/2283) Update generated code for private-preview
  * Add support for new resources `BillingAlertRecovered` and `Profile`
  * Add support for `Reauthorize` method on resource `PaymentIntent`
  * Add support for `Settings` on `QuoteLineActionAddDiscount`, `QuoteLineActionAddItemDiscount`, `QuoteLineActionSetDiscounts`, `QuoteLineActionSetItemsDiscount`, `QuotePreviewSubscriptionSchedulePhaseDiscount`, `QuotePreviewSubscriptionSchedulePhaseItemDiscount`, `SubscriptionSchedulePhaseDiscount`, and `SubscriptionSchedulePhaseItemDiscount`
  * Add support for `SmartDisputes` on `AccountSettingsParams`, `AccountSettings`, `V2CoreAccountConfigurationMerchantParams`, and `V2CoreAccountConfigurationMerchant`
  * Add support for `EmailCustomersOnSuccessfulPayment` on `AccountSettingsPaymentsParams` and `AccountSettingsPayments`
  * Add support for `BalanceUpdateDetails` on `BillingCreditBalanceSummaryBalance`
  * Add support for `Reauthorization` and `ReauthorizeBefore` on `ChargePaymentMethodDetailsCardPresent`, `ChargePaymentMethodDetailsCard`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsCardPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, and `PaymentRecordPaymentMethodDetailsCardPresent`
  * Add support for `Location` and `Reader` on `ChargePaymentMethodDetailsCardPresent`, `ChargePaymentMethodDetailsInteracPresent`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsInteracPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentRecordPaymentMethodDetailsCardPresent`, and `PaymentRecordPaymentMethodDetailsInteracPresent`
  * Add support for `ManagedPayments` on `CheckoutSessionParams`, `CheckoutSession`, `PaymentIntent`, `SetupIntent`, and `Subscription`
  * Add support for new value `lk_vat` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `CheckoutSessionCustomerDetailsTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, `QuotePreviewInvoiceCustomerTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, `TaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
  * Add support for `Digital` on `DelegatedCheckoutRequestedSessionFulfillmentDetailsFulfillmentOptions`, `DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOptionParams`, and `DelegatedCheckoutRequestedSessionFulfillmentDetailsSelectedFulfillmentOption`
  * Add support for `AffiliateAttributions` on `DelegatedCheckoutRequestedSessionConfirmParams`, `DelegatedCheckoutRequestedSessionParams`, and `DelegatedCheckoutRequestedSession`
  * Add support for `FulfillmentType` on `DelegatedCheckoutRequestedSessionLineItemDetail`
  * Add support for `MarketplaceSellerDetails`, `NetworkProfile`, `PrivacyNoticeURL`, `ReturnPolicyURL`, `StorePolicyURL`, and `TermsOfServiceURL` on `DelegatedCheckoutRequestedSessionSellerDetails`
  * Add support for `AmountToCounter` on `DisputeParams`
  * Add support for new values `reserve.hold.created`, `reserve.hold.updated`, `reserve.plan.created`, `reserve.plan.disabled`, `reserve.plan.expired`, `reserve.plan.updated`, and `reserve.release.created` on enum `Event.Type`
  * Add support for new values `terminal_wifi_certificate` and `terminal_wifi_private_key` on enum `File.Purpose`
  * Add support for new value `pay_by_bank` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
  * Add support for `DisplayName` and `ServiceUserNumber` on `MandatePaymentMethodDetailsBacsDebit`
  * Add support for `RequestReauthorization` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardPresent`, and `PaymentIntentPaymentMethodOptionsCard`
  * Add support for `TransactionPurpose` on `PaymentIntentConfirmPaymentMethodOptionsUsBankAccountParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountParams`, and `PaymentIntentPaymentMethodOptionsUsBankAccount`
  * Add support for new value `requires_reauthorization` on enum `PaymentIntent.Status`
  * Add support for `OptionalItems` on `PaymentLinkParams`
  * Add support for new value `billing_schedules_invalid` on enum `QuoteStatusDetailsStaleLastReason.Type`
  * ⚠️ Remove support for `CardIssuerDecline` on `RadarPaymentEvaluationInsights`
  * Add support for `PaymentBehavior` on `SubscriptionItemParams`
  * Add support for `BillingCycleAnchor` on `SubscriptionTrialSettingsEndBehavior`
  * Add support for `Lk` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
  * Add support for `Cellular` and `StripeS710` on `TerminalConfigurationParams` and `TerminalConfiguration`
  * Add support for new values `simulated_stripe_s710` and `stripe_s710` on enum `TerminalReader.DeviceType`
  * Add support for new values `ar_bank_account`, `bt_bank_account`, `co_bank_account`, `cr_bank_account`, `do_bank_account`, `gt_bank_account`, `md_bank_account`, `mk_bank_account`, `mo_bank_account`, `mz_bank_account`, `pe_bank_account`, `pk_bank_account`, `tw_bank_account`, and `uz_bank_account` on enums `V2AccountConfigurationRecipientDataDefaultOutboundDestination.Type` and `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * Add support for `RecipientOnboarding` and `RecipientUpdate` on `V2CoreAccountLinkUseCaseParams` and `V2CoreAccountLinkUseCase`
  * Add support for new values `recipient_onboarding` and `recipient_update` on enum `V2CoreAccountLinkUseCase.Type`
  * Add support for `Consumer` on `V2CoreAccountConfigurationStorerCapabilitiesParams` and `V2CoreAccountConfigurationStorerCapabilities`
  * Add support for new value `consumer.holds_currencies.usd` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for `FundsUsageType` on `V2MoneyManagementFinancialAccountStorageParams` and `V2MoneyManagementFinancialAccountStorage`
  * Add support for `Purpose` on `V2MoneyManagementOutboundPaymentParams` and `V2MoneyManagementOutboundPayment`
  * Add support for `BranchNumber` and `SwiftCode` on `V2MoneyManagementPayoutMethodBankAccount`
  * Add support for new values `dispute`, `inbound_payment_failure`, `inbound_payment`, `india_mdr_processing_fee`, `payment_method_passthrough_fee`, `refund`, and `tax_withholding` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
  * ⚠️ Remove support for values `charge_failure` and `charge` from enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
  * Add support for new value `consumer.holds_currencies.usd` on enum `EventsV2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent.UpdatedCapability`
  * Add support for snapshot event `EventTypeBillingAlertRecovered` with resource `BillingAlertRecovered`
  * Add support for snapshot events `EventTypeReserveHoldCreated` and `EventTypeReserveHoldUpdated` with resource `ReserveHold`
  * Add support for snapshot events `EventTypeReservePlanCreated`, `EventTypeReservePlanDisabled`, `EventTypeReservePlanExpired`, and `EventTypeReservePlanUpdated` with resource `ReservePlan`
  * Add support for snapshot event `EventTypeReserveReleaseCreated` with resource `ReserveRelease`
  * Add support for event notification `V2BillingRateCardCustomPricingUnitOverageRateCreatedEvent` with related object `V2BillingRateCardCustomPricingUnitOverageRate`
  * Add support for event notifications `V2IamStripeAccessGrantApprovedEvent`, `V2IamStripeAccessGrantCanceledEvent`, `V2IamStripeAccessGrantDeniedEvent`, `V2IamStripeAccessGrantRemovedEvent`, `V2IamStripeAccessGrantRequestedEvent`, and `V2IamStripeAccessGrantUpdatedEvent`
  * Add support for error codes `storer_capability_missing` and `storer_capability_not_active` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `QuotePreviewInvoiceLastFinalizationError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`

## 84.5.0-alpha.1 - 2026-02-25
This release uses the API version `2026-01-28.preview`.

* [#2275](https://github.com/stripe/stripe-go/pull/2275) Update generated code for private-preview
  * Add support for new resource `AccountSignals`
  * Add support for `Get` method on resource `AccountSignals`
  * Add support for `AggregationPeriod`, `GroupBy`, and `TriggeredAt` on `BillingAlertTriggered`
  * Add support for `ExternalAccountCollection` on `AccountLinkCollectionOptionsParams`
  * Add support for `FundingSource` on `ApplicationFee`
  * Add support for `Hosted` and `UIMode` on `FinancialConnectionsSessionParams` and `FinancialConnectionsSession`
  * Add support for `URL` on `FinancialConnectionsSession`
  * Add support for `BillingCycleAnchor` on `SubscriptionTrialSettingsEndBehaviorParams`

## 84.4.0-alpha.4 - 2026-02-19
* ⚠️ [#2268](https://github.com/stripe/stripe-go/pull/2268) Update generated code for private-preview
  * Add support for `SpendThreshold` on `BillingAlertParams` and `BillingAlert`
  * ⚠️ Add support for new value `spend_threshold` on enum `BillingAlert.AlertType`
  * Add support for `InvoiceItem`, `ProrationDetails`, `Proration`, and `Subscription` on `InvoiceLineItemParentScheduleDetails`
  * Add support for `Custom` on `PaymentMethodParams`
  * Add support for `PaymentMethodReference` and `Usage` on `PaymentMethodCustom`
  * ⚠️ Change type of `QuoteSubscriptionDataOverridesParams.BillingSchedules` from `emptyable(array(billing_schedules_update_specs))` to `array(billing_schedules_update_specs)`
  * Add support for `OutstandingUsageThrough` and `UnusedTimeFrom` on `SubscriptionPauseBillForParams`
  * ⚠️ Remove support for `OutstandingUsage` and `UnusedTime` on `SubscriptionPauseBillForParams`
  * ⚠️ Remove support for `PaymentBehavior` on `SubscriptionResumeParams`

## 84.4.0-alpha.3 - 2026-02-11
* [#2264](https://github.com/stripe/stripe-go/pull/2264) Update generated code for private-preview
  * Add support for new resources `V2BillingCadenceSpendModifier`, `V2BillingOneTimeItem`, and `V2BillingRateCardCustomPricingUnitOverageRate`
  * Add support for `Del`, `Get`, `List`, and `New` methods on resource `V2BillingRateCardCustomPricingUnitOverageRate`
  * Add support for `Get`, `List`, `New`, and `Update` methods on resource `V2BillingOneTimeItem`
  * Add support for `Get` method on resource `V2BillingCadenceSpendModifier`
  * Add support for `SettlementType` on `ApplicationFee`
  * Add support for `RateCardCustomPricingUnitOverageRateDetails` on `InvoiceItemPricing` and `InvoiceLineItemPricing`
  * Add support for new value `rate_card_custom_pricing_unit_overage_rate_details` on enums `InvoiceItemPricing.Type` and `InvoiceLineItemPricing.Type`
  * Add support for `DefaultSettings` on `InvoiceCreatePreviewScheduleDetailsParams`
  * Add support for `PaymentBehavior` on `SubscriptionResumeParams`
  * Add support for `EffectiveAt` and `SpendModifierRule` on `V2BillingIntentActionApplyParams`, `V2BillingIntentActionApply`, `V2BillingIntentActionRemoveParams`, and `V2BillingIntentActionRemove`
  * Change type of `V2BillingIntentActionApply.Type`, `V2BillingIntentActionApplyParams.Type`, `V2BillingIntentActionRemove.Type`, and `V2BillingIntentActionRemoveParams.Type` from `literal('invoice_discount_rule')` to `enum('invoice_discount_rule'|'spend_modifier_rule')`

## 84.4.0-alpha.2 - 2026-02-04
* [#2262](https://github.com/stripe/stripe-go/pull/2262) Update generated code for private-preview
  * Add support for new resource `V2CoreConnectionSession`
  * Add support for `Get` and `New` methods on resource `V2CoreConnectionSession`
  * Add support for `List` method on resources `V2PaymentsSettlementAllocationIntentSplit` and `V2PaymentsSettlementAllocationIntent`
  * Add support for `AgenticCommerceSettings` on `AccountSessionComponentsParams`
  * Add support for `TerminalHardwareOrders` and `TerminalHardwareShop` on `AccountSessionComponentsParams` and `AccountSessionComponents`
  * Add support for `NetworkCostPassthroughReport` on `AccountSessionComponents`
  * Add support for new values `ae_bank_account`, `ag_bank_account`, `bh_bank_account`, `gm_bank_account`, `hk_bank_account`, `kh_bank_account`, `lc_bank_account`, `mc_bank_account`, `mg_bank_account`, `my_bank_account`, `qa_bank_account`, `rw_bank_account`, `th_bank_account`, `tt_bank_account`, and `vn_bank_account` on enums `V2AccountConfigurationRecipientDataDefaultOutboundDestination.Type` and `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
  * Add support for `CadenceData` on `V2BillingIntentParams` and `V2BillingIntent`
  * Add support for `CancellationDetails` on `V2BillingIntentActionDeactivateParams`, `V2BillingIntentActionDeactivate`, and `V2BillingPricingPlanSubscription`
  * Add support for `ContactPhone` on `V2CoreAccountParams`, `V2CoreAccountTokenParams`, and `V2CoreAccount`
  * Add support for `RegistrationDate` on `V2CoreAccountIdentityBusinessDetailsParams`, `V2CoreAccountIdentityBusinessDetails`, and `V2CoreAccountTokenIdentityBusinessDetailsParams`
  * Add support for new value `gb_vat` on enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
  * Add support for `Reference` on `V2MoneyManagementAdjustment`
  * Add support for `AccruedFees` on `V2MoneyManagementFinancialAccount`
  * Add support for `StartingBalance` on `V2MoneyManagementFinancialAccountPayments`
  * Add support for new value `accrued_fees` on enum `V2MoneyManagementFinancialAccount.Type`
  * Add support for `AccountHolderAddress` and `AccountHolderName` on `V2MoneyManagementFinancialAddressCredentialsUsBankAccount`
  * Add support for `Fingerprint` on `V2MoneyManagementPayoutMethodCard`
  * Add support for `CardSpend` on `V2MoneyManagementReceivedCredit` and `V2MoneyManagementReceivedDebit`
  * Add support for new value `card_spend` on enum `V2MoneyManagementReceivedCredit.Type`
  * Add support for new value `card_spend` on enum `V2MoneyManagementReceivedDebit.Type`
  * Add support for new values `advance`, `anticipation_repayment`, `balance_transfer`, `charge_failure`, `charge`, `climate_order_purchase`, `climate_order_refund`, `connect_collection_transfer`, `connect_reserved_funds`, `contribution`, `dispute_reversal`, `financing_paydown_reversal`, `financing_paydown`, `inbound_transfer_reversal`, `issuing_dispute_fraud_liability_debit`, `issuing_dispute_provisional_credit_reversal`, `issuing_dispute_provisional_credit`, `issuing_dispute`, `minimum_balance_hold`, `network_cost`, `obligation`, `outbound_payment_reversal`, `outbound_transfer_reversal`, `partial_capture_reversal`, `payment_network_reserved_funds`, `platform_earning_refund`, `platform_earning`, `platform_fee`, `received_credit_reversal`, `received_debit_reversal`, `refund_failure`, `risk_reserved_funds`, `stripe_balance_payment_debit_reversal`, `stripe_balance_payment_debit`, `stripe_fee_tax`, `transfer_reversal`, and `unreconciled_customer_funds` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
  * Add support for `ApplicationFeeRefund`, `ApplicationFee`, `Charge`, `Dispute`, `Payout`, `Refund`, `ReserveHold`, `ReserveRelease`, `Topup`, `TransferReversal`, and `Transfer` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
  * Add support for new values `application_fee_refund`, `application_fee`, `charge`, `dispute`, `payout`, `refund`, `reserve_hold`, `reserve_release`, `topup`, `transfer_reversal`, and `transfer` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
  * Add support for error codes `blocked_payout_method` and `unsupported_payout_method` on `BlockedByStripeError`
  * Add support for error code `invalid_payout_method_data` on `InvalidPayoutMethodError`
  * Add support for error code `limit_payout_method` on `QuotaExceededError`

## 84.4.0-alpha.1 - 2026-01-28
This release changes the pinned API version to `2026-01-28.preview`.

* [#2259](https://github.com/stripe/stripe-go/pull/2259) Update generated code for private-preview
  * Add support for new resources `FRMealVouchersOnboarding`, `ReserveHold`, `ReservePlan`, and `ReserveRelease`
  * Add support for `Get`, `List`, `New`, and `Update` methods on resource `FRMealVouchersOnboarding`
  * Add support for `Get` and `List` methods on resources `ReserveHold` and `ReserveRelease`
  * Add support for `Get` method on resource `ReservePlan`
  * Add support for `Pause` method on resource `Subscription`
  * Add support for `ServicePeriodDetails` on `Discount`
  * Add support for `AgenticCommerceSettings` on `AccountSessionComponents`
  * Add support for new value `risk_reserved` on enum `BalanceTransaction.BalanceType`
  * Add support for `ServicePeriod` on `CouponParams` and `Coupon`
  * Add support for new value `service_period` on enum `Coupon.Duration`
  * Change type of `InvoiceItemPricingPriceDetails.Price` and `InvoiceLineItemPricingPriceDetails.Price` from `string` to `expandable($Price)`
  * Add support for `Settings` on `InvoiceCreatePreviewDiscountsParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentDiscountActionAddParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentDiscountActionSetParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionAddDiscountParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionSetDiscountParams`, `InvoiceCreatePreviewScheduleDetailsPhaseDiscountsParams`, `InvoiceCreatePreviewScheduleDetailsPhaseItemDiscountsParams`, `InvoiceCreatePreviewSubscriptionDetailsItemDiscountsParams`, `QuoteLineActionAddDiscountParams`, `QuoteLineActionAddItemDiscountParams`, `QuoteLineActionSetDiscountParams`, `QuoteLineActionSetItemDiscountParams`, `SubscriptionDiscountsParams`, `SubscriptionItemDiscountsParams`, `SubscriptionScheduleAmendAmendmentDiscountActionAddParams`, `SubscriptionScheduleAmendAmendmentDiscountActionSetParams`, `SubscriptionScheduleAmendAmendmentItemActionAddDiscountParams`, `SubscriptionScheduleAmendAmendmentItemActionSetDiscountParams`, `SubscriptionSchedulePhaseDiscountsParams`, and `SubscriptionSchedulePhaseItemDiscountsParams`
  * Add support for `Subtotal` on `InvoiceLineItem`
  * Add support for `BillingCadence` on `SubscriptionListParams`

## 84.3.0-alpha.1 - 2026-01-21
* [#2257](https://github.com/stripe/stripe-go/pull/2257) Update generated code for private-preview
  * Remove support for `Pause` method on resource `Subscription`

## 84.2.0-alpha.3 - 2026-01-14
* [#2250](https://github.com/stripe/stripe-go/pull/2250) Update generated code for private-preview
  * Add support for `RiskDetails` on `DelegatedCheckoutRequestedSession`
  * Remove support for `Description`, `Images`, and `Name` on `DelegatedCheckoutRequestedSessionLineItemDetail`
  * Add support for `Name` on `ProductCatalogTrialOfferParams` and `ProductCatalogTrialOffer`
  * Add support for `LoginFailed` and `RegistrationFailed` on `RadarAccountEvaluationEvents` and `RadarAccountEvaluationParams`
  * Change type of `RadarAccountEvaluationParams.Type` from `literal('registration_succeeded')` to `enum('login_failed'|'login_succeeded'|'registration_failed'|'registration_succeeded')`

## 84.2.0-alpha.2 - 2026-01-07
* [#2236](https://github.com/stripe/stripe-go/pull/2236) Update generated code for private-preview
  * Add support for new resource `TaxLocation`
  * Add support for `Get`, `List`, and `New` methods on resource `TaxLocation`
  * Add support for `Pause` method on resource `Subscription`
  * Add support for `PerformanceLocation` on `CheckoutSessionLineItemPriceDataProductDataTaxDetailsParams`, `InvoiceAddLinesLinePriceDataProductDataTaxDetailsParams`, `InvoiceLineItemPriceDataProductDataTaxDetailsParams`, `InvoiceUpdateLinesLinePriceDataProductDataTaxDetailsParams`, `PaymentLinkLineItemPriceDataProductDataTaxDetailsParams`, `ProductTaxDetailsParams`, `TaxCalculationLineItemParams`, and `TaxCalculationLineItem`
  * Add support for new value `performance` on enums `TaxCalculationLineItemTaxBreakdown.Sourcing`, `TaxCalculationShippingCostTaxBreakdown.Sourcing`, and `TaxTransactionShippingCostTaxBreakdown.Sourcing`
  * Add support for new values `admissions_tax`, `attendance_tax`, `entertainment_tax`, `gross_receipts_tax`, `hospitality_tax`, `luxury_tax`, `resort_tax`, and `tourism_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationTaxBreakdownTaxRateDetails.TaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetails.TaxType`
  * Change type of `DelegatedCheckoutRequestedSessionParams.Metadata` from `map(string: string)` to `emptyable(map(string: string))`
  * Change type of `DelegatedCheckoutRequestedSessionParams.PaymentMethodData` from `payment_method_data` to `emptyable(payment_method_data)`
  * Change type of `DelegatedCheckoutRequestedSessionParams.SharedMetadata` from `map(string: string)` to `emptyable(map(string: string))`
  * Add support for `Subscription` on `InvoiceParentScheduleDetails` and `QuotePreviewInvoiceParentScheduleDetails`
  * Change type of `PaymentIntentConfirmPaymentDetailsBenefitParams.FRMealVoucher`, `PaymentIntentPaymentDetailsBenefitParams.FRMealVoucher`, `SetupIntentConfirmSetupDetailsBenefitParams.FRMealVoucher`, and `SetupIntentSetupDetailsBenefitParams.FRMealVoucher` from `payment_details_benefit_fr_meal_voucher` to `emptyable(payment_details_benefit_fr_meal_voucher)`
  * Add support for `TaxDetails` on `PlanProductParams` and `PriceProductDataParams`
  * Add support for `ExternalReference` on `Plan` and `Price`
  * Add support for new value `phase_start` on enums `QuoteSubscriptionData.PhaseEffectiveAt` and `QuoteSubscriptionDataOverrides.PhaseEffectiveAt`
  * Remove support for value `line_start` from enums `QuoteSubscriptionData.PhaseEffectiveAt` and `QuoteSubscriptionDataOverrides.PhaseEffectiveAt`
  * Add support for `AdmissionsTax`, `AttendanceTax`, `EntertainmentTax`, `GrossReceiptsTax`, `HospitalityTax`, `LuxuryTax`, `ResortTax`, and `TourismTax` on `TaxRegistrationCountryOptionsUs`
  * Add support for new values `admissions_tax`, `attendance_tax`, `entertainment_tax`, `gross_receipts_tax`, `hospitality_tax`, `luxury_tax`, `resort_tax`, and `tourism_tax` on enum `TaxRegistrationCountryOptionsUs.Type`
  * Add support for `Requirements` on `TaxCode`
* [#2245](https://github.com/stripe/stripe-go/pull/2245) Update generated code for private-preview
  * Add support for `TrackingDetails` on `V2MoneyManagementOutboundPayment`
  * Add support for `PaperCheck` on `V2MoneyManagementOutboundPaymentDeliveryOptionsParams` and `V2MoneyManagementOutboundPaymentDeliveryOptions`
  * Add support for event notification `V2CoreAccountIncludingFutureRequirementsUpdatedEvent` with related object `V2CoreAccount`
  * Add support for error code `account_rate_limit_exceeded` on `RateLimitError`

## 84.2.0-alpha.1 - 2025-12-14
This release changes the pinned API version to `2025-12-15.preview`.

* [#2233](https://github.com/stripe/stripe-go/pull/2233) Update generated code for private-preview
  * Add support for new resources `SharedPaymentGrantedToken`, `V2IamAPIKey`, `V2PaymentsSettlementAllocationIntentSplit`, `V2PaymentsSettlementAllocationIntent`, and `V2TaxManualRule`
  * Add support for `Get` method on resource `SharedPaymentGrantedToken`
  * Add support for `New` and `Update` test helper methods on resource `SharedPaymentGrantedToken`
  * Add support for `Deactivate`, `Get`, `List`, `New`, and `Update` methods on resource `V2TaxManualRule`
  * Add support for `Cancel`, `Get`, `New`, `Submit`, and `Update` methods on resource `V2PaymentsSettlementAllocationIntent`
  * Add support for `Cancel`, `Get`, and `New` methods on resource `V2PaymentsSettlementAllocationIntentSplit`
  * Add support for `Expire`, `Get`, `List`, `New`, `Rotate`, and `Update` methods on resource `V2IamAPIKey`
  * Add support for `CheckScanning` on `AccountSessionComponentsParams`
  * Add support for `TaxDetails` on `CheckoutSessionLineItemPriceDataProductDataParams`, `InvoiceAddLinesLinePriceDataProductDataParams`, `InvoiceLineItemPriceDataProductDataParams`, `InvoiceUpdateLinesLinePriceDataProductDataParams`, `PaymentLinkLineItemPriceDataProductDataParams`, and `ProductParams`
  * Add support for `PaymentMethodData` on `DelegatedCheckoutRequestedSessionConfirmParams`
  * Add support for `ProductDetails` on `DelegatedCheckoutRequestedSessionLineItemDetail`
  * Add support for `Wallets` on `IssuingCardListParams`
  * Add support for `PrimaryAccountIdentifier` on `IssuingCardWalletsApplePay` and `IssuingCardWalletsGooglePay`
  * Add support for `SharedPaymentGrantedToken` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, and `PaymentIntent`
  * Add support for new values `al_bank_account`, `am_bank_account`, `bn_bank_account`, `bw_bank_account`, `dz_bank_account`, `gy_bank_account`, `jm_bank_account`, `jo_bank_account`, `kw_bank_account`, `lk_bank_account`, `ma_bank_account`, `om_bank_account`, and `tz_bank_account` on enum `V2AccountConfigurationRecipientDataDefaultOutboundDestination.Type`
  * Add support for `Instant` on `V2AccountConfigurationRecipientDataFeaturesBankAccountsParams`, `V2AccountConfigurationRecipientDataFeaturesBankAccounts`, `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsParams`, and `V2CoreAccountConfigurationRecipientCapabilitiesBankAccounts`
  * Add support for new value `bank_accounts.instant` on enum `V2AccountRequirementImpact.RequiredForFeatures`
  * Add support for `CollectAt` on `V2BillingIntentActionDeactivateParams`, `V2BillingIntentActionDeactivate`, `V2BillingIntentActionModifyParams`, `V2BillingIntentActionModify`, `V2BillingIntentActionSubscribeParams`, and `V2BillingIntentActionSubscribe`
  * Remove support for `BillingDetails` on `V2BillingIntentActionDeactivateParams`, `V2BillingIntentActionDeactivate`, `V2BillingIntentActionModifyParams`, `V2BillingIntentActionModify`, `V2BillingIntentActionSubscribeParams`, and `V2BillingIntentActionSubscribe`
  * Add support for `Overrides` on `V2BillingIntentActionDeactivatePricingPlanSubscriptionDetailsParams`, `V2BillingIntentActionDeactivatePricingPlanSubscriptionDetails`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetailsParams`, `V2BillingIntentActionModifyPricingPlanSubscriptionDetails`, `V2BillingIntentActionSubscribePricingPlanSubscriptionDetailsParams`, and `V2BillingIntentActionSubscribePricingPlanSubscriptionDetails`
  * Remove support for `Requested` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticChargeCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCelticSpendCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankChargeCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialCrossRiverBankSpendCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialLeadPrepaidCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeChargeCard`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripePrepaidCard`, `V2CoreAccountConfigurationRecipientCapabilitiesCryptoWallets`, `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesCryptoWallets`, `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesUsdc`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsCryptoWallets`, and `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersCryptoWallets`
  * Add support for new value `bank_accounts.instant` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for `AlternativeReference` on `V2CoreVaultGbBankAccount`, `V2CoreVaultUsBankAccount`, and `V2MoneyManagementPayoutMethod`
  * Add support for `ManagedBy` and `Payments` on `V2MoneyManagementFinancialAccount`
  * Add support for new value `payments` on enum `V2MoneyManagementFinancialAccount.Type`
  * Add support for `Speed` on `V2MoneyManagementOutboundPaymentDeliveryOptionsParams`, `V2MoneyManagementOutboundPaymentDeliveryOptions`, `V2MoneyManagementOutboundPaymentQuoteDeliveryOptionsParams`, and `V2MoneyManagementOutboundPaymentQuoteDeliveryOptions`
  * Add support for new value `real_time_payout_fee` on enum `V2MoneyManagementOutboundPaymentQuoteEstimatedFee.Type`
  * Add support for `Types` on `V2MoneyManagementFinancialAccountListParams`
  * Add support for new value `bank_accounts.instant` on enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`
  * Add support for `TopImpactedAccounts` on `EventsV2CoreHealthApiErrorFiringEventImpact`, `EventsV2CoreHealthApiErrorResolvedEventImpact`, `EventsV2CoreHealthApiLatencyFiringEventImpact`, `EventsV2CoreHealthApiLatencyResolvedEventImpact`, `EventsV2CoreHealthPaymentMethodErrorFiringEventImpact`, and `EventsV2CoreHealthPaymentMethodErrorResolvedEventImpact`
  * Add support for event notifications `V2CoreHealthSepaDebitDelayedFiringEvent`, `V2CoreHealthSepaDebitDelayedResolvedEvent`, and `V2PaymentsSettlementAllocationIntentNotFoundEvent`
  * Add support for event notifications `V2PaymentsSettlementAllocationIntentCanceledEvent`, `V2PaymentsSettlementAllocationIntentCreatedEvent`, `V2PaymentsSettlementAllocationIntentErroredEvent`, `V2PaymentsSettlementAllocationIntentFundsNotReceivedEvent`, `V2PaymentsSettlementAllocationIntentMatchedEvent`, `V2PaymentsSettlementAllocationIntentSettledEvent`, and `V2PaymentsSettlementAllocationIntentSubmittedEvent` with related object `V2PaymentsSettlementAllocationIntent`
  * Add support for event notifications `V2PaymentsSettlementAllocationIntentSplitCanceledEvent`, `V2PaymentsSettlementAllocationIntentSplitCreatedEvent`, and `V2PaymentsSettlementAllocationIntentSplitSettledEvent` with related object `V2PaymentsSettlementAllocationIntentSplit`
  * Remove support for error code `account_rate_limit_exceeded` on `RateLimitError`

## 84.1.0-alpha.4 - 2025-12-04
* [#2230](https://github.com/stripe/stripe-go/pull/2230) Update generated code for private-preview
  * Add support for `CheckScanning` on `AccountSessionComponents`
  * Add support for `Client` on `V2CoreEventReasonRequest`
  * Add support for `StripeBalancePayment` on `V2MoneyManagementReceivedCredit` and `V2MoneyManagementReceivedDebit`
  * Add support for new value `stripe_balance_payment` on enum `V2MoneyManagementReceivedCredit.Type`
  * Add support for `BalanceTransfer` on `V2MoneyManagementReceivedDebit`
  * Add support for new values `balance_transfer` and `stripe_balance_payment` on enum `V2MoneyManagementReceivedDebit.Type`
  * Add support for `Include` on `V2CoreEventListParams` and `V2CoreEventParams`
* [#2231](https://github.com/stripe/stripe-go/pull/2231) Update generated code for private-preview
  * Add support for event notifications `V2IamApiKeyCreatedEvent`, `V2IamApiKeyDefaultSecretRevealedEvent`, `V2IamApiKeyExpiredEvent`, `V2IamApiKeyPermissionsUpdatedEvent`, `V2IamApiKeyRotatedEvent`, and `V2IamApiKeyUpdatedEvent`

## 84.1.0-alpha.3 - 2025-11-24
* [#2224](https://github.com/stripe/stripe-go/pull/2224) Update generated code for private-preview
  * Add support for new resource `ProductCatalogTrialOffer`
  * Add support for `New` method on resource `ProductCatalogTrialOffer`
  * Remove support for `AmountSubtotalAfterDiscount` on `DelegatedCheckoutRequestedSessionLineItemDetail` and `DelegatedCheckoutRequestedSessionTotalDetails`
  * Remove support for `AmountTotal`, `UnitAmountAfterDiscount`, and `UnitDiscount` on `DelegatedCheckoutRequestedSessionLineItemDetail`
  * Add support for `AmountCartDiscount` and `AmountItemsDiscount` on `DelegatedCheckoutRequestedSessionTotalDetails`
  * Remove support for `AmountDiscount` on `DelegatedCheckoutRequestedSessionTotalDetails`
  * Add support for `PaymentsOrchestration` on `PaymentIntentParams` and `PaymentIntent`

## 84.1.0-alpha.2 - 2025-11-20
This release changes the pinned API version to `2025-11-17.preview`.

* [#2221](https://github.com/stripe/stripe-go/pull/2221) Update generated code for private-preview
  * Add support for new resources `V2CoreAccountPersonToken`, `V2CoreAccountToken`, and `V2MoneyManagementCurrencyConversion`
  * Add support for `Get`, `List`, and `New` methods on resource `V2MoneyManagementCurrencyConversion`
  * Add support for `Get` and `New` methods on resources `V2CoreAccountPersonToken` and `V2CoreAccountToken`
  * Add support for `EffectiveAt` on `InvoiceCreatePreviewScheduleDetailsAmendmentParams`, `InvoiceCreatePreviewScheduleDetailsPhaseParams`, `QuoteLineParams`, `QuoteLine`, `QuotePreviewSubscriptionSchedulePhase`, `SubscriptionScheduleAmendAmendmentParams`, `SubscriptionSchedulePhaseParams`, and `SubscriptionSchedulePhase`
  * Add support for `TrialOffer` on `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionAddParams`, `InvoiceCreatePreviewScheduleDetailsAmendmentItemActionSetParams`, `InvoiceCreatePreviewScheduleDetailsPhaseItemParams`, `QuoteLineActionAddItemParams`, `QuoteLineActionAddItem`, `QuoteLineActionSetItemParams`, `QuoteLineActionSetItems`, `QuotePreviewSubscriptionSchedulePhaseItem`, `SubscriptionScheduleAmendAmendmentItemActionAddParams`, `SubscriptionScheduleAmendAmendmentItemActionSetParams`, `SubscriptionSchedulePhaseItemParams`, and `SubscriptionSchedulePhaseItem`
  * Add support for `AmountDiscount`, `AmountSubtotal`, `AmountTotal`, `UnitAmountAfterDiscount`, and `UnitDiscount` on `DelegatedCheckoutRequestedSessionLineItemDetail`
  * Add support for `AmountSubtotalAfterDiscount` on `DelegatedCheckoutRequestedSessionLineItemDetail` and `DelegatedCheckoutRequestedSessionTotalDetails`
  * Change type of `InvoiceCreatePreviewScheduleDetailsParams.BillingSchedules` from `array(billing_schedules_update_params)` to `emptyable(array(billing_schedules_update_params))`
  * Add support for `CurrentTrial` on `InvoiceCreatePreviewSubscriptionDetailsItemParams`, `SubscriptionItemParams`, and `SubscriptionItem`
  * Change type of `QuoteSubscriptionDataOverrideParams.BillingSchedules` and `QuoteSubscriptionDataParams.BillingSchedules` from `emptyable(array(billing_schedules_create_specs))` to `array(billing_schedules_create_specs)`
  * Change type of `QuoteSubscriptionData.BillingSchedules` and `QuoteSubscriptionDataOverrides.BillingSchedules` from `nullable(array(SubscriptionsResourceBillingSchedules))` to `array(QuotesResourceSubscriptionDataBillingSchedules)`
  * Change type of `QuoteSubscriptionData.PhaseEffectiveAt` and `QuoteSubscriptionDataOverrides.PhaseEffectiveAt` from `nullable(enum('billing_period_start'|'phase_start'))` to `enum('billing_period_start'|'line_start')`
  * Change type of `QuotePreviewSubscriptionSchedule.BillingSchedules` and `SubscriptionSchedule.BillingSchedules` from `nullable(array(SubscriptionsResourceBillingSchedules))` to `array(SubscriptionsResourceBillingSchedules)`
  * Remove support for `AmendmentStart`, `LineStartsAt`, and `Relative` on `SubscriptionBillingScheduleBillFrom`
  * Change type of `SubscriptionBillingScheduleBillFrom.Type` from `enum` to `literal('timestamp')`
  * Remove support for `AmendmentEnd` and `LineEndsAt` on `SubscriptionBillingScheduleBillUntil`
  * Remove support for values `amendment_end`, `line_ends_at`, `schedule_end`, and `upcoming_invoice` from enum `SubscriptionBillingScheduleBillUntil.Type`
  * Change type of `V2BillingServiceActionCreditGrantAmount.Monetary`, `V2BillingServiceActionCreditGrantAmountParams.Monetary`, `V2BillingServiceActionCreditGrantPerTenantAmount.Monetary`, and `V2BillingServiceActionCreditGrantPerTenantAmountParams.Monetary` from `amount` to `an object`
  * Add support for `FutureRequirements` on `V2CoreAccount`
  * Add support for `KonbiniPayments` and `ScriptStatementDescriptor` on `V2CoreAccountConfigurationMerchantParams` and `V2CoreAccountConfigurationMerchant`
  * Add support for `EUR` on `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesParams` and `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrencies`
  * Add support for `RequirementsCollector` on `V2CoreAccountDefaultsResponsibilities`
  * Add support for new value `ar_cuit` on enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
  * Add support for new value `ar_dni` on enums `V2CoreAccountIdentityIndividualIdNumber.Type` and `V2CoreAccountPersonIdNumber.Type`
  * Remove support for `Collector` on `V2CoreAccountRequirements`
  * Add support for new value `holds_currencies.eur` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for new values `payment_method` and `person` on enum `V2CoreAccountRequirementsEntryReference.Type`
  * Remove support for value `resource` from enum `V2CoreAccountRequirementsEntryReference.Type`
  * Remove support for value `future_requirements` from enum `V2CoreAccountRequirementsEntryRequestedReason.Code`
  * Remove support for `V1EventID` on `V2CoreEvent`
  * Remove support for `AmountDetails` and `CaptureMethod` on `V2PaymentsOffSessionPaymentParams` and `V2PaymentsOffSessionPayment`
  * Change type of `V2PaymentsOffSessionPayment.AmountCapturable` from `amount` to `an object`
  * Change type of `V2PaymentsOffSessionPayment.AmountRequested` from `amount` to `an object`
  * Change type of `V2PaymentsOffSessionPaymentParams.Amount` from `amount` to `an object`
  * Remove support for `Destination` on `V2PaymentsOffSessionPaymentCaptureTransferDataParams`
  * Add support for `Created` on `V2CoreEventListParams`
  * Remove support for `GTE`, `Gt`, `LT`, and `Lte` on `V2CoreEventListParams`
  * Add support for `AccountToken` on `V2CoreAccountParams`
  * Add support for `PersonToken` on `V2CoreAccountPersonParams`
  * Add support for `ImpactedRequestsPercentage` on `EventsV2CoreHealthApiErrorFiringEventImpact`, `EventsV2CoreHealthApiErrorResolvedEventImpact`, `EventsV2CoreHealthApiLatencyFiringEventImpact`, `EventsV2CoreHealthApiLatencyResolvedEventImpact`, `EventsV2CoreHealthPaymentMethodErrorFiringEventImpact`, and `EventsV2CoreHealthPaymentMethodErrorResolvedEventImpact`
  * Add support for `Context` and `RelatedObject` on `EventsV2CoreHealthEventGenerationFailureResolvedEventImpact`
  * Remove support for `Account`, `Livemode`, `MissingDeliveryAttempts`, and `RelatedObjectID` on `EventsV2CoreHealthEventGenerationFailureResolvedEventImpact`
  * Change type of `EventsV2CoreHealthFraudRateIncreasedEventImpact.RealizedFraudAmount` from `amount` to `an object`
  * Change type of `EventsV2CoreHealthIssuingAuthorizationRequestErrorsFiringEventImpact.ApprovedAmount`, `EventsV2CoreHealthIssuingAuthorizationRequestErrorsResolvedEventImpact.ApprovedAmount`, `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutFiringEventImpact.ApprovedAmount`, and `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEventImpact.ApprovedAmount` from `amount` to `an object`
  * Change type of `EventsV2CoreHealthIssuingAuthorizationRequestErrorsFiringEventImpact.DeclinedAmount`, `EventsV2CoreHealthIssuingAuthorizationRequestErrorsResolvedEventImpact.DeclinedAmount`, `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutFiringEventImpact.DeclinedAmount`, and `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEventImpact.DeclinedAmount` from `amount` to `an object`
  * Add support for thin events `V2PaymentsOffSessionPaymentAttemptFailedEvent` and `V2PaymentsOffSessionPaymentAttemptStartedEvent` with related object `V2PaymentsOffSessionPayment`
  * Remove support for thin event `V1AccountUpdatedEvent` with related object `Account`
  * Remove support for thin events `V1ApplicationFeeCreatedEvent` and `V1ApplicationFeeRefundedEvent` with related object `ApplicationFee`
  * Remove support for thin events `V1BillingPortalConfigurationCreatedEvent` and `V1BillingPortalConfigurationUpdatedEvent` with related object `BillingPortalConfiguration`
  * Remove support for thin event `V1CapabilityUpdatedEvent` with related object `Capability`
  * Remove support for thin events `V1ChargeCapturedEvent`, `V1ChargeExpiredEvent`, `V1ChargeFailedEvent`, `V1ChargePendingEvent`, `V1ChargeRefundedEvent`, `V1ChargeSucceededEvent`, and `V1ChargeUpdatedEvent` with related object `Charge`
  * Remove support for thin events `V1ChargeDisputeClosedEvent`, `V1ChargeDisputeCreatedEvent`, `V1ChargeDisputeFundsReinstatedEvent`, `V1ChargeDisputeFundsWithdrawnEvent`, and `V1ChargeDisputeUpdatedEvent` with related object `Dispute`
  * Remove support for thin events `V1ChargeRefundUpdatedEvent`, `V1RefundCreatedEvent`, `V1RefundFailedEvent`, and `V1RefundUpdatedEvent` with related object `Refund`
  * Remove support for thin events `V1CheckoutSessionAsyncPaymentFailedEvent`, `V1CheckoutSessionAsyncPaymentSucceededEvent`, `V1CheckoutSessionCompletedEvent`, and `V1CheckoutSessionExpiredEvent` with related object `CheckoutSession`
  * Remove support for thin events `V1ClimateOrderCanceledEvent`, `V1ClimateOrderCreatedEvent`, `V1ClimateOrderDelayedEvent`, `V1ClimateOrderDeliveredEvent`, and `V1ClimateOrderProductSubstitutedEvent` with related object `ClimateOrder`
  * Remove support for thin events `V1ClimateProductCreatedEvent` and `V1ClimateProductPricingUpdatedEvent` with related object `ClimateProduct`
  * Remove support for thin events `V1CouponCreatedEvent`, `V1CouponDeletedEvent`, and `V1CouponUpdatedEvent` with related object `Coupon`
  * Remove support for thin events `V1CreditNoteCreatedEvent`, `V1CreditNoteUpdatedEvent`, and `V1CreditNoteVoidedEvent` with related object `CreditNote`
  * Remove support for thin events `V1CustomerCreatedEvent`, `V1CustomerDeletedEvent`, and `V1CustomerUpdatedEvent` with related object `Customer`
  * Remove support for thin events `V1CustomerSubscriptionCreatedEvent`, `V1CustomerSubscriptionDeletedEvent`, `V1CustomerSubscriptionPausedEvent`, `V1CustomerSubscriptionPendingUpdateAppliedEvent`, `V1CustomerSubscriptionPendingUpdateExpiredEvent`, `V1CustomerSubscriptionResumedEvent`, `V1CustomerSubscriptionTrialWillEndEvent`, and `V1CustomerSubscriptionUpdatedEvent` with related object `Subscription`
  * Remove support for thin events `V1CustomerTaxIdCreatedEvent`, `V1CustomerTaxIdDeletedEvent`, and `V1CustomerTaxIdUpdatedEvent` with related object `TaxID`
  * Remove support for thin event `V1FileCreatedEvent` with related object `File`
  * Remove support for thin events `V1FinancialConnectionsAccountCreatedEvent`, `V1FinancialConnectionsAccountDeactivatedEvent`, `V1FinancialConnectionsAccountDisconnectedEvent`, `V1FinancialConnectionsAccountReactivatedEvent`, `V1FinancialConnectionsAccountRefreshedBalanceEvent`, `V1FinancialConnectionsAccountRefreshedOwnershipEvent`, and `V1FinancialConnectionsAccountRefreshedTransactionsEvent` with related object `FinancialConnectionsAccount`
  * Remove support for thin events `V1IdentityVerificationSessionCanceledEvent`, `V1IdentityVerificationSessionCreatedEvent`, `V1IdentityVerificationSessionProcessingEvent`, `V1IdentityVerificationSessionRedactedEvent`, `V1IdentityVerificationSessionRequiresInputEvent`, and `V1IdentityVerificationSessionVerifiedEvent` with related object `IdentityVerificationSession`
  * Remove support for thin events `V1InvoiceCreatedEvent`, `V1InvoiceDeletedEvent`, `V1InvoiceFinalizationFailedEvent`, `V1InvoiceFinalizedEvent`, `V1InvoiceMarkedUncollectibleEvent`, `V1InvoiceOverdueEvent`, `V1InvoiceOverpaidEvent`, `V1InvoicePaidEvent`, `V1InvoicePaymentActionRequiredEvent`, `V1InvoicePaymentFailedEvent`, `V1InvoicePaymentSucceededEvent`, `V1InvoiceSentEvent`, `V1InvoiceUpcomingEvent`, `V1InvoiceUpdatedEvent`, `V1InvoiceVoidedEvent`, and `V1InvoiceWillBeDueEvent` with related object `Invoice`
  * Remove support for thin event `V1InvoicePaymentPaidEvent` with related object `InvoicePayment`
  * Remove support for thin events `V1InvoiceitemCreatedEvent` and `V1InvoiceitemDeletedEvent` with related object `InvoiceItem`
  * Remove support for thin events `V1IssuingAuthorizationCreatedEvent`, `V1IssuingAuthorizationRequestEvent`, and `V1IssuingAuthorizationUpdatedEvent` with related object `IssuingAuthorization`
  * Remove support for thin events `V1IssuingCardCreatedEvent` and `V1IssuingCardUpdatedEvent` with related object `IssuingCard`
  * Remove support for thin events `V1IssuingCardholderCreatedEvent` and `V1IssuingCardholderUpdatedEvent` with related object `IssuingCardholder`
  * Remove support for thin events `V1IssuingDisputeClosedEvent`, `V1IssuingDisputeCreatedEvent`, `V1IssuingDisputeFundsReinstatedEvent`, `V1IssuingDisputeFundsRescindedEvent`, `V1IssuingDisputeSubmittedEvent`, and `V1IssuingDisputeUpdatedEvent` with related object `IssuingDispute`
  * Remove support for thin events `V1IssuingPersonalizationDesignActivatedEvent`, `V1IssuingPersonalizationDesignDeactivatedEvent`, `V1IssuingPersonalizationDesignRejectedEvent`, and `V1IssuingPersonalizationDesignUpdatedEvent` with related object `IssuingPersonalizationDesign`
  * Remove support for thin events `V1IssuingTokenCreatedEvent` and `V1IssuingTokenUpdatedEvent` with related object `IssuingToken`
  * Remove support for thin events `V1IssuingTransactionCreatedEvent`, `V1IssuingTransactionPurchaseDetailsReceiptUpdatedEvent`, and `V1IssuingTransactionUpdatedEvent` with related object `IssuingTransaction`
  * Remove support for thin event `V1MandateUpdatedEvent` with related object `Mandate`
  * Remove support for thin events `V1PaymentIntentAmountCapturableUpdatedEvent`, `V1PaymentIntentCanceledEvent`, `V1PaymentIntentCreatedEvent`, `V1PaymentIntentPartiallyFundedEvent`, `V1PaymentIntentPaymentFailedEvent`, `V1PaymentIntentProcessingEvent`, `V1PaymentIntentRequiresActionEvent`, and `V1PaymentIntentSucceededEvent` with related object `PaymentIntent`
  * Remove support for thin events `V1PaymentLinkCreatedEvent` and `V1PaymentLinkUpdatedEvent` with related object `PaymentLink`
  * Remove support for thin events `V1PaymentMethodAttachedEvent`, `V1PaymentMethodAutomaticallyUpdatedEvent`, `V1PaymentMethodDetachedEvent`, and `V1PaymentMethodUpdatedEvent` with related object `PaymentMethod`
  * Remove support for thin events `V1PayoutCanceledEvent`, `V1PayoutCreatedEvent`, `V1PayoutFailedEvent`, `V1PayoutPaidEvent`, `V1PayoutReconciliationCompletedEvent`, and `V1PayoutUpdatedEvent` with related object `Payout`
  * Remove support for thin events `V1PersonCreatedEvent`, `V1PersonDeletedEvent`, and `V1PersonUpdatedEvent` with related object `Person`
  * Remove support for thin events `V1PlanCreatedEvent`, `V1PlanDeletedEvent`, and `V1PlanUpdatedEvent` with related object `Plan`
  * Remove support for thin events `V1PriceCreatedEvent`, `V1PriceDeletedEvent`, and `V1PriceUpdatedEvent` with related object `Price`
  * Remove support for thin events `V1ProductCreatedEvent`, `V1ProductDeletedEvent`, and `V1ProductUpdatedEvent` with related object `Product`
  * Remove support for thin events `V1PromotionCodeCreatedEvent` and `V1PromotionCodeUpdatedEvent` with related object `PromotionCode`
  * Remove support for thin events `V1QuoteAcceptedEvent`, `V1QuoteCanceledEvent`, `V1QuoteCreatedEvent`, and `V1QuoteFinalizedEvent` with related object `Quote`
  * Remove support for thin events `V1RadarEarlyFraudWarningCreatedEvent` and `V1RadarEarlyFraudWarningUpdatedEvent` with related object `RadarEarlyFraudWarning`
  * Remove support for thin events `V1ReviewClosedEvent` and `V1ReviewOpenedEvent` with related object `Review`
  * Remove support for thin events `V1SetupIntentCanceledEvent`, `V1SetupIntentCreatedEvent`, `V1SetupIntentRequiresActionEvent`, `V1SetupIntentSetupFailedEvent`, and `V1SetupIntentSucceededEvent` with related object `SetupIntent`
  * Remove support for thin event `V1SigmaScheduledQueryRunCreatedEvent` with related object `SigmaScheduledQueryRun`
  * Remove support for thin events `V1SourceCanceledEvent`, `V1SourceChargeableEvent`, `V1SourceFailedEvent`, and `V1SourceRefundAttributesRequiredEvent` with related object `Source`
  * Remove support for thin events `V1SubscriptionScheduleAbortedEvent`, `V1SubscriptionScheduleCanceledEvent`, `V1SubscriptionScheduleCompletedEvent`, `V1SubscriptionScheduleCreatedEvent`, `V1SubscriptionScheduleExpiringEvent`, `V1SubscriptionScheduleReleasedEvent`, and `V1SubscriptionScheduleUpdatedEvent` with related object `SubscriptionSchedule`
  * Remove support for thin events `V1TaxRateCreatedEvent` and `V1TaxRateUpdatedEvent` with related object `TaxRate`
  * Remove support for thin events `V1TerminalReaderActionFailedEvent`, `V1TerminalReaderActionSucceededEvent`, and `V1TerminalReaderActionUpdatedEvent` with related object `TerminalReader`
  * Remove support for thin events `V1TestHelpersTestClockAdvancingEvent`, `V1TestHelpersTestClockCreatedEvent`, `V1TestHelpersTestClockDeletedEvent`, `V1TestHelpersTestClockInternalFailureEvent`, and `V1TestHelpersTestClockReadyEvent` with related object `TestHelpersTestClock`
  * Remove support for thin events `V1TopupCanceledEvent`, `V1TopupCreatedEvent`, `V1TopupFailedEvent`, `V1TopupReversedEvent`, and `V1TopupSucceededEvent` with related object `Topup`
  * Remove support for thin events `V1TransferCreatedEvent`, `V1TransferReversedEvent`, and `V1TransferUpdatedEvent` with related object `Transfer`

## 84.1.0-alpha.1 - 2025-11-18
This release changes the pinned API version to `2025-11-17.preview`.

* [#2214](https://github.com/stripe/stripe-go/pull/2214) Update generated code for private-preview
  * Add support for new resources `BalanceTransfer` and `RadarAccountEvaluation`
  * Add support for `New` method on resource `BalanceTransfer`
  * Add support for `Get`, `New`, and `Update` methods on resource `RadarAccountEvaluation`
  * Add support for `SpecifiedCommercialTransactionsActURL` on `AccountBusinessProfileParams` and `AccountBusinessProfile`
  * Add support for `PaypayPayments` on `AccountSettingsParams` and `AccountSettings`
  * Change type of `BillingAnalyticsMeterUsageMeterParams.DimensionFilters` from `string` to `array(string)`
  * Change type of `BillingAnalyticsMeterUsageMeterParams.TenantFilters` from `string` to `array(string)`
  * Add support for `PaymentMethodConfiguration` on `BillingPortalConfigurationFeaturesPaymentMethodUpdate`
  * Add support for `CarRentalData`, `FlightData`, and `LodgingData` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmPaymentDetailsParams`, and `PaymentIntentPaymentDetailsParams`
  * Add support for `TransactionID` on `ChargePaymentMethodDetailsIdeal`, `PaymentAttemptRecordPaymentMethodDetailsIdeal`, and `PaymentRecordPaymentMethodDetailsIdeal`
  * Add support for new value `finom` on enums `ConfirmationTokenPaymentMethodPreviewIdeal.Bank`, `PaymentAttemptRecordPaymentMethodDetailsIdeal.Bank`, and `PaymentRecordPaymentMethodDetailsIdeal.Bank`
  * Add support for new value `FNOMNL22` on enums `ConfirmationTokenPaymentMethodPreviewIdeal.BIC`, `PaymentAttemptRecordPaymentMethodDetailsIdeal.BIC`, and `PaymentRecordPaymentMethodDetailsIdeal.BIC`
  * Add support for new value `tokenized_account_number_deactivated` on enums `ConfirmationTokenPaymentMethodPreviewUsBankAccountStatusDetailsBlocked.Reason` and `PaymentMethodUsBankAccountStatusDetailsBlocked.Reason`
  * Add support for `Created` on `CustomerCustomerBalanceTransactionListParams` and `InvoicePaymentListParams`
  * Add support for new values `capital.financing_offer.accepted_other_offer`, `financial_connections.account.account_numbers_updated`, and `financial_connections.account.upcoming_account_number_expiry` on enum `Event.Type`
  * Add support for `AccountNumbers` on `FinancialConnectionsAccount`
  * Add support for `FraudRisk` on `IssuingAuthorizationRiskAssessmentParams`
  * Add support for `LatestFraudWarning` on `IssuingCard`
  * Add support for `SupplementaryPurchaseData` on `OrderPaymentSettingsPaymentMethodOptionsKlarnaParams`, `PaymentIntentConfirmPaymentMethodOptionsKlarnaParams`, and `PaymentIntentPaymentMethodOptionsKlarnaParams`
  * Add support for `CaptureMethod` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, and `PaymentIntentPaymentMethodOptionsCardPresent`
  * Add support for `AllowRedisplay` and `CustomerAccount` on `PaymentMethodListParams`
  * Add support for `MbWay` and `TWINT` on `RefundDestinationDetails`
  * Change type of `SubscriptionScheduleParams.BillingSchedules` from `array(billing_schedules_update_params)` to `emptyable(array(billing_schedules_update_params))`
  * Add support for snapshot events `EventTypeFinancialConnectionsAccountAccountNumbersUpdated` and `EventTypeFinancialConnectionsAccountUpcomingAccountNumberExpiry` with resource `FinancialConnectionsAccount`
* [#2217](https://github.com/stripe/stripe-go/pull/2217) Update generated code for private-preview
  * Add support for `BillingSchedulesActions` on `InvoiceCreatePreviewScheduleDetailsAmendmentParams` and `SubscriptionScheduleAmendAmendmentParams`

## 83.3.0-alpha.2 - 2025-11-13
This release changes the pinned API version to `2025-10-29.preview`.

* [#2204](https://github.com/stripe/stripe-go/pull/2204) Update generated code for private-preview
  * Remove support for resource `V2TaxAutomaticRule`
  * Remove support for `Deactivate`, `Find`, `Get`, `New`, and `Update` methods on resource `V2TaxAutomaticRule`
  * Add support for `SelfReportedIncome` and `SelfReportedMonthlyHousingPayment` on `AccountIndividualParams`, `AccountPersonParams`, `Person`, `TokenAccountIndividualParams`, and `TokenPersonParams`
  * Add support for `BillingSchedules` and `PhaseEffectiveAt` on `QuoteSubscriptionDataOverrideParams`, `QuoteSubscriptionDataOverridesParams`, `QuoteSubscriptionDataOverrides`, `QuoteSubscriptionDataParams`, and `QuoteSubscriptionData`
  * Add support for `BillFrom` on `SubscriptionBillingSchedule`
  * Add support for `AmendmentEnd` and `LineEndsAt` on `SubscriptionBillingScheduleBillUntil`
  * Add support for new values `amendment_end`, `line_ends_at`, `schedule_end`, and `upcoming_invoice` on enum `SubscriptionBillingScheduleBillUntil.Type`
* [#2213](https://github.com/stripe/stripe-go/pull/2213) Update generated code for private-preview
  * Add support for new resource `IssuingProgram`
  * Add support for `Get`, `List`, `New`, and `Update` methods on resource `IssuingProgram`
  * Add support for `Schedule` on `Discount`
  * Add support for `ApplicableFees` on `DelegatedCheckoutRequestedSessionTotalDetails`
  * Add support for `ScheduleDetails` on `InvoiceItemParent`, `InvoiceLineItemParent`, `InvoiceParent`, and `QuotePreviewInvoiceParent`
  * Add support for new value `schedule_details` on enum `InvoiceItemParent.Type`
  * Add support for `BillingSchedules` on `InvoiceCreatePreviewScheduleDetailsParams`, `QuotePreviewSubscriptionSchedule`, `SubscriptionScheduleParams`, and `SubscriptionSchedule`
  * Add support for new value `schedule_details` on enums `InvoiceParent.Type` and `QuotePreviewInvoiceParent.Type`
  * Add support for new value `schedule_details` on enum `InvoiceLineItemParent.Type`
  * Add support for `LatestInvoice` on `QuotePreviewSubscriptionSchedule` and `SubscriptionSchedule`
  * Add support for `PhaseEffectiveAt` on `QuotePreviewSubscriptionScheduleDefaultSettings`, `SubscriptionScheduleDefaultSettingsParams`, and `SubscriptionScheduleDefaultSettings`

## 83.3.0-alpha.1 - 2025-11-06
* [#2200](https://github.com/stripe/stripe-go/pull/2200) Update generated code for private-preview
  * Add support for new resources `TransitBalance`, `V2ReportingReportRun`, `V2ReportingReport`
  * Add support for `Get` and `New` methods on resource `V2ReportingReportRun`
  * Add support for `Get` method on resource `V2ReportingReport`
  * Add support for `New` and `Refill` test helper methods on resource `CapitalFinancingOffer`
  * Add support for `AllocatedFunds` on `Charge`, `PaymentIntentConfirmParams`, and `PaymentIntentParams`
  * Add support for thin events `V2ReportingReportRunCreatedEvent`, `V2ReportingReportRunFailedEvent`, `V2ReportingReportRunSucceededEvent`, and `V2ReportingReportRunUpdatedEvent` with related object `V2ReportingReportRun`

## 83.2.0-alpha.2 - 2025-10-30
* [#2198](https://github.com/stripe/stripe-go/pull/2198) Update generated code for private-preview
  * Add support for `PaymentMethodPreview` on `DelegatedCheckoutRequestedSession`
  * Add support for `OrderID` on `DelegatedCheckoutRequestedSessionOrderDetails`
  * Add support for `Lead` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercial`, `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialParams`, and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercial`
  * Add support for `GlobalAccountHolder` on `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialParams` and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercial`
  * Add support for new value `commercial.lead.prepaid_card` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for new value `commercial.lead.prepaid_card` on enum `EventsV2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEvent.UpdatedCapability`

## 83.2.0-alpha.1 - 2025-10-29
* [#2192](https://github.com/stripe/stripe-go/pull/2192) Update generated code for private-preview
  * Add support for `ReportRefund` method on resource `PaymentRecord`
  * Add support for new value `verification_data_not_found` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
  * Add support for `Tenants` on `BillingAnalyticsMeterUsageRow`
  * Add support for `RepresentativeDeclaration` on `AccountCompanyParams`, `AccountCompany`, and `TokenAccountCompanyParams`
  * Add support for `Transfer` on `ApplicationFeeFeeSource`
  * Add support for new value `transfer` on enum `ApplicationFeeFeeSource.Type`
  * Add support for `TransitBalancesTotal` on `Balance`
  * Add support for new value `transit` on enum `BalanceTransaction.BalanceType`
  * Add support for `TenantGroupByKeys` on `BillingAnalyticsMeterUsageMeterParams`
  * Add support for `PaymentMethodConfiguration` on `BillingPortalConfigurationFeaturesPaymentMethodUpdateParams`
  * Add support for new value `solana` on enums `ChargePaymentMethodDetailsCrypto.Network`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.Network`, and `PaymentRecordPaymentMethodDetailsCrypto.Network`
  * Add support for `PaymentPortalURL` on `ChargePaymentMethodDetailsRechnung`, `PaymentAttemptRecordPaymentMethodDetailsRechnung`, and `PaymentRecordPaymentMethodDetailsRechnung`
  * Add support for `TWINT` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
  * Add support for new value `custom` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
  * Add support for `CustomerSheet`, `MobilePaymentElement`, and `TaxIDElement` on `CustomerSessionComponentsParams` and `CustomerSessionComponents`
  * Add support for `Provider` on `CustomerTax`
  * Remove support for `RiskDetails` on `DelegatedCheckoutRequestedSessionParams`
  * Add support for `RiskDetails` on `DelegatedCheckoutRequestedSessionConfirmParams`
  * Add support for new value `platform_terms_of_service` on enum `File.Purpose`
  * Add support for `StartingAfter` on `PaymentAttemptRecordListParams`
  * Add support for `Reference` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna`, `PaymentIntentAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, `PaymentIntentCaptureAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, `PaymentIntentConfirmAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`
  * Add support for `AllocatedFunds` on `PaymentIntent`
  * Add support for `SubscriptionReference` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna`
  * Add support for `NameCollection` on `PaymentLinkParams` and `PaymentLink`
  * Add support for `Crypto` on `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, and `RefundDestinationDetails`
  * Add support for `MbWay` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
  * Add support for `Custom` on `PaymentMethodParams` and `PaymentMethod`
  * Add support for `ExcludedPaymentMethodTypes` on `SetupIntentParams` and `SetupIntent`
  * Add support for `Tw` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
  * Add support for `Gip` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
  * Add support for `LastSeenAt` on `TerminalReader`
  * Add support for `ApplicationFeeAmount` on `TransferParams` and `Transfer`
  * Add support for `ApplicationFee` on `Transfer`
  * Add support for `HighRiskActivitiesDescription`, `HighRiskActivities`, `MoneyServicesDescription`, `OperatesInProhibitedCountries`, `ParticipatesInRegulatedActivity`, `PurposeOfFundsDescription`, `PurposeOfFunds`, `RegulatedActivity`, `SourceOfFundsDescription`, and `SourceOfFunds` on `V2CoreAccountConfigurationStorerParams` and `V2CoreAccountConfigurationStorer`
  * Add support for `CryptoWallets` on `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddressesParams`, `V2CoreAccountConfigurationStorerCapabilitiesFinancialAddresses`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsParams`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPayments`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfersParams`, and `V2CoreAccountConfigurationStorerCapabilitiesOutboundTransfers`
  * Add support for `Usdc` on `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesParams` and `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrencies`
  * Add support for `CryptoStorer` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams` and `V2CoreAccountIdentityAttestationsTermsOfService`
  * Add support for `ComplianceScreeningDescription` on `V2CoreAccountIdentityBusinessDetailsParams` and `V2CoreAccountIdentityBusinessDetails`
  * Add support for `ExternalAmount` on `V2MoneyManagementReceivedCredit` and `V2MoneyManagementReceivedDebit`
  * Add support for error code `payment_intent_rate_limit_exceeded` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `QuotePreviewInvoiceLastFinalizationError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`

## 83.1.0-alpha.6 - 2025-10-23
* [#2184](https://github.com/stripe/stripe-go/pull/2184) Update generated code for private-preview
  * Add support for new resource `V2BillingPricingPlanSubscriptionComponents`
  * Add support for `Get` method on resource `V2BillingPricingPlanSubscriptionComponents`
  * Add support for `DimensionPayloadKeys` on `BillingMeterParams` and `BillingMeter`
  * Add support for `DimensionFilters` and `DimensionGroupByKeys` on `BillingBillingMeterMeterEventSummaryListParams`
  * Add support for `Dimensions` on `BillingMeterEventSummary`
  * Add support for `FulfillmentDetails` and `PaymentMethodData` on `DelegatedCheckoutRequestedSessionParams`
  * Add support for `LineItemDetails`, `Metadata`, `PaymentMethod`, and `SharedMetadata` on `DelegatedCheckoutRequestedSessionParams` and `DelegatedCheckoutRequestedSession`
  * Add support for `Currency`, `Customer`, and `RiskDetails` on `DelegatedCheckoutRequestedSessionParams`
  * Add support for `SellerDetails` and `SetupFutureUsage` on `DelegatedCheckoutRequestedSessionParams` and `DelegatedCheckoutRequestedSession`
  * Add support for `AmountSubtotal`, `AmountTotal`, `CreatedAt`, `ExpiresAt`, `OrderDetails`, `SharedPaymentIssuedToken`, `Status`, `TotalDetails`, and `UpdatedAt` on `DelegatedCheckoutRequestedSession`
  * Add support for `Address`, `Email`, `FulfillmentOptions`, `Name`, `Phone`, and `SelectedFulfillmentOption` on `DelegatedCheckoutRequestedSessionFulfillmentDetails`
  * Add support for new values `billie`, `crypto`, `kr_card`, `kriya`, `mb_way`, `mondu`, `ng_bank_transfer`, `ng_bank`, `ng_card`, `ng_market`, `ng_ussd`, `ng_wallet`, `payco`, `paypay`, `rechnung`, `samsung_pay`, `satispay`, `scalapay`, `sequra`, `sunbit`, `us_bank_account`, and `vipps` on enums `EventsV2CoreHealthAuthorizationRateDropFiringEventImpact.PaymentMethodType`, `EventsV2CoreHealthAuthorizationRateDropResolvedEventImpact.PaymentMethodType`, `EventsV2CoreHealthPaymentMethodErrorFiringEventImpact.PaymentMethodType`, and `EventsV2CoreHealthPaymentMethodErrorResolvedEventImpact.PaymentMethodType`

## 83.1.0-alpha.5 - 2025-10-21
* [#2183](https://github.com/stripe/stripe-go/pull/2183) Fix URL serialization for array query parameters that affected V2 GET APIs

## 83.1.0-alpha.4 - 2025-10-17
* [#2173](https://github.com/stripe/stripe-go/pull/2173) Update generated code for private-preview
  * Add support for new resources `DelegatedCheckoutRequestedSession` and `IdentityBlocklistEntry`
  * Add support for `Confirm`, `Expire`, `Get`, `New`, and `Update` methods on resource `DelegatedCheckoutRequestedSession`
  * Add support for `Disable`, `Get`, `List`, and `New` methods on resource `IdentityBlocklistEntry`
  * Add support for `BlockedByEntry` on `IdentityVerificationReportDocument`, `IdentityVerificationReportListParams`, and `IdentityVerificationReportSelfie`

## 83.1.0-alpha.3 - 2025-10-09
* [#2155](https://github.com/stripe/stripe-go/pull/2155) Update generated code for private-preview
  * Add support for new resource `PaymentMethodBalance`
  * Add support for `CheckBalance` method on resource `PaymentMethod`
  * Add support for `Benefits` on `Card`, `ChargePaymentMethodDetailsCard`, `ConfirmationTokenPaymentMethodPreviewCard`, and `PaymentMethodCard`
  * Add support for `Benefit` on `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentPaymentDetailsParams`, and `PaymentIntentPaymentDetails`
  * Add support for `SetupDetails` on `SetupIntentConfirmParams`, `SetupIntentParams`, and `SetupIntent`
  * Add support for new value `card_creator` on enum `V2CoreAccount.AppliedConfigurations`
  * Add support for `CardCreator` on `V2CoreAccountConfigurationParams`, `V2CoreAccountConfiguration`, `V2CoreAccountIdentityAttestationsTermsOfServiceParams`, and `V2CoreAccountIdentityAttestationsTermsOfService`
  * Add support for new values `commercial.celtic.charge_card`, `commercial.celtic.spend_card`, `commercial.cross_river_bank.charge_card`, `commercial.cross_river_bank.spend_card`, `commercial.stripe.charge_card`, and `commercial.stripe.prepaid_card` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
  * Add support for new value `card_creator` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Configuration`
  * Add support for thin events `V2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEvent` and `V2CoreAccountIncludingConfigurationCardCreatorUpdatedEvent` with related object `V2CoreAccount`
  * Remove support for thin events `V1CustomerDiscountCreatedEvent`, `V1CustomerDiscountDeletedEvent`, and `V1CustomerDiscountUpdatedEvent` with related object `Discount`

## 83.1.0-alpha.2 - 2025-10-08
* Contains bug fixes and improvements from [v83.0.1](https://github.com/stripe/stripe-go/blob/v83.0.1/CHANGELOG.md#8301---2025-10-08).

## 83.1.0-alpha.1 - 2025-09-30
This release changes the pinned API version to `2025-09-30.preview`.

It is built on top of SDK version 83.0.0 and 83.1.0-beta.1 which contains breaking changes. Please review the changelog for these versions in this page if upgrading from older SDK versions.

* [#2131](https://github.com/stripe/stripe-go/pull/2131) Update generated code for private-preview
  * Add support for new resource `V2MoneyManagementRecipientVerification`
  * Add support for `Acknowledge`, `Get`, `New`, and `RecipientVerifications` methods on resource `V2MoneyManagementRecipientVerification`
  * Add support for `Update` method on resources `V2BillingPricingPlanSubscription` and `V2BillingServiceAction`
  * Add support for `CryptoWallets` on `V2AccountConfigurationRecipientDataFeaturesParams`, `V2AccountConfigurationRecipientDataFeatures`, `V2CoreAccountConfigurationRecipientCapabilitiesParams`, and `V2CoreAccountConfigurationRecipientCapabilities`
  * Add support for new value `crypto` on enum `V2CoreAccountRequirementsEntriesImpactRestrictsCapabilities.Capability`
  * Add support for new value `crypto_wallet` on enum `V2AccountConfigurationRecipientDataDefaultOutboundDestination.Type`
  * Add support for new value `crypto_wallets` on enum `V2AccountConfigurationSupportableFeatures.RecipientData`
  * Add support for new value `crypto_wallets` on enum `V2AccountRequirementsImpact.RequiredForFeatures`
  * Add support for `LookupKey` on `V2BillingCadenceParams` and `V2BillingCadence`
  * Add support for `SettingsData` on `V2BillingCadence`
  * Add support for `V1EventID` on `V2CoreEvent`
  * Add support for `RecipientVerification` on `V2MoneyManagementOutboundPaymentParams`, `V2MoneyManagementOutboundPayment`, `V2MoneyManagementOutboundTransferParams`, and `V2MoneyManagementOutboundTransfer`
  * Add support for `CryptoWallet` on `V2MoneyManagementOutboundSetupIntentPayoutMethodDataParams` and `V2MoneyManagementPayoutMethod`
  * Add support for `CustomPricingUnitDetails` on `V2BillingRateCardRateCustomPricingUnitAmount`, `V2BillingServiceActionCreditGrantAmountCustomPricingUnit`, and `V2BillingServiceActionCreditGrantPerTenantAmountCustomPricingUnit`
  * Add support for `OriginType` on `V2MoneyManagementReceivedDebitBankTransfer`
  * Add support for `CreditGrants` on `BillingAlertCreditBalanceThresholdFilterParams`
  * Add support for `PaymentRecordRefund` and `Type` on `CreditNotePreviewLinesRefundParams`, `CreditNotePreviewRefundParams`, `CreditNoteRefundParams`, and `CreditNoteRefund`
  * Add support for `BillingCadence` on `InvoiceListParams`
  * Add support for `SEPABankAccount` on `V2MoneyManagementFinancialAddressParams`
  * Remove support for `Price` on `V2BillingRateCardRateParams`
  * Add support for `LookupKeys` on `V2BillingCadenceListParams`
  * Change type of `V2BillingCadenceCancelParams.Include`, `V2BillingCadenceListParams.Include`, and `V2BillingCadenceParams.Include` from `literal('invoice_discount_rules')` to `enum('invoice_discount_rules'|'settings_data')`
  * Remove support for `Customer` and `Type` on `V2BillingCadencePayerParams`
  * Add support for new value `crypto_wallets` on enum `EventsAccountConfigurationRecipientDataFeatureStatusUpdatedEvent.FeatureName`
  * Add support for new value `crypto_wallets_v2` on enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`
  * Remove support for `AlertID` on `EventsV2CoreHealthApiErrorResolvedEvent`, `EventsV2CoreHealthApiLatencyResolvedEvent`, `EventsV2CoreHealthAuthorizationRateDropResolvedEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEvent`, `EventsV2CoreHealthPaymentMethodErrorResolvedEvent`, `EventsV2CoreHealthTrafficVolumeDropResolvedEvent`, and `EventsV2CoreHealthWebhookLatencyResolvedEvent`
  * Add support for thin event `V1AccountUpdatedEvent` with related object `V2Account`
  * Add support for thin events `V1ApplicationFeeCreatedEvent`, `V1ApplicationFeeRefundedEvent`, `V1BillingPortalConfigurationCreatedEvent`, `V1BillingPortalConfigurationUpdatedEvent`, `V1CapabilityUpdatedEvent`, `V1ChargeCapturedEvent`, `V1ChargeDisputeClosedEvent`, `V1ChargeDisputeCreatedEvent`, `V1ChargeDisputeFundsReinstatedEvent`, `V1ChargeDisputeFundsWithdrawnEvent`, `V1ChargeDisputeUpdatedEvent`, `V1ChargeExpiredEvent`, `V1ChargeFailedEvent`, `V1ChargePendingEvent`, `V1ChargeRefundUpdatedEvent`, `V1ChargeRefundedEvent`, `V1ChargeSucceededEvent`, `V1ChargeUpdatedEvent`, `V1CheckoutSessionAsyncPaymentFailedEvent`, `V1CheckoutSessionAsyncPaymentSucceededEvent`, `V1CheckoutSessionCompletedEvent`, `V1CheckoutSessionExpiredEvent`, `V1ClimateOrderCanceledEvent`, `V1ClimateOrderCreatedEvent`, `V1ClimateOrderDelayedEvent`, `V1ClimateOrderDeliveredEvent`, `V1ClimateOrderProductSubstitutedEvent`, `V1ClimateProductCreatedEvent`, `V1ClimateProductPricingUpdatedEvent`, `V1CouponCreatedEvent`, `V1CouponDeletedEvent`, `V1CouponUpdatedEvent`, `V1CreditNoteCreatedEvent`, `V1CreditNoteUpdatedEvent`, `V1CreditNoteVoidedEvent`, `V1CustomerCreatedEvent`, `V1CustomerDeletedEvent`, `V1CustomerSubscriptionCreatedEvent`, `V1CustomerSubscriptionDeletedEvent`, `V1CustomerSubscriptionPausedEvent`, `V1CustomerSubscriptionPendingUpdateAppliedEvent`, `V1CustomerSubscriptionPendingUpdateExpiredEvent`, `V1CustomerSubscriptionResumedEvent`, `V1CustomerSubscriptionTrialWillEndEvent`, `V1CustomerSubscriptionUpdatedEvent`, `V1CustomerTaxIdCreatedEvent`, `V1CustomerTaxIdDeletedEvent`, `V1CustomerTaxIdUpdatedEvent`, `V1CustomerUpdatedEvent`, `V1FileCreatedEvent`, `V1FinancialConnectionsAccountCreatedEvent`, `V1FinancialConnectionsAccountDeactivatedEvent`, `V1FinancialConnectionsAccountDisconnectedEvent`, `V1FinancialConnectionsAccountReactivatedEvent`, `V1FinancialConnectionsAccountRefreshedBalanceEvent`, `V1FinancialConnectionsAccountRefreshedOwnershipEvent`, `V1FinancialConnectionsAccountRefreshedTransactionsEvent`, `V1IdentityVerificationSessionCanceledEvent`, `V1IdentityVerificationSessionCreatedEvent`, `V1IdentityVerificationSessionProcessingEvent`, `V1IdentityVerificationSessionRedactedEvent`, `V1IdentityVerificationSessionRequiresInputEvent`, `V1IdentityVerificationSessionVerifiedEvent`, `V1InvoiceCreatedEvent`, `V1InvoiceDeletedEvent`, `V1InvoiceFinalizationFailedEvent`, `V1InvoiceFinalizedEvent`, `V1InvoiceMarkedUncollectibleEvent`, `V1InvoiceOverdueEvent`, `V1InvoiceOverpaidEvent`, `V1InvoicePaidEvent`, `V1InvoicePaymentActionRequiredEvent`, `V1InvoicePaymentFailedEvent`, `V1InvoicePaymentPaidEvent`, `V1InvoicePaymentSucceededEvent`, `V1InvoiceSentEvent`, `V1InvoiceUpcomingEvent`, `V1InvoiceUpdatedEvent`, `V1InvoiceVoidedEvent`, `V1InvoiceWillBeDueEvent`, `V1InvoiceitemCreatedEvent`, `V1InvoiceitemDeletedEvent`, `V1IssuingAuthorizationCreatedEvent`, `V1IssuingAuthorizationRequestEvent`, `V1IssuingAuthorizationUpdatedEvent`, `V1IssuingCardCreatedEvent`, `V1IssuingCardUpdatedEvent`, `V1IssuingCardholderCreatedEvent`, `V1IssuingCardholderUpdatedEvent`, `V1IssuingDisputeClosedEvent`, `V1IssuingDisputeCreatedEvent`, `V1IssuingDisputeFundsReinstatedEvent`, `V1IssuingDisputeFundsRescindedEvent`, `V1IssuingDisputeSubmittedEvent`, `V1IssuingDisputeUpdatedEvent`, `V1IssuingPersonalizationDesignActivatedEvent`, `V1IssuingPersonalizationDesignDeactivatedEvent`, `V1IssuingPersonalizationDesignRejectedEvent`, `V1IssuingPersonalizationDesignUpdatedEvent`, `V1IssuingTokenCreatedEvent`, `V1IssuingTokenUpdatedEvent`, `V1IssuingTransactionCreatedEvent`, `V1IssuingTransactionPurchaseDetailsReceiptUpdatedEvent`, `V1IssuingTransactionUpdatedEvent`, `V1MandateUpdatedEvent`, `V1PaymentIntentAmountCapturableUpdatedEvent`, `V1PaymentIntentCanceledEvent`, `V1PaymentIntentCreatedEvent`, `V1PaymentIntentPartiallyFundedEvent`, `V1PaymentIntentPaymentFailedEvent`, `V1PaymentIntentProcessingEvent`, `V1PaymentIntentRequiresActionEvent`, `V1PaymentIntentSucceededEvent`, `V1PaymentLinkCreatedEvent`, `V1PaymentLinkUpdatedEvent`, `V1PaymentMethodAttachedEvent`, `V1PaymentMethodAutomaticallyUpdatedEvent`, `V1PaymentMethodDetachedEvent`, `V1PaymentMethodUpdatedEvent`, `V1PayoutCanceledEvent`, `V1PayoutCreatedEvent`, `V1PayoutFailedEvent`, `V1PayoutPaidEvent`, `V1PayoutReconciliationCompletedEvent`, `V1PayoutUpdatedEvent`, `V1PersonCreatedEvent`, `V1PersonDeletedEvent`, `V1PersonUpdatedEvent`, `V1PlanCreatedEvent`, `V1PlanDeletedEvent`, `V1PlanUpdatedEvent`, `V1PriceCreatedEvent`, `V1PriceDeletedEvent`, `V1PriceUpdatedEvent`, `V1ProductCreatedEvent`, `V1ProductDeletedEvent`, `V1ProductUpdatedEvent`, `V1PromotionCodeCreatedEvent`, `V1PromotionCodeUpdatedEvent`, `V1QuoteAcceptedEvent`, `V1QuoteCanceledEvent`, `V1QuoteCreatedEvent`, `V1QuoteFinalizedEvent`, `V1RadarEarlyFraudWarningCreatedEvent`, `V1RadarEarlyFraudWarningUpdatedEvent`, `V1RefundCreatedEvent`, `V1RefundFailedEvent`, `V1RefundUpdatedEvent`, `V1ReviewClosedEvent`, `V1ReviewOpenedEvent`, `V1SetupIntentCanceledEvent`, `V1SetupIntentCreatedEvent`, `V1SetupIntentRequiresActionEvent`, `V1SetupIntentSetupFailedEvent`, `V1SetupIntentSucceededEvent`, `V1SigmaScheduledQueryRunCreatedEvent`, `V1SourceCanceledEvent`, `V1SourceChargeableEvent`, `V1SourceFailedEvent`, `V1SourceRefundAttributesRequiredEvent`, `V1SubscriptionScheduleAbortedEvent`, `V1SubscriptionScheduleCanceledEvent`, `V1SubscriptionScheduleCompletedEvent`, `V1SubscriptionScheduleCreatedEvent`, `V1SubscriptionScheduleExpiringEvent`, `V1SubscriptionScheduleReleasedEvent`, `V1SubscriptionScheduleUpdatedEvent`, `V1TaxRateCreatedEvent`, `V1TaxRateUpdatedEvent`, `V1TerminalReaderActionFailedEvent`, `V1TerminalReaderActionSucceededEvent`, `V1TerminalReaderActionUpdatedEvent`, `V1TestHelpersTestClockAdvancingEvent`, `V1TestHelpersTestClockCreatedEvent`, `V1TestHelpersTestClockDeletedEvent`, `V1TestHelpersTestClockInternalFailureEvent`, `V1TestHelpersTestClockReadyEvent`, `V1TopupCanceledEvent`, `V1TopupCreatedEvent`, `V1TopupFailedEvent`, `V1TopupReversedEvent`, `V1TopupSucceededEvent`, `V1TransferCreatedEvent`, `V1TransferReversedEvent`, `V1TransferUpdatedEvent`, `V2CoreHealthIssuingAuthorizationRequestErrorsFiringEvent`, and `V2CoreHealthIssuingAuthorizationRequestErrorsResolvedEvent`
  * Add support for thin event `V2CoreClaimableSandboxCreatedEvent` with related object `V2CoreClaimableSandbox`
  * Add support for thin events `V2MoneyManagementRecipientVerificationCreatedEvent` and `V2MoneyManagementRecipientVerificationUpdatedEvent` with related object `V2MoneyManagementRecipientVerification`
  * Remove support for resources `V2ReportingReportRun`, `V2ReportingReport`
  * Remove support for thin events `V2ReportingReportRunCreatedEvent`, `V2ReportingReportRunFailedEvent`, `V2ReportingReportRunSucceededEvent`, and `V2ReportingReportRunUpdatedEvent` with related object `V2ReportingReportRun`

## 82.6.0-alpha.2 - 2025-09-17
* [#2118](https://github.com/stripe/stripe-go/pull/2118) generate private-preview SDK w/ mid Sept changes
  * Add support for `Get` method on resource `V2CoreClaimableSandbox`
  * Add support for `MonthOfYear` on `V2BillingCadenceBillingCycleMonthParams` and `V2BillingCadenceBillingCycleMonth`
  * Add support for `ClaimedAt`, `ExpiresAt`, `SandboxDetails`, and `Status` on `V2CoreClaimableSandbox`
  * Remove support for `APIKeys` on `V2CoreClaimableSandbox`
  * Add support for new value `current_billing_period_end` on enum `V2BillingIntentActionDeactivateEffectiveAt.Type`
  * Add support for `WillActivateAt` and `WillCancelAt` on `V2BillingPricingPlanSubscriptionServicingStatusTransitions` and `V2BillingRateCardSubscriptionServicingStatusTransitions`
  * Add support for `Category` and `Priority` on `V2BillingServiceActionCreditGrantParams`, `V2BillingServiceActionCreditGrantPerTenantParams`, `V2BillingServiceActionCreditGrantPerTenant`, and `V2BillingServiceActionCreditGrant`
  * Add support for `invoices` on `EventsV2BillingCadenceBilledEvent`
  * Add support for thin events `V2CoreClaimableSandboxClaimedEvent`, `V2CoreClaimableSandboxExpiredEvent`, `V2CoreClaimableSandboxExpiringEvent`, and `V2CoreClaimableSandboxSandboxDetailsOwnerAccountUpdatedEvent` with related object `V2.Core.ClaimableSandbox`
  * Remove support for thin event `V2BillingCadenceErroredEvent` with related object `V2.Billing.Cadence`

## 82.6.0-alpha.1 - 2025-08-27
* [#2110](https://github.com/stripe/stripe-go/pull/2110) Use the right API version 2025-08-27.preview
* [#2106](https://github.com/stripe/stripe-go/pull/2106) Update generated code for private-preview
  * Add support for `AttachCadence` method on resource `Subscription`
  * Add support for `Currency` and `ExternalCustomerID` on `BillingAlertTriggered`
  * Add support for `CustomPricingUnit` on `BillingAlertTriggered`, `BillingCreditBalanceSummaryBalanceAvailableBalance`, `BillingCreditBalanceSummaryBalanceLedgerBalance`, `BillingCreditBalanceTransactionCreditAmount`, `BillingCreditBalanceTransactionDebitAmount`, `BillingCreditGrantAmountParams`, and `BillingCreditGrantAmount`
  * Add support for `Customer` on `BillingAlertListParams`
  * Change type of `BillingAlert.AlertType`, `BillingAlertListParams.AlertType`, and `BillingAlertParams.AlertType` from `literal('usage_threshold')` to `enum('credit_balance_threshold'|'usage_threshold')`
  * Add support for `CreditBalanceThreshold` on `BillingAlertParams` and `BillingAlert`
  * Add support for `BillableItems` on `BillingCreditBalanceSummaryFilterApplicabilityScopeParams`, `BillingCreditGrantApplicabilityConfigScopeParams`, and `BillingCreditGrantApplicabilityConfigScope`
  * Change type of `BillingCreditBalanceSummaryBalanceAvailableBalance.Type`, `BillingCreditBalanceSummaryBalanceLedgerBalance.Type`, `BillingCreditBalanceTransactionCreditAmount.Type`, `BillingCreditBalanceTransactionDebitAmount.Type`, `BillingCreditGrantAmount.Type`, and `BillingCreditGrantAmountParams.Type` from `literal('monetary')` to `enum('custom_pricing_unit'|'monetary')`
  * Add support for `LicenseFeeSubscriptionDetails` and `RateCardSubscriptionDetails` on `InvoiceItemParent` and `InvoiceLineItemParent`
  * Change type of `InvoiceItemParent.Type` from `literal('subscription_details')` to `enum('license_fee_subscription_details'|'rate_card_subscription_details'|'subscription_details')`
  * Add support for `LicenseFeeDetails` and `RateCardRateDetails` on `InvoiceItemPricing` and `InvoiceLineItemPricing`
  * Change type of `InvoiceItemPricing.Type` and `InvoiceLineItemPricing.Type` from `literal('price_details')` to `enum('license_fee_details'|'price_details'|'rate_card_rate_details')`
  * Add support for `BillingCadence` on `InvoiceCreatePreviewParams`, `SubscriptionParams`, and `Subscription`
  * Add support for `BillingCadenceDetails` on `InvoiceParent` and `QuotePreviewInvoiceParent`
  * Add support for new value `billing_cadence_details` on enums `InvoiceParent.Type` and `QuotePreviewInvoiceParent.Type`
  * Add support for new values `license_fee_subscription_details` and `rate_card_subscription_details` on enum `InvoiceLineItemParent.Type`
  * Add support for new resources `V2BillingBillSettingVersion`, `V2BillingBillSetting`, `V2BillingCadence`, `V2BillingCollectionSettingVersion`, `V2BillingCollectionSetting`, `V2BillingCustomPricingUnit`, `V2BillingIntentAction`, `V2BillingIntent`, `V2BillingLicenseFeeSubscription`, `V2BillingLicenseFeeVersion`, `V2BillingLicenseFee`, `V2BillingLicensedItem`, `V2BillingMeteredItem`, `V2BillingPricingPlanComponent`, `V2BillingPricingPlanSubscription`, `V2BillingPricingPlanVersion`, `V2BillingPricingPlan`, `V2BillingProfile`, `V2BillingRateCardRate`, `V2BillingRateCardSubscription`, `V2BillingRateCardVersion`, `V2BillingRateCard`, `V2BillingServiceAction`, `V2CoreClaimableSandbox`, `V2ReportingReportRun`, `V2ReportingReport`, and `V2TaxAutomaticRule`
  * Add support for `Deactivate`, `Find`, `Get`, `New`, and `Update` methods on resource `V2TaxAutomaticRule`
  * Add support for `Get` and `New` methods on resources `V2BillingServiceAction` and `V2ReportingReportRun`
  * Add support for `Get` method on resources `V2BillingLicenseFeeSubscription` and `V2ReportingReport`
  * Add support for `New` method on resource `V2CoreClaimableSandbox`
  * Add support for `Cancel`, `Get`, `List`, `New`, and `Update` methods on resources `V2BillingCadence` and `V2BillingRateCardSubscription`
  * Add support for `Get`, `List`, `New`, and `Update` methods on resources `V2BillingBillSetting`, `V2BillingCollectionSetting`, `V2BillingCustomPricingUnit`, `V2BillingLicenseFee`, `V2BillingLicensedItem`, `V2BillingMeteredItem`, `V2BillingPricingPlan`, `V2BillingProfile`, and `V2BillingRateCard`
  * Add support for `Get` and `List` methods on resources `V2BillingBillSettingVersion`, `V2BillingCollectionSettingVersion`, `V2BillingIntentAction`, `V2BillingLicenseFeeVersion`, `V2BillingPricingPlanSubscription`, `V2BillingPricingPlanVersion`, and `V2BillingRateCardVersion`
  * Add support for `Del`, `Get`, `List`, and `New` methods on resource `V2BillingRateCardRate`
  * Add support for `Del`, `Get`, `List`, `New`, and `Update` methods on resource `V2BillingPricingPlanComponent`
  * Add support for `Cancel`, `Commit`, `Get`, `List`, `New`, `ReleaseReservation`, and `Reserve` methods on resource `V2BillingIntent`
  * Add support for `Changes` on `V2Event`
  * Add support for thin events `V2BillingCadenceBilledEvent`, `V2BillingCadenceCanceledEvent`, `V2BillingCadenceCreatedEvent`, and `V2BillingCadenceErroredEvent` with related object `V2BillingCadence`
  * Add support for thin events `V2BillingLicenseFeeCreatedEvent` and `V2BillingLicenseFeeUpdatedEvent` with related object `V2BillingLicenseFee`
  * Add support for thin event `V2BillingLicenseFeeVersionCreatedEvent` with related object `V2BillingLicenseFeeVersion`
  * Add support for thin events `V2BillingLicensedItemCreatedEvent` and `V2BillingLicensedItemUpdatedEvent` with related object `V2BillingLicensedItem`
  * Add support for thin events `V2BillingMeteredItemCreatedEvent` and `V2BillingMeteredItemUpdatedEvent` with related object `V2BillingMeteredItem`
  * Add support for thin events `V2BillingPricingPlanCreatedEvent` and `V2BillingPricingPlanUpdatedEvent` with related object `V2BillingPricingPlan`
  * Add support for thin events `V2BillingPricingPlanComponentCreatedEvent` and `V2BillingPricingPlanComponentUpdatedEvent` with related object `V2BillingPricingPlanComponent`
  * Add support for thin events `V2BillingPricingPlanSubscriptionCollectionAwaitingCustomerActionEvent`, `V2BillingPricingPlanSubscriptionCollectionCurrentEvent`, `V2BillingPricingPlanSubscriptionCollectionPastDueEvent`, `V2BillingPricingPlanSubscriptionCollectionPausedEvent`, `V2BillingPricingPlanSubscriptionCollectionUnpaidEvent`, `V2BillingPricingPlanSubscriptionServicingActivatedEvent`, `V2BillingPricingPlanSubscriptionServicingCanceledEvent`, and `V2BillingPricingPlanSubscriptionServicingPausedEvent` with related object `V2BillingPricingPlanSubscription`
  * Add support for thin event `V2BillingPricingPlanVersionCreatedEvent` with related object `V2BillingPricingPlanVersion`
  * Add support for thin events `V2BillingRateCardCreatedEvent` and `V2BillingRateCardUpdatedEvent` with related object `V2BillingRateCard`
  * Add support for thin event `V2BillingRateCardRateCreatedEvent` with related object `V2BillingRateCardRate`
  * Add support for thin events `V2BillingRateCardSubscriptionActivatedEvent`, `V2BillingRateCardSubscriptionCanceledEvent`, `V2BillingRateCardSubscriptionCollectionAwaitingCustomerActionEvent`, `V2BillingRateCardSubscriptionCollectionCurrentEvent`, `V2BillingRateCardSubscriptionCollectionPastDueEvent`, `V2BillingRateCardSubscriptionCollectionPausedEvent`, `V2BillingRateCardSubscriptionCollectionUnpaidEvent`, `V2BillingRateCardSubscriptionServicingActivatedEvent`, `V2BillingRateCardSubscriptionServicingCanceledEvent`, and `V2BillingRateCardSubscriptionServicingPausedEvent` with related object `V2BillingRateCardSubscription`
  * Add support for thin event `V2BillingRateCardVersionCreatedEvent` with related object `V2BillingRateCardVersion`
  * Add support for thin events `V2CoreHealthApiErrorFiringEvent`, `V2CoreHealthApiErrorResolvedEvent`, `V2CoreHealthApiLatencyFiringEvent`, `V2CoreHealthApiLatencyResolvedEvent`, `V2CoreHealthAuthorizationRateDropFiringEvent`, `V2CoreHealthAuthorizationRateDropResolvedEvent`, `V2CoreHealthEventGenerationFailureResolvedEvent`, `V2CoreHealthFraudRateIncreasedEvent`, `V2CoreHealthIssuingAuthorizationRequestTimeoutFiringEvent`, `V2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEvent`, `V2CoreHealthPaymentMethodErrorFiringEvent`, `V2CoreHealthPaymentMethodErrorResolvedEvent`, `V2CoreHealthTrafficVolumeDropFiringEvent`, `V2CoreHealthTrafficVolumeDropResolvedEvent`, `V2CoreHealthWebhookLatencyFiringEvent`, and `V2CoreHealthWebhookLatencyResolvedEvent`
  * Add support for thin events `V2ReportingReportRunCreatedEvent`, `V2ReportingReportRunFailedEvent`, `V2ReportingReportRunSucceededEvent`, and `V2ReportingReportRunUpdatedEvent` with related object `V2ReportingReportRun`
  * Add support for error type `RateLimitError`
