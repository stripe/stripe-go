---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/2417
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.4.0
---

* Add support for new resource `BillingFeedbackOption`
* Add support for `Deactivate`, `Get`, `List`, `New`, and `Update` methods on resource `BillingFeedbackOption`
* Add support for `PaymentMethodSettings` on `AccountSessionComponentsParams` and `AccountSessionComponents`
* Add support for `FeedbackOptions` on `BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReasonParams` and `BillingPortalConfigurationFeaturesSubscriptionCancelCancellationReason`
* Add support for `CustomerUpdate` on `BillingPortalSessionFlow`
* Add support for new value `customer_update` on enum `BillingPortalSessionFlow.Type`
* Add support for `FundingSourceGroup` on `ChargePaymentMethodDetailsCardWalletLink` and `ChargePaymentMethodDetailsLink`
* Add support for `FundingTypesBlocked` on `CheckoutSessionPaymentMethodOptionsCardRestrictionsParams` and `CheckoutSessionPaymentMethodOptionsCardRestrictions`
* Add support for `Metadata` on `ConfirmationToken`
* Add support for `ActiveEntitlements` and `CustomerPortal` on `CustomerSessionComponentsParams` and `CustomerSessionComponents`
* Add support for `Country` on `FinancialConnectionsSessionFilters`
* Add support for `FrozenFields` on `InvoiceItem`
* Add support for `Billie` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for new value `billie` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* ⚠️ Remove support for `Cryptogram` on `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure` and `PaymentRecordPaymentMethodDetailsCardThreeDSecure`
* Add support for new value `touch_n_go` on enums `PaymentIntent.AllowedPaymentMethodTypes` and `SetupIntent.AllowedPaymentMethodTypes`
* Add support for `ApplicationFeeAmount`, `ApplicationFeePercent`, `OnBehalfOf`, and `TransferData` on `PaymentLinkParams`
* Add support for `FeedbackOption` on `SubscriptionCancelCancellationDetailsParams`, `SubscriptionCancellationDetailsParams`, and `SubscriptionCancellationDetails`
* Add support for `Igic` on `TaxRegistrationCountryOptionsAtParams`, `TaxRegistrationCountryOptionsAt`, `TaxRegistrationCountryOptionsBeParams`, `TaxRegistrationCountryOptionsBe`, `TaxRegistrationCountryOptionsBgParams`, `TaxRegistrationCountryOptionsBg`, `TaxRegistrationCountryOptionsCyParams`, `TaxRegistrationCountryOptionsCy`, `TaxRegistrationCountryOptionsCzParams`, `TaxRegistrationCountryOptionsCz`, `TaxRegistrationCountryOptionsDeParams`, `TaxRegistrationCountryOptionsDe`, `TaxRegistrationCountryOptionsDkParams`, `TaxRegistrationCountryOptionsDk`, `TaxRegistrationCountryOptionsEeParams`, `TaxRegistrationCountryOptionsEe`, `TaxRegistrationCountryOptionsEsParams`, `TaxRegistrationCountryOptionsEs`, `TaxRegistrationCountryOptionsFiParams`, `TaxRegistrationCountryOptionsFi`, `TaxRegistrationCountryOptionsFrParams`, `TaxRegistrationCountryOptionsFr`, `TaxRegistrationCountryOptionsGrParams`, `TaxRegistrationCountryOptionsGr`, `TaxRegistrationCountryOptionsHrParams`, `TaxRegistrationCountryOptionsHr`, `TaxRegistrationCountryOptionsHuParams`, `TaxRegistrationCountryOptionsHu`, `TaxRegistrationCountryOptionsIeParams`, `TaxRegistrationCountryOptionsIe`, `TaxRegistrationCountryOptionsItParams`, `TaxRegistrationCountryOptionsIt`, `TaxRegistrationCountryOptionsLtParams`, `TaxRegistrationCountryOptionsLt`, `TaxRegistrationCountryOptionsLuParams`, `TaxRegistrationCountryOptionsLu`, `TaxRegistrationCountryOptionsLvParams`, `TaxRegistrationCountryOptionsLv`, `TaxRegistrationCountryOptionsMtParams`, `TaxRegistrationCountryOptionsMt`, `TaxRegistrationCountryOptionsNlParams`, `TaxRegistrationCountryOptionsNl`, `TaxRegistrationCountryOptionsPlParams`, `TaxRegistrationCountryOptionsPl`, `TaxRegistrationCountryOptionsPtParams`, `TaxRegistrationCountryOptionsPt`, `TaxRegistrationCountryOptionsRoParams`, `TaxRegistrationCountryOptionsRo`, `TaxRegistrationCountryOptionsSeParams`, `TaxRegistrationCountryOptionsSe`, `TaxRegistrationCountryOptionsSiParams`, `TaxRegistrationCountryOptionsSi`, `TaxRegistrationCountryOptionsSkParams`, and `TaxRegistrationCountryOptionsSk`
* Add support for error codes `authentication_failure`, `capability_not_active`, `expired_payment_method`, `incorrect_postal_code`, `invalid_canceled_subscription_fields`, and `payment_method_restricted` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, `StripeError`, and `TerminalReaderActionApiError`
