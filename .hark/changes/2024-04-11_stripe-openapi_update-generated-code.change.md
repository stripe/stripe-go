---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1846
is_stripe_api_change: true
released_in_version: 78.1.0
---

* Add support for `AccountManagement` and `NotificationBanner` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for `ExternalAccountCollection` on `AccountSessionComponentsAccountOnboardingFeaturesParams` and `AccountSessionComponentsAccountOnboardingFeatures`
* Add support for new values `billing_policy_remote_function_response_invalid`, `billing_policy_remote_function_timeout`, `billing_policy_remote_function_unexpected_status_code`, and `billing_policy_remote_function_unreachable` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Change type of `BillingMeterEventAdjustmentCancel` from `BillingMeterResourceBillingMeterEventAdjustmentCancel` to `nullable(BillingMeterResourceBillingMeterEventAdjustmentCancel)`
* Add support for `AmazonPay` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `RefundDestinationDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodDataParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for new values `bh_vat`, `kz_bin`, `ng_tin`, and `om_vat` on enums `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, `TaxIdType`, and `TaxTransactionCustomerDetailsTaxIdsType`
* Add support for new value `ownership` on enums `CheckoutSessionPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetch`, `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetch`, `PaymentIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetch`, `SetupIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetch`, and `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetch`
* Add support for new value `amazon_pay` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
* Add support for `NextRefreshAvailableAt` on `FinancialConnectionsAccountOwnershipRefresh`
* Add support for new value `ownership` on enums `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPermissions` and `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPermissions`
