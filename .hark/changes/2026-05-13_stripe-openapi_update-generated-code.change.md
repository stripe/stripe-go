---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2353
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.2.0-alpha.5
---

* Add support for new resources `V2CoreFeeBatch`, `V2CoreFeeEntry`, `V2MoneyManagementDebitDispute`, and `V2MoneyManagementFinancialAccountStatement`
* Add support for `SimulateNetworkLifecyclePreArbitrationResponse` and `SimulateNetworkLifecyclePreArbitrationSubmission` test helper methods on resource `IssuingDispute`
* Add support for `List` method on resource `PaymentLocation`
* Add support for `Get` and `List` methods on resources `V2CoreFeeBatch`, `V2CoreFeeEntry`, and `V2MoneyManagementFinancialAccountStatement`
* Add support for `Get`, `List`, and `New` methods on resource `V2MoneyManagementDebitDispute`
* Add support for `Discounts` on `DelegatedCheckoutRequestedSessionParams` and `DelegatedCheckoutRequestedSession`
* Add support for `AmountSale` on `DelegatedCheckoutRequestedSessionLineItemDetail` and `DelegatedCheckoutRequestedSessionTotalDetails`
* Add support for `AmountDiscount` and `Breakdown` on `DelegatedCheckoutRequestedSessionTotalDetails`
* ⚠️ Remove support for `CheckDepositAddress` on `InvoicePaymentSettingsPaymentMethodOptionsCheckScanParams`, `InvoicePaymentSettingsPaymentMethodOptionsCheckScan`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsCheckScan`, `SubscriptionPaymentSettingsPaymentMethodOptionsCheckScanParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsCheckScan`
* Add support for `PaymentEvaluations` on `PaymentAttemptRecordReportGuaranteedParams`, `PaymentRecordReportPaymentAttemptGuaranteedParams`, and `PaymentRecordReportPaymentGuaranteedParams`
* Add support for `Location` on `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentPaymentDetailsParams`, `SetupIntentConfirmSetupDetailsParams`, and `SetupIntentSetupDetailsParams`
* Add support for `OnboardingDataUpdateAcknowledged` on `PaymentLocationParams`
* Add support for `Customer` on `RadarCustomerEvaluationParams`
* Add support for `Status` on `RadarCustomerEvaluationParams` and `RadarCustomerEvaluation`
* Add support for `PaymentBehavior` on `SubscriptionResumeParams`
* Add support for `DisputeDetails` on `V2MoneyManagementReceivedDebit`
* Add support for new value `debit_dispute` on enums `V2MoneyManagementTransaction.Category` and `V2MoneyManagementTransactionEntryTransactionDetails.Category`
* Add support for `DebitDispute` on `V2MoneyManagementTransactionEntryTransactionDetailsFlow` and `V2MoneyManagementTransactionFlow`
* Add support for new value `debit_dispute` on enums `V2MoneyManagementTransactionEntryTransactionDetailsFlow.Type` and `V2MoneyManagementTransactionFlow.Type`
* Add support for `PaymentAttemptRecord` on `EventsV2PaymentsOffSessionPaymentAttemptFailedEvent` and `EventsV2PaymentsOffSessionPaymentFailedEvent`
* Add support for event notifications `V2MoneyManagementFinancialAccountStatementCreatedEvent` and `V2MoneyManagementFinancialAccountStatementRestatedEvent` with related object `V2MoneyManagementFinancialAccountStatement`
