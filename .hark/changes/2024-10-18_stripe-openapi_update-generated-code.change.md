---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1934
is_stripe_api_change: true
released_in_version: 80.3.0-beta.1
---

* Add support for `AlmaPayments`, `GopayPayments`, `KakaoPayPayments`, `KrCardPayments`, `NaverPayPayments`, `PaycoPayments`, `QrisPayments`, `SamsungPayPayments`, `ShopeepayPayments`, `TreasuryEvolve`, `TreasuryFifthThird`, and `TreasuryGoldmanSachs` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `ScheduleAtPeriodEnd` on `BillingPortalConfigurationFeaturesSubscriptionUpdateParams` and `BillingPortalConfigurationFeaturesSubscriptionUpdate`
* Add support for `Alma` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `RefundDestinationDetails`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `Gopay`, `Qris`, and `Shopeepay` on `ChargePaymentMethodDetails`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new values `alma`, `gopay`, `qris`, and `shopeepay` on enums `ConfirmationTokenPaymentMethodPreviewType` and `PaymentMethodType`
* Add support for `Metadata` on `ForwardingRequestParams`
* Add support for new values `jp_credit_transfer`, `kakao_pay`, `kr_card`, `naver_pay`, and `payco` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Remove support for value `expired` from enum `IssuingAuthorizationStatus`
* Add support for new values `alma`, `gopay`, `qris`, and `shopeepay` on enum `PaymentLinkPaymentMethodTypes`
* Add support for `AmazonPay` on `PaymentMethodDomain`
* Add support for `ExternalReference` on `TaxFormListPayeeParams` and `TaxFormPayee`
* Change type of `TaxFormListPayeeTypeParams` and `TaxFormPayeeType` from `literal('account')` to `enum('account'|'external_reference')`
* Add support for `AuSerr`, `CaMrdp`, `EUDac7`, `GBMrdp`, and `NzMrdp` on `TaxForm`
* Add support for new values `au_serr`, `ca_mrdp`, `eu_dac7`, `gb_mrdp`, and `nz_mrdp` on enum `TaxFormType`
* Add support for `Pln` on `TerminalConfigurationTippingParams` and `TerminalConfigurationTipping`
* Add support for `Bank` on `TreasuryFinancialAccountFeaturesFinancialAddressesAbaParams`, `TreasuryFinancialAccountFeaturesFinancialAddressesAba`, and `TreasuryFinancialAccountUpdateFeaturesFinancialAddressesAbaParams`
