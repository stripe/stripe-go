---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1932
is_stripe_api_change: true
released_in_version: 80.2.0-beta.2
---

* Add support for `SubmitCard` test helper method on resource `Issuing.Card`
* Add support for `Groups` on `AccountParams` and `Account`
* Add support for `DisableStripeUserAuthentication` on `AccountSessionComponentsAccountManagementFeaturesParams`, `AccountSessionComponentsAccountManagementFeatures`, `AccountSessionComponentsAccountOnboardingFeaturesParams`, `AccountSessionComponentsAccountOnboardingFeatures`, `AccountSessionComponentsBalancesFeaturesParams`, `AccountSessionComponentsBalancesFeatures`, `AccountSessionComponentsFinancialAccountFeaturesParams`, `AccountSessionComponentsNotificationBannerFeaturesParams`, `AccountSessionComponentsNotificationBannerFeatures`, `AccountSessionComponentsPayoutsFeaturesParams`, and `AccountSessionComponentsPayoutsFeatures`
* Add support for `CardSpendDisputeManagement` and `SpendControlManagement` on `AccountSessionComponentsIssuingCardsListFeaturesParams`
* Add support for new value `payout_statement_descriptor_profanity` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for `KakaoPay` and `KrCard` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `MandatePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `NaverPay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Payco` and `SamsungPay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new values `by_tin`, `ma_vat`, `md_vat`, `tz_vat`, `uz_tin`, and `uz_vat` on enums `CheckoutSessionCollectedInformationTaxIdsType`, `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `OrderTaxDetailsTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, `TaxIdType`, and `TaxTransactionCustomerDetailsTaxIdsType`
* Add support for new values `kakao_pay`, `kr_card`, `naver_pay`, `payco`, and `samsung_pay` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
* Add support for new value `refund.failed` on enum `EventType`
* Add support for `Metadata` on `ForwardingRequest`
* Add support for new value `expired` on enum `IssuingAuthorizationStatus`
* Add support for `LineItems` on `OrderPaymentSettingsPaymentMethodOptionsPaypalParams`, `OrderPaymentSettingsPaymentMethodOptionsPaypal`, `PaymentIntentConfirmPaymentMethodOptionsPaypalParams`, `PaymentIntentPaymentMethodOptionsPaypalParams`, and `PaymentIntentPaymentMethodOptionsPaypal`
* Add support for new value `retail_delivery_fee` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationTaxBreakdownTaxRateDetailsTaxType`, `TaxRateTaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetailsTaxType`
* Add support for `FlatAmount` and `RateType` on `TaxCalculationTaxBreakdownTaxRateDetails` and `TaxRate`
* Add support for `By`, `Cr`, `Ec`, `Ma`, `Md`, `RU`, `Rs`, `Tz`, and `Uz` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
* Add support for new value `state_retail_delivery_fee` on enum `TaxRegistrationCountryOptionsUsType`
