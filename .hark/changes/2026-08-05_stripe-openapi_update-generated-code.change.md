---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2402
is_stripe_api_change: true
released_in_version: 86.3.0-alpha.2
---

* Add support for new resource `BillingFeedbackOptions`
* Add support for `SequraPayments` on `AccountCapabilities`
* Add support for `FeedbackOptions` on `BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReason`
* Add support for `Sequra` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodPreview`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentPaymentMethodOptions`, and `PaymentRecordPaymentMethodDetails`
* Add support for `RetrievalReferenceNumber` on `ChargePaymentMethodDetailsCardPresent`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsCardPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, and `PaymentRecordPaymentMethodDetailsCardPresent`
* Add support for `PricingGroup` on `ChargePaymentMethodDetailsLink`
* Add support for `TaxRates` on `CheckoutSessionShippingOptionParams`, `CheckoutSessionShippingOption`, and `CheckoutSessionShippingOptionsParams`
* Add support for new value `daikin` on enums `CheckoutSessionAutomaticSurcharge.Provider` and `PaymentLinkAutomaticSurcharge.Provider`
* Add support for `FundingTypesBlocked` on `CheckoutSessionPaymentMethodOptionsCardRestrictions`
* Add support for new value `sequra` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `Healthcare` on `IssuingAuthorizationParams`, `TestHelpersIssuingAuthorizationCapturePurchaseDetailsParams`, `TestHelpersIssuingTransactionCreateForceCapturePurchaseDetailsParams`, and `TestHelpersIssuingTransactionCreateUnlinkedRefundPurchaseDetailsParams`
* Add support for `IsAnomalous` on `PaymentAttemptRecordReportGuaranteedParams`
* Add support for new value `sequra` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `AadeData` on `PaymentIntentPaymentMethodOptionsCardPresent`
* Add support for `FeedbackOption` on `SubscriptionCancellationDetails`
* Add support for `Application` on `V2PaymentsOffSessionPayment`
* Add support for `Status` on `V2MoneyManagementFinancialAccountStatementListParams`
