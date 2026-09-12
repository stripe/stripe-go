---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1746
is_stripe_api_change: true
released_in_version: 75.10.0
---

* Add support for `RedirectOnCompletion`, `ReturnURL`, and `UIMode` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `ClientSecret` on `CheckoutSession`
* Change type of `CheckoutSessionCustomFieldsDropdown` from `nullable(PaymentPagesCheckoutSessionCustomFieldsDropdown)` to `PaymentPagesCheckoutSessionCustomFieldsDropdown`
* Change type of `CheckoutSessionCustomFieldsNumeric` and `CheckoutSessionCustomFieldsText` from `nullable(PaymentPagesCheckoutSessionCustomFieldsNumeric)` to `PaymentPagesCheckoutSessionCustomFieldsNumeric`
* Add support for `PostalCode` on `IssuingAuthorizationVerificationData`
* Change type of `PaymentLinkCustomFieldsDropdown` from `nullable(PaymentLinksResourceCustomFieldsDropdown)` to `PaymentLinksResourceCustomFieldsDropdown`
* Change type of `PaymentLinkCustomFieldsNumeric` and `PaymentLinkCustomFieldsText` from `nullable(PaymentLinksResourceCustomFieldsNumeric)` to `PaymentLinksResourceCustomFieldsNumeric`
* Add support for `Offline` on `TerminalConfigurationParams` and `TerminalConfiguration`
