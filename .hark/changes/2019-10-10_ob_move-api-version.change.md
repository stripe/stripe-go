---
title: Move to API version [`2019-10-08`](https://docs.stripe.com/changelog/2019-10-08) and other changes
pr_url: https://github.com/stripe/stripe-go/pull/951
released_in_version: 65.0.0
---

* [#950](https://github.com/stripe/stripe-go/pull/950) Remove lossy "MarshalJSON" implementations
* [#962](https://github.com/stripe/stripe-go/pull/962) Removed deprecated properties and most todos
  * Removed `GetBalanceTransaction` and `List` from the `balance` package. Prefer using `Get` and `List` in the `balancetransaction` package.
  * Removed `ApplicationFee` from the `charge` and `paymentintent` packages. Prefer using `ApplicationFeeAmount`.
  * Removed `TaxInfo` and related fields from the `customer` packager. Prefer using the `customertaxid` package.
  * Removed unsupported `Customer` parameter on `PaymentMethodParams` and `PaymentMethodDetachParams` in the `paymentmethod` package.
  * Removed `Billing` properties in the `invoice`, `sub` and `subschedule` packages. Prefer using `CollectionMethod`.
  * Removed the `InvoiceBilling` type from the `invoice` package. Prefer using `InvoiceCollectionMethod`.
  * Removed the `SubscriptionBilling` type from the `sub` package. Prefer using `SubscriptionCollectionMethod`.
  * Removed deprecated constants for `PaymentIntentConfirmationMethod` in `paymentintent` package.
  * Removed `OperatorAccount` from Terminal APIs.
* [#960](https://github.com/stripe/stripe-go/pull/960) Remove `issuerfraudrecord` package. Prefer using `earlyfraudwarning`
* [#968](https://github.com/stripe/stripe-go/pull/968) Rename `AccountOpener` to `Representative` and update to latest API version
