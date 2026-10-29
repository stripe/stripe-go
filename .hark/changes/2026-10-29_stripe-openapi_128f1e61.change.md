---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2479
semver_level: minor
is_stripe_api_change: true
released_in_version: 87.1.0
---

* Add support for new resources `V2CoreVaultGBBankAccount`, `V2CoreVaultUSBankAccount`, `V2MoneyManagementAdjustment`, `V2MoneyManagementFinancialAccount`, `V2MoneyManagementFinancialAddressCreditSimulation`, `V2MoneyManagementFinancialAddress`, `V2MoneyManagementInboundTransfer`, `V2MoneyManagementOutboundPayment`, `V2MoneyManagementOutboundSetupIntent`, `V2MoneyManagementPayoutIntent`, `V2MoneyManagementPayoutMethod`, `V2MoneyManagementPayoutMethodsBankAccountSpec`, `V2MoneyManagementReceivedCredit`, `V2MoneyManagementRecipientVerification`, `V2MoneyManagementTransactionEntry`, and `V2MoneyManagementTransaction`
* Add support for `Acknowledge`, `Get`, `New`, and `RecipientVerifications` methods on resource `V2MoneyManagementRecipientVerification`
* Add support for `Get` and `List` methods on resources `V2MoneyManagementAdjustment`, `V2MoneyManagementReceivedCredit`, and `V2MoneyManagementTransactionEntry`
* Add support for `Get`, `List`, `RefreshRegulatoryReceipt`, and `Update` methods on resource `V2MoneyManagementTransaction`
* Add support for `Credit` method on resource `V2MoneyManagementFinancialAddressCreditSimulation`
* Add support for `Get` method on resource `V2MoneyManagementPayoutMethodsBankAccountSpec`
* Add support for `Archive`, `Get`, `List`, and `Unarchive` methods on resource `V2MoneyManagementPayoutMethod`
* Add support for `Cancel`, `Confirm`, `Get`, `List`, `New`, `RefreshFxQuote`, `RefreshRecipientVerification`, and `Update` methods on resource `V2MoneyManagementPayoutIntent`
* Add support for `Cancel`, `Get`, `List`, `New`, and `Update` methods on resource `V2MoneyManagementOutboundSetupIntent`
* Add support for `Cancel`, `Get`, `List`, and `New` methods on resource `V2MoneyManagementOutboundPayment`
* Add support for `Get`, `List`, and `New` methods on resources `V2MoneyManagementFinancialAddress` and `V2MoneyManagementInboundTransfer`
* Add support for `Get`, `List`, and `Update` methods on resource `V2MoneyManagementFinancialAccount`
* Add support for `Archive`, `ConfirmMicrodeposits`, `Get`, `List`, `New`, `SendMicrodeposits`, `Unarchive`, and `Update` methods on resource `V2CoreVaultUSBankAccount`
* Add support for `Archive`, `Get`, `List`, `New`, and `Unarchive` methods on resource `V2CoreVaultGBBankAccount`
* Add support for new values `get_fund` and `nhil` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationTaxBreakdownTaxRateDetails.TaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetails.TaxType`
* Add support for `ReportingChart`, `TerminalHardwareOrders`, and `TerminalHardwareShop` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for `WeroPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `PhysicalAddress` on `AccountCompanyParams`, `AccountCompany`, and `TokenAccountCompanyParams`
* Add support for `ProofOfBusinessPresence` on `AccountDocumentsParams`
* Add support for `SubscriptionPause` on `BillingPortalConfigurationFeaturesParams`, `BillingPortalConfigurationFeatures`, `BillingPortalSessionFlowDataParams`, and `BillingPortalSessionFlow`
* Add support for `AfterExpiration` on `BillingPortalSessionParams` and `BillingPortalSession`
* Add support for new value `subscription_pause` on enum `BillingPortalSessionFlow.Type`
* Add support for `SetupCredentialUsage` on `ChargePaymentMethodDetailsCard`, `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCard`, `SetupIntentConfirmPaymentMethodOptionsCardParams`, `SetupIntentPaymentMethodOptionsCardParams`, and `SetupIntentPaymentMethodOptionsCard`
* Add support for `StoredCredentialUsage` on `ChargePaymentMethodDetailsCard`, `PaymentAttemptRecordPaymentMethodDetailsCard`, `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCard`, and `PaymentRecordPaymentMethodDetailsCard`
* Add support for new value `ousd` on enums `ChargePaymentMethodDetailsCrypto.TokenCurrency`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.TokenCurrency`, and `PaymentRecordPaymentMethodDetailsCrypto.TokenCurrency`
* Add support for `AccountTaxIDs`, `CustomFields`, `Description`, and `Footer` on `CheckoutSessionSubscriptionDataInvoiceSettingsParams`, `PaymentLinkSubscriptionDataInvoiceSettingsParams`, and `PaymentLinkSubscriptionDataInvoiceSettings`
* Add support for `PaymentIntentData` on `CheckoutSessionParams`
* Add support for new values `au_acn`, `au_wpn`, `bw_vat`, `ci_vat`, `cz_dic`, `fr_siren`, `fr_siret`, `fr_spi`, `gh_tin`, `is_kt`, `md_cf`, `mu_vat`, `mz_nuit`, `nz_bn`, `pt_nif`, `sm_coe`, `tt_bir`, and `tt_vat` on enums `CheckoutSessionCustomerDetailsTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, `TaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
* Add support for new value `hi-IN` on enum `CheckoutSession.Locale`
* Add support for new values `fednow` and `rtp` on enum `CustomerCashBalanceTransactionFundedBankTransferUsBankTransfer.Network`
* Add support for `IntendedSubmissionMethod` on `DisputeParams` and `Dispute`
* Add support for `Appeal` on `DisputeEvidenceParams` and `DisputeEvidence`
* Add support for `SubmissionMethod` on `DisputeEvidenceDetails`
* Add support for new values `appeal_needs_response` and `appeal_under_review` on enum `Dispute.Status`
* Add support for new values `three_d_secure.authentication.canceled`, `three_d_secure.authentication.challenge_started`, `three_d_secure.authentication.errored`, `three_d_secure.authentication.failed`, `three_d_secure.authentication.requires_challenge`, `three_d_secure.authentication.requires_submission`, and `three_d_secure.authentication.succeeded` on enum `Event.Type`
* Add support for `PauseSchedules` on `InvoiceCreatePreviewScheduleDetailsParams`
* Add support for new value `sequra` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `FinancialAccount` on `IssuingCardParams`
* Add support for `DeliveredAt` and `ShippedAt` on `IssuingCardShipping`
* Add support for new value `scalapay` on enum `PaymentLink.PaymentMethodTypes`
* Add support for `ReferenceStatus` and `Reference` on `RefundDestinationDetailsAfterpayClearpay`, `RefundDestinationDetailsKlarna`, and `RefundDestinationDetailsPaypal`
* Add support for `Bw`, `Ci`, `Gh`, and `Mu` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
* Add support for `UtilityUsersTax` on `TaxRegistrationCountryOptionsUsParams` and `TaxRegistrationCountryOptionsUs`
* Add support for new values `digital_excise_tax` and `utility_users_tax` on enum `TaxRegistrationCountryOptionsUs.Type`
* Add support for new values `admissions_tax`, `attendance_tax`, `entertainment_tax`, `get_fund`, `gross_receipts_tax`, `hospitality_tax`, `luxury_tax`, `nhil`, `resort_tax`, and `tourism_tax` on enum `TaxRate.TaxType`
* Add support for `ReaderSecurity` and `StripeT600` on `TerminalConfigurationParams` and `TerminalConfiguration`
* Add support for new values `simulated_stripe_t600` and `stripe_t600` on enum `TerminalReader.DeviceType`
* Add support for `BankAccounts` and `Cards` on `V2CoreAccountConfigurationRecipientCapabilitiesParams` and `V2CoreAccountConfigurationRecipientCapabilities`
* Add support for `PayoutMethods` on `V2CoreAccountDefaultsParams` and `V2CoreAccountDefaults`
* Add support for `Alternatives` on `V2CoreAccountFutureRequirements` and `V2CoreAccountRequirements`
* Add support for new values `bank_accounts.ach`, `bank_accounts.fedwire`, `bank_accounts.fps`, `bank_accounts.rtp`, `bank_accounts.sepa`, and `bank_accounts.swift` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for `Customer` and `Subscription` on `EventsV1InvoiceUpcomingEvent`
* Add support for new values `bank_accounts.ach`, `bank_accounts.fedwire`, `bank_accounts.fps`, `bank_accounts.rtp`, `bank_accounts.sepa`, and `bank_accounts.swift` on enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for snapshot events `EventTypeThreeDSecureAuthenticationCanceled`, `EventTypeThreeDSecureAuthenticationChallengeStarted`, `EventTypeThreeDSecureAuthenticationErrored`, `EventTypeThreeDSecureAuthenticationFailed`, `EventTypeThreeDSecureAuthenticationRequiresChallenge`, `EventTypeThreeDSecureAuthenticationRequiresSubmission`, and `EventTypeThreeDSecureAuthenticationSucceeded` with resource `ThreeDSecureAuthentication`
* Add support for event notifications `V1CustomerDiscountSourceChangedEvent`, `V1EntitlementsActiveEntitlementSummaryEntitlementAddedEvent`, and `V1EntitlementsActiveEntitlementSummaryEntitlementRemovedEvent`
* Add support for event notifications `V1CustomerSubscriptionActiveEvent`, `V1CustomerSubscriptionCancelAtChangedEvent`, `V1CustomerSubscriptionDefaultPaymentMethodChangedEvent`, `V1CustomerSubscriptionIncompleteEvent`, `V1CustomerSubscriptionIncompleteExpiredEvent`, `V1CustomerSubscriptionItemAddedEvent`, `V1CustomerSubscriptionItemPriceChangedEvent`, `V1CustomerSubscriptionItemQuantityChangedEvent`, `V1CustomerSubscriptionItemRemovedEvent`, `V1CustomerSubscriptionItemTaxRatesChangedEvent`, `V1CustomerSubscriptionPastDueEvent`, `V1CustomerSubscriptionTrialingEvent`, and `V1CustomerSubscriptionUnpaidEvent` with related object `Subscription`
* Add support for event notifications `V1FinancialConnectionsAccountBalanceRefreshFailedEvent`, `V1FinancialConnectionsAccountBalanceRefreshSucceededEvent`, `V1FinancialConnectionsAccountInferredBalancesRefreshFailedEvent`, `V1FinancialConnectionsAccountInferredBalancesRefreshSucceededEvent`, `V1FinancialConnectionsAccountOwnershipRefreshFailedEvent`, `V1FinancialConnectionsAccountOwnershipRefreshSucceededEvent`, `V1FinancialConnectionsAccountRiskIntelligenceRefreshFailedEvent`, `V1FinancialConnectionsAccountRiskIntelligenceRefreshSucceededEvent`, `V1FinancialConnectionsAccountTransactionsRefreshFailedEvent`, and `V1FinancialConnectionsAccountTransactionsRefreshSucceededEvent` with related object `FinancialConnectionsAccount`
* Add support for event notifications `V1InvoiceDiscountAddedEvent`, `V1InvoiceDiscountRemovedEvent`, `V1InvoiceLineItemAddedEvent`, `V1InvoiceLineItemPricingChangedEvent`, `V1InvoiceLineItemQuantityChangedEvent`, and `V1InvoiceLineItemRemovedEvent` with related object `Invoice`
* Add support for event notifications `V1InvoicePaymentCanceledEvent` and `V1InvoicePaymentOpenEvent` with related object `InvoicePayment`
* Add support for event notifications `V1InvoiceitemAmountChangedEvent`, `V1InvoiceitemDiscountAddedEvent`, and `V1InvoiceitemDiscountRemovedEvent` with related object `InvoiceItem`
* Add support for event notifications `V1PriceActivatedEvent`, `V1PriceDeactivatedEvent`, `V1PriceLookupKeyChangedEvent`, and `V1PriceTaxBehaviorChangedEvent` with related object `Price`
* Add support for event notifications `V1ProductActivatedEvent`, `V1ProductDeactivatedEvent`, `V1ProductDefaultPriceChangedEvent`, and `V1ProductTaxCodeChangedEvent` with related object `Product`
* Add support for event notification `V1PromotionCodeFullyRedeemedEvent` with related object `PromotionCode`
* Add support for event notifications `V1SubscriptionSchedulePhaseAddedEvent`, `V1SubscriptionSchedulePhaseDefaultTaxRatesChangedEvent`, `V1SubscriptionSchedulePhaseEndDateChangedEvent`, `V1SubscriptionSchedulePhaseRemovedEvent`, and `V1SubscriptionSchedulePhaseTrialEndChangedEvent` with related object `SubscriptionSchedule`
* Add support for event notification `V2MoneyManagementAdjustmentCreatedEvent` with related object `V2MoneyManagementAdjustment`
* Add support for event notifications `V2MoneyManagementFinancialAccountCreatedEvent` and `V2MoneyManagementFinancialAccountUpdatedEvent` with related object `V2MoneyManagementFinancialAccount`
* Add support for event notifications `V2MoneyManagementFinancialAddressActivatedEvent` and `V2MoneyManagementFinancialAddressFailedEvent` with related object `V2MoneyManagementFinancialAddress`
* Add support for event notifications `V2MoneyManagementInboundTransferAvailableEvent`, `V2MoneyManagementInboundTransferBankDebitFailedEvent`, `V2MoneyManagementInboundTransferBankDebitProcessingEvent`, `V2MoneyManagementInboundTransferBankDebitQueuedEvent`, `V2MoneyManagementInboundTransferBankDebitReturnedEvent`, and `V2MoneyManagementInboundTransferBankDebitSucceededEvent` with related object `V2MoneyManagementInboundTransfer`
* Add support for event notifications `V2MoneyManagementOutboundPaymentCanceledEvent`, `V2MoneyManagementOutboundPaymentCreatedEvent`, `V2MoneyManagementOutboundPaymentFailedEvent`, `V2MoneyManagementOutboundPaymentPostedEvent`, `V2MoneyManagementOutboundPaymentReturnedEvent`, `V2MoneyManagementOutboundPaymentUnderReviewEvent`, and `V2MoneyManagementOutboundPaymentUpdatedEvent` with related object `V2MoneyManagementOutboundPayment`
* Add support for event notifications `V2MoneyManagementPayoutIntentCanceledEvent`, `V2MoneyManagementPayoutIntentCreatedEvent`, `V2MoneyManagementPayoutIntentPendingEvent`, `V2MoneyManagementPayoutIntentPostedEvent`, `V2MoneyManagementPayoutIntentProcessingEvent`, `V2MoneyManagementPayoutIntentRequiresActionEvent`, and `V2MoneyManagementPayoutIntentUpdatedEvent` with related object `V2MoneyManagementPayoutIntent`
* Add support for event notifications `V2MoneyManagementPayoutMethodCreatedEvent` and `V2MoneyManagementPayoutMethodUpdatedEvent` with related object `V2MoneyManagementPayoutMethod`
* Add support for event notifications `V2MoneyManagementReceivedCreditAvailableEvent`, `V2MoneyManagementReceivedCreditFailedEvent`, `V2MoneyManagementReceivedCreditReturnedEvent`, and `V2MoneyManagementReceivedCreditSucceededEvent` with related object `V2MoneyManagementReceivedCredit`
* Add support for event notifications `V2MoneyManagementRecipientVerificationCreatedEvent` and `V2MoneyManagementRecipientVerificationUpdatedEvent` with related object `V2MoneyManagementRecipientVerification`
* Add support for event notifications `V2MoneyManagementTransactionCreatedEvent` and `V2MoneyManagementTransactionUpdatedEvent` with related object `V2MoneyManagementTransaction`
* Add support for error types `AlreadyCanceledError`, `BlockedByStripeError`, `CannotProceedError`, `CapabilityNotEnabledError`, `ControlledByAlternateResourceError`, `FeatureNotEnabledError`, `FinancialAccountNotOpenError`, `FxQuoteNeedsRefreshError`, `InsufficientFundsError`, `InvalidPaymentMethodError`, `InvalidPayoutMethodError`, `InvalidVaultedCredentialError`, `NotCancelableError`, `QuotaExceededError`, `RecipientNotNotifiableError`, `VerificationAttemptFailedError`, `VerificationExpiredError`, and `VerificationNotInitiatedError`
* Add support for error codes `smart_disputes_not_available`, `subscription_update_trial_after_billing_cycle_anchor`, `three_d_secure_authentication_not_supported`, and `three_d_secure_authentication_upstream_error` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, `StripeError`, and `TerminalReaderActionApiError`
