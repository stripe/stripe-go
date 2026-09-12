---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2347
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.2.0-beta.2
---

* Add support for new resources `V2CommerceProductCatalogImport`, `V2DataReportingQueryRun`, `V2ExtendWorkflowRun`, `V2ExtendWorkflow`, `V2IamActivityLog`, `V2NetworkBusinessProfile`, and `V2OrchestratedCommerceAgreement`
* Add support for `Confirm`, `Get`, `List`, `New`, and `Terminate` methods on resource `V2OrchestratedCommerceAgreement`
* Add support for `Get` and `Me` methods on resource `V2NetworkBusinessProfile`
* Add support for `List` method on resource `V2IamActivityLog`
* Add support for `Get` and `List` methods on resource `V2ExtendWorkflowRun`
* Add support for `Get`, `Invoke`, and `List` methods on resource `V2ExtendWorkflow`
* Add support for `Get` and `New` methods on resources `V2CommerceProductCatalogImport` and `V2DataReportingQueryRun`
* ⚠️ Change type of `V2BillingCadenceSettingsDataCollectionPaymentMethodOptions.Konbini`, `V2BillingCollectionSettingPaymentMethodOptions.Konbini`, `V2BillingCollectionSettingPaymentMethodOptionsParams.Konbini`, and `V2BillingCollectionSettingVersionPaymentMethodOptions.Konbini` from `map(string: dynamic)` to `an object`
* ⚠️ Change type of `V2BillingCadenceSettingsDataCollectionPaymentMethodOptions.SEPADebit`, `V2BillingCollectionSettingPaymentMethodOptions.SEPADebit`, `V2BillingCollectionSettingPaymentMethodOptionsParams.SEPADebit`, and `V2BillingCollectionSettingVersionPaymentMethodOptions.SEPADebit` from `map(string: dynamic)` to `an object`
* Add support for new values `cn_bank_account` and `jp_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* Add support for new values `futsu` and `toza` on enums `V2CoreVaultGbBankAccount.BankAccountType` and `V2MoneyManagementPayoutMethodBankAccount.BankAccountType`
* ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitProcessing` from `map(string: dynamic)` to `an object`
* ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitQueued` from `map(string: dynamic)` to `an object`
* ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitSucceeded` from `map(string: dynamic)` to `an object`
* Add support for new value `payout_method_amount_limit_exceeded` on enum `V2MoneyManagementOutboundTransferStatusDetailsFailed.Reason`
* ⚠️ Add support for new values `inbound_transfer_reversal`, `outbound_payment_reversal`, `outbound_transfer_reversal`, `received_credit_reversal`, `received_debit_reversal`, and `stripe_fee_tax` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* ⚠️ Remove support for value `return` from enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* Change type of `V2CoreBatchJobEndpointParams.HTTPMethod` from `literal('post')` to `enum('delete'|'post')`
* Add support for new value `meter_event_value_too_many_digits` on enums `EventsV1BillingMeterErrorReportTriggeredEventReasonErrorType.Code` and `EventsV1BillingMeterNoMeterFoundEventReasonErrorType.Code`
* Add support for `TreasuryTransaction` on `EventsV2MoneyManagementTransactionCreatedEvent`
* Add support for event notifications `V2CommerceProductCatalogImportsFailedEvent`, `V2CommerceProductCatalogImportsProcessingEvent`, `V2CommerceProductCatalogImportsSucceededEvent`, and `V2CommerceProductCatalogImportsSucceededWithErrorsEvent` with related object `V2CommerceProductCatalogImport`
* Add support for event notifications `V2DataReportingQueryRunCreatedEvent`, `V2DataReportingQueryRunFailedEvent`, `V2DataReportingQueryRunSucceededEvent`, and `V2DataReportingQueryRunUpdatedEvent` with related object `V2DataReportingQueryRun`
* Add support for event notifications `V2ExtendWorkflowRunFailedEvent`, `V2ExtendWorkflowRunStartedEvent`, and `V2ExtendWorkflowRunSucceededEvent` with related object `V2ExtendWorkflowRun`
* Add support for event notifications `V2OrchestratedCommerceAgreementConfirmedEvent`, `V2OrchestratedCommerceAgreementCreatedEvent`, `V2OrchestratedCommerceAgreementPartiallyConfirmedEvent`, and `V2OrchestratedCommerceAgreementTerminatedEvent` with related object `V2OrchestratedCommerceAgreement`
* Add support for error type `CannotProceedError`
