---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1761
is_stripe_api_change: true
released_in_version: 76.4.0-beta.1
---

* Add support for `AttachPaymentIntent` method on resource `Invoice`
* Add support for `RevolutPay` on `ConfirmationTokenPaymentMethodPreview`
* Add support for new value `revolut_pay` on enum `ConfirmationTokenPaymentMethodPreviewType`
* Add support for `Refunds` on `CreditNoteParams`, `CreditNotePreviewLinesParams`, `CreditNotePreviewParams`, and `CreditNote`
* Add support for `PostPaymentAmount` and `PrePaymentAmount` on `CreditNote`
* Add support for new value `invoice.payment.overpaid` on enum `EventType`
* Add support for `ScheduleDetails` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
* Add support for `AmountsDue` on `InvoiceParams` and `Invoice`
* Add support for `Payments` on `Invoice`
* Add support for `Created` on `IssuingPersonalizationDesign`
