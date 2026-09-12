---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2404
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.4.0-alpha.1
---

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
