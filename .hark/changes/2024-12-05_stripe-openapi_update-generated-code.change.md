---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1953
is_stripe_api_change: true
released_in_version: 81.2.0-beta.2
---

* Add support for `AutomaticIndirectTax` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for new values `payout_minimum_balance_hold` and `payout_minimum_balance_release` on enum `BalanceTransactionType`
* Add support for `ReferencePrefix` on `CheckoutSessionPaymentMethodOptionsBacsDebitMandateOptionsParams`, `CheckoutSessionPaymentMethodOptionsBacsDebitMandateOptions`, `CheckoutSessionPaymentMethodOptionsSepaDebitMandateOptionsParams`, `CheckoutSessionPaymentMethodOptionsSepaDebitMandateOptions`, `OrderPaymentSettingsPaymentMethodOptionsSepaDebitMandateOptionsParams`, `OrderPaymentSettingsPaymentMethodOptionsSepaDebitMandateOptions`, `PaymentIntentConfirmPaymentMethodOptionsBacsDebitMandateOptionsParams`, `PaymentIntentConfirmPaymentMethodOptionsSepaDebitMandateOptionsParams`, `PaymentIntentPaymentMethodOptionsBacsDebitMandateOptionsParams`, `PaymentIntentPaymentMethodOptionsBacsDebitMandateOptions`, `PaymentIntentPaymentMethodOptionsSepaDebitMandateOptionsParams`, `PaymentIntentPaymentMethodOptionsSepaDebitMandateOptions`, `SetupIntentConfirmPaymentMethodOptionsBacsDebitMandateOptionsParams`, `SetupIntentConfirmPaymentMethodOptionsSepaDebitMandateOptionsParams`, `SetupIntentPaymentMethodOptionsBacsDebitMandateOptionsParams`, `SetupIntentPaymentMethodOptionsBacsDebitMandateOptions`, `SetupIntentPaymentMethodOptionsSepaDebitMandateOptionsParams`, and `SetupIntentPaymentMethodOptionsSepaDebitMandateOptions`
* Add support for `DisabledReason` on `InvoiceAutomaticTax`, `SubscriptionAutomaticTax`, `SubscriptionScheduleDefaultSettingsAutomaticTax`, and `SubscriptionSchedulePhasesAutomaticTax`
* Add support for `TrialPeriodDays` on `PaymentLinkSubscriptionDataParams`
