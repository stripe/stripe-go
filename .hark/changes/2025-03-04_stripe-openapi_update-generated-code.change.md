---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/1987
is_stripe_api_change: true
released_in_version: 81.5.0-beta.1
---

* Add support for `SucceedInputCollection` and `TimeoutInputCollection` test helper methods on resource `Terminal.Reader`
* Add support for new value `setup_intent_mobile_wallet_unsupported` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Remove support for `Carrier`, `Phone`, and `TrackingNumber` on `CheckoutSessionCollectedInformationShippingDetails`
* Add support for `InterchangeFeesAmount`, `NetTotalAmount`, `NetworkFeesAmount`, `OtherFeesAmount`, `OtherFeesCount`, and `TransactionAmount` on `IssuingSettlement`
* Remove support for `InterchangeFees`, `NetTotal`, `NetworkFees`, and `TransactionVolume` on `IssuingSettlement`
* Add support for `TargetDate` on `OrderPaymentSettingsPaymentMethodOptionsAcssDebitParams`, `OrderPaymentSettingsPaymentMethodOptionsAcssDebit`, `OrderPaymentSettingsPaymentMethodOptionsSepaDebitParams`, and `OrderPaymentSettingsPaymentMethodOptionsSepaDebit`
* Add support for `ACHCreditTransfer`, `ACHDebit`, `ACSSDebit`, `AUBECSDebit`, `Affirm`, `AfterpayClearpay`, `Alipay`, `Alma`, `AmazonPay`, `BACSDebit`, `BLIK`, `Bancontact`, `Boleto`, `CardPresent`, `Card`, `CashApp`, `CustomerBalance`, `EPS`, `FPX`, `Giropay`, `Gopay`, `Grabpay`, `IDBankTransfer`, `IDEAL`, `InteracPresent`, `KakaoPay`, `Klarna`, `Konbini`, `KrCard`, `Link`, `MbWay`, `Mobilepay`, `Multibanco`, `NaverPay`, `OXXO`, `P24`, `PayByBank`, `PayNow`, `Payco`, `Paypal`, `Payto`, `Pix`, `PromptPay`, `Qris`, `Rechnung`, `RevolutPay`, `SEPACreditTransfer`, `SEPADebit`, `SamsungPay`, `Shopeepay`, `Sofort`, `StripeAccount`, `Swish`, `TWINT`, `USBankAccount`, `WeChatPay`, `WeChat`, and `Zip` on `PaymentAttemptRecordPaymentMethodDetails` and `PaymentRecordPaymentMethodDetails`
* Change type of `PaymentAttemptRecordPaymentMethodDetailsCustom` and `PaymentRecordPaymentMethodDetailsCustom` from `nullable(PaymentsPrimitivesPaymentRecordsResourcePaymentMethodDetailsResourceCustomDetails)` to `PaymentsPrimitivesPaymentRecordsResourcePaymentMethodCustomDetails`
* Change type of `PaymentAttemptRecordPaymentMethodDetailsType` and `PaymentRecordPaymentMethodDetailsType` from `literal('custom')` to `string`
* Add support for `Wifi` on `TerminalConfigurationParams` and `TerminalConfiguration`
