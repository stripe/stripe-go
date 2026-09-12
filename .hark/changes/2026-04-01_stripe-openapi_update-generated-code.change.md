---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2335
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.1.0-alpha.2
---

* Add support for new resources `SharedPaymentIssuedToken` and `V2DataReportingQueryRun`
* Add support for `Get` and `New` methods on resource `V2DataReportingQueryRun`
* Add support for `Pause` and `Resume` methods on resource `V2PaymentsOffSessionPayment`
* Add support for `TenantKeys`, `TenantOperator`, and `TenantValues` on `BillingBillingMeterMeterEventSummaryListParams`
* Add support for `MoneyServices` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmPaymentDetailsParams`, and `PaymentIntentPaymentDetailsParams`
* Add support for `PaymentMethodOptions` on `DelegatedCheckoutRequestedSessionParams` and `DelegatedCheckoutRequestedSession`
* ⚠️ Remove support for `PaymentMethodData` on `DelegatedCheckoutRequestedSessionConfirmParams` and `DelegatedCheckoutRequestedSessionParams`
* Add support for `CardBrands` and `PaymentMethodTypes` on `DelegatedCheckoutRequestedSessionSellerDetails`
* ⚠️ Change type of `DelegatedCheckoutRequestedSession.SharedPaymentIssuedToken` from `string` to `expandable($SharedPayment.IssuedToken)`
* Add support for `CheckScan` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for new value `check_scan` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `ProcessorDetails` on `PaymentAttemptRecordReportFailedParams`, `PaymentAttemptRecordReportGuaranteedParams`, `PaymentRecordReportPaymentAttemptFailedParams`, `PaymentRecordReportPaymentAttemptGuaranteedParams`, `PaymentRecordReportPaymentFailedParams`, and `PaymentRecordReportPaymentGuaranteedParams`
* Add support for `PaymentDetails` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardParams`, and `PaymentIntentPaymentMethodOptionsCardPresentParams`
* ⚠️ Remove support for `BillFrom` on `QuotePreviewSubscriptionScheduleBillingSchedule`, `SubscriptionBillingSchedule`, and `SubscriptionScheduleBillingSchedule`
* Add support for `AgentDetails`, `PaymentMethodDetails`, and `RiskDetails` on `SharedPaymentGrantedToken`
* Add support for `PaperChecks` on `V2AccountConfigurationRecipientDataFeaturesParams`, `V2AccountConfigurationRecipientDataFeatures`, `V2CoreAccountConfigurationRecipientCapabilitiesParams`, `V2CoreAccountConfigurationRecipientCapabilities`, `V2CoreAccountConfigurationStorerCapabilitiesOutboundPaymentsParams`, and `V2CoreAccountConfigurationStorerCapabilitiesOutboundPayments`
* Add support for new value `paper_checks` on enum `V2AccountConfigurationSupportableFeatures.RecipientData`
* Add support for new value `paper_checks` on enum `V2AccountRequirementImpact.RequiredForFeatures`
* ⚠️ Change type of `V2BillingCadenceSettingsDataCollectionPaymentMethodOptions.Konbini`, `V2BillingCollectionSettingPaymentMethodOptions.Konbini`, `V2BillingCollectionSettingPaymentMethodOptionsParams.Konbini`, and `V2BillingCollectionSettingVersionPaymentMethodOptions.Konbini` from `map(string: dynamic)` to `an object`
* ⚠️ Change type of `V2BillingCadenceSettingsDataCollectionPaymentMethodOptions.SEPADebit`, `V2BillingCollectionSettingPaymentMethodOptions.SEPADebit`, `V2BillingCollectionSettingPaymentMethodOptionsParams.SEPADebit`, and `V2BillingCollectionSettingVersionPaymentMethodOptions.SEPADebit` from `map(string: dynamic)` to `an object`
* Add support for `ID` on `V2BillingCadenceSpendModifierMaxBillingPeriodSpendAmountCustomPricingUnit`, `V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendAmountCustomPricingUnitParams`, and `V2BillingIntentActionApplySpendModifierRuleMaxBillingPeriodSpendAmountCustomPricingUnit`
* Add support for new values `outbound_payments.paper_checks` and `paper_checks` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new values `bm_crn`, `bo_tin`, `bt_tpn`, `co_nit`, `ec_ruc`, `eg_tin`, `gh_tin`, `gy_tin`, `hn_rtn`, `jm_trn`, `jo_crn`, `ke_pin`, `ky_crn`, `lk_tin`, `mo_tin`, `mv_tin`, `ng_tin`, `pa_ruc`, `ph_tin`, `py_ruc`, `sl_tin`, `sv_nit`, `uy_ruc`, `vg_cn`, and `za_tin` on enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
* Add support for new values `bm_pp`, `bo_ci`, `bt_cid`, `eg_tin`, `gh_pin`, `gy_tin`, `hn_rtn`, `jm_trn`, `jo_pin`, `ky_pp`, `lk_nic`, `mo_bir`, `mt_nic`, `mv_tin`, `pa_ruc`, `ph_tin`, `py_ruc`, `si_pin`, `sv_nit`, and `vg_pp` on enums `V2CoreAccountIdentityIndividualIdNumber.Type` and `V2CoreAccountPersonIdNumber.Type`
* ⚠️ Change type of `V2CoreEventReasonRequestClient.StripeAction` from `map(string: dynamic)` to `an object`
* ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitProcessing` from `map(string: dynamic)` to `an object`
* ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitQueued` from `map(string: dynamic)` to `an object`
* ⚠️ Change type of `V2MoneyManagementInboundTransferTransferHistory.BankDebitSucceeded` from `map(string: dynamic)` to `an object`
* Add support for new values `paper_check_attachment_too_large`, `paper_check_expired`, and `paper_check_undeliverable` on enum `V2MoneyManagementOutboundPaymentStatusDetailsFailed.Reason`
* ⚠️ Remove support for `Town` on `V2MoneyManagementOutboundPaymentTrackingDetailsPaperCheckMailingAddress`
* Add support for new value `payout_method_amount_limit_exceeded` on enum `V2MoneyManagementOutboundTransferStatusDetailsFailed.Reason`
* Add support for `ApplicationFeeAmountRequested` on `V2PaymentsOffSessionPayment`
* ⚠️ Remove support for `CompartmentID` on `V2PaymentsOffSessionPayment`
* Add support for new value `exceeded_retry_window` on enum `V2PaymentsOffSessionPayment.FailureReason`
* Add support for `RetryUntil` on `V2PaymentsOffSessionPaymentRetryDetails`
* Add support for new value `paused` on enum `V2PaymentsOffSessionPayment.Status`
* Add support for `ApplicationFeeAmount` on `V2PaymentsOffSessionPaymentCaptureParams` and `V2PaymentsOffSessionPaymentParams`
* Add support for new value `paper_checks` on enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for new value `outbound_payments.paper_checks` on enum `EventsV2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for `AlertID` on `EventsV2CoreHealthApiErrorResolvedEvent`, `EventsV2CoreHealthApiLatencyResolvedEvent`, `EventsV2CoreHealthAuthorizationRateDropResolvedEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestErrorsFiringEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestErrorsResolvedEvent`, `EventsV2CoreHealthIssuingAuthorizationRequestTimeoutResolvedEvent`, `EventsV2CoreHealthPaymentMethodErrorResolvedEvent`, `EventsV2CoreHealthSepaDebitDelayedFiringEvent`, `EventsV2CoreHealthSepaDebitDelayedResolvedEvent`, `EventsV2CoreHealthTrafficVolumeDropResolvedEvent`, and `EventsV2CoreHealthWebhookLatencyResolvedEvent`
* Add support for `APIKey` on `EventsV2IamApiKeyCreatedEvent`, `EventsV2IamApiKeyDefaultSecretRevealedEvent`, `EventsV2IamApiKeyExpiredEvent`, `EventsV2IamApiKeyPermissionsUpdatedEvent`, `EventsV2IamApiKeyRotatedEvent`, and `EventsV2IamApiKeyUpdatedEvent`
* Add support for `StripeAccessGrant` on `EventsV2IamStripeAccessGrantApprovedEvent`, `EventsV2IamStripeAccessGrantCanceledEvent`, `EventsV2IamStripeAccessGrantDeniedEvent`, `EventsV2IamStripeAccessGrantRemovedEvent`, `EventsV2IamStripeAccessGrantRequestedEvent`, and `EventsV2IamStripeAccessGrantUpdatedEvent`
* Add support for event notifications `V2DataReportingQueryRunCreatedEvent`, `V2DataReportingQueryRunFailedEvent`, `V2DataReportingQueryRunSucceededEvent`, and `V2DataReportingQueryRunUpdatedEvent` with related object `V2DataReportingQueryRun`
* Add support for event notifications `V2PaymentsOffSessionPaymentPausedEvent` and `V2PaymentsOffSessionPaymentResumedEvent` with related object `V2PaymentsOffSessionPayment`
