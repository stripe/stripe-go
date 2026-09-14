---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1760
is_stripe_api_change: true
released_in_version: 76.3.0
---

* Add support for new resource `Tax.Registration`
* Add support for `List`, `New`, and `Update` methods on resource `Registration`
* Add support for `RevolutPay` throughout the API
* Add support for new value `token_card_network_invalid` on enums `InvoiceLastFinalizationErrorCode`, `PaymentIntentLastPaymentErrorCode`, `SetupAttemptSetupErrorCode`, `SetupIntentLastSetupErrorCode`, and `StripeErrorCode`
* Add support for new value `payment_unreconciled` on enum `BalanceTransactionType`
* Add support for `ABA` and `Swift` on `FundingInstructionsBankTransferFinancialAddresses` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddresses`
* Add support for new values `ach`, `domestic_wire_us`, and `swift` on enums `FundingInstructionsBankTransferFinancialAddressesSupportedNetworks` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSupportedNetworks`
* Add support for new values `aba` and `swift` on enums `FundingInstructionsBankTransferFinancialAddressesType` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesType`
* Add support for `URL` on `IssuingAuthorizationMerchantDataParams`, `IssuingAuthorizationMerchantData`, `IssuingTransactionMerchantData`, `TestHelpersIssuingTransactionCreateForceCaptureMerchantDataParams`, and `TestHelpersIssuingTransactionCreateUnlinkedRefundMerchantDataParams`
* Add support for `AuthenticationExemption` and `ThreeDSecure` on `IssuingAuthorizationVerificationDataParams` and `IssuingAuthorizationVerificationData`
* Add support for `Description` on `PaymentLinkPaymentIntentDataParams` and `PaymentLinkPaymentIntentData`
