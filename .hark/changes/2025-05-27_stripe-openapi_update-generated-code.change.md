---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/2060
is_stripe_api_change: true
released_in_version: 82.3.0-beta.1
---

### Breaking changes
* Remove support for deprecated previews
  * Remove support for resources `BillingMeterErrorReport`, `GiftCardsCard`, `GiftCardsTransaction`, and `PrivacyRedactionJobRootObjects`
  * Remove support for `Get`, `List`, `New`, `Update`, and `Validate` methods on resource `GiftCardsCard`
  * Remove support for `Cancel`, `Confirm`, `Get`, `List`, `New`, and `Update` methods on resource `GiftCardsTransaction`
  * Remove support for `Provisioning` on `ProductParams` and `Product`
  * Remove support for snapshot event `EventTypeBillingMeterErrorReportTriggered` with resource `BillingMeterErrorReport`
  * Remove support for error codes `gift_card_balance_insufficient`, `gift_card_code_exists`, and `gift_card_inactive` on `Error` and `QuotePreviewInvoiceLastFinalizationError`
* Remove support for `Credits` on `OrderParams` and `Order`
* Remove support for `AmountRemaining` on `Order`
* Remove support for `AmountCredit` on `OrderTotalDetails`
* Remove support for `AsyncWorkflows` on `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentDecrementAuthorizationParams`, `PaymentIntentIncrementAuthorizationParams`, `PaymentIntentParams`, and `PaymentIntent`
* Remove support for values `credits_attributed_to_debits` and `legacy_prorations` from enums `QuotePreviewSubscriptionSchedule.BillingMode`, `QuoteSubscriptionData.BillingMode`, `Subscription.BillingMode`, and `SubscriptionSchedule.BillingMode`
* Remove support for `StatusDetails` and `Status` on `TaxAssociation`
* Change type of `InvoiceCreatePreviewSubscriptionDetailsParams.CancelAt` and `SubscriptionParams.CancelAt` from `DateTime` to `DateTime | enum('max_period_end'|'min_period_end')`
* Change type of `CheckoutSessionLineItemParams.Quantity` from `emptyable(longInteger)` to `longInteger`
* Change type of `PrivacyRedactionJob.Objects` from `$Privacy.RedactionJobRootObjects` to `RedactionResourceRootObjects`
* Change type of `PrivacyRedactionJob.Status` from `string` to `enum`
* Change type of `PrivacyRedactionJob.ValidationBehavior` from `string` to `enum('error'|'fix')`
* Change type of `PrivacyRedactionJobValidationError.Code` from `string` to `enum`
* Change type of `PrivacyRedactionJobValidationError.ErroringObject` from `map(string: string)` to `RedactionResourceErroringObject`

### Other changes
* Add support for `Migrate` method on resource `Subscription`
* Add support for `Distance`, `PickupLocationName`, `ReturnLocationName`, and `VehicleIdentificationNumber` on `ChargeCapturePaymentDetailsCarRentalParams`, `ChargePaymentDetailsCarRentalParams`, `PaymentIntentCapturePaymentDetailsCarRentalParams`, `PaymentIntentConfirmPaymentDetailsCarRentalParams`, `PaymentIntentPaymentDetailsCarRentalParams`, and `PaymentIntentPaymentDetailsCarRental`
* Add support for `DriverIdentificationNumber` and `DriverTaxNumber` on `ChargeCapturePaymentDetailsCarRentalDriverParams`, `ChargePaymentDetailsCarRentalDriverParams`, `PaymentIntentCapturePaymentDetailsCarRentalDriverParams`, `PaymentIntentConfirmPaymentDetailsCarRentalDriverParams`, `PaymentIntentPaymentDetailsCarRentalDriverParams`, and `PaymentIntentPaymentDetailsCarRentalDriver`
* Add support for `Institution` on `FinancialConnectionsAccount`
* Add support for `Countries` on `FinancialConnectionsInstitution`
* Add support for `Location` and `Reader` on `PaymentAttemptRecordPaymentMethodDetailsAffirm`, `PaymentAttemptRecordPaymentMethodDetailsWechatPay`, `PaymentRecordPaymentMethodDetailsAffirm`, and `PaymentRecordPaymentMethodDetailsWechatPay`
* Add support for `Hooks` on `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentDecrementAuthorizationParams`, `PaymentIntentIncrementAuthorizationParams`, `PaymentIntentParams`, and `PaymentIntent`
* Add support for `CardPresent` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptions`
* Change type of `PaymentRecordReportPaymentAttemptCanceledParams.Metadata`, `PaymentRecordReportPaymentAttemptFailedParams.Metadata`, `PaymentRecordReportPaymentAttemptGuaranteedParams.Metadata`, `PaymentRecordReportPaymentAttemptParams.Metadata`, and `PaymentRecordReportPaymentParams.Metadata` from `map(string: string)` to `emptyable(map(string: string))`
* Add support for `Livemode` on `PrivacyRedactionJob`
* Add support for new values `classic` and `flexible` on enums `QuotePreviewSubscriptionSchedule.BillingMode`, `QuoteSubscriptionData.BillingMode`, `Subscription.BillingMode`, and `SubscriptionSchedule.BillingMode`
* Add support for `BillingThresholds` on `QuotePreviewSubscriptionScheduleDefaultSettings`, `QuotePreviewSubscriptionSchedulePhaseItem`, and `QuotePreviewSubscriptionSchedulePhase`
* Add support for `BillingModeDetails` on `Subscription`
* Add support for `TaxTransactionAttempts` on `TaxAssociation`
* Add support for `ConfirmConfig` on `TerminalReaderActionConfirmPaymentIntent` and `TerminalReaderConfirmPaymentIntentParams`
* Add support for error code `forwarding_api_upstream_error` on `QuotePreviewInvoiceLastFinalizationError`
