---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1769
is_stripe_api_change: true
released_in_version: 76.6.0
---

* Add support for `ElectronicCommerceIndicator` on `ChargePaymentMethodDetailsCardThreeDSecure` and `SetupAttemptPaymentMethodDetailsCardThreeDSecure`
* Add support for `ExemptionIndicatorApplied` and `ExemptionIndicator` on `ChargePaymentMethodDetailsCardThreeDSecure`
* Add support for `TransactionID` on `ChargePaymentMethodDetailsCardThreeDSecure`, `IssuingAuthorizationNetworkData`, `IssuingTransactionNetworkData`, and `SetupAttemptPaymentMethodDetailsCardThreeDSecure`
* Add support for `Offline` on `ChargePaymentMethodDetailsCardPresent`
* Add support for `SystemTraceAuditNumber` on `IssuingAuthorizationNetworkData`
* Add support for `NetworkRiskScore` on `IssuingAuthorizationPendingRequest` and `IssuingAuthorizationRequestHistory`
* Add support for `RequestedAt` on `IssuingAuthorizationRequestHistory`
* Add support for `AuthorizationCode` on `IssuingTransactionNetworkData`
* Add support for `ThreeDSecure` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `SetupIntentConfirmPaymentMethodOptionsCardParams`, and `SetupIntentPaymentMethodOptionsCardParams`
