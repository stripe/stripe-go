---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2351
is_stripe_api_change: true
released_in_version: 85.2.0-alpha.4
---

* Add support for new resource `PaymentLocation`
* Add support for `Del`, `Get`, `New`, and `Update` methods on resource `PaymentLocation`
* Add support for `Protections` on `AccountCapabilitiesCardPaymentsParams` and `Capability`
* Add support for `GiftCard` on `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentPaymentMethodDataParams`, and `SharedPaymentGrantedTokenPaymentMethodDetails`
* Add support for new value `gift_card` on enums `ConfirmationTokenPaymentMethodPreview.Type`, `PaymentMethod.Type`, and `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
* Add support for `Metadata` on `DelegatedCheckoutRequestedSessionConfirmParams`
* Add support for `CreditedItems` on `InvoiceItemProrationDetails`
* Add support for `NetworkLifecycle` on `IssuingDispute`
* Add support for new value `gift_card` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `StatusDetails` on `Subscription`
