---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2378
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.2.0-alpha.1
---

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
