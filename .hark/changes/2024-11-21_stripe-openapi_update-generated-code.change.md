---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1952
is_stripe_api_change: true
released_in_version: 81.2.0-beta.1
---

* Add support for `NetworkAdviceCode` and `NetworkDeclineCode` on `ChargeOutcome`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
* Add support for `Funding` on `ChargePaymentMethodDetailsAmazonPay` and `ChargePaymentMethodDetailsRevolutPay`
* Add support for `AmountRequested` and `PartialAuthorization` on `ChargePaymentMethodDetailsCard`
* Add support for `Metadata` on `CheckoutSessionLineItemsParams` and `LineItem`
* Add support for `LineItems` on `CheckoutSessionParams`, `CheckoutSessionPermissionsUpdateParams`, and `CheckoutSessionPermissionsUpdate`
* Add support for new value `invoice.overpaid` on enum `EventType`
* Add support for `AdjustableQuantity` and `Display` on `LineItem`
* Add support for `RequestPartialAuthorization` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, and `PaymentIntentPaymentMethodOptionsCard`
* Add support for `PaymentMethodOptions` on `PaymentIntentIncrementAuthorizationParams`
