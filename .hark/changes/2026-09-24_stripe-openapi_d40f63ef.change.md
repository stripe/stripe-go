---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2421
semver_level: major
is_stripe_api_change: true
---

* Add support for new resource `RadarBillingEvaluation`
* Add support for `New` method on resource `RadarBillingEvaluation`
* Add support for `List` method on resource `ReservePlan`
* Add support for `AfterExpiration` on `BillingPortalSessionParams` and `BillingPortalSession`
* Add support for new value `fundbox_ca_financing` on enums `CapitalFinancingOffer.DisclaimerVariant` and `CapitalFinancingSummaryDetails.DisclaimerVariant`
* Add support for `SetupCredentialUsage` on `ChargePaymentMethodDetailsCard`, `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCard`, `SetupIntentConfirmPaymentMethodOptionsCardParams`, `SetupIntentPaymentMethodOptionsCardParams`, and `SetupIntentPaymentMethodOptionsCard`
* Add support for `StoredCredentialUsage` on `ChargePaymentMethodDetailsCard`, `PaymentAttemptRecordPaymentMethodDetailsCard`, `PaymentIntentConfirmPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCardParams`, `PaymentIntentPaymentMethodOptionsCard`, and `PaymentRecordPaymentMethodDetailsCard`
* Add support for `ExpiresAt` on `CheckoutSessionPaymentMethodOptionsBlikMandateOptionsParams`, `SubscriptionPaymentSettingsPaymentMethodOptionsBlikMandateOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsBlikMandateOptions`
* ⚠️ Remove support for `ExpiresAfter` on `CheckoutSessionPaymentMethodOptionsBlikMandateOptionsParams`, `SubscriptionPaymentSettingsPaymentMethodOptionsBlikMandateOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptionsBlikMandateOptions`
* Add support for `PaymentIntentData` on `CheckoutSessionParams`
* Add support for `Appeal` on `DisputeEvidenceParams` and `DisputeEvidence`
* Add support for `Livemode` on `FxQuote`
* ⚠️ Remove support for `CaptureMethod` on `PaymentIntentConfirmPaymentMethodOptionsPaypayParams` and `PaymentIntentPaymentMethodOptionsPaypayParams`
* Add support for `Active` on `ProductCatalogTrialOfferListParams`, `ProductCatalogTrialOfferParams`, and `ProductCatalogTrialOffer`
* Add support for `Nickname` on `ProductCatalogTrialOfferParams` and `ProductCatalogTrialOffer`
* ⚠️ Remove support for `Name` on `ProductCatalogTrialOfferParams` and `ProductCatalogTrialOffer`
* Add support for `StatusDetails` on `QuotePreviewInvoice`
* Add support for `CompanyDetails` and `Reference` on `QuotePreviewInvoicePaymentSettingsPaymentMethodOptionsBillie`
* Add support for `PauseSchedules` on `QuotePreviewSubscriptionSchedule`
* Add support for `Destination` on `ReserveHold`, `ReservePlan`, and `ReserveRelease`
* Add support for `ManualRelease` on `ReservePlan`
* Add support for new value `other` on enum `ReservePlan.Status`
* Add support for new values `manual_release` and `other` on enum `ReservePlan.Type`
* ⚠️ Add support for new value `hold_expired` on enum `ReserveRelease.Reason`
* ⚠️ Remove support for value `bulk_hold_expiry` from enum `ReserveRelease.Reason`
* Add support for new value `igic` on enum `TaxRegistrationCountryOptionsEs.Type`
* Add support for error codes `dispute_evidence_page_limit_exceeded`, `financial_connections_consent_locale_invalid`, `financial_connections_consent_locale_unsupported`, and `payment_evaluation_on_api_version_not_supported` on `QuotePreviewInvoiceLastFinalizationError`
