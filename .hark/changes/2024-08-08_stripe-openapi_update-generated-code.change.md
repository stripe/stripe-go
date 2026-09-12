---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1899
is_stripe_api_change: true
released_in_version: 79.7.0
---

* Add support for `Activate`, `Archive`, `Deactivate`, `Get`, `List`, and `New` methods on resource `Billing.Alert`
* Add support for `Get` method on resource `Tax.Calculation`
* Add support for new value `invalid_mandate_reference_prefix_format` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `Type` on `ChargePaymentMethodDetailsCardPresentOffline`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresentOffline`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresentOffline`, and `SetupAttemptPaymentMethodDetailsCardPresentOffline`
* Add support for `Offline` on `ConfirmationTokenPaymentMethodPreviewCardPresent` and `PaymentMethodCardPresent`
* Add support for `RelatedCustomer` on `IdentityVerificationSessionListParams`, `IdentityVerificationSessionParams`, and `IdentityVerificationSession`
* Add support for new value `girocard` on enums `PaymentIntentPaymentMethodOptionsCardNetwork`, `SetupIntentPaymentMethodOptionsCardNetwork`, and `SubscriptionPaymentSettingsPaymentMethodOptionsCardNetwork`
* Add support for new value `financial_addresses.aba.forwarding` on enums `TreasuryFinancialAccountActiveFeatures`, `TreasuryFinancialAccountPendingFeatures`, and `TreasuryFinancialAccountRestrictedFeatures`
