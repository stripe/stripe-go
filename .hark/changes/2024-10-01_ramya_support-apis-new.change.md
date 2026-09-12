---
title: Support for APIs in the new API version 2024-09-30.acacia
pr_url: https://github.com/stripe/stripe-go/pull/1926
is_stripe_api_change: true
released_in_version: 80.0.0
---

This release changes the pinned API version to `2024-09-30.acacia`. Please read the [API Changelog](https://docs.stripe.com/changelog/acacia#2024-09-30.acacia) and carefully review the API changes before upgrading.

### ⚠️ Breaking changes

* Rename `usage_threshold_config` to `usage_threshold` on `BillingAlertParams` and `BillingAlert`
* Remove support for `filter` on `BillingAlertParams` and `BillingAlert`. Use the filters on the `usage_threshold` instead
* Remove support for `CustomerConsentCollected` on `TerminalReaderProcessSetupIntentParams`


### Additions
* Add support for `CustomUnitAmount` on `ProductDefaultPriceDataParams`
* Add support for `AllowRedisplay` on `TerminalReaderProcessPaymentIntentProcessConfigParams` and `TerminalReaderProcessSetupIntentParams`
* Add support for new value `international_transaction` on enum `TreasuryReceivedCreditFailureCode`
* Add method [RawRequest()](https://github.com/stripe/stripe-go/tree/master?tab=readme-ov-file#custom-requests) that takes a HTTP method type, url and relevant parameters to make requests to the Stripe API that are not yet supported in the SDK.
