---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1843
is_stripe_api_change: true
released_in_version: 78.2.0-beta.1
---

* Add support for `Get` method on resources `Entitlements.ActiveEntitlement` and `Entitlements.Feature`
* Add support for `Fees`, `Losses`, `RequirementCollection`, and `StripeDashboard` on `AccountControllerParams`
* Add support for new values `bh_vat`, `kz_bin`, `ng_tin`, and `om_vat` on enum `OrderTaxDetailsTaxIdsType`
* Add support for `HostedVoucherURL` on `PaymentIntentNextActionMultibancoDisplayDetails`
* Add support for `Toggles` on `TerminalReaderActionCollectInputsInputs` and `TerminalReaderCollectInputsInputsParams`
* Add support for `Email`, `Numeric`, `Phone`, and `Text` on `TerminalReaderActionCollectInputsInputs`
