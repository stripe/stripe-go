---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1908
is_stripe_api_change: true
released_in_version: 79.11.0-beta.1
---

* Add support for new resources `Billing.MeterErrorReport` and `Terminal.ReaderCollectedData`
* Add support for `Get` method on resource `ReaderCollectedData`
* Add support for `Recipients` on `AccountSessionComponentsParams`
* Add support for new value `terminal_reader_collected_data_invalid` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `BusinessName`, `Email`, `Phone`, and `TaxIDs` on `CheckoutSessionCollectedInformation`
* Add support for new value `billing.meter_error_report.triggered` on enum `EventType`
* Add support for `RegulatoryReportingFile` on `IssuingCreditUnderwritingRecordCorrectParams`, `IssuingCreditUnderwritingRecordCreateFromProactiveReviewParams`, `IssuingCreditUnderwritingRecordReportDecisionParams`, and `IssuingCreditUnderwritingRecord`
* Add support for new value `mb_way` on enum `PaymentLinkPaymentMethodTypes`
* Remove support for `Rechnung` on `PaymentMethodParams`
