---
title: API Updates for beta branch
pr_link: https://github.com/stripe/stripe-go/pull/1601
is_stripe_api_change: true
released_in_version: 74.8.0-beta.1
---

* Updated stable APIs to the latest version
* Add support for new resource `FinancialConnections.Transaction`
* Add support for `List` method on resource `Transaction`
* Add support for `Prefetch` on `-PaymentMethodOptionsUsBankAccountFinancialConnectionsParams` and `-PaymentMethodOptionsUsBankAccountFinancialConnections` across several APIs.
* * Add support for `InferredBalancesRefresh`, `Subscriptions`, and `TransactionRefresh` on `FinancialConnectionsAccount`
* Add support for `ManualEntry` on `FinancialConnectionsSessionParams` and `FinancialConnectionsSession`
* Add support for `StatusDetails` and `Status` on `FinancialConnectionsSession`
* Add support for new value `ownership` on enums `InvoicePaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPermissions` and `SubscriptionPaymentSettingsPaymentMethodOptionsUsBankAccountFinancialConnectionsPermissions`
* Add support for `AccountNumber` on `PaymentMethodUsBankAccount`
* Remove support for `ID` on `QuoteLinesStartsAtLineEndsAtParams`
