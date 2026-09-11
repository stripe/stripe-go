---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1816
is_stripe_api_change: true
released_in_version: 76.20.0-beta.1
---

* Remove support for resource `Entitlements.Event`
* Change type of `ConfirmationTokenMandateData` from `nullable(ConfirmationTokensResourceMandateData)` to `ConfirmationTokensResourceMandateData`
* Remove support for `Quantity` and `Type` on `EntitlementsFeatureParams` and `EntitlementsFeature`
* Add support for `Livemode` on `IssuingPersonalizationDesign`
* Add support for `ApplicationFeeAmount`, `Description`, `Metadata`, and `TransferData` on `PaymentIntentDecrementAuthorizationParams`
* Add support for `EnableCustomerCancellation` on `TerminalReaderActionCollectPaymentMethodCollectConfig` and `TerminalReaderCollectPaymentMethodCollectConfigParams`
