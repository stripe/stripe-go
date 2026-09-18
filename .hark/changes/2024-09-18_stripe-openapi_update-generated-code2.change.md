---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1918
is_stripe_api_change: true
released_in_version: 79.12.0
---

* Add support for new value `verification_supportability` on enums `AccountFutureRequirementsErrorsCode`, `AccountRequirementsErrorsCode`, `BankAccountFutureRequirementsErrorsCode`, and `BankAccountRequirementsErrorsCode`
* Add support for new value `terminal_reader_invalid_location_for_activation` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `PayerDetails` on `ChargePaymentMethodDetailsKlarna`
* Add support for `AmazonPay` on `DisputePaymentMethodDetails`
* Add support for new value `amazon_pay` on enum `DisputePaymentMethodDetailsType`
* Add support for `AutomaticallyFinalizesAt` on `Invoice`
* Add support for `StateSalesTax` on `TaxRegistrationCountryOptionsUsParams` and `TaxRegistrationCountryOptionsUs`
