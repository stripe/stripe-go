---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1853
is_stripe_api_change: true
released_in_version: 78.5.0
---

* Add support for new value `shipping_address_invalid` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `Paypal` on `DisputePaymentMethodDetails`
* Change type of `DisputePaymentMethodDetailsType` from `literal('card')` to `enum('card'|'paypal')`
* Change type of `EntitlementsFeatureMetadataParams` from `map(string: string)` to `emptyable(map(string: string))`
* Add support for `PaymentMethodTypes` on `PaymentIntentConfirmParams`
* Add support for `ShipFromDetails` on `TaxCalculationParams`, `TaxCalculation`, and `TaxTransaction`
* Add support for `Bh`, `Eg`, `Ge`, `Ke`, `Kz`, `Ng`, and `Om` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
