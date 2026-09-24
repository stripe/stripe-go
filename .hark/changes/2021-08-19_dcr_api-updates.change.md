---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1328
is_stripe_api_change: true
released_in_version: 72.61.0
---

* Add support for new TaxId type: `au_arn`
* Add support for `InteracPresent` on `ChargePaymentMethodDetails`
* Add support for `SepaCreditTransfer` on `ChargePaymentMethodDetails`
* Codegen related changes:
  * Moved `ShippingDetails` into `address.go`
  * Add support for `Object` and `Order` to `Charge`
  * Renamed `ReviewReasonType` enum to `ReviewReason` but added a type alias to preserve backwards compatibility
