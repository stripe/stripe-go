---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2396
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.2.0-alpha.5
---

* Add support for new resources `BillingAlertNotification` and `CryptoDepositAddress`
* Add support for `Get`, `List`, and `New` methods on resource `CryptoDepositAddress`
* Add support for `List` method on resource `BillingAlertNotification`
* Add support for new values `partner_disabled_dispute_rate`, `partner_disabled_responsibilities`, `partner_disabled_restricted_business`, and `partner_disabled_suspected_fraud` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
* Add support for new value `data_share_only` on enums `ChargePaymentMethodDetailsCardThreeDSecure.Result`, `PaymentAttemptRecordPaymentMethodDetailsCardThreeDSecure.Result`, `PaymentRecordPaymentMethodDetailsCardThreeDSecure.Result`, and `SetupAttemptPaymentMethodDetailsCardThreeDSecure.Result`
* Add support for `Vipps` on `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for new value `vipps` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for new value `sui` on enum `CryptoCustomerConsumerWallet.Network`
* Add support for new value `sui` on enum `CryptoOnrampSessionTransactionDetails.DestinationNetwork`
* Add support for new value `sui` on enum `CryptoOnrampSessionTransactionDetails.DestinationNetworks`
* Add support for `Sui` on `CryptoOnrampSessionTransactionDetailsWalletAddresses`
* Add support for `UseStripeSDK` on `DelegatedCheckoutRequestedSessionConfirmParams`
* Add support for new value `mb_way` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `EvCharging` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCard`, `PaymentIntentAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentCaptureAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentConfirmAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentDecrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsCardParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsCardParams`
* ⚠️ Change type of `PaymentIntent.AllowedPaymentMethodTypes` from `string` to `enum`
* Add support for new value `vipps` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for `TaxItems` on `PaymentIntentPaymentDetailsCarRentalDatumTotalTax`, `PaymentIntentPaymentDetailsFlightDatumTotalTax`, and `PaymentIntentPaymentDetailsLodgingDatumTotalTax`
* ⚠️ Remove support for `Taxes` on `PaymentIntentPaymentDetailsCarRentalDatumTotalTax`, `PaymentIntentPaymentDetailsFlightDatumTotalTax`, and `PaymentIntentPaymentDetailsLodgingDatumTotalTax`
* Add support for `Card` on `RadarPaymentEvaluationPaymentDetailsPaymentMethodDetailsParams`
* ⚠️ Remove support for `ACSSDebit`, `AUBECSDebit`, `AfterpayClearpay`, `Alipay`, `Alma`, `AmazonPay`, `BACSDebit`, `BLIK`, `Bancontact`, `Billie`, `Bizum`, `Boleto`, `CardPresent`, `CashApp`, `Crypto`, `CustomerBalance`, `EPS`, `FPX`, `GiftCard`, `Giropay`, `Gopay`, `Grabpay`, `IDBankTransfer`, `IDEAL`, `InteracPresent`, `KakaoPay`, `Konbini`, `KrCard`, `MbWay`, `Mobilepay`, `Multibanco`, `NaverPay`, `NzBankAccount`, `OXXO`, `P24`, `PayByBank`, `PayNow`, `Payco`, `Paypal`, `Paypay`, `Payto`, `Pix`, `PromptPay`, `Qris`, `Rechnung`, `RevolutPay`, `SEPADebit`, `SamsungPay`, `Satispay`, `Scalapay`, `Shopeepay`, `Sofort`, `StripeBalance`, `Sunbit`, `Swish`, `TWINT`, `Tamara`, `USBankAccount`, `Upi`, `WeChatPay`, and `Zip` on `SharedPaymentGrantedTokenPaymentMethodDetails`
* ⚠️ Add support for new value `shop_pay` on enum `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
* ⚠️ Remove support for values `acss_debit`, `afterpay_clearpay`, `alipay`, `alma`, `amazon_pay`, `au_becs_debit`, `bacs_debit`, `bancontact`, `billie`, `bizum`, `blik`, `boleto`, `card_present`, `cashapp`, `crypto`, `custom`, `customer_balance`, `eps`, `fpx`, `gift_card`, `giropay`, `gopay`, `grabpay`, `id_bank_transfer`, `ideal`, `interac_present`, `kakao_pay`, `konbini`, `kr_card`, `mb_way`, `mobilepay`, `multibanco`, `naver_pay`, `nz_bank_account`, `oxxo`, `p24`, `pay_by_bank`, `payco`, `paynow`, `paypal`, `paypay`, `payto`, `pix`, `promptpay`, `qris`, `rechnung`, `revolut_pay`, `samsung_pay`, `satispay`, `scalapay`, `sepa_debit`, `shopeepay`, `sofort`, `stripe_balance`, `sunbit`, `swish`, `tamara`, `twint`, `upi`, `us_bank_account`, `wechat_pay`, and `zip` from enum `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
* Add support for `SpendCard` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripeParams` and `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialStripe`
* Add support for new value `commercial.stripe.spend_card` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new value `payment_delinquency_exposure` on enum `V2SignalsAccountSignal.Type`
* Add support for new value `commercial.stripe.spend_card` on enum `EventsV2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEvent.UpdatedCapability`
