---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1920
is_stripe_api_change: true
released_in_version: 79.13.0-beta.1
---

* Remove support for resource `QuotePhase`
* Remove support for `Get` and `ListLineItems` methods on resource `QuotePhase`
* Add support for `SendMoney` and `TransferBalance` on `AccountSessionComponentsFinancialAccountFeaturesParams`
* Add support for new value `rechnung` on enum `PaymentLinkPaymentMethodTypes`
