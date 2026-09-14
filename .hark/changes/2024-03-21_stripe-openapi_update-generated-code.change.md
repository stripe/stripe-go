---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1828
is_stripe_api_change: true
released_in_version: 76.22.0
---

* Add support for new resources `ConfirmationToken` and `Forwarding.Request`
* Add support for `Get` method on resource `ConfirmationToken`
* Add support for `Get`, `List`, and `New` methods on resource `Request`
* Add support for `MobilepayPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for new values `forwarding_api_inactive`, `forwarding_api_invalid_parameter`, `forwarding_api_upstream_connection_error`, and `forwarding_api_upstream_connection_timeout` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `Mobilepay` on `ChargePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `PaymentReference` on `ChargePaymentMethodDetailsUsBankAccount`
* Add support for `ConfirmationToken` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, `SetupIntentConfirmParams`, and `SetupIntentParams`
* Add support for new value `mobilepay` on enum `PaymentMethodType`
* Add support for `Name` on `TerminalConfigurationParams` and `TerminalConfiguration`
* Add support for `Payout` on `TreasuryReceivedDebitLinkedFlows`
