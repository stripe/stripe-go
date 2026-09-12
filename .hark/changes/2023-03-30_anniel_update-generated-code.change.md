---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1635
is_stripe_api_change: true
released_in_version: 74.14.0
---

* Remove support for `New` method on resource `Tax.Transaction`
  * This is not a breaking change, as this method was deprecated before the Tax Transactions API was released in favor of the `CreateFromCalculation` method.
* Add support for `ExportLicenseID` and `ExportPurposeCode` on `AccountCompanyParams`, `AccountCompany`, and `TokenAccountCompanyParams`
* Remove support for value `deleted` from enum `InvoiceStatus`
  * This is not a breaking change, as the value was never returned or accepted as input.
* Add support for `AmountTip` on `TestHelpersTerminalReaderPresentPaymentMethodParams`
