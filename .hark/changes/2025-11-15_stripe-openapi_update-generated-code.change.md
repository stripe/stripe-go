---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/2206
is_stripe_api_change: true
released_in_version: 84.0.0
---

* Add support for new resources `TaxAssociation` and `TerminalOnboardingLink`
* Add support for `Find` method on resource `TaxAssociation`
* Add support for `New` method on resource `TerminalOnboardingLink`
* Add support for `PaymentMethodConfiguration` on `BillingPortalConfigurationFeaturesPaymentMethodUpdate`
* Add support for `TransactionID` on `ChargePaymentMethodDetailsIdeal`, `PaymentAttemptRecordPaymentMethodDetailsIdeal`, and `PaymentRecordPaymentMethodDetailsIdeal`
* Add support for new value `finom` on enums `ConfirmationTokenPaymentMethodPreviewIdeal.Bank`, `PaymentAttemptRecordPaymentMethodDetailsIdeal.Bank`, and `PaymentRecordPaymentMethodDetailsIdeal.Bank`
* Add support for new value `FNOMNL22` on enums `ConfirmationTokenPaymentMethodPreviewIdeal.BIC`, `PaymentAttemptRecordPaymentMethodDetailsIdeal.BIC`, and `PaymentRecordPaymentMethodDetailsIdeal.BIC`
* Add support for new value `tokenized_account_number_deactivated` on enums `ConfirmationTokenPaymentMethodPreviewUsBankAccountStatusDetailsBlocked.Reason` and `PaymentMethodUsBankAccountStatusDetailsBlocked.Reason`
* Add support for `Created` on `CustomerCustomerBalanceTransactionListParams` and `InvoicePaymentListParams`
* Add support for new values `financial_connections.account.account_numbers_updated` and `financial_connections.account.upcoming_account_number_expiry` on enum `Event.Type`
* Add support for `AccountNumbers` on `FinancialConnectionsAccount`
* Add support for `FraudRisk` on `IssuingAuthorizationRiskAssessmentParams`
* Add support for `LatestFraudWarning` on `IssuingCard`
* Add support for `Hooks` on `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentIncrementAuthorizationParams`, `PaymentIntentParams`, and `PaymentIntent`
* Add support for `MbWay` and `TWINT` on `RefundDestinationDetails`
* Add support for snapshot events `EventTypeFinancialConnectionsAccountAccountNumbersUpdated` and `EventTypeFinancialConnectionsAccountUpcomingAccountNumberExpiry` with resource `FinancialConnectionsAccount`
