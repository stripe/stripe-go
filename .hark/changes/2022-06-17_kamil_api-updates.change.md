---
title: API Updates
pr_link: https://github.com/stripe/stripe-go/pull/1477
is_stripe_api_change: true
released_in_version: 72.115.0
---

* Add support for `FundCashBalance` test helper method on resource `Customer`
* Add support for `StatementDescriptorPrefixKana` and `StatementDescriptorPrefixKanji` on `AccountSettingsCardPaymentsParams`, `AccountSettingsCardPayments`, and `AccountSettingsPayments`
* Add support for `StatementDescriptorSuffixKana` and `StatementDescriptorSuffixKanji` on `CheckoutSessionPaymentMethodOptionsCardParams`, `CheckoutSessionPaymentMethodOptionsCard`, `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, and `PaymentIntentPaymentMethodOptionsCard`
* Add support for `TotalExcludingTax` on `CreditNote`
* Change type of `CustomerInvoiceSettingsRenderingOptionsParams` from `rendering_options_param` to `emptyStringable(rendering_options_param)`
* Add support for `RenderingOptions` on `CustomerInvoiceSettings` and `Invoice`
