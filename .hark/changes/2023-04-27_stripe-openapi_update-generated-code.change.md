---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1645
is_stripe_api_change: true
released_in_version: 74.17.0-beta.1
---

* Add support for `BillingCycleAnchor` and `ProrationBehavior` on `CheckoutSessionSubscriptionDataParams`
* Add support for `TerminalID` on `IssuingAuthorizationMerchantData` and `IssuingTransactionMerchantData`
* Add support for `Metadata` on `PaymentIntentCaptureParams`
* Add support for `Checks` on `SetupAttemptPaymentMethodDetailsCard`
* Add support for `TaxBreakdown` on `TaxCalculationShippingCost` and `TaxTransactionShippingCost`
* Change type of `TaxRegistrationActiveFromParams` and `TaxRegistrationExpiresAtParams` from `longInteger` to `longInteger | literal('now')`
