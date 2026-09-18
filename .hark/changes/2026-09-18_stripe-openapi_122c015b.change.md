---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2439
semver_level: major
is_stripe_api_change: true
---

* Add support for new resources `V2ProvisioningEligibility`, `V2ProvisioningPaymentMethodRequest`, `V2ProvisioningPaymentProfile`, `V2ProvisioningProject`, `V2ProvisioningProviderConnectionRequest`, `V2ProvisioningProviderConnection`, `V2ProvisioningProviderServiceDetail`, `V2ProvisioningProvider`, and `V2ProvisioningResource`
* ⚠️ Remove support for resource `RadarBillingEvaluation`
* ⚠️ Remove support for `New` method on resource `RadarBillingEvaluation`
* Add support for `Get`, `Link`, `New`, `Remove`, `RotateCredentials`, `SubmitInformation`, `Unlink`, and `Update` methods on resource `V2ProvisioningResource`
* Add support for `Get`, `New`, and `SubmitInformation` methods on resource `V2ProvisioningProviderConnectionRequest`
* Add support for `List` and `Unlink` methods on resource `V2ProvisioningProviderConnection`
* Add support for `New` method on resources `V2ProvisioningPaymentMethodRequest` and `V2ProvisioningProject`
* Add support for `Get` and `Update` methods on resource `V2ProvisioningPaymentProfile`
* Add support for `Get` method on resource `V2ProvisioningEligibility`
* Add support for new values `invalid_address_cmra_address` and `invalid_address_registered_agent_address` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
* Add support for `BLIKRecurringPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `SequraPayments` on `AccountCapabilitiesParams`
* Add support for `Capital` on `AccountSettings`
* Add support for `PayoutMethod` on `BalanceInstantAvailableNetAvailable`
* Add support for `DestinationCurrency` on `BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrencyParams` and `BalanceSettingsPaymentsPayoutsAutomaticTransferRulesByCurrency`
* Add support for `TotalDueAmount` on `CapitalFinancingOfferAcceptedTerms` and `CapitalFinancingSummaryDetails`
* Add support for `IncrementalIntervalTargetAmount` and `StartsAt` on `CapitalFinancingSummaryDetailsCurrentRepaymentInterval`
* Add support for `SetupCredentialUsage` on `ChargePaymentMethodDetailsCard`, `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCard`, `SetupIntentConfirmPaymentMethodOptionsCardParams`, `SetupIntentPaymentMethodOptionsCardParams`, and `SetupIntentPaymentMethodOptionsCard`
* Add support for `StoredCredentialUsage` on `ChargePaymentMethodDetailsCard`, `PaymentAttemptRecordPaymentMethodDetailsCard`, `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCard`, and `PaymentRecordPaymentMethodDetailsCard`
* Add support for `PaymentMethodOptions` on `CheckoutSessionApproveParams`
* Add support for `PaymentReservation` on `CheckoutSession`
* Add support for `Custom` on `CheckoutSessionCurrentAttemptPaymentMethodDetails`
* Change type of `CheckoutSessionItem.Subscription` from `nullable(PaymentPagesCheckoutSessionSubscription)` to `PaymentPagesCheckoutSessionSubscription`
* Add support for `PaymentMethodPreselect` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`
* Add support for `BIC`, `IBANLast4`, and `Network` on `CustomerCashBalanceTransactionFundedBankTransferGbBankTransfer`
* Add support for `Appeal` on `DisputeEvidenceParams`
* Add support for new values `apps.install.created`, `apps.install.deleted`, and `apps.install.updated` on enum `Event.Type`
* Add support for new value `expired` on enum `FinancialConnectionsAccountAccountNumbers.Status`
* Add support for `PreCollectedConsent` on `FinancialConnectionsSession`
* Add support for `FinancialActivity` on `FinancialConnectionsTransactionClassifications`
* Add support for `InvoicingRules` on `InvoiceItemParams`
* Add support for new value `touch_n_go` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* ⚠️ Change type of `MandatePaymentMethodDetailsBlik.Type` from `enum('off_session'|'on_session')` to `literal('off_session')`
* Add support for `Sequra` on `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentMethodConfigurationParams`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `MandateOptions` on `PaymentIntentConfirmPaymentMethodOptionsBlikParams`, `PaymentIntentPaymentMethodOptionsBlikParams`, and `PaymentIntentPaymentMethodOptionsBlik`
* Change type of `PaymentIntentConfirmPaymentMethodOptionsBlikParams.SetupFutureUsage` and `PaymentIntentPaymentMethodOptionsBlikParams.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
* ⚠️ Remove support for `CaptureMethod` on `PaymentIntentConfirmPaymentMethodOptionsPaypayParams` and `PaymentIntentPaymentMethodOptionsPaypayParams`
* ⚠️ Change type of `PaymentIntentPaymentMethodOptionsBlik.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session')`
* ⚠️ Remove support for `Payto` on `PaymentMethodParams`
* Add support for `PayoutMethodOptions` on `Payout`
* Add support for new value `rerouted` on enum `RadarPaymentEvaluationOutcome.Type`
* Add support for `BLIK` on `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for `ExpiresAt` on `SubscriptionPaymentSettingsPaymentMethodOptionsBlikMandateOptionsParams`
* ⚠️ Remove support for `ExpiresAfter` on `SubscriptionPaymentSettingsPaymentMethodOptionsBlikMandateOptionsParams`
* Add support for `CancelAtPeriodEnd` on `SubscriptionPendingUpdate`
* Add support for `TamperState` on `TerminalReaderListParams`
* Add support for new value `developer` on enums `EventsV2CoreAccountLinkReturnedEvent.Configurations`, `V2CoreAccountLinkUseCaseAccountOnboarding.Configurations`, and `V2CoreAccountLinkUseCaseAccountUpdate.Configurations`
* Add support for new value `developer` on enum `V2CoreAccount.AppliedConfigurations`
* Add support for `Developer` on `V2CoreAccountConfigurationParams` and `V2CoreAccountConfiguration`
* Add support for new value `projects` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new value `developer` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Configuration` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Configuration`
* Add support for snapshot events `EventTypeAppsInstallCreated`, `EventTypeAppsInstallDeleted`, and `EventTypeAppsInstallUpdated` with resource `AppsInstall`
