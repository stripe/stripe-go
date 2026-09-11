---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1838
is_stripe_api_change: true
released_in_version: 76.24.0
---

* Change type of `CheckoutSessionPaymentMethodOptionsSwishReferenceParams` from `emptyable(string)` to `string`
* Add support for `SubscriptionItem` on `Discount`
* Add support for `Email` and `Phone` on `IdentityVerificationReport`, `IdentityVerificationSessionOptionsParams`, `IdentityVerificationSessionOptions`, and `IdentityVerificationSessionVerifiedOutputs`
* Add support for `VerificationFlow` on `IdentityVerificationReport`, `IdentityVerificationSessionParams`, and `IdentityVerificationSession`
* Add support for new value `verification_flow` on enums `IdentityVerificationReportType` and `IdentityVerificationSessionType`
* Add support for `ProvidedDetails` on `IdentityVerificationSessionParams` and `IdentityVerificationSession`
* Add support for new values `email_unverified_other`, `email_verification_declined`, `phone_unverified_other`, and `phone_verification_declined` on enum `IdentityVerificationSessionLastErrorCode`
* Add support for `PromotionCode` on `InvoiceDiscountsParams`, `InvoiceItemDiscountsParams`, and `QuoteDiscountsParams`
* Add support for `Discounts` on `InvoiceUpcomingLinesSubscriptionItemsParams`, `InvoiceUpcomingSubscriptionItemsParams`, `QuoteLineItemsParams`, `SubscriptionAddInvoiceItemsParams`, `SubscriptionItemParams`, `SubscriptionItem`, `SubscriptionItemsParams`, `SubscriptionParams`, `SubscriptionSchedulePhasesAddInvoiceItemsParams`, `SubscriptionSchedulePhasesAddInvoiceItems`, `SubscriptionSchedulePhasesItemsParams`, `SubscriptionSchedulePhasesItems`, `SubscriptionSchedulePhasesParams`, `SubscriptionSchedulePhases`, and `Subscription`
* Add support for `AllowedMerchantCountries` and `BlockedMerchantCountries` on `IssuingCardSpendingControlsParams`, `IssuingCardSpendingControls`, `IssuingCardholderSpendingControlsParams`, and `IssuingCardholderSpendingControls`
* Add support for `Zip` on `PaymentMethodConfigurationParams` and `PaymentMethodConfiguration`
* Add support for `Offline` on `SetupAttemptPaymentMethodDetailsCardPresent`
* Add support for `CardPresent` on `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for new value `mobile_phone_reader` on enum `TerminalReaderDeviceType`
