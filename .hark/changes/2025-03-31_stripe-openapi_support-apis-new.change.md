---
title: Support for APIs in the new API version 2025-03-31.basil
pr_link: https://github.com/stripe/stripe-go/pull/1992
is_stripe_api_change: true
released_in_version: 82.0.0
---

This release changes the pinned API version to `2025-03-31.basil`.

### ⚠️ Breaking changes due to changes in the Stripe API

Please review details for the breaking changes and alternatives in the [Stripe API changelog](https://docs.stripe.com/changelog/basil) before upgrading.

* Remove support for resources `UsageRecordSummary` and `UsageRecord`
* Remove support for `New` method on resource `UsageRecord`
* Remove support for `List` method on resource `UsageRecordSummary`
* Remove support for `UpcomingLines` and `Upcoming` methods on resource `Invoice`
* Remove support for `UsageRecordSummaries` method on resource `SubscriptionItem`
* Remove support for `Invoice` on `Charge` and `PaymentIntent`
* Remove support for `ShippingDetails` on `CheckoutSession`
* Remove support for `Carrier`, `Phone`, and `TrackingNumber` on `CheckoutSessionCollectedInformationShippingDetails`
* Remove support for `Refund` on `CreditNoteParams`, `CreditNotePreviewLinesParams`, `CreditNotePreviewParams`, and `CreditNote`
* Remove support for `TaxAmounts` on `CreditNoteLineItem`, `CreditNote`, and `InvoiceLineItem`
* Remove support for `AmountExcludingTax` and `UnitAmountExcludingTax` on `CreditNoteLineItem` and `InvoiceLineItem`
* Remove support for `Coupon` on `CustomerParams`, `InvoiceCreatePreviewParams`, `InvoiceCreatePreviewScheduleDetailsPhasesParams`, `SubscriptionParams`, `SubscriptionSchedulePhasesParams`, and `SubscriptionSchedulePhases`
* Remove support for `PromotionCode` on `CustomerParams` and `SubscriptionParams`
* Remove support for `Price` on `InvoiceAddLinesLinesParams`, `InvoiceItemParams`, `InvoiceItem`, `InvoiceLineItemParams`, `InvoiceLineItem`, and `InvoiceUpdateLinesLinesParams`
* Remove support for `BillingThresholds` on `InvoiceCreatePreviewScheduleDetailsPhasesItemsParams`, `InvoiceCreatePreviewScheduleDetailsPhasesParams`, `InvoiceCreatePreviewSubscriptionDetailsItemsParams`, `SubscriptionItemParams`, `SubscriptionItem`, `SubscriptionItemsParams`, `SubscriptionParams`, `SubscriptionScheduleDefaultSettingsParams`, `SubscriptionScheduleDefaultSettings`, `SubscriptionSchedulePhasesItemsParams`, `SubscriptionSchedulePhasesItems`, `SubscriptionSchedulePhasesParams`, `SubscriptionSchedulePhases`, and `Subscription`
* Remove support for `ApplicationFeeAmount`, `Charge`, `PaidOutOfBand`, `Paid`, `PaymentIntent`, `Quote`, `Subscription`, `SubscriptionDetails`, `SubscriptionProrationDate`, `Tax`, `TotalTaxAmounts`, and `TransferData` on `Invoice`
* Remove support for `Discount` on `Invoice` and `Subscription`
* Remove support for `InvoiceItem`, `ProrationDetails`, `Proration`, `TaxRates`, and `Type` on `InvoiceLineItem`
* Remove support for `Plan` and `SubscriptionItem` on `InvoiceItem` and `InvoiceLineItem`
* Remove support for `UnitAmount` on `InvoiceItemParams` and `InvoiceItem`
* Remove support for `Subscription` and `UnitAmountDecimal` on `InvoiceItem`
* Remove support for `NaverPay` on `PaymentMethodParams`
* Remove support for `AggregateUsage` on `PlanParams`, `Plan`, `PriceRecurringParams`, and `PriceRecurring`
* Remove support for `CurrentPeriodEnd` and `CurrentPeriodStart` on `Subscription`

### ⚠️ Other Breaking changes in the SDK

* [#1999](https://github.com/stripe/stripe-go/pull/1999) Upgrade to go 1.18
    * Go version 1.18 or later is now required to address security vulnerabilities in Go <= 1.17. In particular, HTTP/2 is enabled now by default for all users (it was disabled for Go <= 1.14).
* [#1998](https://github.com/stripe/stripe-go/pull/1998) Breaking changes to support V2
    * Renamed the `stripe.Amount` type in the `stripe.Balance` object to `stripe.BalanceAmount`
    * Changed the signature of the `CallRaw` method in the `stripe.Backend` interface to accept a `[]byte` instead of `*form.Values` in its fourth argument. Call sites can safely replace a `*form.Values` argument `v` with `[]byte(v.Encode())` (`Encode` is `nil`-safe).

### Additions to Stripe API

* Add support for new resource `InvoicePayment`
* Add support for `Get` and `List` methods on resource `InvoicePayment`
* Add support for `BilliePayments`, `NzBankAccountBECSDebitPayments`, and `SatispayPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `HostedPaymentMethodSave` on `AccountSettingsInvoicesParams` and `AccountSettingsInvoices`
* Add support for `Invoices` on `AccountSettingsParams`
* Add support for new values `forwarding_api_retryable_upstream_error` and `setup_intent_mobile_wallet_unsupported` on enums `InvoiceLastFinalizationError.Code`, `PaymentIntentLastPaymentError.Code`, `SetupAttemptSetupError.Code`, `SetupIntentLastSetupError.Code`, and `StripeError.Code`
* Add support for new values `stripe_balance_payment_debit_reversal` and `stripe_balance_payment_debit` on enum `BalanceTransaction.Type`
* Add support for new values `information_missing`, `invalid_signator`, `verification_failed_authorizer_authority`, and `verification_rejected_ownership_exemption_reason` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
* Add support for new value `last` on enum `BillingMeterDefaultAggregation.Formula`
* Add support for `PresentmentDetails` on `Charge`, `CheckoutSession`, `PaymentIntent`, and `Refund`
* Add support for `Billie` and `Satispay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `NzBankAccount` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `MandatePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `OptionalItems` on `CheckoutSessionParams`, `CheckoutSession`, `PaymentLinkParams`, and `PaymentLink`
* Add support for `Permissions` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `ShippingOptions` on `CheckoutSessionParams`
* Add support for new value `custom` on enum `CheckoutSession.UIMode`
* Add support for `BuyerID` on `ConfirmationTokenPaymentMethodPreviewNaverPay` and `PaymentMethodNaverPay`
* Add support for new values `billie`, `nz_bank_account`, and `satispay` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `Refunds` on `CreditNoteParams`, `CreditNotePreviewLinesParams`, `CreditNotePreviewParams`, and `CreditNote`
* Add support for `TotalTaxes` on `CreditNote` and `Invoice`
* Add support for `Taxes` on `CreditNoteLineItem` and `InvoiceLineItem`
* Add support for `TaxabilityReason` on `InvoiceAddLinesLinesTaxAmountsParams`, `InvoiceLineItemTaxAmountsParams`, and `InvoiceUpdateLinesLinesTaxAmountsParams`
* Add support for `JurisdictionLevel` on `InvoiceAddLinesLinesTaxAmountsTaxRateDataParams`, `InvoiceLineItemTaxAmountsTaxRateDataParams`, and `InvoiceUpdateLinesLinesTaxAmountsTaxRateDataParams`
* Add support for `AmountOverpaid`, `ConfirmationSecret`, and `Payments` on `Invoice`
* Add support for `Parent` on `InvoiceItem`, `InvoiceLineItem`, and `Invoice`
* Add support for new values `klarna` and `nz_bank_account` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `CheckoutSession` on `CustomerBalanceTransaction`
* Add support for new values `checkout_session_subscription_payment_canceled` and `checkout_session_subscription_payment` on enum `CustomerBalanceTransaction.Type`
* Add support for new value `invoice.overpaid` on enum `Event.Type`
* Add support for `Pricing` on `InvoiceAddLinesLinesParams`, `InvoiceItemParams`, `InvoiceItem`, `InvoiceLineItemParams`, `InvoiceLineItem`, and `InvoiceUpdateLinesLinesParams`
* Add support for `Wifi` on `TerminalConfigurationParams` and `TerminalConfiguration`
* Add support for `NzBankTransfer` on `RefundDestinationDetails`
* Add support for new value `canceled` on enum `Review.ClosedReason`
* Add support for `CurrentPeriodEnd` and `CurrentPeriodStart` on `SubscriptionItem`
* Add support for `NaverPay` on `MandatePaymentMethodDetails` and `SetupAttemptPaymentMethodDetails`
* Add support for `SetupFutureUsage` on `PaymentIntentConfirmPaymentMethodOptionsNaverPayParams`, `PaymentIntentPaymentMethodOptionsNaverPayParams`, and `PaymentIntentPaymentMethodOptionsNaverPay`
* Add support for new value `expired` on enum `PaymentIntent.CancellationReason`
* Add support for `DefaultValue` on `PaymentLinkCustomFieldsDropdownParams`, `PaymentLinkCustomFieldsDropdown`, `PaymentLinkCustomFieldsNumericParams`, `PaymentLinkCustomFieldsNumeric`, `PaymentLinkCustomFieldsTextParams`, and `PaymentLinkCustomFieldsText`
* Add support for new values `billie` and `satispay` on enum `PaymentLink.PaymentMethodTypes`
