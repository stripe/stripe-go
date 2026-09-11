---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1862
is_stripe_api_change: true
released_in_version: 78.7.0
---

* Add support for `FeeSource` on `ApplicationFee`
* Add support for `NetAvailable` on `BalanceInstantAvailable`
* Add support for `PreferredLocales` on `ChargePaymentMethodDetailsCardPresent`, `ConfirmationTokenPaymentMethodPreviewCardPresent`, and `PaymentMethodCardPresent`
* Add support for `Klarna` on `DisputePaymentMethodDetails`
* Add support for new value `klarna` on enum `DisputePaymentMethodDetailsType`
* Add support for `Archived` and `LookupKey` on `EntitlementsFeatureListParams`
* Add support for `NoValidAuthorization` on `IssuingDisputeEvidenceParams` and `IssuingDisputeEvidence`
* Add support for `LossReason` on `IssuingDispute`
* Add support for new value `no_valid_authorization` on enum `IssuingDisputeEvidenceReason`
* Add support for `Routing` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, and `PaymentIntentPaymentMethodOptionsCardPresent`
* Add support for `ApplicationFeeAmount` and `ApplicationFee` on `Payout`
* Add support for `StripeS700` on `TerminalConfigurationParams` and `TerminalConfiguration`
