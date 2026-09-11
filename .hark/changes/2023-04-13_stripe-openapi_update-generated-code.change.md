---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1639
is_stripe_api_change: true
released_in_version: 74.16.0-beta.2
---

* Add support for `CollectPaymentMethod` and `ConfirmPaymentIntent` methods on resource `Terminal.Reader`
* Add support for `PaypalPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for new value `REVOIE23` on enums `ChargePaymentMethodDetailsIdealBic`, `PaymentMethodIdealBic`, and `SetupAttemptPaymentMethodDetailsIdealBic`
* Add support for `CollectPaymentMethod` and `ConfirmPaymentIntent` on `TerminalReaderAction`
* Add support for `StripeAccount` on `TerminalReaderActionProcessPaymentIntent` and `TerminalReaderActionRefundPayment`
* Add support for new values `collect_payment_method` and `confirm_payment_intent` on enum `TerminalReaderActionType`
