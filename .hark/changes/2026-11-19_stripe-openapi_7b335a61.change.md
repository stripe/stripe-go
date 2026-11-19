---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2497
semver_level: major
is_stripe_api_change: true
released_in_version: 87.3.0-alpha.1
---

* Add support for new resources `V2BillingGroupInvoice`, `V2BillingGroupSpendModifier`, and `V2BillingGroup`
* ⚠️ Remove support for resources `V2BillingCadenceSpendModifier` and `V2BillingContractPricingLineQuantityChange`
* Add support for `New` and `Update` methods on resource `InvoiceRenderingTemplate`
* Add support for `Cancel`, `Get`, `List`, `New`, and `Update` methods on resource `V2BillingGroup`
* Add support for `Get` method on resource `V2BillingGroupSpendModifier`
* Add support for `New` method on resource `V2BillingGroupInvoice`
* ⚠️ Remove support for `Get` method on resource `V2BillingCadenceSpendModifier`
* ⚠️ Remove support for `BillingCadence` on `BillingAlertSpendThresholdFiltersParams`, `BillingAlertSpendThresholdFilters`, `BillingAlertTriggeredGroupBy`, `InvoiceCreatePreviewParams`, `InvoiceListParams`, `SubscriptionListParams`, `SubscriptionParams`, `Subscription`, `V2BillingPricingPlanSubscriptionListParams`, `V2BillingPricingPlanSubscription`, `V2BillingRateCardSubscriptionListParams`, `V2BillingRateCardSubscriptionParams`, and `V2BillingRateCardSubscription`
* ⚠️ Add support for new value `billing_group` on enum `BillingAlertTriggeredGroupBy.Type`
* ⚠️ Remove support for value `billing_cadence` from enum `BillingAlertTriggeredGroupBy.Type`
* Add support for `ScheduledFee` on `V2BillingContractActionAddTermParams`, `V2BillingContractActionUpdateTermParams`, `V2BillingContractTermParams`, and `V2BillingContractTerm`
* ⚠️ Remove support for `AdHocProduct` on `V2BillingContractActionAddTermParams`, `V2BillingContractActionUpdateTermParams`, `V2BillingContractTermParams`, and `V2BillingContractTerm`
* ⚠️ Add support for new value `scheduled_fee` on enum `V2BillingContractTerm.Type`
* ⚠️ Remove support for value `ad_hoc_product` from enum `V2BillingContractTerm.Type`
* ⚠️ Remove support for `Cadence` on `BillingAlertNotification`, `BillingBillingAlertAlertNotificationListParams`, `V2BillingIntentParams`, and `V2BillingIntent`
* ⚠️ Add support for new value `billing_group` on enum `BillingAlertSpendThreshold.GroupBy`
* ⚠️ Remove support for value `billing_cadence` from enum `BillingAlertSpendThreshold.GroupBy`
* ⚠️ Remove support for `PricingToken` on `CheckoutSessionParams`
* Add support for `PassThroughFeeCost` on `CheckoutSession`
* Add support for new value `dealerfx` on enums `CheckoutSessionAutomaticSurcharge.Provider` and `PaymentLinkAutomaticSurcharge.Provider`
* Add support for `AmountPassThroughFee` on `CheckoutSessionTotalDetails`
* Add support for `LinkSharedData` on `IdentityVerificationSessionParams`
* Add support for `LineItemPolicy`, `RenderingSettings`, and `SummaryItemPolicies` on `InvoiceRenderingTemplate`
* Add support for `BillingGroup` on `InvoiceCreatePreviewParams`, `InvoiceListParams`, `V2BillingIntentParams`, `V2BillingIntent`, `V2BillingPricingPlanSubscriptionListParams`, `V2BillingPricingPlanSubscription`, `V2BillingRateCardSubscriptionListParams`, and `V2BillingRateCardSubscription`
* Add support for `AppliedCreditMemos` on `Invoice` and `QuotePreviewInvoice`
* Add support for `BillingGroupDetails` on `InvoiceParent` and `QuotePreviewInvoiceParent`
* ⚠️ Remove support for `BillingCadenceDetails` on `InvoiceParent` and `QuotePreviewInvoiceParent`
* ⚠️ Add support for new value `billing_group_details` on enums `InvoiceParent.Type` and `QuotePreviewInvoiceParent.Type`
* ⚠️ Remove support for value `billing_cadence_details` from enums `InvoiceParent.Type` and `QuotePreviewInvoiceParent.Type`
* Add support for `TokenServiceProviderID` on `IssuingAuthorizationTokenDetailsNetworkDataVisa` and `IssuingTokenNetworkDataVisa`
* Add support for `AmountType`, `Amount`, `Currency`, `Description`, `EndDate`, `IntervalCount`, `Interval`, `Reference`, and `StartDate` on `MandatePaymentMethodDetailsCardIndia`
* Add support for `ThreeDSecure` on `PaymentAttemptRecordReportFailedPaymentMethodDetailsCardParams` and `PaymentRecordReportPaymentAttemptFailedPaymentMethodDetailsCardParams`
* Add support for `ProviderSignatureData` and `ProviderSignature` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentAadeDataParams`, `PaymentIntentPaymentMethodOptionsCardPresentAadeDataParams`, and `PaymentIntentPaymentMethodOptionsCardPresentAadeData`
* ⚠️ Remove support for `MarkData` and `SignedMark` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentAadeDataParams`, `PaymentIntentPaymentMethodOptionsCardPresentAadeDataParams`, and `PaymentIntentPaymentMethodOptionsCardPresentAadeData`
* Add support for `RequestMulticapture` on `PaymentIntentConfirmPaymentMethodOptionsVippsParams`, `PaymentIntentPaymentMethodOptionsVippsParams`, and `PaymentIntentPaymentMethodOptionsVipps`
* Add support for `IsFreeTrialTransaction` on `PaymentIntentConfirmRadarOptionsParams` and `PaymentIntentRadarOptionsParams`
* Add support for `RadarOptions` on `PaymentIntent`
* Add support for `AeonPay`, `AuPay`, `Dbarai`, and `Merpay` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
* Add support for `RequestedSignals` on `RadarPaymentEvaluationParams` and `RadarPaymentEvaluation`
* Add support for `GiftCard` on `RefundDestinationDetails`
* Add support for `DaysAfterFulfillment` on `ReservePlanRollingRelease`
* Add support for `Rendering` on `SubscriptionScheduleDefaultSettingsInvoiceSettingsParams`
* Add support for `DebitBehavior` on `SubscriptionPaymentSettingsPaymentMethodOptionsCardParams`, `SubscriptionPaymentSettingsPaymentMethodOptionsCard`, `SubscriptionPaymentSettingsPaymentMethodOptionsPixParams`, `SubscriptionPaymentSettingsPaymentMethodOptionsPix`, `SubscriptionPaymentSettingsPaymentMethodOptionsUpiParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsUpi`
* Add support for `PresentmentDetails` on `SubscriptionPendingUpdate`
* ⚠️ Remove support for `DeviceHealthDetails` and `DeviceHealth` on `TerminalReader`
* Add support for `BillingGroupData` on `V2BillingIntentParams` and `V2BillingIntent`
* ⚠️ Remove support for `CadenceData` on `V2BillingIntentParams` and `V2BillingIntent`
* Add support for new value `group` on enum `V2BillingIntentActionApplyInvoiceDiscountRule.AppliesTo`
* Add support for new value `account_dashboard` on enum `V2CoreAccountLinkUseCase.Type`
* Add support for `TouchNGoPayments`, `TruemoneyPayments`, `WeChatPayMobileWebPayments`, `WeChatPayPayments`, and `WeroPayments` on `V2CoreAccountConfigurationMerchantCapabilitiesParams` and `V2CoreAccountConfigurationMerchantCapabilities`
* Add support for `BusinessDemandDepositAccountFifthThirdStorage` on `V2CoreAccountConfigurationMoneyManagerCapabilitiesParams` and `V2CoreAccountConfigurationMoneyManagerCapabilities`
* Add support for `Tips` on `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsParams` and `V2CoreAccountConfigurationRecipientCapabilitiesBankAccounts`
* Add support for `BusinessDemandDepositAccountFifthThird` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams` and `V2CoreAccountIdentityAttestationsTermsOfService`
* Add support for `Channel` on `V2CoreClaimableSandboxParams` and `V2CoreClaimableSandbox`
* ⚠️ Remove support for `AppChannel` on `V2CoreClaimableSandboxParams` and `V2CoreClaimableSandbox`
* Add support for `ReturnURL` on `V2CoreClaimableSandboxOnboardingLinkDetailsParams`, `V2CoreClaimableSandboxOnboardingLinkDetails`, and `V2CoreClaimableSandboxRenewOnboardingLinkOnboardingLinkDetailsParams`
* ⚠️ Change type of `V2CoreClaimableSandboxSandboxDetailsApiKeys.Mcp` from `string` to `an object`
* ⚠️ Change type of `V2CoreClaimableSandboxSandboxDetailsApiKeys.Publishable` from `string` to `an object`
* ⚠️ Change type of `V2CoreClaimableSandboxSandboxDetailsApiKeys.Secret` from `string` to `an object`
* Add support for `CustomPaymentMethodAdapterDegraded` on `V2CoreHealthAlertHistoryEntry` and `V2CoreHealthAlert`
* Add support for new value `custom_payment_method_adapter_degraded` on enums `V2CoreHealthAlert.Type` and `V2CoreHealthAlertHistoryEntry.Type`
* Add support for `FundsHoldingStructure` on `V2MoneyManagementFinancialAccountStorageParams` and `V2MoneyManagementFinancialAccountStorage`
* Add support for `AUBECSDebit` on `V2MoneyManagementInboundTransferFromPaymentMethod`
* Add support for new value `tips` on enums `V2MoneyManagementOutboundPaymentQuoteEstimatedFeeNetworkFeeDetails.Network` and `V2MoneyManagementPayoutIntentEstimatedFeeNetworkFeeDetails.Network`
* Add support for new value `tips` on enums `V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccount.PreferredNetworks`, `V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccount.PreferredNetworks`, `V2MoneyManagementOutboundTransferToPayoutMethodOptionsBankAccount.PreferredNetworks`, and `V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccount.PreferredNetworks`
* Add support for `Network` on `V2MoneyManagementReceivedDebitBankTransfer`
* Add support for new values `inbound_funding_reversal` and `inbound_funding` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* Add support for new value `inbound_funding` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
* Add support for `SettlementTiming` on `V2MoneyManagementFundingIntentParams`
* Add support for `Include` on `V2CoreClaimableSandboxParams`
* ⚠️ Add support for `BillingGroup` on `V2BillingRateCardSubscriptionParams`
* ⚠️ Remove support for `OneTimeFees`, `PricingLines`, and `PricingOverrides` on `V2BillingContractParams`
* Add support for `ContractNumber` on `V2BillingContractParams`
* ⚠️ Remove support for `OneTimeFeeActions`, `PricingLineActions`, and `PricingOverrideActions` on `V2BillingContractParams`
* ⚠️ Remove support for `CancelPricingLines` on `V2BillingContractCancelParams`
* Add support for new values `business_demand_deposit_account_fifth_third_storage.inbound.usd` and `business_demand_deposit_account_fifth_third_storage.outbound.usd` on enum `EventsV2CoreAccountIncludingConfigurationMoneyManagerCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for new value `bank_accounts.tips` on enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for `Charge` on `EventsV2PaymentsOffSessionPaymentAttemptFailedEvent`, `EventsV2PaymentsOffSessionPaymentFailedEvent`, and `EventsV2PaymentsOffSessionPaymentSucceededEvent`
* Add support for `LastAttemptError` on `EventsV2PaymentsOffSessionPaymentAttemptFailedEvent` and `EventsV2PaymentsOffSessionPaymentFailedEvent`
* Add support for `PaymentAttemptRecord` on `EventsV2PaymentsOffSessionPaymentSucceededEvent`
* Add support for event notifications `V2BillingContractTermPricingOverrideEndedEvent`, `V2BillingContractTermPricingOverrideStartedEvent`, `V2BillingContractTermRecurringPriceEndedEvent`, `V2BillingContractTermRecurringPriceStartedEvent`, `V2BillingContractTermScheduledFeeServicePeriodEndedEvent`, and `V2BillingContractTermScheduledFeeServicePeriodStartedEvent` with related object `V2BillingContractTerm`
* Add support for event notifications `V2CoreHealthCustomPaymentMethodAdapterDegradedFiringEvent` and `V2CoreHealthCustomPaymentMethodAdapterDegradedResolvedEvent`
* Add support for event notifications `V2SignalsPaymentRetryEvaluationCanceledEvent`, `V2SignalsPaymentRetryEvaluationCompletedEvent`, `V2SignalsPaymentRetryEvaluationCreatedEvent`, `V2SignalsPaymentRetryEvaluationExpiredEvent`, and `V2SignalsPaymentRetryEvaluationUpdatedEvent` with related object `V2SignalsPaymentRetryEvaluation`
* ⚠️ Remove support for event notifications `V2BillingCadenceBilledEvent`, `V2BillingCadenceCanceledEvent`, and `V2BillingCadenceCreatedEvent` with related object `V2BillingCadence`
* ⚠️ Add error type `RateLimitError` on `V2CoreClaimableSandbox#New`
* Add support for error code `scheduled_cancellation_conflicts_with_updated_terms` on `CannotProceedError`
