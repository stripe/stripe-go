---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1546
is_stripe_api_change: true
released_in_version: 73.9.0
---

* Add support for `Pix` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `FromInvoice` on `InvoiceParams` and `Invoice`
* Add support for `LatestRevision` on `Invoice`
* Add support for `Amount` on `IssuingDisputeParams`
* Add support for `PixDisplayQRCode` on `PaymentIntentNextAction`
* Add support for new value `pix` on enum `PaymentLinkPaymentMethodTypes`
* Add support for new value `pix` on enum `PaymentMethodType`
* Add support for `Created` on `TreasuryCreditReversal` and `TreasuryDebitReversal`
