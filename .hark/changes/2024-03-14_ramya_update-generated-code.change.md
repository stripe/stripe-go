---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1824
is_stripe_api_change: true
released_in_version: 76.21.0
---

* Add support for new resources `Issuing.PersonalizationDesign` and `Issuing.PhysicalBundle`
* Add support for `Get`, `List`, `New`, and `Update` methods on resource `PersonalizationDesign`
* Add support for `Get` and `List` methods on resource `PhysicalBundle`
* Add support for `PersonalizationDesign` on `IssuingCardListParams`, `IssuingCardParams`, and `IssuingCard`
* Change type of `SubscriptionApplicationFeePercentParams` from `number` to `emptyStringable(number)`
* Add support for `SEPADebit` on `SubscriptionPaymentSettingsPaymentMethodOptionsParams` and `SubscriptionPaymentSettingsPaymentMethodOptions`
