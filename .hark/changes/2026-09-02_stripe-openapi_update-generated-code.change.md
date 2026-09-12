---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2420
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.5.0-alpha.2
---

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
