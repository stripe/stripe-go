---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2250
is_stripe_api_change: true
released_in_version: 84.2.0-alpha.3
---

* Add support for `RiskDetails` on `DelegatedCheckoutRequestedSession`
* Remove support for `Description`, `Images`, and `Name` on `DelegatedCheckoutRequestedSessionLineItemDetail`
* Add support for `Name` on `ProductCatalogTrialOfferParams` and `ProductCatalogTrialOffer`
* Add support for `LoginFailed` and `RegistrationFailed` on `RadarAccountEvaluationEvents` and `RadarAccountEvaluationParams`
* Change type of `RadarAccountEvaluationParams.Type` from `literal('registration_succeeded')` to `enum('login_failed'|'login_succeeded'|'registration_failed'|'registration_succeeded')`
