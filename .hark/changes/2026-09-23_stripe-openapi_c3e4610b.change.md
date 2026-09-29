---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2439
semver_level: major
is_stripe_api_change: true
released_in_version: 86.5.0-alpha.5
---

* Add support for new resources `FinancialConnectionsConsent`, `V2MoneyManagementFinancialAccountWalletExportCredentials`, `V2MoneyManagementFinancialAccountWalletExport`, `V2ProvisioningEligibility`, `V2ProvisioningPaymentMethodRequest`, `V2ProvisioningPaymentProfile`, `V2ProvisioningProject`, `V2ProvisioningProviderConnectionRequest`, `V2ProvisioningProviderConnection`, `V2ProvisioningProviderServiceDetail`, `V2ProvisioningProvider`, and `V2ProvisioningResource`
* ⚠️ Remove support for resource `RadarBillingEvaluation`
* Add support for `Get` and `New` methods on resource `FinancialConnectionsConsent`
* ⚠️ Remove support for `New` method on resource `RadarBillingEvaluation`
* Add support for `Get`, `Link`, `New`, `Remove`, `RotateCredentials`, `SubmitInformation`, `Unlink`, and `Update` methods on resource `V2ProvisioningResource`
* Add support for `Get`, `New`, and `SubmitInformation` methods on resource `V2ProvisioningProviderConnectionRequest`
* Add support for `List` and `Unlink` methods on resource `V2ProvisioningProviderConnection`
* Add support for `New` method on resources `V2ProvisioningPaymentMethodRequest` and `V2ProvisioningProject`
* Add support for `Get` and `UpdateLimit` methods on resource `V2ProvisioningPaymentProfile`
* Add support for `Get` method on resource `V2ProvisioningEligibility`
* Add support for `ExportCredentials` and `Get` methods on resource `V2MoneyManagementFinancialAccountWalletExport`
* Add support for new values `invalid_address_cmra_address` and `invalid_address_registered_agent_address` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
* Add support for new values `digital_excise_tax`, `recycling_fee`, and `utility_users_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetails.TaxType`, `TaxCalculationTaxBreakdownTaxRateDetails.TaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetails.TaxType`
* Add support for `BLIKRecurringPayments` on `AccountCapabilities`
* Add support for `Capital` on `AccountSettings`
* Add support for `PayoutMethod` on `BalanceInstantAvailableNetAvailable`
* Add support for `DestinationCurrency` on `BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyParams` and `BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrency`
* Add support for `TotalDueAmount` on `CapitalFinancingOfferAcceptedTerms` and `CapitalFinancingSummaryDetails`
* Add support for `IncrementalIntervalTargetAmount` and `StartsAt` on `CapitalFinancingSummaryDetailsCurrentRepaymentInterval`
* Add support for `SetupCredentialUsage` on `ChargePaymentMethodDetailsCard`, `PaymentIntentPaymentMethodOptionsCard`, and `SetupIntentPaymentMethodOptionsCard`
* Add support for `StoredCredentialUsage` on `ChargePaymentMethodDetailsCard`, `PaymentAttemptRecordPaymentMethodDetailsCard`, `PaymentIntentPaymentMethodOptionsCard`, and `PaymentRecordPaymentMethodDetailsCard`
* Add support for `PaymentMethodOptions` on `CheckoutSessionApproveParams`
* Add support for `PaymentReservation` on `CheckoutSession`
* Add support for `Custom` on `CheckoutSessionCurrentAttemptPaymentMethodDetails`
* Change type of `CheckoutSessionItem.Subscription` from `nullable(PaymentPagesCheckoutSessionSubscription)` to `PaymentPagesCheckoutSessionSubscription`
* Add support for `PaymentMethodPreselect` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`
* Add support for `BIC`, `IBANLast4`, and `Network` on `CustomerCashBalanceTransactionFundedBankTransferGbBankTransfer`
* Add support for new values `apps.install.created`, `apps.install.deleted`, and `apps.install.updated` on enum `Event.Type`
* Add support for new values `expired` and `pending` on enum `FinancialConnectionsAccountAccountNumbers.Status`
* Add support for `PreCollectedConsent` on `FinancialConnectionsSession`
* Add support for `FinancialActivity` on `FinancialConnectionsTransactionClassifications`
* ⚠️ Remove support for `Credit` on `FinancialConnectionsTransactionClassifications`
* Add support for new value `touch_n_go` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `Fuels` on `IssuingAuthorization`
* ⚠️ Change type of `MandatePaymentMethodDetailsBlik.Type` from `enum('off_session'|'on_session')` to `literal('off_session')`
* Add support for `MandateOptions` on `PaymentIntentPaymentMethodOptionsBlik`
* ⚠️ Change type of `PaymentIntentPaymentMethodOptionsBlik.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
* Add support for `PayoutMethodOptions` on `Payout`
* Add support for new value `rerouted` on enum `RadarPaymentEvaluationOutcome.Type`
* Add support for `BLIK` on `SetupAttemptPaymentMethodDetails` and `SetupIntentPaymentMethodOptions`
* Add support for `CancelAtPeriodEnd` on `SubscriptionPendingUpdate`
* Add support for `TamperState` on `TerminalReaderListParams`
* Add support for `OriginPaymentMethodOptions` on `TreasuryInboundTransferParams`
* Add support for `ACH` on `TreasuryInboundTransferOriginPaymentMethodDetailsUsBankAccount`
* Add support for `CollectionStatusTransitions` and `CollectionStatus` on `V2BillingContract`
* Add support for new value `developer` on enums `EventsV2CoreAccountLinkReturnedEvent.Configurations`, `V2CoreAccountLinkUseCaseAccountOnboarding.Configurations`, and `V2CoreAccountLinkUseCaseAccountUpdate.Configurations`
* Add support for new value `developer` on enum `V2CoreAccount.AppliedConfigurations`
* Add support for `Developer` on `V2CoreAccountConfigurationParams` and `V2CoreAccountConfiguration`
* Add support for new value `apple_pay` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* Add support for new value `projects` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new value `developer` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Configuration` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Configuration`
* Add support for `SkipExportableBalances` on `V2MoneyManagementFinancialAccountCloseForwardingSettingsParams` and `V2MoneyManagementFinancialAccountStatusDetailsClosedForwardingSettings`
* Add support for `Crypto` on `V2MoneyManagementFinancialAccountStorageParams` and `V2MoneyManagementFinancialAccountStorage`
* Add support for `Addenda` on `V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsAchParams`, `V2MoneyManagementOutboundPaymentQuoteToPayoutMethodOptionsBankAccountPreferredNetworkOptionsAch`, `V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsAchParams`, `V2MoneyManagementOutboundPaymentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsAch`, `V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsAchParams`, and `V2MoneyManagementPayoutIntentToPayoutMethodOptionsBankAccountPreferredNetworkOptionsAch`
* Add support for `PreferredNetworkOptions` on `V2MoneyManagementOutboundTransferToPayoutMethodOptionsBankAccountParams` and `V2MoneyManagementOutboundTransferToPayoutMethodOptionsBankAccount`
* Add support for `ApplePay` on `V2MoneyManagementOutboundSetupIntentPayoutMethodDataParams` and `V2MoneyManagementPayoutMethod`
* Add support for new value `apple_pay` on enum `V2MoneyManagementPayoutMethod.Type`
* Add support for `NetworkDetails` on `V2MoneyManagementReceivedCreditBankTransfer`
* Add support for new value `crypto_wallet_export` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* Change type of `V2MoneyManagementFinancialAccountListParams.Include` and `V2MoneyManagementFinancialAccountParams.Include` from `literal('payments.balance_by_funds_type')` to `enum('payments.balance_by_funds_type'|'storage.crypto')`
* Add support for `ForwardingSettings` on `V2MoneyManagementFinancialAccountParams`
* Add support for snapshot events `EventTypeAppsInstallCreated`, `EventTypeAppsInstallDeleted`, and `EventTypeAppsInstallUpdated` with resource `AppsInstall`
* Add support for event notifications `V2BillingContractCollectionBlockedEvent`, `V2BillingContractCollectionCurrentEvent`, `V2BillingContractCollectionPastDueEvent`, and `V2BillingContractCollectionUnpaidEvent` with related object `V2BillingContract`
* Add support for event notifications `V2CoreVaultNetworkTokenActivatedEvent`, `V2CoreVaultNetworkTokenAuthorizationRequirementsChangedEvent`, `V2CoreVaultNetworkTokenDeactivatedEvent`, `V2CoreVaultNetworkTokenDetailsUpdatedEvent`, and `V2CoreVaultNetworkTokenSuspendedEvent` with related object `V2CoreVaultNetworkToken`
* Add support for event notifications `V2MoneyManagementFinancialAccountWalletExportCompletedEvent`, `V2MoneyManagementFinancialAccountWalletExportPendingEvent`, and `V2MoneyManagementFinancialAccountWalletExportReadyEvent` with related object `V2MoneyManagementFinancialAccount`
* Add support for error type `ServiceUnavailableError`
