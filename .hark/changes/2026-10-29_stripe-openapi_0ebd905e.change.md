---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2463
semver_level: major
is_stripe_api_change: true
released_in_version: 87.2.0-beta.1
---

* Add support for new resources `PaymentPlan`, `V2DataQueryRun`, `V2DataReportRun`, `V2DataReport`, and `V2DataSchema`
* Add support for `Get`, `List`, `New`, and `Update` methods on resource `PaymentPlan`
* Add support for `Get` and `List` methods on resources `V2DataReport` and `V2DataSchema`
* Add support for `Get` and `New` methods on resources `V2DataQueryRun` and `V2DataReportRun`
* ⚠️ Remove support for `AcknowledgeConfirmationOfPayee` and `InitiateConfirmationOfPayee` methods on resource `V2CoreVaultGBBankAccount`
* Add support for new value `issuing.application_rejected_because_prohibited_business` on enum `AccountNotice.Reason`
* Add support for `SettlementCurrencies` on `BalanceSettingsPaymentsParams` and `BalanceSettingsPayments`
* Add support for `DefaultSettlementCurrency` on `BalanceSettingsPayments`
* Add support for `Wero` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new value `mandiri` on enums `ChargePaymentMethodDetailsIdBankTransfer.Bank`, `ConfirmationTokenPaymentMethodPreviewIdBankTransfer.Bank`, `PaymentAttemptRecordPaymentMethodDetailsIdBankTransfer.Bank`, `PaymentMethodIdBankTransfer.Bank`, `PaymentRecordPaymentMethodDetailsIdBankTransfer.Bank`, and `SetupAttemptPaymentMethodDetailsIdBankTransfer.Bank`
* Add support for new values `au_acn`, `au_wpn`, `bw_vat`, `ci_vat`, `cz_dic`, `fr_siren`, `fr_siret`, `fr_spi`, `gh_tin`, `is_kt`, `md_cf`, `mu_vat`, `mz_nuit`, `nz_bn`, `pt_nif`, `sm_coe`, `tt_bir`, and `tt_vat` on enums `CheckoutSessionCollectedInformationTaxId.Type`, `OrderTaxDetailsTaxId.Type`, and `QuotePreviewInvoiceCustomerTaxIds.Type`
* Add support for `ExternalToken` on `ConfirmationTokenPaymentMethodDataUsBankAccountParams`, `PaymentIntentConfirmPaymentMethodDataUsBankAccountParams`, `PaymentIntentPaymentMethodDataUsBankAccountParams`, `PaymentMethodUsBankAccountParams`, `SetupIntentConfirmPaymentMethodDataUsBankAccountParams`, `SetupIntentPaymentMethodDataUsBankAccountParams`, and `TreasuryOutboundPaymentDestinationPaymentMethodDataUsBankAccountParams`
* Add support for new value `wero` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* ⚠️ Change type of `DisputeEvidence.Appeal` from `DisputeAppeal` to `nullable(DisputeAppeal)`
* Add support for `PaymentPlan` on `Invoice`
* Add support for new value `wero` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for new value `sequra` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
* ⚠️ Remove support for `DefaultOutboundDestination` on `V2CoreAccountConfigurationRecipientParams` and `V2CoreAccountConfigurationRecipient`
* Add support for `CryptoWallets` on `V2CoreAccountConfigurationRecipientCapabilitiesParams` and `V2CoreAccountConfigurationRecipientCapabilities`
* Add support for `ACH`, `FPS`, `Fedwire`, `RTP`, `SEPA`, and `Swift` on `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsParams` and `V2CoreAccountConfigurationRecipientCapabilitiesBankAccounts`
* ⚠️ Remove support for `Wire` on `V2CoreAccountConfigurationRecipientCapabilitiesBankAccountsParams` and `V2CoreAccountConfigurationRecipientCapabilitiesBankAccounts`
* Add support for new value `physical` on enum `V2CoreAccountIdentityBusinessDetailsAdditionalAddress.Purpose`
* Add support for `ProofOfBusinessPresence` on `V2CoreAccountIdentityBusinessDetailsDocumentsParams`, `V2CoreAccountIdentityBusinessDetailsDocuments`, and `V2CoreAccountTokenIdentityBusinessDetailsDocumentsParams`
* Add support for `PartialOutputFile` on `V2CoreBatchJobStatusDetailsBatchFailed`
* Add support for `Fingerprint` on `V2CoreVaultGbBankAccount`, `V2CoreVaultUsBankAccount`, and `V2MoneyManagementPayoutMethodBankAccount`
* ⚠️ Remove support for values `futsu` and `toza` from enum `V2CoreVaultGbBankAccount.BankAccountType`
* Add support for `FedwireFingerprint` on `V2CoreVaultUsBankAccount`
* ⚠️ Remove support for values `balance_exchange` and `inbound_payment` from enum `V2MoneyManagementAdjustmentAdjustedFlow.Type`
* Add support for `Address` on `V2MoneyManagementFinancialAddressBankAccountAba`
* ⚠️ Remove support for `AccountHolderAddress` on `V2MoneyManagementFinancialAddressBankAccountAba`
* ⚠️ Remove support for `Last4` on `V2MoneyManagementFinancialAddressBankAccountAba`, `V2MoneyManagementFinancialAddressBankAccountIban`, and `V2MoneyManagementFinancialAddressBankAccountSortCode`
* ⚠️ Remove support for `Country` on `V2MoneyManagementFinancialAddressBankAccountIban`
* Add support for `TaxAmount` on `V2MoneyManagementOutboundPaymentQuoteEstimatedFee`
* ⚠️ Remove support for values `becs`, `eft`, `npp`, and `sepa_instant` from enum `V2MoneyManagementOutboundPaymentQuoteEstimatedFeeNetworkFeeDetails.Network`
* Add support for new value `next_day_payout_fee` on enum `V2MoneyManagementOutboundPaymentQuoteEstimatedFee.Type`
* Add support for `PayoutMethodOptions` on `V2MoneyManagementOutboundPaymentQuoteToParams`, `V2MoneyManagementOutboundPaymentQuoteTo`, `V2MoneyManagementOutboundPaymentToParams`, `V2MoneyManagementOutboundPaymentTo`, `V2MoneyManagementOutboundTransferToParams`, and `V2MoneyManagementOutboundTransferTo`
* Add support for `PayoutIntent` and `TrackingDetails` on `V2MoneyManagementOutboundPayment`
* Add support for `RecipientVerification` on `V2MoneyManagementOutboundPaymentParams` and `V2MoneyManagementOutboundPayment`
* Add support for `CryptoWallet` on `V2MoneyManagementOutboundSetupIntentPayoutMethodDataParams` and `V2MoneyManagementPayoutMethod`
* Add support for `Network` on `V2MoneyManagementReceivedCreditBankTransfer`
* ⚠️ Remove support for `GBBankAccount`, `SEPABankAccount`, and `USBankAccount` on `V2MoneyManagementReceivedCreditBankTransfer`
* ⚠️ Remove support for `Network` on `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountAba`, `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountIban`, and `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCode`
* Add support for `FeeBatch` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
* ⚠️ Remove support for `FeeTransaction` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
* ⚠️ Add support for new value `fee_batch` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
* ⚠️ Remove support for value `fee_transaction` from enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
* Add support for `Metadata` and `RegulatoryReceipt` on `V2MoneyManagementTransaction`
* Add support for `ContactEmail` on `V2SignalsAccountActivityAccountDetailsDataParams`, `V2SignalsAccountActivityAccountDetailsData`, `V2SignalsAccountEvaluationAccountDetailsDataParams`, and `V2SignalsAccountEvaluationAccountDetailsData`
* Add support for `StatementDescriptor` on `V2MoneyManagementOutboundPaymentParams` and `V2MoneyManagementOutboundTransferParams`
* ⚠️ Remove support for `CreatedGTE`, `CreatedGt`, `CreatedLT`, and `CreatedLte` on `V2MoneyManagementOutboundPaymentListParams`
* ⚠️ Change type of `V2MoneyManagementOutboundPaymentListParams.Created` from `DateTime` to `an object`
* ⚠️ Remove support for `SettlementCurrency` on `V2MoneyManagementFinancialAddressParams`
* ⚠️ Remove support for `DepositInsuranceEligibility` on `V2MoneyManagementFinancialAccountStorageParams`
* Add support for snapshot events `EventTypePaymentPlanCreated`, `EventTypePaymentPlanInstallmentDue`, `EventTypePaymentPlanInstallmentPaid`, `EventTypePaymentPlanInstallmentWillBeDue`, and `EventTypePaymentPlanUpdated` with resource `PaymentPlan`
* Add support for event notifications `V2DataQueryRunCreatedEvent`, `V2DataQueryRunFailedEvent`, `V2DataQueryRunSucceededEvent`, and `V2DataQueryRunUpdatedEvent` with related object `V2DataQueryRun`
* Add support for event notifications `V2DataReportRunCreatedEvent`, `V2DataReportRunFailedEvent`, `V2DataReportRunSucceededEvent`, and `V2DataReportRunUpdatedEvent` with related object `V2DataReportRun`
* Add support for error codes `smart_disputes_not_available`, `subscription_update_trial_after_billing_cycle_anchor`, `three_d_secure_authentication_not_supported`, and `three_d_secure_authentication_upstream_error` on `QuotePreviewInvoiceLastFinalizationError`
* ⚠️ Remove support for error code `blocked_payout_method_card` on `BlockedByStripeError`
* Add support for error code `us_bank_account_microdeposits_requires_routing_number` on `CannotProceedError`
* Add support for error codes `gb_bank_account_cannot_be_unarchived` and `us_bank_account_cannot_be_unarchived` on `ControlledByAlternateResourceError`
* ⚠️ Remove support for error code `outbound_flow_from_closed_financial_account_unsupported` on `FeatureNotEnabledError`
* ⚠️ Remove support for error code `invalid_payout_method_card` on `InvalidPayoutMethodError`
* Add support for error code `payout_intent_not_cancelable` on `NotCancelableError`
* ⚠️ Remove support for error code `archived_payout_method_card` on `QuotaExceededError`
