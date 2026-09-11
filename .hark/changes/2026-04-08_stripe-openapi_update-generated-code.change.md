---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2339
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.1.0-alpha.3
---

* Add support for `PaymentRecord` on `ApplicationFeeFeeSource`
* Add support for `FleetData` on `ChargeCapturePaymentDetailsParams`, `ChargePaymentDetailsParams`, `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsCard`, `PaymentIntentAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentCaptureAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentCapturePaymentDetailsParams`, `PaymentIntentConfirmAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentDecrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsCardParams`, `PaymentIntentPaymentDetailsParams`, and `PaymentIntentPaymentDetails`
* Add support for `BeneficiaryAccount`, `BeneficiaryDetails`, `SenderAccount`, and `SenderDetails` on `ChargeCapturePaymentDetailsMoneyServicesAccountFundingParams`, `ChargePaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentCapturePaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentConfirmPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentDetailsMoneyServicesAccountFundingParams`
* Change type of `ChargeCapturePaymentDetailsMoneyServicesParams.TransactionType`, `ChargePaymentDetailsMoneyServicesParams.TransactionType`, `PaymentIntentCapturePaymentDetailsMoneyServicesParams.TransactionType`, `PaymentIntentConfirmPaymentDetailsMoneyServicesParams.TransactionType`, and `PaymentIntentPaymentDetailsMoneyServicesParams.TransactionType` from `literal('account_funding')` to `emptyable(literal('account_funding'))`
* Add support for new value `requires_action` on enum `DelegatedCheckoutRequestedSession.Status`
* Add support for `Bizum` on `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for new value `bizum` on enums `InvoicePaymentSettings.PaymentMethodTypes`, `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`, and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `QuantityPrecision` on `PaymentIntentAmountDetailsLineItem`, `PaymentIntentAmountDetailsLineItemsParams`, `PaymentIntentCaptureAmountDetailsLineItemsParams`, `PaymentIntentConfirmAmountDetailsLineItemsParams`, `PaymentIntentDecrementAuthorizationAmountDetailsLineItemsParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsParams`
* Add support for `LiquidAsset` and `Wallet` on `PaymentIntentConfirmPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`
* Add support for `SharedPaymentGrantedToken` on `PaymentMethod`
* ⚠️ Change type of `RadarCustomerEvaluation.EventType` from `string` to `enum('login'|'registration')`
* ⚠️ Change type of `RadarCustomerEvaluationSignalsAccountSharing.RiskLevel` and `RadarCustomerEvaluationSignalsMultiAccounting.RiskLevel` from `string` to `enum`
* Add support for `Data` on `RadarPaymentEvaluationClientDeviceMetadataDetailsParams` and `RadarPaymentEvaluationClientDeviceMetadataDetails`
* Add support for `Sunbit` on `SharedPaymentGrantedTokenPaymentMethodDetails`
* Add support for new value `sunbit` on enum `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
* ⚠️ Remove support for values `bm_crn`, `bo_tin`, `bt_tpn`, `co_nit`, `ec_ruc`, `eg_tin`, `gh_tin`, `gy_tin`, `hn_rtn`, `jm_trn`, `jo_crn`, `ke_pin`, `ky_crn`, `lk_tin`, `mo_tin`, `mv_tin`, `ng_tin`, `pa_ruc`, `ph_tin`, `py_ruc`, `sl_tin`, `sv_nit`, `uy_ruc`, `vg_cn`, and `za_tin` from enum `V2CoreAccountIdentityBusinessDetailsIdNumber.Type`
* ⚠️ Remove support for values `bm_pp`, `bo_ci`, `bt_cid`, `eg_tin`, `gh_pin`, `gy_tin`, `hn_rtn`, `jm_trn`, `jo_pin`, `ky_pp`, `lk_nic`, `mo_bir`, `mt_nic`, `mv_tin`, `pa_ruc`, `ph_tin`, `py_ruc`, `si_pin`, `sv_nit`, and `vg_pp` from enums `V2CoreAccountIdentityIndividualIdNumber.Type` and `V2CoreAccountPersonIdNumber.Type`
* Add support for error type `CannotProceedError`
