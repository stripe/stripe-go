---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1931
is_stripe_api_change: true
released_in_version: 81.0.0
---

* Add support for new resource `V2.EventDestinations`
* Add support for `New`, `Retrieve`, `Update`, `List`, `Delete`, `Disable`, `Enable` and `Ping` methods on resource `V2.EventDestinations`
* Add support for `SubmitCard` test helper method on resource `Issuing.Card`
* Add support for `Groups` on `AccountParams` and `Account`
* Add support for `AlmaPayments`, `KakaoPayPayments`, `KrCardPayments`, `NaverPayPayments`, `PaycoPayments`, and `SamsungPayPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `DisableStripeUserAuthentication` on `AccountSessionComponentsAccountManagementFeaturesParams`, `AccountSessionComponentsAccountManagementFeatures`, `AccountSessionComponentsAccountOnboardingFeaturesParams`, `AccountSessionComponentsAccountOnboardingFeatures`, `AccountSessionComponentsBalancesFeaturesParams`, `AccountSessionComponentsBalancesFeatures`, `AccountSessionComponentsNotificationBannerFeaturesParams`, `AccountSessionComponentsNotificationBannerFeatures`, `AccountSessionComponentsPayoutsFeaturesParams`, and `AccountSessionComponentsPayoutsFeatures`
* Add support for `ScheduleAtPeriodEnd` on `BillingPortalConfigurationFeaturesSubscriptionUpdateParams` and `BillingPortalConfigurationFeaturesSubscriptionUpdate`
* Add support for `Alma` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `RefundDestinationDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `KakaoPay` and `KrCard` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `MandatePaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `NaverPay` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Payco` and `SamsungPay` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new values `by_tin`, `ma_vat`, `md_vat`, `tz_vat`, `uz_tin`, and `uz_vat` on enums `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, `TaxIdType`, and `TaxTransactionCustomerDetailsTaxIdsType`
* Add support for new values `alma`, `kakao_pay`, `kr_card`, `naver_pay`, `payco`, and `samsung_pay` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
* Add support for `EnhancedEvidence` on `DisputeEvidenceParams` and `DisputeEvidence`
* Add support for `EnhancedEligibilityTypes` on `Dispute`
* Add support for `EnhancedEligibility` on `DisputeEvidenceDetails`
* Add support for new values `issuing_transaction.purchase_details_receipt_updated` and `refund.failed` on enum `EventType`
* Add support for `Metadata` on `ForwardingRequestParams` and `ForwardingRequest`
* Add support for `AutomaticallyFinalizesAt` on `InvoiceParams`
* Add support for new values `jp_credit_transfer`, `kakao_pay`, `kr_card`, `naver_pay`, and `payco` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Add support for new value `alma` on enum `PaymentLinkPaymentMethodTypes`
* Add support for `AmazonPay` on `PaymentMethodDomain`
* Change type of `RefundNextActionDisplayDetails` from `nullable(RefundNextActionDisplayDetails)` to `RefundNextActionDisplayDetails`
* Add support for new value `retail_delivery_fee` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationTaxBreakdownTaxRateDetailsTaxType`, `TaxRateTaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetailsTaxType`
* Add support for `FlatAmount` and `RateType` on `TaxCalculationTaxBreakdownTaxRateDetails` and `TaxRate`
* Add support for `By`, `Cr`, `Ec`, `Ma`, `Md`, `RU`, `Rs`, `Tz`, and `Uz` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
* Add support for new value `state_retail_delivery_fee` on enum `TaxRegistrationCountryOptionsUsType`
* Add support for `Pln` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
