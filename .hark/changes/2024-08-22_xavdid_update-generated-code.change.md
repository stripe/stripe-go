---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1907
is_stripe_api_change: true
released_in_version: 79.9.0-beta.2
---

* Add support for `MbWayPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `MbWay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new value `mb_way` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
* Remove support for values `accepted`, `partner_rejected`, and `submitted` from enum `DisputeEvidenceDetailsEnhancedEligibilityVisaCompellingEvidence3Status`
* Add support for new value `hr_oib` on enum `OrderTaxDetailsTaxIdsType`
* Remove support for `Phases` on `QuoteParams`
* Remove support for `FromSchedule` on `QuoteSubscriptionDataParams`
