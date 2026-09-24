---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1551
is_stripe_api_change: true
released_in_version: 73.12.0
---

* Add support for new value `invalid_dob_age_under_18` on enums `AccountFutureRequirementsErrorsCode`, `AccountRequirementsErrorsCode`, `CapabilityFutureRequirementsErrorsCode`, `CapabilityRequirementsErrorsCode`, `PersonFutureRequirementsErrorsCode`, and `PersonRequirementsErrorsCode`
* Add support for new value `bank_of_china` on enums `ChargePaymentMethodDetailsFpxBank` and `PaymentMethodFpxBank`
* Add support for `Klarna` on `SetupAttemptPaymentMethodDetails`
