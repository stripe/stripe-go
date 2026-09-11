---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1882
is_breaking: true
is_stripe_api_change: true
released_in_version: 79.3.0-beta.1
---

* ⚠️ Remove support for `PaymentMethodUpdate` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`. Users are expected to completely migrate from using `payment_method_update`.
* Add support for new resource `FinancialConnections.Institution`
* Add support for `Get` and `List` methods on resource `Institution`
* Add support for `Institution` on `CheckoutSessionPaymentMethodOptionsUsBankAccountFinancialConnectionsFilters`, `FinancialConnectionsSessionFiltersParams`, `FinancialConnectionsSessionFilters`, `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsFilters`, `PaymentIntentConfirmPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsFilters`, `SetupIntentConfirmPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, `SetupIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, `SetupIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsFilters`, `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsFiltersParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsFilters`
* Add support for `PaymentMethodAllowRedisplayFilters`, `PaymentMethodRedisplayLimit`, `PaymentMethodRedisplay`, and `PaymentMethodSaveUsage` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`
* Add support for new value `balance` on enum `FinancialConnectionsAccountSubscriptions`
