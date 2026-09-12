---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2245
is_stripe_api_change: true
released_in_version: 84.2.0-alpha.2
---

* Add support for `TrackingDetails` on `V2MoneyManagementOutboundPayment`
* Add support for `PaperCheck` on `V2MoneyManagementOutboundPaymentDeliveryOptionsParams` and `V2MoneyManagementOutboundPaymentDeliveryOptions`
* Add support for event notification `V2CoreAccountIncludingFutureRequirementsUpdatedEvent` with related object `V2CoreAccount`
* Add support for error code `account_rate_limit_exceeded` on `RateLimitError`
