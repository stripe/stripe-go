---
title: Deprecate Go methods and Params
pr_link: https://github.com/stripe/stripe-go/pull/1856
released_in_version: 78.5.0
---

- Mark as deprecated the `Approve` and `Decline` methods on `issuing/authorization/client.go`.  Instead, [respond directly to the webhook request to approve an authorization](https://stripe.com/docs/issuing/controls/real-time-authorizations#authorization-handling).
- Mark as deprecated the `persistent_token` property on `ConfirmationTokenPaymentMethodPreviewLink.persistent_token`, `PaymentIntentPaymentMethodOptionsLink`, `PaymentIntentPaymentMethodOptionsLinkParams`, `PaymentMethodLink`, `SetupIntentPaymentMethodOptionsCard`, `SetupIntentPaymentMethodOptionsLinkParams`. This is a legacy parameter that no longer has any function.
