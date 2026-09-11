---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1963
is_stripe_api_change: true
released_in_version: 81.3.0-beta.2
---

* Add support for `PayByBankPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `DirectorshipDeclaration` on `AccountCompanyParams` and `TokenAccountCompanyParams`
* Add support for `ProofOfUltimateBeneficialOwnership` on `AccountDocumentsParams`
* Add support for `TaxThresholdMonitoring` on `AccountSessionComponentsParams`
* Add support for `FinancialAccountTransactions`, `FinancialAccount`, `IssuingCard`, and `IssuingCardsList` on `AccountSessionComponents`
* Add support for new value `always_invoice` on enum `BillingPortalConfigurationFeaturesSubscriptionCancelProrationBehavior`
* Add support for `PayByBank` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Discounts` on `CheckoutSession`
* Add support for new value `SD` on enums `CheckoutSessionShippingAddressCollectionAllowedCountries` and `PaymentLinkShippingAddressCollectionAllowedCountries`
* Add support for new value `pay_by_bank` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
* Add support for `PhoneNumberCollection` on `PaymentLinkParams`
* Add support for new value `pay_by_bank` on enum `PaymentLinkPaymentMethodTypes`
* Add support for `Jpy` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
