---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2462
semver_level: major
is_stripe_api_change: true
---

* Add support for new resources `RadarRule` and `V2MoneyManagementFundingSession`
* Add support for `New` method on resource `V2MoneyManagementFundingSession`
* ⚠️ Change type of `ChargeOutcome.Rule` from `RadarRule` to `$Radar.Rule`
* Add support for new value `ousd` on enums `ChargePaymentMethodDetailsCrypto.TokenCurrency`, `PaymentAttemptRecordPaymentMethodDetailsCrypto.TokenCurrency`, and `PaymentRecordPaymentMethodDetailsCrypto.TokenCurrency`
* Add support for `PaymentSettings` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `OnBehalfOf` on `CheckoutSession`
* Add support for new values `fednow` and `rtp` on enum `CustomerCashBalanceTransactionFundedBankTransferUsBankTransfer.Network`
* Add support for `Fuels` on `IssuingTransactionPurchaseDetails`
* Add support for `Fleet` on `PaymentIntentConfirmPaymentMethodOptionsCardPresentParams`, `PaymentIntentPaymentMethodOptionsCardPresentParams`, and `PaymentIntentPaymentMethodOptionsCardPresent`
* Add support for `SubscriptionReference` on `PaymentIntentConfirmPaymentMethodOptionsPaypayParams`, `PaymentIntentPaymentMethodOptionsPaypayParams`, and `PaymentIntentPaymentMethodOptionsPaypay`
* Add support for `USBankAccount` on `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsParams` and `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetails`
* Change type of `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetailsParams.MoneyMovementType` from `literal('card')` to `enum('card'|'us_bank_account')`
* Add support for `Rules` on `RadarPaymentEvaluation`
* ⚠️ Change type of `RadarPaymentEvaluationPaymentDetailsMoneyMovementDetails.MoneyMovementType` from `literal('card')` to `enum('card'|'us_bank_account')`
* Add support for new values `request_three_d_secure` and `reroute` on enum `RadarPaymentEvaluation.RecommendedAction`
* Add support for `BankInitiatedReturn` on `RadarPaymentEvaluationSignals`
* Add support for `Account` on `V2MoneyManagementFinancialAddressListParams`, `V2MoneyManagementFinancialAddressParams`, and `V2MoneyManagementFinancialAddress`
* Add support for new values `bre_b` and `pix` on enum `V2MoneyManagementFinancialAddressBankAccount.Type`
* Add support for `SupportedNetworkDetails` on `V2MoneyManagementFinancialAddressCryptoWallet`
* Add support for new value `bitcoin` on enums `V2MoneyManagementFinancialAddressCryptoWallet.Network` and `V2MoneyManagementReceivedCreditCryptoWalletTransferCryptoWallet.Network`
* Add support for `NetworkDetails` on `V2MoneyManagementInboundTransferParams` and `V2MoneyManagementInboundTransfer`
* Add support for `OriginatingCryptoWallet`, `TokenCurrency`, and `TransactionHash` on `V2MoneyManagementReceivedCreditCryptoWalletTransfer`
* Add support for `Customer` and `Subscription` on `EventsV1InvoiceUpcomingEvent`
