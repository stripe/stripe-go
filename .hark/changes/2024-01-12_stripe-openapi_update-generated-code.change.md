---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1796
is_stripe_api_change: true
released_in_version: 76.13.0
---

* Add support for new resource `CustomerSession`
* Add support for `New` method on resource `CustomerSession`
* Remove support for values `obligation_inbound`, `obligation_payout_failure`, `obligation_payout`, and `obligation_reversal_outbound` from enum `BalanceTransactionType`
* Remove support for `Expand` on `BankAccountParams` and `CardParams`
* Add support for `AccountType`, `DefaultForCurrency`, and `Documents` on `BankAccountParams` and `CardParams`
* Remove support for `Owner` on `BankAccountParams` and `CardParams`
* Change type of `BankAccountAccountHolderTypeParams` and `CardAccountHolderTypeParams` from `enum('company'|'individual')` to `emptyStringable(enum('company'|'individual'))`
* Add support for new values `eps` and `p24` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Add support for `BillingCycleAnchorConfig` on `SubscriptionParams` and `Subscription`
