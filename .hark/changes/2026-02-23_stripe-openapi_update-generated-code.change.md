---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/2277
is_stripe_api_change: true
released_in_version: 84.4.0
---

* Add support for new resources `ReserveHold`, `ReservePlan`, and `ReserveRelease`
* Add support for `Location` and `Reader` on `ChargePaymentMethodDetailsCardPresent`, `ChargePaymentMethodDetailsInteracPresent`, `ConfirmationTokenPaymentMethodPreviewCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsCardPresent`, `PaymentAttemptRecordPaymentMethodDetailsInteracPresent`, `PaymentMethodCardGeneratedFromPaymentMethodDetailsCardPresent`, `PaymentRecordPaymentMethodDetailsCardPresent`, and `PaymentRecordPaymentMethodDetailsInteracPresent`
* Add support for new value `lk_vat` on enums `CheckoutSessionCustomerDetailsTaxIds.Type`, `TaxCalculationCustomerDetailsTaxId.Type`, `TaxId.Type`, and `TaxTransactionCustomerDetailsTaxId.Type`
* Add support for new values `reserve.hold.created`, `reserve.hold.updated`, `reserve.plan.created`, `reserve.plan.disabled`, `reserve.plan.expired`, `reserve.plan.updated`, and `reserve.release.created` on enum `Event.Type`
* Add support for new values `terminal_wifi_certificate` and `terminal_wifi_private_key` on enum `File.Purpose`
* Add support for new value `pay_by_bank` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `DisplayName` and `ServiceUserNumber` on `MandatePaymentMethodDetailsBacsDebit`
* Add support for `TransactionPurpose` on `PaymentIntentConfirmPaymentMethodOptionsUsBankAccountParams`, `PaymentIntentPaymentMethodOptionsUsBankAccountParams`, and `PaymentIntentPaymentMethodOptionsUsBankAccount`
* Add support for `OptionalItems` on `PaymentLinkParams`
* Remove support for unused `CardIssuerDecline` on `RadarPaymentEvaluationInsights`
* Add support for `PaymentBehavior` on `SubscriptionItemParams`
* Add support for `Lk` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
* Add support for `Cellular` and `StripeS710` on `TerminalConfigurationParams` and `TerminalConfiguration`
* Add support for new values `simulated_stripe_s710` and `stripe_s710` on enum `TerminalReader.DeviceType`
* Add support for snapshot events `EventTypeReserveHoldCreated` and `EventTypeReserveHoldUpdated` with resource `ReserveHold`
* Add support for snapshot events `EventTypeReservePlanCreated`, `EventTypeReservePlanDisabled`, `EventTypeReservePlanExpired`, and `EventTypeReservePlanUpdated` with resource `ReservePlan`
* Add support for snapshot event `EventTypeReserveReleaseCreated` with resource `ReserveRelease`
* Add support for error codes `storer_capability_missing` and `storer_capability_not_active` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
