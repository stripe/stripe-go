---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1605
is_stripe_api_change: true
released_in_version: 74.8.0
---

* Add support for `RefundPayment` method on resource `Terminal.Reader`
* Add support for new value `name` on enum `BillingPortalConfigurationFeaturesCustomerUpdateAllowedUpdates`
* Add support for `CustomFields` on `CheckoutSessionParams`, `CheckoutSession`, `PaymentLinkParams`, and `PaymentLink`
* Add support for `InteracPresent` on `TestHelpersTerminalReaderPresentPaymentMethodParams`
* Change type of `TestHelpersTerminalReaderPresentPaymentMethodTypeParams` from `literal('card_present')` to `enum('card_present'|'interac_present')`
* Add support for `RefundPayment` on `TerminalReaderAction`
* Add support for new value `refund_payment` on enum `TerminalReaderActionType`
