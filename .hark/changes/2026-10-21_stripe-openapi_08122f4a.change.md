---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2475
semver_level: major
is_stripe_api_change: true
released_in_version: 87.1.0-alpha.4
---

* Add support for new resources `RiskForm` and `V2DataSearchResultList`
* Add support for `Get` and `Update` methods on resource `RiskForm`
* Add support for `Search` method on resource `V2DataSearchResultList`
* Add support for `BusinessProfileManagement` and `ReportingChart` on `AccountSessionComponents`
* Change type of `AccountSessionComponents.TerminalHardwareOrders` from `nullable(ConnectEmbeddedBaseConfig)` to `ConnectEmbeddedBaseConfig`
* Change type of `AccountSessionComponents.TerminalHardwareShop` from `nullable(ConnectEmbeddedBaseConfig)` to `ConnectEmbeddedBaseConfig`
* Add support for `SubscriptionPause` on `BillingPortalConfigurationFeatures`
* Add support for new value `mandiri` on enums `ChargePaymentMethodDetailsIdBankTransfer.Bank`, `ConfirmationTokenPaymentMethodPreviewIdBankTransfer.Bank`, `PaymentAttemptRecordPaymentMethodDetailsIdBankTransfer.Bank`, `PaymentMethodIdBankTransfer.Bank`, `PaymentRecordPaymentMethodDetailsIdBankTransfer.Bank`, and `SetupAttemptPaymentMethodDetailsIdBankTransfer.Bank`
* Add support for `MarketingAttributions` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `MetronomeCommit` on `CheckoutSessionItemParams` and `CheckoutSessionItem`
* Add support for `AdjustableQuantity` on `CheckoutSessionItemSubscriptionItemParams`
* Add support for new values `au_acn`, `au_wpn`, `cz_dic`, `fr_siren`, `fr_siret`, `fr_spi`, `is_kt`, `md_cf`, `mu_vat`, `mz_nuit`, `nz_bn`, `pt_nif`, `sm_coe`, `tt_bir`, and `tt_vat` on enums `CheckoutSessionCollectedInformationTaxId.Type`, `CheckoutSessionCustomerDetailsTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, `QuotePreviewInvoiceCustomerTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
* Add support for new value `metronome_commit` on enum `CheckoutSessionItem.Type`
* Add support for `MandateOptions` on `CheckoutSessionPaymentMethodOptionsCard`
* ⚠️ Change type of `DisputeEvidence.Appeal` from `DisputeAppeal` to `nullable(DisputeAppeal)`
* ⚠️ Remove support for `Details` on `IdentityVerificationReportEmail`
* ⚠️ Remove support for value `email` from enums `IdentityVerificationReport.Type` and `IdentityVerificationSession.Type`
* Add support for `MerchantAmountExchangeRate` on `IssuingAuthorizationParams`, `TestHelpersIssuingAuthorizationCaptureParams`, and `TestHelpersIssuingTransactionCreateForceCaptureParams`
* Add support for `Type` on `IssuingAuthorizationParams` and `IssuingAuthorization`
* Add support for `SettlementDetails` on `TestHelpersIssuingAuthorizationCaptureParams` and `TestHelpersIssuingTransactionCreateForceCaptureParams`
* ⚠️ Remove support for `Healthcare` on `TestHelpersIssuingAuthorizationCapturePurchaseDetailsParams`, `TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsParams`, and `TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsParams`
* Add support for `ProductEligibilityInquiry` on `IssuingAuthorization`
* Add support for `AccountTaxIDs`, `CustomFields`, `Description`, and `Footer` on `PaymentLinkSubscriptionDataInvoiceSettings`
* Change type of `RadarPaymentEvaluationSignals.EarlyFraudWarning` from `nullable(Payment Evaluation Signal)` to `Payment Evaluation Signal`
* Add support for `ReferenceStatus` and `Reference` on `RefundDestinationDetailsAfterpayClearpay`, `RefundDestinationDetailsKlarna`, and `RefundDestinationDetailsPaypal`
* Add support for `Bw`, `Ci`, `Gh`, and `Mu` on `TaxRegistrationCountryOptions`
* Add support for new value `account_takeover` on enum `V2CoreAccountEvaluation.EvaluationsTriggered`
* Add support for `CZK`, `NOK`, `PLN`, and `SEK` on `V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInboundParams`, `V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageInbound`, `V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutboundParams`, and `V2CoreAccountConfigurationMoneyManagerCapabilitiesBusinessStorageOutbound`
* Add support for `StablecoinCustodial` on `V2CoreAccountConfigurationRecipientCapabilitiesStripeBalanceParams` and `V2CoreAccountConfigurationRecipientCapabilitiesStripeBalance`
* Add support for new value `stripe_balance.stablecoin_custodial` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for `AccountTakeoverConfirmation`, `AccountTakeoverReport`, `CredentialChange`, `MfaAddition`, and `MfaRemoval` on `V2SignalsAccountActivityParams`, `V2SignalsAccountActivity`, and `V2SignalsAccountEvaluationAccountActivityDetailsDataParams`
* Add support for new values `account_takeover_confirmation`, `account_takeover_mitigation`, `account_takeover_report`, `credential_change`, `mfa_addition`, and `mfa_removal` on enum `V2SignalsAccountActivity.Type`
* Add support for `AccountTakeover` on `V2SignalsAccountEvaluationEvaluatedSignals` and `V2SignalsAccountSignal`
* Add support for new value `account_takeover` on enum `V2SignalsAccountEvaluation.PendingSignals`
* Add support for new value `account_takeover` on enum `V2SignalsAccountEvaluation.RequestedSignals`
* Add support for `CustomerRisk` on `V2SignalsAccountSignal`
* Add support for new values `account_takeover` and `customer_risk` on enum `V2SignalsAccountSignal.Type`
* Add support for new values `business_storage.inbound.czk`, `business_storage.inbound.nok`, `business_storage.inbound.pln`, `business_storage.inbound.sek`, `business_storage.outbound.czk`, `business_storage.outbound.nok`, `business_storage.outbound.pln`, and `business_storage.outbound.sek` on enum `EventsV2CoreAccountIncludingConfigurationMoneyManagerCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for new value `stripe_balance.stablecoin_custodial` on enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for event notification `V2SignalsAccountSignalCustomerRiskReadyEvent` with related object `V2SignalsAccountSignal`
