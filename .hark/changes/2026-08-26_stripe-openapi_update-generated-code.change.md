---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2414
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.5.0-alpha.1
---

* Add support for new resource `CustomerTaxExemption`
* Add support for `Del`, `Get`, `List`, and `New` methods on resource `CustomerTaxExemption`
* Add support for `Details` on `AccountFutureRequirementsErrors`, `AccountRequirementsErrors`, `BankAccountFutureRequirementsErrors`, `BankAccountRequirementsErrors`, `CapabilityFutureRequirementsError`, and `PersonFutureRequirementsError`
* ⚠️ Remove support for `SequraPayments` on `AccountCapabilities`
* Add support for `SubscriptionPause` on `BillingPortalSessionFlowDataParams`
* ⚠️ Remove support for `Sequra` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentPaymentMethodOptions`, and `PaymentRecordPaymentMethodDetails`
* Add support for `EnablementDetails` on `CheckoutSessionAutomaticTax`
* ⚠️ Remove support for value `sequra` from enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `Credit` on `FinancialConnectionsTransactionClassifications`
* Change type of `FinancialConnectionsTransactionClassifications.MoneyMovement` from `nullable(BankConnectionsResourceTransactionResourceClassificationsLabels)` to `BankConnectionsResourceTransactionResourceClassificationsLabels`
* Change type of `FinancialConnectionsTransactionClassifications.PersonalFinance` from `nullable(BankConnectionsResourceTransactionResourceClassificationsLabels)` to `BankConnectionsResourceTransactionResourceClassificationsLabels`
* Add support for `UserConsent` on `IdentityVerificationSessionParams`
* Add support for `CompanyDetails` on `InvoicePaymentSettingsPaymentMethodOptionsBillie`, `PaymentIntentConfirmPaymentMethodOptionsBillieParams`, `PaymentIntentPaymentMethodOptionsBillieParams`, `PaymentIntentPaymentMethodOptionsBillie`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBillie`, and `SubscriptionPaymentSettingsPaymentMethodOptionsBillie`
* Add support for `Reference` on `InvoicePaymentSettingsPaymentMethodOptionsBillie`, `PaymentIntentConfirmPaymentMethodOptionsBillieParams`, `PaymentIntentPaymentMethodOptionsBillieParams`, `PaymentIntentPaymentMethodOptionsBillie`, and `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBillie`
* Add support for `PosCondition` on `IssuingAuthorizationParams` and `IssuingAuthorization`
* Add support for `CryptoWallet` on `IssuingCardParams` and `IssuingCard`
* Add support for `PaymentEvaluations` and `PaymentMethodDetails` on `PaymentAttemptRecordReportAuthorizedParams`
* Add support for `AadeData` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams` and `PaymentIntentPaymentMethodOptionsCardPresentParams`
* ⚠️ Remove support for `CancelAtPeriodEnd` on `SubscriptionPendingUpdate`
* Add support for `BLIKRecurringPayments` on `V2CoreAccountConfigurationMerchantCapabilitiesParams` and `V2CoreAccountConfigurationMerchantCapabilities`
* Add support for `UserAccess` on `V2IamActivityLogDetails`
* Add support for new value `user_access` on enum `V2IamActivityLogDetails.Type`
* Add support for new value `user_access_started` on enum `V2IamActivityLog.Type`
* Add support for new value `blik_recurring_payments` on enum `EventsV2CoreAccountIncludingConfigurationMerchantCapabilityStatusUpdatedEvent.UpdatedCapability`
