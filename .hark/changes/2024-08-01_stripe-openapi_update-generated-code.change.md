---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1897
is_breaking: true
is_stripe_api_change: true
released_in_version: 79.6.0
---

* Add support for new resources `Billing.AlertTriggered` and `Billing.Alert`
* Add support for new value `charge_exceeds_transaction_limit` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* ⚠️ Remove support for `AuthorizationCode` on `ChargePaymentMethodDetailsCard`. This was accidentally released last week.
* Add support for new value `billing.alert.triggered` on enum `EventType`
