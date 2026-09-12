---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1858
is_stripe_api_change: true
released_in_version: 78.6.0
---

* Add support for `Update` test helper method on resources `Treasury.OutboundPayment` and `Treasury.OutboundTransfer`
* Add support for `AllowRedisplay` on `ConfirmationTokenPaymentMethodPreview` and `PaymentMethod`
* Add support for new values `treasury.outbound_payment.tracking_details_updated` and `treasury.outbound_transfer.tracking_details_updated` on enum `EventType`
* Add support for `PreviewMode` on `InvoiceCreatePreviewParams`, `InvoiceUpcomingLinesParams`, and `InvoiceUpcomingParams`
* Add support for `TrackingDetails` on `TreasuryOutboundPayment` and `TreasuryOutboundTransfer`
