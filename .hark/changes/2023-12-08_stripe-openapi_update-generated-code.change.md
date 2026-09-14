---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1776
is_stripe_api_change: true
released_in_version: 76.9.0-beta.1
---

* Add support for `Get` method on resource `FinancialConnections.Transaction`
* Remove support for `IssuingCard` and `IssuingCardsList` on `AccountSessionComponentsParams`
* Add support for `PaymentMethodRemove`, `PaymentMethodSave`, and `PaymentMethodSetAsDefault` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`
* Remove support for `PaymentMethodDetach` and `PaymentMethodSetAsCustomerDefault` on `CustomerSessionComponentsPaymentElementFeaturesParams` and `CustomerSessionComponentsPaymentElementFeatures`
