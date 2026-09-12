---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2345
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.2.0-alpha.1
---

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
