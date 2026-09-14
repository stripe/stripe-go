---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1847
is_stripe_api_change: true
released_in_version: 78.2.0
---

* Add support for new resource `Entitlements.ActiveEntitlementSummary`
* Add support for `Balances` and `PayoutsList` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for new value `entitlements.active_entitlement_summary.updated` on enum `EventType`
* Remove support for `Config` on `ForwardingRequestParams` and `ForwardingRequest`. This field is no longer used by the Forwarding Request API.
* Add support for `CaptureMethod` on `PaymentIntentConfirmPaymentMethodOptionsRevolutPayParams`, `PaymentIntentPaymentMethodOptionsRevolutPayParams`, and `PaymentIntentPaymentMethodOptionsRevolutPay`
* Add support for `Swish` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
