---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/2063
is_stripe_api_change: true
released_in_version: 82.2.0
---

* Add support for `AttachPayment` method on resource `Invoice`
* Add support for `CollectInputs` method on resource `TerminalReader`
* Add support for `SucceedInputCollection` and `TimeoutInputCollection` test helper methods on resource `TerminalReader`
* Add support for `PixPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `DisputesList` and `PaymentDisputes` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for `RefundAndDisputePrefunding` on `Balance`
* Add support for `BalanceType` on `BalanceTransaction`
* Add support for `Location` and `Reader` on `ChargePaymentMethodDetailsAffirm` and `ChargePaymentMethodDetailsWechatPay`
* Add support for `PaymentMethodRemove` on `CheckoutSessionSavedPaymentMethodOptionsParams`
* Add support for `SetupFutureUsage` on `CheckoutSessionPaymentMethodOptionsNaverPay`
* Add support for `PostPaymentAmount` and `PrePaymentAmount` on `CreditNote`
* Add support for new value `mixed` on enum `CreditNote.Type`
* Add support for new value `invoice_payment.paid` on enum `Event.Type`
* Add support for `Sex`, `UnparsedPlaceOfBirth`, and `UnparsedSex` on `IdentityVerificationReportDocument` and `IdentityVerificationSessionVerifiedOutputs`
* Add support for `BillingThresholds` on `InvoiceCreatePreviewScheduleDetailsPhaseItemParams`, `InvoiceCreatePreviewScheduleDetailsPhaseParams`, `InvoiceCreatePreviewSubscriptionDetailsItemParams`, `SubscriptionItemParams`, `SubscriptionItem`, `SubscriptionParams`, `SubscriptionScheduleDefaultSettingsParams`, `SubscriptionScheduleDefaultSettings`, `SubscriptionSchedulePhaseItemParams`, `SubscriptionSchedulePhaseItem`, `SubscriptionSchedulePhaseParams`, `SubscriptionSchedulePhase`, and `Subscription`
* Add support for `Satispay` on `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptionsParams`, and `PaymentIntentPaymentMethodOptions`
* Add support for `CaptureMethod` on `PaymentIntentPaymentMethodOptionsBillie`
* Add support for `KakaoPay`, `KrCard`, `NaverPay`, `Payco`, and `SamsungPay` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
* Add support for `NetworkDeclineCode` on `RefundDestinationDetailsPaypal`
* Add support for `Metadata` on `TaxCalculationLineItemParams` and `TaxCalculationLineItem`
* Add support for `ReturnURL` on `TerminalReaderActionProcessPaymentIntentProcessConfig` and `TerminalReaderProcessPaymentIntentProcessConfigParams`
* Add support for `CollectInputs` on `TerminalReaderAction`
* Add support for new value `collect_inputs` on enum `TerminalReaderAction.Type`
* Add support for new value `simulated_stripe_s700` on enum `TerminalReader.DeviceType`
* Add support for snapshot event `EventTypeInvoicePaymentPaid` with resource `InvoicePayment`
* Add support for error code `forwarding_api_upstream_error` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
