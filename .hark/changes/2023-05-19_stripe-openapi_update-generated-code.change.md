---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1657
is_stripe_api_change: true
released_in_version: 74.19.0
---

* Add support for `SubscriptionUpdateConfirm` and `SubscriptionUpdate` on `BillingPortalSessionFlowDataParams` and `BillingPortalSessionFlow`
* Add support for new values `subscription_update_confirm` and `subscription_update` on enum `BillingPortalSessionFlowType`
* Add support for `Link` on `ChargePaymentMethodDetailsCardWallet` and `PaymentMethodCardWallet`
* Add support for `BuyerID` and `Cashtag` on `ChargePaymentMethodDetailsCashapp` and `PaymentMethodCashapp`
* Add support for new values `amusement_tax` and `communications_tax` on enum `TaxRateTaxType`
