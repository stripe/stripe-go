---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1798
is_stripe_api_change: true
released_in_version: 76.14.0
---

* Add support for new value `nn` on enums `ChargePaymentMethodDetailsIdealBank`, `PaymentMethodIdealBank`, and `SetupAttemptPaymentMethodDetailsIdealBank`
* Add support for `Issuer` on `InvoiceParams`, `InvoiceUpcomingLinesParams`, `InvoiceUpcomingParams`, and `Invoice`
* Add support for `Liability` on `InvoiceAutomaticTaxParams`, `InvoiceAutomaticTax`, `InvoiceUpcomingAutomaticTaxParams`, `InvoiceUpcomingLinesAutomaticTaxParams`, `SubscriptionAutomaticTaxParams`, and `SubscriptionAutomaticTax`
* Add support for `OnBehalfOf` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
* Add support for `PIN` on `IssuingCardParams`
* Add support for `RevocationReason` on `MandatePaymentMethodDetailsBacsDebit`
* Add support for `CustomerBalance` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
* Add support for `InvoiceSettings` on `SubscriptionParams`
