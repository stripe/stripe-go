---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2078
is_stripe_api_change: true
released_in_version: 82.3.0
---

* Add support for `Migrate` method on resource `Subscription`
* Add support for `CollectPaymentMethod` and `ConfirmPaymentIntent` methods on resource `TerminalReader`
* Add support for `CryptoPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `ProofOfAddress` on `AccountDocumentsParams`
* Add support for `MonthlyPayoutDays` and `WeeklyPayoutDays` on `AccountSettingsPayoutsScheduleParams` and `AccountSettingsPayoutsSchedule`
* Add support for `Crypto` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Subscriptions` on `CheckoutSessionPaymentMethodOptionsKlarnaParams`, `PaymentIntentConfirmPaymentMethodOptionsKlarnaParams`, and `PaymentIntentPaymentMethodOptionsKlarnaParams`
* Add support for `BillingMode` on `CheckoutSessionSubscriptionDataParams`, `InvoiceCreatePreviewScheduleDetailsParams`, `InvoiceCreatePreviewSubscriptionDetailsParams`, `QuoteSubscriptionDataParams`, `QuoteSubscriptionData`, `SubscriptionParams`, `SubscriptionScheduleParams`, `SubscriptionSchedule`, and `Subscription`
* Change type of `ConfirmationTokenPaymentMethodOptionsCardInstallmentsPlan.Type`, `ConfirmationTokenPaymentMethodOptionsCardInstallmentsPlanParams.Type`, `InvoicePaymentSettingsPaymentMethodOptionsCardInstallmentsPlanParams.Type`, `PaymentIntentConfirmPaymentMethodOptionsCardInstallmentsPlanParams.Type`, `PaymentIntentPaymentMethodOptionsCardInstallmentsPlan.Type`, and `PaymentIntentPaymentMethodOptionsCardInstallmentsPlanParams.Type` from `literal('fixed_count')` to `enum('bonus'|'fixed_count'|'revolving')`
* Add support for new value `buut` on enum `ConfirmationTokenPaymentMethodPreviewIdeal.Bank`
* Add support for new value `BUUTNL2A` on enum `ConfirmationTokenPaymentMethodPreviewIdeal.BIC`
* Add support for new value `crypto` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Change type of `Dispute.EnhancedEligibilityTypes` from `literal('visa_compelling_evidence_3')` to `enum('visa_compelling_evidence_3'|'visa_compliance')`
* Add support for new value `compliance` on enum `DisputePaymentMethodDetailsCard.CaseType`
* Add support for new value `terminal.reader.action_updated` on enum `Event.Type`
* Add support for `RelatedPerson` on `IdentityVerificationSessionParams` and `IdentityVerificationSession`
* Add support for `Matching` on `IdentityVerificationSessionOptions`
* Add support for new value `crypto` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `Klarna` on `MandatePaymentMethodDetails`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptionsParams`, and `SetupIntentPaymentMethodOptions`
* Add support for `OnDemand` on `PaymentIntentConfirmPaymentMethodOptionsKlarnaParams` and `PaymentIntentPaymentMethodOptionsKlarnaParams`
* Change type of `PaymentIntentConfirmPaymentMethodOptionsKlarnaParams.SetupFutureUsage`, `PaymentIntentPaymentMethodOptionsKlarna.SetupFutureUsage`, and `PaymentIntentPaymentMethodOptionsKlarnaParams.SetupFutureUsage` from `literal('none')` to `enum('none'|'off_session'|'on_session')`
* Add support for `Ua` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
* Change type of `TerminalLocationParams.DisplayName` from `string` to `emptyable(string)`
* Add support for `CollectPaymentMethod` and `ConfirmPaymentIntent` on `TerminalReaderAction`
* Add support for new values `collect_payment_method` and `confirm_payment_intent` on enum `TerminalReaderAction.Type`
* Add support for `Status` on `TreasuryFinancialAccountListParams`
* Add support for snapshot event `EventTypeTerminalReaderActionUpdated` with resource `TerminalReader`
