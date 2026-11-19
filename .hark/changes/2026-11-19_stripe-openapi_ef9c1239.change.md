---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2496
semver_level: minor
is_stripe_api_change: true
released_in_version: 87.2.0
---

* Add support for new resource `V2CoreClaimableSandbox`
* Add support for `Get`, `New`, and `RenewOnboardingLink` methods on resource `V2CoreClaimableSandbox`
* Add support for `FundingSources` on `ChargePaymentMethodDetailsCardWalletLink` and `ChargePaymentMethodDetailsLink`
* Add support for new value `tron` on enums `ChargePaymentMethodDetailsCrypto.Network`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.Network`, and `PaymentRecordPaymentMethodDetailsCrypto.Network`
* Add support for `Fingerprint` on `ChargePaymentMethodDetailsPaypay`, `PaymentAttemptRecordPaymentMethodDetailsPaypay`, and `PaymentRecordPaymentMethodDetailsPaypay`
* Add support for `AddressCollectionPrecision` on `CheckoutSessionAutomaticTaxParams` and `CheckoutSessionAutomaticTax`
* Add support for `CarRentalData` and `LodgingData` on `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentPaymentDetailsParams`, and `PaymentIntentPaymentDetails`
* Change type of `PaymentLinkPaymentIntentDataParams.CaptureMethod` from `enum('automatic'|'automatic_async'|'manual')` to `emptyable(enum('automatic'|'automatic_async'|'manual'))`
* Add support for `CaptureMethod` on `PaymentLinkPaymentIntentDataParams`
* Add support for `LossMessage` on `ProductMarketingFeatureParams`, `ProductMarketingFeature`, and `ProductMarketingFeaturesParams`
* Add support for `Application` on `Refund`
* Add support for `Bizum` and `Upi` on `RefundDestinationDetails`
* Add support for `TaxInvoiceMonth` on `ReportingReportRunParametersParams` and `ReportingReportRunParameters`
* Add support for `Igic` on `TaxRegistrationCountryOptionsEsParams` and `TaxRegistrationCountryOptionsEs`
* Add support for new value `igic` on enum `TaxRegistrationCountryOptionsEs.Type`
* Add support for new values `stripe_s700_devkit`, `stripe_s710_devkit`, `stripe_t600_devkit`, `verifone_ux700_devkit`, and `verifone_v660p_devkit` on enum `TerminalReader.DeviceType`
* Add support for `NetworkDetails` on `TestHelpersTreasuryInboundTransferFailFailureDetailsParams`, `TreasuryInboundTransferFailureDetails`, `TreasuryOutboundPaymentReturnedDetails`, and `TreasuryOutboundTransferReturnedDetails`
* Add support for `FailureDetails` on `TestHelpersTreasuryInboundTransferReturnInboundTransferParams`
* Add support for `AccountDashboard` on `V2CoreAccountLinkUseCaseParams` and `V2CoreAccountLinkUseCase`
* Add support for new value `account_dashboard` on enum `V2CoreAccountLinkUseCase.Type`
* Add support for `PaypayPayments` on `V2CoreAccountConfigurationMerchantCapabilitiesParams`, `V2CoreAccountConfigurationMerchantCapabilities`, `V2CoreAccountConfigurationMerchantParams`, and `V2CoreAccountConfigurationMerchant`
* Add support for `BilliePayments`, `MbWayPayments`, `NzBankAccountBECSDebitPayments`, and `PixPayments` on `V2CoreAccountConfigurationMerchantCapabilitiesParams` and `V2CoreAccountConfigurationMerchantCapabilities`
* Add support for `AdditionalMaterial` on `V2CoreAccountConfigurationMerchantKonbiniPaymentsParams` and `V2CoreAccountConfigurationMerchantKonbiniPayments`
* Add support for new values `billie_payments`, `mb_way_payments`, `nz_bank_account_becs_debit_payments`, `paypay_payments`, and `pix_payments` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new values `billie_payments`, `mb_way_payments`, `nz_bank_account_becs_debit_payments`, `paypay_payments`, and `pix_payments` on enum `EventsV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for event notifications `V2CoreClaimableSandboxClaimedEvent`, `V2CoreClaimableSandboxCreatedEvent`, `V2CoreClaimableSandboxExpiredEvent`, `V2CoreClaimableSandboxExpiringEvent`, and `V2CoreClaimableSandboxUpdatedEvent` with related object `V2CoreClaimableSandbox`
* Add support for error code `inactive_api_operation_confirmation_required` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, `StripeError`, and `TerminalReaderActionApiError`
* Add support for error codes `claimable_sandbox_creation_rate_limit_exceeded` and `claimable_sandbox_testing_creation_rate_limit_exceeded` on `RateLimitError`
