---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1705
is_breaking: true
is_stripe_api_change: true
released_in_version: 75.0.0
---

* ⚠️Add support for new values `verification_directors_mismatch`, `verification_document_directors_mismatch`, `verification_extraneous_directors`, and `verification_missing_directors` on enums `AccountFutureRequirementsErrorsCode`, `AccountRequirementsErrorsCode`, `BankAccountFutureRequirementsErrorsCode`, and `BankAccountRequirementsErrorsCode`
* Remove support for `AvailableOn` on `BalanceTransactionListParams`
  * Use of this parameter is discouraged. You may use [`.AddExtra`](https://github.com/stripe/stripe-go#parameters) if sending the parameter is still required.
* ⚠️Remove support for `Destination` on `Charge`
  * Please use `TransferData` or `OnBehalfOf` instead.
* ⚠️Remove support for `AlternateStatementDescriptors` and `Dispute` on `Charge`
  * Use of these parameters is discouraged.
* ⚠️Remove support for `ShippingRates` on `CheckoutSessionParams`
  * Please use `ShippingParams` instead.
* ⚠️Remove support for `Coupon` and `TrialFromPlan` on `CheckoutSessionSubscriptionDataParams`
  * Please [migrate to the Prices API](https://stripe.com/docs/billing/migration/migrating-prices), or use [`.AddExtra`](https://github.com/stripe/stripe-go#parameters) if sending the parameter is still required.
* ⚠️Remove support for value `charge_refunded` from enum `DisputeStatus`
* ⚠️Remove support for `BLIK` on `MandatePaymentMethodDetails`, `PaymentMethodParams`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
    * These fields were mistakenly released.
* ⚠️Remove support for `ACSSDebit`, `AUBECSDebit`, `Affirm`, `BACSDebit`, `CashApp`, `SEPADebit`, and `Zip` on `PaymentMethodParams`
    * These fields were empty hashes.
* ⚠️Remove support for `Country` on `PaymentMethodLink`
    * This field was not fully operational.
* ⚠️Remove support for `Recurring` on `PriceParams`
    * This property should be set on create only.
* ⚠️Remove support for `Attributes`, `Caption`, and `DeactivateOn` on `ProductParams` and `Product`
  * These fields are not fully operational.
