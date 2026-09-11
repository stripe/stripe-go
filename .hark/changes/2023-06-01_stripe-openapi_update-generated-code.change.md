---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1664
is_stripe_api_change: true
released_in_version: 74.21.0
---

* Add support for `Numeric` and `Text` on `CheckoutSessionCustomFieldsParams` and `PaymentLinkCustomFieldsParams`
* Add support for `MaximumLength` and `MinimumLength` on `CheckoutSessionCustomFieldsNumeric` and `CheckoutSessionCustomFieldsText`
* Add support for new values `aba` and `swift` on enums `CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypes` and `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypes`
* Add support for new value `us_bank_transfer` on enums `CheckoutSessionPaymentMethodOptionsCustomerBalanceBankTransferType`, `PaymentIntentNextActionDisplayBankTransferInstructionsType`, and `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferType`
* Add support for `PreferredLocales` on `IssuingCardholderParams` and `IssuingCardholder`
* Add support for `Description`, `IIN`, and `Issuer` on `PaymentMethodCardPresent` and `PaymentMethodInteracPresent`
* Add support for `PayerEmail` on `PaymentMethodPaypal`
