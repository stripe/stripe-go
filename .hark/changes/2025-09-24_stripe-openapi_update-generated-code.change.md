---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2123
is_stripe_api_change: true
released_in_version: 83.1.0-beta.1
---

* Add support for new resources `BillingAnalyticsMeterUsageRow`, `BillingAnalyticsMeterUsage`, `V2BillingBillSettingVersion`, `V2BillingBillSetting`, `V2BillingCadence`, `V2BillingCollectionSettingVersion`, `V2BillingCollectionSetting`, and `V2BillingProfile`
* Remove support for resources `BillingMeterUsageRow` and `BillingMeterUsage`
* Add support for `Get` method on resource `BillingAnalyticsMeterUsage`
* Remove support for `Get` method on resource `BillingMeterUsage`
* Add support for `ReportPaymentAttemptInformational` method on resource `PaymentRecord`
* Add support for `Get`, `List`, `New`, and `Update` methods on resources `V2BillingBillSetting`, `V2BillingCollectionSetting`, and `V2BillingProfile`
* Add support for `Get` and `List` methods on resources `V2BillingBillSettingVersion` and `V2BillingCollectionSettingVersion`
* Add support for `Cancel`, `Get`, `List`, `New`, and `Update` methods on resource `V2BillingCadence`
* Add support for new value `crypto_wallet` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* Add support for `Profile` on `V2CoreAccountDefaultsParams` and `V2CoreAccountDefaults`
* Add support for `IP` on `V2CoreAccountIdentityAttestationsDirectorshipDeclaration`, `V2CoreAccountIdentityAttestationsOwnershipDeclaration`, `V2CoreAccountIdentityAttestationsTermsOfServiceAccountParams`, `V2CoreAccountIdentityAttestationsTermsOfServiceAccount`, `V2CoreAccountIdentityAttestationsTermsOfServiceStorerParams`, `V2CoreAccountIdentityAttestationsTermsOfServiceStorer`, `V2CoreAccountIdentityIndividualAdditionalTermsOfServiceAccount`, `V2CorePersonAdditionalTermsOfServiceAccountParams`, and `V2CorePersonAdditionalTermsOfServiceAccount`
* Remove support for `IP` on `V2CoreAccountIdentityAttestationsDirectorshipDeclaration`, `V2CoreAccountIdentityAttestationsOwnershipDeclaration`, `V2CoreAccountIdentityAttestationsTermsOfServiceAccountParams`, `V2CoreAccountIdentityAttestationsTermsOfServiceAccount`, `V2CoreAccountIdentityAttestationsTermsOfServiceStorerParams`, `V2CoreAccountIdentityAttestationsTermsOfServiceStorer`, `V2CoreAccountIdentityIndividualAdditionalTermsOfServiceAccount`, `V2CorePersonAdditionalTermsOfServiceAccountParams`, and `V2CorePersonAdditionalTermsOfServiceAccount`
* Remove support for `DoingBusinessAs`, `ProductDescription`, and `URL` on `V2CoreAccountIdentityBusinessDetailsParams` and `V2CoreAccountIdentityBusinessDetails`
* Add support for `SettlementCurrency` on `V2MoneyManagementFinancialAddress`
* Add support for `SEPABankAccount` on `V2MoneyManagementFinancialAddressCredentials` and `V2MoneyManagementReceivedCreditBankTransfer`
* Add support for new value `sepa_bank_account` on enum `V2MoneyManagementFinancialAddressCredentials.Type`
* Add support for `AmountDetails` and `PaymentsOrchestration` on `V2PaymentsOffSessionPaymentParams` and `V2PaymentsOffSessionPayment`
* Add support for new value `authorization_expired` on enum `V2PaymentsOffSessionPayment.FailureReason`
* Add support for `RetryPolicy` on `V2PaymentsOffSessionPaymentRetryDetailsParams` and `V2PaymentsOffSessionPaymentRetryDetails`
* Add support for new values `heuristic` and `scheduled` on enum `V2PaymentsOffSessionPaymentRetryDetails.RetryStrategy`
* Change type of `V2MoneyManagementOutboundPaymentQuoteFxQuote.LockDuration` from `literal('five_minutes')` to `enum('five_minutes'|'none')`
* Add support for new value `none` on enum `V2MoneyManagementOutboundPaymentQuoteFxQuote.LockStatus`
* Add support for new value `crypto_wallet` on enum `V2MoneyManagementPayoutMethod.Type`
* Add support for `OriginType` on `V2MoneyManagementReceivedCreditBankTransfer`
* Remove support for `PaymentMethodType` on `V2MoneyManagementReceivedCreditBankTransfer`
* Add support for `MinimumBalanceByCurrency` on `BalanceSettingsPaymentsPayoutsParams` and `BalanceSettingsPaymentsPayouts`
* Change type of `BalanceSettingsPaymentsSettlementTimingParams.DelayDaysOverride` from `longInteger` to `emptyable(longInteger)`
* Remove support for values `saturday` and `sunday` from enum `BalanceSettingsPaymentsPayoutsSchedule.WeeklyPayoutDays`
* Add support for `DelayDaysOverride` on `BalanceSettingsPaymentsSettlementTiming`
* Add support for `AutomaticTax` and `InvoiceCreation` on `CheckoutSessionParams`
* Add support for `UnitLabel` on `CheckoutSessionLineItemPriceDataProductDataParams`
* Add support for `InvoiceSettings` on `CheckoutSessionSubscriptionDataParams`
* Add support for `IntendedSubmissionMethod` on `DisputeParams` and `Dispute`
* Change type of `DisputeSmartDisputes.RecommendedEvidence` from `string` to `array(string)`
* Add support for new value `prevented` on enum `Dispute.Status`
* Add support for `Pix` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for `BillingSchedules` on `InvoiceCreatePreviewSubscriptionDetailsParams`, `SubscriptionParams`, and `Subscription`
* Add support for new value `pix` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `Paypay` on `PaymentAttemptRecordPaymentMethodDetails` and `PaymentRecordPaymentMethodDetails`
* Add support for `Wallet` on `PaymentAttemptRecordPaymentMethodDetailsCard` and `PaymentRecordPaymentMethodDetailsCard`
* Add support for `Flexible` on `QuotePreviewSubscriptionScheduleBillingMode`
* Add support for `BilledUntil` on `SubscriptionItem`
* Add support for `MandateData` and `PaymentMethodOptions` on `V2PaymentsOffSessionPaymentParams`
* Add support for `Type` on `V2MoneyManagementFinancialAddressParams`
* Remove support for `Currency` on `V2MoneyManagementFinancialAddressParams`
* Add support for new values `financial_addressses.crypto_wallets`, `holds_currencies.usdc`, `outbound_payments.crypto_wallets`, and `outbound_transfers.crypto_wallets` on enum `EventsV2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for thin event `V2BillingBillSettingUpdatedEvent` with related object `V2BillingBillSetting`
* Add support for error type `RateLimitError`
* Add support for error codes `financial_connections_account_pending_account_numbers` and `financial_connections_account_unavailable_account_numbers` on `QuotePreviewInvoiceLastFinalizationError`
* Add support for error code `invalid_payout_method_crypto_wallet` on `InvalidPayoutMethodError`
