---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1775
is_stripe_api_change: true
released_in_version: 76.8.0
---

* Add support for `PaymentDetails`, `Payments`, and `Payouts` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for `Features` on `AccountSessionComponentsAccountOnboardingParams` and `AccountSessionComponentsAccountOnboarding`
* Add support for new values `customer_tax_location_invalid` and `financial_connections_no_successful_transaction_refresh` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for new values `payment_network_reserve_hold` and `payment_network_reserve_release` on enum `BalanceTransactionType`
* Remove support for value `various` from enum `ClimateSupplierRemovalPathway`
* Remove support for values `challenge_only` and `challenge` from enum `PaymentIntentPaymentMethodOptionsCardRequestThreeDSecure`
* Add support for `InactiveMessage` and `Restrictions` on `PaymentLinkParams` and `PaymentLink`
* Add support for `TransferGroup` on `PaymentLinkPaymentIntentDataParams` and `PaymentLinkPaymentIntentData`
* Add support for `TrialSettings` on `PaymentLinkSubscriptionDataParams` and `PaymentLinkSubscriptionData`
