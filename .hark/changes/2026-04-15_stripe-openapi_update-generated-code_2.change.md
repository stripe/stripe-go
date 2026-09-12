---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2343
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.1.0-alpha.4
---

* Add support for new resources `V2CoreWorkflowRun` and `V2CoreWorkflow`
* Add support for `ReportAuthorized` method on resource `PaymentAttemptRecord`
* Add support for `Get` and `List` methods on resource `V2CoreWorkflowRun`
* Add support for `Get`, `Invoke`, and `List` methods on resource `V2CoreWorkflow`
* Add support for `NextAction` and `Status` on `SharedPaymentIssuedToken`
* ⚠️ Remove support for `NetworkID` on `SharedPaymentIssuedTokenSellerDetails`
* Add support for `Bills` on `AccountSessionComponents`
* Add support for `SettlementCurrencies` on `BalanceSettingsPaymentsParams` and `BalanceSettingsPayments`
* Add support for `DefaultSettlementCurrency` on `BalanceSettingsPayments`
* Add support for `AccountFunding` on `ChargePaymentMethodDetailsCard`
* Add support for `AutomaticSurcharge` on `CheckoutSessionParams`, `CheckoutSession`, `PaymentLinkParams`, and `PaymentLink`
* Add support for `Bizum` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for `SurchargeCost` on `CheckoutSession`
* Add support for `AmountSurcharge` on `CheckoutSessionTotalDetails`
* Add support for `SharedPaymentGrantedToken` on `ConfirmationTokenPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Details` on `IdentityVerificationReportEmail`
* Add support for new value `email` on enums `IdentityVerificationReport.Type` and `IdentityVerificationSession.Type`
* Add support for `Confirm` on `IdentityVerificationSessionParams`
* Add support for `Subscription` on `InvoiceItemParentScheduleDetails`
* ⚠️ Remove support for `SharedPaymentGrantedToken` on `PaymentIntentConfirmParams` and `PaymentIntentParams`
* Add support for `MoneyServices` on `PaymentIntentPaymentDetails`
* ⚠️ Remove support for `ExternalReference` on `Plan`
