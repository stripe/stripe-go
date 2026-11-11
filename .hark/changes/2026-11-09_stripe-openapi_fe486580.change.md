---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2491
semver_level: major
is_stripe_api_change: true
released_in_version: 87.2.0-alpha.3
---

* Add support for new resource `V2MoneyManagementSpendingMandate`
* Add support for `Cancel`, `Get`, and `List` methods on resource `V2MoneyManagementSpendingMandate`
* Add support for `CardholderMessage` on `ChargePaymentMethodDetailsCardPresentReceipt`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresentReceipt`, `PaymentAttemptRecordPaymentMethodDetailsCardPresentReceipt`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresentReceipt`, and `PaymentRecordPaymentMethodDetailsCardPresentReceipt`
* Add support for `PricingToken` on `CheckoutSessionParams`
* Add support for `MarketingAttributions` on `CustomerParams`
* Add support for `FraudulentSelfie` on `IdentityVerificationReportSignals`
* Add support for `TossPay` on `MandatePaymentMethodDetails`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, and `SetupAttemptPaymentMethodDetails`
* Add support for `PaymentEvaluations` on `PaymentAttemptRecordReportAuthenticatedParams`
* Add support for `ThreeDSecure` on `PaymentAttemptRecordReportAuthorizedPaymentMethodDetailsCardParams`, `PaymentAttemptRecordReportGuaranteedPaymentMethodDetailsCardParams`, `PaymentRecordReportPaymentAttemptGuaranteedPaymentMethodDetailsCardParams`, `PaymentRecordReportPaymentAttemptPaymentMethodDetailsCardParams`, and `PaymentRecordReportPaymentPaymentMethodDetailsCardParams`
* Add support for `WeChatPayRedirectToMiniProgram` on `PaymentIntentNextAction`
* Add support for `CustomerPresence` and `PaymentType` on `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsSepaDebitParams` and `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsSepaDebit`
* Add support for `FinancialConnections` on `SetupIntentConfirmPaymentMethodOptionsBacsDebitParams`, `SetupIntentPaymentMethodOptionsBacsDebitParams`, and `SetupIntentPaymentMethodOptionsBacsDebit`
* Add support for `ReferralFeeDetails` on `SharedPaymentGrantedToken`, `SharedPaymentIssuedTokenParams`, and `SharedPaymentIssuedToken`
* Add support for `DeviceHealthDetails` and `DeviceHealth` on `TerminalReader`
* Add support for new values `check_attachment_too_large`, `check_expired`, and `check_undeliverable` on enum `V2MoneyManagementOutboundPaymentStatusDetailsFailed.Reason`
* Add support for `CheckSend` on `V2MoneyManagementOutboundPaymentToPayoutMethodOptionsParams`, `V2MoneyManagementOutboundPaymentToPayoutMethodOptions`, `V2MoneyManagementOutboundPaymentTrackingDetails`, `V2MoneyManagementOutboundSetupIntentPayoutMethodDataParams`, `V2MoneyManagementPayoutIntentToPayoutMethodOptionsParams`, `V2MoneyManagementPayoutIntentToPayoutMethodOptions`, and `V2MoneyManagementPayoutMethod`
* Add support for new values `check_attachment_too_large`, `check_expired`, and `check_undeliverable` on enum `V2MoneyManagementPayoutIntentNextActionHandleFailure.FailureReason`
* Add support for new value `check_send` on enum `V2MoneyManagementPayoutMethod.Type`
* Add support for `SpendingMandate` on `V2MoneyManagementReceivedDebitBankTransfer` and `V2MoneyManagementReceivedDebitListParams`
* Add support for event notifications `V2MoneyManagementSpendingMandateCanceledEvent`, `V2MoneyManagementSpendingMandateCreatedEvent`, `V2MoneyManagementSpendingMandateExpiredEvent`, `V2MoneyManagementSpendingMandatePendingCancellationEvent`, and `V2MoneyManagementSpendingMandateUpdatedEvent` with related object `V2MoneyManagementSpendingMandate`
* ⚠️ Add error type `MerchantNotGatedError` on `V2CoreVaultNetworkToken#CreateFromCredential`
