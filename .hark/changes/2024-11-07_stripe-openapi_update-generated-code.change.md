---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/1949
is_stripe_api_change: true
released_in_version: 81.1.0-beta.2
---

* Add support for new resources `Issuing.FraudLiabilityDebit`, `PaymentAttemptRecord`, and `PaymentRecord`
* Add support for `Get` and `List` methods on resources `FraudLiabilityDebit` and `PaymentAttemptRecord`
* Add support for `Get`, `ReportPaymentAttemptCanceled`, `ReportPaymentAttemptFailed`, `ReportPaymentAttemptGuaranteed`, `ReportPaymentAttempt`, and `ReportPayment` methods on resource `PaymentRecord`
* Change type of `AccountFutureRequirementsDisabledReason` and `AccountRequirementsDisabledReason` from `string` to `enum`
* Remove support for `MoneyMovement` on `AccountSessionComponentsFinancialAccountFeaturesParams`
* Add support for `CardManagement`, `CardSpendDisputeManagement`, `CardholderManagement`, and `SpendControlManagement` on `AccountSessionComponentsIssuingCardFeaturesParams`
* Add support for `DisableStripeUserAuthentication` on `AccountSessionComponentsIssuingCardsListFeaturesParams`
* Add support for `AdaptivePricing` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `MandateOptions` on `CheckoutSessionPaymentMethodOptionsBacsDebitParams`, `CheckoutSessionPaymentMethodOptionsBacsDebit`, `CheckoutSessionPaymentMethodOptionsSepaDebitParams`, and `CheckoutSessionPaymentMethodOptionsSepaDebit`
* Add support for `RequestDecrementalAuthorization`, `RequestExtendedAuthorization`, `RequestIncrementalAuthorization`, `RequestMulticapture`, and `RequestOvercapture` on `CheckoutSessionPaymentMethodOptionsCardParams` and `CheckoutSessionPaymentMethodOptionsCard`
* Add support for `CaptureMethod` on `CheckoutSessionPaymentMethodOptionsKakaoPayParams`, `CheckoutSessionPaymentMethodOptionsKrCardParams`, `CheckoutSessionPaymentMethodOptionsNaverPayParams`, `CheckoutSessionPaymentMethodOptionsPaycoParams`, and `CheckoutSessionPaymentMethodOptionsSamsungPayParams`
* Add support for new value `li_vat` on enums `CheckoutSessionCollectedInformationTaxIdsType`, `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `OrderTaxDetailsTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, `TaxIdType`, and `TaxTransactionCustomerDetailsTaxIdsType`
* Add support for new values `invoice.payment_attempt_required` and `issuing_fraud_liability_debit.created` on enum `EventType`
* Add support for `AccountHolderAddress`, `AccountHolderName`, `AccountType`, and `BankAddress` on `FundingInstructionsBankTransferFinancialAddressesAba`, `FundingInstructionsBankTransferFinancialAddressesSwift`, `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesAba`, and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSwift`
* Add support for `PaymentRecordData` and `PaymentRecord` on `InvoiceAttachPaymentParams`
* Remove support for `OutOfBandPayment` on `InvoiceAttachPaymentParams`
* Add support for `AmountOverpaid` on `Invoice`
* Add support for new value `custom` on enums `InvoicePaymentSettingsPaymentMethodTypes` and `SubscriptionPaymentSettingsPaymentMethodTypes`
* Add support for `MerchantAmount` and `MerchantCurrency` on `IssuingAuthorizationParams`
* Add support for new value `link` on enums `PaymentIntentPaymentMethodOptionsCardNetwork`, `SetupIntentPaymentMethodOptionsCardNetwork`, and `SubscriptionPaymentSettingsPaymentMethodOptionsCardNetwork`
* Add support for `SubmitType` on `PaymentLinkParams`
* Add support for new value `service_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationTaxBreakdownTaxRateDetailsTaxType`, `TaxRateTaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetailsTaxType`
