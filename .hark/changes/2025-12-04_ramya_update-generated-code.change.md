---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2230
is_stripe_api_change: true
released_in_version: 84.1.0-alpha.4
---

* Add support for `CheckScanning` on `AccountSessionComponents`
* Add support for `Client` on `V2CoreEventReasonRequest`
* Add support for `StripeBalancePayment` on `V2MoneyManagementReceivedCredit` and `V2MoneyManagementReceivedDebit`
* Add support for new value `stripe_balance_payment` on enum `V2MoneyManagementReceivedCredit.Type`
* Add support for `BalanceTransfer` on `V2MoneyManagementReceivedDebit`
* Add support for new values `balance_transfer` and `stripe_balance_payment` on enum `V2MoneyManagementReceivedDebit.Type`
* Add support for `Include` on `V2CoreEventListParams` and `V2CoreEventParams`
