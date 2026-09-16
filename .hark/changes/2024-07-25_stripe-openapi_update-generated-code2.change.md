---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1892
is_stripe_api_change: true
released_in_version: 79.5.0
---

* Add support for `Update` method on resource `Checkout.Session`
* Add support for `TransactionID` on `ChargePaymentMethodDetailsAffirm`
* Add support for `BuyerID` on `ChargePaymentMethodDetailsBlik`
* Add support for `AuthorizationCode` on `ChargePaymentMethodDetailsCard`
* Add support for `BrandProduct` on `ChargePaymentMethodDetailsCardPresent`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `ConfirmationTokenPaymentMethodPreviewCardPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, and `PaymentMethodCardPresent`
* Add support for `NetworkTransactionID` on `ChargePaymentMethodDetailsCardPresent`, `ChargePaymentMethodDetailsInteracPresent`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, and `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`
* Add support for `CaseType` on `DisputePaymentMethodDetailsCard`
* Add support for new values `invoice.overdue` and `invoice.will_be_due` on enum `EventType`
* Add support for `TWINT` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
