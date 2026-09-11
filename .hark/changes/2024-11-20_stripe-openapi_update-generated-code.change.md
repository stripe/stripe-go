---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1951
is_stripe_api_change: true
released_in_version: 81.1.0
---

* Add support for `Respond` test helper method on resource `Issuing.Authorization`
* Add support for `Authorizer` on `AccountPersonsRelationshipParams` and `TokenPersonRelationshipParams`
* Change type of `AccountFutureRequirementsDisabledReason` and `AccountRequirementsDisabledReason` from `string` to `enum`
* Add support for `AdaptivePricing` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for `MandateOptions` on `CheckoutSessionPaymentMethodOptionsBacsDebitParams`, `CheckoutSessionPaymentMethodOptionsBacsDebit`, `CheckoutSessionPaymentMethodOptionsSepaDebitParams`, and `CheckoutSessionPaymentMethodOptionsSepaDebit`
* Add support for `RequestExtendedAuthorization`, `RequestIncrementalAuthorization`, `RequestMulticapture`, and `RequestOvercapture` on `CheckoutSessionPaymentMethodOptionsCardParams` and `CheckoutSessionPaymentMethodOptionsCard`
* Add support for `CaptureMethod` on `CheckoutSessionPaymentMethodOptionsKakaoPayParams`, `CheckoutSessionPaymentMethodOptionsKrCardParams`, `CheckoutSessionPaymentMethodOptionsNaverPayParams`, `CheckoutSessionPaymentMethodOptionsPaycoParams`, and `CheckoutSessionPaymentMethodOptionsSamsungPayParams`
* Add support for new value `li_vat` on enums `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, `TaxIdType`, and `TaxTransactionCustomerDetailsTaxIdsType`
* Add support for new value `subscribe` on enums `CheckoutSessionSubmitType` and `PaymentLinkSubmitType`
* Add support for new value `financial_account_statement` on enum `FilePurpose`
* Add support for `AccountHolderAddress`, `AccountHolderName`, `AccountType`, and `BankAddress` on `FundingInstructionsBankTransferFinancialAddressesAba`, `FundingInstructionsBankTransferFinancialAddressesSwift`, `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesAba`, and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSwift`
* Add support for `MerchantAmount` and `MerchantCurrency` on `IssuingAuthorizationParams`
* Add support for `FraudChallenges` and `VerifiedByFraudChallenge` on `IssuingAuthorization`
* Add support for new value `link` on enums `PaymentIntentPaymentMethodOptionsCardNetwork`, `SetupIntentPaymentMethodOptionsCardNetwork`, and `SubscriptionPaymentSettingsPaymentMethodOptionsCardNetwork`
* Add support for `SubmitType` on `PaymentLinkParams`
* Add support for `TraceID` on `Payout`
* Add support for `NetworkDeclineCode` on `RefundDestinationDetailsBlik` and `RefundDestinationDetailsSwish`
* Add support for new value `service_tax` on enums `TaxCalculationLineItemTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationShippingCostTaxBreakdownTaxRateDetailsTaxType`, `TaxCalculationTaxBreakdownTaxRateDetailsTaxType`, `TaxRateTaxType`, and `TaxTransactionShippingCostTaxBreakdownTaxRateDetailsTaxType`
