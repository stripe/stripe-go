---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1528
is_stripe_api_change: true
released_in_version: 73.3.0
---

* Add support for new resource `CustomerCashBalanceTransaction`
* Remove support for value `paypal` from enum `OrderPaymentSettingsPaymentMethodTypes`
* Add support for `Currency` on `PaymentLink`
* Add support for `Network` on `SetupIntentConfirmPaymentMethodOptionsCardParams`, `SetupIntentPaymentMethodOptionsCardParams`, `SubscriptionPaymentSettingsPaymentMethodOptionsCardParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsCard`
* Change type of `TopupSource` from `$Source` to `nullable($Source)`
