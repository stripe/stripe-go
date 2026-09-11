---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/2362
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.1.0
---

* Add support for `ReleaseDetails` on `ReserveHold`
* Add support for new value `tax_fund` on enum `BalanceTransaction.Type`
* Add support for `BuyerID` on `ChargePaymentMethodDetailsBizum`, `ConfirmationTokenPaymentMethodPreviewBizum`, `ConfirmationTokenPaymentMethodPreviewBlik`, `PaymentAttemptRecordPaymentMethodDetailsBizum`, `PaymentMethodBizum`, `PaymentMethodBlik`, and `PaymentRecordPaymentMethodDetailsBizum`
* Add support for `TransactionLinkID` on `ChargePaymentMethodDetailsCard`
* Add support for new value `sui` on enums `ChargePaymentMethodDetailsCrypto.Network`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.Network`, and `PaymentRecordPaymentMethodDetailsCrypto.Network`
* Add support for new value `usdsui` on enums `ChargePaymentMethodDetailsCrypto.TokenCurrency`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.TokenCurrency`, and `PaymentRecordPaymentMethodDetailsCrypto.TokenCurrency`
* Add support for `Fingerprint` on `ChargePaymentMethodDetailsPix`, `ConfirmationTokenPaymentMethodPreviewPix`, `PaymentMethodPix`, and `SetupAttemptPaymentMethodDetailsPix`
* Add support for `Sunbit` on `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptionsParams`, and `PaymentIntentPaymentMethodOptions`
* Add support for `BillingCycleAnchorConfig` on `CheckoutSessionSubscriptionDataParams`
* Add support for `WeChatPay` on `CheckoutSessionPaymentMethodOptions`
* Add support for `MastercardCompliance` on `DisputeEvidenceDetailsEnhancedEligibility`, `DisputeEvidenceEnhancedEvidenceParams`, and `DisputeEvidenceEnhancedEvidence`
* Add support for new value `mastercard_compliance` on enum `Dispute.EnhancedEligibilityTypes`
* Add support for `StatusDetails` on `FinancialConnectionsAccount`
* Add support for new value `validated` on enum `IdentityVerificationSessionRedaction.Status`
* Add support for new value `satispay` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* ⚠️ Remove support for `StoredCredentialUsage` on `PaymentAttemptRecordPaymentMethodDetailsCard` and `PaymentRecordPaymentMethodDetailsCard`
* Add support for `SetupFutureUsage` on `PaymentIntentConfirmPaymentMethodOptionsSatispayParams`, `PaymentIntentPaymentMethodOptionsSatispayParams`, and `PaymentIntentPaymentMethodOptionsSatispay`
* Add support for `Satispay` on `SetupAttemptPaymentMethodDetails`
* Add support for `CustomFields`, `Description`, and `Footer` on `SubscriptionInvoiceSettingsParams` and `SubscriptionInvoiceSettings`
* Add support for `PaymentMethodOptions` and `PaymentMethod` on `TopupParams`
* Add support for `Mode` on `V2CommerceProductCatalogImport`
* Add support for new value `promotion` on enum `V2CommerceProductCatalogImport.FeedType`
* Add support for `SunbitPayments` on `V2CoreAccountConfigurationMerchantCapabilitiesParams` and `V2CoreAccountConfigurationMerchantCapabilities`
* Add support for `CryptoMoneyManager` and `MoneyManager` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams`
* ⚠️ Remove support for `CryptoStorer` and `Storer` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams`
* Add support for new value `sunbit_payments` on enum `EventsV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent.UpdatedCapability`
* Add support for error codes `anomalous_money_movement_request`, `failed_tax_calculation`, `financial_account_balance_does_not_support_currency`, `financial_account_capability_not_enabled`, and `financial_account_capability_restricted` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, `StripeError`, and `TerminalReaderActionApiError`
