---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2380
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.3.0-beta.1
---

* Add support for `Get` and `List` methods on resource `ProductCatalogTrialOffer`
* Add support for `TaxItems` on `ChargeCapturePaymentDetailsCarRentalDataTotalTaxParams`, `ChargeCapturePaymentDetailsFlightDataTotalTaxParams`, `ChargeCapturePaymentDetailsLodgingDataTotalTaxParams`, `ChargePaymentDetailsCarRentalDataTotalTaxParams`, `ChargePaymentDetailsFlightDataTotalTaxParams`, `ChargePaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentPaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentPaymentDetailsCarRentalDatumTotalTax`, `PaymentIntentPaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentPaymentDetailsFlightDatumTotalTax`, `PaymentIntentPaymentDetailsLodgingDataTotalTaxParams`, and `PaymentIntentPaymentDetailsLodgingDatumTotalTax`
* ⚠️ Remove support for `Taxes` on `ChargeCapturePaymentDetailsCarRentalDataTotalTaxParams`, `ChargeCapturePaymentDetailsFlightDataTotalTaxParams`, `ChargeCapturePaymentDetailsLodgingDataTotalTaxParams`, `ChargePaymentDetailsCarRentalDataTotalTaxParams`, `ChargePaymentDetailsFlightDataTotalTaxParams`, `ChargePaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentCapturePaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentConfirmPaymentDetailsLodgingDataTotalTaxParams`, `PaymentIntentPaymentDetailsCarRentalDataTotalTaxParams`, `PaymentIntentPaymentDetailsCarRentalDatumTotalTax`, `PaymentIntentPaymentDetailsFlightDataTotalTaxParams`, `PaymentIntentPaymentDetailsFlightDatumTotalTax`, `PaymentIntentPaymentDetailsLodgingDataTotalTaxParams`, and `PaymentIntentPaymentDetailsLodgingDatumTotalTax`
* Add support for `TaxID` on `CheckoutSessionCollectedInformation`
* ⚠️ Remove support for `TaxIDs` on `CheckoutSessionCollectedInformation`
* Add support for `Mode` on `FinancialConnectionsSessionManualEntry`
* Add support for `Name` on `IssuingCardholderParams`
* Add support for new value `ic_nif` on enums `OrderTaxDetailsTaxId.Type` and `QuotePreviewInvoiceCustomerTaxIds.Type`
* Add support for new values `alipay` and `mb_way` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
* Add support for `CustomFields`, `Description`, and `Footer` on `QuotePreviewSubscriptionScheduleDefaultSettingsInvoiceSettings` and `QuotePreviewSubscriptionSchedulePhaseInvoiceSettings`
* Add support for `Trial` on `QuotePreviewSubscriptionSchedulePhase`
* ⚠️ Remove support for `ACSSDebit`, `AUBECSDebit`, `AfterpayClearpay`, `Alipay`, `Alma`, `AmazonPay`, `BACSDebit`, `BLIK`, `Bancontact`, `Billie`, `Bizum`, `Boleto`, `CardPresent`, `CashApp`, `Crypto`, `CustomerBalance`, `EPS`, `FPX`, `Giropay`, `Gopay`, `Grabpay`, `IDBankTransfer`, `IDEAL`, `InteracPresent`, `KakaoPay`, `Konbini`, `KrCard`, `MbWay`, `Mobilepay`, `Multibanco`, `NaverPay`, `NzBankAccount`, `OXXO`, `P24`, `PayByBank`, `PayNow`, `Payco`, `Paypal`, `Paypay`, `Payto`, `Pix`, `PromptPay`, `Qris`, `Rechnung`, `RevolutPay`, `SEPADebit`, `SamsungPay`, `Satispay`, `Scalapay`, `Shopeepay`, `Sofort`, `StripeBalance`, `Sunbit`, `Swish`, `TWINT`, `USBankAccount`, `Upi`, `WeChatPay`, and `Zip` on `SharedPaymentGrantedTokenPaymentMethodDetails`
* ⚠️ Remove support for values `acss_debit`, `afterpay_clearpay`, `alipay`, `alma`, `amazon_pay`, `au_becs_debit`, `bacs_debit`, `bancontact`, `billie`, `bizum`, `blik`, `boleto`, `card_present`, `cashapp`, `crypto`, `custom`, `customer_balance`, `eps`, `fpx`, `giropay`, `gopay`, `grabpay`, `id_bank_transfer`, `ideal`, `interac_present`, `kakao_pay`, `konbini`, `kr_card`, `mb_way`, `mobilepay`, `multibanco`, `naver_pay`, `nz_bank_account`, `oxxo`, `p24`, `pay_by_bank`, `payco`, `paynow`, `paypal`, `paypay`, `payto`, `pix`, `promptpay`, `qris`, `rechnung`, `revolut_pay`, `samsung_pay`, `satispay`, `scalapay`, `sepa_debit`, `shopeepay`, `sofort`, `stripe_balance`, `sunbit`, `swish`, `twint`, `upi`, `us_bank_account`, `wechat_pay`, and `zip` from enum `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
* Add support for `UseStripeSDK` on `SharedPaymentIssuedTokenParams` and `SharedPaymentIssuedToken`
* Add support for `RedirectToURL` on `SharedPaymentIssuedTokenNextAction`
* ⚠️ Change type of `SharedPaymentIssuedTokenNextAction.Type` from `literal('use_stripe_sdk')` to `enum('redirect_to_url'|'use_stripe_sdk')`
* Add support for `Livemode` on `TaxLocation`
* Add support for `Source` on `V2IamActivityLogDetailsUserRoles`
* Add support for `Payout` on `V2MoneyManagementReceivedCreditBalanceTransfer`
* ⚠️ Remove support for `PayoutV1` on `V2MoneyManagementReceivedCreditBalanceTransfer`
* Add support for new value `payout` on enum `V2MoneyManagementReceivedCreditBalanceTransfer.Type`
* Add support for error codes `us_bank_account_microdeposits_cannot_be_confirmed` and `us_bank_account_microdeposits_cannot_be_sent` on `ControlledByAlternateResourceError`
