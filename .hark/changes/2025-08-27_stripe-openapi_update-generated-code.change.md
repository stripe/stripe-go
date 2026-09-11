---
title: Update generated code.
pr_link: https://github.com/stripe/stripe-go/pull/2101
is_stripe_api_change: true
released_in_version: 82.5.0
---

* Add support for `BalanceReport`, `PayoutDetails`, and `PayoutReconciliationReport` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for `Name` on `BillingPortalConfigurationParams` and `BillingPortalConfiguration`
* Add support for `Installments` on `ChargePaymentMethodDetailsAlma`
* Add support for `TransactionID` on `ChargePaymentMethodDetailsAlma`, `ChargePaymentMethodDetailsAmazonPay`, `ChargePaymentMethodDetailsBillie`, `ChargePaymentMethodDetailsKakaoPay`, `ChargePaymentMethodDetailsKrCard`, `ChargePaymentMethodDetailsNaverPay`, `ChargePaymentMethodDetailsPayco`, `ChargePaymentMethodDetailsRevolutPay`, `ChargePaymentMethodDetailsSamsungPay`, and `ChargePaymentMethodDetailsSatispay`
* Add support for `Location` and `Reader` on `ChargePaymentMethodDetailsPaynow`
* Add support for `AmountIncludesIof` on `CheckoutSessionPaymentMethodOptionsPixParams`, `CheckoutSessionPaymentMethodOptionsPix`, `PaymentIntentConfirmPaymentMethodOptionsPixParams`, `PaymentIntentPaymentMethodOptionsPixParams`, and `PaymentIntentPaymentMethodOptionsPix`
* Add support for new values `block` and `resolution` on enum `DisputePaymentMethodDetailsCard.CaseType`
* Add support for new value `terminal_android_apk` on enum `File.Purpose`
* Add support for `Metadata` and `Period` on `InvoiceCreatePreviewScheduleDetailsPhaseAddInvoiceItemParams`, `SubscriptionAddInvoiceItemParams`, `SubscriptionSchedulePhaseAddInvoiceItemParams`, and `SubscriptionSchedulePhaseAddInvoiceItem`
* Add support for `ExpMonth` and `ExpYear` on `IssuingCardParams`
* Add support for `ExcludedPaymentMethodTypes` on `PaymentIntentParams` and `PaymentIntent`
* Add support for `PayoutMethod` on `PayoutParams` and `Payout`
* Add support for `Mxn` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
* Add support for `Card` on `TestHelpersTerminalReaderPresentPaymentMethodParams`
* Add support for error codes `customer_session_expired` and `india_recurring_payment_mandate_canceled` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
