---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2034
is_stripe_api_change: true
released_in_version: 82.2.0-beta.1
---

This release changes the pinned API version to `2025-04-30.preview`.

* Add support for `BillingMode` on `CheckoutSessionSubscriptionDataParams`, `InvoiceCreatePreviewScheduleDetailsParams`, `InvoiceCreatePreviewSubscriptionDetailsParams`, `QuotePreviewSubscriptionSchedule`, `QuoteSubscriptionDataParams`, `QuoteSubscriptionData`, `SubscriptionParams`, `SubscriptionScheduleParams`, `SubscriptionSchedule`, and `Subscription`
* Add support for new values `aw_tin`, `az_tin`, `bd_bin`, `bf_ifu`, `bj_ifu`, `cm_niu`, `cv_nif`, `et_tin`, `kg_tin`, and `la_tin` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, and `QuotePreviewInvoiceCustomerTaxIds.Type`
* Add support for `AccountNumber` on `ConfirmationTokenPaymentMethodPreviewAcssDebit` and `PaymentMethodAcssDebit`
* Add support for new value `balance_settings.updated` on enum `Event.Type`
