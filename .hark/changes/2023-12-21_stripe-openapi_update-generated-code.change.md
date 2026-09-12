---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1789
is_stripe_api_change: true
released_in_version: 76.11.0
---

* Add support for new resource `FinancialConnections.Transaction`
* Add support for `Get` and `List` methods on resource `Transaction`
* Add support for `Subscribe` and `Unsubscribe` methods on resource `FinancialConnections.Account`
* Add support for `Features` on `AccountSessionComponentsPayoutsParams`
* Add support for `EditPayoutSchedule`, `InstantPayouts`, and `StandardPayouts` on `AccountSessionComponentsPayoutsFeatures`
* Change type of `CheckoutSessionPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetchParams`, `CheckoutSessionPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetch`, `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetchParams`, `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetch`, `PaymentIntentConfirmPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetchParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetchParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetch`, `SetupIntentConfirmPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetchParams`, `SetupIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetchParams`, `SetupIntentPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetch`, `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetchParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPrefetch` from `literal('balances')` to `enum('balances'|'transactions')`
* Add support for new value `financial_connections.account.refreshed_transactions` on enum `EventType`
* Add support for `Subscriptions` and `TransactionRefresh` on `FinancialConnectionsAccount`
* Add support for `NextRefreshAvailableAt` on `FinancialConnectionsAccountBalanceRefresh`
* Add support for new value `transactions` on enum `FinancialConnectionsSessionPrefetch`
* Add support for new value `unknown` on enum `IssuingAuthorizationVerificationDataAuthenticationExemptionType`
* Add support for new value `challenge` on enums `PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure` and `SetupIntentPaymentMethodOptionsCardRequestThreeDSecure`
* Add support for `RevolutPay` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
* Change type of `QuoteInvoiceSettings` from `nullable(InvoiceSettingQuoteSetting)` to `InvoiceSettingQuoteSetting`
* Add support for `DestinationDetails` on `Refund`
