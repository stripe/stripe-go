---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2474
semver_level: major
is_stripe_api_change: true
---

* ⚠️ Remove support for `Capture` method on resource `V2PaymentsOffSessionPayment`
* ⚠️ Remove support for `AcknowledgeConfirmationOfPayee` and `InitiateConfirmationOfPayee` methods on resource `V2CoreVaultGBBankAccount`
* Add support for `WeChatPayMobileWebPayments` on `AccountSettingsParams` and `AccountSettings`
* ⚠️ Remove support for `WeChatPayPayments` on `AccountSettingsParams` and `AccountSettings`
* Add support for `SettlementReserved` on `Balance`
* Add support for new value `settlement_reserved` on enum `BalanceTransaction.BalanceType`
* Add support for `Carecredit`, `Getflex`, and `Sezzle` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `MandateOptions` on `CheckoutSessionPaymentMethodOptionsCardParams`
* Add support for new value `auto` on enum `CheckoutSession.PaymentMethodCollection`
* Add support for new values `carecredit`, `getflex`, and `sezzle` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for new values `three_d_secure.authentication.canceled`, `three_d_secure.authentication.challenge_started`, `three_d_secure.authentication.errored`, `three_d_secure.authentication.failed`, `three_d_secure.authentication.requires_challenge`, `three_d_secure.authentication.requires_submission`, and `three_d_secure.authentication.succeeded` on enum `Event.Type`
* Add support for new values `carecredit`, `getflex`, and `sezzle` on enums `PaymentIntent.AllowedPaymentMethodTypes` and `SetupIntent.AllowedPaymentMethodTypes`
* Add support for new values `carecredit`, `getflex`, and `sezzle` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `ContactEmail` on `V2CoreAccountEvaluationAccountDataParams`, `V2CoreAccountEvaluationAccountData`, `V2SignalsAccountActivityAccountDetailsDataParams`, `V2SignalsAccountActivityAccountDetailsData`, `V2SignalsAccountEvaluationAccountDetailsDataParams`, and `V2SignalsAccountEvaluationAccountDetailsData`
* Add support for `BreB`, `Nip`, and `Pix` on `V2MoneyManagementFinancialAddressBankAccount` and `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccount`
* Add support for new values `bre_b`, `nip`, and `pix` on enum `V2MoneyManagementReceivedCreditBankTransferOriginatingBankAccount.Type`
* ⚠️ Remove support for `AmountCapturable` on `V2PaymentsOffSessionPayment`
* ⚠️ Remove support for `Capture` on `V2PaymentsOffSessionPaymentParams` and `V2PaymentsOffSessionPayment`
* Add support for snapshot events `EventTypeThreeDSecureAuthenticationCanceled`, `EventTypeThreeDSecureAuthenticationChallengeStarted`, `EventTypeThreeDSecureAuthenticationErrored`, `EventTypeThreeDSecureAuthenticationFailed`, `EventTypeThreeDSecureAuthenticationRequiresChallenge`, `EventTypeThreeDSecureAuthenticationRequiresSubmission`, and `EventTypeThreeDSecureAuthenticationSucceeded` with resource `ThreeDSecureAuthentication`
* ⚠️ Remove support for event notification `V2PaymentsOffSessionPaymentRequiresCaptureEvent` with related object `V2PaymentsOffSessionPayment`
