---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1813
is_stripe_api_change: true
released_in_version: 76.18.0-beta.1
---

* Add support for `DecrementAuthorization` method on resource `PaymentIntent`
* Add support for `PaytoPayments` and `TWINTPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `Payto` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodPreview`, `MandatePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodDataParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for `TWINT` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `DecrementalAuthorization` on `ChargePaymentMethodDetailsCard`
* Add support for `DisplayBrand` on `ConfirmationTokenPaymentMethodPreviewCard`
* Add support for new values `payto` and `twint` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
* Add support for new value `no_voec` on enum `OrderTaxDetailsTaxIdsType`
* Add support for `RequestDecrementalAuthorization` on `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, and `PaymentIntentPaymentMethodOptionsCard`
