---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1690
is_stripe_api_change: true
released_in_version: 74.29.0-beta.1
---

* Remove support for values `excluded_territory`, `jurisdiction_unsupported`, and `vat_exempt` from enums `OrderShippingCostTaxesTaxabilityReason`, `OrderTotalDetailsBreakdownTaxesTaxabilityReason`, and `QuotePhaseTotalDetailsBreakdownTaxesTaxabilityReason`
* Add support for new value `ro_tin` on enum `OrderTaxDetailsTaxIdsType`
* Add support for new values `email`, `numeric`, `phone`, and `text` on enum `TerminalReaderActionCollectInputsInputsType`
