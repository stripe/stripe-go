---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1729
is_stripe_api_change: true
released_in_version: 75.6.0
---

* Add support for `Capture`, `Expire`, `Increment`, `New`, and `Reverse` test helper methods on resource `Issuing.Authorization`
* Add support for `CreateForceCapture`, `CreateUnlinkedRefund`, and `Refund` test helper methods on resource `Issuing.Transaction`
* Add support for new value `stripe_tax_inactive` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `Nonce` on `EphemeralKeyParams`
* Add support for `CashbackAmount` on `IssuingAuthorizationAmountDetails`, `IssuingAuthorizationPendingRequestAmountDetails`, `IssuingAuthorizationRequestHistoryAmountDetails`, and `IssuingTransactionAmountDetails`
* Add support for `SerialNumber` on `TerminalReaderListParams`
