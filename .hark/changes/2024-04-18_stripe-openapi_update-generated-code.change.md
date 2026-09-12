---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1849
is_stripe_api_change: true
released_in_version: 78.3.0
---

* Add support for `CreatePreview` method on resource `Invoice`
* Add support for `PaymentMethodData` on `CheckoutSessionParams`
* Add support for `SavedPaymentMethodOptions` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `Mobilepay` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for `AllowRedisplay` on `ConfirmationTokenPaymentMethodDataParams`, `CustomerListPaymentMethodsParams`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodParams`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `ScheduleDetails` and `SubscriptionDetails` on `InvoiceUpcomingLinesParams` and `InvoiceUpcomingParams`
