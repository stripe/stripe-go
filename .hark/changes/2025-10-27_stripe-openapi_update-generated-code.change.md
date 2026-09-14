---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2147
is_stripe_api_change: true
released_in_version: 83.1.0
---

* Add support for new resources `PaymentAttemptRecord`, `PaymentIntentAmountDetailsLineItem`, and `PaymentRecord`
* Add support for `Get` and `List` methods on resource `PaymentAttemptRecord`
* Add support for `Get`, `ReportPaymentAttemptCanceled`, `ReportPaymentAttemptFailed`, `ReportPaymentAttemptGuaranteed`, `ReportPaymentAttemptInformational`, `ReportPaymentAttempt`, `ReportPayment`, and `ReportRefund` methods on resource `PaymentRecord`
* Add support for `List` method on resource `PaymentIntentAmountDetailsLineItem`
* Add support for `RepresentativeDeclaration` on `AccountCompanyParams`, `AccountCompany`, and `TokenAccountCompanyParams`
* Add support for `PaymentMethodConfiguration` on `BillingPortalConfigurationFeaturesPaymentMethodUpdateParams`
* Add support for new value `solana` on enum `ChargePaymentMethodDetailsCrypto.Network`
* Add support for `TWINT` on `CheckoutSessionPaymentMethodOptionsParams` and `CheckoutSessionPaymentMethodOptions`
* Add support for new value `custom` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `PaymentRecordRefund` and `Type` on `CreditNotePreviewLinesRefundParams`, `CreditNotePreviewRefundParams`, `CreditNoteRefundParams`, and `CreditNoteRefund`
* Add support for `CustomerSheet` and `MobilePaymentElement` on `CustomerSessionComponentsParams` and `CustomerSessionComponents`
* Add support for `Provider` on `CustomerTax`
* Add support for new values `balance_settings.updated` and `invoice.payment_attempt_required` on enum `Event.Type`
* Add support for new value `platform_terms_of_service` on enum `File.Purpose`
* Add support for `PaymentRecord` on `InvoiceAttachPaymentParams`, `InvoicePaymentListPaymentParams`, and `InvoicePaymentPayment`
* Change type of `InvoicePaymentListPaymentParams.Type` from `literal('payment_intent')` to `enum('payment_intent'|'payment_record')`
* Add support for new value `custom` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `AmountDetails` on `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentIncrementAuthorizationParams`, and `PaymentIntentParams`
* Add support for `PaymentDetails` on `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentIncrementAuthorizationParams`, `PaymentIntentParams`, and `PaymentIntent`
* Add support for `DiscountAmount`, `LineItems`, `Shipping`, and `Tax` on `PaymentIntentAmountDetails`
* Add support for `NameCollection` on `PaymentLinkParams` and `PaymentLink`
* Add support for new value `mb_way` on enum `PaymentLink.PaymentMethodTypes`
* Add support for `Crypto` on `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, and `RefundDestinationDetails`
* Add support for `MbWay` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
* Add support for `Custom` on `PaymentMethodParams` and `PaymentMethod`
* Add support for `ExcludedPaymentMethodTypes` on `SetupIntentParams` and `SetupIntent`
* Add support for `Tw` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
* Add support for `Gip` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
* Add support for `LastSeenAt` on `TerminalReader`
* Add support for `GTE`, `Gt`, `LT`, `Lte`, and `Types` on `V2CoreEventListParams`
* Add support for snapshot event `EventTypeBalanceSettingsUpdated` with resource `BalanceSettings`
* Add support for snapshot event `EventTypeInvoicePaymentAttemptRequired` with resource `Invoice`
* Add support for error code `payment_intent_rate_limit_exceeded` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
