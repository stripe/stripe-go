---
title: API Updates
pr_url: https://github.com/stripe/stripe-go/pull/1471
is_stripe_api_change: true
released_in_version: 72.112.0
---

* Add support for `RadarOptions` on `ChargeParams`, `Charge`, `PaymentIntentConfirmParams`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentMethodParams`, `PaymentMethod`, `SetupIntentConfirmPaymentMethodDataParams`, and `SetupIntentPaymentMethodDataParams`
* Add support for `AccountHolderName`, `AccountNumber`, `AccountType`, `BankCode`, `BankName`, `BranchCode`, and `BranchName` on `FundingInstructionsBankTransferFinancialAddressesZengin` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesZengin`
* Change type of `OrderPaymentSettingsPaymentMethodOptionsCustomerBalanceBankTransferType` and `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferType` from `enum` to `literal('jp_bank_transfer')`
* Add support for `Network` on `SetupIntentPaymentMethodOptionsCard`
* Add support for new value `simulated_wisepos_e` on enum `TerminalReaderDeviceType`
