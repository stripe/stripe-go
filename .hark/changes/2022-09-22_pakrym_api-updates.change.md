---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1547
is_breaking: true
is_stripe_api_change: true
released_in_version: 73.10.0
---

* Add support for `TermsOfService` on `CheckoutSessionConsentCollectionParams`, `CheckoutSessionConsentCollection`, `CheckoutSessionConsent`, `PaymentLinkConsentCollectionParams`, and `PaymentLinkConsentCollection`
* ⚠️ Remove support for `Plan` on `CheckoutSessionPaymentMethodOptionsCardInstallmentsParams`. The property was mistakenly released and never worked.
* Add support for `StatementDescriptor` on `PaymentIntentIncrementAuthorizationParams`
