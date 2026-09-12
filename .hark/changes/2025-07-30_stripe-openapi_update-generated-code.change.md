---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2083
is_stripe_api_change: true
released_in_version: 82.5.0-beta.1
---

* Add support for new resources `BillingMeterUsageRow`, `BillingMeterUsage`, and `TerminalOnboardingLink`
* Add support for `Get` method on resource `BillingMeterUsage`
* Add support for `New` method on resource `TerminalOnboardingLink`
* Add support for `MonthlyPayoutDays` and `WeeklyPayoutDays` on `BalanceSettingsPayoutsScheduleParams` and `BalanceSettingsPayoutsSchedule`
* Remove support for `MonthlyAnchor` and `WeeklyAnchor` on `BalanceSettingsPayoutsScheduleParams` and `BalanceSettingsPayoutsSchedule`
* Add support for `DelayDaysOverride` on `BalanceSettingsSettlementTimingParams`
* Remove support for `DelayDays` on `BalanceSettingsSettlementTimingParams`
* Add support for `UpdateDiscounts` on `CheckoutSessionPermissionsParams`
* Add support for `Discounts` and `SubscriptionData` on `CheckoutSessionParams`
* Add support for `SmartDisputes` on `Dispute`
* Add support for `Upi` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for new value `upi` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `TransactionID` on `PaymentAttemptRecordPaymentMethodDetailsCashapp` and `PaymentRecordPaymentMethodDetailsCashapp`
* Add support for `AmountDetails` on `PaymentIntentCaptureParams`, `PaymentIntentConfirmParams`, `PaymentIntentIncrementAuthorizationParams`, and `PaymentIntentParams`
* Add support for `PaymentDetails` on `PaymentIntentIncrementAuthorizationParams`
* Add support for `Storer` on `V2CoreAccountIdentityAttestationsTermsOfServiceParams` and `V2CoreAccountIdentityAttestationsTermsOfService`
* Add support for `CollectionOptions` on `V2CoreAccountLinkUseCaseAccountOnboardingParams`, `V2CoreAccountLinkUseCaseAccountOnboarding`, `V2CoreAccountLinkUseCaseAccountUpdateParams`, and `V2CoreAccountLinkUseCaseAccountUpdate`
* Change type of `V2CoreAccountLinkUseCaseAccountOnboarding.Configurations`, `V2CoreAccountLinkUseCaseAccountOnboardingParams.Configurations`, `V2CoreAccountLinkUseCaseAccountUpdate.Configurations`, and `V2CoreAccountLinkUseCaseAccountUpdateParams.Configurations` from `literal('recipient')` to `enum('customer'|'merchant'|'recipient'|'storer')`
* Add support for `BankAccountType` on `V2MoneyManagementPayoutMethodBankAccount`
* Add support for thin event `V2CoreAccountLinkReturnedEvent`
* Add support for thin event `V2MoneyManagementPayoutMethodUpdatedEvent` with related object `V2MoneyManagementPayoutMethod`
* Remove support for thin event `V2CoreAccountLinkCompletedEvent`
* Remove support for thin event `V2OffSessionPaymentRequiresCaptureEvent` with related object `V2PaymentsOffSessionPayment`
