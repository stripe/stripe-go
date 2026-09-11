---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/2132
is_stripe_api_change: true
released_in_version: 83.1.0-beta.1
---

* Add support for new value `billing_cadence_details` on enums `InvoiceParent.Type` and `QuotePreviewInvoiceParent.Type`
* Add support for `AttachCadence` method on resource `Subscription`
* Add support for `BillingCadenceDetails` on `InvoiceParent` and `QuotePreviewInvoiceParent`
* Add support for `BillingCadence` on `InvoiceCreatePreviewParams`, `SubscriptionParams`, and `Subscription`
