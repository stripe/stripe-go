---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1743
is_stripe_api_change: true
released_in_version: 75.9.0
---

* Add support for new resource `Issuing.Token`
* Add support for `Get`, `List`, and `Update` methods on resource `Token`
* Add support for `AmountAuthorized`, `ExtendedAuthorization`, `IncrementalAuthorization`, `Multicapture`, and `Overcapture` on `ChargePaymentMethodDetailsCard`
* Add support for `Token` on `IssuingAuthorization` and `IssuingTransaction`
* Add support for `AuthorizationCode` on `IssuingAuthorizationRequestHistory`
* Add support for `RequestExtendedAuthorization`, `RequestMulticapture`, and `RequestOvercapture` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, and `PaymentIntentPaymentMethodOptionsCard`
* Add support for `RequestIncrementalAuthorization` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, and `PaymentIntentPaymentMethodOptionsCard`
* Add support for `FinalCapture` on `PaymentIntentCaptureParams`
* Add support for `Metadata` on `PaymentLinkPaymentIntentDataParams`, `PaymentLinkPaymentIntentData`, `PaymentLinkSubscriptionDataParams`, and `PaymentLinkSubscriptionData`
* Add support for `StatementDescriptorSuffix` and `StatementDescriptor` on `PaymentLinkPaymentIntentDataParams` and `PaymentLinkPaymentIntentData`
* Add support for `PaymentIntentData` and `SubscriptionData` on `PaymentLinkParams`
