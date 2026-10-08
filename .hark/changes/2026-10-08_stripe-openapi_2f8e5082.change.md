---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2462
semver_level: major
is_stripe_api_change: true
---

* Add support for new resources `RadarRule`, `V2MoneyManagementFundingSession`, and `V2MoneyManagementInboundTransferMandate`
* Add support for `Cancel`, `Get`, `List`, and `New` methods on resource `V2MoneyManagementInboundTransferMandate`
* Add support for `New` method on resource `V2MoneyManagementFundingSession`
* Add support for `ExcludedPayoutDestinations` on `AccountSettingsCapitalParams`
* Add support for `WeroPayments` on `AccountCapabilities`
* ⚠️ Change type of `ChargeOutcome.Rule` from `RadarRule` to `$Radar.Rule`
* Add support for new value `ousd` on enums `ChargePaymentMethodDetailsCrypto.TokenCurrency`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.TokenCurrency`, and `PaymentRecordPaymentMethodDetailsCrypto.TokenCurrency`
* Add support for `Location` and `Reader` on `ChargePaymentMethodDetailsSwish`, `PaymentAttemptRecordPaymentMethodDetailsSwish`, and `PaymentRecordPaymentMethodDetailsSwish`
* Add support for `PaymentSettings` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `OnBehalfOf` on `CheckoutSession`
* Add support for new values `fednow` and `rtp` on enum `CustomerCashBalanceTransactionFundedBankTransferUsBankTransfer.Network`
* Add support for `FlexibleCredential` on `IssuingAuthorization`
* Add support for `Fuels` on `IssuingTransactionPurchaseDetails`
* Add support for `USBankAccount` on `PaymentAttemptRecordReportFailedPaymentMethodDetailsParams`, `PaymentRecordReportPaymentAttemptFailedPaymentMethodDetailsParams`, `PaymentRecordReportPaymentAttemptPaymentMethodDetailsParams`, `PaymentRecordReportPaymentPaymentMethodDetailsParams`, `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsParams`, and `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetails`
* Change type of `PaymentAttemptRecordReportFailedPaymentMethodDetailsParams.Type` and `PaymentRecordReportPaymentAttemptFailedPaymentMethodDetailsParams.Type` from `literal('card')` to `enum('card'|'us_bank_account')`
* Add support for `Fleet` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, and `PaymentIntentPaymentMethodOptionsCardPresent`
* Add support for `SubscriptionReference` on `PaymentIntentConfirmPaymentMethodOptionsPaypayParams`, `PaymentIntentPaymentMethodOptionsPaypayParams`, and `PaymentIntentPaymentMethodOptionsPaypay`
* Add support for `EnablementDetails` on `QuotePreviewSubscriptionScheduleDefaultSettingsAutomaticTax`, `QuotePreviewSubscriptionSchedulePhaseAutomaticTax`, and `SubscriptionAutomaticTax`
* Change type of `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsParams.MoneyMovementType` from `literal('card')` to `enum('card'|'us_bank_account')`
* Add support for `Rules` on `RadarPaymentEvaluation`
* ⚠️ Change type of `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetails.MoneyMovementType` from `literal('card')` to `enum('card'|'us_bank_account')`
* Add support for new values `request_three_d_secure` and `reroute` on enum `RadarPaymentEvaluation.RecommendedAction`
* Add support for `BankInitiatedReturn` on `RadarPaymentEvaluationSignals`
* Add support for new value `hour` on enums `SharedPaymentGrantedTokenUsageLimitsRecurring.Interval` and `SharedPaymentIssuedTokenUsageLimitsRecurring.Interval`
* Add support for `UtilityUsersTax` on `TaxRegistrationCountryOptionsUs`
* Add support for new values `digital_excise_tax` and `utility_users_tax` on enum `TaxRegistrationCountryOptionsUs.Type`
* Add support for `EnableCustomerCancellation` on `TerminalReaderActivateGiftCardParams`, `TerminalReaderCashoutGiftCardParams`, `TerminalReaderCheckGiftCardBalanceParams`, and `TerminalReaderReloadGiftCardParams`
* Add support for `VippsPayments` on `V2CoreAccountConfigurationMerchantCapabilitiesParams` and `V2CoreAccountConfigurationMerchantCapabilities`
* Add support for `BusinessCustodialStorage` on `V2CoreAccountConfigurationMoneyManagerCapabilitiesParams` and `V2CoreAccountConfigurationMoneyManagerCapabilities`
* Add support for `Offramp` and `Onramp` on `V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPaymentsParams`, `V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundPayments`, `V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfersParams`, `V2CoreAccountConfigurationMoneyManagerCapabilitiesOutboundTransfers`, `V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCreditsParams`, and `V2CoreAccountConfigurationMoneyManagerCapabilitiesReceivedCredits`
* Add support for `Pix` on `V2CoreAccountConfigurationRecipientCapabilitiesParams`, `V2CoreAccountConfigurationRecipientCapabilities`, `V2MoneyManagementOutboundSetupIntentPayoutMethodDataParams`, and `V2MoneyManagementPayoutMethod`
* Add support for new value `pix` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* Add support for new values `pix` and `vipps_payments` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for `Account` on `V2MoneyManagementFinancialAddressListParams`, `V2MoneyManagementFinancialAddressParams`, and `V2MoneyManagementFinancialAddress`
* Add support for new values `bre_b`, `nip`, and `pix` on enum `V2MoneyManagementFinancialAddressBankAccount.Type`
* Add support for `SupportedNetworkDetails` on `V2MoneyManagementFinancialAddressCryptoWallet`
* Add support for new value `bitcoin` on enums `V2MoneyManagementFinancialAddressCryptoWallet.Network` and `V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWallet.Network`
* Add support for `NetworkDetails` on `V2MoneyManagementInboundTransferParams` and `V2MoneyManagementInboundTransfer`
* Add support for `BACSDebit` on `V2MoneyManagementInboundTransferFromPaymentMethod`
* Add support for new value `pix` on enum `V2MoneyManagementPayoutMethod.Type`
* Add support for `BIC` on `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountAba` and `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccountSortCode`
* Add support for `OriginatingCryptoWallet`, `TokenCurrency`, and `TransactionHash` on `V2MoneyManagementReceivedCreditCryptoWalletTransfer`
* Add support for `Invoices` on `V2TaxIntegrationConfigurationParams` and `V2TaxIntegrationConfiguration`
* ⚠️ Remove support for `Account` on `V2RiskInquiryListParams`
* Add support for `Customer` and `Subscription` on `EventsV1InvoiceUpcomingEvent`
* Add support for new value `vipps_payments` on enum `EventsV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for new values `business_custodial_storage.inbound.ousd`, `business_custodial_storage.inbound.usdc`, `business_custodial_storage.outbound.ousd`, `business_custodial_storage.outbound.usdc`, `outbound_payments.offramp.bank_accounts.brl`, `outbound_payments.offramp.bank_accounts.cop`, `outbound_payments.offramp.bank_accounts.eur`, `outbound_payments.offramp.bank_accounts.gbp`, `outbound_payments.offramp.bank_accounts.mxn`, `outbound_payments.offramp.bank_accounts.usd`, `outbound_payments.onramp.crypto_wallets.brl`, `outbound_payments.onramp.crypto_wallets.cop`, `outbound_payments.onramp.crypto_wallets.eur`, `outbound_payments.onramp.crypto_wallets.gbp`, `outbound_payments.onramp.crypto_wallets.mxn`, `outbound_payments.onramp.crypto_wallets.usd`, `outbound_transfers.offramp.bank_accounts.brl`, `outbound_transfers.offramp.bank_accounts.cop`, `outbound_transfers.offramp.bank_accounts.eur`, `outbound_transfers.offramp.bank_accounts.gbp`, `outbound_transfers.offramp.bank_accounts.mxn`, `outbound_transfers.offramp.bank_accounts.usd`, `outbound_transfers.onramp.crypto_wallets.brl`, `outbound_transfers.onramp.crypto_wallets.cop`, `outbound_transfers.onramp.crypto_wallets.eur`, `outbound_transfers.onramp.crypto_wallets.gbp`, `outbound_transfers.onramp.crypto_wallets.mxn`, `outbound_transfers.onramp.crypto_wallets.usd`, `received_credits.offramp.bank_accounts.brl`, `received_credits.offramp.bank_accounts.cop`, `received_credits.offramp.bank_accounts.eur`, `received_credits.offramp.bank_accounts.gbp`, `received_credits.offramp.bank_accounts.mxn`, `received_credits.offramp.bank_accounts.usd`, `received_credits.onramp.crypto_wallets.brl`, `received_credits.onramp.crypto_wallets.cop`, `received_credits.onramp.crypto_wallets.eur`, `received_credits.onramp.crypto_wallets.gbp`, `received_credits.onramp.crypto_wallets.mxn`, and `received_credits.onramp.crypto_wallets.usd` on enum `EventsV2CoreAccountIncludingConfigurationMoneyManagerCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for new value `pix` on enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for event notifications `V2MoneyManagementInboundTransferMandateActivatedEvent`, `V2MoneyManagementInboundTransferMandateCreatedEvent`, `V2MoneyManagementInboundTransferMandateExpiredEvent`, `V2MoneyManagementInboundTransferMandateRefusedEvent`, and `V2MoneyManagementInboundTransferMandateRevokedEvent` with related object `V2MoneyManagementInboundTransferMandate`
* Add support for error code `service_unavailable` on `ServiceUnavailableError`
