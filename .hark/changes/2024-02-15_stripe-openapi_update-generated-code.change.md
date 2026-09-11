---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1812
is_stripe_api_change: true
released_in_version: 76.17.0
---

* Add support for `Networks` on `Card`, `PaymentMethodCardParams`, and `TokenCardParams`
* Add support for new value `no_voec` on enums `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, `TaxIdType`, and `TaxTransactionCustomerDetailsTaxIdsType`
* Add support for new value `financial_connections.account.refreshed_ownership` on enum `EventType`
* Add support for `DisplayBrand` on `PaymentMethodCard`
