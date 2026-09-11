---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2388
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.2.0-alpha.4
---

* ⚠️ Remove support for resource `FRMealVouchersOnboarding`
* ⚠️ Remove support for `Get`, `List`, `New`, and `Update` methods on resource `FRMealVouchersOnboarding`
* Add support for `New` method on resource `PaymentRecord`
* Add support for new value `chaps` on enums `FundingInstructionsBankTransferFinancialAddress.SupportedNetworks` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddress.SupportedNetworks`
* ⚠️ Remove support for `FinancialAccountsTransactions`, `FinancialAccounts`, and `RecipientsList` on `AccountSessionComponentsParams`
* Add support for `SmartDisputesManagement` on `AccountSessionComponentsDisputesListFeatures`, `AccountSessionComponentsPaymentDetailsFeatures`, `AccountSessionComponentsPaymentDisputesFeatures`, and `AccountSessionComponentsPaymentsFeatures`
* Add support for new value `ic_nif` on enums `CheckoutSessionCollectedInformationTaxId.Type`, `CheckoutSessionCustomerDetailsTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, `QuotePreviewInvoiceCustomerTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
* Add support for new values `financial_connections.account.expected_deactivation_date_updated`, `financial_connections.account.supported_payment_method_types_updated`, `financial_connections.account.upcoming_deactivation`, `financial_connections.authorization.expected_deactivation_date_updated`, and `financial_connections.authorization.upcoming_deactivation` on enum `Event.Type`
* Add support for `Mode` on `FinancialConnectionsSessionManualEntry`
* Add support for new values `alipay` and `sequra` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for new value `stripe_internal_error` on enum `IssuingAuthorizationRequestHistory.Reason`
* Add support for `BusinessName` on `IssuingCardShipping`
* Add support for new value `correos` on enum `IssuingCardShipping.Carrier`
* ⚠️ Change type of `IssuingTransactionNetworkData.TraceID` from `IssuingTransactionTraceId` to `nullable(IssuingTransactionTraceId)`
* Add support for `PauseSchedules` on `QuotePreviewSubscriptionSchedule`, `SubscriptionScheduleParams`, and `SubscriptionSchedule`
* Add support for `Trial` on `QuotePreviewSubscriptionSchedulePhase` and `SubscriptionSchedulePhase`
* Add support for `PaymentRecord` on `RefundParams`
* Add support for `RedirectToURL` on `SharedPaymentIssuedTokenNextAction`
* ⚠️ Change type of `SharedPaymentIssuedTokenNextAction.Type` from `literal('use_stripe_sdk')` to `enum('redirect_to_url'|'use_stripe_sdk')`
* Add support for snapshot events `EventTypeFinancialConnectionsAccountExpectedDeactivationDateUpdated`, `EventTypeFinancialConnectionsAccountSupportedPaymentMethodTypesUpdated`, and `EventTypeFinancialConnectionsAccountUpcomingDeactivation` with resource `FinancialConnectionsAccount`
* Add support for snapshot events `EventTypeFinancialConnectionsAuthorizationExpectedDeactivationDateUpdated` and `EventTypeFinancialConnectionsAuthorizationUpcomingDeactivation` with resource `FinancialConnectionsAuthorization`
