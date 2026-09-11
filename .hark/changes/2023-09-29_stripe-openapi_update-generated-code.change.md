---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1740
is_stripe_api_change: true
released_in_version: 75.9.0-beta.1
---

* Rename resources `Issuing.CardDesign` and `Issuing.CardBundle` to `Issuing.PersonalizationDesign` and `Issuing.PhysicalBundle`
* Add support for `Features` on `AccountSessionComponentsAccountOnboardingParams`, `AccountSessionComponentsPaymentDetailsParams`, `AccountSessionComponentsPaymentDetails`, `AccountSessionComponentsPaymentsParams`, `AccountSessionComponentsPayments`, `AccountSessionComponentsPayoutsParams`, and `AccountSessionComponentsPayouts`
* Add support for `Reason` on `Event`
