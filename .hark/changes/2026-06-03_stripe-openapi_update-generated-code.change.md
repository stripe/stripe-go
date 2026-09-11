---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2365
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.3.0-alpha.2
---

* Add support for new resources `DelegatedCheckoutOrderEvent`, `DelegatedCheckoutOrder`, `V2BillingContractLicensePricingQuantityChange`, `V2BillingContract`, and `V2SignalsAccountSignal`
* Add support for `Get` method on resource `DelegatedCheckoutOrder`
* Add support for `ListOrders` method on resource `DelegatedCheckoutRequestedSession`
* Add support for `Get` and `List` methods on resource `V2SignalsAccountSignal`
* Add support for `Activate`, `Cancel`, `Get`, `List`, `New`, and `Update` methods on resource `V2BillingContract`
* Add support for `BirthAddress` on `AccountIndividualParams`, `AccountPersonParams`, `Person`, `TokenAccountIndividualParams`, and `TokenPersonParams`
* Change type of `ChargeCapturePaymentDetailsMoneyServicesParams.TransactionType`, `ChargePaymentDetailsMoneyServicesParams.TransactionType`, `PaymentIntentCapturePaymentDetailsMoneyServicesParams.TransactionType`, `PaymentIntentConfirmPaymentDetailsMoneyServicesParams.TransactionType`, and `PaymentIntentPaymentDetailsMoneyServicesParams.TransactionType` from `literal('account_funding')` to `enum('account_funding'|'debt_repayment')`
* Add support for new value `proserv` on enums `CheckoutSessionAutomaticSurcharge.Provider` and `PaymentLinkAutomaticSurcharge.Provider`
* Add support for `ProvisioningDecision` and `TokenType` on `IssuingAuthorizationTokenDetails` and `IssuingToken`
* Add support for `TokenDecisionRecommendation` on `IssuingAuthorizationTokenDetailsNetworkDataVisa` and `IssuingTokenNetworkDataVisa`
* Add support for `Language` on `IssuingTokenNetworkDataDevice`
* Add support for `DigitalAssetCategory` on `PaymentIntentConfirmPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentConfirmPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`, `PaymentIntentPaymentMethodOptionsCardPaymentDetailsMoneyServicesAccountFundingParams`, and `PaymentIntentPaymentMethodOptionsCardPresentPaymentDetailsMoneyServicesAccountFundingParams`
* Add support for `StaticAddress` on `PaymentIntentConfirmPaymentMethodOptionsCryptoDepositOptionsParams`, `PaymentIntentPaymentMethodOptionsCryptoDepositOptionsParams`, and `PaymentIntentPaymentMethodOptionsCryptoDepositOptions`
* Add support for `PaymentReference` on `PaymentIntentPaymentsOrchestrationParams`
* ⚠️ Remove support for `PaymentDetails` on `PaymentIntentPaymentsOrchestrationParams`
* ⚠️ Change type of `PaymentIntentPaymentDetailsMoneyServices.TransactionType` from `literal('account_funding')` to `enum('account_funding'|'debt_repayment')`
* ⚠️ Add support for `EndingBefore`, `Limit`, and `StartingAfter` on `PaymentLocationListParams`
* Add support for `Schema` on `V2DataReportingQueryRunResultFile` and `V2ReportingReportRunResultFile`
* Add support for new value `payout_method_amount_limit_exceeded` on enum `V2MoneyManagementOutboundPaymentStatusDetailsFailed.Reason`
* Add support for `Include` on `V2DataReportingQueryRunParams` and `V2ReportingReportRunParams`
* Add support for `RequirementsCollector` on `V2CoreAccountDefaultsResponsibilitiesParams`
* Add support for event notification `V2SignalsAccountSignalMerchantDelinquencyReadyEvent` with related object `V2SignalsAccountSignal`
